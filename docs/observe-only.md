# Observe-Only Resources (Data Sources)

Crossplane supports an **observe-only** mode where a managed resource acts as a
read-only data source: the provider reads the remote object's state and surfaces
it in `.status.atProvider`, but never creates, updates, or deletes it.

## How to use observe-only resources

1. Set `spec.managementPolicies: ["Observe"]` so the reconciler skips Create /
   Update / Delete.
2. Set the `crossplane.io/external-name` annotation to the UUID of the existing
   gridscale resource you want to observe.
3. Leave `spec.forProvider` empty (`{}`) unless the resource type requires an
   input (e.g. `BackupList` requires `storageUUID`).

```yaml
apiVersion: gridscale.gridscale.platformrelay.io/v1alpha1
kind: Sshkey
metadata:
  name: my-sshkey-observer
  annotations:
    crossplane.io/external-name: "00000000-0000-0000-0000-000000000000"
spec:
  managementPolicies:
    - Observe
  providerConfigRef:
    name: default
  forProvider: {}
```

After the first successful reconciliation, the SSH key's attributes are
available under `.status.atProvider`:

```yaml
status:
  atProvider:
    name: my-key
    sshkey: "ssh-ed25519 AAAA..."
    createTime: "2024-01-01T00:00:00Z"
```

## Resources that support observe-only

All 19 managed resource types support observe-only mode by setting
`managementPolicies: ["Observe"]`.

### Cluster-scoped (`gridscale.gridscale.platformrelay.io/v1alpha1`)

| Terraform resource                  | Kind              |
|-------------------------------------|-------------------|
| `gridscale_backupschedule`          | `Backupschedule`  |
| `gridscale_firewall`                | `Firewall`        |
| `gridscale_ipv4`                    | `IPv4`            |
| `gridscale_ipv6`                    | `IPv6`            |
| `gridscale_isoimage`                | `Isoimage`        |
| `gridscale_k8s`                     | `K8S`             |
| `gridscale_loadbalancer`            | `Loadbalancer`    |
| `gridscale_network`                 | `Network`         |
| `gridscale_paas`                    | `Paas`            |
| `gridscale_server`                  | `Server`          |
| `gridscale_snapshot`                | `Snapshot`        |
| `gridscale_snapshotschedule`        | `Snapshotschedule`|
| `gridscale_sshkey`                  | `Sshkey`          |
| `gridscale_storage`                 | `Storage`         |
| `gridscale_template`                | `Template`        |

| Terraform resource                  | API group                                    | Kind              |
|-------------------------------------|----------------------------------------------|-------------------|
| `gridscale_marketplace_application` | `marketplace.gridscale.platformrelay.io`     | `Application`     |
| `gridscale_object_storage_accesskey`| `object.gridscale.platformrelay.io`          | `StorageAccesskey`|
| `gridscale_paas_securityzone`       | `paas.gridscale.platformrelay.io`            | `Securityzone`    |
| `gridscale_ssl_certificate`         | `ssl.gridscale.platformrelay.io`             | `Certificate`     |

For namespaced variants, use the `.m.` groups (e.g.
`gridscale.gridscale.m.platformrelay.io`) and add `namespace:` to the metadata.

## Custom observe-only resources

These two resources are implemented as native HTTP controllers (not upjet) and
are pure data sources with no Terraform equivalent.

### BackupList

Lists all storage backups for a given storage UUID.

**API group (cluster):** `storage.gridscale.platformrelay.io/v1alpha1`
**API group (namespaced):** `storage.gridscale.m.platformrelay.io/v1alpha1`
**Kind:** `BackupList`

**Required input:**

| Field                    | Description                                  |
|--------------------------|----------------------------------------------|
| `spec.forProvider.storageUUID` | UUID of the storage to list backups for |

**Example:**

```yaml
apiVersion: storage.gridscale.platformrelay.io/v1alpha1
kind: BackupList
metadata:
  name: my-storage-backups
spec:
  managementPolicies:
    - Observe
  providerConfigRef:
    name: default
  forProvider:
    storageUUID: "00000000-0000-0000-0000-000000000000"
```

**Observed attributes** (`.status.atProvider.storageBackups[]`):

| Field        | Type    | Description                     |
|--------------|---------|---------------------------------|
| `objectUuid` | string  | UUID of the backup              |
| `name`       | string  | Human-readable name             |
| `capacity`   | number  | Backup capacity in GB           |
| `createTime` | string  | ISO-8601 creation timestamp     |

### PublicNetwork

Returns the account's public network details. No inputs are required.

**API group (cluster):** `gridscale.gridscale.platformrelay.io/v1alpha1`
**API group (namespaced):** `gridscale.gridscale.m.platformrelay.io/v1alpha1`
**Kind:** `PublicNetwork`

**Example:**

```yaml
apiVersion: gridscale.gridscale.platformrelay.io/v1alpha1
kind: PublicNetwork
metadata:
  name: my-public-network
spec:
  managementPolicies:
    - Observe
  providerConfigRef:
    name: default
  forProvider: {}
```

**Observed attributes** (`.status.atProvider`):

| Field             | Type     | Description                      |
|-------------------|----------|----------------------------------|
| `name`            | string   | Network name                     |
| `status`          | string   | Network status                   |
| `networkType`     | string   | Always `"public"`                |
| `locationUUID`    | string   | Location UUID                    |
| `locationName`    | string   | Location display name            |
| `locationCountry` | string   | Country code                     |
| `locationIATA`    | string   | IATA airport code                |
| `l2Security`      | boolean  | Whether L2 security is enabled   |
| `deleteBlock`     | boolean  | Whether deletion is blocked      |
| `labels`          | string[] | Assigned labels                  |
| `createTime`      | string   | ISO-8601 creation timestamp      |
| `changeTime`      | string   | ISO-8601 last-change timestamp   |

## Note on attribute fidelity

For upjet-backed resources (the 19 listed above), nested block types from the
Terraform schema appear as nested structs in `.status.atProvider`. The field
names follow camelCase JSON conventions matching the Terraform attribute names.
For the two custom resources (`BackupList`, `PublicNetwork`), field names are
derived directly from the gridscale REST API response and documented above.
