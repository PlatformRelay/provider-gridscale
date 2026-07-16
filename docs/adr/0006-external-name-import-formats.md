# ADR-0006: External-name strategy and the two non-UUID import formats

**Theme:** 01 · Foundations · **Status:** Accepted · **Date:** 2026-07-16 · **Realises:** E7-S03
(audit finding ARCH-3, downgraded per `PROVIDER-DOCS-RESEARCH-2026-07-15.md` Q4)

## Context

Every generated resource in this provider uses Upjet's `IdentifierFromProvider` external-name
strategy, applied uniformly in [`config/external_name.go`](../../config/external_name.go) via the
`".+"` include-list. The original rationale comment claimed all gridscale objects are identified by
a "provider-assigned UUID". That is inaccurate for two resources, and the inaccuracy matters to
operators: the external-name is exactly the string an operator must supply when importing/adopting
an existing object (`crossplane.io/external-name` annotation), so a wrong mental model ("it's always
a UUID") leads to failed imports of object-storage resources.

The audit (ARCH-3) initially framed this as a possible *strategy* bug — the schema-only view shows
`object_storage_bucket` has no UUID `id`. Primary-source research against the upstream
`terraform-provider-gridscale` source settled it: the strategy is correct; only the documentation
was wrong.

## Decision

**Keep `IdentifierFromProvider` for all 32 resources — no `NameAsIdentifier`, no custom `GetIDFn` —
and document the two non-UUID external-name/import formats.**

`IdentifierFromProvider` stores whatever the upstream provider passes to Terraform's `SetId()` and
uses it verbatim as the external-name. What that string is per resource:

| Resources | External-name / import format | Upstream evidence |
| --- | --- | --- |
| 30 of 32 (server, storage, network, ipv4, k8s, paas, …) | provider-assigned UUID, e.g. `d.SetId(response.ObjectUUID)` (`storage_clone`, `storage_import`) | upstream resource sources |
| `object_storage_bucket` | composite **`<s3_host>/<bucket_name>`** — `id := fmt.Sprintf("%s/%s", s3HostStr, bucketNameStr); d.SetId(id)` | `resource_gridscale_bucket.go:249-250`; importer is `ImportStatePassthrough` (lines 36-37) |
| `object_storage_accesskey` | the **`<access_key>`** string — `d.SetId(response.AccessKey.AccessKey)` | `resource_gridscale_objectstorage.go:105` |

Why no override is needed:

- Upjet's `IdentifierFromProvider` makes no UUID assumption — it round-trips the Terraform ID as an
  opaque string. Composite and key-string IDs work unchanged.
- Both object-storage resources use `ImportStatePassthrough` upstream, so setting the
  `crossplane.io/external-name` annotation to the formats above makes import/observe work as-is.
- `NameAsIdentifier` would be *wrong* here: none of the 32 resources are identified by a
  user-chosen `name` argument at the Terraform level, and a custom `GetIDFn` would only re-derive
  what `SetId` already stored.

Operator guidance: to import an existing bucket, set the external-name annotation to
`<s3_host>/<bucket_name>` (e.g. `gos3.io/my-bucket`); to import an existing access key, set it to
the access-key string. All other resources take the gridscale object UUID.

## Consequences

- The rationale comment in `config/external_name.go` now names the two non-UUID cases and points
  here; the misleading "every resource is a UUID" claim is gone.
- No behavioural change: the uniform `IdentifierFromProvider` + `".+"` include-list stays, so the
  generated tree, CRDs, and controllers are untouched (regression guard: existing config tests and
  `make check-diff` stay green).
- Future resource additions inherit the correct default automatically; the only maintenance duty is
  to check upstream's `SetId` when a new resource's ID is not a UUID and extend the table above.

## Counterpoints considered

- *"Switch object-storage to `NameAsIdentifier` so operators can pick the external-name."* Rejected:
  the Terraform ID is not a name argument; diverging from `SetId`'s value would break Observe (the
  ID Upjet feeds back to Terraform must be the one upstream understands via
  `ImportStatePassthrough`).
- *"Add a custom `GetIDFn` to normalise the bucket ID to just `<bucket_name>`."* Rejected: upstream
  parses `<s3_host>/<bucket_name>` on read/delete; stripping the host would hand Terraform an ID it
  cannot resolve.
- *"Leave the comment as-is; it's only a comment."* Rejected: the external-name string is
  operator-facing API surface (the import handle). ARCH-3 exists because the comment actively
  teaches the wrong import format for two resources.
