// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	backupschedule "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/backupschedule"
	filesystem "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/filesystem"
	firewall "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/firewall"
	ipv4 "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/ipv4"
	ipv6 "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/ipv6"
	isoimage "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/isoimage"
	k8s "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/k8s"
	loadbalancer "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/loadbalancer"
	mariadb "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/mariadb"
	memcached "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/memcached"
	mysql "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/mysql"
	network "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/network"
	paas "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/paas"
	postgresql "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/postgresql"
	server "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/server"
	snapshot "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/snapshot"
	snapshotschedule "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/snapshotschedule"
	sqlserver "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/sqlserver"
	sshkey "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/sshkey"
	storage "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/storage"
	template "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/gridscale/template"
	application "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/marketplace/application"
	applicationimport "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/marketplace/applicationimport"
	mysql8 "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/mysql8/mysql8"
	storageaccesskey "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/object/storageaccesskey"
	storagebucket "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/object/storagebucket"
	securityzone "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/paas/securityzone"
	providerconfig "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/providerconfig"
	cache "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/redis/cache"
	store "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/redis/store"
	certificate "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/ssl/certificate"
	clone "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/storage/clone"
	storageimport "github.com/PlatformRelay/provider-gridscale/internal/controller/cluster/storage/storageimport"
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
