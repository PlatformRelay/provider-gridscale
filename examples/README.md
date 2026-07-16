# Curated examples

Curated, minimal, apply-able example manifests for the gridscale Crossplane
provider — one per major API group. Unlike the auto-generated manifests in
[`examples-generated/`](../examples-generated) (which carry unresolved
`${...}` HCL interpolations), these are hand-written, cleaned up, and safe to
`kubectl apply` in principle. They also double as
[uptest](https://github.com/crossplane/uptest) fixtures — examples-as-tests.

Each manifest maps directly to a CRD under
[`package/crds/`](../package/crds) and to the CRD reference documentation:
the `apiVersion` / `kind` identify the group and resource.

## Layout

These are **cluster-scoped** resources, grouped by API group:

| File | Group | Kind |
| --- | --- | --- |
| `gridscale/server.yaml` | `gridscale.gridscale.platformrelay.io` | `Server` |
| `gridscale/network.yaml` | `gridscale.gridscale.platformrelay.io` | `Network` |
| `gridscale/storage.yaml` | `gridscale.gridscale.platformrelay.io` | `Storage` |
| `gridscale/sshkey.yaml` | `gridscale.gridscale.platformrelay.io` | `Sshkey` |
| `gridscale/ipv4.yaml` | `gridscale.gridscale.platformrelay.io` | `IPv4` |
| `redis/cache.yaml` | `redis.gridscale.platformrelay.io` | `Cache` |
| `paas/securityzone.yaml` | `paas.gridscale.platformrelay.io` | `Securityzone` |
| `mysql8/mysql8.yaml` | `mysql8.gridscale.platformrelay.io` | `MySQL8` |
| `storage/storageimport.yaml` | `storage.gridscale.platformrelay.io` | `StorageImport` |
| `ssl/certificate.yaml` | `ssl.gridscale.platformrelay.io` | `Certificate` |
| `object/storagebucket.yaml` | `object.gridscale.platformrelay.io` | `StorageBucket` |

## Usage

1. Install the provider and create a `ProviderConfig` named `default`. See the
   provider-config examples:
   - Cluster-scoped: [`cluster/providerconfig/`](cluster/providerconfig)
   - Namespaced: [`namespaced/providerconfig/`](namespaced/providerconfig)

   Every manifest here references it via:

   ```yaml
   spec:
     providerConfigRef:
       name: default
   ```

2. Apply an example:

   ```sh
   kubectl apply -f examples/gridscale/network.yaml
   ```

3. Watch it reconcile:

   ```sh
   kubectl get network.gridscale.gridscale.platformrelay.io example-network
   ```

Replace placeholder values (SSH public key, `storageBackupId`, etc.) with real
values before applying against a live gridscale account.

## Notes on coverage

- The five gridscale primitives (`Server`, `Network`, `Storage`, `Sshkey`,
  `IPv4`) are standalone and independently apply-able.
- `redis/Cache`, `paas/Securityzone` and `mysql8/MySQL8` are standalone: their
  `networkUuid` field is optional (gridscale provisions a dedicated security
  zone when omitted). Prefer `networkUuid` over the deprecated
  `securityZoneUuid` if you do set a network.
- `storage/StorageImport` needs a `storageBackupId` — a cross-resource
  dependency on an existing storage backup. A placeholder UUID is used; replace
  it before applying.
- `ssl/Certificate` and `object/StorageBucket` reference external material
  (certificate/key, S3 access/secret keys) through `secretRef`s pointing at a
  Kubernetes Secret you create first — the same pattern as the provider-config
  `secret.yaml.tmpl`. Create the referenced Secret before applying.

The following groups/resources are intentionally **not** covered here because a
minimal manifest cannot be independently apply-able:

- `marketplace/Application` — requires an object-storage image path
  (`objectStoragePath`) plus publishing metadata; not a lightweight fixture.
- `storage/Clone` — requires a `sourceStorageId` reference to an existing
  storage.
- `object/StorageAccesskey` — a credential-derivation resource with no minimal
  standalone shape.

For these, see the shapes in [`examples-generated/`](../examples-generated).
