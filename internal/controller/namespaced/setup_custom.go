// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	publicnetwork "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/publicnetwork"
	backuplist "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/storage/backuplist"
)

// SetupCustom adds non-upjet controllers.
func SetupCustom(mgr ctrl.Manager, o tjcontroller.Options) error {
	for _, setup := range []func(ctrl.Manager, tjcontroller.Options) error{
		backuplist.Setup,
		publicnetwork.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupCustomGated adds non-upjet controllers with CRD gating.
func SetupCustomGated(mgr ctrl.Manager, o tjcontroller.Options) error {
	for _, setup := range []func(ctrl.Manager, tjcontroller.Options) error{
		backuplist.SetupGated,
		publicnetwork.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
