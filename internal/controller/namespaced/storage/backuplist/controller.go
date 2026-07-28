// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package backuplist

import (
	"context"
	"fmt"
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

	v1alpha1 "github.com/PlatformRelay/provider-gridscale/apis/namespaced/storage/v1alpha1"
	"github.com/PlatformRelay/provider-gridscale/internal/clients"
)

const (
	errNotBackupList = "managed resource is not a BackupList"
	errList          = "cannot list storage backups"
)

type connector struct {
	kube    client.Client
	setupFn terraform.SetupFn
}

func (c *connector) Connect(ctx context.Context, mg xpresource.Managed) (managed.ExternalClient, error) {
	bl, ok := mg.(*v1alpha1.BackupList)
	if !ok {
		return nil, errors.New(errNotBackupList)
	}
	setup, err := c.setupFn(ctx, c.kube, bl)
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

// backupListResponse models the gridscale API response for storage backups.
type backupListResponse struct {
	StorageBackups []struct {
		ObjectUUID string  `json:"object_uuid"`
		Name       string  `json:"name"`
		Capacity   float64 `json:"capacity"`
		CreateTime string  `json:"create_time"`
	} `json:"storage_backups"`
}

func (e *external) Observe(ctx context.Context, mg xpresource.Managed) (managed.ExternalObservation, error) {
	bl, ok := mg.(*v1alpha1.BackupList)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotBackupList)
	}

	var resp backupListResponse
	path := fmt.Sprintf("/objects/storages/%s/backups", bl.Spec.ForProvider.StorageUUID)
	if err := e.client.Get(ctx, path, &resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, errList)
	}

	entries := make([]v1alpha1.StorageBackupEntry, 0, len(resp.StorageBackups))
	for _, b := range resp.StorageBackups {
		entries = append(entries, v1alpha1.StorageBackupEntry{
			ObjectUUID: b.ObjectUUID,
			Name:       b.Name,
			Capacity:   b.Capacity,
			CreateTime: b.CreateTime,
		})
	}
	bl.Status.AtProvider.StorageBackups = entries
	bl.SetConditions(xpv1.Available())

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func (e *external) Create(_ context.Context, _ xpresource.Managed) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, errors.New("BackupList is observe-only and does not support Create")
}

func (e *external) Update(_ context.Context, _ xpresource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, errors.New("BackupList is observe-only and does not support Update")
}

func (e *external) Delete(_ context.Context, _ xpresource.Managed) (managed.ExternalDelete, error) {
	// observe-only: nothing to delete remotely.
	return managed.ExternalDelete{}, nil
}

func (e *external) Disconnect(_ context.Context) error {
	return nil
}

// Setup adds a controller that reconciles BackupList resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.BackupList_GroupVersionKind.String())
	opts := []managed.ReconcilerOption{
		managed.WithExternalConnecter(&connector{
			kube:    mgr.GetClient(),
			setupFn: o.SetupFn,
		}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithTimeout(1 * time.Minute),
		managed.WithPollInterval(o.PollInterval),
		managed.WithManagementPolicies(),
	}
	if o.PollJitter != 0 {
		opts = append(opts, managed.WithPollJitterHook(o.PollJitter))
	}

	r := managed.NewReconciler(mgr, xpresource.ManagedKind(v1alpha1.BackupList_GroupVersionKind), opts...)
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(xpresource.DesiredStateChanged()).
		For(&v1alpha1.BackupList{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// SetupGated registers the controller behind a CRD-existence gate.
func SetupGated(mgr ctrl.Manager, o tjcontroller.Options) error {
	o.Options.Gate.Register(func() {
		if err := Setup(mgr, o); err != nil {
			mgr.GetLogger().Error(err, "unable to setup reconciler", "gvk", v1alpha1.BackupList_GroupVersionKind.String())
		}
	}, v1alpha1.BackupList_GroupVersionKind)
	return nil
}
