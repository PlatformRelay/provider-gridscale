// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package publicnetwork

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache/informertest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	ctrlconfig "sigs.k8s.io/controller-runtime/pkg/config"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/terraform"

	v1alpha1 "github.com/PlatformRelay/provider-gridscale/apis/namespaced/gridscale/v1alpha1"
	otherv1alpha1 "github.com/PlatformRelay/provider-gridscale/apis/namespaced/storage/v1alpha1"
	"github.com/PlatformRelay/provider-gridscale/internal/clients"
)

// recordingGate captures what SetupGated registers so the gating contract can be
// asserted rather than merely executed.
type recordingGate struct {
	callback func()
	gvks     []schema.GroupVersionKind
}

func (g *recordingGate) Register(callback func(), gvks ...schema.GroupVersionKind) {
	g.callback = callback
	g.gvks = gvks
}

func (g *recordingGate) Set(schema.GroupVersionKind, bool) bool { return false }

// newFakeManager builds a controller-runtime Manager backed by a fake client and
// fake informers. It is never started, so the unreachable rest.Config host is
// never dialled — it exists only so Setup() can register a controller against it.
func newFakeManager(t *testing.T) ctrl.Manager {
	t.Helper()
	s := runtime.NewScheme()
	if err := v1alpha1.AddToScheme(s); err != nil {
		t.Fatalf("add scheme: %v", err)
	}
	fc := fake.NewClientBuilder().WithScheme(s).Build()
	// Controller-name uniqueness is enforced process-wide, so several tests each
	// setting up the same GVK would collide. These managers are never started.
	skipNameValidation := true
	mgr, err := ctrl.NewManager(&rest.Config{Host: "http://127.0.0.1:1"}, ctrl.Options{
		Scheme:                 s,
		Controller:             ctrlconfig.Controller{SkipNameValidation: &skipNameValidation},
		Metrics:                metricsserver.Options{BindAddress: "0"},
		HealthProbeBindAddress: "0",
		MapperProvider: func(*rest.Config, *http.Client) (meta.RESTMapper, error) {
			return meta.NewDefaultRESTMapper(nil), nil
		},
		NewClient: func(*rest.Config, client.Options) (client.Client, error) { return fc, nil },
		NewCache: func(*rest.Config, cache.Options) (cache.Cache, error) {
			return &informertest.FakeInformers{Scheme: s}, nil
		},
	})
	if err != nil {
		t.Fatalf("new manager: %v", err)
	}
	return mgr
}

func testOptions(gate xpcontroller.Gate) tjcontroller.Options {
	return tjcontroller.Options{Options: xpcontroller.Options{
		Logger:            logging.NewNopLogger(),
		GlobalRateLimiter: ratelimiter.NewGlobal(1),
		PollInterval:      time.Minute,
		Gate:              gate,
	}}
}

// externalFor wires an external against a test HTTP server standing in for the
// gridscale API.
func externalFor(t *testing.T, handler http.HandlerFunc) *external {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return &external{client: clients.NewGridscaleClient("uuid", "token", srv.URL)}
}

func TestConnect_RejectsWrongManagedType(t *testing.T) {
	c := &connector{setupFn: func(context.Context, client.Client, xpresource.Managed) (terraform.Setup, error) {
		t.Error("setupFn must not be called for a mismatched managed type")
		return terraform.Setup{}, nil
	}}

	_, err := c.Connect(context.Background(), &otherv1alpha1.BackupList{})
	if err == nil {
		t.Fatal("Connect must reject a managed resource of the wrong type")
	}
	if !strings.Contains(err.Error(), errNotPublicNetwork) {
		t.Errorf("error = %q, want it to contain %q", err, errNotPublicNetwork)
	}
}

func TestConnect_PropagatesSetupError(t *testing.T) {
	sentinel := errors.New("provider setup exploded")
	c := &connector{setupFn: func(context.Context, client.Client, xpresource.Managed) (terraform.Setup, error) {
		return terraform.Setup{}, sentinel
	}}

	_, err := c.Connect(context.Background(), &v1alpha1.PublicNetwork{})
	if err == nil {
		t.Fatal("Connect must surface the setupFn error")
	}
	if !errors.Is(err, sentinel) {
		t.Errorf("error = %v, want it to wrap %v", err, sentinel)
	}
}

// TestConnect_BuildsClientFromProviderSetup asserts that Connect actually
// threads uuid/token/api_url out of the terraform Setup and into the API client,
// by observing the request the resulting client makes.
func TestConnect_BuildsClientFromProviderSetup(t *testing.T) {
	var gotUUID, gotToken string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUUID = r.Header.Get("X-Auth-UserID")
		gotToken = r.Header.Get("X-Auth-Token")
		json.NewEncoder(w).Encode(networksResponse{}) //nolint:errcheck
	}))
	defer srv.Close()

	c := &connector{setupFn: func(context.Context, client.Client, xpresource.Managed) (terraform.Setup, error) {
		return terraform.Setup{Configuration: terraform.ProviderConfiguration{
			"uuid":    "user-uuid",
			"token":   "secret-token",
			"api_url": srv.URL,
		}}, nil
	}}

	ext, err := c.Connect(context.Background(), &v1alpha1.PublicNetwork{})
	if err != nil {
		t.Fatalf("Connect: %v", err)
	}
	if _, err := ext.Observe(context.Background(), &v1alpha1.PublicNetwork{}); err != nil {
		t.Fatalf("Observe through connected client: %v", err)
	}
	if gotUUID != "user-uuid" {
		t.Errorf("X-Auth-UserID = %q, want %q", gotUUID, "user-uuid")
	}
	if gotToken != "secret-token" {
		t.Errorf("X-Auth-Token = %q, want %q", gotToken, "secret-token")
	}
}

func TestObserve_RejectsWrongManagedType(t *testing.T) {
	ext := externalFor(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("no API call should be made for a mismatched managed type")
	})

	_, err := ext.Observe(context.Background(), &otherv1alpha1.BackupList{})
	if err == nil {
		t.Fatal("Observe must reject a managed resource of the wrong type")
	}
	if !strings.Contains(err.Error(), errNotPublicNetwork) {
		t.Errorf("error = %q, want it to contain %q", err, errNotPublicNetwork)
	}
}

func TestObserve_SurfacesAPIError(t *testing.T) {
	ext := externalFor(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	_, err := ext.Observe(context.Background(), &v1alpha1.PublicNetwork{})
	if err == nil {
		t.Fatal("Observe must surface a non-2xx API response as an error")
	}
	if !strings.Contains(err.Error(), errGetNetworks) {
		t.Errorf("error = %q, want it to contain %q", err, errGetNetworks)
	}
}

func TestDelete_IsNoOp(t *testing.T) {
	ext := &external{}
	if _, err := ext.Delete(context.Background(), &v1alpha1.PublicNetwork{}); err != nil {
		t.Errorf("Delete on an observe-only resource must be a no-op, got %v", err)
	}
}

func TestDisconnect_IsNoOp(t *testing.T) {
	ext := &external{}
	if err := ext.Disconnect(context.Background()); err != nil {
		t.Errorf("Disconnect must be a no-op, got %v", err)
	}
}

func TestSetup_RegistersControllerWithManager(t *testing.T) {
	if err := Setup(newFakeManager(t), testOptions(nil)); err != nil {
		t.Fatalf("Setup: %v", err)
	}
}

func TestSetup_HonoursPollJitter(t *testing.T) {
	o := testOptions(nil)
	o.PollJitter = time.Second
	if err := Setup(newFakeManager(t), o); err != nil {
		t.Fatalf("Setup with PollJitter: %v", err)
	}
}

func TestSetupGated_RegistersGVKAndDefersSetup(t *testing.T) {
	mgr := newFakeManager(t)
	g := &recordingGate{}

	if err := SetupGated(mgr, testOptions(g)); err != nil {
		t.Fatalf("SetupGated: %v", err)
	}
	if g.callback == nil {
		t.Fatal("SetupGated must register a callback with the gate")
	}
	want := v1alpha1.PublicNetwork_GroupVersionKind
	if len(g.gvks) != 1 || g.gvks[0] != want {
		t.Fatalf("registered gvks = %v, want [%v]", g.gvks, want)
	}

	// The gate fires the callback once the CRD exists; it must set the
	// controller up against the manager without panicking.
	g.callback()
}
