# ADR-0005: Cross-resource references — wire UUID-keyed edges only

**Theme:** 01 · Foundations · **Status:** Accepted · **Date:** 2026-07-16 · **Realises:** E7-S01
(audit finding ARCH-1, P1) + research doc `PROVIDER-DOCS-RESEARCH-2026-07-15.md` §Q3

## Context

Upstream terraform-provider-gridscale wires resources together through ~30 plain-string UUID fields
(`server.ipv4`, `loadbalancer.forwarding_rule[].certificate_uuid`, PaaS `network_uuid`, …). Without
`config.Reference{}` entries the generated CRDs expose only the raw string field, so users must
hardcode UUIDs, and Upjet's example generator emits literal HCL interpolations
(`${gridscale_ipv4.server.ip}`) that Crossplane treats as opaque strings — five generated example
manifests carried them (audit ARCH-1).

Two properties of this provider constrain the design:

- Every gridscale resource uses `IdentifierFromProvider` (see `config/external_name.go`), so the
  external name **is** the provider-assigned UUID for 30 of 32 resources. The default reference
  extractor returns the external name — which is exactly the value a `*_uuid` field needs.
- The two object-storage resources have **non-UUID** external names (`<s3_host>/<bucket_name>`,
  `<access_key>`), and some upstream edges interpolate non-UUID attributes (an IPv4 object's `.ip`)
  or target deprecated fields (`security_zone_uuid`).

## Decision

**Wire `config.Reference{TerraformName: …}` in [`config/references.go`](../../config/references.go)
for every UUID-keyed edge; use a custom extractor for the one IP-valued edge; skip deprecated and
non-UUID-keyed targets.** `TerraformName` (not the deprecated `Type` Go-path form) is used
throughout, so kind/version changes propagate automatically. The wiring is locked by
`config/reference_test.go` (TDD, ADR-0002) on both the cluster and namespaced providers.

### Edges wired (default extractor = external name = UUID)

| Resource | Field path | → Target |
| --- | --- | --- |
| `gridscale_server` | `ipv4` / `ipv6` / `isoimage` | `gridscale_ipv4` / `gridscale_ipv6` / `gridscale_isoimage` |
| `gridscale_server` | `network.object_uuid` | `gridscale_network` |
| `gridscale_server` | `network.firewall_template_uuid` | `gridscale_firewall` |
| `gridscale_server` | `storage.object_uuid` | `gridscale_storage` |
| `gridscale_loadbalancer` | `listen_ipv4_uuid` / `listen_ipv6_uuid` | `gridscale_ipv4` / `gridscale_ipv6` |
| `gridscale_loadbalancer` | `forwarding_rule.certificate_uuid` | `gridscale_ssl_certificate` |
| `gridscale_template` | `snapshot_uuid` | `gridscale_snapshot` |
| `gridscale_storage_clone` | `source_storage_id` | `gridscale_storage` |
| `gridscale_snapshot` / `gridscale_snapshotschedule` / `gridscale_backupschedule` | `storage_uuid` | `gridscale_storage` |
| PaaS family (`k8s`, `paas`, `postgresql`, `mysql`, `mysql8_0`, `mariadb`, `memcached`, `redis_cache`, `redis_store`, `sqlserver`, `filesystem`) | `network_uuid` | `gridscale_network` |

### Edge wired with a custom extractor

| Resource | Field path | → Target | Extractor |
| --- | --- | --- | --- |
| `gridscale_loadbalancer` | `backend_server.host` | `gridscale_ipv4` | `resource.ExtractParamPath("ip", true)` |

`host` holds the backend's IP address, not a UUID; the default extractor would resolve the IPv4
object's UUID — the wrong value. `ExtractParamPath("ip", true)` reads the referenced IPv4's observed
address from `status.atProvider.ip` instead. This was the last remaining `${…}` interpolation in
`examples-generated/` (loadbalancer + ssl-certificate examples).

### Edges deliberately skipped

- **`security_zone_uuid` (all PaaS resources) → `gridscale_paas_securityzone`** — upstream
  deprecates the field in favor of `network_uuid` (`paas.html.md:39`); wiring a Ref/Selector would
  actively steer users onto an API gridscale is removing. The plain string field remains usable for
  legacy setups. Guarded by a negative assertion in `config/reference_test.go`.
- **`object_storage_bucket.access_key`/`secret_key` → `object_storage_accesskey`** — both fields
  are `sensitive: true` in the schema, so upjet generates them as Kubernetes `*SecretRef` fields,
  not plain strings; `config.Reference{}` does not apply to secret-ref fields. Additionally the
  target's external name is the access-key string, not a UUID (story scope: UUID-keyed targets
  only). Users wire these through the accesskey's connection secret.
- **`ssl_certificate` itself has no ipv4-referencing field** — the interpolation seen in its
  generated example belonged to the dependent loadbalancer manifest (`backend_server.host`, wired
  above). Verified against `config/schema.json`: no skipped-but-existing field.

### Residual interpolations

None — `grep -rEn '\$\{' examples-generated/` returns nothing after regeneration.

## Consequences

- All 5 affected example manifests now render `*Ref`/`*Selector` stanzas; users compose resources by
  name/label instead of copy-pasting UUIDs.
- The generated CRDs gain sibling `*Ref`/`*Selector` fields for every wired path (and the PaaS CRDs
  **lose** the previously-generated `securityZoneUuidRef`/`securityZoneUuidSelector` fields — an
  intentional pre-1.0 API change; the CRD golden contracts in `config/testdata/crd-contract/` were
  refreshed accordingly).
- Per the codegen boundary (ADR-0003) the change is config-only; `apis/**`, `package/crds/**` and
  `examples-generated/**` were regenerated, never hand-edited.

## Counterpoints considered

- *"Wire `security_zone_uuid` too — the field still exists."* Rejected: deprecation hygiene
  (research Q5, story E4-S05) says examples and reference UX must steer to `network_uuid`; the
  deprecated field stays usable as a raw string.
- *"Use `Type` with explicit Go paths."* Rejected: `Type` is deprecated in upjet v2;
  `TerraformName` survives kind/version/group changes without hand-maintenance.
- *"Skip `backend_server.host` (non-UUID) for symmetry with the skipped object-storage edge."*
  Rejected: unlike the object-storage case, upjet ships a purpose-built extractor
  (`ExtractParamPath`) that resolves the observed `.ip` correctly, and leaving it out would leave
  the acceptance criterion (interpolation-free examples) unmet for 4 manifests.
