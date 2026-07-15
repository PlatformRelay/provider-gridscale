# API Reference

## Packages
- [gridscale.gridscale.m.platformrelay.io/v1alpha1](#gridscalegridscalemplatformrelayiov1alpha1)
- [gridscale.gridscale.platformrelay.io/v1alpha1](#gridscalegridscaleplatformrelayiov1alpha1)
- [gridscale.m.platformrelay.io/v1alpha1](#gridscalemplatformrelayiov1alpha1)
- [gridscale.m.platformrelay.io/v1beta1](#gridscalemplatformrelayiov1beta1)
- [gridscale.platformrelay.io/v1alpha1](#gridscaleplatformrelayiov1alpha1)
- [gridscale.platformrelay.io/v1beta1](#gridscaleplatformrelayiov1beta1)
- [marketplace.gridscale.m.platformrelay.io/v1alpha1](#marketplacegridscalemplatformrelayiov1alpha1)
- [marketplace.gridscale.platformrelay.io/v1alpha1](#marketplacegridscaleplatformrelayiov1alpha1)
- [mysql8.gridscale.m.platformrelay.io/v1alpha1](#mysql8gridscalemplatformrelayiov1alpha1)
- [mysql8.gridscale.platformrelay.io/v1alpha1](#mysql8gridscaleplatformrelayiov1alpha1)
- [object.gridscale.m.platformrelay.io/v1alpha1](#objectgridscalemplatformrelayiov1alpha1)
- [object.gridscale.platformrelay.io/v1alpha1](#objectgridscaleplatformrelayiov1alpha1)
- [paas.gridscale.m.platformrelay.io/v1alpha1](#paasgridscalemplatformrelayiov1alpha1)
- [paas.gridscale.platformrelay.io/v1alpha1](#paasgridscaleplatformrelayiov1alpha1)
- [redis.gridscale.m.platformrelay.io/v1alpha1](#redisgridscalemplatformrelayiov1alpha1)
- [redis.gridscale.platformrelay.io/v1alpha1](#redisgridscaleplatformrelayiov1alpha1)
- [ssl.gridscale.m.platformrelay.io/v1alpha1](#sslgridscalemplatformrelayiov1alpha1)
- [ssl.gridscale.platformrelay.io/v1alpha1](#sslgridscaleplatformrelayiov1alpha1)
- [storage.gridscale.m.platformrelay.io/v1alpha1](#storagegridscalemplatformrelayiov1alpha1)
- [storage.gridscale.platformrelay.io/v1alpha1](#storagegridscaleplatformrelayiov1alpha1)


## gridscale.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [Backupschedule](#backupschedule)
- [BackupscheduleList](#backupschedulelist)
- [Filesystem](#filesystem)
- [FilesystemList](#filesystemlist)
- [Firewall](#firewall)
- [FirewallList](#firewalllist)
- [IPv4](#ipv4)
- [IPv4List](#ipv4list)
- [IPv6](#ipv6)
- [IPv6List](#ipv6list)
- [Isoimage](#isoimage)
- [IsoimageList](#isoimagelist)
- [K8S](#k8s)
- [K8SList](#k8slist)
- [Loadbalancer](#loadbalancer)
- [LoadbalancerList](#loadbalancerlist)
- [Mariadb](#mariadb)
- [MariadbList](#mariadblist)
- [Memcached](#memcached)
- [MemcachedList](#memcachedlist)
- [MySQL](#mysql)
- [MySQLList](#mysqllist)
- [Network](#network)
- [NetworkList](#networklist)
- [Paas](#paas)
- [PaasList](#paaslist)
- [Postgresql](#postgresql)
- [PostgresqlList](#postgresqllist)
- [Server](#server)
- [ServerList](#serverlist)
- [Snapshot](#snapshot)
- [SnapshotList](#snapshotlist)
- [Snapshotschedule](#snapshotschedule)
- [SnapshotscheduleList](#snapshotschedulelist)
- [Sqlserver](#sqlserver)
- [SqlserverList](#sqlserverlist)
- [Sshkey](#sshkey)
- [SshkeyList](#sshkeylist)
- [Storage](#storage)
- [StorageList](#storagelist)
- [Template](#template)
- [TemplateList](#templatelist)





#### AutoAssignedServersObservation







_Appears in:_
- [NetworkObservation_2](#networkobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `ip` _string_ | IP which is assigned to the server. |  |  |
| `serverUuid` _string_ | UUID of the server. |  |  |




#### BackendServerInitParameters







_Appears in:_
- [LoadbalancerInitParameters](#loadbalancerinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | A valid domain or an IP address of a server. |  |  |
| `proxyProtocol` _string_ | The proxy protocol version. The proxy protocol is disabled by default and the valid version is either v1 or v2. |  |  |
| `weight` _float_ | The backend host weight. Default: 100. |  |  |


#### BackendServerObservation







_Appears in:_
- [LoadbalancerObservation](#loadbalancerobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | A valid domain or an IP address of a server. |  |  |
| `proxyProtocol` _string_ | The proxy protocol version. The proxy protocol is disabled by default and the valid version is either v1 or v2. |  |  |
| `weight` _float_ | The backend host weight. Default: 100. |  |  |


#### BackendServerParameters







_Appears in:_
- [LoadbalancerParameters](#loadbalancerparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | A valid domain or an IP address of a server. |  | Optional: \{\} <br /> |
| `proxyProtocol` _string_ | The proxy protocol version. The proxy protocol is disabled by default and the valid version is either v1 or v2. |  | Optional: \{\} <br /> |
| `weight` _float_ | The backend host weight. Default: 100. |  | Optional: \{\} <br /> |


#### Backupschedule



Backupschedule is the Schema for the Backupschedules API. <no value>



_Appears in:_
- [BackupscheduleList](#backupschedulelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Backupschedule` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[BackupscheduleSpec](#backupschedulespec)_ |  |  |  |
| `status` _[BackupscheduleStatus](#backupschedulestatus)_ |  |  |  |


#### BackupscheduleInitParameters







_Appears in:_
- [BackupscheduleSpec](#backupschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `active` _boolean_ | The status of the schedule active or not |  |  |
| `backupLocationUuid` _string_ | UUID of the location where your backup is stored. |  |  |
| `keepBackups` _float_ | The amount of storage backups to keep before overwriting the last created backup |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the storage backup schedule will be run. Format: "2006-01-02 15:04:05" |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `storageUuid` _string_ | UUID of the storage used to create storage backups |  |  |


#### BackupscheduleList



BackupscheduleList contains a list of Backupschedules





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `BackupscheduleList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Backupschedule](#backupschedule) array_ |  |  |  |


#### BackupscheduleObservation







_Appears in:_
- [BackupscheduleStatus](#backupschedulestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `active` _boolean_ | The status of the schedule active or not |  |  |
| `backupLocationName` _string_ | The human-readable name of backup location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `backupLocationUuid` _string_ | UUID of the location where your backup is stored. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `keepBackups` _float_ | The amount of storage backups to keep before overwriting the last created backup |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the storage backup schedule will be run. Format: "2006-01-02 15:04:05" |  |  |
| `nextRuntimeComputed` _string_ | The date and time that the storage backup schedule will be run. This date and time is computed by gridscale's server. |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `storageBackups` _[StorageBackupsObservation](#storagebackupsobservation) array_ | Related backups |  |  |
| `storageUuid` _string_ | UUID of the storage used to create storage backups |  |  |


#### BackupscheduleParameters







_Appears in:_
- [BackupscheduleSpec](#backupschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `active` _boolean_ | The status of the schedule active or not |  | Optional: \{\} <br /> |
| `backupLocationUuid` _string_ | UUID of the location where your backup is stored. |  | Optional: \{\} <br /> |
| `keepBackups` _float_ | The amount of storage backups to keep before overwriting the last created backup |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object |  | Optional: \{\} <br /> |
| `nextRuntime` _string_ | The date and time that the storage backup schedule will be run. Format: "2006-01-02 15:04:05" |  | Optional: \{\} <br /> |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  | Optional: \{\} <br /> |
| `storageUuid` _string_ | UUID of the storage used to create storage backups |  | Optional: \{\} <br /> |


#### BackupscheduleSpec



BackupscheduleSpec defines the desired state of Backupschedule



_Appears in:_
- [Backupschedule](#backupschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[BackupscheduleParameters](#backupscheduleparameters)_ |  |  |  |
| `initProvider` _[BackupscheduleInitParameters](#backupscheduleinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### BackupscheduleStatus



BackupscheduleStatus defines the observed state of Backupschedule.



_Appears in:_
- [Backupschedule](#backupschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[BackupscheduleObservation](#backupscheduleobservation)_ |  |  |  |


#### Filesystem



Filesystem is the Schema for the Filesystems API. <no value>



_Appears in:_
- [FilesystemList](#filesystemlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Filesystem` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[FilesystemSpec](#filesystemspec)_ |  |  |  |
| `status` _[FilesystemStatus](#filesystemstatus)_ |  |  |  |


#### FilesystemInitParameters







_Appears in:_
- [FilesystemSpec](#filesystemspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowedIpRanges` _string array_ | Allowed CIDR block or IP address in CIDR notation. |  |  |
| `anonGid` _float_ | Target group id when root squash is active. |  |  |
| `anonUid` _float_ | Target user id when root squash is active. |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Filesystem service. |  |  |
| `release` _string_ | The Filesystem service release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Filesystem service releases. |  |  |
| `rootSquash` _boolean_ | Map root user/group ownership to anon_uid/anon_gid |  |  |
| `securityZoneUuid` _string_ | Security zone UUID linked to Filesystem service. |  |  |


#### FilesystemList



FilesystemList contains a list of Filesystems





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `FilesystemList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Filesystem](#filesystem) array_ |  |  |  |


#### FilesystemObservation







_Appears in:_
- [FilesystemStatus](#filesystemstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowedIpRanges` _string array_ | Allowed CIDR block or IP address in CIDR notation. |  |  |
| `anonGid` _float_ | Target group id when root squash is active. |  |  |
| `anonUid` _float_ | Target user id when root squash is active. |  |  |
| `changeTime` _string_ | Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `listenPort` _[ListenPortObservation](#listenportobservation) array_ | The port numbers where this Filesystem service accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Filesystem service. |  |  |
| `release` _string_ | The Filesystem service release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Filesystem service releases. |  |  |
| `rootSquash` _boolean_ | Map root user/group ownership to anon_uid/anon_gid |  |  |
| `securityZoneUuid` _string_ | Security zone UUID linked to Filesystem service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Filesystem service uses. |  |  |
| `status` _string_ | Current status of Filesystem service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that Filesystem service is in use. |  |  |


#### FilesystemParameters







_Appears in:_
- [FilesystemSpec](#filesystemspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowedIpRanges` _string array_ | Allowed CIDR block or IP address in CIDR notation. |  | Optional: \{\} <br /> |
| `anonGid` _float_ | Target group id when root squash is active. |  | Optional: \{\} <br /> |
| `anonUid` _float_ | Target user id when root squash is active. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Filesystem service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Filesystem service release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Filesystem service releases. |  | Optional: \{\} <br /> |
| `rootSquash` _boolean_ | Map root user/group ownership to anon_uid/anon_gid |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | Security zone UUID linked to Filesystem service. |  | Optional: \{\} <br /> |


#### FilesystemSpec



FilesystemSpec defines the desired state of Filesystem



_Appears in:_
- [Filesystem](#filesystem)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[FilesystemParameters](#filesystemparameters)_ |  |  |  |
| `initProvider` _[FilesystemInitParameters](#filesysteminitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### FilesystemStatus



FilesystemStatus defines the observed state of Filesystem.



_Appears in:_
- [Filesystem](#filesystem)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[FilesystemObservation](#filesystemobservation)_ |  |  |  |


#### Firewall



Firewall is the Schema for the Firewalls API. Manages a firewall in gridscale.



_Appears in:_
- [FirewallList](#firewalllist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Firewall` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[FirewallSpec](#firewallspec)_ |  |  |  |
| `status` _[FirewallStatus](#firewallstatus)_ |  |  |  |


#### FirewallInitParameters







_Appears in:_
- [FirewallSpec](#firewallspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `rulesV4In` _[RulesV4InInitParameters](#rulesv4ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[RulesV4OutInitParameters](#rulesv4outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[RulesV6InInitParameters](#rulesv6ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[RulesV6OutInitParameters](#rulesv6outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |


#### FirewallList



FirewallList contains a list of Firewalls





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `FirewallList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Firewall](#firewall) array_ |  |  |  |


#### FirewallObservation







_Appears in:_
- [FirewallStatus](#firewallstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | The date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `createTime` _string_ | The date and time the object was initially created.<br />The date and time the object was initially created. |  |  |
| `description` _string_ | Description of the firewall.<br />Description of the Firewall. |  |  |
| `id` _string_ | The UUID of the firewall. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationName` _string_ | The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `network` _[NetworkObservation](#networkobservation) array_ | The information about networks which are related to this firewall. |  |  |
| `private` _boolean_ | The object is private, the value will be true. Otherwise the value will be false.<br />The object is private, the value will be true. Otherwise the value will be false. |  |  |
| `rulesV4In` _[RulesV4InObservation](#rulesv4inobservation) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[RulesV4OutObservation](#rulesv4outobservation) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[RulesV6InObservation](#rulesv6inobservation) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[RulesV6OutObservation](#rulesv6outobservation) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |
| `status` _string_ | Status indicates the status of the object.<br />Status indicates the status of the object |  |  |


#### FirewallParameters







_Appears in:_
- [FirewallSpec](#firewallspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `rulesV4In` _[RulesV4InParameters](#rulesv4inparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV4Out` _[RulesV4OutParameters](#rulesv4outparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV6In` _[RulesV6InParameters](#rulesv6inparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |
| `rulesV6Out` _[RulesV6OutParameters](#rulesv6outparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |


#### FirewallSpec



FirewallSpec defines the desired state of Firewall



_Appears in:_
- [Firewall](#firewall)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[FirewallParameters](#firewallparameters)_ |  |  |  |
| `initProvider` _[FirewallInitParameters](#firewallinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### FirewallStatus



FirewallStatus defines the observed state of Firewall.



_Appears in:_
- [Firewall](#firewall)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[FirewallObservation](#firewallobservation)_ |  |  |  |


#### ForwardingRuleInitParameters







_Appears in:_
- [LoadbalancerInitParameters](#loadbalancerinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateUuid` _string_ | The UUID of a custom certificate.<br />The UUID of a custom certificate. |  |  |
| `letsencryptSsl` _string_ | A valid domain name that points to the loadbalancer's IP address.<br />A valid domain name that points to the loadbalancer's IP address. |  |  |
| `listenPort` _float_ | Specifies the entry port of the load balancer.<br />Specifies the entry port of the load balancer. |  |  |
| `mode` _string_ | Supports HTTP and TCP mode. Valid values: http, tcp.<br />Supports HTTP and TCP mode. Valid values: http, tcp. |  |  |
| `targetPort` _float_ | Specifies the exit port that the load balancer uses to forward the traffic to the backend server.<br />Specifies the exit port that the load balancer uses to forward the traffic to the backend server. |  |  |


#### ForwardingRuleObservation







_Appears in:_
- [LoadbalancerObservation](#loadbalancerobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateUuid` _string_ | The UUID of a custom certificate.<br />The UUID of a custom certificate. |  |  |
| `letsencryptSsl` _string_ | A valid domain name that points to the loadbalancer's IP address.<br />A valid domain name that points to the loadbalancer's IP address. |  |  |
| `listenPort` _float_ | Specifies the entry port of the load balancer.<br />Specifies the entry port of the load balancer. |  |  |
| `mode` _string_ | Supports HTTP and TCP mode. Valid values: http, tcp.<br />Supports HTTP and TCP mode. Valid values: http, tcp. |  |  |
| `targetPort` _float_ | Specifies the exit port that the load balancer uses to forward the traffic to the backend server.<br />Specifies the exit port that the load balancer uses to forward the traffic to the backend server. |  |  |


#### ForwardingRuleParameters







_Appears in:_
- [LoadbalancerParameters](#loadbalancerparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateUuid` _string_ | The UUID of a custom certificate.<br />The UUID of a custom certificate. |  | Optional: \{\} <br /> |
| `letsencryptSsl` _string_ | A valid domain name that points to the loadbalancer's IP address.<br />A valid domain name that points to the loadbalancer's IP address. |  | Optional: \{\} <br /> |
| `listenPort` _float_ | Specifies the entry port of the load balancer.<br />Specifies the entry port of the load balancer. |  | Optional: \{\} <br /> |
| `mode` _string_ | Supports HTTP and TCP mode. Valid values: http, tcp.<br />Supports HTTP and TCP mode. Valid values: http, tcp. |  | Optional: \{\} <br /> |
| `targetPort` _float_ | Specifies the exit port that the load balancer uses to forward the traffic to the backend server.<br />Specifies the exit port that the load balancer uses to forward the traffic to the backend server. |  | Optional: \{\} <br /> |


#### HardwareProfileConfigInitParameters







_Appears in:_
- [ServerInitParameters_2](#serverinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hypervExtensions` _boolean_ | Boolean. |  |  |
| `machinetype` _string_ | Allowed values: "i440fx", "q35_bios", "q35_uefi". |  |  |
| `nestedVirtualization` _boolean_ | Boolean. |  |  |
| `networkModel` _string_ | Allowed values: "e1000", "e1000e", "virtio", "vmxnet3". |  |  |
| `serialInterface` _boolean_ | Boolean. |  |  |
| `serverRenice` _boolean_ | Boolean. |  |  |
| `storageDevice` _string_ | Allowed values: "ide", "sata", "virtio_scsi", "virtio_block". |  |  |
| `usbController` _string_ | Allowed values: "nec_xhci", "piix3_uhci". |  |  |


#### HardwareProfileConfigObservation







_Appears in:_
- [ServerObservation_2](#serverobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hypervExtensions` _boolean_ | Boolean. |  |  |
| `machinetype` _string_ | Allowed values: "i440fx", "q35_bios", "q35_uefi". |  |  |
| `nestedVirtualization` _boolean_ | Boolean. |  |  |
| `networkModel` _string_ | Allowed values: "e1000", "e1000e", "virtio", "vmxnet3". |  |  |
| `serialInterface` _boolean_ | Boolean. |  |  |
| `serverRenice` _boolean_ | Boolean. |  |  |
| `storageDevice` _string_ | Allowed values: "ide", "sata", "virtio_scsi", "virtio_block". |  |  |
| `usbController` _string_ | Allowed values: "nec_xhci", "piix3_uhci". |  |  |


#### HardwareProfileConfigParameters







_Appears in:_
- [ServerParameters_2](#serverparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hypervExtensions` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `machinetype` _string_ | Allowed values: "i440fx", "q35_bios", "q35_uefi". |  | Optional: \{\} <br /> |
| `nestedVirtualization` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `networkModel` _string_ | Allowed values: "e1000", "e1000e", "virtio", "vmxnet3". |  | Optional: \{\} <br /> |
| `serialInterface` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `serverRenice` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `storageDevice` _string_ | Allowed values: "ide", "sata", "virtio_scsi", "virtio_block". |  | Optional: \{\} <br /> |
| `usbController` _string_ | Allowed values: "nec_xhci", "piix3_uhci". |  | Optional: \{\} <br /> |


#### IPv4



IPv4 is the Schema for the IPv4s API. Manages an IPv4 address in gridscale.



_Appears in:_
- [IPv4List](#ipv4list)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv4` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[IPv4Spec](#ipv4spec)_ |  |  |  |
| `status` _[IPv4Status](#ipv4status)_ |  |  |  |


#### IPv4InitParameters







_Appears in:_
- [IPv4Spec](#ipv4spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |


#### IPv4List



IPv4List contains a list of IPv4s





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv4List` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[IPv4](#ipv4) array_ |  |  |  |


#### IPv4Observation







_Appears in:_
- [IPv4Status](#ipv4status)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill.<br />Defines the price for the current period since the last bill. |  |  |
| `deleteBlock` _boolean_ | Defines if the object is administratively blocked. If true, it can not be deleted by the user.<br />Defines if the object is administratively blocked. If true, it can not be deleted by the user. |  |  |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `id` _string_ |  |  |  |
| `ip` _string_ | Defines the IP address.<br />Defines the IP address. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | See Argument Reference above.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `prefix` _string_ | The network address and the subnet. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### IPv4Parameters







_Appears in:_
- [IPv4Spec](#ipv4spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  | Optional: \{\} <br /> |


#### IPv4Spec



IPv4Spec defines the desired state of IPv4



_Appears in:_
- [IPv4](#ipv4)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[IPv4Parameters](#ipv4parameters)_ |  |  |  |
| `initProvider` _[IPv4InitParameters](#ipv4initparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### IPv4Status



IPv4Status defines the observed state of IPv4.



_Appears in:_
- [IPv4](#ipv4)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[IPv4Observation](#ipv4observation)_ |  |  |  |


#### IPv6



IPv6 is the Schema for the IPv6s API. Manages an IPv6 address in gridscale.



_Appears in:_
- [IPv6List](#ipv6list)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv6` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[IPv6Spec](#ipv6spec)_ |  |  |  |
| `status` _[IPv6Status](#ipv6status)_ |  |  |  |


#### IPv6InitParameters







_Appears in:_
- [IPv6Spec](#ipv6spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |


#### IPv6List



IPv6List contains a list of IPv6s





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv6List` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[IPv6](#ipv6) array_ |  |  |  |


#### IPv6Observation







_Appears in:_
- [IPv6Status](#ipv6status)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill.<br />Defines the price for the current period since the last bill. |  |  |
| `deleteBlock` _boolean_ | Defines if the object is administratively blocked. If true, it can not be deleted by the user.<br />Defines if the object is administratively blocked. If true, it can not be deleted by the user. |  |  |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `id` _string_ |  |  |  |
| `ip` _string_ | Defines the IP address.<br />Defines the IP address. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this resource is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `prefix` _string_ | The network address and the subnet. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### IPv6Parameters







_Appears in:_
- [IPv6Spec](#ipv6spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  | Optional: \{\} <br /> |


#### IPv6Spec



IPv6Spec defines the desired state of IPv6



_Appears in:_
- [IPv6](#ipv6)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[IPv6Parameters](#ipv6parameters)_ |  |  |  |
| `initProvider` _[IPv6InitParameters](#ipv6initparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### IPv6Status



IPv6Status defines the observed state of IPv6.



_Appears in:_
- [IPv6](#ipv6)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[IPv6Observation](#ipv6observation)_ |  |  |  |


#### Isoimage



Isoimage is the Schema for the Isoimages API. <no value>



_Appears in:_
- [IsoimageList](#isoimagelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Isoimage` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[IsoimageSpec](#isoimagespec)_ |  |  |  |
| `status` _[IsoimageStatus](#isoimagestatus)_ |  |  |  |


#### IsoimageInitParameters







_Appears in:_
- [IsoimageSpec](#isoimagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sourceUrl` _string_ | Contains the source URL of the ISO image that it was originally fetched from. |  |  |


#### IsoimageList



IsoimageList contains a list of Isoimages





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IsoimageList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Isoimage](#isoimage) array_ |  |  |  |


#### IsoimageObservation







_Appears in:_
- [IsoimageStatus](#isoimagestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The capacity of a storage/ISO image/template/snapshot in GB. |  |  |
| `changeTime` _string_ | The date and time of the last object change. |  |  |
| `createTime` _string_ | The date and time the object was initially created. |  |  |
| `currentPrice` _float_ | Defines the price for the current period since the last bill. |  |  |
| `description` _string_ | Description of the ISO image. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `private` _boolean_ | The object is private, the value will be true. Otherwise the value will be false. |  |  |
| `server` _[ServerObservation](#serverobservation) array_ | The information about servers which are related to this ISO image. |  |  |
| `sourceUrl` _string_ | Contains the source URL of the ISO image that it was originally fetched from. |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `usageInMinutes` _float_ | Total minutes the object has been running. |  |  |
| `version` _string_ | Upstream version of the ISO image release |  |  |


#### IsoimageParameters







_Appears in:_
- [IsoimageSpec](#isoimagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `sourceUrl` _string_ | Contains the source URL of the ISO image that it was originally fetched from. |  | Optional: \{\} <br /> |


#### IsoimageSpec



IsoimageSpec defines the desired state of Isoimage



_Appears in:_
- [Isoimage](#isoimage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[IsoimageParameters](#isoimageparameters)_ |  |  |  |
| `initProvider` _[IsoimageInitParameters](#isoimageinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### IsoimageStatus



IsoimageStatus defines the observed state of Isoimage.



_Appears in:_
- [Isoimage](#isoimage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[IsoimageObservation](#isoimageobservation)_ |  |  |  |


#### K8S



K8S is the Schema for the K8Ss API. Manages a k8s cluster in gridscale.



_Appears in:_
- [K8SList](#k8slist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `K8S` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[K8SSpec](#k8sspec)_ |  |  |  |
| `status` _[K8SStatus](#k8sstatus)_ |  |  |  |


#### K8SInitParameters







_Appears in:_
- [K8SSpec](#k8sspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `auditLogEnabled` _boolean_ | Enable Kubernetes audit logs. |  |  |
| `auditLogLevel` _string_ | Audit log level. |  |  |
| `clusterCidr` _string_ | (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).<br />The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16" |  |  |
| `clusterTrafficEncryption` _boolean_ | Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.<br />Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false. |  |  |
| `gskVersion` _string_ | The gridscale's Kubernetes version of this instance (e.g. "1.30.3-gs0"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.<br />The gridscale k8s PaaS version (issued by gridscale) of this instance. |  |  |
| `k8sHubble` _boolean_ | Enable Hubble for the k8s cluster.<br />Enables Hubble Integration. |  |  |
| `kubeApiserverLogEnabled` _boolean_ | Enable kube-apiserver logs. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `logDelivery` _boolean_ | Enable control plane log delivery. |  |  |
| `logDeliveryAccessKey` _string_ | Access key used to authenticate against Object Storage endpoint. |  |  |
| `logDeliveryBucket` _string_ | Bucket to upload logs to. |  |  |
| `logDeliveryEndpoint` _string_ | Object Storage endpoint URL the bucket is located on. |  |  |
| `logDeliveryInterval` _float_ | Time interval (in min), at which log files will be delivered, unless file size limit is reached first. |  |  |
| `logDeliverySecretKey` _string_ | Secret key used to authenticate against Object Storage endpoint. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `nodePool` _[NodePoolInitParameters](#nodepoolinitparameters) array_ | The collection of node pool specifications. Mutiple node pools can be defined with multiple node_pool blocks. The node pool block supports the following arguments:<br />Define a node pool and its attributes. |  |  |
| `oidcCaPem` _string_ | Custom CA from customer in pem format as string.<br />Custom CA from customer in pem format as string. |  |  |
| `oidcClientId` _string_ | A client ID that all tokens must be issued for.<br />A client ID that all tokens must be issued for. |  |  |
| `oidcEnabled` _boolean_ | Enable OIDC for the k8s cluster.<br />Disable or enable OIDC |  |  |
| `oidcGroupsClaim` _string_ | JWT claim to use as the user's group.<br />JWT claim to use as the user's group. |  |  |
| `oidcGroupsPrefix` _string_ | Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.<br />Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra. |  |  |
| `oidcIssuerUrl` _string_ | URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.<br />URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted. |  |  |
| `oidcRequiredClaim` _string_ | A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.<br />A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2 |  |  |
| `oidcSigningAlgs` _string_ | The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.<br />The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'. |  |  |
| `oidcUsernameClaim` _string_ | JWT claim to use as the user name.<br />JWT claim to use as the user name. |  |  |
| `oidcUsernamePrefix` _string_ | Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.<br />Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. |  |  |
| `release` _string_ | The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.<br />The k8s release of this instance. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.<br />Security zone UUID linked to PaaS service. |  |  |
| `surgeNode` _boolean_ | Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).<br />Enable surge node to avoid resources shortage during the cluster upgrade. |  |  |


#### K8SList



K8SList contains a list of K8Ss





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `K8SList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[K8S](#k8s) array_ |  |  |  |




#### K8SListenPortObservation







_Appears in:_
- [K8SObservation](#k8sobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ | Name of the node pool. |  |  |
| `port` _float_ |  |  |  |




#### K8SObservation







_Appears in:_
- [K8SStatus](#k8sstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `auditLogEnabled` _boolean_ | Enable Kubernetes audit logs. |  |  |
| `auditLogLevel` _string_ | Audit log level. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Time of the last change |  |  |
| `clusterCidr` _string_ | (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).<br />The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16" |  |  |
| `clusterTrafficEncryption` _boolean_ | Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.<br />Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false. |  |  |
| `createTime` _string_ | The time the object was created.<br />Time this service was created. |  |  |
| `gskVersion` _string_ | The gridscale's Kubernetes version of this instance (e.g. "1.30.3-gs0"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.<br />The gridscale k8s PaaS version (issued by gridscale) of this instance. |  |  |
| `id` _string_ |  |  |  |
| `k8sHubble` _boolean_ | Enable Hubble for the k8s cluster.<br />Enables Hubble Integration. |  |  |
| `k8sPrivateNetworkUuid` _string_ | Private network UUID which k8s nodes are attached to. It can be used to attach other PaaS/VMs.<br />Private network UUID which k8s nodes are attached to. It can be used to attach other PaaS/VMs. |  |  |
| `kubeApiserverLogEnabled` _boolean_ | Enable kube-apiserver logs. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[K8SListenPortObservation](#k8slistenportobservation) array_ | The port number where this k8s service accepts connections. |  |  |
| `logDelivery` _boolean_ | Enable control plane log delivery. |  |  |
| `logDeliveryAccessKey` _string_ | Access key used to authenticate against Object Storage endpoint. |  |  |
| `logDeliveryBucket` _string_ | Bucket to upload logs to. |  |  |
| `logDeliveryEndpoint` _string_ | Object Storage endpoint URL the bucket is located on. |  |  |
| `logDeliveryInterval` _float_ | Time interval (in min), at which log files will be delivered, unless file size limit is reached first. |  |  |
| `logDeliverySecretKey` _string_ | Secret key used to authenticate against Object Storage endpoint. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | DEPRECATED  Network UUID containing security zone, which is linked to the k8s cluster.<br />Network UUID containing security zone |  |  |
| `nodePool` _[NodePoolObservation](#nodepoolobservation) array_ | The collection of node pool specifications. Mutiple node pools can be defined with multiple node_pool blocks. The node pool block supports the following arguments:<br />Define a node pool and its attributes. |  |  |
| `oidcCaPem` _string_ | Custom CA from customer in pem format as string.<br />Custom CA from customer in pem format as string. |  |  |
| `oidcClientId` _string_ | A client ID that all tokens must be issued for.<br />A client ID that all tokens must be issued for. |  |  |
| `oidcEnabled` _boolean_ | Enable OIDC for the k8s cluster.<br />Disable or enable OIDC |  |  |
| `oidcGroupsClaim` _string_ | JWT claim to use as the user's group.<br />JWT claim to use as the user's group. |  |  |
| `oidcGroupsPrefix` _string_ | Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.<br />Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra. |  |  |
| `oidcIssuerUrl` _string_ | URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.<br />URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted. |  |  |
| `oidcRequiredClaim` _string_ | A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.<br />A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2 |  |  |
| `oidcSigningAlgs` _string_ | The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.<br />The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'. |  |  |
| `oidcUsernameClaim` _string_ | JWT claim to use as the user name.<br />JWT claim to use as the user name. |  |  |
| `oidcUsernamePrefix` _string_ | Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.<br />Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. |  |  |
| `release` _string_ | The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.<br />The k8s release of this instance. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.<br />Security zone UUID linked to PaaS service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that k8s service uses.g. the k8s service is upgraded by gridscale staffs).<br />PaaS service template identifier for this service. |  |  |
| `status` _string_ | status indicates the status of the object.<br />Current status of the service |  |  |
| `surgeNode` _boolean_ | Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).<br />Enable surge node to avoid resources shortage during the cluster upgrade. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use.<br />Number of minutes that PaaS service is in use |  |  |


#### K8SParameters







_Appears in:_
- [K8SSpec](#k8sspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `auditLogEnabled` _boolean_ | Enable Kubernetes audit logs. |  | Optional: \{\} <br /> |
| `auditLogLevel` _string_ | Audit log level. |  | Optional: \{\} <br /> |
| `clusterCidr` _string_ | (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).<br />The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16" |  | Optional: \{\} <br /> |
| `clusterTrafficEncryption` _boolean_ | Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.<br />Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false. |  | Optional: \{\} <br /> |
| `gskVersion` _string_ | The gridscale's Kubernetes version of this instance (e.g. "1.30.3-gs0"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.<br />The gridscale k8s PaaS version (issued by gridscale) of this instance. |  | Optional: \{\} <br /> |
| `k8sHubble` _boolean_ | Enable Hubble for the k8s cluster.<br />Enables Hubble Integration. |  | Optional: \{\} <br /> |
| `kubeApiserverLogEnabled` _boolean_ | Enable kube-apiserver logs. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `logDelivery` _boolean_ | Enable control plane log delivery. |  | Optional: \{\} <br /> |
| `logDeliveryAccessKey` _string_ | Access key used to authenticate against Object Storage endpoint. |  | Optional: \{\} <br /> |
| `logDeliveryBucket` _string_ | Bucket to upload logs to. |  | Optional: \{\} <br /> |
| `logDeliveryEndpoint` _string_ | Object Storage endpoint URL the bucket is located on. |  | Optional: \{\} <br /> |
| `logDeliveryInterval` _float_ | Time interval (in min), at which log files will be delivered, unless file size limit is reached first. |  | Optional: \{\} <br /> |
| `logDeliverySecretKey` _string_ | Secret key used to authenticate against Object Storage endpoint. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `nodePool` _[NodePoolParameters](#nodepoolparameters) array_ | The collection of node pool specifications. Mutiple node pools can be defined with multiple node_pool blocks. The node pool block supports the following arguments:<br />Define a node pool and its attributes. |  | Optional: \{\} <br /> |
| `oidcCaPem` _string_ | Custom CA from customer in pem format as string.<br />Custom CA from customer in pem format as string. |  | Optional: \{\} <br /> |
| `oidcClientId` _string_ | A client ID that all tokens must be issued for.<br />A client ID that all tokens must be issued for. |  | Optional: \{\} <br /> |
| `oidcEnabled` _boolean_ | Enable OIDC for the k8s cluster.<br />Disable or enable OIDC |  | Optional: \{\} <br /> |
| `oidcGroupsClaim` _string_ | JWT claim to use as the user's group.<br />JWT claim to use as the user's group. |  | Optional: \{\} <br /> |
| `oidcGroupsPrefix` _string_ | Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.<br />Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra. |  | Optional: \{\} <br /> |
| `oidcIssuerUrl` _string_ | URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.<br />URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted. |  | Optional: \{\} <br /> |
| `oidcRequiredClaim` _string_ | A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.<br />A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2 |  | Optional: \{\} <br /> |
| `oidcSigningAlgs` _string_ | The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.<br />The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'. |  | Optional: \{\} <br /> |
| `oidcUsernameClaim` _string_ | JWT claim to use as the user name.<br />JWT claim to use as the user name. |  | Optional: \{\} <br /> |
| `oidcUsernamePrefix` _string_ | Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.<br />Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. |  | Optional: \{\} <br /> |
| `release` _string_ | The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.<br />The k8s release of this instance. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.<br />Security zone UUID linked to PaaS service. |  | Optional: \{\} <br /> |
| `surgeNode` _boolean_ | Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).<br />Enable surge node to avoid resources shortage during the cluster upgrade. |  | Optional: \{\} <br /> |


#### K8SSpec



K8SSpec defines the desired state of K8S



_Appears in:_
- [K8S](#k8s)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[K8SParameters](#k8sparameters)_ |  |  |  |
| `initProvider` _[K8SInitParameters](#k8sinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### K8SStatus



K8SStatus defines the observed state of K8S.



_Appears in:_
- [K8S](#k8s)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[K8SObservation](#k8sobservation)_ |  |  |  |


#### LabelsInitParameters







_Appears in:_
- [NodePoolInitParameters](#nodepoolinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `key` _string_ | The key of the label. |  |  |
| `value` _string_ | The value of the label. |  |  |


#### LabelsObservation







_Appears in:_
- [NodePoolObservation](#nodepoolobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `key` _string_ | The key of the label. |  |  |
| `value` _string_ | The value of the label. |  |  |


#### LabelsParameters







_Appears in:_
- [NodePoolParameters](#nodepoolparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `key` _string_ | The key of the label. |  | Optional: \{\} <br /> |
| `value` _string_ | The value of the label. |  | Optional: \{\} <br /> |




#### ListenPortObservation







_Appears in:_
- [FilesystemObservation](#filesystemobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ |  |  |  |
| `name` _string_ |  |  |  |
| `port` _float_ |  |  |  |




#### Loadbalancer



Loadbalancer is the Schema for the Loadbalancers API. Manage a loadbalancer in gridscale.



_Appears in:_
- [LoadbalancerList](#loadbalancerlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Loadbalancer` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[LoadbalancerSpec](#loadbalancerspec)_ |  |  |  |
| `status` _[LoadbalancerStatus](#loadbalancerstatus)_ |  |  |  |


#### LoadbalancerInitParameters







_Appears in:_
- [LoadbalancerSpec](#loadbalancerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `algorithm` _string_ | The algorithm used to process requests. Accepted values: roundrobin/leastconn.<br />The algorithm used to process requests. Accepted values: roundrobin/leastconn. |  |  |
| `backendServer` _[BackendServerInitParameters](#backendserverinitparameters) array_ | The servers that the load balancer can communicate with.<br />List of backend servers. |  |  |
| `forwardingRule` _[ForwardingRuleInitParameters](#forwardingruleinitparameters) array_ | The forwarding rules of the load balancer.<br />List of forwarding rules for the Load balancer. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenIpv4Uuid` _string_ | The UUID of the IPv4 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv4 address the Load balancer will listen to for incoming requests. |  |  |
| `listenIpv6Uuid` _string_ | The UUID of the IPv6 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv6 address the Load balancer will listen to for incoming requests. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `redirectHttpToHttps` _boolean_ | Whether the load balancer is forced to redirect requests from HTTP to HTTPS.<br />Whether the Load balancer is forced to redirect requests from HTTP to HTTPS |  |  |
| `status` _string_ | The status of the load balancer.<br />Status indicates the status of the object. |  |  |


#### LoadbalancerList



LoadbalancerList contains a list of Loadbalancers





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `LoadbalancerList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Loadbalancer](#loadbalancer) array_ |  |  |  |


#### LoadbalancerObservation







_Appears in:_
- [LoadbalancerStatus](#loadbalancerstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `algorithm` _string_ | The algorithm used to process requests. Accepted values: roundrobin/leastconn.<br />The algorithm used to process requests. Accepted values: roundrobin/leastconn. |  |  |
| `backendServer` _[BackendServerObservation](#backendserverobservation) array_ | The servers that the load balancer can communicate with.<br />List of backend servers. |  |  |
| `forwardingRule` _[ForwardingRuleObservation](#forwardingruleobservation) array_ | The forwarding rules of the load balancer.<br />List of forwarding rules for the Load balancer. |  |  |
| `id` _string_ | The UUID of the load balancer. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenIpv4Uuid` _string_ | The UUID of the IPv4 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv4 address the Load balancer will listen to for incoming requests. |  |  |
| `listenIpv6Uuid` _string_ | The UUID of the IPv6 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv6 address the Load balancer will listen to for incoming requests. |  |  |
| `locationUuid` _string_ | The location this load balancer is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `redirectHttpToHttps` _boolean_ | Whether the load balancer is forced to redirect requests from HTTP to HTTPS.<br />Whether the Load balancer is forced to redirect requests from HTTP to HTTPS |  |  |
| `status` _string_ | The status of the load balancer.<br />Status indicates the status of the object. |  |  |


#### LoadbalancerParameters







_Appears in:_
- [LoadbalancerSpec](#loadbalancerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `algorithm` _string_ | The algorithm used to process requests. Accepted values: roundrobin/leastconn.<br />The algorithm used to process requests. Accepted values: roundrobin/leastconn. |  | Optional: \{\} <br /> |
| `backendServer` _[BackendServerParameters](#backendserverparameters) array_ | The servers that the load balancer can communicate with.<br />List of backend servers. |  | Optional: \{\} <br /> |
| `forwardingRule` _[ForwardingRuleParameters](#forwardingruleparameters) array_ | The forwarding rules of the load balancer.<br />List of forwarding rules for the Load balancer. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `listenIpv4Uuid` _string_ | The UUID of the IPv4 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv4 address the Load balancer will listen to for incoming requests. |  | Optional: \{\} <br /> |
| `listenIpv6Uuid` _string_ | The UUID of the IPv6 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv6 address the Load balancer will listen to for incoming requests. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `redirectHttpToHttps` _boolean_ | Whether the load balancer is forced to redirect requests from HTTP to HTTPS.<br />Whether the Load balancer is forced to redirect requests from HTTP to HTTPS |  | Optional: \{\} <br /> |
| `status` _string_ | The status of the load balancer.<br />Status indicates the status of the object. |  | Optional: \{\} <br /> |


#### LoadbalancerSpec



LoadbalancerSpec defines the desired state of Loadbalancer



_Appears in:_
- [Loadbalancer](#loadbalancer)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[LoadbalancerParameters](#loadbalancerparameters)_ |  |  |  |
| `initProvider` _[LoadbalancerInitParameters](#loadbalancerinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### LoadbalancerStatus



LoadbalancerStatus defines the observed state of Loadbalancer.



_Appears in:_
- [Loadbalancer](#loadbalancer)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[LoadbalancerObservation](#loadbalancerobservation)_ |  |  |  |


#### Mariadb



Mariadb is the Schema for the Mariadbs API. Manage a MariaDB service in gridscale.



_Appears in:_
- [MariadbList](#mariadblist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Mariadb` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MariadbSpec](#mariadbspec)_ |  |  |  |
| `status` _[MariadbStatus](#mariadbstatus)_ |  |  |  |


#### MariadbInitParameters







_Appears in:_
- [MariadbSpec](#mariadbspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `mariadbBinlogFormat` _string_ | MariaDB parameter: Binary Logging Format. Default: "MIXED".<br />Binary Logging Format. |  |  |
| `mariadbDefaultTimeZone` _string_ | MariaDB parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mariadbLogBin` _boolean_ | MariaDB parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mariadbMaxAllowedPacket` _string_ | MariaDB parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbMaxConnections` _float_ | MariaDB parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mariadbQueryCache` _boolean_ | MariaDB parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mariadbQueryCacheLimit` _string_ | MariaDB parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbQueryCacheSize` _string_ | MariaDB parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbSqlMode` _string_ | MariaDB parameter: SQL Mode. Default: "NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO".<br />SQL Mode. |  |  |
| `mariadbServerId` _float_ | MariaDB parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MariaDB service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MariaDB service. |  |  |
| `release` _string_ | The MariaDB release of this instance. For convenience, please use gscloud to get the list of available MariaDB service releases.<br />The MariaDB release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MariaDB service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MariaDB service. |  |  |


#### MariadbList



MariadbList contains a list of Mariadbs





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MariadbList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Mariadb](#mariadb) array_ |  |  |  |




#### MariadbListenPortObservation







_Appears in:_
- [MariadbObservation](#mariadbobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MariadbObservation







_Appears in:_
- [MariadbStatus](#mariadbstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[MariadbListenPortObservation](#mariadblistenportobservation) array_ | The port numbers where this MariaDB service accepts connections.<br />The port numbers where this MariaDB service accepts connections. |  |  |
| `mariadbBinlogFormat` _string_ | MariaDB parameter: Binary Logging Format. Default: "MIXED".<br />Binary Logging Format. |  |  |
| `mariadbDefaultTimeZone` _string_ | MariaDB parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mariadbLogBin` _boolean_ | MariaDB parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mariadbMaxAllowedPacket` _string_ | MariaDB parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbMaxConnections` _float_ | MariaDB parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mariadbQueryCache` _boolean_ | MariaDB parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mariadbQueryCacheLimit` _string_ | MariaDB parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbQueryCacheSize` _string_ | MariaDB parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbSqlMode` _string_ | MariaDB parameter: SQL Mode. Default: "NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO".<br />SQL Mode. |  |  |
| `mariadbServerId` _float_ | MariaDB parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MariaDB service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MariaDB service. |  |  |
| `release` _string_ | The MariaDB release of this instance. For convenience, please use gscloud to get the list of available MariaDB service releases.<br />The MariaDB release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MariaDB service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MariaDB service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that MariaDB service uses.<br />PaaS service template that MariaDB service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MariaDB service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MariaDB service is in use. |  |  |


#### MariadbParameters







_Appears in:_
- [MariadbSpec](#mariadbspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `mariadbBinlogFormat` _string_ | MariaDB parameter: Binary Logging Format. Default: "MIXED".<br />Binary Logging Format. |  | Optional: \{\} <br /> |
| `mariadbDefaultTimeZone` _string_ | MariaDB parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  | Optional: \{\} <br /> |
| `mariadbLogBin` _boolean_ | MariaDB parameter: Binary Logging. Default: false.<br />Binary Logging. |  | Optional: \{\} <br /> |
| `mariadbMaxAllowedPacket` _string_ | MariaDB parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mariadbMaxConnections` _float_ | MariaDB parameter: Max Connections. Default: 4000.<br />Max Connections. |  | Optional: \{\} <br /> |
| `mariadbQueryCache` _boolean_ | MariaDB parameter: Enable query cache. Default: true.<br />Enable query cache. |  | Optional: \{\} <br /> |
| `mariadbQueryCacheLimit` _string_ | MariaDB parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mariadbQueryCacheSize` _string_ | MariaDB parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mariadbSqlMode` _string_ | MariaDB parameter: SQL Mode. Default: "NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO".<br />SQL Mode. |  | Optional: \{\} <br /> |
| `mariadbServerId` _float_ | MariaDB parameter: Server Id. Default: 1.<br />Server Id. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of MariaDB service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MariaDB service. |  | Optional: \{\} <br /> |
| `release` _string_ | The MariaDB release of this instance. For convenience, please use gscloud to get the list of available MariaDB service releases.<br />The MariaDB release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MariaDB service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MariaDB service. |  | Optional: \{\} <br /> |


#### MariadbSpec



MariadbSpec defines the desired state of Mariadb



_Appears in:_
- [Mariadb](#mariadb)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[MariadbParameters](#mariadbparameters)_ |  |  |  |
| `initProvider` _[MariadbInitParameters](#mariadbinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MariadbStatus



MariadbStatus defines the observed state of Mariadb.



_Appears in:_
- [Mariadb](#mariadb)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MariadbObservation](#mariadbobservation)_ |  |  |  |


#### Memcached



Memcached is the Schema for the Memcacheds API. Manage a Memcached service in gridscale



_Appears in:_
- [MemcachedList](#memcachedlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Memcached` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MemcachedSpec](#memcachedspec)_ |  |  |  |
| `status` _[MemcachedStatus](#memcachedstatus)_ |  |  |  |


#### MemcachedInitParameters







_Appears in:_
- [MemcachedSpec](#memcachedspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Memcached service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Memcached service. |  |  |
| `release` _string_ | The Memcached release of this instance. For convenience, please use gscloud to get the list of available Memcached service releases.<br />The Memcached release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Memcached service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Memcached service. |  |  |


#### MemcachedList



MemcachedList contains a list of Memcacheds





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MemcachedList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Memcached](#memcached) array_ |  |  |  |




#### MemcachedListenPortObservation







_Appears in:_
- [MemcachedObservation](#memcachedobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MemcachedObservation







_Appears in:_
- [MemcachedStatus](#memcachedstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[MemcachedListenPortObservation](#memcachedlistenportobservation) array_ | The port numbers where this Memcached service accepts connections.<br />The port numbers where this Memcached service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Memcached service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Memcached service. |  |  |
| `release` _string_ | The Memcached release of this instance. For convenience, please use gscloud to get the list of available Memcached service releases.<br />The Memcached release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Memcached service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Memcached service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Memcached service uses.<br />PaaS service template that Memcached service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of Memcached service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that Memcached service is in use. |  |  |


#### MemcachedParameters







_Appears in:_
- [MemcachedSpec](#memcachedspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Memcached service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Memcached service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Memcached release of this instance. For convenience, please use gscloud to get the list of available Memcached service releases.<br />The Memcached release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Memcached service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Memcached service. |  | Optional: \{\} <br /> |


#### MemcachedSpec



MemcachedSpec defines the desired state of Memcached



_Appears in:_
- [Memcached](#memcached)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[MemcachedParameters](#memcachedparameters)_ |  |  |  |
| `initProvider` _[MemcachedInitParameters](#memcachedinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MemcachedStatus



MemcachedStatus defines the observed state of Memcached.



_Appears in:_
- [Memcached](#memcached)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MemcachedObservation](#memcachedobservation)_ |  |  |  |


#### MySQL



MySQL is the Schema for the MySQLs API. Manage a MySQL service in gridscale.



_Appears in:_
- [MySQLList](#mysqllist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQL` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MySQLSpec](#mysqlspec)_ |  |  |  |
| `status` _[MySQLStatus](#mysqlstatus)_ |  |  |  |


#### MySQLInitParameters







_Appears in:_
- [MySQLSpec](#mysqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlBinlogFormat` _string_ | mysql parameter: Binary Logging Format. Default: "ROW".<br />Binary Logging Format. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlLogBin` _boolean_ | mysql parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlQueryCache` _boolean_ | mysql parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mysqlQueryCacheLimit` _string_ | mysql parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlQueryCacheSize` _string_ | mysql parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `mysqlServerId` _float_ | mysql parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |


#### MySQLList



MySQLList contains a list of MySQLs





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQLList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[MySQL](#mysql) array_ |  |  |  |




#### MySQLListenPortObservation







_Appears in:_
- [MySQLObservation](#mysqlobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MySQLObservation







_Appears in:_
- [MySQLStatus](#mysqlstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[MySQLListenPortObservation](#mysqllistenportobservation) array_ | The port numbers where this mysql service accepts connections.<br />The port numbers where this MySQL service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlBinlogFormat` _string_ | mysql parameter: Binary Logging Format. Default: "ROW".<br />Binary Logging Format. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlLogBin` _boolean_ | mysql parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlQueryCache` _boolean_ | mysql parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mysqlQueryCacheLimit` _string_ | mysql parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlQueryCacheSize` _string_ | mysql parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `mysqlServerId` _float_ | mysql parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that mysql service uses.<br />PaaS service template that MySQL service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MySQL service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MySQL service is in use. |  |  |


#### MySQLParameters







_Appears in:_
- [MySQLSpec](#mysqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `mysqlBinlogFormat` _string_ | mysql parameter: Binary Logging Format. Default: "ROW".<br />Binary Logging Format. |  | Optional: \{\} <br /> |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  | Optional: \{\} <br /> |
| `mysqlLogBin` _boolean_ | mysql parameter: Binary Logging. Default: false.<br />Binary Logging. |  | Optional: \{\} <br /> |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  | Optional: \{\} <br /> |
| `mysqlQueryCache` _boolean_ | mysql parameter: Enable query cache. Default: true.<br />Enable query cache. |  | Optional: \{\} <br /> |
| `mysqlQueryCacheLimit` _string_ | mysql parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlQueryCacheSize` _string_ | mysql parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  | Optional: \{\} <br /> |
| `mysqlServerId` _float_ | mysql parameter: Server Id. Default: 1.<br />Server Id. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  | Optional: \{\} <br /> |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  | Optional: \{\} <br /> |


#### MySQLSpec



MySQLSpec defines the desired state of MySQL



_Appears in:_
- [MySQL](#mysql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[MySQLParameters](#mysqlparameters)_ |  |  |  |
| `initProvider` _[MySQLInitParameters](#mysqlinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MySQLStatus



MySQLStatus defines the observed state of MySQL.



_Appears in:_
- [MySQL](#mysql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MySQLObservation](#mysqlobservation)_ |  |  |  |


#### Network



Network is the Schema for the Networks API. Manages a network in gridscale.



_Appears in:_
- [NetworkList](#networklist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Network` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[NetworkSpec](#networkspec)_ |  |  |  |
| `status` _[NetworkStatus](#networkstatus)_ |  |  |  |




#### NetworkInitParameters_2







_Appears in:_
- [NetworkSpec](#networkspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `dhcpActive` _boolean_ | Enable DHCP.<br />Enable DHCP. |  |  |
| `dhcpDns` _string_ | The IP address reserved and communicated by the dhcp service to be the default gateway.<br />DHCP DNS. |  |  |
| `dhcpGateway` _string_ | The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />The IP address reserved and communicated by the dhcp service to be the default gateway. |  |  |
| `dhcpRange` _string_ | DHCP DNS. If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically.<br />The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically. |  |  |
| `dhcpReservedSubnet` _string array_ | Subrange within the IP range.<br />Subrange within the IP range. |  |  |
| `l2security` _boolean_ | Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.<br />MAC spoofing protection - filters layer2 and ARP traffic based on source MAC |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |


#### NetworkList



NetworkList contains a list of Networks





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `NetworkList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Network](#network) array_ |  |  |  |


#### NetworkObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `createTime` _string_ | The date and time the object was initially created. |  |  |
| `networkName` _string_ | Name of the network. |  |  |
| `networkUuid` _string_ | The object UUID or id of the network. |  |  |
| `objectName` _string_ | Name of the firewall. |  |  |
| `objectUuid` _string_ | The object UUID or id of the firewall. |  |  |


#### NetworkObservation_2







_Appears in:_
- [NetworkStatus](#networkstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoAssignedServers` _[AutoAssignedServersObservation](#autoassignedserversobservation) array_ | A list of server UUIDs with the corresponding IPs that are designated by the DHCP server.<br />A list of server UUIDs with the corresponding IPs that are designated by the DHCP server. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created |  |  |
| `dhcpActive` _boolean_ | Enable DHCP.<br />Enable DHCP. |  |  |
| `dhcpDns` _string_ | The IP address reserved and communicated by the dhcp service to be the default gateway.<br />DHCP DNS. |  |  |
| `dhcpGateway` _string_ | The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />The IP address reserved and communicated by the dhcp service to be the default gateway. |  |  |
| `dhcpRange` _string_ | DHCP DNS. If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically.<br />The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically. |  |  |
| `dhcpReservedSubnet` _string array_ | Subrange within the IP range.<br />Subrange within the IP range. |  |  |
| `deleteBlock` _boolean_ | If deleting this network is allowed.<br />If deleting this network is allowed |  |  |
| `id` _string_ |  |  |  |
| `l2security` _boolean_ | Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.<br />MAC spoofing protection - filters layer2 and ARP traffic based on source MAC |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this network is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkType` _string_ | The type of this network, can be mpls, breakout or network.<br />The type of this network, can be mpls, breakout or network. |  |  |
| `pinnedServers` _[PinnedServersObservation](#pinnedserversobservation) array_ | A list of server UUIDs with the corresponding IPs that are designated by the user.<br />A list of server UUIDs with the corresponding IPs that are designated by the user. |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |




#### NetworkParameters_2







_Appears in:_
- [NetworkSpec](#networkspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `dhcpActive` _boolean_ | Enable DHCP.<br />Enable DHCP. |  | Optional: \{\} <br /> |
| `dhcpDns` _string_ | The IP address reserved and communicated by the dhcp service to be the default gateway.<br />DHCP DNS. |  | Optional: \{\} <br /> |
| `dhcpGateway` _string_ | The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />The IP address reserved and communicated by the dhcp service to be the default gateway. |  | Optional: \{\} <br /> |
| `dhcpRange` _string_ | DHCP DNS. If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically.<br />The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically. |  | Optional: \{\} <br /> |
| `dhcpReservedSubnet` _string array_ | Subrange within the IP range.<br />Subrange within the IP range. |  | Optional: \{\} <br /> |
| `l2security` _boolean_ | Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.<br />MAC spoofing protection - filters layer2 and ARP traffic based on source MAC |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |


#### NetworkRulesV4InInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4InObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4InParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkRulesV4OutInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4OutObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4OutParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkRulesV6InInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6InObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6InParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkRulesV6OutInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6OutObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6OutParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkSpec



NetworkSpec defines the desired state of Network



_Appears in:_
- [Network](#network)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[NetworkParameters_2](#networkparameters_2)_ |  |  |  |
| `initProvider` _[NetworkInitParameters_2](#networkinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### NetworkStatus



NetworkStatus defines the observed state of Network.



_Appears in:_
- [Network](#network)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[NetworkObservation_2](#networkobservation_2)_ |  |  |  |


#### NodePoolInitParameters







_Appears in:_
- [K8SInitParameters](#k8sinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cores` _float_ | Cores per worker node.<br />Cores per worker node. |  |  |
| `labels` _[LabelsInitParameters](#labelsinitparameters) array_ | List of labels to be applied to the nodes of this pool. Check the product documentation for details<br />List of labels to be applied to the nodes of this pool. |  |  |
| `memory` _float_ | Memory per worker node (in GiB).<br />Memory per worker node (in GiB). |  |  |
| `name` _string_ | Name of the node pool.<br />Name of node pool. |  |  |
| `nodeCount` _float_ | Number of worker nodes.<br />Number of worker nodes. |  |  |
| `rocketStorage` _float_ | Rocket storage per worker node (in GiB).<br />Rocket storage per worker node (in GiB). |  |  |
| `storage` _float_ | Storage per worker node (in GiB).<br />Storage per worker node (in GiB). |  |  |
| `storageType` _string_ | Storage type (one of storage, storage_high, storage_insane).<br />Storage type. |  |  |
| `taints` _[TaintsInitParameters](#taintsinitparameters) array_ | List of taints to be applied to the nodes of this pool. Check the product documentation for details<br />List of taints to be applied to the nodes of this pool. |  |  |


#### NodePoolObservation







_Appears in:_
- [K8SObservation](#k8sobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cores` _float_ | Cores per worker node.<br />Cores per worker node. |  |  |
| `labels` _[LabelsObservation](#labelsobservation) array_ | List of labels to be applied to the nodes of this pool. Check the product documentation for details<br />List of labels to be applied to the nodes of this pool. |  |  |
| `memory` _float_ | Memory per worker node (in GiB).<br />Memory per worker node (in GiB). |  |  |
| `name` _string_ | Name of the node pool.<br />Name of node pool. |  |  |
| `nodeCount` _float_ | Number of worker nodes.<br />Number of worker nodes. |  |  |
| `rocketStorage` _float_ | Rocket storage per worker node (in GiB).<br />Rocket storage per worker node (in GiB). |  |  |
| `storage` _float_ | Storage per worker node (in GiB).<br />Storage per worker node (in GiB). |  |  |
| `storageType` _string_ | Storage type (one of storage, storage_high, storage_insane).<br />Storage type. |  |  |
| `taints` _[TaintsObservation](#taintsobservation) array_ | List of taints to be applied to the nodes of this pool. Check the product documentation for details<br />List of taints to be applied to the nodes of this pool. |  |  |


#### NodePoolParameters







_Appears in:_
- [K8SParameters](#k8sparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cores` _float_ | Cores per worker node.<br />Cores per worker node. |  | Optional: \{\} <br /> |
| `labels` _[LabelsParameters](#labelsparameters) array_ | List of labels to be applied to the nodes of this pool. Check the product documentation for details<br />List of labels to be applied to the nodes of this pool. |  | Optional: \{\} <br /> |
| `memory` _float_ | Memory per worker node (in GiB).<br />Memory per worker node (in GiB). |  | Optional: \{\} <br /> |
| `name` _string_ | Name of the node pool.<br />Name of node pool. |  | Optional: \{\} <br /> |
| `nodeCount` _float_ | Number of worker nodes.<br />Number of worker nodes. |  | Optional: \{\} <br /> |
| `rocketStorage` _float_ | Rocket storage per worker node (in GiB).<br />Rocket storage per worker node (in GiB). |  | Optional: \{\} <br /> |
| `storage` _float_ | Storage per worker node (in GiB).<br />Storage per worker node (in GiB). |  | Optional: \{\} <br /> |
| `storageType` _string_ | Storage type (one of storage, storage_high, storage_insane).<br />Storage type. |  | Optional: \{\} <br /> |
| `taints` _[TaintsParameters](#taintsparameters) array_ | List of taints to be applied to the nodes of this pool. Check the product documentation for details<br />List of taints to be applied to the nodes of this pool. |  | Optional: \{\} <br /> |


#### ObjectStorageExportInitParameters







_Appears in:_
- [SnapshotInitParameters](#snapshotinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Access key |  |  |
| `bucket` _string_ | Bucket name |  |  |
| `host` _string_ | Host of object storage. Must be of URL type. E.g: https://gos3.io |  |  |
| `object` _string_ | Name of file (include file path) |  |  |
| `private` _boolean_ |  |  |  |
| `secretKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Secret key |  |  |


#### ObjectStorageExportObservation







_Appears in:_
- [SnapshotObservation](#snapshotobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bucket` _string_ | Bucket name |  |  |
| `host` _string_ | Host of object storage. Must be of URL type. E.g: https://gos3.io |  |  |
| `object` _string_ | Name of file (include file path) |  |  |
| `private` _boolean_ |  |  |  |
| `status` _string_ |  |  |  |


#### ObjectStorageExportParameters







_Appears in:_
- [SnapshotParameters](#snapshotparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Access key |  | Optional: \{\} <br /> |
| `bucket` _string_ | Bucket name |  | Optional: \{\} <br /> |
| `host` _string_ | Host of object storage. Must be of URL type. E.g: https://gos3.io |  | Optional: \{\} <br /> |
| `object` _string_ | Name of file (include file path) |  | Optional: \{\} <br /> |
| `private` _boolean_ |  |  | Optional: \{\} <br /> |
| `secretKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Secret key |  | Optional: \{\} <br /> |


#### Paas



Paas is the Schema for the Paass API. Manages a PaaS in gridscale.



_Appears in:_
- [PaasList](#paaslist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Paas` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[PaasSpec](#paasspec)_ |  |  |  |
| `status` _[PaasStatus](#paasstatus)_ |  |  |  |


#### PaasInitParameters







_Appears in:_
- [PaasSpec](#paasspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `parameter` _[ParameterInitParameters](#parameterinitparameters) array_ | See Argument Reference above.<br />Parameter for PaaS service |  |  |
| `resourceLimit` _[ResourceLimitInitParameters](#resourcelimitinitparameters) array_ | A list of service resource limits..<br />Resource for PaaS service |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PaaS service |  |  |
| `serviceTemplateUuid` _string_ | The template used to create the service.<br />Template that PaaS service uses |  |  |


#### PaasList



PaasList contains a list of Paass





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `PaasList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Paas](#paas) array_ |  |  |  |




#### PaasListenPortObservation







_Appears in:_
- [PaasObservation](#paasobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### PaasObservation







_Appears in:_
- [PaasStatus](#paasstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change |  |  |
| `createTime` _string_ | Time of the creation.<br />Time of the creation |  |  |
| `currentPrice` _float_ | Current price of PaaS service.<br />Current price of PaaS service |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[PaasListenPortObservation](#paaslistenportobservation) array_ | Ports that PaaS service listens to.<br />Ports that PaaS service listens to |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `parameter` _[ParameterObservation](#parameterobservation) array_ | See Argument Reference above.<br />Parameter for PaaS service |  |  |
| `resourceLimit` _[ResourceLimitObservation](#resourcelimitobservation) array_ | A list of service resource limits..<br />Resource for PaaS service |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PaaS service |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | The template used to create the service.<br />Template that PaaS service uses |  |  |
| `serviceTemplateUuidComputed` _string_ | Template that PaaS service uses.<br />Template that PaaS service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of PaaS service |  |  |
| `usageInMinute` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that PaaS service is in use |  |  |


#### PaasParameters







_Appears in:_
- [PaasSpec](#paasspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `parameter` _[ParameterParameters](#parameterparameters) array_ | See Argument Reference above.<br />Parameter for PaaS service |  | Optional: \{\} <br /> |
| `resourceLimit` _[ResourceLimitParameters](#resourcelimitparameters) array_ | A list of service resource limits..<br />Resource for PaaS service |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PaaS service |  | Optional: \{\} <br /> |
| `serviceTemplateUuid` _string_ | The template used to create the service.<br />Template that PaaS service uses |  | Optional: \{\} <br /> |


#### PaasSpec



PaasSpec defines the desired state of Paas



_Appears in:_
- [Paas](#paas)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[PaasParameters](#paasparameters)_ |  |  |  |
| `initProvider` _[PaasInitParameters](#paasinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### PaasStatus



PaasStatus defines the observed state of Paas.



_Appears in:_
- [Paas](#paas)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[PaasObservation](#paasobservation)_ |  |  |  |


#### ParameterInitParameters







_Appears in:_
- [PaasInitParameters](#paasinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `param` _string_ | Name of parameter. |  |  |
| `type` _string_ | Primitive type of the parameter: bool, int (better use float for int case), float, string. |  |  |
| `value` _string_ | Value of the corresponding parameter. |  |  |


#### ParameterObservation







_Appears in:_
- [PaasObservation](#paasobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `param` _string_ | Name of parameter. |  |  |
| `type` _string_ | Primitive type of the parameter: bool, int (better use float for int case), float, string. |  |  |
| `value` _string_ | Value of the corresponding parameter. |  |  |


#### ParameterParameters







_Appears in:_
- [PaasParameters](#paasparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `param` _string_ | Name of parameter. |  | Optional: \{\} <br /> |
| `type` _string_ | Primitive type of the parameter: bool, int (better use float for int case), float, string. |  | Optional: \{\} <br /> |
| `value` _string_ | Value of the corresponding parameter. |  | Optional: \{\} <br /> |




#### PinnedServersObservation







_Appears in:_
- [NetworkObservation_2](#networkobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `ip` _string_ | IP which is assigned to the server. |  |  |
| `serverUuid` _string_ | UUID of the server. |  |  |




#### Postgresql



Postgresql is the Schema for the Postgresqls API. Manage a PostgreSQL service in gridscale.



_Appears in:_
- [PostgresqlList](#postgresqllist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Postgresql` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[PostgresqlSpec](#postgresqlspec)_ |  |  |  |
| `status` _[PostgresqlStatus](#postgresqlstatus)_ |  |  |  |


#### PostgresqlInitParameters







_Appears in:_
- [PostgresqlSpec](#postgresqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of PostgreSQL service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of PostgreSQL service. |  |  |
| `pgauditLogAccessKey` _string_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  |  |
| `pgauditLogBucket` _string_ | Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key.<br />Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key. |  |  |
| `pgauditLogRotationFrequency` _float_ | Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. Default is 5 minutes.<br />Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. |  |  |
| `pgauditLogSecretKey` _string_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  |  |
| `pgauditLogServerUrl` _string_ | Object Storage server URL the bucket is located on.<br />Object Storage server URL the bucket is located on. |  |  |
| `release` _string_ | The PostgreSQL release of this instance. For convenience, please use gscloud to get the list of available PostgreSQL service releases.<br />The PostgreSQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available PostgreSQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PostgreSQL service. |  |  |


#### PostgresqlList



PostgresqlList contains a list of Postgresqls





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `PostgresqlList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Postgresql](#postgresql) array_ |  |  |  |




#### PostgresqlListenPortObservation







_Appears in:_
- [PostgresqlObservation](#postgresqlobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### PostgresqlObservation







_Appears in:_
- [PostgresqlStatus](#postgresqlstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[PostgresqlListenPortObservation](#postgresqllistenportobservation) array_ | The port numbers where this PostgreSQL service accepts connections.<br />The port numbers where this PostgreSQL service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of PostgreSQL service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of PostgreSQL service. |  |  |
| `pgauditLogAccessKey` _string_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  |  |
| `pgauditLogBucket` _string_ | Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key.<br />Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key. |  |  |
| `pgauditLogRotationFrequency` _float_ | Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. Default is 5 minutes.<br />Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. |  |  |
| `pgauditLogSecretKey` _string_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  |  |
| `pgauditLogServerUrl` _string_ | Object Storage server URL the bucket is located on.<br />Object Storage server URL the bucket is located on. |  |  |
| `release` _string_ | The PostgreSQL release of this instance. For convenience, please use gscloud to get the list of available PostgreSQL service releases.<br />The PostgreSQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available PostgreSQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PostgreSQL service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that PostgreSQL service uses.<br />PaaS service template that PostgreSQL service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of PostgreSQL service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that PostgreSQL service is in use. |  |  |


#### PostgresqlParameters







_Appears in:_
- [PostgresqlSpec](#postgresqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of PostgreSQL service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of PostgreSQL service. |  | Optional: \{\} <br /> |
| `pgauditLogAccessKey` _string_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `pgauditLogBucket` _string_ | Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key.<br />Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key. |  | Optional: \{\} <br /> |
| `pgauditLogRotationFrequency` _float_ | Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. Default is 5 minutes.<br />Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. |  | Optional: \{\} <br /> |
| `pgauditLogSecretKey` _string_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `pgauditLogServerUrl` _string_ | Object Storage server URL the bucket is located on.<br />Object Storage server URL the bucket is located on. |  | Optional: \{\} <br /> |
| `release` _string_ | The PostgreSQL release of this instance. For convenience, please use gscloud to get the list of available PostgreSQL service releases.<br />The PostgreSQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available PostgreSQL service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PostgreSQL service. |  | Optional: \{\} <br /> |


#### PostgresqlSpec



PostgresqlSpec defines the desired state of Postgresql



_Appears in:_
- [Postgresql](#postgresql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[PostgresqlParameters](#postgresqlparameters)_ |  |  |  |
| `initProvider` _[PostgresqlInitParameters](#postgresqlinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### PostgresqlStatus



PostgresqlStatus defines the observed state of Postgresql.



_Appears in:_
- [Postgresql](#postgresql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[PostgresqlObservation](#postgresqlobservation)_ |  |  |  |


#### ResourceLimitInitParameters







_Appears in:_
- [PaasInitParameters](#paasinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `limit` _float_ | The maximum number of the specific resource your service can use. |  |  |
| `resource` _string_ | The name of the resource you would like to cap. |  |  |


#### ResourceLimitObservation







_Appears in:_
- [PaasObservation](#paasobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `limit` _float_ | The maximum number of the specific resource your service can use. |  |  |
| `resource` _string_ | The name of the resource you would like to cap. |  |  |


#### ResourceLimitParameters







_Appears in:_
- [PaasParameters](#paasparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `limit` _float_ | The maximum number of the specific resource your service can use. |  | Optional: \{\} <br /> |
| `resource` _string_ | The name of the resource you would like to cap. |  | Optional: \{\} <br /> |


#### RollbackInitParameters







_Appears in:_
- [SnapshotInitParameters](#snapshotinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `id` _string_ | ID of the rollback request. Each rollback request has to have a unique id. ID can be any string value. |  |  |


#### RollbackObservation







_Appears in:_
- [SnapshotObservation](#snapshotobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `id` _string_ | ID of the rollback request. Each rollback request has to have a unique id. ID can be any string value. |  |  |
| `rollbackTime` _string_ |  |  |  |
| `status` _string_ |  |  |  |


#### RollbackParameters







_Appears in:_
- [SnapshotParameters](#snapshotparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `id` _string_ | ID of the rollback request. Each rollback request has to have a unique id. ID can be any string value. |  | Optional: \{\} <br /> |


#### RulesV4InInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4InObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4InParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### RulesV4OutInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4OutObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4OutParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### RulesV6InInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6InObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6InParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### RulesV6OutInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6OutObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6OutParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### S3BackupInitParameters







_Appears in:_
- [SqlserverInitParameters](#sqlserverinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `backupAccessKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  |  |
| `backupBucket` _string_ | Object Storage bucket to upload backups to and restore backups from.<br />Object Storage bucket to upload backups to and restore backups from. |  |  |
| `backupRetention` _float_ | Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage).<br />Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage). |  |  |
| `backupSecretKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  |  |
| `backupServerUrl` _string_ | Object Storage server URL the bucket is located on. Note: Currently, only object storage host "https://gos3.io/" is supported.<br />Object Storage server URL the bucket is located on. |  |  |


#### S3BackupObservation







_Appears in:_
- [SqlserverObservation](#sqlserverobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `backupBucket` _string_ | Object Storage bucket to upload backups to and restore backups from.<br />Object Storage bucket to upload backups to and restore backups from. |  |  |
| `backupRetention` _float_ | Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage).<br />Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage). |  |  |
| `backupServerUrl` _string_ | Object Storage server URL the bucket is located on. Note: Currently, only object storage host "https://gos3.io/" is supported.<br />Object Storage server URL the bucket is located on. |  |  |


#### S3BackupParameters







_Appears in:_
- [SqlserverParameters](#sqlserverparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `backupAccessKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `backupBucket` _string_ | Object Storage bucket to upload backups to and restore backups from.<br />Object Storage bucket to upload backups to and restore backups from. |  | Optional: \{\} <br /> |
| `backupRetention` _float_ | Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage).<br />Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage). |  | Optional: \{\} <br /> |
| `backupSecretKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `backupServerUrl` _string_ | Object Storage server URL the bucket is located on. Note: Currently, only object storage host "https://gos3.io/" is supported.<br />Object Storage server URL the bucket is located on. |  | Optional: \{\} <br /> |


#### Server



Server is the Schema for the Servers API. Manages a server in gridscale.



_Appears in:_
- [ServerList](#serverlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Server` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ServerSpec](#serverspec)_ |  |  |  |
| `status` _[ServerStatus](#serverstatus)_ |  |  |  |




#### ServerInitParameters_2







_Appears in:_
- [ServerSpec](#serverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoRecovery` _boolean_ | If the server should be auto-started in case of a failure (default=true).<br />If the server should be auto-started in case of a failure (default=true). |  |  |
| `availabilityZone` _string_ | Defines which Availability-Zone the Server is placed.<br />Defines which Availability-Zone the Server is placed. |  |  |
| `cores` _float_ | The number of server cores.<br />The number of server cores. |  |  |
| `hardwareProfile` _string_ | The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. If it is not set, the backend will set it by default. Check the official docs.<br />Specifies the hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `hardwareProfileConfig` _[HardwareProfileConfigInitParameters](#hardwareprofileconfiginitparameters) array_ | Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. Note: If hardware_profile_config is set, all fields of hardware_profile_config MUST be set.<br />Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `ipv4` _string_ | The UUID of the IPv4 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `ipv6` _string_ | The UUID of the IPv6 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `isoimage` _string_ | The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in the expert panel. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `memory` _float_ | The amount of server memory in GB.<br />The amount of server memory in GB. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `network` _[ServerNetworkInitParameters](#servernetworkinitparameters) array_ | Connects a network to the server. The network ordering of the server corresponds to the order of the networks in the server resource block. |  |  |
| `power` _boolean_ | The power state of the server. Set this to true to will boot the server, false will shut it down.<br />The number of server cores. |  |  |
| `storage` _[StorageInitParameters](#storageinitparameters) array_ | Connects a storage to the server. **NOTE: The first storage is always the boot device.<br />A list of storages attached to the server. The first storage in the list is always set as the boot storage of the server. |  |  |
| `userDataBase64` _string_ | For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition.<br />For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition. |  |  |


#### ServerList



ServerList contains a list of Servers





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ServerList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Server](#server) array_ |  |  |  |


#### ServerNetworkInitParameters







_Appears in:_
- [ServerInitParameters_2](#serverinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  |  |
| `firewallTemplateUuid` _string_ | The UUID of firewall template. |  |  |
| `ip` _string_ | Manually assign DHCP IP to the server (if applicable).<br />Manually assign DHCP IP to the server. |  |  |
| `objectUuid` _string_ | The object UUID or id of the network. |  |  |
| `ordering` _float_ | DEPRECATED  Defines the ordering of the network interfaces. Lower numbers have lower PCI-IDs. |  |  |
| `rulesV4In` _[NetworkRulesV4InInitParameters](#networkrulesv4ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[NetworkRulesV4OutInitParameters](#networkrulesv4outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[NetworkRulesV6InInitParameters](#networkrulesv6ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[NetworkRulesV6OutInitParameters](#networkrulesv6outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |


#### ServerNetworkObservation







_Appears in:_
- [ServerObservation_2](#serverobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoAssignedIp` _string_ | DHCP IP which is automatically assigned to the server (if applicable).<br />DHCP IP which is automatically assigned to the server. |  |  |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created. |  |  |
| `firewallTemplateUuid` _string_ | The UUID of firewall template. |  |  |
| `ip` _string_ | Manually assign DHCP IP to the server (if applicable).<br />Manually assign DHCP IP to the server. |  |  |
| `mac` _string_ | network_mac defines the MAC address of the network interface. |  |  |
| `networkType` _string_ | One of network, network_high, network_insane. |  |  |
| `objectName` _string_ | Name of the storage. |  |  |
| `objectUuid` _string_ | The object UUID or id of the network. |  |  |
| `ordering` _float_ | DEPRECATED  Defines the ordering of the network interfaces. Lower numbers have lower PCI-IDs. |  |  |
| `rulesV4In` _[NetworkRulesV4InObservation](#networkrulesv4inobservation) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[NetworkRulesV4OutObservation](#networkrulesv4outobservation) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[NetworkRulesV6InObservation](#networkrulesv6inobservation) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[NetworkRulesV6OutObservation](#networkrulesv6outobservation) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |


#### ServerNetworkParameters







_Appears in:_
- [ServerParameters_2](#serverparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  | Optional: \{\} <br /> |
| `firewallTemplateUuid` _string_ | The UUID of firewall template. |  | Optional: \{\} <br /> |
| `ip` _string_ | Manually assign DHCP IP to the server (if applicable).<br />Manually assign DHCP IP to the server. |  | Optional: \{\} <br /> |
| `objectUuid` _string_ | The object UUID or id of the network. |  | Optional: \{\} <br /> |
| `ordering` _float_ | DEPRECATED  Defines the ordering of the network interfaces. Lower numbers have lower PCI-IDs. |  | Optional: \{\} <br /> |
| `rulesV4In` _[NetworkRulesV4InParameters](#networkrulesv4inparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV4Out` _[NetworkRulesV4OutParameters](#networkrulesv4outparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV6In` _[NetworkRulesV6InParameters](#networkrulesv6inparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |
| `rulesV6Out` _[NetworkRulesV6OutParameters](#networkrulesv6outparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |


#### ServerObservation







_Appears in:_
- [IsoimageObservation](#isoimageobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ |  |  |  |
| `createTime` _string_ |  |  |  |
| `objectName` _string_ |  |  |  |
| `objectUuid` _string_ |  |  |  |


#### ServerObservation_2







_Appears in:_
- [ServerStatus](#serverstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoRecovery` _boolean_ | If the server should be auto-started in case of a failure (default=true).<br />If the server should be auto-started in case of a failure (default=true). |  |  |
| `availabilityZone` _string_ | Defines which Availability-Zone the Server is placed.<br />Defines which Availability-Zone the Server is placed. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change. |  |  |
| `consoleToken` _string_ | The token used by the panel to open the websocket VNC connection to the server console.<br />The token used by the panel to open the websocket VNC connection to the server console. |  |  |
| `cores` _float_ | The number of server cores.<br />The number of server cores. |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `hardwareProfile` _string_ | The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. If it is not set, the backend will set it by default. Check the official docs.<br />Specifies the hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `hardwareProfileConfig` _[HardwareProfileConfigObservation](#hardwareprofileconfigobservation) array_ | Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. Note: If hardware_profile_config is set, all fields of hardware_profile_config MUST be set.<br />Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `id` _string_ | UUID of the server. |  |  |
| `ipv4` _string_ | The UUID of the IPv4 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `ipv6` _string_ | The UUID of the IPv6 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `isoimage` _string_ | The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in the expert panel. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `legacy` _boolean_ | Legacy-Hardware emulation instead of virtio hardware. If enabled, hot-plugging cores, memory, storage, network, etc. will not work, but the server will most likely run every x86 compatible operating system. This mode comes with a performance penalty, as emulated hardware does not benefit from the virtio driver infrastructure.<br />Legacy-Hardware emulation instead of virtio hardware. If enabled, hotplugging cores, memory, storage, network, etc. will not work, but the server will most likely run every x86 compatible operating system. This mode comes with a performance penalty, as emulated hardware does not benefit from the virtio driver infrastructure. |  |  |
| `locationUuid` _string_ | The location this server is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `memory` _float_ | The amount of server memory in GB.<br />The amount of server memory in GB. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `network` _[ServerNetworkObservation](#servernetworkobservation) array_ | Connects a network to the server. The network ordering of the server corresponds to the order of the networks in the server resource block. |  |  |
| `power` _boolean_ | The power state of the server. Set this to true to will boot the server, false will shut it down.<br />The number of server cores. |  |  |
| `status` _string_ | Status indicates the status of the object. |  |  |
| `storage` _[StorageObservation](#storageobservation) array_ | Connects a storage to the server. **NOTE: The first storage is always the boot device.<br />A list of storages attached to the server. The first storage in the list is always set as the boot storage of the server. |  |  |
| `usageInMinutesCores` _float_ | Total minutes of cores used. |  |  |
| `usageInMinutesMemory` _float_ | Total minutes of memory used. |  |  |
| `userDataBase64` _string_ | For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition.<br />For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition. |  |  |




#### ServerParameters_2







_Appears in:_
- [ServerSpec](#serverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoRecovery` _boolean_ | If the server should be auto-started in case of a failure (default=true).<br />If the server should be auto-started in case of a failure (default=true). |  | Optional: \{\} <br /> |
| `availabilityZone` _string_ | Defines which Availability-Zone the Server is placed.<br />Defines which Availability-Zone the Server is placed. |  | Optional: \{\} <br /> |
| `cores` _float_ | The number of server cores.<br />The number of server cores. |  | Optional: \{\} <br /> |
| `hardwareProfile` _string_ | The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. If it is not set, the backend will set it by default. Check the official docs.<br />Specifies the hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  | Optional: \{\} <br /> |
| `hardwareProfileConfig` _[HardwareProfileConfigParameters](#hardwareprofileconfigparameters) array_ | Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. Note: If hardware_profile_config is set, all fields of hardware_profile_config MUST be set.<br />Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  | Optional: \{\} <br /> |
| `ipv4` _string_ | The UUID of the IPv4 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  | Optional: \{\} <br /> |
| `ipv6` _string_ | The UUID of the IPv6 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  | Optional: \{\} <br /> |
| `isoimage` _string_ | The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in the expert panel. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `memory` _float_ | The amount of server memory in GB.<br />The amount of server memory in GB. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `network` _[ServerNetworkParameters](#servernetworkparameters) array_ | Connects a network to the server. The network ordering of the server corresponds to the order of the networks in the server resource block. |  | Optional: \{\} <br /> |
| `power` _boolean_ | The power state of the server. Set this to true to will boot the server, false will shut it down.<br />The number of server cores. |  | Optional: \{\} <br /> |
| `storage` _[StorageParameters](#storageparameters) array_ | Connects a storage to the server. **NOTE: The first storage is always the boot device.<br />A list of storages attached to the server. The first storage in the list is always set as the boot storage of the server. |  | Optional: \{\} <br /> |
| `userDataBase64` _string_ | For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition.<br />For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition. |  | Optional: \{\} <br /> |


#### ServerSpec



ServerSpec defines the desired state of Server



_Appears in:_
- [Server](#server)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[ServerParameters_2](#serverparameters_2)_ |  |  |  |
| `initProvider` _[ServerInitParameters_2](#serverinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### ServerStatus



ServerStatus defines the observed state of Server.



_Appears in:_
- [Server](#server)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[ServerObservation_2](#serverobservation_2)_ |  |  |  |


#### Snapshot



Snapshot is the Schema for the Snapshots API. <no value>



_Appears in:_
- [SnapshotList](#snapshotlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Snapshot` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SnapshotSpec](#snapshotspec)_ |  |  |  |
| `status` _[SnapshotStatus](#snapshotstatus)_ |  |  |  |


#### SnapshotInitParameters







_Appears in:_
- [SnapshotSpec](#snapshotspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `objectStorageExport` _[ObjectStorageExportInitParameters](#objectstorageexportinitparameters) array_ | Export snapshot to a object storage |  |  |
| `rollback` _[RollbackInitParameters](#rollbackinitparameters) array_ | Returns a storage to the state of the selected Snapshot. |  |  |
| `storageUuid` _string_ | UUID of the storage used to create this snapshot |  |  |


#### SnapshotList



SnapshotList contains a list of Snapshots





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SnapshotList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Snapshot](#snapshot) array_ |  |  |  |


#### SnapshotObservation







_Appears in:_
- [SnapshotStatus](#snapshotstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The capacity of a storage/ISO image/template/snapshot in GB |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows<br />the product_no of the license (see the /prices endpoint for more details) |  |  |
| `locationCountry` _string_ | The human-readable name of the location |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The human-readable name of the location |  |  |
| `locationUuid` _string_ | The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `objectStorageExport` _[ObjectStorageExportObservation](#objectstorageexportobservation) array_ | Export snapshot to a object storage |  |  |
| `rollback` _[RollbackObservation](#rollbackobservation) array_ | Returns a storage to the state of the selected Snapshot. |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `storageUuid` _string_ | UUID of the storage used to create this snapshot |  |  |
| `usageInMinutes` _float_ | Total minutes the object has been running |  |  |


#### SnapshotParameters







_Appears in:_
- [SnapshotSpec](#snapshotspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object |  | Optional: \{\} <br /> |
| `objectStorageExport` _[ObjectStorageExportParameters](#objectstorageexportparameters) array_ | Export snapshot to a object storage |  | Optional: \{\} <br /> |
| `rollback` _[RollbackParameters](#rollbackparameters) array_ | Returns a storage to the state of the selected Snapshot. |  | Optional: \{\} <br /> |
| `storageUuid` _string_ | UUID of the storage used to create this snapshot |  | Optional: \{\} <br /> |


#### SnapshotSpec



SnapshotSpec defines the desired state of Snapshot



_Appears in:_
- [Snapshot](#snapshot)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[SnapshotParameters](#snapshotparameters)_ |  |  |  |
| `initProvider` _[SnapshotInitParameters](#snapshotinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SnapshotStatus



SnapshotStatus defines the observed state of Snapshot.



_Appears in:_
- [Snapshot](#snapshot)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SnapshotObservation](#snapshotobservation)_ |  |  |  |


#### Snapshotschedule



Snapshotschedule is the Schema for the Snapshotschedules API. <no value>



_Appears in:_
- [SnapshotscheduleList](#snapshotschedulelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Snapshotschedule` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SnapshotscheduleSpec](#snapshotschedulespec)_ |  |  |  |
| `status` _[SnapshotscheduleStatus](#snapshotschedulestatus)_ |  |  |  |


#### SnapshotscheduleInitParameters







_Appears in:_
- [SnapshotscheduleSpec](#snapshotschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `keepSnapshots` _float_ | The amount of Snapshots to keep before overwriting the last created Snapshot |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the snapshot schedule will be run |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `storageUuid` _string_ | UUID of the storage used to create snapshots |  |  |


#### SnapshotscheduleList



SnapshotscheduleList contains a list of Snapshotschedules





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SnapshotscheduleList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Snapshotschedule](#snapshotschedule) array_ |  |  |  |


#### SnapshotscheduleObservation







_Appears in:_
- [SnapshotscheduleStatus](#snapshotschedulestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `keepSnapshots` _float_ | The amount of Snapshots to keep before overwriting the last created Snapshot |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the snapshot schedule will be run |  |  |
| `nextRuntimeComputed` _string_ | The date and time that the snapshot schedule will be run. This date and time is computed by gridscale's server. |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `snapshot` _[SnapshotscheduleSnapshotObservation](#snapshotschedulesnapshotobservation) array_ | Related snashots |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `storageUuid` _string_ | UUID of the storage used to create snapshots |  |  |


#### SnapshotscheduleParameters







_Appears in:_
- [SnapshotscheduleSpec](#snapshotschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `keepSnapshots` _float_ | The amount of Snapshots to keep before overwriting the last created Snapshot |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object |  | Optional: \{\} <br /> |
| `nextRuntime` _string_ | The date and time that the snapshot schedule will be run |  | Optional: \{\} <br /> |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  | Optional: \{\} <br /> |
| `storageUuid` _string_ | UUID of the storage used to create snapshots |  | Optional: \{\} <br /> |




#### SnapshotscheduleSnapshotObservation







_Appears in:_
- [SnapshotscheduleObservation](#snapshotscheduleobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `createTime` _string_ |  |  |  |
| `name` _string_ |  |  |  |
| `objectUuid` _string_ |  |  |  |




#### SnapshotscheduleSpec



SnapshotscheduleSpec defines the desired state of Snapshotschedule



_Appears in:_
- [Snapshotschedule](#snapshotschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[SnapshotscheduleParameters](#snapshotscheduleparameters)_ |  |  |  |
| `initProvider` _[SnapshotscheduleInitParameters](#snapshotscheduleinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SnapshotscheduleStatus



SnapshotscheduleStatus defines the observed state of Snapshotschedule.



_Appears in:_
- [Snapshotschedule](#snapshotschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SnapshotscheduleObservation](#snapshotscheduleobservation)_ |  |  |  |


#### Sqlserver



Sqlserver is the Schema for the Sqlservers API. Manage a MS SQL server service in gridscale.



_Appears in:_
- [SqlserverList](#sqlserverlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Sqlserver` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SqlserverSpec](#sqlserverspec)_ |  |  |  |
| `status` _[SqlserverStatus](#sqlserverstatus)_ |  |  |  |


#### SqlserverInitParameters







_Appears in:_
- [SqlserverSpec](#sqlserverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MS SQL server service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MS SQL Server. |  |  |
| `release` _string_ | The MS SQL server release of this instance. For convenience, please use gscloud to get the list of available MS SQL server service releases.<br />The MS SQL Server release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MS SQL Server releases. |  |  |
| `s3Backup` _[S3BackupInitParameters](#s3backupinitparameters) array_ | Allow backup/restore MS SQL server to/from a S3 bucket.<br />Allow backup/restore MS SQL server to/from a S3 bucket. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MS SQL Server. |  |  |


#### SqlserverList



SqlserverList contains a list of Sqlservers





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SqlserverList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Sqlserver](#sqlserver) array_ |  |  |  |




#### SqlserverListenPortObservation







_Appears in:_
- [SqlserverObservation](#sqlserverobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### SqlserverObservation







_Appears in:_
- [SqlserverStatus](#sqlserverstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[SqlserverListenPortObservation](#sqlserverlistenportobservation) array_ | The port numbers where this MS SQL server service accepts connections.<br />The port numbers where this MS SQL Server accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MS SQL server service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MS SQL Server. |  |  |
| `release` _string_ | The MS SQL server release of this instance. For convenience, please use gscloud to get the list of available MS SQL server service releases.<br />The MS SQL Server release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MS SQL Server releases. |  |  |
| `s3Backup` _[S3BackupObservation](#s3backupobservation) array_ | Allow backup/restore MS SQL server to/from a S3 bucket.<br />Allow backup/restore MS SQL server to/from a S3 bucket. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MS SQL Server. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that MS SQL server service uses.<br />PaaS service template that MS SQL Server uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MS SQL Server. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MS SQL Server is in use. |  |  |


#### SqlserverParameters







_Appears in:_
- [SqlserverSpec](#sqlserverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of MS SQL server service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MS SQL Server. |  | Optional: \{\} <br /> |
| `release` _string_ | The MS SQL server release of this instance. For convenience, please use gscloud to get the list of available MS SQL server service releases.<br />The MS SQL Server release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MS SQL Server releases. |  | Optional: \{\} <br /> |
| `s3Backup` _[S3BackupParameters](#s3backupparameters) array_ | Allow backup/restore MS SQL server to/from a S3 bucket.<br />Allow backup/restore MS SQL server to/from a S3 bucket. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MS SQL Server. |  | Optional: \{\} <br /> |


#### SqlserverSpec



SqlserverSpec defines the desired state of Sqlserver



_Appears in:_
- [Sqlserver](#sqlserver)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[SqlserverParameters](#sqlserverparameters)_ |  |  |  |
| `initProvider` _[SqlserverInitParameters](#sqlserverinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SqlserverStatus



SqlserverStatus defines the observed state of Sqlserver.



_Appears in:_
- [Sqlserver](#sqlserver)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SqlserverObservation](#sqlserverobservation)_ |  |  |  |


#### Sshkey



Sshkey is the Schema for the Sshkeys API. Manages an SSH public key in gridscale.



_Appears in:_
- [SshkeyList](#sshkeylist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Sshkey` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SshkeySpec](#sshkeyspec)_ |  |  |  |
| `status` _[SshkeyStatus](#sshkeystatus)_ |  |  |  |


#### SshkeyInitParameters







_Appears in:_
- [SshkeySpec](#sshkeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sshkey` _string_ | This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).<br />sshkey_string is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1) |  |  |


#### SshkeyList



SshkeyList contains a list of Sshkeys





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SshkeyList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Sshkey](#sshkey) array_ |  |  |  |


#### SshkeyObservation







_Appears in:_
- [SshkeyStatus](#sshkeystatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sshkey` _string_ | This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).<br />sshkey_string is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1) |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |


#### SshkeyParameters







_Appears in:_
- [SshkeySpec](#sshkeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `sshkey` _string_ | This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).<br />sshkey_string is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1) |  | Optional: \{\} <br /> |


#### SshkeySpec



SshkeySpec defines the desired state of Sshkey



_Appears in:_
- [Sshkey](#sshkey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[SshkeyParameters](#sshkeyparameters)_ |  |  |  |
| `initProvider` _[SshkeyInitParameters](#sshkeyinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SshkeyStatus



SshkeyStatus defines the observed state of Sshkey.



_Appears in:_
- [Sshkey](#sshkey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SshkeyObservation](#sshkeyobservation)_ |  |  |  |


#### Storage



Storage is the Schema for the Storages API. Manages a storage in gridscale.



_Appears in:_
- [StorageList](#storagelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Storage` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageSpec](#storagespec)_ |  |  |  |
| `status` _[StorageStatus](#storagestatus)_ |  |  |  |




#### StorageBackupsObservation







_Appears in:_
- [BackupscheduleObservation](#backupscheduleobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `createTime` _string_ |  |  |  |
| `name` _string_ |  |  |  |
| `objectUuid` _string_ |  |  |  |




#### StorageInitParameters







_Appears in:_
- [ServerInitParameters_2](#serverinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `objectUuid` _string_ | The object UUID or id of the storage. |  |  |


#### StorageInitParameters_2







_Appears in:_
- [StorageSpec](#storagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `rollbackFromBackupUuid` _string_ | Rollback the storage from a specific storage backup.<br />Rollback the storage from a specific storage backup. |  |  |
| `storageType` _string_ | (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `storageVariant` _string_ | Storage variant (one of local or distributed). Default: "distributed".<br />Storage variant (one of local or distributed). |  |  |
| `template` _[TemplateInitParameters](#templateinitparameters) array_ | List of labels in the format [ "label1", "label2" ]. |  |  |


#### StorageList



StorageList contains a list of Storages





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Storage](#storage) array_ |  |  |  |


#### StorageObservation







_Appears in:_
- [ServerObservation_2](#serverobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  |  |
| `bus` _float_ | The SCSI bus id. The SCSI defines transmission routes like Serial Attached SCSI (SAS), Fibre Channel and iSCSI. Each SCSI device is addressed via a specific number. Each SCSI bus can have multiple SCSI devices connected to it. |  |  |
| `capacity` _float_ | Capacity of the storage (GB). |  |  |
| `controller` _float_ | Defines the SCSI controller id. The SCSI defines transmission routes such as Serial Attached SCSI (SAS), Fibre Channel and iSCSI. |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `lun` _float_ | Is the common SCSI abbreviation of the Logical Unit Number. A lun is a unique identifier for a single disk or a composite of disks. |  |  |
| `objectName` _string_ | Name of the storage. |  |  |
| `objectUuid` _string_ | The object UUID or id of the storage. |  |  |
| `storageType` _string_ | Indicates the speed of the storage. This may be (storage, storage_high or storage_insane). |  |  |
| `target` _float_ | Defines the SCSI target ID. The target ID is a device (e.g. disk). |  |  |


#### StorageObservation_2







_Appears in:_
- [StorageStatus](#storagestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `locationUuid` _string_ | The location this storage is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `parentUuid` _string_ |  |  |  |
| `rollbackFromBackupUuid` _string_ | Rollback the storage from a specific storage backup.<br />Rollback the storage from a specific storage backup. |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |
| `storageType` _string_ | (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `storageVariant` _string_ | Storage variant (one of local or distributed). Default: "distributed".<br />Storage variant (one of local or distributed). |  |  |
| `template` _[TemplateObservation](#templateobservation) array_ | List of labels in the format [ "label1", "label2" ]. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### StorageParameters







_Appears in:_
- [ServerParameters_2](#serverparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `objectUuid` _string_ | The object UUID or id of the storage. |  | Optional: \{\} <br /> |


#### StorageParameters_2







_Appears in:_
- [StorageSpec](#storagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `rollbackFromBackupUuid` _string_ | Rollback the storage from a specific storage backup.<br />Rollback the storage from a specific storage backup. |  | Optional: \{\} <br /> |
| `storageType` _string_ | (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  | Optional: \{\} <br /> |
| `storageVariant` _string_ | Storage variant (one of local or distributed). Default: "distributed".<br />Storage variant (one of local or distributed). |  | Optional: \{\} <br /> |
| `template` _[TemplateParameters](#templateparameters) array_ | List of labels in the format [ "label1", "label2" ]. |  | Optional: \{\} <br /> |


#### StorageSpec



StorageSpec defines the desired state of Storage



_Appears in:_
- [Storage](#storage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[StorageParameters_2](#storageparameters_2)_ |  |  |  |
| `initProvider` _[StorageInitParameters_2](#storageinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageStatus



StorageStatus defines the observed state of Storage.



_Appears in:_
- [Storage](#storage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageObservation_2](#storageobservation_2)_ |  |  |  |


#### TaintsInitParameters







_Appears in:_
- [NodePoolInitParameters](#nodepoolinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `effect` _string_ | The effect of the taint. |  |  |
| `key` _string_ | The key of the taint. |  |  |
| `value` _string_ | The value of the taint. |  |  |


#### TaintsObservation







_Appears in:_
- [NodePoolObservation](#nodepoolobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `effect` _string_ | The effect of the taint. |  |  |
| `key` _string_ | The key of the taint. |  |  |
| `value` _string_ | The value of the taint. |  |  |


#### TaintsParameters







_Appears in:_
- [NodePoolParameters](#nodepoolparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `effect` _string_ | The effect of the taint. |  | Optional: \{\} <br /> |
| `key` _string_ | The key of the taint. |  | Optional: \{\} <br /> |
| `value` _string_ | The value of the taint. |  | Optional: \{\} <br /> |


#### Template



Template is the Schema for the Templates API. Manages a template in gridscale.



_Appears in:_
- [TemplateList](#templatelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Template` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[TemplateSpec](#templatespec)_ |  |  |  |
| `status` _[TemplateStatus](#templatestatus)_ |  |  |  |


#### TemplateInitParameters







_Appears in:_
- [StorageInitParameters_2](#storageinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostname` _string_ | The hostname of the installed server (ignored for private templates and public windows templates). |  |  |
| `passwordSecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The root (Linux) or Administrator (Windows) password to set for the installed storage. Valid only for public templates. The password has to be either plain-text or a crypt string (modular crypt format - MCF). |  |  |
| `passwordType` _string_ | (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates). |  |  |
| `sshkeys` _string array_ | (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user. |  |  |
| `templateUuid` _string_ | The UUID of a template. This can be found in the the page Template by clicking more on the template or by using a gridscale_template datasource. |  |  |


#### TemplateInitParameters_2







_Appears in:_
- [TemplateSpec](#templatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `name` _string_ | The exact name of the template as show in the page Template.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `snapshotUuid` _string_ | Snapshot uuid for template.<br />Snapshot UUID for template. |  |  |


#### TemplateList



TemplateList contains a list of Templates





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `TemplateList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Template](#template) array_ |  |  |  |


#### TemplateObservation







_Appears in:_
- [StorageObservation_2](#storageobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostname` _string_ | The hostname of the installed server (ignored for private templates and public windows templates). |  |  |
| `passwordType` _string_ | (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates). |  |  |
| `sshkeys` _string array_ | (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user. |  |  |
| `templateUuid` _string_ | The UUID of a template. This can be found in the the page Template by clicking more on the template or by using a gridscale_template datasource. |  |  |


#### TemplateObservation_2







_Appears in:_
- [TemplateStatus](#templatestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The capacity of a storage/ISO Image/template/snapshot in GB.<br />The capacity of a storage/ISO image/template/snapshot in GB. |  |  |
| `changeTime` _string_ | The date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `createTime` _string_ | The date and time the object was initially created.<br />The date and time the object was initially created. |  |  |
| `currentPrice` _float_ | Defines the price for the current period since the last bill.<br />Defines the price for the current period since the last bill. |  |  |
| `description` _string_ | Description of the template.<br />Description of the template. |  |  |
| `distro` _string_ | The OS distribution that the template contains.<br />The Linux distribution of this template (if applicable). |  |  |
| `id` _string_ | The UUID of the template. |  |  |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  |  |
| `name` _string_ | The exact name of the template as show in the page Template.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `ostype` _string_ | The operating system installed in the template.<br />The operating system installed in the template. |  |  |
| `private` _boolean_ | The object is private, the value will be true. Otherwise the value will be false.<br />The object is private, the value will be true. Otherwise the value will be false. |  |  |
| `snapshotUuid` _string_ | Snapshot uuid for template.<br />Snapshot UUID for template. |  |  |
| `status` _string_ | Status indicates the status of the object.<br />Status indicates the status of the object. |  |  |
| `usageInMinutes` _float_ | Total minutes the object has been running.<br />Total minutes the object has been running. |  |  |
| `version` _string_ | The version of the template. |  |  |


#### TemplateParameters







_Appears in:_
- [StorageParameters_2](#storageparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostname` _string_ | The hostname of the installed server (ignored for private templates and public windows templates). |  | Optional: \{\} <br /> |
| `passwordSecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The root (Linux) or Administrator (Windows) password to set for the installed storage. Valid only for public templates. The password has to be either plain-text or a crypt string (modular crypt format - MCF). |  | Optional: \{\} <br /> |
| `passwordType` _string_ | (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates). |  | Optional: \{\} <br /> |
| `sshkeys` _string array_ | (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user. |  | Optional: \{\} <br /> |
| `templateUuid` _string_ | The UUID of a template. This can be found in the the page Template by clicking more on the template or by using a gridscale_template datasource. |  | Optional: \{\} <br /> |


#### TemplateParameters_2







_Appears in:_
- [TemplateSpec](#templatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The exact name of the template as show in the page Template.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `snapshotUuid` _string_ | Snapshot uuid for template.<br />Snapshot UUID for template. |  | Optional: \{\} <br /> |


#### TemplateSpec



TemplateSpec defines the desired state of Template



_Appears in:_
- [Template](#template)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[TemplateParameters_2](#templateparameters_2)_ |  |  |  |
| `initProvider` _[TemplateInitParameters_2](#templateinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### TemplateStatus



TemplateStatus defines the observed state of Template.



_Appears in:_
- [Template](#template)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[TemplateObservation_2](#templateobservation_2)_ |  |  |  |



## gridscale.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [Backupschedule](#backupschedule)
- [BackupscheduleList](#backupschedulelist)
- [Filesystem](#filesystem)
- [FilesystemList](#filesystemlist)
- [Firewall](#firewall)
- [FirewallList](#firewalllist)
- [IPv4](#ipv4)
- [IPv4List](#ipv4list)
- [IPv6](#ipv6)
- [IPv6List](#ipv6list)
- [Isoimage](#isoimage)
- [IsoimageList](#isoimagelist)
- [K8S](#k8s)
- [K8SList](#k8slist)
- [Loadbalancer](#loadbalancer)
- [LoadbalancerList](#loadbalancerlist)
- [Mariadb](#mariadb)
- [MariadbList](#mariadblist)
- [Memcached](#memcached)
- [MemcachedList](#memcachedlist)
- [MySQL](#mysql)
- [MySQLList](#mysqllist)
- [Network](#network)
- [NetworkList](#networklist)
- [Paas](#paas)
- [PaasList](#paaslist)
- [Postgresql](#postgresql)
- [PostgresqlList](#postgresqllist)
- [Server](#server)
- [ServerList](#serverlist)
- [Snapshot](#snapshot)
- [SnapshotList](#snapshotlist)
- [Snapshotschedule](#snapshotschedule)
- [SnapshotscheduleList](#snapshotschedulelist)
- [Sqlserver](#sqlserver)
- [SqlserverList](#sqlserverlist)
- [Sshkey](#sshkey)
- [SshkeyList](#sshkeylist)
- [Storage](#storage)
- [StorageList](#storagelist)
- [Template](#template)
- [TemplateList](#templatelist)





#### AutoAssignedServersObservation







_Appears in:_
- [NetworkObservation_2](#networkobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `ip` _string_ | IP which is assigned to the server. |  |  |
| `serverUuid` _string_ | UUID of the server. |  |  |




#### BackendServerInitParameters







_Appears in:_
- [LoadbalancerInitParameters](#loadbalancerinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | A valid domain or an IP address of a server. |  |  |
| `proxyProtocol` _string_ | The proxy protocol version. The proxy protocol is disabled by default and the valid version is either v1 or v2. |  |  |
| `weight` _float_ | The backend host weight. Default: 100. |  |  |


#### BackendServerObservation







_Appears in:_
- [LoadbalancerObservation](#loadbalancerobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | A valid domain or an IP address of a server. |  |  |
| `proxyProtocol` _string_ | The proxy protocol version. The proxy protocol is disabled by default and the valid version is either v1 or v2. |  |  |
| `weight` _float_ | The backend host weight. Default: 100. |  |  |


#### BackendServerParameters







_Appears in:_
- [LoadbalancerParameters](#loadbalancerparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | A valid domain or an IP address of a server. |  | Optional: \{\} <br /> |
| `proxyProtocol` _string_ | The proxy protocol version. The proxy protocol is disabled by default and the valid version is either v1 or v2. |  | Optional: \{\} <br /> |
| `weight` _float_ | The backend host weight. Default: 100. |  | Optional: \{\} <br /> |


#### Backupschedule



Backupschedule is the Schema for the Backupschedules API. <no value>



_Appears in:_
- [BackupscheduleList](#backupschedulelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Backupschedule` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[BackupscheduleSpec](#backupschedulespec)_ |  |  |  |
| `status` _[BackupscheduleStatus](#backupschedulestatus)_ |  |  |  |


#### BackupscheduleInitParameters







_Appears in:_
- [BackupscheduleSpec](#backupschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `active` _boolean_ | The status of the schedule active or not |  |  |
| `backupLocationUuid` _string_ | UUID of the location where your backup is stored. |  |  |
| `keepBackups` _float_ | The amount of storage backups to keep before overwriting the last created backup |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the storage backup schedule will be run. Format: "2006-01-02 15:04:05" |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `storageUuid` _string_ | UUID of the storage used to create storage backups |  |  |


#### BackupscheduleList



BackupscheduleList contains a list of Backupschedules





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `BackupscheduleList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Backupschedule](#backupschedule) array_ |  |  |  |


#### BackupscheduleObservation







_Appears in:_
- [BackupscheduleStatus](#backupschedulestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `active` _boolean_ | The status of the schedule active or not |  |  |
| `backupLocationName` _string_ | The human-readable name of backup location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `backupLocationUuid` _string_ | UUID of the location where your backup is stored. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `keepBackups` _float_ | The amount of storage backups to keep before overwriting the last created backup |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the storage backup schedule will be run. Format: "2006-01-02 15:04:05" |  |  |
| `nextRuntimeComputed` _string_ | The date and time that the storage backup schedule will be run. This date and time is computed by gridscale's server. |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `storageBackups` _[StorageBackupsObservation](#storagebackupsobservation) array_ | Related backups |  |  |
| `storageUuid` _string_ | UUID of the storage used to create storage backups |  |  |


#### BackupscheduleParameters







_Appears in:_
- [BackupscheduleSpec](#backupschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `active` _boolean_ | The status of the schedule active or not |  | Optional: \{\} <br /> |
| `backupLocationUuid` _string_ | UUID of the location where your backup is stored. |  | Optional: \{\} <br /> |
| `keepBackups` _float_ | The amount of storage backups to keep before overwriting the last created backup |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object |  | Optional: \{\} <br /> |
| `nextRuntime` _string_ | The date and time that the storage backup schedule will be run. Format: "2006-01-02 15:04:05" |  | Optional: \{\} <br /> |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  | Optional: \{\} <br /> |
| `storageUuid` _string_ | UUID of the storage used to create storage backups |  | Optional: \{\} <br /> |


#### BackupscheduleSpec



BackupscheduleSpec defines the desired state of Backupschedule



_Appears in:_
- [Backupschedule](#backupschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[BackupscheduleParameters](#backupscheduleparameters)_ |  |  |  |
| `initProvider` _[BackupscheduleInitParameters](#backupscheduleinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### BackupscheduleStatus



BackupscheduleStatus defines the observed state of Backupschedule.



_Appears in:_
- [Backupschedule](#backupschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[BackupscheduleObservation](#backupscheduleobservation)_ |  |  |  |


#### Filesystem



Filesystem is the Schema for the Filesystems API. <no value>



_Appears in:_
- [FilesystemList](#filesystemlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Filesystem` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[FilesystemSpec](#filesystemspec)_ |  |  |  |
| `status` _[FilesystemStatus](#filesystemstatus)_ |  |  |  |


#### FilesystemInitParameters







_Appears in:_
- [FilesystemSpec](#filesystemspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowedIpRanges` _string array_ | Allowed CIDR block or IP address in CIDR notation. |  |  |
| `anonGid` _float_ | Target group id when root squash is active. |  |  |
| `anonUid` _float_ | Target user id when root squash is active. |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Filesystem service. |  |  |
| `release` _string_ | The Filesystem service release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Filesystem service releases. |  |  |
| `rootSquash` _boolean_ | Map root user/group ownership to anon_uid/anon_gid |  |  |
| `securityZoneUuid` _string_ | Security zone UUID linked to Filesystem service. |  |  |


#### FilesystemList



FilesystemList contains a list of Filesystems





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `FilesystemList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Filesystem](#filesystem) array_ |  |  |  |


#### FilesystemObservation







_Appears in:_
- [FilesystemStatus](#filesystemstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowedIpRanges` _string array_ | Allowed CIDR block or IP address in CIDR notation. |  |  |
| `anonGid` _float_ | Target group id when root squash is active. |  |  |
| `anonUid` _float_ | Target user id when root squash is active. |  |  |
| `changeTime` _string_ | Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `listenPort` _[ListenPortObservation](#listenportobservation) array_ | The port numbers where this Filesystem service accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Filesystem service. |  |  |
| `release` _string_ | The Filesystem service release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Filesystem service releases. |  |  |
| `rootSquash` _boolean_ | Map root user/group ownership to anon_uid/anon_gid |  |  |
| `securityZoneUuid` _string_ | Security zone UUID linked to Filesystem service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Filesystem service uses. |  |  |
| `status` _string_ | Current status of Filesystem service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that Filesystem service is in use. |  |  |


#### FilesystemParameters







_Appears in:_
- [FilesystemSpec](#filesystemspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `allowedIpRanges` _string array_ | Allowed CIDR block or IP address in CIDR notation. |  | Optional: \{\} <br /> |
| `anonGid` _float_ | Target group id when root squash is active. |  | Optional: \{\} <br /> |
| `anonUid` _float_ | Target user id when root squash is active. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Filesystem service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Filesystem service release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Filesystem service releases. |  | Optional: \{\} <br /> |
| `rootSquash` _boolean_ | Map root user/group ownership to anon_uid/anon_gid |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | Security zone UUID linked to Filesystem service. |  | Optional: \{\} <br /> |


#### FilesystemSpec



FilesystemSpec defines the desired state of Filesystem



_Appears in:_
- [Filesystem](#filesystem)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[FilesystemParameters](#filesystemparameters)_ |  |  |  |
| `initProvider` _[FilesystemInitParameters](#filesysteminitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### FilesystemStatus



FilesystemStatus defines the observed state of Filesystem.



_Appears in:_
- [Filesystem](#filesystem)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[FilesystemObservation](#filesystemobservation)_ |  |  |  |


#### Firewall



Firewall is the Schema for the Firewalls API. Manages a firewall in gridscale.



_Appears in:_
- [FirewallList](#firewalllist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Firewall` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[FirewallSpec](#firewallspec)_ |  |  |  |
| `status` _[FirewallStatus](#firewallstatus)_ |  |  |  |


#### FirewallInitParameters







_Appears in:_
- [FirewallSpec](#firewallspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `rulesV4In` _[RulesV4InInitParameters](#rulesv4ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[RulesV4OutInitParameters](#rulesv4outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[RulesV6InInitParameters](#rulesv6ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[RulesV6OutInitParameters](#rulesv6outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |


#### FirewallList



FirewallList contains a list of Firewalls





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `FirewallList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Firewall](#firewall) array_ |  |  |  |


#### FirewallObservation







_Appears in:_
- [FirewallStatus](#firewallstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | The date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `createTime` _string_ | The date and time the object was initially created.<br />The date and time the object was initially created. |  |  |
| `description` _string_ | Description of the firewall.<br />Description of the Firewall. |  |  |
| `id` _string_ | The UUID of the firewall. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationName` _string_ | The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `network` _[NetworkObservation](#networkobservation) array_ | The information about networks which are related to this firewall. |  |  |
| `private` _boolean_ | The object is private, the value will be true. Otherwise the value will be false.<br />The object is private, the value will be true. Otherwise the value will be false. |  |  |
| `rulesV4In` _[RulesV4InObservation](#rulesv4inobservation) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[RulesV4OutObservation](#rulesv4outobservation) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[RulesV6InObservation](#rulesv6inobservation) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[RulesV6OutObservation](#rulesv6outobservation) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |
| `status` _string_ | Status indicates the status of the object.<br />Status indicates the status of the object |  |  |


#### FirewallParameters







_Appears in:_
- [FirewallSpec](#firewallspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `rulesV4In` _[RulesV4InParameters](#rulesv4inparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV4Out` _[RulesV4OutParameters](#rulesv4outparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV6In` _[RulesV6InParameters](#rulesv6inparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |
| `rulesV6Out` _[RulesV6OutParameters](#rulesv6outparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |


#### FirewallSpec



FirewallSpec defines the desired state of Firewall



_Appears in:_
- [Firewall](#firewall)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[FirewallParameters](#firewallparameters)_ |  |  |  |
| `initProvider` _[FirewallInitParameters](#firewallinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### FirewallStatus



FirewallStatus defines the observed state of Firewall.



_Appears in:_
- [Firewall](#firewall)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[FirewallObservation](#firewallobservation)_ |  |  |  |


#### ForwardingRuleInitParameters







_Appears in:_
- [LoadbalancerInitParameters](#loadbalancerinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateUuid` _string_ | The UUID of a custom certificate.<br />The UUID of a custom certificate. |  |  |
| `letsencryptSsl` _string_ | A valid domain name that points to the loadbalancer's IP address.<br />A valid domain name that points to the loadbalancer's IP address. |  |  |
| `listenPort` _float_ | Specifies the entry port of the load balancer.<br />Specifies the entry port of the load balancer. |  |  |
| `mode` _string_ | Supports HTTP and TCP mode. Valid values: http, tcp.<br />Supports HTTP and TCP mode. Valid values: http, tcp. |  |  |
| `targetPort` _float_ | Specifies the exit port that the load balancer uses to forward the traffic to the backend server.<br />Specifies the exit port that the load balancer uses to forward the traffic to the backend server. |  |  |


#### ForwardingRuleObservation







_Appears in:_
- [LoadbalancerObservation](#loadbalancerobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateUuid` _string_ | The UUID of a custom certificate.<br />The UUID of a custom certificate. |  |  |
| `letsencryptSsl` _string_ | A valid domain name that points to the loadbalancer's IP address.<br />A valid domain name that points to the loadbalancer's IP address. |  |  |
| `listenPort` _float_ | Specifies the entry port of the load balancer.<br />Specifies the entry port of the load balancer. |  |  |
| `mode` _string_ | Supports HTTP and TCP mode. Valid values: http, tcp.<br />Supports HTTP and TCP mode. Valid values: http, tcp. |  |  |
| `targetPort` _float_ | Specifies the exit port that the load balancer uses to forward the traffic to the backend server.<br />Specifies the exit port that the load balancer uses to forward the traffic to the backend server. |  |  |


#### ForwardingRuleParameters







_Appears in:_
- [LoadbalancerParameters](#loadbalancerparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateUuid` _string_ | The UUID of a custom certificate.<br />The UUID of a custom certificate. |  | Optional: \{\} <br /> |
| `letsencryptSsl` _string_ | A valid domain name that points to the loadbalancer's IP address.<br />A valid domain name that points to the loadbalancer's IP address. |  | Optional: \{\} <br /> |
| `listenPort` _float_ | Specifies the entry port of the load balancer.<br />Specifies the entry port of the load balancer. |  | Optional: \{\} <br /> |
| `mode` _string_ | Supports HTTP and TCP mode. Valid values: http, tcp.<br />Supports HTTP and TCP mode. Valid values: http, tcp. |  | Optional: \{\} <br /> |
| `targetPort` _float_ | Specifies the exit port that the load balancer uses to forward the traffic to the backend server.<br />Specifies the exit port that the load balancer uses to forward the traffic to the backend server. |  | Optional: \{\} <br /> |


#### HardwareProfileConfigInitParameters







_Appears in:_
- [ServerInitParameters_2](#serverinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hypervExtensions` _boolean_ | Boolean. |  |  |
| `machinetype` _string_ | Allowed values: "i440fx", "q35_bios", "q35_uefi". |  |  |
| `nestedVirtualization` _boolean_ | Boolean. |  |  |
| `networkModel` _string_ | Allowed values: "e1000", "e1000e", "virtio", "vmxnet3". |  |  |
| `serialInterface` _boolean_ | Boolean. |  |  |
| `serverRenice` _boolean_ | Boolean. |  |  |
| `storageDevice` _string_ | Allowed values: "ide", "sata", "virtio_scsi", "virtio_block". |  |  |
| `usbController` _string_ | Allowed values: "nec_xhci", "piix3_uhci". |  |  |


#### HardwareProfileConfigObservation







_Appears in:_
- [ServerObservation_2](#serverobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hypervExtensions` _boolean_ | Boolean. |  |  |
| `machinetype` _string_ | Allowed values: "i440fx", "q35_bios", "q35_uefi". |  |  |
| `nestedVirtualization` _boolean_ | Boolean. |  |  |
| `networkModel` _string_ | Allowed values: "e1000", "e1000e", "virtio", "vmxnet3". |  |  |
| `serialInterface` _boolean_ | Boolean. |  |  |
| `serverRenice` _boolean_ | Boolean. |  |  |
| `storageDevice` _string_ | Allowed values: "ide", "sata", "virtio_scsi", "virtio_block". |  |  |
| `usbController` _string_ | Allowed values: "nec_xhci", "piix3_uhci". |  |  |


#### HardwareProfileConfigParameters







_Appears in:_
- [ServerParameters_2](#serverparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hypervExtensions` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `machinetype` _string_ | Allowed values: "i440fx", "q35_bios", "q35_uefi". |  | Optional: \{\} <br /> |
| `nestedVirtualization` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `networkModel` _string_ | Allowed values: "e1000", "e1000e", "virtio", "vmxnet3". |  | Optional: \{\} <br /> |
| `serialInterface` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `serverRenice` _boolean_ | Boolean. |  | Optional: \{\} <br /> |
| `storageDevice` _string_ | Allowed values: "ide", "sata", "virtio_scsi", "virtio_block". |  | Optional: \{\} <br /> |
| `usbController` _string_ | Allowed values: "nec_xhci", "piix3_uhci". |  | Optional: \{\} <br /> |


#### IPv4



IPv4 is the Schema for the IPv4s API. Manages an IPv4 address in gridscale.



_Appears in:_
- [IPv4List](#ipv4list)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv4` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[IPv4Spec](#ipv4spec)_ |  |  |  |
| `status` _[IPv4Status](#ipv4status)_ |  |  |  |


#### IPv4InitParameters







_Appears in:_
- [IPv4Spec](#ipv4spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |


#### IPv4List



IPv4List contains a list of IPv4s





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv4List` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[IPv4](#ipv4) array_ |  |  |  |


#### IPv4Observation







_Appears in:_
- [IPv4Status](#ipv4status)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill.<br />Defines the price for the current period since the last bill. |  |  |
| `deleteBlock` _boolean_ | Defines if the object is administratively blocked. If true, it can not be deleted by the user.<br />Defines if the object is administratively blocked. If true, it can not be deleted by the user. |  |  |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `id` _string_ |  |  |  |
| `ip` _string_ | Defines the IP address.<br />Defines the IP address. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | See Argument Reference above.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `prefix` _string_ | The network address and the subnet. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### IPv4Parameters







_Appears in:_
- [IPv4Spec](#ipv4spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  | Optional: \{\} <br /> |


#### IPv4Spec



IPv4Spec defines the desired state of IPv4



_Appears in:_
- [IPv4](#ipv4)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[IPv4Parameters](#ipv4parameters)_ |  |  |  |
| `initProvider` _[IPv4InitParameters](#ipv4initparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### IPv4Status



IPv4Status defines the observed state of IPv4.



_Appears in:_
- [IPv4](#ipv4)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[IPv4Observation](#ipv4observation)_ |  |  |  |


#### IPv6



IPv6 is the Schema for the IPv6s API. Manages an IPv6 address in gridscale.



_Appears in:_
- [IPv6List](#ipv6list)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv6` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[IPv6Spec](#ipv6spec)_ |  |  |  |
| `status` _[IPv6Status](#ipv6status)_ |  |  |  |


#### IPv6InitParameters







_Appears in:_
- [IPv6Spec](#ipv6spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |


#### IPv6List



IPv6List contains a list of IPv6s





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IPv6List` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[IPv6](#ipv6) array_ |  |  |  |


#### IPv6Observation







_Appears in:_
- [IPv6Status](#ipv6status)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill.<br />Defines the price for the current period since the last bill. |  |  |
| `deleteBlock` _boolean_ | Defines if the object is administratively blocked. If true, it can not be deleted by the user.<br />Defines if the object is administratively blocked. If true, it can not be deleted by the user. |  |  |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  |  |
| `id` _string_ |  |  |  |
| `ip` _string_ | Defines the IP address.<br />Defines the IP address. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this resource is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `prefix` _string_ | The network address and the subnet. |  |  |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### IPv6Parameters







_Appears in:_
- [IPv6Spec](#ipv6spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `failover` _boolean_ | Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server.<br />Sets failover mode for this IP. If true, then this IP is no longer available for DHCP and can no longer be related to any server. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `reverseDns` _string_ | Defines the reverse DNS entry for the IP address (PTR Resource Record).<br />Defines the reverse DNS entry for the IP address (PTR Resource Record). |  | Optional: \{\} <br /> |


#### IPv6Spec



IPv6Spec defines the desired state of IPv6



_Appears in:_
- [IPv6](#ipv6)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[IPv6Parameters](#ipv6parameters)_ |  |  |  |
| `initProvider` _[IPv6InitParameters](#ipv6initparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### IPv6Status



IPv6Status defines the observed state of IPv6.



_Appears in:_
- [IPv6](#ipv6)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[IPv6Observation](#ipv6observation)_ |  |  |  |


#### Isoimage



Isoimage is the Schema for the Isoimages API. <no value>



_Appears in:_
- [IsoimageList](#isoimagelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Isoimage` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[IsoimageSpec](#isoimagespec)_ |  |  |  |
| `status` _[IsoimageStatus](#isoimagestatus)_ |  |  |  |


#### IsoimageInitParameters







_Appears in:_
- [IsoimageSpec](#isoimagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sourceUrl` _string_ | Contains the source URL of the ISO image that it was originally fetched from. |  |  |


#### IsoimageList



IsoimageList contains a list of Isoimages





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `IsoimageList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Isoimage](#isoimage) array_ |  |  |  |


#### IsoimageObservation







_Appears in:_
- [IsoimageStatus](#isoimagestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The capacity of a storage/ISO image/template/snapshot in GB. |  |  |
| `changeTime` _string_ | The date and time of the last object change. |  |  |
| `createTime` _string_ | The date and time the object was initially created. |  |  |
| `currentPrice` _float_ | Defines the price for the current period since the last bill. |  |  |
| `description` _string_ | Description of the ISO image. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `private` _boolean_ | The object is private, the value will be true. Otherwise the value will be false. |  |  |
| `server` _[ServerObservation](#serverobservation) array_ | The information about servers which are related to this ISO image. |  |  |
| `sourceUrl` _string_ | Contains the source URL of the ISO image that it was originally fetched from. |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `usageInMinutes` _float_ | Total minutes the object has been running. |  |  |
| `version` _string_ | Upstream version of the ISO image release |  |  |


#### IsoimageParameters







_Appears in:_
- [IsoimageSpec](#isoimagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `sourceUrl` _string_ | Contains the source URL of the ISO image that it was originally fetched from. |  | Optional: \{\} <br /> |


#### IsoimageSpec



IsoimageSpec defines the desired state of Isoimage



_Appears in:_
- [Isoimage](#isoimage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[IsoimageParameters](#isoimageparameters)_ |  |  |  |
| `initProvider` _[IsoimageInitParameters](#isoimageinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### IsoimageStatus



IsoimageStatus defines the observed state of Isoimage.



_Appears in:_
- [Isoimage](#isoimage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[IsoimageObservation](#isoimageobservation)_ |  |  |  |


#### K8S



K8S is the Schema for the K8Ss API. Manages a k8s cluster in gridscale.



_Appears in:_
- [K8SList](#k8slist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `K8S` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[K8SSpec](#k8sspec)_ |  |  |  |
| `status` _[K8SStatus](#k8sstatus)_ |  |  |  |


#### K8SInitParameters







_Appears in:_
- [K8SSpec](#k8sspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `auditLogEnabled` _boolean_ | Enable Kubernetes audit logs. |  |  |
| `auditLogLevel` _string_ | Audit log level. |  |  |
| `clusterCidr` _string_ | (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).<br />The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16" |  |  |
| `clusterTrafficEncryption` _boolean_ | Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.<br />Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false. |  |  |
| `gskVersion` _string_ | The gridscale's Kubernetes version of this instance (e.g. "1.30.3-gs0"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.<br />The gridscale k8s PaaS version (issued by gridscale) of this instance. |  |  |
| `k8sHubble` _boolean_ | Enable Hubble for the k8s cluster.<br />Enables Hubble Integration. |  |  |
| `kubeApiserverLogEnabled` _boolean_ | Enable kube-apiserver logs. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `logDelivery` _boolean_ | Enable control plane log delivery. |  |  |
| `logDeliveryAccessKey` _string_ | Access key used to authenticate against Object Storage endpoint. |  |  |
| `logDeliveryBucket` _string_ | Bucket to upload logs to. |  |  |
| `logDeliveryEndpoint` _string_ | Object Storage endpoint URL the bucket is located on. |  |  |
| `logDeliveryInterval` _float_ | Time interval (in min), at which log files will be delivered, unless file size limit is reached first. |  |  |
| `logDeliverySecretKey` _string_ | Secret key used to authenticate against Object Storage endpoint. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `nodePool` _[NodePoolInitParameters](#nodepoolinitparameters) array_ | The collection of node pool specifications. Mutiple node pools can be defined with multiple node_pool blocks. The node pool block supports the following arguments:<br />Define a node pool and its attributes. |  |  |
| `oidcCaPem` _string_ | Custom CA from customer in pem format as string.<br />Custom CA from customer in pem format as string. |  |  |
| `oidcClientId` _string_ | A client ID that all tokens must be issued for.<br />A client ID that all tokens must be issued for. |  |  |
| `oidcEnabled` _boolean_ | Enable OIDC for the k8s cluster.<br />Disable or enable OIDC |  |  |
| `oidcGroupsClaim` _string_ | JWT claim to use as the user's group.<br />JWT claim to use as the user's group. |  |  |
| `oidcGroupsPrefix` _string_ | Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.<br />Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra. |  |  |
| `oidcIssuerUrl` _string_ | URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.<br />URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted. |  |  |
| `oidcRequiredClaim` _string_ | A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.<br />A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2 |  |  |
| `oidcSigningAlgs` _string_ | The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.<br />The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'. |  |  |
| `oidcUsernameClaim` _string_ | JWT claim to use as the user name.<br />JWT claim to use as the user name. |  |  |
| `oidcUsernamePrefix` _string_ | Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.<br />Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. |  |  |
| `release` _string_ | The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.<br />The k8s release of this instance. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.<br />Security zone UUID linked to PaaS service. |  |  |
| `surgeNode` _boolean_ | Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).<br />Enable surge node to avoid resources shortage during the cluster upgrade. |  |  |


#### K8SList



K8SList contains a list of K8Ss





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `K8SList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[K8S](#k8s) array_ |  |  |  |




#### K8SListenPortObservation







_Appears in:_
- [K8SObservation](#k8sobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ | Name of the node pool. |  |  |
| `port` _float_ |  |  |  |




#### K8SObservation







_Appears in:_
- [K8SStatus](#k8sstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `auditLogEnabled` _boolean_ | Enable Kubernetes audit logs. |  |  |
| `auditLogLevel` _string_ | Audit log level. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Time of the last change |  |  |
| `clusterCidr` _string_ | (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).<br />The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16" |  |  |
| `clusterTrafficEncryption` _boolean_ | Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.<br />Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false. |  |  |
| `createTime` _string_ | The time the object was created.<br />Time this service was created. |  |  |
| `gskVersion` _string_ | The gridscale's Kubernetes version of this instance (e.g. "1.30.3-gs0"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.<br />The gridscale k8s PaaS version (issued by gridscale) of this instance. |  |  |
| `id` _string_ |  |  |  |
| `k8sHubble` _boolean_ | Enable Hubble for the k8s cluster.<br />Enables Hubble Integration. |  |  |
| `k8sPrivateNetworkUuid` _string_ | Private network UUID which k8s nodes are attached to. It can be used to attach other PaaS/VMs.<br />Private network UUID which k8s nodes are attached to. It can be used to attach other PaaS/VMs. |  |  |
| `kubeApiserverLogEnabled` _boolean_ | Enable kube-apiserver logs. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[K8SListenPortObservation](#k8slistenportobservation) array_ | The port number where this k8s service accepts connections. |  |  |
| `logDelivery` _boolean_ | Enable control plane log delivery. |  |  |
| `logDeliveryAccessKey` _string_ | Access key used to authenticate against Object Storage endpoint. |  |  |
| `logDeliveryBucket` _string_ | Bucket to upload logs to. |  |  |
| `logDeliveryEndpoint` _string_ | Object Storage endpoint URL the bucket is located on. |  |  |
| `logDeliveryInterval` _float_ | Time interval (in min), at which log files will be delivered, unless file size limit is reached first. |  |  |
| `logDeliverySecretKey` _string_ | Secret key used to authenticate against Object Storage endpoint. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | DEPRECATED  Network UUID containing security zone, which is linked to the k8s cluster.<br />Network UUID containing security zone |  |  |
| `nodePool` _[NodePoolObservation](#nodepoolobservation) array_ | The collection of node pool specifications. Mutiple node pools can be defined with multiple node_pool blocks. The node pool block supports the following arguments:<br />Define a node pool and its attributes. |  |  |
| `oidcCaPem` _string_ | Custom CA from customer in pem format as string.<br />Custom CA from customer in pem format as string. |  |  |
| `oidcClientId` _string_ | A client ID that all tokens must be issued for.<br />A client ID that all tokens must be issued for. |  |  |
| `oidcEnabled` _boolean_ | Enable OIDC for the k8s cluster.<br />Disable or enable OIDC |  |  |
| `oidcGroupsClaim` _string_ | JWT claim to use as the user's group.<br />JWT claim to use as the user's group. |  |  |
| `oidcGroupsPrefix` _string_ | Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.<br />Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra. |  |  |
| `oidcIssuerUrl` _string_ | URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.<br />URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted. |  |  |
| `oidcRequiredClaim` _string_ | A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.<br />A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2 |  |  |
| `oidcSigningAlgs` _string_ | The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.<br />The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'. |  |  |
| `oidcUsernameClaim` _string_ | JWT claim to use as the user name.<br />JWT claim to use as the user name. |  |  |
| `oidcUsernamePrefix` _string_ | Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.<br />Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. |  |  |
| `release` _string_ | The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.<br />The k8s release of this instance. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.<br />Security zone UUID linked to PaaS service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that k8s service uses.g. the k8s service is upgraded by gridscale staffs).<br />PaaS service template identifier for this service. |  |  |
| `status` _string_ | status indicates the status of the object.<br />Current status of the service |  |  |
| `surgeNode` _boolean_ | Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).<br />Enable surge node to avoid resources shortage during the cluster upgrade. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use.<br />Number of minutes that PaaS service is in use |  |  |


#### K8SParameters







_Appears in:_
- [K8SSpec](#k8sspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `auditLogEnabled` _boolean_ | Enable Kubernetes audit logs. |  | Optional: \{\} <br /> |
| `auditLogLevel` _string_ | Audit log level. |  | Optional: \{\} <br /> |
| `clusterCidr` _string_ | (Immutable) The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If the cluster CIDR is not set, the cluster will use "10.244.0.0/16" as it default (even though the cluster_cidr in the k8s resource is empty).<br />The cluster CIDR that will be used to generate the CIDR of nodes, services, and pods. The allowed CIDR prefix length is /16. If this field is empty, the default value is "10.244.0.0/16" |  | Optional: \{\} <br /> |
| `clusterTrafficEncryption` _boolean_ | Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false.<br />Enables cluster encryption via wireguard if true. Only available for GSK version 1.29 and above. Default is false. |  | Optional: \{\} <br /> |
| `gskVersion` _string_ | The gridscale's Kubernetes version of this instance (e.g. "1.30.3-gs0"). Define which gridscale k8s version will be used to create the cluster. For convenience, please use gscloud to get the list of available gridscale k8s version. NOTE: Either gsk_version or release is set at a time.<br />The gridscale k8s PaaS version (issued by gridscale) of this instance. |  | Optional: \{\} <br /> |
| `k8sHubble` _boolean_ | Enable Hubble for the k8s cluster.<br />Enables Hubble Integration. |  | Optional: \{\} <br /> |
| `kubeApiserverLogEnabled` _boolean_ | Enable kube-apiserver logs. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `logDelivery` _boolean_ | Enable control plane log delivery. |  | Optional: \{\} <br /> |
| `logDeliveryAccessKey` _string_ | Access key used to authenticate against Object Storage endpoint. |  | Optional: \{\} <br /> |
| `logDeliveryBucket` _string_ | Bucket to upload logs to. |  | Optional: \{\} <br /> |
| `logDeliveryEndpoint` _string_ | Object Storage endpoint URL the bucket is located on. |  | Optional: \{\} <br /> |
| `logDeliveryInterval` _float_ | Time interval (in min), at which log files will be delivered, unless file size limit is reached first. |  | Optional: \{\} <br /> |
| `logDeliverySecretKey` _string_ | Secret key used to authenticate against Object Storage endpoint. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `nodePool` _[NodePoolParameters](#nodepoolparameters) array_ | The collection of node pool specifications. Mutiple node pools can be defined with multiple node_pool blocks. The node pool block supports the following arguments:<br />Define a node pool and its attributes. |  | Optional: \{\} <br /> |
| `oidcCaPem` _string_ | Custom CA from customer in pem format as string.<br />Custom CA from customer in pem format as string. |  | Optional: \{\} <br /> |
| `oidcClientId` _string_ | A client ID that all tokens must be issued for.<br />A client ID that all tokens must be issued for. |  | Optional: \{\} <br /> |
| `oidcEnabled` _boolean_ | Enable OIDC for the k8s cluster.<br />Disable or enable OIDC |  | Optional: \{\} <br /> |
| `oidcGroupsClaim` _string_ | JWT claim to use as the user's group.<br />JWT claim to use as the user's group. |  | Optional: \{\} <br /> |
| `oidcGroupsPrefix` _string_ | Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.<br />Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra. |  | Optional: \{\} <br /> |
| `oidcIssuerUrl` _string_ | URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted.<br />URL of the provider that allows the API server to discover public signing keys. Only URLs that use the https:// scheme are accepted. |  | Optional: \{\} <br /> |
| `oidcRequiredClaim` _string_ | A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2.<br />A key=value pair that describes a required claim in the ID Token. Multiple claims can be set like this: key1=value1,key2=value2 |  | Optional: \{\} <br /> |
| `oidcSigningAlgs` _string_ | The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'.<br />The signing algorithms accepted. Default is 'RS256'. Other option is 'RS512'. |  | Optional: \{\} <br /> |
| `oidcUsernameClaim` _string_ | JWT claim to use as the user name.<br />JWT claim to use as the user name. |  | Optional: \{\} <br /> |
| `oidcUsernamePrefix` _string_ | Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing.<br />Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this flag isn't provided and --oidc-username-claim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of --oidc-issuer-url. The value - can be used to disable all prefixing. |  | Optional: \{\} <br /> |
| `release` _string_ | The Kubernetes release of this instance. Define which release will be used to create the cluster. For convenience, please use gscloud to get the list of available releases. NOTE: Either gsk_version or release is set at a time.<br />The k8s release of this instance. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  Security zone UUID linked to the Kubernetes resource. If security_zone_uuid is not set, the default security zone will be created (if it doesn't exist) and linked. A change of this argument necessitates the re-creation of the resource.<br />Security zone UUID linked to PaaS service. |  | Optional: \{\} <br /> |
| `surgeNode` _boolean_ | Enable surge node to avoid resources shortage during the cluster upgrade (Default: true).<br />Enable surge node to avoid resources shortage during the cluster upgrade. |  | Optional: \{\} <br /> |


#### K8SSpec



K8SSpec defines the desired state of K8S



_Appears in:_
- [K8S](#k8s)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[K8SParameters](#k8sparameters)_ |  |  |  |
| `initProvider` _[K8SInitParameters](#k8sinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### K8SStatus



K8SStatus defines the observed state of K8S.



_Appears in:_
- [K8S](#k8s)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[K8SObservation](#k8sobservation)_ |  |  |  |


#### LabelsInitParameters







_Appears in:_
- [NodePoolInitParameters](#nodepoolinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `key` _string_ | The key of the label. |  |  |
| `value` _string_ | The value of the label. |  |  |


#### LabelsObservation







_Appears in:_
- [NodePoolObservation](#nodepoolobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `key` _string_ | The key of the label. |  |  |
| `value` _string_ | The value of the label. |  |  |


#### LabelsParameters







_Appears in:_
- [NodePoolParameters](#nodepoolparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `key` _string_ | The key of the label. |  | Optional: \{\} <br /> |
| `value` _string_ | The value of the label. |  | Optional: \{\} <br /> |




#### ListenPortObservation







_Appears in:_
- [FilesystemObservation](#filesystemobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ |  |  |  |
| `name` _string_ |  |  |  |
| `port` _float_ |  |  |  |




#### Loadbalancer



Loadbalancer is the Schema for the Loadbalancers API. Manage a loadbalancer in gridscale.



_Appears in:_
- [LoadbalancerList](#loadbalancerlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Loadbalancer` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[LoadbalancerSpec](#loadbalancerspec)_ |  |  |  |
| `status` _[LoadbalancerStatus](#loadbalancerstatus)_ |  |  |  |


#### LoadbalancerInitParameters







_Appears in:_
- [LoadbalancerSpec](#loadbalancerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `algorithm` _string_ | The algorithm used to process requests. Accepted values: roundrobin/leastconn.<br />The algorithm used to process requests. Accepted values: roundrobin/leastconn. |  |  |
| `backendServer` _[BackendServerInitParameters](#backendserverinitparameters) array_ | The servers that the load balancer can communicate with.<br />List of backend servers. |  |  |
| `forwardingRule` _[ForwardingRuleInitParameters](#forwardingruleinitparameters) array_ | The forwarding rules of the load balancer.<br />List of forwarding rules for the Load balancer. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenIpv4Uuid` _string_ | The UUID of the IPv4 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv4 address the Load balancer will listen to for incoming requests. |  |  |
| `listenIpv6Uuid` _string_ | The UUID of the IPv6 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv6 address the Load balancer will listen to for incoming requests. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `redirectHttpToHttps` _boolean_ | Whether the load balancer is forced to redirect requests from HTTP to HTTPS.<br />Whether the Load balancer is forced to redirect requests from HTTP to HTTPS |  |  |
| `status` _string_ | The status of the load balancer.<br />Status indicates the status of the object. |  |  |


#### LoadbalancerList



LoadbalancerList contains a list of Loadbalancers





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `LoadbalancerList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Loadbalancer](#loadbalancer) array_ |  |  |  |


#### LoadbalancerObservation







_Appears in:_
- [LoadbalancerStatus](#loadbalancerstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `algorithm` _string_ | The algorithm used to process requests. Accepted values: roundrobin/leastconn.<br />The algorithm used to process requests. Accepted values: roundrobin/leastconn. |  |  |
| `backendServer` _[BackendServerObservation](#backendserverobservation) array_ | The servers that the load balancer can communicate with.<br />List of backend servers. |  |  |
| `forwardingRule` _[ForwardingRuleObservation](#forwardingruleobservation) array_ | The forwarding rules of the load balancer.<br />List of forwarding rules for the Load balancer. |  |  |
| `id` _string_ | The UUID of the load balancer. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenIpv4Uuid` _string_ | The UUID of the IPv4 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv4 address the Load balancer will listen to for incoming requests. |  |  |
| `listenIpv6Uuid` _string_ | The UUID of the IPv6 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv6 address the Load balancer will listen to for incoming requests. |  |  |
| `locationUuid` _string_ | The location this load balancer is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `redirectHttpToHttps` _boolean_ | Whether the load balancer is forced to redirect requests from HTTP to HTTPS.<br />Whether the Load balancer is forced to redirect requests from HTTP to HTTPS |  |  |
| `status` _string_ | The status of the load balancer.<br />Status indicates the status of the object. |  |  |


#### LoadbalancerParameters







_Appears in:_
- [LoadbalancerSpec](#loadbalancerspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `algorithm` _string_ | The algorithm used to process requests. Accepted values: roundrobin/leastconn.<br />The algorithm used to process requests. Accepted values: roundrobin/leastconn. |  | Optional: \{\} <br /> |
| `backendServer` _[BackendServerParameters](#backendserverparameters) array_ | The servers that the load balancer can communicate with.<br />List of backend servers. |  | Optional: \{\} <br /> |
| `forwardingRule` _[ForwardingRuleParameters](#forwardingruleparameters) array_ | The forwarding rules of the load balancer.<br />List of forwarding rules for the Load balancer. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `listenIpv4Uuid` _string_ | The UUID of the IPv4 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv4 address the Load balancer will listen to for incoming requests. |  | Optional: \{\} <br /> |
| `listenIpv6Uuid` _string_ | The UUID of the IPv6 address the load balancer will listen to for incoming requests.<br />The UUID of the IPv6 address the Load balancer will listen to for incoming requests. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `redirectHttpToHttps` _boolean_ | Whether the load balancer is forced to redirect requests from HTTP to HTTPS.<br />Whether the Load balancer is forced to redirect requests from HTTP to HTTPS |  | Optional: \{\} <br /> |
| `status` _string_ | The status of the load balancer.<br />Status indicates the status of the object. |  | Optional: \{\} <br /> |


#### LoadbalancerSpec



LoadbalancerSpec defines the desired state of Loadbalancer



_Appears in:_
- [Loadbalancer](#loadbalancer)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[LoadbalancerParameters](#loadbalancerparameters)_ |  |  |  |
| `initProvider` _[LoadbalancerInitParameters](#loadbalancerinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### LoadbalancerStatus



LoadbalancerStatus defines the observed state of Loadbalancer.



_Appears in:_
- [Loadbalancer](#loadbalancer)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[LoadbalancerObservation](#loadbalancerobservation)_ |  |  |  |


#### Mariadb



Mariadb is the Schema for the Mariadbs API. Manage a MariaDB service in gridscale.



_Appears in:_
- [MariadbList](#mariadblist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Mariadb` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MariadbSpec](#mariadbspec)_ |  |  |  |
| `status` _[MariadbStatus](#mariadbstatus)_ |  |  |  |


#### MariadbInitParameters







_Appears in:_
- [MariadbSpec](#mariadbspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `mariadbBinlogFormat` _string_ | MariaDB parameter: Binary Logging Format. Default: "MIXED".<br />Binary Logging Format. |  |  |
| `mariadbDefaultTimeZone` _string_ | MariaDB parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mariadbLogBin` _boolean_ | MariaDB parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mariadbMaxAllowedPacket` _string_ | MariaDB parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbMaxConnections` _float_ | MariaDB parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mariadbQueryCache` _boolean_ | MariaDB parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mariadbQueryCacheLimit` _string_ | MariaDB parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbQueryCacheSize` _string_ | MariaDB parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbSqlMode` _string_ | MariaDB parameter: SQL Mode. Default: "NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO".<br />SQL Mode. |  |  |
| `mariadbServerId` _float_ | MariaDB parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MariaDB service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MariaDB service. |  |  |
| `release` _string_ | The MariaDB release of this instance. For convenience, please use gscloud to get the list of available MariaDB service releases.<br />The MariaDB release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MariaDB service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MariaDB service. |  |  |


#### MariadbList



MariadbList contains a list of Mariadbs





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MariadbList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Mariadb](#mariadb) array_ |  |  |  |




#### MariadbListenPortObservation







_Appears in:_
- [MariadbObservation](#mariadbobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MariadbObservation







_Appears in:_
- [MariadbStatus](#mariadbstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[MariadbListenPortObservation](#mariadblistenportobservation) array_ | The port numbers where this MariaDB service accepts connections.<br />The port numbers where this MariaDB service accepts connections. |  |  |
| `mariadbBinlogFormat` _string_ | MariaDB parameter: Binary Logging Format. Default: "MIXED".<br />Binary Logging Format. |  |  |
| `mariadbDefaultTimeZone` _string_ | MariaDB parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mariadbLogBin` _boolean_ | MariaDB parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mariadbMaxAllowedPacket` _string_ | MariaDB parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbMaxConnections` _float_ | MariaDB parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mariadbQueryCache` _boolean_ | MariaDB parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mariadbQueryCacheLimit` _string_ | MariaDB parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbQueryCacheSize` _string_ | MariaDB parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mariadbSqlMode` _string_ | MariaDB parameter: SQL Mode. Default: "NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO".<br />SQL Mode. |  |  |
| `mariadbServerId` _float_ | MariaDB parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MariaDB service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MariaDB service. |  |  |
| `release` _string_ | The MariaDB release of this instance. For convenience, please use gscloud to get the list of available MariaDB service releases.<br />The MariaDB release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MariaDB service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MariaDB service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that MariaDB service uses.<br />PaaS service template that MariaDB service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MariaDB service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MariaDB service is in use. |  |  |


#### MariadbParameters







_Appears in:_
- [MariadbSpec](#mariadbspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `mariadbBinlogFormat` _string_ | MariaDB parameter: Binary Logging Format. Default: "MIXED".<br />Binary Logging Format. |  | Optional: \{\} <br /> |
| `mariadbDefaultTimeZone` _string_ | MariaDB parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  | Optional: \{\} <br /> |
| `mariadbLogBin` _boolean_ | MariaDB parameter: Binary Logging. Default: false.<br />Binary Logging. |  | Optional: \{\} <br /> |
| `mariadbMaxAllowedPacket` _string_ | MariaDB parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mariadbMaxConnections` _float_ | MariaDB parameter: Max Connections. Default: 4000.<br />Max Connections. |  | Optional: \{\} <br /> |
| `mariadbQueryCache` _boolean_ | MariaDB parameter: Enable query cache. Default: true.<br />Enable query cache. |  | Optional: \{\} <br /> |
| `mariadbQueryCacheLimit` _string_ | MariaDB parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mariadbQueryCacheSize` _string_ | MariaDB parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mariadbSqlMode` _string_ | MariaDB parameter: SQL Mode. Default: "NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO".<br />SQL Mode. |  | Optional: \{\} <br /> |
| `mariadbServerId` _float_ | MariaDB parameter: Server Id. Default: 1.<br />Server Id. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MariaDB instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of MariaDB service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MariaDB service. |  | Optional: \{\} <br /> |
| `release` _string_ | The MariaDB release of this instance. For convenience, please use gscloud to get the list of available MariaDB service releases.<br />The MariaDB release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MariaDB service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MariaDB service. |  | Optional: \{\} <br /> |


#### MariadbSpec



MariadbSpec defines the desired state of Mariadb



_Appears in:_
- [Mariadb](#mariadb)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[MariadbParameters](#mariadbparameters)_ |  |  |  |
| `initProvider` _[MariadbInitParameters](#mariadbinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MariadbStatus



MariadbStatus defines the observed state of Mariadb.



_Appears in:_
- [Mariadb](#mariadb)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MariadbObservation](#mariadbobservation)_ |  |  |  |


#### Memcached



Memcached is the Schema for the Memcacheds API. Manage a Memcached service in gridscale



_Appears in:_
- [MemcachedList](#memcachedlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Memcached` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MemcachedSpec](#memcachedspec)_ |  |  |  |
| `status` _[MemcachedStatus](#memcachedstatus)_ |  |  |  |


#### MemcachedInitParameters







_Appears in:_
- [MemcachedSpec](#memcachedspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Memcached service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Memcached service. |  |  |
| `release` _string_ | The Memcached release of this instance. For convenience, please use gscloud to get the list of available Memcached service releases.<br />The Memcached release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Memcached service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Memcached service. |  |  |


#### MemcachedList



MemcachedList contains a list of Memcacheds





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MemcachedList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Memcached](#memcached) array_ |  |  |  |




#### MemcachedListenPortObservation







_Appears in:_
- [MemcachedObservation](#memcachedobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MemcachedObservation







_Appears in:_
- [MemcachedStatus](#memcachedstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[MemcachedListenPortObservation](#memcachedlistenportobservation) array_ | The port numbers where this Memcached service accepts connections.<br />The port numbers where this Memcached service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Memcached service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Memcached service. |  |  |
| `release` _string_ | The Memcached release of this instance. For convenience, please use gscloud to get the list of available Memcached service releases.<br />The Memcached release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Memcached service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Memcached service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Memcached service uses.<br />PaaS service template that Memcached service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of Memcached service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that Memcached service is in use. |  |  |


#### MemcachedParameters







_Appears in:_
- [MemcachedSpec](#memcachedspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The Memcached instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Memcached service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Memcached service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Memcached release of this instance. For convenience, please use gscloud to get the list of available Memcached service releases.<br />The Memcached release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Memcached service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Memcached service. |  | Optional: \{\} <br /> |


#### MemcachedSpec



MemcachedSpec defines the desired state of Memcached



_Appears in:_
- [Memcached](#memcached)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[MemcachedParameters](#memcachedparameters)_ |  |  |  |
| `initProvider` _[MemcachedInitParameters](#memcachedinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MemcachedStatus



MemcachedStatus defines the observed state of Memcached.



_Appears in:_
- [Memcached](#memcached)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MemcachedObservation](#memcachedobservation)_ |  |  |  |


#### MySQL



MySQL is the Schema for the MySQLs API. Manage a MySQL service in gridscale.



_Appears in:_
- [MySQLList](#mysqllist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQL` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MySQLSpec](#mysqlspec)_ |  |  |  |
| `status` _[MySQLStatus](#mysqlstatus)_ |  |  |  |


#### MySQLInitParameters







_Appears in:_
- [MySQLSpec](#mysqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlBinlogFormat` _string_ | mysql parameter: Binary Logging Format. Default: "ROW".<br />Binary Logging Format. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlLogBin` _boolean_ | mysql parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlQueryCache` _boolean_ | mysql parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mysqlQueryCacheLimit` _string_ | mysql parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlQueryCacheSize` _string_ | mysql parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `mysqlServerId` _float_ | mysql parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |


#### MySQLList



MySQLList contains a list of MySQLs





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQLList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[MySQL](#mysql) array_ |  |  |  |




#### MySQLListenPortObservation







_Appears in:_
- [MySQLObservation](#mysqlobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MySQLObservation







_Appears in:_
- [MySQLStatus](#mysqlstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[MySQLListenPortObservation](#mysqllistenportobservation) array_ | The port numbers where this mysql service accepts connections.<br />The port numbers where this MySQL service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlBinlogFormat` _string_ | mysql parameter: Binary Logging Format. Default: "ROW".<br />Binary Logging Format. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlLogBin` _boolean_ | mysql parameter: Binary Logging. Default: false.<br />Binary Logging. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlQueryCache` _boolean_ | mysql parameter: Enable query cache. Default: true.<br />Enable query cache. |  |  |
| `mysqlQueryCacheLimit` _string_ | mysql parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlQueryCacheSize` _string_ | mysql parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `mysqlServerId` _float_ | mysql parameter: Server Id. Default: 1.<br />Server Id. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that mysql service uses.<br />PaaS service template that MySQL service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MySQL service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MySQL service is in use. |  |  |


#### MySQLParameters







_Appears in:_
- [MySQLSpec](#mysqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `mysqlBinlogFormat` _string_ | mysql parameter: Binary Logging Format. Default: "ROW".<br />Binary Logging Format. |  | Optional: \{\} <br /> |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  | Optional: \{\} <br /> |
| `mysqlLogBin` _boolean_ | mysql parameter: Binary Logging. Default: false.<br />Binary Logging. |  | Optional: \{\} <br /> |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  | Optional: \{\} <br /> |
| `mysqlQueryCache` _boolean_ | mysql parameter: Enable query cache. Default: true.<br />Enable query cache. |  | Optional: \{\} <br /> |
| `mysqlQueryCacheLimit` _string_ | mysql parameter: Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 1M.<br />Query Cache Limit. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlQueryCacheSize` _string_ | mysql parameter: Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 128M.<br />Query Cache Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  | Optional: \{\} <br /> |
| `mysqlServerId` _float_ | mysql parameter: Server Id. Default: 1.<br />Server Id. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  | Optional: \{\} <br /> |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  | Optional: \{\} <br /> |


#### MySQLSpec



MySQLSpec defines the desired state of MySQL



_Appears in:_
- [MySQL](#mysql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[MySQLParameters](#mysqlparameters)_ |  |  |  |
| `initProvider` _[MySQLInitParameters](#mysqlinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MySQLStatus



MySQLStatus defines the observed state of MySQL.



_Appears in:_
- [MySQL](#mysql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MySQLObservation](#mysqlobservation)_ |  |  |  |


#### Network



Network is the Schema for the Networks API. Manages a network in gridscale.



_Appears in:_
- [NetworkList](#networklist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Network` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[NetworkSpec](#networkspec)_ |  |  |  |
| `status` _[NetworkStatus](#networkstatus)_ |  |  |  |




#### NetworkInitParameters_2







_Appears in:_
- [NetworkSpec](#networkspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `dhcpActive` _boolean_ | Enable DHCP.<br />Enable DHCP. |  |  |
| `dhcpDns` _string_ | The IP address reserved and communicated by the dhcp service to be the default gateway.<br />DHCP DNS. |  |  |
| `dhcpGateway` _string_ | The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />The IP address reserved and communicated by the dhcp service to be the default gateway. |  |  |
| `dhcpRange` _string_ | DHCP DNS. If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically.<br />The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically. |  |  |
| `dhcpReservedSubnet` _string array_ | Subrange within the IP range.<br />Subrange within the IP range. |  |  |
| `l2security` _boolean_ | Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.<br />MAC spoofing protection - filters layer2 and ARP traffic based on source MAC |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |


#### NetworkList



NetworkList contains a list of Networks





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `NetworkList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Network](#network) array_ |  |  |  |


#### NetworkObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `createTime` _string_ | The date and time the object was initially created. |  |  |
| `networkName` _string_ | Name of the network. |  |  |
| `networkUuid` _string_ | The object UUID or id of the network. |  |  |
| `objectName` _string_ | Name of the firewall. |  |  |
| `objectUuid` _string_ | The object UUID or id of the firewall. |  |  |


#### NetworkObservation_2







_Appears in:_
- [NetworkStatus](#networkstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoAssignedServers` _[AutoAssignedServersObservation](#autoassignedserversobservation) array_ | A list of server UUIDs with the corresponding IPs that are designated by the DHCP server.<br />A list of server UUIDs with the corresponding IPs that are designated by the DHCP server. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created |  |  |
| `dhcpActive` _boolean_ | Enable DHCP.<br />Enable DHCP. |  |  |
| `dhcpDns` _string_ | The IP address reserved and communicated by the dhcp service to be the default gateway.<br />DHCP DNS. |  |  |
| `dhcpGateway` _string_ | The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />The IP address reserved and communicated by the dhcp service to be the default gateway. |  |  |
| `dhcpRange` _string_ | DHCP DNS. If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically.<br />The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically. |  |  |
| `dhcpReservedSubnet` _string array_ | Subrange within the IP range.<br />Subrange within the IP range. |  |  |
| `deleteBlock` _boolean_ | If deleting this network is allowed.<br />If deleting this network is allowed |  |  |
| `id` _string_ |  |  |  |
| `l2security` _boolean_ | Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.<br />MAC spoofing protection - filters layer2 and ARP traffic based on source MAC |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this network is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkType` _string_ | The type of this network, can be mpls, breakout or network.<br />The type of this network, can be mpls, breakout or network. |  |  |
| `pinnedServers` _[PinnedServersObservation](#pinnedserversobservation) array_ | A list of server UUIDs with the corresponding IPs that are designated by the user.<br />A list of server UUIDs with the corresponding IPs that are designated by the user. |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |




#### NetworkParameters_2







_Appears in:_
- [NetworkSpec](#networkspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `dhcpActive` _boolean_ | Enable DHCP.<br />Enable DHCP. |  | Optional: \{\} <br /> |
| `dhcpDns` _string_ | The IP address reserved and communicated by the dhcp service to be the default gateway.<br />DHCP DNS. |  | Optional: \{\} <br /> |
| `dhcpGateway` _string_ | The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />The IP address reserved and communicated by the dhcp service to be the default gateway. |  | Optional: \{\} <br /> |
| `dhcpRange` _string_ | DHCP DNS. If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically.<br />The general IP Range configured for this network (/24 for private networks). If it is not set, gridscale internal default range is used.<br />If it is not set and DHCP is enabled, dhcp_range will be set by the backend automatically. |  | Optional: \{\} <br /> |
| `dhcpReservedSubnet` _string array_ | Subrange within the IP range.<br />Subrange within the IP range. |  | Optional: \{\} <br /> |
| `l2security` _boolean_ | Defines information about MAC spoofing protection (filters layer2 and ARP traffic based on MAC source). It can only be (de-)activated on a private network - the public network always has l2security enabled. It will be true if the network is public, and false if the network is private.<br />MAC spoofing protection - filters layer2 and ARP traffic based on source MAC |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |


#### NetworkRulesV4InInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4InObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4InParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkRulesV4OutInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4OutObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV4OutParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkRulesV6InInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6InObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6InParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkRulesV6OutInitParameters







_Appears in:_
- [ServerNetworkInitParameters](#servernetworkinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6OutObservation







_Appears in:_
- [ServerNetworkObservation](#servernetworkobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### NetworkRulesV6OutParameters







_Appears in:_
- [ServerNetworkParameters](#servernetworkparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### NetworkSpec



NetworkSpec defines the desired state of Network



_Appears in:_
- [Network](#network)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[NetworkParameters_2](#networkparameters_2)_ |  |  |  |
| `initProvider` _[NetworkInitParameters_2](#networkinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### NetworkStatus



NetworkStatus defines the observed state of Network.



_Appears in:_
- [Network](#network)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[NetworkObservation_2](#networkobservation_2)_ |  |  |  |


#### NodePoolInitParameters







_Appears in:_
- [K8SInitParameters](#k8sinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cores` _float_ | Cores per worker node.<br />Cores per worker node. |  |  |
| `labels` _[LabelsInitParameters](#labelsinitparameters) array_ | List of labels to be applied to the nodes of this pool. Check the product documentation for details<br />List of labels to be applied to the nodes of this pool. |  |  |
| `memory` _float_ | Memory per worker node (in GiB).<br />Memory per worker node (in GiB). |  |  |
| `name` _string_ | Name of the node pool.<br />Name of node pool. |  |  |
| `nodeCount` _float_ | Number of worker nodes.<br />Number of worker nodes. |  |  |
| `rocketStorage` _float_ | Rocket storage per worker node (in GiB).<br />Rocket storage per worker node (in GiB). |  |  |
| `storage` _float_ | Storage per worker node (in GiB).<br />Storage per worker node (in GiB). |  |  |
| `storageType` _string_ | Storage type (one of storage, storage_high, storage_insane).<br />Storage type. |  |  |
| `taints` _[TaintsInitParameters](#taintsinitparameters) array_ | List of taints to be applied to the nodes of this pool. Check the product documentation for details<br />List of taints to be applied to the nodes of this pool. |  |  |


#### NodePoolObservation







_Appears in:_
- [K8SObservation](#k8sobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cores` _float_ | Cores per worker node.<br />Cores per worker node. |  |  |
| `labels` _[LabelsObservation](#labelsobservation) array_ | List of labels to be applied to the nodes of this pool. Check the product documentation for details<br />List of labels to be applied to the nodes of this pool. |  |  |
| `memory` _float_ | Memory per worker node (in GiB).<br />Memory per worker node (in GiB). |  |  |
| `name` _string_ | Name of the node pool.<br />Name of node pool. |  |  |
| `nodeCount` _float_ | Number of worker nodes.<br />Number of worker nodes. |  |  |
| `rocketStorage` _float_ | Rocket storage per worker node (in GiB).<br />Rocket storage per worker node (in GiB). |  |  |
| `storage` _float_ | Storage per worker node (in GiB).<br />Storage per worker node (in GiB). |  |  |
| `storageType` _string_ | Storage type (one of storage, storage_high, storage_insane).<br />Storage type. |  |  |
| `taints` _[TaintsObservation](#taintsobservation) array_ | List of taints to be applied to the nodes of this pool. Check the product documentation for details<br />List of taints to be applied to the nodes of this pool. |  |  |


#### NodePoolParameters







_Appears in:_
- [K8SParameters](#k8sparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `cores` _float_ | Cores per worker node.<br />Cores per worker node. |  | Optional: \{\} <br /> |
| `labels` _[LabelsParameters](#labelsparameters) array_ | List of labels to be applied to the nodes of this pool. Check the product documentation for details<br />List of labels to be applied to the nodes of this pool. |  | Optional: \{\} <br /> |
| `memory` _float_ | Memory per worker node (in GiB).<br />Memory per worker node (in GiB). |  | Optional: \{\} <br /> |
| `name` _string_ | Name of the node pool.<br />Name of node pool. |  | Optional: \{\} <br /> |
| `nodeCount` _float_ | Number of worker nodes.<br />Number of worker nodes. |  | Optional: \{\} <br /> |
| `rocketStorage` _float_ | Rocket storage per worker node (in GiB).<br />Rocket storage per worker node (in GiB). |  | Optional: \{\} <br /> |
| `storage` _float_ | Storage per worker node (in GiB).<br />Storage per worker node (in GiB). |  | Optional: \{\} <br /> |
| `storageType` _string_ | Storage type (one of storage, storage_high, storage_insane).<br />Storage type. |  | Optional: \{\} <br /> |
| `taints` _[TaintsParameters](#taintsparameters) array_ | List of taints to be applied to the nodes of this pool. Check the product documentation for details<br />List of taints to be applied to the nodes of this pool. |  | Optional: \{\} <br /> |


#### ObjectStorageExportInitParameters







_Appears in:_
- [SnapshotInitParameters](#snapshotinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Access key |  |  |
| `bucket` _string_ | Bucket name |  |  |
| `host` _string_ | Host of object storage. Must be of URL type. E.g: https://gos3.io |  |  |
| `object` _string_ | Name of file (include file path) |  |  |
| `private` _boolean_ |  |  |  |
| `secretKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Secret key |  |  |


#### ObjectStorageExportObservation







_Appears in:_
- [SnapshotObservation](#snapshotobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bucket` _string_ | Bucket name |  |  |
| `host` _string_ | Host of object storage. Must be of URL type. E.g: https://gos3.io |  |  |
| `object` _string_ | Name of file (include file path) |  |  |
| `private` _boolean_ |  |  |  |
| `status` _string_ |  |  |  |


#### ObjectStorageExportParameters







_Appears in:_
- [SnapshotParameters](#snapshotparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Access key |  | Optional: \{\} <br /> |
| `bucket` _string_ | Bucket name |  | Optional: \{\} <br /> |
| `host` _string_ | Host of object storage. Must be of URL type. E.g: https://gos3.io |  | Optional: \{\} <br /> |
| `object` _string_ | Name of file (include file path) |  | Optional: \{\} <br /> |
| `private` _boolean_ |  |  | Optional: \{\} <br /> |
| `secretKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Secret key |  | Optional: \{\} <br /> |


#### Paas



Paas is the Schema for the Paass API. Manages a PaaS in gridscale.



_Appears in:_
- [PaasList](#paaslist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Paas` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[PaasSpec](#paasspec)_ |  |  |  |
| `status` _[PaasStatus](#paasstatus)_ |  |  |  |


#### PaasInitParameters







_Appears in:_
- [PaasSpec](#paasspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `parameter` _[ParameterInitParameters](#parameterinitparameters) array_ | See Argument Reference above.<br />Parameter for PaaS service |  |  |
| `resourceLimit` _[ResourceLimitInitParameters](#resourcelimitinitparameters) array_ | A list of service resource limits..<br />Resource for PaaS service |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PaaS service |  |  |
| `serviceTemplateUuid` _string_ | The template used to create the service.<br />Template that PaaS service uses |  |  |


#### PaasList



PaasList contains a list of Paass





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `PaasList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Paas](#paas) array_ |  |  |  |




#### PaasListenPortObservation







_Appears in:_
- [PaasObservation](#paasobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### PaasObservation







_Appears in:_
- [PaasStatus](#paasstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change |  |  |
| `createTime` _string_ | Time of the creation.<br />Time of the creation |  |  |
| `currentPrice` _float_ | Current price of PaaS service.<br />Current price of PaaS service |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[PaasListenPortObservation](#paaslistenportobservation) array_ | Ports that PaaS service listens to.<br />Ports that PaaS service listens to |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `parameter` _[ParameterObservation](#parameterobservation) array_ | See Argument Reference above.<br />Parameter for PaaS service |  |  |
| `resourceLimit` _[ResourceLimitObservation](#resourcelimitobservation) array_ | A list of service resource limits..<br />Resource for PaaS service |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PaaS service |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | The template used to create the service.<br />Template that PaaS service uses |  |  |
| `serviceTemplateUuidComputed` _string_ | Template that PaaS service uses.<br />Template that PaaS service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of PaaS service |  |  |
| `usageInMinute` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that PaaS service is in use |  |  |


#### PaasParameters







_Appears in:_
- [PaasSpec](#paasspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `parameter` _[ParameterParameters](#parameterparameters) array_ | See Argument Reference above.<br />Parameter for PaaS service |  | Optional: \{\} <br /> |
| `resourceLimit` _[ResourceLimitParameters](#resourcelimitparameters) array_ | A list of service resource limits..<br />Resource for PaaS service |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PaaS service |  | Optional: \{\} <br /> |
| `serviceTemplateUuid` _string_ | The template used to create the service.<br />Template that PaaS service uses |  | Optional: \{\} <br /> |


#### PaasSpec



PaasSpec defines the desired state of Paas



_Appears in:_
- [Paas](#paas)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[PaasParameters](#paasparameters)_ |  |  |  |
| `initProvider` _[PaasInitParameters](#paasinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### PaasStatus



PaasStatus defines the observed state of Paas.



_Appears in:_
- [Paas](#paas)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[PaasObservation](#paasobservation)_ |  |  |  |


#### ParameterInitParameters







_Appears in:_
- [PaasInitParameters](#paasinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `param` _string_ | Name of parameter. |  |  |
| `type` _string_ | Primitive type of the parameter: bool, int (better use float for int case), float, string. |  |  |
| `value` _string_ | Value of the corresponding parameter. |  |  |


#### ParameterObservation







_Appears in:_
- [PaasObservation](#paasobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `param` _string_ | Name of parameter. |  |  |
| `type` _string_ | Primitive type of the parameter: bool, int (better use float for int case), float, string. |  |  |
| `value` _string_ | Value of the corresponding parameter. |  |  |


#### ParameterParameters







_Appears in:_
- [PaasParameters](#paasparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `param` _string_ | Name of parameter. |  | Optional: \{\} <br /> |
| `type` _string_ | Primitive type of the parameter: bool, int (better use float for int case), float, string. |  | Optional: \{\} <br /> |
| `value` _string_ | Value of the corresponding parameter. |  | Optional: \{\} <br /> |




#### PinnedServersObservation







_Appears in:_
- [NetworkObservation_2](#networkobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `ip` _string_ | IP which is assigned to the server. |  |  |
| `serverUuid` _string_ | UUID of the server. |  |  |




#### Postgresql



Postgresql is the Schema for the Postgresqls API. Manage a PostgreSQL service in gridscale.



_Appears in:_
- [PostgresqlList](#postgresqllist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Postgresql` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[PostgresqlSpec](#postgresqlspec)_ |  |  |  |
| `status` _[PostgresqlStatus](#postgresqlstatus)_ |  |  |  |


#### PostgresqlInitParameters







_Appears in:_
- [PostgresqlSpec](#postgresqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of PostgreSQL service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of PostgreSQL service. |  |  |
| `pgauditLogAccessKey` _string_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  |  |
| `pgauditLogBucket` _string_ | Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key.<br />Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key. |  |  |
| `pgauditLogRotationFrequency` _float_ | Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. Default is 5 minutes.<br />Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. |  |  |
| `pgauditLogSecretKey` _string_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  |  |
| `pgauditLogServerUrl` _string_ | Object Storage server URL the bucket is located on.<br />Object Storage server URL the bucket is located on. |  |  |
| `release` _string_ | The PostgreSQL release of this instance. For convenience, please use gscloud to get the list of available PostgreSQL service releases.<br />The PostgreSQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available PostgreSQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PostgreSQL service. |  |  |


#### PostgresqlList



PostgresqlList contains a list of Postgresqls





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `PostgresqlList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Postgresql](#postgresql) array_ |  |  |  |




#### PostgresqlListenPortObservation







_Appears in:_
- [PostgresqlObservation](#postgresqlobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### PostgresqlObservation







_Appears in:_
- [PostgresqlStatus](#postgresqlstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[PostgresqlListenPortObservation](#postgresqllistenportobservation) array_ | The port numbers where this PostgreSQL service accepts connections.<br />The port numbers where this PostgreSQL service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of PostgreSQL service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of PostgreSQL service. |  |  |
| `pgauditLogAccessKey` _string_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  |  |
| `pgauditLogBucket` _string_ | Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key.<br />Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key. |  |  |
| `pgauditLogRotationFrequency` _float_ | Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. Default is 5 minutes.<br />Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. |  |  |
| `pgauditLogSecretKey` _string_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  |  |
| `pgauditLogServerUrl` _string_ | Object Storage server URL the bucket is located on.<br />Object Storage server URL the bucket is located on. |  |  |
| `release` _string_ | The PostgreSQL release of this instance. For convenience, please use gscloud to get the list of available PostgreSQL service releases.<br />The PostgreSQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available PostgreSQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PostgreSQL service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that PostgreSQL service uses.<br />PaaS service template that PostgreSQL service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of PostgreSQL service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that PostgreSQL service is in use. |  |  |


#### PostgresqlParameters







_Appears in:_
- [PostgresqlSpec](#postgresqlspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The PostgreSQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of PostgreSQL service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of PostgreSQL service. |  | Optional: \{\} <br /> |
| `pgauditLogAccessKey` _string_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `pgauditLogBucket` _string_ | Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key.<br />Object Storage bucket to upload audit logs to. For pgAudit to be enabled these additional parameters need to be configured: pgaudit_log_server_url, pgaudit_log_access_key, pgaudit_log_secret_key. |  | Optional: \{\} <br /> |
| `pgauditLogRotationFrequency` _float_ | Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. Default is 5 minutes.<br />Rotation (in minutes) for audit logs. Logs are uploaded to Object Storage once rotated. |  | Optional: \{\} <br /> |
| `pgauditLogSecretKey` _string_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `pgauditLogServerUrl` _string_ | Object Storage server URL the bucket is located on.<br />Object Storage server URL the bucket is located on. |  | Optional: \{\} <br /> |
| `release` _string_ | The PostgreSQL release of this instance. For convenience, please use gscloud to get the list of available PostgreSQL service releases.<br />The PostgreSQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available PostgreSQL service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to PostgreSQL service. |  | Optional: \{\} <br /> |


#### PostgresqlSpec



PostgresqlSpec defines the desired state of Postgresql



_Appears in:_
- [Postgresql](#postgresql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[PostgresqlParameters](#postgresqlparameters)_ |  |  |  |
| `initProvider` _[PostgresqlInitParameters](#postgresqlinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### PostgresqlStatus



PostgresqlStatus defines the observed state of Postgresql.



_Appears in:_
- [Postgresql](#postgresql)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[PostgresqlObservation](#postgresqlobservation)_ |  |  |  |


#### ResourceLimitInitParameters







_Appears in:_
- [PaasInitParameters](#paasinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `limit` _float_ | The maximum number of the specific resource your service can use. |  |  |
| `resource` _string_ | The name of the resource you would like to cap. |  |  |


#### ResourceLimitObservation







_Appears in:_
- [PaasObservation](#paasobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `limit` _float_ | The maximum number of the specific resource your service can use. |  |  |
| `resource` _string_ | The name of the resource you would like to cap. |  |  |


#### ResourceLimitParameters







_Appears in:_
- [PaasParameters](#paasparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `limit` _float_ | The maximum number of the specific resource your service can use. |  | Optional: \{\} <br /> |
| `resource` _string_ | The name of the resource you would like to cap. |  | Optional: \{\} <br /> |


#### RollbackInitParameters







_Appears in:_
- [SnapshotInitParameters](#snapshotinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `id` _string_ | ID of the rollback request. Each rollback request has to have a unique id. ID can be any string value. |  |  |


#### RollbackObservation







_Appears in:_
- [SnapshotObservation](#snapshotobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `id` _string_ | ID of the rollback request. Each rollback request has to have a unique id. ID can be any string value. |  |  |
| `rollbackTime` _string_ |  |  |  |
| `status` _string_ |  |  |  |


#### RollbackParameters







_Appears in:_
- [SnapshotParameters](#snapshotparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `id` _string_ | ID of the rollback request. Each rollback request has to have a unique id. ID can be any string value. |  | Optional: \{\} <br /> |


#### RulesV4InInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4InObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4InParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### RulesV4OutInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4OutObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV4OutParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### RulesV6InInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6InObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6InParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### RulesV6OutInitParameters







_Appears in:_
- [FirewallInitParameters](#firewallinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6OutObservation







_Appears in:_
- [FirewallObservation](#firewallobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  |  |
| `comment` _string_ | Comment.<br />Comment. |  |  |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  |  |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  |  |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  |  |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  |  |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  |  |


#### RulesV6OutParameters







_Appears in:_
- [FirewallParameters](#firewallparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `action` _string_ | This defines what the firewall will do. Either accept or drop.<br />This defines what the firewall will do. Either accept or drop. |  | Optional: \{\} <br /> |
| `comment` _string_ | Comment.<br />Comment. |  | Optional: \{\} <br /> |
| `dstCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then all IPs have access to this service. |  | Optional: \{\} <br /> |
| `dstPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |
| `order` _float_ | The order at which the firewall will compare packets against its rules. A packet will be compared against the first rule, it will either allow it to pass or block it and it won't be matched against any other rules. However, if it does no match the rule, then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound).<br />The order at which the firewall will compare packets against its rules.<br />A packet will be compared against the first rule, it will either allow it to pass or block it<br />and it won't be matched against any other rules. However, if it does no match the rule,<br />then it will proceed onto rule 2. Packets that do not match any rules are blocked by default (Only for inbound). |  | Optional: \{\} <br /> |
| `protocol` _string_ | Either 'udp' or 'tcp'.<br />Either 'udp' or 'tcp' |  | Optional: \{\} <br /> |
| `srcCidr` _string_ | Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs.<br />Either an IPv4/6 address or and IP Network in CIDR format. If this field is empty then this service has access to all IPs. |  | Optional: \{\} <br /> |
| `srcPort` _string_ | A Number between 1 and 65535, port ranges are separated by a colon for FTP.<br />A Number between 1 and 65535, port ranges are seperated by a colon for FTP |  | Optional: \{\} <br /> |


#### S3BackupInitParameters







_Appears in:_
- [SqlserverInitParameters](#sqlserverinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `backupAccessKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  |  |
| `backupBucket` _string_ | Object Storage bucket to upload backups to and restore backups from.<br />Object Storage bucket to upload backups to and restore backups from. |  |  |
| `backupRetention` _float_ | Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage).<br />Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage). |  |  |
| `backupSecretKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  |  |
| `backupServerUrl` _string_ | Object Storage server URL the bucket is located on. Note: Currently, only object storage host "https://gos3.io/" is supported.<br />Object Storage server URL the bucket is located on. |  |  |


#### S3BackupObservation







_Appears in:_
- [SqlserverObservation](#sqlserverobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `backupBucket` _string_ | Object Storage bucket to upload backups to and restore backups from.<br />Object Storage bucket to upload backups to and restore backups from. |  |  |
| `backupRetention` _float_ | Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage).<br />Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage). |  |  |
| `backupServerUrl` _string_ | Object Storage server URL the bucket is located on. Note: Currently, only object storage host "https://gos3.io/" is supported.<br />Object Storage server URL the bucket is located on. |  |  |


#### S3BackupParameters







_Appears in:_
- [SqlserverParameters](#sqlserverparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `backupAccessKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Access key used to authenticate against Object Storage server.<br />Access key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `backupBucket` _string_ | Object Storage bucket to upload backups to and restore backups from.<br />Object Storage bucket to upload backups to and restore backups from. |  | Optional: \{\} <br /> |
| `backupRetention` _float_ | Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage).<br />Retention (in seconds) for local originals of backups. (0 for immediate removal once uploaded to Object Storage (default), higher values for delayed removal after the given time and once uploaded to Object Storage). |  | Optional: \{\} <br /> |
| `backupSecretKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | Secret key used to authenticate against Object Storage server.<br />Secret key used to authenticate against Object Storage server. |  | Optional: \{\} <br /> |
| `backupServerUrl` _string_ | Object Storage server URL the bucket is located on. Note: Currently, only object storage host "https://gos3.io/" is supported.<br />Object Storage server URL the bucket is located on. |  | Optional: \{\} <br /> |


#### Server



Server is the Schema for the Servers API. Manages a server in gridscale.



_Appears in:_
- [ServerList](#serverlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Server` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ServerSpec](#serverspec)_ |  |  |  |
| `status` _[ServerStatus](#serverstatus)_ |  |  |  |




#### ServerInitParameters_2







_Appears in:_
- [ServerSpec](#serverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoRecovery` _boolean_ | If the server should be auto-started in case of a failure (default=true).<br />If the server should be auto-started in case of a failure (default=true). |  |  |
| `availabilityZone` _string_ | Defines which Availability-Zone the Server is placed.<br />Defines which Availability-Zone the Server is placed. |  |  |
| `cores` _float_ | The number of server cores.<br />The number of server cores. |  |  |
| `hardwareProfile` _string_ | The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. If it is not set, the backend will set it by default. Check the official docs.<br />Specifies the hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `hardwareProfileConfig` _[HardwareProfileConfigInitParameters](#hardwareprofileconfiginitparameters) array_ | Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. Note: If hardware_profile_config is set, all fields of hardware_profile_config MUST be set.<br />Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `ipv4` _string_ | The UUID of the IPv4 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `ipv6` _string_ | The UUID of the IPv6 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `isoimage` _string_ | The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in the expert panel. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `memory` _float_ | The amount of server memory in GB.<br />The amount of server memory in GB. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `network` _[ServerNetworkInitParameters](#servernetworkinitparameters) array_ | Connects a network to the server. The network ordering of the server corresponds to the order of the networks in the server resource block. |  |  |
| `power` _boolean_ | The power state of the server. Set this to true to will boot the server, false will shut it down.<br />The number of server cores. |  |  |
| `storage` _[StorageInitParameters](#storageinitparameters) array_ | Connects a storage to the server. **NOTE: The first storage is always the boot device.<br />A list of storages attached to the server. The first storage in the list is always set as the boot storage of the server. |  |  |
| `userDataBase64` _string_ | For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition.<br />For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition. |  |  |


#### ServerList



ServerList contains a list of Servers





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ServerList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Server](#server) array_ |  |  |  |


#### ServerNetworkInitParameters







_Appears in:_
- [ServerInitParameters_2](#serverinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  |  |
| `firewallTemplateUuid` _string_ | The UUID of firewall template. |  |  |
| `ip` _string_ | Manually assign DHCP IP to the server (if applicable).<br />Manually assign DHCP IP to the server. |  |  |
| `objectUuid` _string_ | The object UUID or id of the network. |  |  |
| `ordering` _float_ | DEPRECATED  Defines the ordering of the network interfaces. Lower numbers have lower PCI-IDs. |  |  |
| `rulesV4In` _[NetworkRulesV4InInitParameters](#networkrulesv4ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[NetworkRulesV4OutInitParameters](#networkrulesv4outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[NetworkRulesV6InInitParameters](#networkrulesv6ininitparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[NetworkRulesV6OutInitParameters](#networkrulesv6outinitparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |


#### ServerNetworkObservation







_Appears in:_
- [ServerObservation_2](#serverobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoAssignedIp` _string_ | DHCP IP which is automatically assigned to the server (if applicable).<br />DHCP IP which is automatically assigned to the server. |  |  |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created. |  |  |
| `firewallTemplateUuid` _string_ | The UUID of firewall template. |  |  |
| `ip` _string_ | Manually assign DHCP IP to the server (if applicable).<br />Manually assign DHCP IP to the server. |  |  |
| `mac` _string_ | network_mac defines the MAC address of the network interface. |  |  |
| `networkType` _string_ | One of network, network_high, network_insane. |  |  |
| `objectName` _string_ | Name of the storage. |  |  |
| `objectUuid` _string_ | The object UUID or id of the network. |  |  |
| `ordering` _float_ | DEPRECATED  Defines the ordering of the network interfaces. Lower numbers have lower PCI-IDs. |  |  |
| `rulesV4In` _[NetworkRulesV4InObservation](#networkrulesv4inobservation) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  |  |
| `rulesV4Out` _[NetworkRulesV4OutObservation](#networkrulesv4outobservation) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  |  |
| `rulesV6In` _[NetworkRulesV6InObservation](#networkrulesv6inobservation) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  |  |
| `rulesV6Out` _[NetworkRulesV6OutObservation](#networkrulesv6outobservation) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  |  |


#### ServerNetworkParameters







_Appears in:_
- [ServerParameters_2](#serverparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  | Optional: \{\} <br /> |
| `firewallTemplateUuid` _string_ | The UUID of firewall template. |  | Optional: \{\} <br /> |
| `ip` _string_ | Manually assign DHCP IP to the server (if applicable).<br />Manually assign DHCP IP to the server. |  | Optional: \{\} <br /> |
| `objectUuid` _string_ | The object UUID or id of the network. |  | Optional: \{\} <br /> |
| `ordering` _float_ | DEPRECATED  Defines the ordering of the network interfaces. Lower numbers have lower PCI-IDs. |  | Optional: \{\} <br /> |
| `rulesV4In` _[NetworkRulesV4InParameters](#networkrulesv4inparameters) array_ | Firewall template rules for inbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV4Out` _[NetworkRulesV4OutParameters](#networkrulesv4outparameters) array_ | Firewall template rules for outbound traffic - covers ipv4 addresses. |  | Optional: \{\} <br /> |
| `rulesV6In` _[NetworkRulesV6InParameters](#networkrulesv6inparameters) array_ | Firewall template rules for inbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |
| `rulesV6Out` _[NetworkRulesV6OutParameters](#networkrulesv6outparameters) array_ | Firewall template rules for outbound traffic - covers ipv6 addresses. |  | Optional: \{\} <br /> |


#### ServerObservation







_Appears in:_
- [IsoimageObservation](#isoimageobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ |  |  |  |
| `createTime` _string_ |  |  |  |
| `objectName` _string_ |  |  |  |
| `objectUuid` _string_ |  |  |  |


#### ServerObservation_2







_Appears in:_
- [ServerStatus](#serverstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoRecovery` _boolean_ | If the server should be auto-started in case of a failure (default=true).<br />If the server should be auto-started in case of a failure (default=true). |  |  |
| `availabilityZone` _string_ | Defines which Availability-Zone the Server is placed.<br />Defines which Availability-Zone the Server is placed. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change. |  |  |
| `consoleToken` _string_ | The token used by the panel to open the websocket VNC connection to the server console.<br />The token used by the panel to open the websocket VNC connection to the server console. |  |  |
| `cores` _float_ | The number of server cores.<br />The number of server cores. |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `hardwareProfile` _string_ | The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. If it is not set, the backend will set it by default. Check the official docs.<br />Specifies the hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `hardwareProfileConfig` _[HardwareProfileConfigObservation](#hardwareprofileconfigobservation) array_ | Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. Note: If hardware_profile_config is set, all fields of hardware_profile_config MUST be set.<br />Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  |  |
| `id` _string_ | UUID of the server. |  |  |
| `ipv4` _string_ | The UUID of the IPv4 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `ipv6` _string_ | The UUID of the IPv6 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  |  |
| `isoimage` _string_ | The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in the expert panel. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `legacy` _boolean_ | Legacy-Hardware emulation instead of virtio hardware. If enabled, hot-plugging cores, memory, storage, network, etc. will not work, but the server will most likely run every x86 compatible operating system. This mode comes with a performance penalty, as emulated hardware does not benefit from the virtio driver infrastructure.<br />Legacy-Hardware emulation instead of virtio hardware. If enabled, hotplugging cores, memory, storage, network, etc. will not work, but the server will most likely run every x86 compatible operating system. This mode comes with a performance penalty, as emulated hardware does not benefit from the virtio driver infrastructure. |  |  |
| `locationUuid` _string_ | The location this server is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `memory` _float_ | The amount of server memory in GB.<br />The amount of server memory in GB. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `network` _[ServerNetworkObservation](#servernetworkobservation) array_ | Connects a network to the server. The network ordering of the server corresponds to the order of the networks in the server resource block. |  |  |
| `power` _boolean_ | The power state of the server. Set this to true to will boot the server, false will shut it down.<br />The number of server cores. |  |  |
| `status` _string_ | Status indicates the status of the object. |  |  |
| `storage` _[StorageObservation](#storageobservation) array_ | Connects a storage to the server. **NOTE: The first storage is always the boot device.<br />A list of storages attached to the server. The first storage in the list is always set as the boot storage of the server. |  |  |
| `usageInMinutesCores` _float_ | Total minutes of cores used. |  |  |
| `usageInMinutesMemory` _float_ | Total minutes of memory used. |  |  |
| `userDataBase64` _string_ | For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition.<br />For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition. |  |  |




#### ServerParameters_2







_Appears in:_
- [ServerSpec](#serverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `autoRecovery` _boolean_ | If the server should be auto-started in case of a failure (default=true).<br />If the server should be auto-started in case of a failure (default=true). |  | Optional: \{\} <br /> |
| `availabilityZone` _string_ | Defines which Availability-Zone the Server is placed.<br />Defines which Availability-Zone the Server is placed. |  | Optional: \{\} <br /> |
| `cores` _float_ | The number of server cores.<br />The number of server cores. |  | Optional: \{\} <br /> |
| `hardwareProfile` _string_ | The hardware profile of the Server. Options are default, legacy, nested, cisco_csr, sophos_utm, f5_bigip and q35 at the moment of writing. If it is not set, the backend will set it by default. Check the official docs.<br />Specifies the hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  | Optional: \{\} <br /> |
| `hardwareProfileConfig` _[HardwareProfileConfigParameters](#hardwareprofileconfigparameters) array_ | Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. Note: If hardware_profile_config is set, all fields of hardware_profile_config MUST be set.<br />Specifies the custom hardware settings for the virtual machine. Note: hardware_profile and hardware_profile_config parameters can't be used at the same time. |  | Optional: \{\} <br /> |
| `ipv4` _string_ | The UUID of the IPv4 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  | Optional: \{\} <br /> |
| `ipv6` _string_ | The UUID of the IPv6 address of the server. (***NOTE: The server will NOT automatically be connected to the public network; to give it access to the internet, please add server to the public network.) |  | Optional: \{\} <br /> |
| `isoimage` _string_ | The UUID of an ISO image in gridscale. The server will automatically boot from the ISO if one was added. The UUIDs of ISO images can be found in the expert panel. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `memory` _float_ | The amount of server memory in GB.<br />The amount of server memory in GB. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `network` _[ServerNetworkParameters](#servernetworkparameters) array_ | Connects a network to the server. The network ordering of the server corresponds to the order of the networks in the server resource block. |  | Optional: \{\} <br /> |
| `power` _boolean_ | The power state of the server. Set this to true to will boot the server, false will shut it down.<br />The number of server cores. |  | Optional: \{\} <br /> |
| `storage` _[StorageParameters](#storageparameters) array_ | Connects a storage to the server. **NOTE: The first storage is always the boot device.<br />A list of storages attached to the server. The first storage in the list is always set as the boot storage of the server. |  | Optional: \{\} <br /> |
| `userDataBase64` _string_ | For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition.<br />For system configuration on first boot. May contain cloud-config data or shell scripting, encoded as base64 string. Supported tools are cloud-init, Cloudbase-init, and Ignition. |  | Optional: \{\} <br /> |


#### ServerSpec



ServerSpec defines the desired state of Server



_Appears in:_
- [Server](#server)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[ServerParameters_2](#serverparameters_2)_ |  |  |  |
| `initProvider` _[ServerInitParameters_2](#serverinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### ServerStatus



ServerStatus defines the observed state of Server.



_Appears in:_
- [Server](#server)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[ServerObservation_2](#serverobservation_2)_ |  |  |  |


#### Snapshot



Snapshot is the Schema for the Snapshots API. <no value>



_Appears in:_
- [SnapshotList](#snapshotlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Snapshot` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SnapshotSpec](#snapshotspec)_ |  |  |  |
| `status` _[SnapshotStatus](#snapshotstatus)_ |  |  |  |


#### SnapshotInitParameters







_Appears in:_
- [SnapshotSpec](#snapshotspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `objectStorageExport` _[ObjectStorageExportInitParameters](#objectstorageexportinitparameters) array_ | Export snapshot to a object storage |  |  |
| `rollback` _[RollbackInitParameters](#rollbackinitparameters) array_ | Returns a storage to the state of the selected Snapshot. |  |  |
| `storageUuid` _string_ | UUID of the storage used to create this snapshot |  |  |


#### SnapshotList



SnapshotList contains a list of Snapshots





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SnapshotList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Snapshot](#snapshot) array_ |  |  |  |


#### SnapshotObservation







_Appears in:_
- [SnapshotStatus](#snapshotstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The capacity of a storage/ISO image/template/snapshot in GB |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows<br />the product_no of the license (see the /prices endpoint for more details) |  |  |
| `locationCountry` _string_ | The human-readable name of the location |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The human-readable name of the location |  |  |
| `locationUuid` _string_ | The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `objectStorageExport` _[ObjectStorageExportObservation](#objectstorageexportobservation) array_ | Export snapshot to a object storage |  |  |
| `rollback` _[RollbackObservation](#rollbackobservation) array_ | Returns a storage to the state of the selected Snapshot. |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `storageUuid` _string_ | UUID of the storage used to create this snapshot |  |  |
| `usageInMinutes` _float_ | Total minutes the object has been running |  |  |


#### SnapshotParameters







_Appears in:_
- [SnapshotSpec](#snapshotspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object |  | Optional: \{\} <br /> |
| `objectStorageExport` _[ObjectStorageExportParameters](#objectstorageexportparameters) array_ | Export snapshot to a object storage |  | Optional: \{\} <br /> |
| `rollback` _[RollbackParameters](#rollbackparameters) array_ | Returns a storage to the state of the selected Snapshot. |  | Optional: \{\} <br /> |
| `storageUuid` _string_ | UUID of the storage used to create this snapshot |  | Optional: \{\} <br /> |


#### SnapshotSpec



SnapshotSpec defines the desired state of Snapshot



_Appears in:_
- [Snapshot](#snapshot)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[SnapshotParameters](#snapshotparameters)_ |  |  |  |
| `initProvider` _[SnapshotInitParameters](#snapshotinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SnapshotStatus



SnapshotStatus defines the observed state of Snapshot.



_Appears in:_
- [Snapshot](#snapshot)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SnapshotObservation](#snapshotobservation)_ |  |  |  |


#### Snapshotschedule



Snapshotschedule is the Schema for the Snapshotschedules API. <no value>



_Appears in:_
- [SnapshotscheduleList](#snapshotschedulelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Snapshotschedule` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SnapshotscheduleSpec](#snapshotschedulespec)_ |  |  |  |
| `status` _[SnapshotscheduleStatus](#snapshotschedulestatus)_ |  |  |  |


#### SnapshotscheduleInitParameters







_Appears in:_
- [SnapshotscheduleSpec](#snapshotschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `keepSnapshots` _float_ | The amount of Snapshots to keep before overwriting the last created Snapshot |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the snapshot schedule will be run |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `storageUuid` _string_ | UUID of the storage used to create snapshots |  |  |


#### SnapshotscheduleList



SnapshotscheduleList contains a list of Snapshotschedules





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SnapshotscheduleList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Snapshotschedule](#snapshotschedule) array_ |  |  |  |


#### SnapshotscheduleObservation







_Appears in:_
- [SnapshotscheduleStatus](#snapshotschedulestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `keepSnapshots` _float_ | The amount of Snapshots to keep before overwriting the last created Snapshot |  |  |
| `labels` _string array_ | List of labels. |  |  |
| `name` _string_ | The human-readable name of the object |  |  |
| `nextRuntime` _string_ | The date and time that the snapshot schedule will be run |  |  |
| `nextRuntimeComputed` _string_ | The date and time that the snapshot schedule will be run. This date and time is computed by gridscale's server. |  |  |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  |  |
| `snapshot` _[SnapshotscheduleSnapshotObservation](#snapshotschedulesnapshotobservation) array_ | Related snashots |  |  |
| `status` _string_ | Status indicates the status of the object |  |  |
| `storageUuid` _string_ | UUID of the storage used to create snapshots |  |  |


#### SnapshotscheduleParameters







_Appears in:_
- [SnapshotscheduleSpec](#snapshotschedulespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `keepSnapshots` _float_ | The amount of Snapshots to keep before overwriting the last created Snapshot |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object |  | Optional: \{\} <br /> |
| `nextRuntime` _string_ | The date and time that the snapshot schedule will be run |  | Optional: \{\} <br /> |
| `runInterval` _float_ | The interval at which the schedule will run (in minutes) |  | Optional: \{\} <br /> |
| `storageUuid` _string_ | UUID of the storage used to create snapshots |  | Optional: \{\} <br /> |




#### SnapshotscheduleSnapshotObservation







_Appears in:_
- [SnapshotscheduleObservation](#snapshotscheduleobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `createTime` _string_ |  |  |  |
| `name` _string_ |  |  |  |
| `objectUuid` _string_ |  |  |  |




#### SnapshotscheduleSpec



SnapshotscheduleSpec defines the desired state of Snapshotschedule



_Appears in:_
- [Snapshotschedule](#snapshotschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[SnapshotscheduleParameters](#snapshotscheduleparameters)_ |  |  |  |
| `initProvider` _[SnapshotscheduleInitParameters](#snapshotscheduleinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SnapshotscheduleStatus



SnapshotscheduleStatus defines the observed state of Snapshotschedule.



_Appears in:_
- [Snapshotschedule](#snapshotschedule)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SnapshotscheduleObservation](#snapshotscheduleobservation)_ |  |  |  |


#### Sqlserver



Sqlserver is the Schema for the Sqlservers API. Manage a MS SQL server service in gridscale.



_Appears in:_
- [SqlserverList](#sqlserverlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Sqlserver` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SqlserverSpec](#sqlserverspec)_ |  |  |  |
| `status` _[SqlserverStatus](#sqlserverstatus)_ |  |  |  |


#### SqlserverInitParameters







_Appears in:_
- [SqlserverSpec](#sqlserverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MS SQL server service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MS SQL Server. |  |  |
| `release` _string_ | The MS SQL server release of this instance. For convenience, please use gscloud to get the list of available MS SQL server service releases.<br />The MS SQL Server release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MS SQL Server releases. |  |  |
| `s3Backup` _[S3BackupInitParameters](#s3backupinitparameters) array_ | Allow backup/restore MS SQL server to/from a S3 bucket.<br />Allow backup/restore MS SQL server to/from a S3 bucket. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MS SQL Server. |  |  |


#### SqlserverList



SqlserverList contains a list of Sqlservers





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SqlserverList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Sqlserver](#sqlserver) array_ |  |  |  |




#### SqlserverListenPortObservation







_Appears in:_
- [SqlserverObservation](#sqlserverobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### SqlserverObservation







_Appears in:_
- [SqlserverStatus](#sqlserverstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[SqlserverListenPortObservation](#sqlserverlistenportobservation) array_ | The port numbers where this MS SQL server service accepts connections.<br />The port numbers where this MS SQL Server accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of MS SQL server service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MS SQL Server. |  |  |
| `release` _string_ | The MS SQL server release of this instance. For convenience, please use gscloud to get the list of available MS SQL server service releases.<br />The MS SQL Server release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MS SQL Server releases. |  |  |
| `s3Backup` _[S3BackupObservation](#s3backupobservation) array_ | Allow backup/restore MS SQL server to/from a S3 bucket.<br />Allow backup/restore MS SQL server to/from a S3 bucket. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MS SQL Server. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that MS SQL server service uses.<br />PaaS service template that MS SQL Server uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MS SQL Server. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MS SQL Server is in use. |  |  |


#### SqlserverParameters







_Appears in:_
- [SqlserverSpec](#sqlserverspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of MS SQL server service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MS SQL Server. |  | Optional: \{\} <br /> |
| `release` _string_ | The MS SQL server release of this instance. For convenience, please use gscloud to get the list of available MS SQL server service releases.<br />The MS SQL Server release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MS SQL Server releases. |  | Optional: \{\} <br /> |
| `s3Backup` _[S3BackupParameters](#s3backupparameters) array_ | Allow backup/restore MS SQL server to/from a S3 bucket.<br />Allow backup/restore MS SQL server to/from a S3 bucket. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MS SQL Server. |  | Optional: \{\} <br /> |


#### SqlserverSpec



SqlserverSpec defines the desired state of Sqlserver



_Appears in:_
- [Sqlserver](#sqlserver)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[SqlserverParameters](#sqlserverparameters)_ |  |  |  |
| `initProvider` _[SqlserverInitParameters](#sqlserverinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SqlserverStatus



SqlserverStatus defines the observed state of Sqlserver.



_Appears in:_
- [Sqlserver](#sqlserver)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SqlserverObservation](#sqlserverobservation)_ |  |  |  |


#### Sshkey



Sshkey is the Schema for the Sshkeys API. Manages an SSH public key in gridscale.



_Appears in:_
- [SshkeyList](#sshkeylist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Sshkey` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SshkeySpec](#sshkeyspec)_ |  |  |  |
| `status` _[SshkeyStatus](#sshkeystatus)_ |  |  |  |


#### SshkeyInitParameters







_Appears in:_
- [SshkeySpec](#sshkeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sshkey` _string_ | This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).<br />sshkey_string is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1) |  |  |


#### SshkeyList



SshkeyList contains a list of Sshkeys





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SshkeyList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Sshkey](#sshkey) array_ |  |  |  |


#### SshkeyObservation







_Appears in:_
- [SshkeyStatus](#sshkeystatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change |  |  |
| `createTime` _string_ | The time the object was created.<br />The date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sshkey` _string_ | This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).<br />sshkey_string is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1) |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |


#### SshkeyParameters







_Appears in:_
- [SshkeySpec](#sshkeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `sshkey` _string_ | This is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1).<br />sshkey_string is the OpenSSH public key string (all key types are supported => ed25519, ecdsa, dsa, rsa, rsa1) |  | Optional: \{\} <br /> |


#### SshkeySpec



SshkeySpec defines the desired state of Sshkey



_Appears in:_
- [Sshkey](#sshkey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[SshkeyParameters](#sshkeyparameters)_ |  |  |  |
| `initProvider` _[SshkeyInitParameters](#sshkeyinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SshkeyStatus



SshkeyStatus defines the observed state of Sshkey.



_Appears in:_
- [Sshkey](#sshkey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SshkeyObservation](#sshkeyobservation)_ |  |  |  |


#### Storage



Storage is the Schema for the Storages API. Manages a storage in gridscale.



_Appears in:_
- [StorageList](#storagelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Storage` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageSpec](#storagespec)_ |  |  |  |
| `status` _[StorageStatus](#storagestatus)_ |  |  |  |




#### StorageBackupsObservation







_Appears in:_
- [BackupscheduleObservation](#backupscheduleobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `createTime` _string_ |  |  |  |
| `name` _string_ |  |  |  |
| `objectUuid` _string_ |  |  |  |




#### StorageInitParameters







_Appears in:_
- [ServerInitParameters_2](#serverinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `objectUuid` _string_ | The object UUID or id of the storage. |  |  |


#### StorageInitParameters_2







_Appears in:_
- [StorageSpec](#storagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `rollbackFromBackupUuid` _string_ | Rollback the storage from a specific storage backup.<br />Rollback the storage from a specific storage backup. |  |  |
| `storageType` _string_ | (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `storageVariant` _string_ | Storage variant (one of local or distributed). Default: "distributed".<br />Storage variant (one of local or distributed). |  |  |
| `template` _[TemplateInitParameters](#templateinitparameters) array_ | List of labels in the format [ "label1", "label2" ]. |  |  |


#### StorageList



StorageList contains a list of Storages





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Storage](#storage) array_ |  |  |  |


#### StorageObservation







_Appears in:_
- [ServerObservation_2](#serverobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bootdevice` _boolean_ | Make this network the boot device. This can only be set for one network. |  |  |
| `bus` _float_ | The SCSI bus id. The SCSI defines transmission routes like Serial Attached SCSI (SAS), Fibre Channel and iSCSI. Each SCSI device is addressed via a specific number. Each SCSI bus can have multiple SCSI devices connected to it. |  |  |
| `capacity` _float_ | Capacity of the storage (GB). |  |  |
| `controller` _float_ | Defines the SCSI controller id. The SCSI defines transmission routes such as Serial Attached SCSI (SAS), Fibre Channel and iSCSI. |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `lun` _float_ | Is the common SCSI abbreviation of the Logical Unit Number. A lun is a unique identifier for a single disk or a composite of disks. |  |  |
| `objectName` _string_ | Name of the storage. |  |  |
| `objectUuid` _string_ | The object UUID or id of the storage. |  |  |
| `storageType` _string_ | Indicates the speed of the storage. This may be (storage, storage_high or storage_insane). |  |  |
| `target` _float_ | Defines the SCSI target ID. The target ID is a device (e.g. disk). |  |  |


#### StorageObservation_2







_Appears in:_
- [StorageStatus](#storagestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `locationUuid` _string_ | The location this storage is placed. The location of a resource is determined by it's project.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `parentUuid` _string_ |  |  |  |
| `rollbackFromBackupUuid` _string_ | Rollback the storage from a specific storage backup.<br />Rollback the storage from a specific storage backup. |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |
| `storageType` _string_ | (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `storageVariant` _string_ | Storage variant (one of local or distributed). Default: "distributed".<br />Storage variant (one of local or distributed). |  |  |
| `template` _[TemplateObservation](#templateobservation) array_ | List of labels in the format [ "label1", "label2" ]. |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### StorageParameters







_Appears in:_
- [ServerParameters_2](#serverparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `objectUuid` _string_ | The object UUID or id of the storage. |  | Optional: \{\} <br /> |


#### StorageParameters_2







_Appears in:_
- [StorageSpec](#storagespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `rollbackFromBackupUuid` _string_ | Rollback the storage from a specific storage backup.<br />Rollback the storage from a specific storage backup. |  | Optional: \{\} <br /> |
| `storageType` _string_ | (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  | Optional: \{\} <br /> |
| `storageVariant` _string_ | Storage variant (one of local or distributed). Default: "distributed".<br />Storage variant (one of local or distributed). |  | Optional: \{\} <br /> |
| `template` _[TemplateParameters](#templateparameters) array_ | List of labels in the format [ "label1", "label2" ]. |  | Optional: \{\} <br /> |


#### StorageSpec



StorageSpec defines the desired state of Storage



_Appears in:_
- [Storage](#storage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[StorageParameters_2](#storageparameters_2)_ |  |  |  |
| `initProvider` _[StorageInitParameters_2](#storageinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageStatus



StorageStatus defines the observed state of Storage.



_Appears in:_
- [Storage](#storage)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageObservation_2](#storageobservation_2)_ |  |  |  |


#### TaintsInitParameters







_Appears in:_
- [NodePoolInitParameters](#nodepoolinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `effect` _string_ | The effect of the taint. |  |  |
| `key` _string_ | The key of the taint. |  |  |
| `value` _string_ | The value of the taint. |  |  |


#### TaintsObservation







_Appears in:_
- [NodePoolObservation](#nodepoolobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `effect` _string_ | The effect of the taint. |  |  |
| `key` _string_ | The key of the taint. |  |  |
| `value` _string_ | The value of the taint. |  |  |


#### TaintsParameters







_Appears in:_
- [NodePoolParameters](#nodepoolparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `effect` _string_ | The effect of the taint. |  | Optional: \{\} <br /> |
| `key` _string_ | The key of the taint. |  | Optional: \{\} <br /> |
| `value` _string_ | The value of the taint. |  | Optional: \{\} <br /> |


#### Template



Template is the Schema for the Templates API. Manages a template in gridscale.



_Appears in:_
- [TemplateList](#templatelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Template` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[TemplateSpec](#templatespec)_ |  |  |  |
| `status` _[TemplateStatus](#templatestatus)_ |  |  |  |


#### TemplateInitParameters







_Appears in:_
- [StorageInitParameters_2](#storageinitparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostname` _string_ | The hostname of the installed server (ignored for private templates and public windows templates). |  |  |
| `passwordSecretRef` _[SecretKeySelector](#secretkeyselector)_ | The root (Linux) or Administrator (Windows) password to set for the installed storage. Valid only for public templates. The password has to be either plain-text or a crypt string (modular crypt format - MCF). |  |  |
| `passwordType` _string_ | (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates). |  |  |
| `sshkeys` _string array_ | (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user. |  |  |
| `templateUuid` _string_ | The UUID of a template. This can be found in the the page Template by clicking more on the template or by using a gridscale_template datasource. |  |  |


#### TemplateInitParameters_2







_Appears in:_
- [TemplateSpec](#templatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `name` _string_ | The exact name of the template as show in the page Template.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `snapshotUuid` _string_ | Snapshot uuid for template.<br />Snapshot UUID for template. |  |  |


#### TemplateList



TemplateList contains a list of Templates





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `TemplateList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Template](#template) array_ |  |  |  |


#### TemplateObservation







_Appears in:_
- [StorageObservation_2](#storageobservation_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostname` _string_ | The hostname of the installed server (ignored for private templates and public windows templates). |  |  |
| `passwordType` _string_ | (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates). |  |  |
| `sshkeys` _string array_ | (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user. |  |  |
| `templateUuid` _string_ | The UUID of a template. This can be found in the the page Template by clicking more on the template or by using a gridscale_template datasource. |  |  |


#### TemplateObservation_2







_Appears in:_
- [TemplateStatus](#templatestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The capacity of a storage/ISO Image/template/snapshot in GB.<br />The capacity of a storage/ISO image/template/snapshot in GB. |  |  |
| `changeTime` _string_ | The date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `createTime` _string_ | The date and time the object was initially created.<br />The date and time the object was initially created. |  |  |
| `currentPrice` _float_ | Defines the price for the current period since the last bill.<br />Defines the price for the current period since the last bill. |  |  |
| `description` _string_ | Description of the template.<br />Description of the template. |  |  |
| `distro` _string_ | The OS distribution that the template contains.<br />The Linux distribution of this template (if applicable). |  |  |
| `id` _string_ | The UUID of the template. |  |  |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  |  |
| `name` _string_ | The exact name of the template as show in the page Template.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `ostype` _string_ | The operating system installed in the template.<br />The operating system installed in the template. |  |  |
| `private` _boolean_ | The object is private, the value will be true. Otherwise the value will be false.<br />The object is private, the value will be true. Otherwise the value will be false. |  |  |
| `snapshotUuid` _string_ | Snapshot uuid for template.<br />Snapshot UUID for template. |  |  |
| `status` _string_ | Status indicates the status of the object.<br />Status indicates the status of the object. |  |  |
| `usageInMinutes` _float_ | Total minutes the object has been running.<br />Total minutes the object has been running. |  |  |
| `version` _string_ | The version of the template. |  |  |


#### TemplateParameters







_Appears in:_
- [StorageParameters_2](#storageparameters_2)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostname` _string_ | The hostname of the installed server (ignored for private templates and public windows templates). |  | Optional: \{\} <br /> |
| `passwordSecretRef` _[SecretKeySelector](#secretkeyselector)_ | The root (Linux) or Administrator (Windows) password to set for the installed storage. Valid only for public templates. The password has to be either plain-text or a crypt string (modular crypt format - MCF). |  | Optional: \{\} <br /> |
| `passwordType` _string_ | (one of plain, crypt) Required if password is set (ignored for private templates and public Windows templates). |  | Optional: \{\} <br /> |
| `sshkeys` _string array_ | (array of any - minItems: 0) Public Linux templates only! The UUIDs of SSH keys to be added for the root user. |  | Optional: \{\} <br /> |
| `templateUuid` _string_ | The UUID of a template. This can be found in the the page Template by clicking more on the template or by using a gridscale_template datasource. |  | Optional: \{\} <br /> |


#### TemplateParameters_2







_Appears in:_
- [TemplateSpec](#templatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The exact name of the template as show in the page Template.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `snapshotUuid` _string_ | Snapshot uuid for template.<br />Snapshot UUID for template. |  | Optional: \{\} <br /> |


#### TemplateSpec



TemplateSpec defines the desired state of Template



_Appears in:_
- [Template](#template)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[TemplateParameters_2](#templateparameters_2)_ |  |  |  |
| `initProvider` _[TemplateInitParameters_2](#templateinitparameters_2)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### TemplateStatus



TemplateStatus defines the observed state of Template.



_Appears in:_
- [Template](#template)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[TemplateObservation_2](#templateobservation_2)_ |  |  |  |



## gridscale.m.platformrelay.io/v1alpha1

Package v1alpha1 contains the core resources of the gridscale jet provider.




## gridscale.m.platformrelay.io/v1beta1

Package v1beta1 contains the core resources of the gridscale upjet provider.

### Resource Types
- [ClusterProviderConfig](#clusterproviderconfig)
- [ClusterProviderConfigList](#clusterproviderconfiglist)
- [ProviderConfig](#providerconfig)
- [ProviderConfigList](#providerconfiglist)
- [ProviderConfigUsage](#providerconfigusage)
- [ProviderConfigUsageList](#providerconfigusagelist)



#### ClusterProviderConfig



A ClusterProviderConfig configures a Gridscale provider.



_Appears in:_
- [ClusterProviderConfigList](#clusterproviderconfiglist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.m.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ClusterProviderConfig` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ProviderConfigSpec](#providerconfigspec)_ |  |  |  |
| `status` _[ProviderConfigStatus](#providerconfigstatus)_ |  |  |  |


#### ClusterProviderConfigList



ClusterProviderConfigList contains a list of ProviderConfig.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.m.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ClusterProviderConfigList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ClusterProviderConfig](#clusterproviderconfig) array_ |  |  |  |


#### ProviderConfig



A ProviderConfig configures a Gridscale provider.



_Appears in:_
- [ProviderConfigList](#providerconfiglist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.m.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfig` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ProviderConfigSpec](#providerconfigspec)_ |  |  |  |
| `status` _[ProviderConfigStatus](#providerconfigstatus)_ |  |  |  |


#### ProviderConfigList



ProviderConfigList contains a list of ProviderConfig.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.m.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfigList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ProviderConfig](#providerconfig) array_ |  |  |  |


#### ProviderConfigSpec



A ProviderConfigSpec defines the desired state of a ProviderConfig.



_Appears in:_
- [ClusterProviderConfig](#clusterproviderconfig)
- [ProviderConfig](#providerconfig)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `credentials` _[ProviderCredentials](#providercredentials)_ | Credentials required to authenticate to this provider. |  |  |


#### ProviderConfigStatus



A ProviderConfigStatus reflects the observed state of a ProviderConfig.



_Appears in:_
- [ClusterProviderConfig](#clusterproviderconfig)
- [ProviderConfig](#providerconfig)



#### ProviderConfigUsage



A ProviderConfigUsage indicates that a resource is using a ProviderConfig.



_Appears in:_
- [ProviderConfigUsageList](#providerconfigusagelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.m.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfigUsage` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |


#### ProviderConfigUsageList



ProviderConfigUsageList contains a list of ProviderConfigUsage





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.m.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfigUsageList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ProviderConfigUsage](#providerconfigusage) array_ |  |  |  |


#### ProviderCredentials



ProviderCredentials required to authenticate.



_Appears in:_
- [ProviderConfigSpec](#providerconfigspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `source` _[CredentialsSource](#credentialssource)_ | Source of the provider credentials. |  | Enum: [None Secret InjectedIdentity Environment Filesystem] <br /> |



## gridscale.platformrelay.io/v1alpha1

Package v1alpha1 contains the core resources of the gridscale jet provider.




## gridscale.platformrelay.io/v1beta1

Package v1beta1 contains the core resources of the gridscale upjet provider.

### Resource Types
- [ProviderConfig](#providerconfig)
- [ProviderConfigList](#providerconfiglist)
- [ProviderConfigUsage](#providerconfigusage)
- [ProviderConfigUsageList](#providerconfigusagelist)



#### ProviderConfig



A ProviderConfig configures a Gridscale provider.



_Appears in:_
- [ProviderConfigList](#providerconfiglist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfig` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ProviderConfigSpec](#providerconfigspec)_ |  |  |  |
| `status` _[ProviderConfigStatus](#providerconfigstatus)_ |  |  |  |


#### ProviderConfigList



ProviderConfigList contains a list of ProviderConfig.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfigList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ProviderConfig](#providerconfig) array_ |  |  |  |


#### ProviderConfigSpec



A ProviderConfigSpec defines the desired state of a ProviderConfig.



_Appears in:_
- [ProviderConfig](#providerconfig)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `credentials` _[ProviderCredentials](#providercredentials)_ | Credentials required to authenticate to this provider. |  |  |


#### ProviderConfigStatus



A ProviderConfigStatus reflects the observed state of a ProviderConfig.



_Appears in:_
- [ProviderConfig](#providerconfig)



#### ProviderConfigUsage



A ProviderConfigUsage indicates that a resource is using a ProviderConfig.



_Appears in:_
- [ProviderConfigUsageList](#providerconfigusagelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfigUsage` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |


#### ProviderConfigUsageList



ProviderConfigUsageList contains a list of ProviderConfigUsage





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `gridscale.platformrelay.io/v1beta1` | | |
| `kind` _string_ | `ProviderConfigUsageList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ProviderConfigUsage](#providerconfigusage) array_ |  |  |  |


#### ProviderCredentials



ProviderCredentials required to authenticate.



_Appears in:_
- [ProviderConfigSpec](#providerconfigspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `source` _[CredentialsSource](#credentialssource)_ | Source of the provider credentials. |  | Enum: [None Secret InjectedIdentity Environment Filesystem] <br /> |



## marketplace.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [Application](#application)
- [ApplicationImport](#applicationimport)
- [ApplicationImportList](#applicationimportlist)
- [ApplicationList](#applicationlist)



#### Application



Application is the Schema for the Applications API. <no value>



_Appears in:_
- [ApplicationList](#applicationlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Application` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ApplicationSpec](#applicationspec)_ |  |  |  |
| `status` _[ApplicationStatus](#applicationstatus)_ |  |  |  |


#### ApplicationImport



ApplicationImport is the Schema for the ApplicationImports API. <no value>



_Appears in:_
- [ApplicationImportList](#applicationimportlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ApplicationImport` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ApplicationImportSpec](#applicationimportspec)_ |  |  |  |
| `status` _[ApplicationImportStatus](#applicationimportstatus)_ |  |  |  |


#### ApplicationImportInitParameters







_Appears in:_
- [ApplicationImportSpec](#applicationimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `importUniqueHash` _string_ | Hash of a specific marketplace application that you want to import |  |  |


#### ApplicationImportList



ApplicationImportList contains a list of ApplicationImports





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ApplicationImportList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ApplicationImport](#applicationimport) array_ |  |  |  |


#### ApplicationImportObservation







_Appears in:_
- [ApplicationImportStatus](#applicationimportstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `importUniqueHash` _string_ | Hash of a specific marketplace application that you want to import |  |  |
| `isApplicationOwner` _boolean_ | Whether the you are the owner of application or not |  |  |
| `isPublishGlobal` _boolean_ | Whether a template is published to other partner or not |  |  |
| `isPublishGlobalRequested` _boolean_ | Whether a partner wants their tenant template published to other partners |  |  |
| `isPublishRequested` _boolean_ | Whether the tenants want their template to be published or not |  |  |
| `isPublished` _boolean_ | Whether the template is published by the partner to their tenant |  |  |
| `metaAdvices` _string_ |  |  |  |
| `metaAuthor` _string_ |  |  |  |
| `metaComponents` _string array_ |  |  |  |
| `metaFeatures` _string_ |  |  |  |
| `metaHints` _string_ |  |  |  |
| `metaIcon` _string_ |  |  |  |
| `metaLicense` _string_ |  |  |  |
| `metaOs` _string_ |  |  |  |
| `metaOverview` _string_ |  |  |  |
| `metaTermsOfUse` _string_ |  |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  |  |
| `publishGlobalRequestedDate` _string_ | The date when a partner requested their tenants template to be published |  |  |
| `publishRequestedDate` _string_ | The date when the tenant requested their template to be published |  |  |
| `publishedDate` _string_ | The date when the template is published into other tenant in the same partner |  |  |
| `publishedGlobalDate` _string_ | The date when a template is published to other partner |  |  |
| `setupCores` _float_ | Number of server's cores |  |  |
| `setupMemory` _float_ | The capacity of server's memory in GB |  |  |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  |  |
| `status` _string_ | status indicates the status of the object |  |  |
| `type` _string_ | The type of template |  |  |
| `uniqueHash` _string_ | Unique hash to allow user to import the self-created marketplace application |  |  |


#### ApplicationImportParameters







_Appears in:_
- [ApplicationImportSpec](#applicationimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `importUniqueHash` _string_ | Hash of a specific marketplace application that you want to import |  | Optional: \{\} <br /> |


#### ApplicationImportSpec



ApplicationImportSpec defines the desired state of ApplicationImport



_Appears in:_
- [ApplicationImport](#applicationimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[ApplicationImportParameters](#applicationimportparameters)_ |  |  |  |
| `initProvider` _[ApplicationImportInitParameters](#applicationimportinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### ApplicationImportStatus



ApplicationImportStatus defines the observed state of ApplicationImport.



_Appears in:_
- [ApplicationImport](#applicationimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[ApplicationImportObservation](#applicationimportobservation)_ |  |  |  |


#### ApplicationInitParameters







_Appears in:_
- [ApplicationSpec](#applicationspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  |  |
| `metaAdvices` _string_ |  |  |  |
| `metaAuthor` _string_ |  |  |  |
| `metaComponents` _string array_ |  |  |  |
| `metaFeatures` _string_ |  |  |  |
| `metaHints` _string_ |  |  |  |
| `metaIcon` _string_ |  |  |  |
| `metaLicense` _string_ |  |  |  |
| `metaOs` _string_ |  |  |  |
| `metaOverview` _string_ |  |  |  |
| `metaTermsOfUse` _string_ |  |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  |  |
| `publish` _boolean_ | Whether you want to publish your application or not |  |  |
| `setupCores` _float_ | Number of server's cores |  |  |
| `setupMemory` _float_ | The capacity of server's memory in GB |  |  |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  |  |


#### ApplicationList



ApplicationList contains a list of Applications





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ApplicationList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Application](#application) array_ |  |  |  |


#### ApplicationObservation







_Appears in:_
- [ApplicationStatus](#applicationstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `isApplicationOwner` _boolean_ | Whether the you are the owner of application or not |  |  |
| `isPublishGlobal` _boolean_ | Whether a template is published to other partner or not |  |  |
| `isPublishGlobalRequested` _boolean_ | Whether a partner wants their tenant template published to other partners |  |  |
| `isPublishRequested` _boolean_ | Whether the tenants want their template to be published or not |  |  |
| `isPublished` _boolean_ | Whether the template is published by the partner to their tenant |  |  |
| `metaAdvices` _string_ |  |  |  |
| `metaAuthor` _string_ |  |  |  |
| `metaComponents` _string array_ |  |  |  |
| `metaFeatures` _string_ |  |  |  |
| `metaHints` _string_ |  |  |  |
| `metaIcon` _string_ |  |  |  |
| `metaLicense` _string_ |  |  |  |
| `metaOs` _string_ |  |  |  |
| `metaOverview` _string_ |  |  |  |
| `metaTermsOfUse` _string_ |  |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  |  |
| `publish` _boolean_ | Whether you want to publish your application or not |  |  |
| `publishGlobalRequestedDate` _string_ | The date when a partner requested their tenants template to be published |  |  |
| `publishRequestedDate` _string_ | The date when the tenant requested their template to be published |  |  |
| `publishedDate` _string_ | The date when the template is published into other tenant in the same partner |  |  |
| `publishedGlobalDate` _string_ | The date when a template is published to other partner |  |  |
| `setupCores` _float_ | Number of server's cores |  |  |
| `setupMemory` _float_ | The capacity of server's memory in GB |  |  |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  |  |
| `status` _string_ | status indicates the status of the object |  |  |
| `type` _string_ | The type of template |  |  |
| `uniqueHash` _string_ | Unique hash to allow user to import the self-created marketplace application |  |  |


#### ApplicationParameters







_Appears in:_
- [ApplicationSpec](#applicationspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  | Optional: \{\} <br /> |
| `metaAdvices` _string_ |  |  | Optional: \{\} <br /> |
| `metaAuthor` _string_ |  |  | Optional: \{\} <br /> |
| `metaComponents` _string array_ |  |  | Optional: \{\} <br /> |
| `metaFeatures` _string_ |  |  | Optional: \{\} <br /> |
| `metaHints` _string_ |  |  | Optional: \{\} <br /> |
| `metaIcon` _string_ |  |  | Optional: \{\} <br /> |
| `metaLicense` _string_ |  |  | Optional: \{\} <br /> |
| `metaOs` _string_ |  |  | Optional: \{\} <br /> |
| `metaOverview` _string_ |  |  | Optional: \{\} <br /> |
| `metaTermsOfUse` _string_ |  |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  | Optional: \{\} <br /> |
| `publish` _boolean_ | Whether you want to publish your application or not |  | Optional: \{\} <br /> |
| `setupCores` _float_ | Number of server's cores |  | Optional: \{\} <br /> |
| `setupMemory` _float_ | The capacity of server's memory in GB |  | Optional: \{\} <br /> |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  | Optional: \{\} <br /> |


#### ApplicationSpec



ApplicationSpec defines the desired state of Application



_Appears in:_
- [Application](#application)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[ApplicationParameters](#applicationparameters)_ |  |  |  |
| `initProvider` _[ApplicationInitParameters](#applicationinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### ApplicationStatus



ApplicationStatus defines the observed state of Application.



_Appears in:_
- [Application](#application)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[ApplicationObservation](#applicationobservation)_ |  |  |  |



## marketplace.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [Application](#application)
- [ApplicationImport](#applicationimport)
- [ApplicationImportList](#applicationimportlist)
- [ApplicationList](#applicationlist)



#### Application



Application is the Schema for the Applications API. <no value>



_Appears in:_
- [ApplicationList](#applicationlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Application` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ApplicationSpec](#applicationspec)_ |  |  |  |
| `status` _[ApplicationStatus](#applicationstatus)_ |  |  |  |


#### ApplicationImport



ApplicationImport is the Schema for the ApplicationImports API. <no value>



_Appears in:_
- [ApplicationImportList](#applicationimportlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ApplicationImport` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ApplicationImportSpec](#applicationimportspec)_ |  |  |  |
| `status` _[ApplicationImportStatus](#applicationimportstatus)_ |  |  |  |


#### ApplicationImportInitParameters







_Appears in:_
- [ApplicationImportSpec](#applicationimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `importUniqueHash` _string_ | Hash of a specific marketplace application that you want to import |  |  |


#### ApplicationImportList



ApplicationImportList contains a list of ApplicationImports





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ApplicationImportList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ApplicationImport](#applicationimport) array_ |  |  |  |


#### ApplicationImportObservation







_Appears in:_
- [ApplicationImportStatus](#applicationimportstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `importUniqueHash` _string_ | Hash of a specific marketplace application that you want to import |  |  |
| `isApplicationOwner` _boolean_ | Whether the you are the owner of application or not |  |  |
| `isPublishGlobal` _boolean_ | Whether a template is published to other partner or not |  |  |
| `isPublishGlobalRequested` _boolean_ | Whether a partner wants their tenant template published to other partners |  |  |
| `isPublishRequested` _boolean_ | Whether the tenants want their template to be published or not |  |  |
| `isPublished` _boolean_ | Whether the template is published by the partner to their tenant |  |  |
| `metaAdvices` _string_ |  |  |  |
| `metaAuthor` _string_ |  |  |  |
| `metaComponents` _string array_ |  |  |  |
| `metaFeatures` _string_ |  |  |  |
| `metaHints` _string_ |  |  |  |
| `metaIcon` _string_ |  |  |  |
| `metaLicense` _string_ |  |  |  |
| `metaOs` _string_ |  |  |  |
| `metaOverview` _string_ |  |  |  |
| `metaTermsOfUse` _string_ |  |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  |  |
| `publishGlobalRequestedDate` _string_ | The date when a partner requested their tenants template to be published |  |  |
| `publishRequestedDate` _string_ | The date when the tenant requested their template to be published |  |  |
| `publishedDate` _string_ | The date when the template is published into other tenant in the same partner |  |  |
| `publishedGlobalDate` _string_ | The date when a template is published to other partner |  |  |
| `setupCores` _float_ | Number of server's cores |  |  |
| `setupMemory` _float_ | The capacity of server's memory in GB |  |  |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  |  |
| `status` _string_ | status indicates the status of the object |  |  |
| `type` _string_ | The type of template |  |  |
| `uniqueHash` _string_ | Unique hash to allow user to import the self-created marketplace application |  |  |


#### ApplicationImportParameters







_Appears in:_
- [ApplicationImportSpec](#applicationimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `importUniqueHash` _string_ | Hash of a specific marketplace application that you want to import |  | Optional: \{\} <br /> |


#### ApplicationImportSpec



ApplicationImportSpec defines the desired state of ApplicationImport



_Appears in:_
- [ApplicationImport](#applicationimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[ApplicationImportParameters](#applicationimportparameters)_ |  |  |  |
| `initProvider` _[ApplicationImportInitParameters](#applicationimportinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### ApplicationImportStatus



ApplicationImportStatus defines the observed state of ApplicationImport.



_Appears in:_
- [ApplicationImport](#applicationimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[ApplicationImportObservation](#applicationimportobservation)_ |  |  |  |


#### ApplicationInitParameters







_Appears in:_
- [ApplicationSpec](#applicationspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  |  |
| `metaAdvices` _string_ |  |  |  |
| `metaAuthor` _string_ |  |  |  |
| `metaComponents` _string array_ |  |  |  |
| `metaFeatures` _string_ |  |  |  |
| `metaHints` _string_ |  |  |  |
| `metaIcon` _string_ |  |  |  |
| `metaLicense` _string_ |  |  |  |
| `metaOs` _string_ |  |  |  |
| `metaOverview` _string_ |  |  |  |
| `metaTermsOfUse` _string_ |  |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  |  |
| `publish` _boolean_ | Whether you want to publish your application or not |  |  |
| `setupCores` _float_ | Number of server's cores |  |  |
| `setupMemory` _float_ | The capacity of server's memory in GB |  |  |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  |  |


#### ApplicationList



ApplicationList contains a list of Applications





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `marketplace.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `ApplicationList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Application](#application) array_ |  |  |  |


#### ApplicationObservation







_Appears in:_
- [ApplicationStatus](#applicationstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created |  |  |
| `id` _string_ |  |  |  |
| `isApplicationOwner` _boolean_ | Whether the you are the owner of application or not |  |  |
| `isPublishGlobal` _boolean_ | Whether a template is published to other partner or not |  |  |
| `isPublishGlobalRequested` _boolean_ | Whether a partner wants their tenant template published to other partners |  |  |
| `isPublishRequested` _boolean_ | Whether the tenants want their template to be published or not |  |  |
| `isPublished` _boolean_ | Whether the template is published by the partner to their tenant |  |  |
| `metaAdvices` _string_ |  |  |  |
| `metaAuthor` _string_ |  |  |  |
| `metaComponents` _string array_ |  |  |  |
| `metaFeatures` _string_ |  |  |  |
| `metaHints` _string_ |  |  |  |
| `metaIcon` _string_ |  |  |  |
| `metaLicense` _string_ |  |  |  |
| `metaOs` _string_ |  |  |  |
| `metaOverview` _string_ |  |  |  |
| `metaTermsOfUse` _string_ |  |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  |  |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  |  |
| `publish` _boolean_ | Whether you want to publish your application or not |  |  |
| `publishGlobalRequestedDate` _string_ | The date when a partner requested their tenants template to be published |  |  |
| `publishRequestedDate` _string_ | The date when the tenant requested their template to be published |  |  |
| `publishedDate` _string_ | The date when the template is published into other tenant in the same partner |  |  |
| `publishedGlobalDate` _string_ | The date when a template is published to other partner |  |  |
| `setupCores` _float_ | Number of server's cores |  |  |
| `setupMemory` _float_ | The capacity of server's memory in GB |  |  |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  |  |
| `status` _string_ | status indicates the status of the object |  |  |
| `type` _string_ | The type of template |  |  |
| `uniqueHash` _string_ | Unique hash to allow user to import the self-created marketplace application |  |  |


#### ApplicationParameters







_Appears in:_
- [ApplicationSpec](#applicationspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `category` _string_ | Category of marketplace application. Accepted values: "CMS", "project management", "Adminpanel", "Collaboration", "Cloud Storage", "Archiving" |  | Optional: \{\} <br /> |
| `metaAdvices` _string_ |  |  | Optional: \{\} <br /> |
| `metaAuthor` _string_ |  |  | Optional: \{\} <br /> |
| `metaComponents` _string array_ |  |  | Optional: \{\} <br /> |
| `metaFeatures` _string_ |  |  | Optional: \{\} <br /> |
| `metaHints` _string_ |  |  | Optional: \{\} <br /> |
| `metaIcon` _string_ |  |  | Optional: \{\} <br /> |
| `metaLicense` _string_ |  |  | Optional: \{\} <br /> |
| `metaOs` _string_ |  |  | Optional: \{\} <br /> |
| `metaOverview` _string_ |  |  | Optional: \{\} <br /> |
| `metaTermsOfUse` _string_ |  |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters |  | Optional: \{\} <br /> |
| `objectStoragePath` _string_ | Path to the images for the application, must be in .gz format and started with s3// |  | Optional: \{\} <br /> |
| `publish` _boolean_ | Whether you want to publish your application or not |  | Optional: \{\} <br /> |
| `setupCores` _float_ | Number of server's cores |  | Optional: \{\} <br /> |
| `setupMemory` _float_ | The capacity of server's memory in GB |  | Optional: \{\} <br /> |
| `setupStorageCapacity` _float_ | The capacity of server's storage in GB |  | Optional: \{\} <br /> |


#### ApplicationSpec



ApplicationSpec defines the desired state of Application



_Appears in:_
- [Application](#application)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[ApplicationParameters](#applicationparameters)_ |  |  |  |
| `initProvider` _[ApplicationInitParameters](#applicationinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### ApplicationStatus



ApplicationStatus defines the observed state of Application.



_Appears in:_
- [Application](#application)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[ApplicationObservation](#applicationobservation)_ |  |  |  |



## mysql8.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [MySQL8](#mysql8)
- [MySQL8List](#mysql8list)





#### ListenPortObservation







_Appears in:_
- [MySQL8Observation](#mysql8observation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MySQL8



MySQL8 is the Schema for the MySQL8s API. Manage a MySQL 8.0 service in gridscale.



_Appears in:_
- [MySQL8List](#mysql8list)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `mysql8.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQL8` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MySQL8Spec](#mysql8spec)_ |  |  |  |
| `status` _[MySQL8Status](#mysql8status)_ |  |  |  |


#### MySQL8InitParameters







_Appears in:_
- [MySQL8Spec](#mysql8spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |


#### MySQL8List



MySQL8List contains a list of MySQL8s





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `mysql8.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQL8List` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[MySQL8](#mysql8) array_ |  |  |  |


#### MySQL8Observation







_Appears in:_
- [MySQL8Status](#mysql8status)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[ListenPortObservation](#listenportobservation) array_ | The port numbers where this mysql service accepts connections.<br />The port numbers where this MySQL service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that mysql service uses.<br />PaaS service template that MySQL service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MySQL service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MySQL service is in use. |  |  |


#### MySQL8Parameters







_Appears in:_
- [MySQL8Spec](#mysql8spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  | Optional: \{\} <br /> |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  | Optional: \{\} <br /> |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  | Optional: \{\} <br /> |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  | Optional: \{\} <br /> |


#### MySQL8Spec



MySQL8Spec defines the desired state of MySQL8



_Appears in:_
- [MySQL8](#mysql8)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[MySQL8Parameters](#mysql8parameters)_ |  |  |  |
| `initProvider` _[MySQL8InitParameters](#mysql8initparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MySQL8Status



MySQL8Status defines the observed state of MySQL8.



_Appears in:_
- [MySQL8](#mysql8)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MySQL8Observation](#mysql8observation)_ |  |  |  |



## mysql8.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [MySQL8](#mysql8)
- [MySQL8List](#mysql8list)





#### ListenPortObservation







_Appears in:_
- [MySQL8Observation](#mysql8observation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### MySQL8



MySQL8 is the Schema for the MySQL8s API. Manage a MySQL 8.0 service in gridscale.



_Appears in:_
- [MySQL8List](#mysql8list)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `mysql8.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQL8` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[MySQL8Spec](#mysql8spec)_ |  |  |  |
| `status` _[MySQL8Status](#mysql8status)_ |  |  |  |


#### MySQL8InitParameters







_Appears in:_
- [MySQL8Spec](#mysql8spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |


#### MySQL8List



MySQL8List contains a list of MySQL8s





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `mysql8.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `MySQL8List` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[MySQL8](#mysql8) array_ |  |  |  |


#### MySQL8Observation







_Appears in:_
- [MySQL8Status](#mysql8status)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[ListenPortObservation](#listenportobservation) array_ | The port numbers where this mysql service accepts connections.<br />The port numbers where this MySQL service accepts connections. |  |  |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  |  |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  |  |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  |  |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  |  |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  |  |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that mysql service uses.<br />PaaS service template that MySQL service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of MySQL service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that MySQL service is in use. |  |  |


#### MySQL8Parameters







_Appears in:_
- [MySQL8Spec](#mysql8spec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `maxCoreCount` _float_ | Maximum CPU core count. The mysql instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and max_core_count.<br />Maximum CPU core count. The MySQL instance's CPU core count will be autoscaled based on the workload. The number of cores stays between 1 and `max_core_count`. |  | Optional: \{\} <br /> |
| `mysqlDefaultTimeZone` _string_ | mysql parameter: Server Timezone. Default: UTC.<br />Server Timezone. |  | Optional: \{\} <br /> |
| `mysqlMaxAllowedPacket` _string_ | mysql parameter: Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). Default: 64M.<br />Max Allowed Packet Size. Format: xM (where x is an integer, M stands for unit: k(kB), M(MB), G(GB)). |  | Optional: \{\} <br /> |
| `mysqlMaxConnections` _float_ | mysql parameter: Max Connections. Default: 4000.<br />Max Connections. |  | Optional: \{\} <br /> |
| `mysqlSqlMode` _string_ | mysql parameter: SQL Mode. Default: "ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION".<br />SQL Mode. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of mysql service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of MySQL service. |  | Optional: \{\} <br /> |
| `release` _string_ | The mysql release of this instance. For convenience, please use gscloud to get the list of available mysql service releases.<br />The MySQL release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available MySQL service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to MySQL service. |  | Optional: \{\} <br /> |


#### MySQL8Spec



MySQL8Spec defines the desired state of MySQL8



_Appears in:_
- [MySQL8](#mysql8)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[MySQL8Parameters](#mysql8parameters)_ |  |  |  |
| `initProvider` _[MySQL8InitParameters](#mysql8initparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### MySQL8Status



MySQL8Status defines the observed state of MySQL8.



_Appears in:_
- [MySQL8](#mysql8)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[MySQL8Observation](#mysql8observation)_ |  |  |  |



## object.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [StorageAccesskey](#storageaccesskey)
- [StorageAccesskeyList](#storageaccesskeylist)
- [StorageBucket](#storagebucket)
- [StorageBucketList](#storagebucketlist)



#### LifecycleRuleInitParameters







_Appears in:_
- [StorageBucketInitParameters](#storagebucketinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  |  |
| `expirationDays` _float_ |  |  |  |
| `id` _string_ |  |  |  |
| `incompleteUploadExpirationDays` _float_ |  |  |  |
| `noncurrentVersionExpirationDays` _float_ |  |  |  |
| `prefix` _string_ |  |  |  |


#### LifecycleRuleObservation







_Appears in:_
- [StorageBucketObservation](#storagebucketobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  |  |
| `expirationDays` _float_ |  |  |  |
| `id` _string_ |  |  |  |
| `incompleteUploadExpirationDays` _float_ |  |  |  |
| `noncurrentVersionExpirationDays` _float_ |  |  |  |
| `prefix` _string_ |  |  |  |


#### LifecycleRuleParameters







_Appears in:_
- [StorageBucketParameters](#storagebucketparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  | Optional: \{\} <br /> |
| `expirationDays` _float_ |  |  | Optional: \{\} <br /> |
| `id` _string_ |  |  | Optional: \{\} <br /> |
| `incompleteUploadExpirationDays` _float_ |  |  | Optional: \{\} <br /> |
| `noncurrentVersionExpirationDays` _float_ |  |  | Optional: \{\} <br /> |
| `prefix` _string_ |  |  | Optional: \{\} <br /> |


#### StorageAccesskey



StorageAccesskey is the Schema for the StorageAccesskeys API. <no value>



_Appears in:_
- [StorageAccesskeyList](#storageaccesskeylist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageAccesskey` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageAccesskeySpec](#storageaccesskeyspec)_ |  |  |  |
| `status` _[StorageAccesskeyStatus](#storageaccesskeystatus)_ |  |  |  |


#### StorageAccesskeyInitParameters







_Appears in:_
- [StorageAccesskeySpec](#storageaccesskeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `comment` _string_ | Comment for the access_key. |  |  |
| `userUuid` _string_ | If a user_uuid is set, a user-specific key will get created.<br />If no user_uuid is set along a user with write-access to the contract will still only create<br />a user-specific key for themselves while a user with admin-access to the contract will create<br />a contract-level admin key. |  |  |


#### StorageAccesskeyList



StorageAccesskeyList contains a list of StorageAccesskeys





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageAccesskeyList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[StorageAccesskey](#storageaccesskey) array_ |  |  |  |


#### StorageAccesskeyObservation







_Appears in:_
- [StorageAccesskeyStatus](#storageaccesskeystatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `comment` _string_ | Comment for the access_key. |  |  |
| `id` _string_ |  |  |  |
| `userUuid` _string_ | If a user_uuid is set, a user-specific key will get created.<br />If no user_uuid is set along a user with write-access to the contract will still only create<br />a user-specific key for themselves while a user with admin-access to the contract will create<br />a contract-level admin key. |  |  |


#### StorageAccesskeyParameters







_Appears in:_
- [StorageAccesskeySpec](#storageaccesskeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `comment` _string_ | Comment for the access_key. |  | Optional: \{\} <br /> |
| `userUuid` _string_ | If a user_uuid is set, a user-specific key will get created.<br />If no user_uuid is set along a user with write-access to the contract will still only create<br />a user-specific key for themselves while a user with admin-access to the contract will create<br />a contract-level admin key. |  | Optional: \{\} <br /> |


#### StorageAccesskeySpec



StorageAccesskeySpec defines the desired state of StorageAccesskey



_Appears in:_
- [StorageAccesskey](#storageaccesskey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[StorageAccesskeyParameters](#storageaccesskeyparameters)_ |  |  |  |
| `initProvider` _[StorageAccesskeyInitParameters](#storageaccesskeyinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageAccesskeyStatus



StorageAccesskeyStatus defines the observed state of StorageAccesskey.



_Appears in:_
- [StorageAccesskey](#storageaccesskey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageAccesskeyObservation](#storageaccesskeyobservation)_ |  |  |  |


#### StorageBucket



StorageBucket is the Schema for the StorageBuckets API. <no value>



_Appears in:_
- [StorageBucketList](#storagebucketlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageBucket` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageBucketSpec](#storagebucketspec)_ |  |  |  |
| `status` _[StorageBucketStatus](#storagebucketstatus)_ |  |  |  |


#### StorageBucketInitParameters







_Appears in:_
- [StorageBucketSpec](#storagebucketspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The object storage secret_key. |  |  |
| `bucketName` _string_ | The name of the bucket. |  |  |
| `lifecycleRule` _[LifecycleRuleInitParameters](#lifecycleruleinitparameters) array_ |  |  |  |
| `s3Host` _string_ | The S3 host. |  |  |
| `secretKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The object storage access_key. |  |  |


#### StorageBucketList



StorageBucketList contains a list of StorageBuckets





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageBucketList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[StorageBucket](#storagebucket) array_ |  |  |  |


#### StorageBucketObservation







_Appears in:_
- [StorageBucketStatus](#storagebucketstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bucketName` _string_ | The name of the bucket. |  |  |
| `id` _string_ |  |  |  |
| `lifecycleRule` _[LifecycleRuleObservation](#lifecycleruleobservation) array_ |  |  |  |
| `s3Host` _string_ | The S3 host. |  |  |


#### StorageBucketParameters







_Appears in:_
- [StorageBucketSpec](#storagebucketspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The object storage secret_key. |  | Optional: \{\} <br /> |
| `bucketName` _string_ | The name of the bucket. |  | Optional: \{\} <br /> |
| `lifecycleRule` _[LifecycleRuleParameters](#lifecycleruleparameters) array_ |  |  | Optional: \{\} <br /> |
| `s3Host` _string_ | The S3 host. |  | Optional: \{\} <br /> |
| `secretKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The object storage access_key. |  | Optional: \{\} <br /> |


#### StorageBucketSpec



StorageBucketSpec defines the desired state of StorageBucket



_Appears in:_
- [StorageBucket](#storagebucket)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[StorageBucketParameters](#storagebucketparameters)_ |  |  |  |
| `initProvider` _[StorageBucketInitParameters](#storagebucketinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageBucketStatus



StorageBucketStatus defines the observed state of StorageBucket.



_Appears in:_
- [StorageBucket](#storagebucket)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageBucketObservation](#storagebucketobservation)_ |  |  |  |



## object.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [StorageAccesskey](#storageaccesskey)
- [StorageAccesskeyList](#storageaccesskeylist)
- [StorageBucket](#storagebucket)
- [StorageBucketList](#storagebucketlist)



#### LifecycleRuleInitParameters







_Appears in:_
- [StorageBucketInitParameters](#storagebucketinitparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  |  |
| `expirationDays` _float_ |  |  |  |
| `id` _string_ |  |  |  |
| `incompleteUploadExpirationDays` _float_ |  |  |  |
| `noncurrentVersionExpirationDays` _float_ |  |  |  |
| `prefix` _string_ |  |  |  |


#### LifecycleRuleObservation







_Appears in:_
- [StorageBucketObservation](#storagebucketobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  |  |
| `expirationDays` _float_ |  |  |  |
| `id` _string_ |  |  |  |
| `incompleteUploadExpirationDays` _float_ |  |  |  |
| `noncurrentVersionExpirationDays` _float_ |  |  |  |
| `prefix` _string_ |  |  |  |


#### LifecycleRuleParameters







_Appears in:_
- [StorageBucketParameters](#storagebucketparameters)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  | Optional: \{\} <br /> |
| `expirationDays` _float_ |  |  | Optional: \{\} <br /> |
| `id` _string_ |  |  | Optional: \{\} <br /> |
| `incompleteUploadExpirationDays` _float_ |  |  | Optional: \{\} <br /> |
| `noncurrentVersionExpirationDays` _float_ |  |  | Optional: \{\} <br /> |
| `prefix` _string_ |  |  | Optional: \{\} <br /> |


#### StorageAccesskey



StorageAccesskey is the Schema for the StorageAccesskeys API. <no value>



_Appears in:_
- [StorageAccesskeyList](#storageaccesskeylist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageAccesskey` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageAccesskeySpec](#storageaccesskeyspec)_ |  |  |  |
| `status` _[StorageAccesskeyStatus](#storageaccesskeystatus)_ |  |  |  |


#### StorageAccesskeyInitParameters







_Appears in:_
- [StorageAccesskeySpec](#storageaccesskeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `comment` _string_ | Comment for the access_key. |  |  |
| `userUuid` _string_ | If a user_uuid is set, a user-specific key will get created.<br />If no user_uuid is set along a user with write-access to the contract will still only create<br />a user-specific key for themselves while a user with admin-access to the contract will create<br />a contract-level admin key. |  |  |


#### StorageAccesskeyList



StorageAccesskeyList contains a list of StorageAccesskeys





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageAccesskeyList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[StorageAccesskey](#storageaccesskey) array_ |  |  |  |


#### StorageAccesskeyObservation







_Appears in:_
- [StorageAccesskeyStatus](#storageaccesskeystatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `comment` _string_ | Comment for the access_key. |  |  |
| `id` _string_ |  |  |  |
| `userUuid` _string_ | If a user_uuid is set, a user-specific key will get created.<br />If no user_uuid is set along a user with write-access to the contract will still only create<br />a user-specific key for themselves while a user with admin-access to the contract will create<br />a contract-level admin key. |  |  |


#### StorageAccesskeyParameters







_Appears in:_
- [StorageAccesskeySpec](#storageaccesskeyspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `comment` _string_ | Comment for the access_key. |  | Optional: \{\} <br /> |
| `userUuid` _string_ | If a user_uuid is set, a user-specific key will get created.<br />If no user_uuid is set along a user with write-access to the contract will still only create<br />a user-specific key for themselves while a user with admin-access to the contract will create<br />a contract-level admin key. |  | Optional: \{\} <br /> |


#### StorageAccesskeySpec



StorageAccesskeySpec defines the desired state of StorageAccesskey



_Appears in:_
- [StorageAccesskey](#storageaccesskey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[StorageAccesskeyParameters](#storageaccesskeyparameters)_ |  |  |  |
| `initProvider` _[StorageAccesskeyInitParameters](#storageaccesskeyinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageAccesskeyStatus



StorageAccesskeyStatus defines the observed state of StorageAccesskey.



_Appears in:_
- [StorageAccesskey](#storageaccesskey)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageAccesskeyObservation](#storageaccesskeyobservation)_ |  |  |  |


#### StorageBucket



StorageBucket is the Schema for the StorageBuckets API. <no value>



_Appears in:_
- [StorageBucketList](#storagebucketlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageBucket` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageBucketSpec](#storagebucketspec)_ |  |  |  |
| `status` _[StorageBucketStatus](#storagebucketstatus)_ |  |  |  |


#### StorageBucketInitParameters







_Appears in:_
- [StorageBucketSpec](#storagebucketspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | The object storage secret_key. |  |  |
| `bucketName` _string_ | The name of the bucket. |  |  |
| `lifecycleRule` _[LifecycleRuleInitParameters](#lifecycleruleinitparameters) array_ |  |  |  |
| `s3Host` _string_ | The S3 host. |  |  |
| `secretKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | The object storage access_key. |  |  |


#### StorageBucketList



StorageBucketList contains a list of StorageBuckets





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `object.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageBucketList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[StorageBucket](#storagebucket) array_ |  |  |  |


#### StorageBucketObservation







_Appears in:_
- [StorageBucketStatus](#storagebucketstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `bucketName` _string_ | The name of the bucket. |  |  |
| `id` _string_ |  |  |  |
| `lifecycleRule` _[LifecycleRuleObservation](#lifecycleruleobservation) array_ |  |  |  |
| `s3Host` _string_ | The S3 host. |  |  |


#### StorageBucketParameters







_Appears in:_
- [StorageBucketSpec](#storagebucketspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `accessKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | The object storage secret_key. |  | Optional: \{\} <br /> |
| `bucketName` _string_ | The name of the bucket. |  | Optional: \{\} <br /> |
| `lifecycleRule` _[LifecycleRuleParameters](#lifecycleruleparameters) array_ |  |  | Optional: \{\} <br /> |
| `s3Host` _string_ | The S3 host. |  | Optional: \{\} <br /> |
| `secretKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | The object storage access_key. |  | Optional: \{\} <br /> |


#### StorageBucketSpec



StorageBucketSpec defines the desired state of StorageBucket



_Appears in:_
- [StorageBucket](#storagebucket)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[StorageBucketParameters](#storagebucketparameters)_ |  |  |  |
| `initProvider` _[StorageBucketInitParameters](#storagebucketinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageBucketStatus



StorageBucketStatus defines the observed state of StorageBucket.



_Appears in:_
- [StorageBucket](#storagebucket)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageBucketObservation](#storagebucketobservation)_ |  |  |  |



## paas.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [Securityzone](#securityzone)
- [SecurityzoneList](#securityzonelist)



#### Securityzone



Securityzone is the Schema for the Securityzones API. Manages a security zone in gridscale.



_Appears in:_
- [SecurityzoneList](#securityzonelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `paas.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Securityzone` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SecurityzoneSpec](#securityzonespec)_ |  |  |  |
| `status` _[SecurityzoneStatus](#securityzonestatus)_ |  |  |  |


#### SecurityzoneInitParameters







_Appears in:_
- [SecurityzoneSpec](#securityzonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object |  |  |


#### SecurityzoneList



SecurityzoneList contains a list of Securityzones





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `paas.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SecurityzoneList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Securityzone](#securityzone) array_ |  |  |  |


#### SecurityzoneObservation







_Appears in:_
- [SecurityzoneStatus](#securityzonestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created.<br />Defines the date and time the object was initially created |  |  |
| `id` _string_ | The UUID of the security zone. |  |  |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `locationCountry` _string_ | The human-readable name of the location's country.<br />The human-readable name of the location |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The human-readable name of the location.<br />The human-readable name of the location |  |  |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object |  |  |
| `relations` _string array_ | List of PaaS services' UUIDs relating to the security zone.<br />List of PaaS services' UUIDs relating to the security zone |  |  |
| `status` _string_ | Status indicates the status of the object.<br />Status indicates the status of the object |  |  |


#### SecurityzoneParameters







_Appears in:_
- [SecurityzoneSpec](#securityzonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  | Optional: \{\} <br /> |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object |  | Optional: \{\} <br /> |


#### SecurityzoneSpec



SecurityzoneSpec defines the desired state of Securityzone



_Appears in:_
- [Securityzone](#securityzone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[SecurityzoneParameters](#securityzoneparameters)_ |  |  |  |
| `initProvider` _[SecurityzoneInitParameters](#securityzoneinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SecurityzoneStatus



SecurityzoneStatus defines the observed state of Securityzone.



_Appears in:_
- [Securityzone](#securityzone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SecurityzoneObservation](#securityzoneobservation)_ |  |  |  |



## paas.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [Securityzone](#securityzone)
- [SecurityzoneList](#securityzonelist)



#### Securityzone



Securityzone is the Schema for the Securityzones API. Manages a security zone in gridscale.



_Appears in:_
- [SecurityzoneList](#securityzonelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `paas.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Securityzone` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[SecurityzoneSpec](#securityzonespec)_ |  |  |  |
| `status` _[SecurityzoneStatus](#securityzonestatus)_ |  |  |  |


#### SecurityzoneInitParameters







_Appears in:_
- [SecurityzoneSpec](#securityzonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object |  |  |


#### SecurityzoneList



SecurityzoneList contains a list of Securityzones





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `paas.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `SecurityzoneList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Securityzone](#securityzone) array_ |  |  |  |


#### SecurityzoneObservation







_Appears in:_
- [SecurityzoneStatus](#securityzonestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change |  |  |
| `createTime` _string_ | Defines the date and time the object was initially created.<br />Defines the date and time the object was initially created |  |  |
| `id` _string_ | The UUID of the security zone. |  |  |
| `labels` _string array_ | List of labels.<br />List of labels. |  |  |
| `locationCountry` _string_ | The human-readable name of the location's country.<br />The human-readable name of the location |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier |  |  |
| `locationName` _string_ | The human-readable name of the location.<br />The human-readable name of the location |  |  |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object |  |  |
| `relations` _string array_ | List of PaaS services' UUIDs relating to the security zone.<br />List of PaaS services' UUIDs relating to the security zone |  |  |
| `status` _string_ | Status indicates the status of the object.<br />Status indicates the status of the object |  |  |


#### SecurityzoneParameters







_Appears in:_
- [SecurityzoneSpec](#securityzonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels.<br />List of labels. |  | Optional: \{\} <br /> |
| `locationUuid` _string_ | The location this object is placed.<br />The location this object is placed. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object |  | Optional: \{\} <br /> |


#### SecurityzoneSpec



SecurityzoneSpec defines the desired state of Securityzone



_Appears in:_
- [Securityzone](#securityzone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[SecurityzoneParameters](#securityzoneparameters)_ |  |  |  |
| `initProvider` _[SecurityzoneInitParameters](#securityzoneinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### SecurityzoneStatus



SecurityzoneStatus defines the observed state of Securityzone.



_Appears in:_
- [Securityzone](#securityzone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[SecurityzoneObservation](#securityzoneobservation)_ |  |  |  |



## redis.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [Cache](#cache)
- [CacheList](#cachelist)
- [Store](#store)
- [StoreList](#storelist)



#### Cache



Cache is the Schema for the Caches API. Manage a Redis cache service in gridscale.



_Appears in:_
- [CacheList](#cachelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Cache` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[CacheSpec](#cachespec)_ |  |  |  |
| `status` _[CacheStatus](#cachestatus)_ |  |  |  |


#### CacheInitParameters







_Appears in:_
- [CacheSpec](#cachespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis cache service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis cache service. |  |  |
| `release` _string_ | The Redis cache release of this instance. For convenience, please use gscloud to get the list of available Redis cache service releases.<br />The RedisCache release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis cache service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis cache service. |  |  |


#### CacheList



CacheList contains a list of Caches





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `CacheList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Cache](#cache) array_ |  |  |  |


#### CacheObservation







_Appears in:_
- [CacheStatus](#cachestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[ListenPortObservation](#listenportobservation) array_ | The port numbers where this Redis cache service accepts connections.<br />The port numbers where this Redis cache service accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis cache service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis cache service. |  |  |
| `release` _string_ | The Redis cache release of this instance. For convenience, please use gscloud to get the list of available Redis cache service releases.<br />The RedisCache release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis cache service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis cache service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Redis cache service uses.<br />PaaS service template that Redis cache service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of Redis cache service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that Redis cache service is in use. |  |  |


#### CacheParameters







_Appears in:_
- [CacheSpec](#cachespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Redis cache service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis cache service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Redis cache release of this instance. For convenience, please use gscloud to get the list of available Redis cache service releases.<br />The RedisCache release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis cache service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis cache service. |  | Optional: \{\} <br /> |


#### CacheSpec



CacheSpec defines the desired state of Cache



_Appears in:_
- [Cache](#cache)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[CacheParameters](#cacheparameters)_ |  |  |  |
| `initProvider` _[CacheInitParameters](#cacheinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### CacheStatus



CacheStatus defines the observed state of Cache.



_Appears in:_
- [Cache](#cache)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[CacheObservation](#cacheobservation)_ |  |  |  |




#### ListenPortObservation







_Appears in:_
- [CacheObservation](#cacheobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### Store



Store is the Schema for the Stores API. Manage a Redis store service in gridscale.



_Appears in:_
- [StoreList](#storelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Store` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StoreSpec](#storespec)_ |  |  |  |
| `status` _[StoreStatus](#storestatus)_ |  |  |  |


#### StoreInitParameters







_Appears in:_
- [StoreSpec](#storespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis store service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis store service. |  |  |
| `release` _string_ | The Redis store release of this instance. For convenience, please use gscloud to get the list of available Redis store service releases.<br />The RedisStore release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis store service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis store service. |  |  |


#### StoreList



StoreList contains a list of Stores





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StoreList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Store](#store) array_ |  |  |  |




#### StoreListenPortObservation







_Appears in:_
- [StoreObservation](#storeobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### StoreObservation







_Appears in:_
- [StoreStatus](#storestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[StoreListenPortObservation](#storelistenportobservation) array_ | The port numbers where this Redis store service accepts connections.<br />The port numbers where this Redis store service accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis store service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis store service. |  |  |
| `release` _string_ | The Redis store release of this instance. For convenience, please use gscloud to get the list of available Redis store service releases.<br />The RedisStore release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis store service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis store service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Redis store service uses.<br />PaaS service template that Redis store service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of Redis store service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that Redis store service is in use. |  |  |


#### StoreParameters







_Appears in:_
- [StoreSpec](#storespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Redis store service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis store service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Redis store release of this instance. For convenience, please use gscloud to get the list of available Redis store service releases.<br />The RedisStore release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis store service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis store service. |  | Optional: \{\} <br /> |


#### StoreSpec



StoreSpec defines the desired state of Store



_Appears in:_
- [Store](#store)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[StoreParameters](#storeparameters)_ |  |  |  |
| `initProvider` _[StoreInitParameters](#storeinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StoreStatus



StoreStatus defines the observed state of Store.



_Appears in:_
- [Store](#store)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StoreObservation](#storeobservation)_ |  |  |  |



## redis.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [Cache](#cache)
- [CacheList](#cachelist)
- [Store](#store)
- [StoreList](#storelist)



#### Cache



Cache is the Schema for the Caches API. Manage a Redis cache service in gridscale.



_Appears in:_
- [CacheList](#cachelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Cache` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[CacheSpec](#cachespec)_ |  |  |  |
| `status` _[CacheStatus](#cachestatus)_ |  |  |  |


#### CacheInitParameters







_Appears in:_
- [CacheSpec](#cachespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis cache service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis cache service. |  |  |
| `release` _string_ | The Redis cache release of this instance. For convenience, please use gscloud to get the list of available Redis cache service releases.<br />The RedisCache release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis cache service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis cache service. |  |  |


#### CacheList



CacheList contains a list of Caches





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `CacheList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Cache](#cache) array_ |  |  |  |


#### CacheObservation







_Appears in:_
- [CacheStatus](#cachestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[ListenPortObservation](#listenportobservation) array_ | The port numbers where this Redis cache service accepts connections.<br />The port numbers where this Redis cache service accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis cache service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis cache service. |  |  |
| `release` _string_ | The Redis cache release of this instance. For convenience, please use gscloud to get the list of available Redis cache service releases.<br />The RedisCache release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis cache service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis cache service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Redis cache service uses.<br />PaaS service template that Redis cache service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of Redis cache service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that Redis cache service is in use. |  |  |


#### CacheParameters







_Appears in:_
- [CacheSpec](#cachespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Redis cache service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis cache service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Redis cache release of this instance. For convenience, please use gscloud to get the list of available Redis cache service releases.<br />The RedisCache release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis cache service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis cache service. |  | Optional: \{\} <br /> |


#### CacheSpec



CacheSpec defines the desired state of Cache



_Appears in:_
- [Cache](#cache)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[CacheParameters](#cacheparameters)_ |  |  |  |
| `initProvider` _[CacheInitParameters](#cacheinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### CacheStatus



CacheStatus defines the observed state of Cache.



_Appears in:_
- [Cache](#cache)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[CacheObservation](#cacheobservation)_ |  |  |  |




#### ListenPortObservation







_Appears in:_
- [CacheObservation](#cacheobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### Store



Store is the Schema for the Stores API. Manage a Redis store service in gridscale.



_Appears in:_
- [StoreList](#storelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Store` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StoreSpec](#storespec)_ |  |  |  |
| `status` _[StoreStatus](#storestatus)_ |  |  |  |


#### StoreInitParameters







_Appears in:_
- [StoreSpec](#storespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis store service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis store service. |  |  |
| `release` _string_ | The Redis store release of this instance. For convenience, please use gscloud to get the list of available Redis store service releases.<br />The RedisStore release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis store service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis store service. |  |  |


#### StoreList



StoreList contains a list of Stores





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `redis.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StoreList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Store](#store) array_ |  |  |  |




#### StoreListenPortObservation







_Appears in:_
- [StoreObservation](#storeobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `host` _string_ | Host address. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `port` _float_ |  |  |  |




#### StoreObservation







_Appears in:_
- [StoreStatus](#storestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Time of the last change.<br />Time of the last change. |  |  |
| `createTime` _string_ | Date time this service has been created.<br />Date time this service has been created. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `listenPort` _[StoreListenPortObservation](#storelistenportobservation) array_ | The port numbers where this Redis store service accepts connections.<br />The port numbers where this Redis store service accepts connections. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  |  |
| `performanceClass` _string_ | Performance class of Redis store service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis store service. |  |  |
| `release` _string_ | The Redis store release of this instance. For convenience, please use gscloud to get the list of available Redis store service releases.<br />The RedisStore release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis store service releases. |  |  |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis store service. |  |  |
| `serviceTemplateCategory` _string_ | The template service's category used to create the service.<br />The template service's category used to create the service. |  |  |
| `serviceTemplateUuid` _string_ | PaaS service template that Redis store service uses.<br />PaaS service template that Redis store service uses. |  |  |
| `status` _string_ | Current status of PaaS service.<br />Current status of Redis store service. |  |  |
| `usageInMinutes` _float_ | Number of minutes that PaaS service is in use.<br />Number of minutes that Redis store service is in use. |  |  |


#### StoreParameters







_Appears in:_
- [StoreSpec](#storespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `networkUuid` _string_ | The UUID of the network that the service is attached to.<br />The UUID of the network that the service is attached to. |  | Optional: \{\} <br /> |
| `performanceClass` _string_ | Performance class of Redis store service. Available performance classes at the time of writing: standard, high, insane, ultra.<br />Performance class of Redis store service. |  | Optional: \{\} <br /> |
| `release` _string_ | The Redis store release of this instance. For convenience, please use gscloud to get the list of available Redis store service releases.<br />The RedisStore release of this instance.\n<br />For convenience, please use gscloud https://github.com/gridscale/gscloud to get the list of available Redis store service releases. |  | Optional: \{\} <br /> |
| `securityZoneUuid` _string_ | DEPRECATED  The UUID of the security zone that the service is attached to.<br />Security zone UUID linked to Redis store service. |  | Optional: \{\} <br /> |


#### StoreSpec



StoreSpec defines the desired state of Store



_Appears in:_
- [Store](#store)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[StoreParameters](#storeparameters)_ |  |  |  |
| `initProvider` _[StoreInitParameters](#storeinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StoreStatus



StoreStatus defines the observed state of Store.



_Appears in:_
- [Store](#store)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StoreObservation](#storeobservation)_ |  |  |  |



## ssl.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [Certificate](#certificate)
- [CertificateList](#certificatelist)



#### Certificate



Certificate is the Schema for the Certificates API. Manages a TLS/SSL certificate resource in gridscale.



_Appears in:_
- [CertificateList](#certificatelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `ssl.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Certificate` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[CertificateSpec](#certificatespec)_ |  |  |  |
| `status` _[CertificateStatus](#certificatestatus)_ |  |  |  |


#### CertificateInitParameters







_Appears in:_
- [CertificateSpec](#certificatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateChainSecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate.<br />The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `leafCertificateSecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The PEM-formatted public SSL of the SSL certificate.<br />The PEM-formatted public SSL of the SSL certificate. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `privateKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The PEM-formatted private-key of the SSL certificate.<br />The PEM-formatted private-key of the SSL certificate. |  |  |


#### CertificateList



CertificateList contains a list of Certificates





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `ssl.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `CertificateList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Certificate](#certificate) array_ |  |  |  |


#### CertificateObservation







_Appears in:_
- [CertificateStatus](#certificatestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `commonName` _string_ | The common domain name of the SSL certificate.<br />The common domain name of the SSL certificate. |  |  |
| `createTime` _string_ | The date and time the object was initially created.<br />The date and time the object was initially created. |  |  |
| `fingerprints` _[FingerprintsObservation](#fingerprintsobservation) array_ | Defines a list of unique identifiers generated from the MD5, SHA-1, and SHA-256 fingerprints of the certificate.<br />Defines a list of unique identifiers generated from the MD5, SHA-1, and SHA-256 fingerprints of the certificate. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `notValidAfter` _string_ | Defines the date after which the certificate is not valid.<br />Defines the date after which the certificate is not valid. |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |


#### CertificateParameters







_Appears in:_
- [CertificateSpec](#certificatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateChainSecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate.<br />The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `leafCertificateSecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The PEM-formatted public SSL of the SSL certificate.<br />The PEM-formatted public SSL of the SSL certificate. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `privateKeySecretRef` _[LocalSecretKeySelector](#localsecretkeyselector)_ | The PEM-formatted private-key of the SSL certificate.<br />The PEM-formatted private-key of the SSL certificate. |  | Optional: \{\} <br /> |


#### CertificateSpec



CertificateSpec defines the desired state of Certificate



_Appears in:_
- [Certificate](#certificate)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[CertificateParameters](#certificateparameters)_ |  |  |  |
| `initProvider` _[CertificateInitParameters](#certificateinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### CertificateStatus



CertificateStatus defines the observed state of Certificate.



_Appears in:_
- [Certificate](#certificate)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[CertificateObservation](#certificateobservation)_ |  |  |  |




#### FingerprintsObservation







_Appears in:_
- [CertificateObservation](#certificateobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `md5` _string_ | MD5 fingerprint of the certificate. |  |  |
| `sha1` _string_ | SHA1 fingerprint of the certificate. |  |  |
| `sha256` _string_ | SHA256 fingerprint of the certificate. |  |  |





## ssl.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [Certificate](#certificate)
- [CertificateList](#certificatelist)



#### Certificate



Certificate is the Schema for the Certificates API. Manages a TLS/SSL certificate resource in gridscale.



_Appears in:_
- [CertificateList](#certificatelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `ssl.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Certificate` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[CertificateSpec](#certificatespec)_ |  |  |  |
| `status` _[CertificateStatus](#certificatestatus)_ |  |  |  |


#### CertificateInitParameters







_Appears in:_
- [CertificateSpec](#certificatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateChainSecretRef` _[SecretKeySelector](#secretkeyselector)_ | The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate.<br />The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `leafCertificateSecretRef` _[SecretKeySelector](#secretkeyselector)_ | The PEM-formatted public SSL of the SSL certificate.<br />The PEM-formatted public SSL of the SSL certificate. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `privateKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | The PEM-formatted private-key of the SSL certificate.<br />The PEM-formatted private-key of the SSL certificate. |  |  |


#### CertificateList



CertificateList contains a list of Certificates





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `ssl.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `CertificateList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Certificate](#certificate) array_ |  |  |  |


#### CertificateObservation







_Appears in:_
- [CertificateStatus](#certificatestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />The date and time of the last object change. |  |  |
| `commonName` _string_ | The common domain name of the SSL certificate.<br />The common domain name of the SSL certificate. |  |  |
| `createTime` _string_ | The date and time the object was initially created.<br />The date and time the object was initially created. |  |  |
| `fingerprints` _[FingerprintsObservation](#fingerprintsobservation) array_ | Defines a list of unique identifiers generated from the MD5, SHA-1, and SHA-256 fingerprints of the certificate.<br />Defines a list of unique identifiers generated from the MD5, SHA-1, and SHA-256 fingerprints of the certificate. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `notValidAfter` _string_ | Defines the date after which the certificate is not valid.<br />Defines the date after which the certificate is not valid. |  |  |
| `status` _string_ | status indicates the status of the object. |  |  |


#### CertificateParameters







_Appears in:_
- [CertificateSpec](#certificatespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `certificateChainSecretRef` _[SecretKeySelector](#secretkeyselector)_ | The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate.<br />The PEM-formatted full-chain between the certificate authority and the domain's SSL certificate. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `leafCertificateSecretRef` _[SecretKeySelector](#secretkeyselector)_ | The PEM-formatted public SSL of the SSL certificate.<br />The PEM-formatted public SSL of the SSL certificate. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `privateKeySecretRef` _[SecretKeySelector](#secretkeyselector)_ | The PEM-formatted private-key of the SSL certificate.<br />The PEM-formatted private-key of the SSL certificate. |  | Optional: \{\} <br /> |


#### CertificateSpec



CertificateSpec defines the desired state of Certificate



_Appears in:_
- [Certificate](#certificate)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[CertificateParameters](#certificateparameters)_ |  |  |  |
| `initProvider` _[CertificateInitParameters](#certificateinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### CertificateStatus



CertificateStatus defines the observed state of Certificate.



_Appears in:_
- [Certificate](#certificate)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[CertificateObservation](#certificateobservation)_ |  |  |  |




#### FingerprintsObservation







_Appears in:_
- [CertificateObservation](#certificateobservation)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `md5` _string_ | MD5 fingerprint of the certificate. |  |  |
| `sha1` _string_ | SHA1 fingerprint of the certificate. |  |  |
| `sha256` _string_ | SHA256 fingerprint of the certificate. |  |  |





## storage.gridscale.m.platformrelay.io/v1alpha1


### Resource Types
- [Clone](#clone)
- [CloneList](#clonelist)
- [StorageImport](#storageimport)
- [StorageImportList](#storageimportlist)



#### Clone



Clone is the Schema for the Clones API. Make a clone of an existing storage instance.



_Appears in:_
- [CloneList](#clonelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Clone` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[CloneSpec](#clonespec)_ |  |  |  |
| `status` _[CloneStatus](#clonestatus)_ |  |  |  |


#### CloneInitParameters







_Appears in:_
- [CloneSpec](#clonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The default value is inherited from the source storage instance. A desired name is possible. The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sourceStorageId` _string_ | The ID of a storage instance which will be cloned.<br />ID of the storage instance that will be cloned. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |


#### CloneList



CloneList contains a list of Clones





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `CloneList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Clone](#clone) array_ |  |  |  |


#### CloneObservation







_Appears in:_
- [CloneStatus](#clonestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `locationUuid` _string_ | The location this resource is placed. The location of a resource is determined by it's project.<br />Identifies the data center this object belongs to. |  |  |
| `name` _string_ | The default value is inherited from the source storage instance. A desired name is possible. The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `parentUuid` _string_ |  |  |  |
| `sourceStorageId` _string_ | The ID of a storage instance which will be cloned.<br />ID of the storage instance that will be cloned. |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### CloneParameters







_Appears in:_
- [CloneSpec](#clonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The default value is inherited from the source storage instance. A desired name is possible. The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `sourceStorageId` _string_ | The ID of a storage instance which will be cloned.<br />ID of the storage instance that will be cloned. |  | Optional: \{\} <br /> |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  | Optional: \{\} <br /> |


#### CloneSpec



CloneSpec defines the desired state of Clone



_Appears in:_
- [Clone](#clone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[CloneParameters](#cloneparameters)_ |  |  |  |
| `initProvider` _[CloneInitParameters](#cloneinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### CloneStatus



CloneStatus defines the observed state of Clone.



_Appears in:_
- [Clone](#clone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[CloneObservation](#cloneobservation)_ |  |  |  |


#### StorageImport



StorageImport is the Schema for the StorageImports API. Make a clone of an existing storage instance.



_Appears in:_
- [StorageImportList](#storageimportlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageImport` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageImportSpec](#storageimportspec)_ |  |  |  |
| `status` _[StorageImportStatus](#storageimportstatus)_ |  |  |  |


#### StorageImportInitParameters







_Appears in:_
- [StorageImportSpec](#storageimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `storageBackupId` _string_ | ID of the storage backup that will be used to create a new storage from.<br />ID of the storage backup that will be used to create a new storage from. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |


#### StorageImportList



StorageImportList contains a list of StorageImports





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.m.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageImportList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[StorageImport](#storageimport) array_ |  |  |  |


#### StorageImportObservation







_Appears in:_
- [StorageImportStatus](#storageimportstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `locationUuid` _string_ | The location this resource is placed. The location of a resource is determined by it's project.<br />Identifies the data center this object belongs to. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `parentUuid` _string_ |  |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |
| `storageBackupId` _string_ | ID of the storage backup that will be used to create a new storage from.<br />ID of the storage backup that will be used to create a new storage from. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### StorageImportParameters







_Appears in:_
- [StorageImportSpec](#storageimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `storageBackupId` _string_ | ID of the storage backup that will be used to create a new storage from.<br />ID of the storage backup that will be used to create a new storage from. |  | Optional: \{\} <br /> |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  | Optional: \{\} <br /> |


#### StorageImportSpec



StorageImportSpec defines the desired state of StorageImport



_Appears in:_
- [StorageImport](#storageimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[LocalSecretReference](#localsecretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[ProviderConfigReference](#providerconfigreference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ kind:ClusterProviderConfig name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] | Enum: [Observe Create Update Delete LateInitialize *] <br /> |
| `forProvider` _[StorageImportParameters](#storageimportparameters)_ |  |  |  |
| `initProvider` _[StorageImportInitParameters](#storageimportinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageImportStatus



StorageImportStatus defines the observed state of StorageImport.



_Appears in:_
- [StorageImport](#storageimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageImportObservation](#storageimportobservation)_ |  |  |  |



## storage.gridscale.platformrelay.io/v1alpha1


### Resource Types
- [Clone](#clone)
- [CloneList](#clonelist)
- [StorageImport](#storageimport)
- [StorageImportList](#storageimportlist)



#### Clone



Clone is the Schema for the Clones API. Make a clone of an existing storage instance.



_Appears in:_
- [CloneList](#clonelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `Clone` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[CloneSpec](#clonespec)_ |  |  |  |
| `status` _[CloneStatus](#clonestatus)_ |  |  |  |


#### CloneInitParameters







_Appears in:_
- [CloneSpec](#clonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The default value is inherited from the source storage instance. A desired name is possible. The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `sourceStorageId` _string_ | The ID of a storage instance which will be cloned.<br />ID of the storage instance that will be cloned. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |


#### CloneList



CloneList contains a list of Clones





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `CloneList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Clone](#clone) array_ |  |  |  |


#### CloneObservation







_Appears in:_
- [CloneStatus](#clonestatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `locationUuid` _string_ | The location this resource is placed. The location of a resource is determined by it's project.<br />Identifies the data center this object belongs to. |  |  |
| `name` _string_ | The default value is inherited from the source storage instance. A desired name is possible. The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `parentUuid` _string_ |  |  |  |
| `sourceStorageId` _string_ | The ID of a storage instance which will be cloned.<br />ID of the storage instance that will be cloned. |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### CloneParameters







_Appears in:_
- [CloneSpec](#clonespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The default value is inherited from the source storage instance. A desired name is possible. The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `sourceStorageId` _string_ | The ID of a storage instance which will be cloned.<br />ID of the storage instance that will be cloned. |  | Optional: \{\} <br /> |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  | Optional: \{\} <br /> |


#### CloneSpec



CloneSpec defines the desired state of Clone



_Appears in:_
- [Clone](#clone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[CloneParameters](#cloneparameters)_ |  |  |  |
| `initProvider` _[CloneInitParameters](#cloneinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### CloneStatus



CloneStatus defines the observed state of Clone.



_Appears in:_
- [Clone](#clone)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[CloneObservation](#cloneobservation)_ |  |  |  |


#### StorageImport



StorageImport is the Schema for the StorageImports API. Make a clone of an existing storage instance.



_Appears in:_
- [StorageImportList](#storageimportlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageImport` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[StorageImportSpec](#storageimportspec)_ |  |  |  |
| `status` _[StorageImportStatus](#storageimportstatus)_ |  |  |  |


#### StorageImportInitParameters







_Appears in:_
- [StorageImportSpec](#storageimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `storageBackupId` _string_ | ID of the storage backup that will be used to create a new storage from.<br />ID of the storage backup that will be used to create a new storage from. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |


#### StorageImportList



StorageImportList contains a list of StorageImports





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `storage.gridscale.platformrelay.io/v1alpha1` | | |
| `kind` _string_ | `StorageImportList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[StorageImport](#storageimport) array_ |  |  |  |


#### StorageImportObservation







_Appears in:_
- [StorageImportStatus](#storageimportstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  |  |
| `changeTime` _string_ | Defines the date and time of the last object change.<br />Defines the date and time of the last object change. |  |  |
| `createTime` _string_ | The time the object was created.<br />Defines the date and time the object was initially created. |  |  |
| `currentPrice` _float_ | The price for the current period since the last bill. |  |  |
| `id` _string_ |  |  |  |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  |  |
| `lastUsedTemplate` _string_ | Indicates the UUID of the last used template on this storage (inherited from snapshots). |  |  |
| `licenseProductNo` _float_ | If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details).<br />If a template has been used that requires a license key (e.g. Windows Servers) this shows the product_no of the license (see the /prices endpoint for more details). |  |  |
| `locationCountry` _string_ | Two digit country code (ISO 3166-2) of the location where this object is placed.<br />Two digit country code (ISO 3166-2) of the location where this object is placed. |  |  |
| `locationIata` _string_ | Uses IATA airport code, which works as a location identifier.<br />Uses IATA airport code, which works as a location identifier. |  |  |
| `locationName` _string_ | The location name.<br />The human-readable name of the location. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `locationUuid` _string_ | The location this resource is placed. The location of a resource is determined by it's project.<br />Identifies the data center this object belongs to. |  |  |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  |  |
| `parentUuid` _string_ |  |  |  |
| `status` _string_ | status indicates the status of the object.<br />status indicates the status of the object. |  |  |
| `storageBackupId` _string_ | ID of the storage backup that will be used to create a new storage from.<br />ID of the storage backup that will be used to create a new storage from. |  |  |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  |  |
| `usageInMinutes` _float_ | The amount of minutes the IP address has been in use. |  |  |


#### StorageImportParameters







_Appears in:_
- [StorageImportSpec](#storageimportspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `capacity` _float_ | The default value is inherited from the source storage instance. A desired capacity is possible. Required (integer - minimum: 1 - maximum: 4096).<br />The capacity of a storage in GB. |  | Optional: \{\} <br /> |
| `labels` _string array_ | List of labels in the format [ "label1", "label2" ].<br />List of labels. |  | Optional: \{\} <br /> |
| `name` _string_ | The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters.<br />The human-readable name of the object. It supports the full UTF-8 character set, with a maximum of 64 characters. |  | Optional: \{\} <br /> |
| `storageBackupId` _string_ | ID of the storage backup that will be used to create a new storage from.<br />ID of the storage backup that will be used to create a new storage from. |  | Optional: \{\} <br /> |
| `storageType` _string_ | The default value is inherited from the source storage instance. A desired storage type is possible. (one of storage, storage_high, storage_insane).<br />(one of storage, storage_high, storage_insane) |  | Optional: \{\} <br /> |


#### StorageImportSpec



StorageImportSpec defines the desired state of StorageImport



_Appears in:_
- [StorageImport](#storageimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `writeConnectionSecretToRef` _[SecretReference](#secretreference)_ | WriteConnectionSecretToReference specifies the namespace and name of a<br />Secret to which any connection details for this managed resource should<br />be written. Connection details frequently include the endpoint, username,<br />and password required to connect to the managed resource. |  |  |
| `providerConfigRef` _[Reference](#reference)_ | ProviderConfigReference specifies how the provider that will be used to<br />create, observe, update, and delete this managed resource should be<br />configured. | \{ name:default \} |  |
| `managementPolicies` _[ManagementPolicies](#managementpolicies)_ | THIS IS A BETA FIELD. It is on by default but can be opted out<br />through a Crossplane feature flag.<br />ManagementPolicies specify the array of actions Crossplane is allowed to<br />take on the managed and external resources.<br />This field is planned to replace the DeletionPolicy field in a future<br />release. Currently, both could be set independently and non-default<br />values would be honored if the feature flag is enabled. If both are<br />custom, the DeletionPolicy field will be ignored.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223<br />and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md | [*] |  |
| `deletionPolicy` _[DeletionPolicy](#deletionpolicy)_ | DeletionPolicy specifies what will happen to the underlying external<br />when this managed resource is deleted - either "Delete" or "Orphan" the<br />external resource.<br />This field is planned to be deprecated in favor of the ManagementPolicies<br />field in a future release. Currently, both could be set independently and<br />non-default values would be honored if the feature flag is enabled.<br />See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223 | Delete | Enum: [Orphan Delete] <br /> |
| `forProvider` _[StorageImportParameters](#storageimportparameters)_ |  |  |  |
| `initProvider` _[StorageImportInitParameters](#storageimportinitparameters)_ | THIS IS A BETA FIELD. It will be honored<br />unless the Management Policies feature flag is disabled.<br />InitProvider holds the same fields as ForProvider, with the exception<br />of Identifier and other resource reference fields. The fields that are<br />in InitProvider are merged into ForProvider when the resource is created.<br />The same fields are also added to the terraform ignore_changes hook, to<br />avoid updating them after creation. This is useful for fields that are<br />required on creation, but we do not desire to update them after creation,<br />for example because of an external controller is managing them, like an<br />autoscaler. |  |  |


#### StorageImportStatus



StorageImportStatus defines the observed state of StorageImport.



_Appears in:_
- [StorageImport](#storageimport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `atProvider` _[StorageImportObservation](#storageimportobservation)_ |  |  |  |


