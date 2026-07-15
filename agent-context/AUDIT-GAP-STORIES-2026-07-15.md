# AUDIT-GAP-STORIES — 2026-07-15

Untracked audit findings from [`archive/audits/HEALTH-AUDIT-2026-07-15.md`](archive/audits/HEALTH-AUDIT-2026-07-15.md) translated into INVEST stories. Every finding is grounded against `origin/main` code; stories match the `E<n>-S<nn>` / `REQ-E<n>-S<nn>-<mm>` scheme and the `openspec/config.yaml` U/E/M/D taxonomy (one **Test:** artifact + one **Verify:** command per REQ).

> **Status: FOR OPERATOR REVIEW — do not auto-merge.** Proposed story IDs below are *proposals*; they are NOT yet written into `docs/ROADMAP.md` / `agent-context/BACKLOG.md` (that is the E1 design lane's territory). A concurrent `/agent-loop-auto` session already laned the top-4 findings (L1-CRED, L2-TEST, L3-README — see `coordination/OPERATOR-BOARD.md`); this doc covers only the untracked remainder.

## Traceability — every finding → tracking status

| Finding | Sev | Tracking status | Where |
| --- | --- | --- | --- |
| CRED-1 | P0 | **LANED** | OPERATOR-BOARD L1-CRED |
| CRED-2 | P0 | **LANED** | OPERATOR-BOARD L1-CRED |
| DOC-1 | P1 | **LANED** | OPERATOR-BOARD L1-CRED |
| TEST-1 | P1 | **LANED** | OPERATOR-BOARD L2-TEST |
| DOC-3 | P2 | **LANED** | OPERATOR-BOARD L3-README |
| DOC-2 | P2 | **ROADMAP-E4** | ≈ E4-S02 (curated example per group) |
| TEST-2 | P2 | **ROADMAP-E2** | ≈ E2-S05 (enable e2e on `/test-examples` + nightly) |
| DOC-4 | P3 | **ROADMAP-E6** (surfaced) | ≈ E6-S03; operator decision on maintainer (OPERATOR-BOARD deferred) |
| DIR-3 | P3 | **WITHDRAWN-FIXED** | commit `797adad` |
| DIR-1 | — | **WITHDRAWN-FIXED** | docs exist on main |
| **ARCH-1** | P1 | **NEW → E2-S06** | cross-resource `config.Reference{}` wiring |
| **ARCH-2** | P2 | **NEW → E2-S07** | `provider-metadata.yaml` malformed keys (codegen input) |
| **SEC-1** | P2 | **NEW → E5-S07** | golangci global gosec suppression |
| **DIR-2 (changelog + Go drift)** | P2 | **NEW → E5-S08** | scaffold-identity tidy |
| **DIR-2 (goimports prefix)** | P2 | **NEW → E5-S09** | repo-wide import reformat — **run alone** |
| **ARCH-3** | P3 | **NEW → E2-S08** (downgraded) | external-name **rationale comment** fix for object-storage — code strategy is correct |
| **TEST-3** | P3 | **NEW → E5-S10** | dormant `schema-version-diff` gate |

**Epic placement note:** ARCH-1/2/3 are *config-correctness* work (they change `config/*.go` and codegen inputs, not test files), so they do **not** belong on E2's `config/*_test.go` path-set — but they are the hand-authored `config/` surface E2's ADR-0002 exists to protect, and each ships with a test at the right tier. Rather than open a new epic for three stories, they slot into **E2** as new stories with an **explicit config-source path-set** (`config/*.go`, `config/provider-metadata.yaml`, `examples-generated/**`) distinct from E2's existing test path-set. This is flagged as a **lane collision** below (they share `config/` with L2-TEST). If the operator prefers isolation, promote them to a new epic **E7 — Config correctness**; the story bodies are epic-agnostic.

---

### E2-S06 — Wire cross-resource references in `config/` (P1)
**As a** provider consumer **I want** related resources (server↔ipv4/network/storage, loadbalancer↔ssl-cert/ipv4/ipv6, ssl-certificate backend↔ipv4, storage-clone↔storage, template↔snapshot) to resolve by Crossplane reference **so that** generated examples apply cleanly and can pass e2e instead of carrying literal `${...}` HCL that Crossplane treats as opaque strings.

Resolves **ARCH-1** (`config/` — no `Reference{}` present, grep-confirmed). Concrete symptom confirmed in 5 resources across both scopes: `examples-generated/{cluster,namespaced}/ssl/v1alpha1/certificate.yaml`, `.../gridscale/v1alpha1/loadbalancer.yaml`, `.../gridscale/v1alpha1/server.yaml`, `.../gridscale/v1alpha1/template.yaml`, `.../storage/v1alpha1/clone.yaml` — all carry unresolved interpolations e.g. `${gridscale_ipv4.server.ip}`, `${gridscale_ssl_certificate.ssl-certificate-john.id}`, `${gridscale_snapshot.foo.id}`, `${gridscale_storage.storage-clone-john.id}`.

**Acceptance criteria**
- Given a `config.Resource` for `gridscale_server`, when the provider is configured, then its `ipv4` / `network[].object_uuid` / `storage[].object_uuid` fields carry `config.Reference{}` to the respective Ipv4/Network/Storage kinds (with `Selector`/`Ref` sibling fields generated).
- Given `gridscale_loadbalancer` and `gridscale_ssl_certificate`, when configured, then the ssl-certificate / ipv4 / ipv6 edges are wired as references.
- Given the example generator runs, when it emits the five affected manifests, then no field value contains a literal `${` interpolation (they become `...Ref`/`...Selector` or resolved values).
- Given a reference is added for a field that no longer exists in `schema.json` (edge/error), when `make generate` runs, then generation fails loudly (unknown field path) rather than silently no-op'ing.

**Done when**
- [ ] REQ-E2-S06-01 — references configured for the server/loadbalancer/ssl/clone/template edges. **Level:** M · **Test:** `config/reference_test.go` (unit-asserts each configured `Resource.References[path].Type` is non-empty for the target fields) · **Verify:** `go test ./config/... -run TestReferences`
- [ ] REQ-E2-S06-02 — regenerated examples are interpolation-free. **Level:** M · **Test:** `examples-generated/**` (regenerated) · **Verify:** `! grep -rEn '\$\{' examples-generated/`
- [ ] Docs updated: `docs/adr/` note or `config/` comment recording which edges are referenced and why (UUID-keyed targets only).
- [ ] Gates green: `make generate` clean, `make reviewable`, `make check-diff`, `make test.spec`.

**Touches:** `config/gridscale.go` (or new `config/references.go`), `config/provider.go`, `examples-generated/**`
**Depends on:** none (independent of L1-CRED; unblocks E2-S05 e2e)
**Not in scope:** external-name strategy changes (that is E2-S08 / ARCH-3); wiring refs for name-keyed targets whose `id` is not a stable UUID (defer); the e2e-lab curation in E2-S04.

---

### E2-S07 — Fix malformed `provider-metadata.yaml` resource keys (codegen input) (P2)
**As a** provider maintainer **I want** every top-level key under `resources:` in `provider-metadata.yaml` to be a real `gridscale_*` resource name **so that** upjet's doc/example scraper matches metadata to resources instead of dropping the eight title-keyed entries — which is why those exact resources have no generated examples.

Resolves **ARCH-2** (`config/provider-metadata.yaml:3` `ISO Image:` and 7 more title-keyed entries; corrupted field text at `:10` `"...storage/ISO Image/ISO Image/snapshot..."`). Confirmed title-keyed (non-`gridscale_`) entries: `ISO Image`, `imported marketplace application`, `marketplace application`, `object storage access key`, `object storage bucket`, `storage backup schedule`, `storage snapshot`, `storage snapshot schedule`.

**Root-cause hypothesis:** `provider-metadata.yaml` is scraper output (`apis/generate.go` `//go:generate ... scraper -o ../config/provider-metadata.yaml`), `//go:embed`ed at `config/provider.go:18`. The eight entries are keyed by human title because the corresponding upstream `terraform-provider-gridscale` doc pages lack the frontmatter `page_title`/resource-name mapping the scraper keys on. A hand-edit of the yaml is clobbered on the next scrape — so the durable fix is either a scraper post-process step in `apis/generate.go` (rename title keys → `gridscale_*` deterministically) **or** an upstream doc PR adding the missing frontmatter (per MEMORY: "upstream gridscale bug → open PR").

**Non-goals:** hand-editing the embedded yaml as the sole fix (would regress on regen); adding the missing *examples* themselves (that is DOC-2 ≈ E4-S02); re-scraping unrelated fields.

**Acceptance criteria**
- Given the regenerated `provider-metadata.yaml`, when its top-level `resources:` keys are enumerated, then every key equals a `gridscale_*` name present in `schema.json` `resource_schemas` (32 entries) — no key contains a space or capitalized title text.
- Given the fix is a post-process (not a raw edit), when `make generate` / the scraper re-runs, then the keys stay `gridscale_*` (no regression).
- Given an upstream doc page still lacks frontmatter (edge/error), when the post-process runs, then it maps the title deterministically or fails with a named, unmapped-title error rather than emitting a title key silently.

**Done when**
- [ ] REQ-E2-S07-01 — all metadata keys are canonical resource names. **Level:** M · **Test:** `config/metadata_test.go` (parses embedded `providerMetadata`, asserts `keys ⊆ schema.json resource_schemas` and none matches `[A-Z ]`) · **Verify:** `go test ./config/... -run TestMetadataKeys`
- [ ] REQ-E2-S07-02 — regeneration is idempotent on keys. **Level:** M · **Test:** post-process step in `apis/generate.go` (or upstream-PR link recorded in ADR) · **Verify:** `make generate && git diff --exit-code config/provider-metadata.yaml`
- [ ] Docs updated: ADR / `apis/generate.go` comment recording the title→name mapping source; if upstream PR route, link it.
- [ ] Gates green: `make reviewable`, `make check-diff`, `make test.spec`.

**Touches:** `config/provider-metadata.yaml`, `apis/generate.go` (post-process), `config/metadata_test.go`
**Depends on:** none
**Not in scope:** the 8 missing example manifests (DOC-2 / E4-S02).

---

### E2-S08 — Correct the object-storage external-name rationale comment + document import formats (P3)
**As a** provider operator **I want** the external-name rationale to accurately describe how `object_storage_bucket` and `object_storage_accesskey` are identified **so that** I import/adopt them with the right external-name string instead of assuming a UUID that gridscale never assigns for these two resources.

Resolves **ARCH-3** (`config/external_name.go:7-9` misleading "provider-assigned UUID" rationale comment).

> **⚠ Reconciled with primary-source research (`PROVIDER-DOCS-RESEARCH-2026-07-15.md`, Q4) — SCOPE DOWNGRADED.** The original "verify/correct the strategy" framing (schema-only view: bucket has no UUID `id`) was **over-scoped**. Reading the actual upstream provider source shows `IdentifierFromProvider` is **correct for all 32 resources, including object-storage** — Upjet stores whatever Terraform's `SetId` assigns and uses it verbatim as the external-name, and both object-storage resources use `ImportStatePassthrough`, so import/observe already work. Evidence: `object_storage_bucket` → `d.SetId("<s3_host>/<bucket_name>")` (`resource_gridscale_bucket.go:249-250`, `ImportStatePassthrough` line 36-37); `object_storage_accesskey` → `d.SetId(<access_key>)` (`resource_gridscale_objectstorage.go:105`); `storage_clone`/`storage_import` → `d.SetId(<UUID>)`. **No `NameAsIdentifier`/custom `GetIDFn` is needed.** This is now a documentation-correctness story, not a code fix.

**Acceptance criteria**
- Given `config/external_name.go`, when its rationale comment is read, then it no longer claims every resource is identified by a "provider-assigned UUID" — it notes that `object_storage_bucket` uses `<s3_host>/<bucket_name>` and `object_storage_accesskey` uses the access-key string.
- Given an operator importing an existing bucket, when they consult provider docs, then the external-name format `<s3_host>/<bucket_name>` is documented (and `<access_key>` for access keys).
- Given the 30 UUID-keyed resources, when external-name resolves, then behaviour is unchanged (no code change — regression guard only).

**Done when**
- [ ] REQ-E2-S08-01 — rationale comment corrected. **Level:** M · **Test:** `config/external_name.go:7-9` (comment) · **Verify:** `! grep -q 'provider-assigned UUID' config/external_name.go` (or equivalent — comment names the two non-UUID cases)
- [ ] REQ-E2-S08-02 — import formats documented. **Level:** D · **Test:** `docs/adr/` (external-name ADR) or resource docs · **Verify:** doc mentions `<s3_host>/<bucket_name>` and `<access_key>`
- [ ] Gates green: `make reviewable`.

**Touches:** `config/external_name.go` (comment only), `docs/adr/` or `docs/api/`
**Depends on:** none. **No longer touches `external_name.go` logic** — comment-only, so the E2-S01 collision risk is minimal (still, coordinate if E2-S01 rewrites the file header).
**Not in scope:** any `NameAsIdentifier`/`GetIDFn` change (confirmed unnecessary); the include-list `.+`.

---

### E5-S07 — Scope gosec suppressions to tests only (P2)
**As a** security reviewer **I want** gosec `G101` (hardcoded credentials) and `G104` (unhandled errors) enforced on hand-authored code **so that** a real credential leak or swallowed error in `config/`, `internal/clients/`, or `cmd/` is caught by CI instead of being globally excluded.

Resolves **SEC-1** (`.golangci.yml:102-107` — `G101` and `G104` excluded for **all** files, generated and hand-authored alike).

**Acceptance criteria**
- Given `.golangci.yml`, when the exclusion rules are read, then `G101`/`G104` are suppressed only for generated paths (`zz_.+\.go$`) and `_test.go`, not globally.
- Given a hand-authored file with a hardcoded credential literal (edge/error), when `golangci-lint run` executes, then `G101` fires and the lint job fails.
- Given the current clean hand-authored surface, when the narrowed config runs, then lint stays green (no pre-existing violations unmasked, or any that surface are fixed in-scope).

**Done when**
- [ ] REQ-E5-S07-01 — suppressions narrowed to generated/test paths. **Level:** M · **Test:** `.golangci.yml` (the scoped exclusion rules) · **Verify:** `golangci-lint run ./config/... ./internal/... ./cmd/...`
- [ ] Docs updated: comment in `.golangci.yml` stating why the two rules are test/generated-only.
- [ ] Gates green: `make lint` / `make reviewable`.

**Touches:** `.golangci.yml` (lines ~102-107)
**Depends on:** none — but **collides on `.golangci.yml` with E5-S09** (goimports prefix, line 126). Sequence E5-S07 and E5-S09 (do not run concurrently on the same file); E5-S07 first is fine since it does not reformat imports.
**Not in scope:** the goimports local-prefix (E5-S09); adding new linters (E5-S03).

---

### E5-S08 — Scaffold-identity tidy: changelog label + Go-version drift (P2)
**As a** maintainer **so that** the provider reports its own identity and builds on one Go toolchain, **I want** the AWS-template changelog label and the three-way Go-version drift fixed.

Resolves **DIR-2** (partial): `cmd/provider/main.go:211` emits `provider-upjet-aws:%s` in the change-logger provider version; Go version drifts across `Makefile:47` (`1.24`), `.github/workflows/ci.yml:13` (`1.25`), `.github/workflows/e2e.yaml:12` (`1.24`).

**Acceptance criteria**
- Given the change-logger, when it reports a provider version, then the label reads `provider-gridscale:%s` (not `provider-upjet-aws`).
- Given the three build configs, when their Go versions are read, then all three agree on one pinned version.
- Given a future drift (edge/error), when CI runs, then a meta check flags mismatched Go versions rather than silently building on divergent toolchains.

**Done when**
- [ ] REQ-E5-S08-01 — changelog label is gridscale. **Level:** M · **Test:** `cmd/provider/main.go:211` · **Verify:** `grep -q 'provider-gridscale:%s' cmd/provider/main.go && ! grep -q 'provider-upjet-aws' cmd/provider/main.go`
- [ ] REQ-E5-S08-02 — single Go version. **Level:** M · **Test:** `Makefile`, `.github/workflows/ci.yml`, `.github/workflows/e2e.yaml` · **Verify:** a one-line check asserting the Go version string is identical in all three (e.g. `hack/check-go-version.sh` or inline `grep`)
- [ ] Gates green: `make build`, `make test.spec`.

**Touches:** `cmd/provider/main.go`, `Makefile`, `.github/workflows/ci.yml`, `.github/workflows/e2e.yaml`
**Depends on:** none. Touches `.github/workflows/**` (E5 path-set) + `cmd/` + `Makefile`. Per BACKLOG "prefer adding over editing `ci.yml`" — here a single-line env bump is unavoidable; keep the diff minimal.
**Not in scope:** goimports prefix (E5-S09); bumping the pinned Go version to something new (pick the value already in majority use); the BSL-pinned `TERRAFORM_VERSION`.

---

### E5-S09 — Fix goimports local-prefix (repo-wide reformat) (P2)
**As a** maintainer **I want** the goimports local-prefix to be `github.com/PlatformRelay/provider-gridscale` **so that** import grouping reflects this module and `make lint` / `make reviewable` order imports correctly.

Resolves **DIR-2** (partial): `.golangci.yml:126` still `github.com/crossplane/upjet-provider-template`. BACKLOG note ("golangci local-prefix is stale … belongs in E5-S03 or a tidy commit") is subsumed here.

**Acceptance criteria**
- Given `.golangci.yml`, when the goimports `local-prefixes` is read, then it is `github.com/PlatformRelay/provider-gridscale`.
- Given the corrected prefix, when `goimports -local` runs repo-wide, then hand-authored imports are regrouped and `make lint` is green.
- Given a stray upstream-template import path anywhere in config (edge), when lint runs, then it is regrouped, not left in the third-party block.

**Done when**
- [ ] REQ-E5-S09-01 — prefix corrected + imports reformatted. **Level:** M · **Test:** `.golangci.yml:126` · **Verify:** `grep -q 'github.com/PlatformRelay/provider-gridscale' .golangci.yml && make lint`
- [ ] Gates green: `make reviewable`.

**Touches:** `.golangci.yml` (line 126) **and every hand-authored `.go` file's import block** (repo-wide reformat).
**Depends on:** none — **RUN ALONE.** This is the flagged repo-wide reformat: it rewrites imports across the whole tree and will collide with **every** in-flight Go lane (L1-CRED, L2-TEST, E2-S06/S07/S08). Land it in its own batch when no other Go lane is open. Also collides on `.golangci.yml` with **E5-S07** — sequence the two.
**Not in scope:** the gosec suppressions (E5-S07); Go-version / changelog fixes (E5-S08).

---

### E5-S10 — Activate the dormant `schema-version-diff` gate (P3)
**As a** maintainer **I want** the native-schema-drift check to actually run on push/PR **so that** an upstream `terraform-provider-gridscale` schema change that breaks generated CRDs is caught, instead of the gate silently never executing.

Resolves **TEST-3** (`.github/workflows/ci.yml:53-56` — the `report-breaking-changes` job's "Report native schema version changes" step is guarded by `if: ${{ inputs.upjet-based-provider }}`, a `workflow_call` input never provided by the `push` / `pull_request` triggers at `ci.yml:4-8`, so the step is dead).

**Root-cause hypothesis:** the step was copied from a reusable-workflow (`workflow_call`) template where `inputs.upjet-based-provider` is set by the caller; on this repo's direct push/PR triggers that input is always empty → the `if` is always false.

**Acceptance criteria**
- Given a PR that modifies the generated schema, when CI runs, then the `schema-version-diff` step executes (guard replaced with a real condition — e.g. gate on modified schema/CRD files, matching the sibling `crddiff` step at `ci.yml:47-52`).
- Given a PR with no schema change (edge), when CI runs, then the step is skipped cleanly (no spurious failure).
- Given `make schema-version-diff` is unavailable or errors (edge/error), when the step runs, then it fails visibly with the tool's message, not a silent pass.

**Done when**
- [ ] REQ-E5-S10-01 — step guard fires on real triggers. **Level:** M · **Test:** `.github/workflows/ci.yml:53-56` · **Verify:** `! grep -q 'inputs.upjet-based-provider' .github/workflows/ci.yml` and the step's `if` references a push/PR-available condition (e.g. `steps.modified-crds.outputs.any_changed`)
- [ ] Docs updated: comment in `ci.yml` explaining when `schema-version-diff` runs.
- [ ] Gates green: workflow lints (`actionlint`) / CI dry-run.

**Touches:** `.github/workflows/ci.yml` (lines 53-56)
**Depends on:** none. E5 `.github/workflows/**` path-set; no collision with E5-S07/S08/S09.
**Not in scope:** making the check PR-blocking vs advisory (operator policy); the `crddiff` step (already active).

---

## Lane / collision map (against BACKLOG disjoint-path graph)

| Story | Path-set | Lane | Collision flag |
| --- | --- | --- | --- |
| E2-S06 (ARCH-1) | `config/*.go`, `examples-generated/**` | config-source (new) | Shares `config/` with L2-TEST/E2-S01 — sequence, don't run concurrently |
| E2-S07 (ARCH-2) | `config/provider-metadata.yaml`, `apis/generate.go`, `config/metadata_test.go` | config-source | Low — distinct files |
| E2-S08 (ARCH-3) | `config/external_name.go` (comment only), `docs/` | docs | Low — comment/docs only after research downgrade; light coordinate with E2-S01 |
| E5-S07 (SEC-1) | `.golangci.yml` (102-107) | E5 | **Collides with E5-S09** on `.golangci.yml` — sequence |
| E5-S08 (DIR-2) | `cmd/`, `Makefile`, `ci.yml`, `e2e.yaml` | E5 | Low — env/label edits |
| E5-S09 (DIR-2 goimports) | `.golangci.yml` + **all `*.go`** | E5 | **RUN ALONE** — repo-wide reformat collides with every Go lane |
| E5-S10 (TEST-3) | `.github/workflows/ci.yml` (53-56) | E5 | Low — isolated step |

## Recommendation (do first)
**E2-S06 (ARCH-1)** — highest untracked severity (P1), and it is the one that unblocks e2e: with credentials landing on L1-CRED, the remaining reason examples can't pass uptest is the five manifests carrying literal `${...}`. Wiring references removes that blocker and makes E2-S04/S05 real. Run **E5-S09 (goimports) alone and last** among the Go lanes to avoid a repo-wide reformat collision.
