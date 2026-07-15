// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backupschedule "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/backupschedule"
	filesystem "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/filesystem"
	firewall "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/firewall"
	ipv4 "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/ipv4"
	ipv6 "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/ipv6"
	isoimage "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/isoimage"
	k8s "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/k8s"
	loadbalancer "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/loadbalancer"
	mariadb "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/mariadb"
	memcached "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/memcached"
	mysql "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/mysql"
	network "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/network"
	paas "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/paas"
	postgresql "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/postgresql"
	server "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/server"
	snapshot "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/snapshot"
	snapshotschedule "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/snapshotschedule"
	sqlserver "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/sqlserver"
	sshkey "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/sshkey"
	storage "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/storage"
	template "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/gridscale/template"
	application "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/marketplace/application"
	applicationimport "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/marketplace/applicationimport"
	mysql8 "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/mysql8/mysql8"
	storageaccesskey "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/object/storageaccesskey"
	storagebucket "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/object/storagebucket"
	securityzone "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/paas/securityzone"
	providerconfig "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/providerconfig"
	cache "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/redis/cache"
	store "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/redis/store"
	certificate "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/ssl/certificate"
	clone "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/storage/clone"
	storageimport "github.com/PlatformRelay/provider-gridscale/internal/controller/namespaced/storage/storageimport"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupschedule.Setup,
		filesystem.Setup,
		firewall.Setup,
		ipv4.Setup,
		ipv6.Setup,
		isoimage.Setup,
		k8s.Setup,
		loadbalancer.Setup,
		mariadb.Setup,
		memcached.Setup,
		mysql.Setup,
		network.Setup,
		paas.Setup,
		postgresql.Setup,
		server.Setup,
		snapshot.Setup,
		snapshotschedule.Setup,
		sqlserver.Setup,
		sshkey.Setup,
		storage.Setup,
		template.Setup,
		application.Setup,
		applicationimport.Setup,
		mysql8.Setup,
		storageaccesskey.Setup,
		storagebucket.Setup,
		securityzone.Setup,
		providerconfig.Setup,
		cache.Setup,
		store.Setup,
		certificate.Setup,
		clone.Setup,
		storageimport.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		backupschedule.SetupGated,
		filesystem.SetupGated,
		firewall.SetupGated,
		ipv4.SetupGated,
		ipv6.SetupGated,
		isoimage.SetupGated,
		k8s.SetupGated,
		loadbalancer.SetupGated,
		mariadb.SetupGated,
		memcached.SetupGated,
		mysql.SetupGated,
		network.SetupGated,
		paas.SetupGated,
		postgresql.SetupGated,
		server.SetupGated,
		snapshot.SetupGated,
		snapshotschedule.SetupGated,
		sqlserver.SetupGated,
		sshkey.SetupGated,
		storage.SetupGated,
		template.SetupGated,
		application.SetupGated,
		applicationimport.SetupGated,
		mysql8.SetupGated,
		storageaccesskey.SetupGated,
		storagebucket.SetupGated,
		securityzone.SetupGated,
		providerconfig.SetupGated,
		cache.SetupGated,
		store.SetupGated,
		certificate.SetupGated,
		clone.SetupGated,
		storageimport.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
