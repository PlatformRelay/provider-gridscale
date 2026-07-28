// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package publicnetwork

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1alpha1 "github.com/PlatformRelay/provider-gridscale/apis/cluster/gridscale/v1alpha1"
	"github.com/PlatformRelay/provider-gridscale/internal/clients"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newTestExternal(t *testing.T, handler http.HandlerFunc) (*external, func()) {
	t.Helper()
	srv := httptest.NewServer(handler)
	c := clients.NewGridscaleClient("uuid", "token", srv.URL)
	return &external{client: c}, srv.Close
}

func TestObserve_PublicNetworkFound(t *testing.T) {
	resp := networksResponse{
		Networks: map[string]networkEntry{
			"abc": {
				Name:         "public-net",
				NetworkType:  "public",
				Status:       "active",
				LocationUUID: "loc-uuid",
			},
		},
	}
	ext, cleanup := newTestExternal(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(resp)
	})
	defer cleanup()

	pn := &v1alpha1.PublicNetwork{ObjectMeta: metav1.ObjectMeta{Name: "test"}}
	obs, err := ext.Observe(context.Background(), pn)
	if err != nil {
		t.Fatalf("Observe error: %v", err)
	}
	if !obs.ResourceExists {
		t.Error("ResourceExists should be true")
	}
	if pn.Status.AtProvider.Name == nil || *pn.Status.AtProvider.Name != "public-net" {
		t.Errorf("unexpected name: %v", pn.Status.AtProvider.Name)
	}
	if pn.Status.AtProvider.NetworkType == nil || *pn.Status.AtProvider.NetworkType != "public" {
		t.Errorf("unexpected network type: %v", pn.Status.AtProvider.NetworkType)
	}
}

func TestObserve_OnlyPrivateNetworks(t *testing.T) {
	resp := networksResponse{
		Networks: map[string]networkEntry{
			"xyz": {Name: "private-net", NetworkType: "private"},
		},
	}
	ext, cleanup := newTestExternal(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(resp)
	})
	defer cleanup()

	pn := &v1alpha1.PublicNetwork{ObjectMeta: metav1.ObjectMeta{Name: "test"}}
	obs, err := ext.Observe(context.Background(), pn)
	if err != nil {
		t.Fatalf("Observe error: %v", err)
	}
	if obs.ResourceExists {
		t.Error("ResourceExists should be false when no public network found")
	}
}

func TestObserve_EmptyNetworkList(t *testing.T) {
	resp := networksResponse{Networks: map[string]networkEntry{}}
	ext, cleanup := newTestExternal(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(resp)
	})
	defer cleanup()

	pn := &v1alpha1.PublicNetwork{ObjectMeta: metav1.ObjectMeta{Name: "test"}}
	obs, err := ext.Observe(context.Background(), pn)
	if err != nil {
		t.Fatalf("Observe error: %v", err)
	}
	if obs.ResourceExists {
		t.Error("ResourceExists should be false for empty network list")
	}
}

func TestCreate_ReturnsError(t *testing.T) {
	ext := &external{client: clients.NewGridscaleClient("u", "t", "http://localhost:1")}
	_, err := ext.Create(context.Background(), &v1alpha1.PublicNetwork{})
	if err == nil {
		t.Fatal("Create should return observe-only error")
	}
}

func TestUpdate_ReturnsError(t *testing.T) {
	ext := &external{client: clients.NewGridscaleClient("u", "t", "http://localhost:1")}
	_, err := ext.Update(context.Background(), &v1alpha1.PublicNetwork{})
	if err == nil {
		t.Fatal("Update should return observe-only error")
	}
}
