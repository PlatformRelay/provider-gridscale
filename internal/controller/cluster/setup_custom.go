// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"

	backuplist "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/storage/backuplist"
)

// SetupCustom adds non-upjet controllers.
func SetupCustom(mgr ctrl.Manager, o tjcontroller.Options) error {
	for _, setup := range []func(ctrl.Manager, tjcontroller.Options) error{
		backuplist.Setup,
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
