// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package publicnetwork

import (
	"context"
	"time"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/terraform"

	v1alpha1 "github.com/PlatformRelay/provider-gridscale/apis/namespaced/gridscale/v1alpha1"
	"github.com/PlatformRelay/provider-gridscale/internal/clients"
)

const (
	errNotPublicNetwork = "managed resource is not a PublicNetwork"
	errGetNetworks      = "cannot get networks"
)

type connector struct {
	kube    client.Client
	setupFn terraform.SetupFn
}

func (c *connector) Connect(ctx context.Context, mg xpresource.Managed) (managed.ExternalClient, error) {
	pn, ok := mg.(*v1alpha1.PublicNetwork)
	if !ok {
		return nil, errors.New(errNotPublicNetwork)
	}
	setup, err := c.setupFn(ctx, c.kube, pn)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get provider setup")
	}
	userUUID, _ := setup.Configuration["uuid"].(string)
	token, _ := setup.Configuration["token"].(string)
	apiURL, _ := setup.Configuration["api_url"].(string)
	return &external{
		client: clients.NewGridscaleClient(userUUID, token, apiURL),
	}, nil
}

type external struct {
	client *clients.GridscaleClient
}

// networkEntry models a single network in the gridscale API response.
type networkEntry struct {
	Name            string   `json:"name"`
	Status          string   `json:"status"`
	NetworkType     string   `json:"network_type"`
	LocationUUID    string   `json:"location_uuid"`
	LocationName    string   `json:"location_name"`
	LocationCountry string   `json:"location_country"`
	LocationIATA    string   `json:"location_iata"`
	L2Security      bool     `json:"l2security"`
	DeleteBlock     bool     `json:"delete_block"`
	Labels          []string `json:"labels"`
	CreateTime      string   `json:"create_time"`
	ChangeTime      string   `json:"change_time"`
}

// networksResponse models the gridscale API response for GET /objects/networks.
type networksResponse struct {
	Networks map[string]networkEntry `json:"networks"`
}

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }

func (e *external) Observe(ctx context.Context, mg xpresource.Managed) (managed.ExternalObservation, error) {
	pn, ok := mg.(*v1alpha1.PublicNetwork)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotPublicNetwork)
	}

	var resp networksResponse
	if err := e.client.Get(ctx, "/objects/networks", &resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errGetNetworks)
	}

	for _, net := range resp.Networks {
		if net.NetworkType != "public" {
			continue
		}
		obs := &pn.Status.AtProvider
		obs.Name = strPtr(net.Name)
		obs.Status = strPtr(net.Status)
		obs.NetworkType = strPtr(net.NetworkType)
		obs.LocationUUID = strPtr(net.LocationUUID)
		obs.LocationName = strPtr(net.LocationName)
		obs.LocationCountry = strPtr(net.LocationCountry)
		obs.LocationIATA = strPtr(net.LocationIATA)
		obs.L2Security = boolPtr(net.L2Security)
		obs.DeleteBlock = boolPtr(net.DeleteBlock)
		obs.Labels = net.Labels
		obs.CreateTime = strPtr(net.CreateTime)
		obs.ChangeTime = strPtr(net.ChangeTime)
		pn.SetConditions(xpv1.Available())
		return managed.ExternalObservation{
			ResourceExists:   true,
			ResourceUpToDate: true,
		}, nil
	}

	return managed.ExternalObservation{ResourceExists: false}, nil
}

func (e *external) Create(_ context.Context, _ xpresource.Managed) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, errors.New("PublicNetwork is observe-only and does not support Create")
}

func (e *external) Update(_ context.Context, _ xpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, errors.New("PublicNetwork is observe-only and does not support Update")
}

func (e *external) Delete(_ context.Context, _ xpresource.Managed) (managed.ExternalDelete, error) {
	// observe-only: nothing to delete remotely.
	return managed.ExternalDelete{}, nil
}

func (e *external) Disconnect(_ context.Context) error {
	return nil
}

// Setup adds a controller that reconciles PublicNetwork resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.PublicNetwork_GroupVersionKind.String())
	opts := []managed.ReconcilerOption{
		managed.WithExternalConnector(&connector{
			kube:    mgr.GetClient(),
			setupFn: o.SetupFn,
		}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))), //nolint:staticcheck // suppress until crossplane-runtime offers the new recorder api
		managed.WithTimeout(1 * time.Minute),
		managed.WithPollInterval(o.PollInterval),
		managed.WithManagementPolicies(),
	}
	if o.PollJitter != 0 {
		opts = append(opts, managed.WithPollJitterHook(o.PollJitter))
	}

	r := managed.NewReconciler(mgr, xpresource.ManagedKind(v1alpha1.PublicNetwork_GroupVersionKind), opts...)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(xpresource.DesiredStateChanged()).
		For(&v1alpha1.PublicNetwork{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// SetupGated registers the controller behind a CRD-existence gate.
func SetupGated(mgr ctrl.Manager, o tjcontroller.Options) error {
	o.Gate.Register(func() {
		if err := Setup(mgr, o); err != nil {
			mgr.GetLogger().Error(err, "unable to setup reconciler", "gvk", v1alpha1.PublicNetwork_GroupVersionKind.String())
		}
	}, v1alpha1.PublicNetwork_GroupVersionKind)
	return nil
}
