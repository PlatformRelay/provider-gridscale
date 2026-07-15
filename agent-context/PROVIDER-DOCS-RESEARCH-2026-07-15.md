# PROVIDER-DOCS-RESEARCH — 2026-07-15

**Subject:** gridscale Terraform provider + REST API vs. this Crossplane/Upjet `provider-gridscale` — contradictions, coverage gaps, improvements.
**Method:** local ground-truth (`config/`, `.work/gridscale/gridscale/website/`) cross-checked against upstream source (`github.com/gridscale/terraform-provider-gridscale`), GitHub releases/issues, and the provider config doc. Every claim cited to a URL or `path:line`.

> **Status: FOR OPERATOR REVIEW.** This complements the baseline audit ([`archive/audits/HEALTH-AUDIT-2026-07-15.md`](archive/audits/HEALTH-AUDIT-2026-07-15.md)) and the audit gap-stories ([`AUDIT-GAP-STORIES-2026-07-15.md`](AUDIT-GAP-STORIES-2026-07-15.md)). Proposed epics/stories at the bottom are *proposals* — not written into `docs/ROADMAP.md`/`BACKLOG.md`.

**Headline:** The provider is pinned to the current upstream release (v2.3.0, 2025-12-08) and generates all 32 resources — but as shipped it **cannot authenticate** (credentials never reach Terraform — corroborates CRED-1/2), exposes **zero datasources** (21 available), wires **zero cross-resource references** (~30 UUID edges), and ships **provider-config/secret examples with the wrong auth keys**.

---

## Q1. Coverage gaps (resources, datasources, post-v2.3.0)

### 1.1 — Zero datasources generated (21 available). **HIGH**
- **Claim:** The baked schema declares **21 datasources**; the provider generates **none**. Upjet generates resources only, and the include-list `.+` (`config/external_name.go:31`) matches resource schemas.
- **Evidence:** `config/schema.json` `data_source_schemas` = `backup_list, backupschedule, firewall, ipv4, ipv6, isoimage, k8s, loadbalancer, marketplace_application, network, object_storage_accesskey, paas, paas_securityzone, public_network, server, snapshot, snapshotschedule, sshkey, ssl_certificate, storage, template` (21). No datasource types under `apis/`. Upstream ships all (`gridscale/terraform-provider-gridscale/gridscale/datasource_*.go`).
- **Impact:** Users cannot reference pre-existing gridscale objects (existing network/storage/template, PaaS `service_template_uuid`) by lookup; every reference must be a hardcoded UUID or a managed resource.
- **Action:** Datasource-support epic (proposed **E8**), or document the omission.

### 1.2 — Datasource-only surfaces absent. **LOW**
- **Claim:** `gridscale_backup_list` and `gridscale_public_network` exist only as datasources (no resource twin) — entirely absent from the provider.
- **Evidence:** `config/schema.json` (present in `data_source_schemas`, absent from `resource_schemas`); upstream `datasource_gridscale_backup.go`, `datasource_gridscale_publicnetwork.go`.
- **Action:** Fold into E8; lookup-only by nature.

### 1.3 — Provider is current; no resources added after v2.3.0. **NONE (clarification)**
- **Claim:** v2.3.0 (2025-12-08) IS the latest upstream release; provider pinned to it. 2.2.4/2.2.5/2.3.0 added **no new resources/datasources** — only k8s pool taints/labels (2.3.0, PR #474) + dependency/CI bumps.
- **Evidence:** `gh release list --repo gridscale/terraform-provider-gridscale` → `v2.3.0 Latest 2025-12-08`. Note: the vendored `.work/.../CHANGELOG.md` tops out at 2.2.3 — a stale-checkout artifact, **not** a coverage gap.
- **Action:** None for coverage. Optional: refresh the vendored checkout so future audits diff cleanly.

---

## Q2. Auth model contradiction (corroborates CRED-1/CRED-2) — **CRITICAL / BLOCKER**

### 2.1 — Credentials never passed to Terraform
- **Claim:** `terraform.Setup.Configuration` assignment is **commented out** → no credentials (or any provider config) reach the Terraform provider; every reconcile fails on auth regardless of secret contents.
- **Evidence:** `internal/clients/gridscale.go:54-57` — the `ps.Configuration = map[string]any{ "username": ..., "password": ... }` block is inside `/* ... */`; `return ps, nil` follows immediately.
- **Impact:** Blocker — provider non-functional end to end. (= audit CRED-1.)

### 2.2 — Wrong auth keys even once re-enabled (`username`/`password` vs `uuid`/`token`)
- **Claim:** gridscale's provider config uses `uuid`, `token`, `api_url` (+ tunables) — **not** `username`/`password`. Both the code (2.1) and the shipped secret templates use `username`/`password`, absent from the gridscale provider schema.
- **Evidence:**
  - `config/schema.json` `provider_schemas.provider.block.attributes` = `uuid`, `token`, `api_url`, `http_headers`, `max_n_retries`, `request_delay_interval`. No `username`/`password`.
  - `.work/gridscale/gridscale/website/docs/index.html.markdown` — `provider "gridscale" { uuid = ...; token = ... }`; env fallbacks `GRIDSCALE_UUID`, `GRIDSCALE_TOKEN`, `GRIDSCALE_URL`; default `api_url = https://api.gridscale.io`.
  - `examples/cluster/providerconfig/secret.yaml.tmpl` + `examples/namespaced/providerconfig/secret.yaml.tmpl` both contain `{"username":"admin","password":"t0ps3cr3t11"}`.
- **Impact:** Even after 2.1 is fixed, auth still fails; operator-facing examples mislead. (= audit CRED-2.)
- **Action:** `ps.Configuration = {"uuid": creds["uuid"], "token": creds["token"], "api_url": creds["api_url"]}` (api_url optional) and rewrite both secret templates to `{"uuid":"...","token":"...","api_url":"https://api.gridscale.io"}`.

### 2.3 — Provider-config tunables unreachable. **MEDIUM (fold into 2.1/2.2)**
- **Claim:** `max_n_retries`, `request_delay_interval`, `http_headers`, `api_url` are valid provider args but unreachable while `Configuration` is empty. gridscale defaults `max_n_retries=1` (low for a rate-limited API).
- **Evidence:** `config/schema.json` provider attributes; `index.html.markdown` (defaults `request_delay_interval=1000`, `max_n_retries=1`).
- **Action:** Pass these optional keys through in the same fix as 2.1 (**belongs on L1-CRED**).

---

## Q3. Cross-resource references not wired (corroborates ARCH-1). **HIGH**

The provider wires **zero** `config.Reference`s. Upstream docs describe ~30 UUID fields referencing other resources; users must hardcode UUIDs today. Source docs: `.work/gridscale/gridscale/website/docs/r/`. Concrete edges:

| From resource | Field | → References | Doc |
|---|---|---|---|
| `server` | `storage[].object_uuid` | `storage` / `storage_clone` / `storage_import` | `server.html.md` |
| `server` | `network[].object_uuid` | `network` | `server.html.md` |
| `server` | `network[].firewall_template_uuid` | `firewall` | `server.html.md` |
| `server` | `ipv4` | `ipv4` | `server.html.md` |
| `server` | `ipv6` | `ipv6` | `server.html.md` |
| `loadbalancer` | `certificate_uuid` | `ssl_certificate` | `loadbalancer.html.md` |
| `backupschedule` | `storage_uuid` | `storage` | `backupschedule.html.md` |
| `snapshot` | `storage_uuid` | `storage` | `snapshot.html.md` |
| `snapshotschedule` | `storage_uuid` | `storage` | `snapshotschedule.html.md` |
| `k8s`/`paas`/`postgresql`/`mysql`/`mysql8_0`/`mariadb`/`memcached`/`redis_cache`/`redis_store`/`sqlserver`/`filesystem` | `network_uuid` | `network` | `k8s.html.md`, `paas.html.md`, … |
| (same PaaS set) | `security_zone_uuid` *(DEPRECATED)* | `paas_securityzone` | `paas.html.md:39` |
| `object_storage_bucket` | `access_key`/`secret_key` | `object_storage_accesskey` (`.access_key`/`.secret_key`) | `object_storage_bucket.html.md` |

- **Action:** "Wire cross-resource references" — see gap-story **E2-S06** (this research enriches it with the full edge list). Prefer `network_uuid` over deprecated `security_zone_uuid`.

---

## Q4. External-name / import correctness (recalibrates audit ARCH-3). **LOW (documentation, not code)**

- **Claim:** `IdentifierFromProvider` is **correct for all 32 resources** — but the rationale comment in `external_name.go` ("provider-assigned UUID") is inaccurate for the two object-storage resources, whose Terraform IDs are **not UUIDs**.
- **Evidence (upstream source):**
  - `object_storage_bucket`: `id := fmt.Sprintf("%s/%s", s3HostStr, bucketNameStr); d.SetId(id)` — composite `s3host/bucketname` (`resource_gridscale_bucket.go:249-250`); `Importer: ImportStatePassthrough` (36-37).
  - `object_storage_accesskey`: `d.SetId(response.AccessKey.AccessKey)` — the access-key string (`resource_gridscale_objectstorage.go:105`).
  - `storage_clone`/`storage_import`: `d.SetId(response.ObjectUUID)` — genuine UUIDs.
- **Impact:** No code fix — `IdentifierFromProvider` stores whatever TF assigns; both object-storage resources use `ImportStatePassthrough`, so import/observe work. Risk is operator confusion only. **This corrects the schema-only gap-story E2-S08, which is now a comment/docs fix.**
- **Action:** See gap-story **E2-S08** (downgraded). Fix comment at `config/external_name.go:7-9`; document import formats `<s3_host>/<bucket_name>` and `<access_key>`.

---

## Q5. Deprecations / pitfalls

- **Deprecated fields (late-init / drift risk). MEDIUM**
  - `security_zone_uuid` deprecated across **all** PaaS resources (paas, k8s, filesystem, postgresql, mysql, mysql8_0, mariadb, memcached, redis_cache, redis_store, sqlserver) — use `network_uuid` (`paas.html.md:39`, `k8s.html.md:49`).
  - `gridscale_memcached` — **entire resource deprecated** (`memcached.html.md:6,11`).
  - `gridscale_mysql` (5.7) — deprecated; migrate to `mysql8_0` (`mysql.html.md:11`).
  - `k8s.network_uuid` deprecated (`k8s.html.md:172`); `server.network[].ordering` deprecated (`server.html.md:134`).
- **BSL / Terraform 1.5.7 pin is deliberate. NONE (confirm, don't touch)**
  - `Makefile:7` `TERRAFORM_VERSION ?= 1.5.7`; `Makefile:112` errors if ≥1.6.0 ("BSL license"). Upstream provider itself is MPL-2.0. Keep the pin.
- **Rate-limit defaults.** gridscale defaults `max_n_retries=1`, `request_delay_interval=1000ms` — unreachable while Configuration is commented out (see 2.3).
- **Upstream open items worth a PR (operator interest). INFO** — see [[upstream-gridscale-bugs-open-pr]]
  - Docs bugs (easy PRs): #200 wrong `gridscale_paas_securityzone` sidebar name; #194 wrong PostgreSQL resource page name.
  - Feature gaps mirrored here: #187 "Add location resource/datasource" (no `location` resource — locations are UUID-only fields); #188 "Add storage backup location to backup schedule."
  - Evidence: `gh issue list --repo gridscale/terraform-provider-gridscale`.

---

## Q6. Improvement ideas

- **management-policies / initProvider late-init:** deprecated `security_zone_uuid` on PaaS is `Optional/Computed` and late-initialized — with references unwired this causes UUID drift. Wiring references (Q3) + documenting `initProvider` for computed PaaS fields reduces reconcile churn. MEDIUM.
- **Sensitive output handling:** redis/paas/memcached expose `username`/`password`/`kubeconfig` connection details (`GetConnectionDetailsMapping` in generated `zz_*_terraformed.go`). Verify these land in the connection secret; document.
- **PaaS ergonomics:** `service_template_uuid` is a required opaque UUID on every PaaS resource with no datasource to resolve it — the datasource epic (Q1) materially improves UX.
- **Testing:** no example manifests exercise auth; an e2e/uptest smoke against real `uuid`/`token` would have caught CRED-2 at build time.

---

## Sources / caveats
- Upstream **datasource docs are not vendored** (`.work/.../website/docs/d/` empty); datasource ground-truth from `config/schema.json` `data_source_schemas` + upstream `datasource_*.go` filenames (GitHub API).
- The registry docs page returned no parseable body via WebFetch; fell back to vendored `index.html.markdown`, `config/schema.json`, upstream source, and `gh` release/issue APIs.
- Vendored `CHANGELOG.md` stale (tops at 2.2.3); release facts from `gh release list/view`.

---

# Proposed backlog (fleshed) — research-derived, FOR REVIEW

Traceability: research item → disposition.

| Research item | Disposition |
| --- | --- |
| #1 Fix authentication (Q2) | **Already LANED** → OPERATOR-BOARD L1-CRED. Extend L1 to also pass `api_url` + tunables (2.3). |
| #3 Wire cross-resource references (Q3) | **Already storied** → gap-story E2-S06 (enriched with the edge table above). |
| #4 Correct object-storage external-name docs (Q4) | **Already storied** → gap-story E2-S08 (downgraded to comment/docs). |
| #2 Datasource support (Q1) | **NEW EPIC → E8** (below). |
| #5 Deprecation hygiene (Q5) | **NEW → E4-S05** (below). |
| #6 Provider tunables & rate limits (Q2.3) | **Fold into L1-CRED**; tracked as a REQ note (below). |
| #7 Upstream PR hunt (Q5) | **NEW → E6-S06** (below). |
| #8 e2e auth smoke test (Q6) | **NEW → E2-S09** (extends E2-S05). |

---

## E8 (NEW EPIC) — Datasource support

**Problem:** The provider exposes 0 of 21 gridscale datasources, so operators can't reference pre-existing objects (network, storage, template, PaaS `service_template_uuid`) by lookup — every cross-reference must be a managed resource or a hardcoded UUID. **Non-goal:** re-implementing resources as datasources; generating datasources gridscale doesn't offer.

### E8-S01 — Generate observe-only datasources for the 21 gridscale data sources (P2)
**As a** provider consumer **I want** gridscale datasources (network, storage, template, ipv4, ssl_certificate, paas, k8s, …) as observe-only Crossplane resources **so that** I can resolve existing objects' UUIDs without importing them as managed resources.
- Resolves research **Q1.1/Q1.2** (`config/schema.json` `data_source_schemas` = 21; none generated).
- **Acceptance:**
  - Given the Upjet config, when generation runs, then observe-only types are produced for the 21 `data_source_schemas` entries (incl. datasource-only `backup_list`, `public_network`).
  - Given a datasource CR referencing an existing network by name/UUID, when reconciled, then its status surfaces the object's fields (read-only; no create/update/delete).
  - Given a datasource for which gridscale requires a lookup key that's absent (edge), when reconciled, then it reports a clear "not found / ambiguous" condition.
- **Done when:**
  - [ ] REQ-E8-S01-01 — datasources generated. **Level:** M · **Test:** `apis/**` datasource types exist for the 21 · **Verify:** `make generate && make reviewable`
  - [ ] REQ-E8-S01-02 — one datasource e2e smoke. **Level:** E · **Test:** `examples/**/<datasource>.yaml` · **Verify:** uptest on `/test-examples`
  - [ ] ADR recording the datasource generation approach (Upjet observe-only vs custom).
- **Touches:** `config/`, `apis/**` (generated), `examples/**`. **Depends on:** L1-CRED (needs auth to e2e). **Sequence:** after references (E2-S06) so datasource+reference UX is coherent.

> Scope check: confirm current Upjet/`crossplane-runtime v2` datasource-generation support before committing; if unsupported, E8 becomes "document the omission + track upstream Upjet feature."

## E4-S05 (NEW) — Deprecation hygiene: steer off deprecated fields/resources (P2)
**As a** provider consumer **I want** examples and references to use non-deprecated gridscale fields **so that** I don't build on APIs gridscale is removing.
- Resolves research **Q5** (`security_zone_uuid`, `memcached`, `mysql` 5.7, `k8s.network_uuid`, `server.network[].ordering` deprecations).
- **Acceptance:**
  - Given generated/curated examples, when they set a PaaS network, then they use `network_uuid`, not deprecated `security_zone_uuid`.
  - Given provider docs/README resource matrix, when a resource/field is upstream-deprecated (`memcached`, `mysql` 5.7), then it is flagged deprecated with the recommended replacement (`mysql8_0`).
  - Given a new reference edge (E2-S06), when the field has a deprecated twin, then the reference targets the non-deprecated field.
- **Done when:**
  - [ ] REQ-E4-S05-01 — examples off deprecated fields. **Level:** M · **Test:** `examples/**` · **Verify:** `! grep -rn 'security_zone_uuid' examples/`
  - [ ] REQ-E4-S05-02 — deprecations flagged in docs. **Level:** D · **Test:** README/`docs/api/` matrix · **Verify:** doc lists deprecated resources + replacements
- **Touches:** `examples/**`, `README.md`/`docs/`. **Depends on:** coordinate with E2-S06 (references) and E3-S02 (README). **Collision:** README shared with E3-S02/E3-S04 — sequence.

## E6-S06 (NEW) — Upstream doc/feature PR hunt (P3)
**As a** maintainer **I want** the known upstream `terraform-provider-gridscale` doc bugs filed as PRs and feature gaps tracked **so that** our provider benefits from upstream fixes (standing operator interest — [[upstream-gridscale-bugs-open-pr]]).
- Resolves research **Q5** upstream items (#200, #194 doc bugs; #187, #188 feature gaps).
- **Acceptance:**
  - Given upstream issue #200 (wrong `paas_securityzone` sidebar name) and #194 (wrong PostgreSQL page name), when addressed, then a doc-fix PR is opened upstream (or a reason-not-to recorded).
  - Given feature gaps #187 (location resource/datasource) and #188 (backup location), when triaged, then each is recorded in `decisions.md`/BACKLOG as track-upstream-or-implement.
- **Done when:**
  - [ ] REQ-E6-S06-01 — upstream doc-fix PR(s) opened or declined-with-reason. **Level:** M · **Test:** PR link recorded in `decisions.md` · **Verify:** manual — PR URL present
  - [ ] REQ-E6-S06-02 — feature gaps triaged. **Level:** M · **Test:** BACKLOG/decisions entry · **Verify:** entries exist for #187/#188
- **Touches:** upstream repo (PRs) + `agent-context/decisions.md`. **Depends on:** none. **Note:** read-only w.r.t. this repo's code.

## E2-S09 (NEW) — e2e auth smoke test (P2, extends E2-S05)
**As a** maintainer **I want** an e2e smoke that authenticates with real `uuid`/`token` and creates one trivial resource **so that** a credential-wiring regression like CRED-1/2 fails CI instead of shipping.
- Resolves research **Q6** (no example exercises auth; would have caught CRED-2 at build time).
- **Acceptance:**
  - Given `UPTEST_CLOUD_CREDENTIALS` with real `uuid`/`token`, when the `/test-examples` e2e runs, then a minimal resource (e.g. `sshkey` or `network`) is created, becomes Ready, and is deleted.
  - Given the credentials are unwired or mis-keyed (regression), when the smoke runs, then it fails with an auth error (not a false pass).
  - Given no creds in a fork PR (edge), when e2e is gated, then the smoke is skipped, not failed.
- **Done when:**
  - [ ] REQ-E2-S09-01 — auth smoke wired into e2e. **Level:** E · **Test:** `examples/**/<minimal>.yaml` annotated for uptest · **Verify:** e2e run on `/test-examples`
  - [ ] REQ-E2-S09-02 — documents the creds contract (`uuid`/`token`/`api_url`). **Level:** D · **Test:** `examples/**/providerconfig/*.tmpl` + docs · **Verify:** creds contract documented
- **Touches:** `examples/**`, `.github/workflows/e2e.yaml`. **Depends on:** **L1-CRED** (creds must be wired first) and E2-S05 (e2e enablement). **Sequence:** after L1-CRED lands.
