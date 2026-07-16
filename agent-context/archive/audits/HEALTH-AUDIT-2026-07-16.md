# PlatformRelay Health & Direction Audit — provider-gridscale

- **Date:** 2026-07-16
- **Auditor:** read-only replayable audit (`/replayable-audit`)
- **Repo:** /Users/A242168/Projects/PlatformRelay/provider-gridscale
- **Commit audited:** `main` @ `6222151` (all 6 CI workflows reported green: CI, CodeQL, Coverage, Govulncheck, Gitleaks, Scorecard)
- **Prior baseline:** `HEALTH-AUDIT-2026-07-15.md` (`origin/main` @ `00e1c6c`)
- **Scope:** Upjet v2 Crossplane provider wrapping `gridscale/gridscale` Terraform provider. Generated `apis/`/`internal/controller/`/`package/crds/` boilerplate NOT audited line-by-line; findings target the hand-authored surface plus generated-metadata artifacts. Stale `.claude/worktrees/**` trees ignored (not part of `main`).

## Verdict: NEEDS-WORK (for v0.1.0)

No functional or P0/P1 blocker remains: the credential path is fixed and live-validated, the coverage floor is now CI-enforced, the generation boundary is machine-gated (arch-lint green), and the full security/supply-chain surface (govulncheck, CodeQL, Scorecard, gitleaks, keyless cosign + SBOM) has landed. Trajectory is strong and professional.

The operator's bar for cutting v0.1.0 is **zero UNBLOCKED backlog items**. That bar is not yet met: several UNBLOCKED hygiene/config items remain open (ARCH-1 references, ARCH-2 provider-metadata, DIR-2a/2c scaffold leftovers, TEST-3 dormant gate, DOC-2 examples, DOC-5 hand-maintained README matrix). None require the operator. Hence **NEEDS-WORK**, not AT-RISK, and not READY.

---

## Gate re-runs this session (not cached where noted)

| Gate | Command | Result |
|------|---------|--------|
| Unit tests | `go test ./config/... ./internal/clients/...` | PASS — config 100.0%, clients 87.1% |
| Coverage floor | `make coverage` (fresh gate) | **PASS — total 90.5% / floor 70%** |
| Vet | `go vet ./config/... ./internal/clients/... ./cmd/...` | exit 0 (clean) |
| Arch boundary | `go run go-arch-lint check` | **PASS — "OK - No warnings found"** (exit 0) |
| Vulnerabilities | `govulncheck ./config/... ./internal/clients/... ./cmd/...` | **0 affecting** (1 in imports, not called) |

---

## Dimension 1 — Architecture

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| ARCH-1 | P1 | **Still open.** No cross-resource references configured anywhere; all resources take raw UUID inputs. `config/*.go` contains zero `Reference{}`/`References:` entries (grep-confirmed). Generated examples still carry unresolved `${gridscale_ipv4...}` HCL interpolations that Crossplane treats as literals, so those examples cannot apply/uptest. | `config/gridscale.go` (no refs); `config/external_name.go` (no refs); `examples-generated/cluster/ssl/v1alpha1/certificate.yaml` |
| ARCH-2 | P2 | **Still open.** `config/provider-metadata.yaml` remains title-keyed for 8 entries (`ISO Image:`, `imported marketplace application:`, `marketplace application:`, `object storage access key:`, `object storage bucket:`, `storage backup schedule:`, `storage snapshot:`, `storage snapshot schedule:`) instead of `gridscale_*` keys; only 23 of 31 entries are `gridscale_`-keyed. Corrupted field text persists (`"...storage/ISO Image/ISO Image/snapshot..."`). This is a generated scrape artifact from upstream TF docs; fixing it is a config/regen lane (see UNBLOCKED). | `config/provider-metadata.yaml:3,10,1300,1341,1382,1396,1416,1441,1475` |
| ARCH-3 | P3 | **Resolved — verified correct (not a defect).** `IdentifierFromProvider` (via `.+`) is correct for the name-suspect resources. Schema check: `gridscale_object_storage_bucket.id` is `computed` (provider-assigned UUID) and `bucket_name` is a create-input, not the import ID; `storage_import`, `storage_clone`, `object_storage_accesskey` all expose a computed `id`. Minor residue: the unit test spot-checks only `storage_import`, not bucket/clone (coverage gap, not a bug). | `config/external_name.go:19-26`; `config/schema.json` (bucket `id: computed`); `config/external_name_test.go:29` |
| ARCH-4 | PASS | **Improved — now machine-gated.** Generation boundary is clean AND enforced: `.go-arch-lint.yml` models `apis`/`controller` as components so any hand-authored import of the generated tree is a hard violation; `make arch-lint` runs green. Dual cluster+namespaced scope still correctly wired. | `.go-arch-lint.yml:1-80`; arch-lint exit 0 |

## Dimension 2 — Tests

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| TEST-1 | **FIXED** | Was P1 (zero tests, vacuous gate). Now: 6 hand-authored test files (external-name, kind overrides, credential resolution, CRD golden-contract, fuzz). Coverage floor is **CI-enforced** via `make coverage` (`COVERAGE_MIN=70`) in `coverage.yml`; fresh run 90.5%. Race + tidy + arch targets exist. | `config/*_test.go`; `internal/clients/*_test.go`; `Makefile:299-313`; `.github/workflows/coverage.yml:35-39` |
| TEST-2 | P2 (operator-blocked) | **Still open.** e2e (`uptest`) wired but non-blocking: gated behind `/test-examples` maintainer comment + `UPTEST_CLOUD_CREDENTIALS`/`UPTEST_DATASOURCE` secrets. Cannot pass PR-blocking today; needs live lab creds as CI secrets AND ARCH-1 references so example interpolations resolve. | `.github/workflows/e2e.yaml:40,142,148` |
| TEST-3 | P3 | **Still open.** `report-breaking-changes` → `schema-version-diff` step still guarded by `if: ${{ inputs.upjet-based-provider }}`, an input never set by push/pull_request, so the native-schema-drift check never runs (dormant gate). | `.github/workflows/ci.yml:54` |

## Dimension 3 — Security

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| CRED-1 | **FIXED** (confirmed) | Credentials now wired: `ps.Configuration = buildProviderConfiguration(creds)` maps `uuid`/`token` (+ optional `api_url`). Live-validated at baseline. | `internal/clients/gridscale.go:59,68-77` |
| CRED-2 | **FIXED** (confirmed) | Correct keys `uuid`/`token` throughout; both secret templates use `"uuid"`/`"token"`. | `internal/clients/gridscale.go:27-29`; `examples/{cluster,namespaced}/providerconfig/secret.yaml.tmpl` |
| SEC-1 | **FIXED** | Was P2 (global gosec G101+G104 suppression). Now G101 is narrowly scoped to `internal/features/features.go` (a feature-flag const false-positive); G104 global suppression removed. gosec stays active on all other production code. | `.golangci.yml:102-108` |
| SEC-2 | PASS | Supply chain strong: all actions SHA-pinned; e2e uses `issue_comment` + admin/write permission gate (not `pull_request_target`); secrets referenced not echoed. Full E5 landed: govulncheck, CodeQL, Scorecard, gitleaks workflows; keyless cosign sign + SBOM attest on publish (with a visible warning when dispatched without a version). Codecov is **tokenless OIDC** (no `CODECOV_TOKEN` secret needed). | `.github/workflows/{govulncheck,codeql,scorecard,gitleaks,coverage}.yml`; `.github/workflows/publish-provider-package.yml:38-130`; `.golangci.yml`/`coverage.yml:19,45` |

## Dimension 4 — Docs

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| DOC-1 | **FIXED** | ProviderConfig example now uses `gridscale.platformrelay.io/v1beta1`. | `examples/cluster/providerconfig/providerconfig.yaml:1` |
| DOC-2 | P2 | **Still open (partial).** `examples-generated/` covers only 6 of ~20 resource groups (`gridscale, mysql8, paas, redis, ssl, storage`; 46 files). Mitigated by curated `examples/` (one per major group), but generated examples remain incomplete and carry the ARCH-1 `${...}` interpolations. | `examples-generated/cluster/` (6 groups) |
| DOC-3 | **FIXED** | README overhauled (logo, badges, quick-start, resource matrix, community links, install). No placeholder text. | `README.md:1-233` |
| DOC-4 | **FIXED** | Maintainer assigned: `@konih` (Konrad Heimel) in OWNERS.md, CODEOWNERS, and `package/crossplane.yaml`. | `OWNERS.md:10`; `CODEOWNERS:16`; `package/crossplane.yaml:6` |
| DOC-5 | P3 | **NEW.** README "Supported resources" is a hand-maintained API table (32 resources across 8 groups). GUIDELINES §7 says "CRD API reference and examples are generated; never hand-maintain API tables." Table is currently **accurate** (cross-checked vs `package/crds/`), but it is a drift-risk: a codegen change that adds/removes/renames a resource will silently desync unless someone hand-edits it. Recommend generating it or adding a CI staleness check. | `README.md:168-177` vs `package/crds/` (34 cluster CRD files); `agent-context/GUIDELINES.md:59` |

## Dimension 5 — Direction

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| DIR-2a | P2 | **Still open** (still-open portion of DIR-2). ChangeLogger provider label still hardcoded `provider-upjet-aws:%s` — every change-log event this provider emits is mislabelled as AWS. | `cmd/provider/main.go:211` |
| DIR-2b | **FIXED** | goimports local-prefix corrected to `github.com/PlatformRelay/provider-gridscale`. | `.golangci.yml:126-127` |
| DIR-2c | P2 | **Still open.** Three-way Go version drift persists: `go.mod` `go 1.25.9` / `toolchain go1.26.5`, `ci.yml`/`coverage.yml` `GO_VERSION 1.25`, `Makefile:47` `GO_REQUIRED_VERSION 1.24`, `e2e.yaml:12` `GO_VERSION 1.24`. Not a runtime break (GOTOOLCHAIN=auto), but a consistency/reproducibility gap. | `Makefile:47`; `.github/workflows/e2e.yaml:12`; `go.mod:3,5` |
| DIR-3 | **FIXED** | local-deploy download URL fixed (GitHub releases). | `Makefile` |
| DIR-4 | Info | No upstream `terraform-provider-gridscale` bug markers / workarounds surfaced in hand-authored config this run (no HACK/BUG/workaround in `config/`, `internal/clients/`, `cmd/`). Nothing to PR upstream from this audit. | (searched) |

---

## Delta vs 2026-07-15

### Fixed since last run (this run's genuine 07-15 -> 07-16 delta)
- **TEST-1 (P1 -> FIXED):** coverage floor now CI-enforced (`make coverage` in `coverage.yml`, floor 70, fresh 90.5%). At baseline it was "unit+contract+fuzz tests exist, CI enforcement pending". The pending piece landed (`5d054e5`/`05114ed`).
- **SEC-1 (P2 -> FIXED):** gosec G101 narrowed to `features.go:107`; global G104 suppression removed (`.golangci.yml:102-108`).
- **DOC-4 (P3 -> FIXED):** maintainer `@konih` assigned (`f79d1a0`).
- **DIR-2b (part of P2 -> FIXED):** goimports local-prefix corrected (`1e9cfe1`).

### Regressed
- **None.** All baseline-fixed items (CRED-1/2, DOC-1, DOC-3, DIR-3) remain fixed and re-verified this run.

### New
- **DOC-5 (P3):** README resource matrix is a hand-maintained API table — GUIDELINES §7 drift-risk. Currently accurate.

### Still open (carried forward)
- **ARCH-1 (P1):** no cross-resource references.
- **ARCH-2 (P2):** provider-metadata title-keying + corrupted fields.
- **DOC-2 (P2):** generated examples cover 6 of ~20 groups.
- **DIR-2a (P2):** `provider-upjet-aws:%s` changelog label.
- **DIR-2c (P2):** three-way Go version drift.
- **TEST-2 (P2, operator-blocked):** e2e non-blocking, needs live lab creds.
- **TEST-3 (P3):** dormant `schema-version-diff` gate.

### Resolved-as-correct (closed, not a defect)
- **ARCH-3:** `IdentifierFromProvider` verified correct via schema (`bucket.id` computed).

### Count reconciliation
- Newly FIXED this run: **4** (TEST-1, SEC-1, DOC-4, DIR-2b).
- Newly CLOSED-as-verified: **1** (ARCH-3).
- NEW: **1** (DOC-5).
- REGRESSED: **0**.
- Still OPEN: **8** = carried-forward **7** (ARCH-1, ARCH-2, DOC-2, DIR-2a, DIR-2c, TEST-2, TEST-3) + new **1** (DOC-5) — of which **7 UNBLOCKED** (ARCH-1, ARCH-2, DOC-2, DIR-2a, DIR-2c, TEST-3, DOC-5), **1 OPERATOR-BLOCKED** (TEST-2).

---

## UNBLOCKED items (actionable now, no operator input) — prioritized

1. **DIR-2a (P2)** — `cmd/provider/main.go:211`: change `"provider-upjet-aws:%s"` to `"provider-gridscale:%s"` (change-log events are mislabelled as AWS). One-line edit in `cmd/` (non-generated).
2. **ARCH-1 (P1)** — `config/gridscale.go`: add `Reference{}` wiring for the obvious edges (server<->ipv4/ipv6/storage, loadbalancer<->ssl_certificate/ipv4, k8s<->network), then `make generate`. Resolves the `${...}` interpolations blocking uptest examples.
3. **ARCH-2 (P2)** — `config/provider-metadata.yaml:3,10,1300+`: re-key the 8 title-keyed entries to `gridscale_*` (or regenerate metadata from the correct source) and fix the corrupted `capacity` field text. Config/regen lane.
4. **DIR-2c (P2)** — `Makefile:47`, `.github/workflows/e2e.yaml:12`: raise `GO_REQUIRED_VERSION`/e2e `GO_VERSION` to `1.25` to match `ci.yml`/`coverage.yml`/`go.mod`. Config edit.
5. **DOC-2 (P2)** — `examples-generated/`: generate/curate at least one applyable example per remaining API group (marketplace, object, plus missing gridscale kinds); depends partly on ARCH-1 so interpolations resolve.
6. **TEST-3 (P3)** — `.github/workflows/ci.yml:54`: the `schema-version-diff` gate is dead (`if: inputs.upjet-based-provider` never set). Either set the input for this upjet provider or drop the dormant step so the gate is honest.
7. **DOC-5 (P3)** — `README.md:168-177`: replace the hand-maintained resource matrix with a generated table or add a CI staleness check (GUIDELINES §7). Alternatively annotate it as generated-from-CRDs.
8. **ARCH-3 residue (P3, optional)** — `config/external_name_test.go:29`: extend the include-list spot-check to also assert `object_storage_bucket` and `storage_clone` map through `IdentifierFromProvider` (closes the coverage gap; the behaviour is already correct).

## OPERATOR-BLOCKED items

1. **TEST-2 / E2-S04 / E2-S05 (P2)** — uptest e2e needs live gridscale lab credentials wired as CI secrets (`UPTEST_CLOUD_CREDENTIALS`, `UPTEST_DATASOURCE`). Only the operator can provision the lab project + add the repo secrets. `.github/workflows/e2e.yaml:142,148`.
2. **Release dispatch / signing (process, not code)** — cutting v0.1.0 requires an operator-dispatched publish with an explicit version input so keyless cosign sign + SBOM attest run (the workflow warns and skips signing without it). `.github/workflows/publish-provider-package.yml:65`. (Not a code defect — a release action.)

> Note: Codecov (tokenless OIDC) and cosign (keyless) are NOT operator-blocked — no secret required. The baseline hypothesis that these would need operator secrets does not hold.

## Top 3 to fix next

1. **ARCH-1 (P1) — cross-resource references.** The single most valuable remaining lane: unblocks applyable examples (DOC-2) and makes uptest (TEST-2) runnable once creds land. `config/gridscale.go`. -> `/write-story`
2. **DIR-2a (P2) — `provider-upjet-aws` changelog label.** Trivial one-liner, but it mislabels every change-log event as AWS in a provider about to be published under the gridscale identity. `cmd/provider/main.go:211`. -> `/write-story`
3. **ARCH-2 (P2) — provider-metadata title-keying + corrupted fields.** Correctness of generated docs/metadata; a config/regen lane. `config/provider-metadata.yaml`. -> `/write-story`

## Proposed backlog stories (hand to `/write-story`)
- **STORY: Wire cross-resource references (ARCH-1)** — add `Reference{}` for server/loadbalancer/k8s edges in `config/gridscale.go`; regenerate; assert generated examples no longer contain `${...}`.
- **STORY: Fix changelog provider label (DIR-2a)** — `provider-upjet-aws:%s` -> `provider-gridscale:%s`; add a unit/M-level assertion on the label.
- **STORY: Repair provider-metadata keys (ARCH-2)** — re-key title entries to `gridscale_*`, fix corrupted field text, add a check that all 31/32 entries are `gridscale_`-keyed.
- **STORY: Unify Go version (DIR-2c)** — single source of truth for Go version across `go.mod`/`ci`/`coverage`/`e2e`/`Makefile`.
- **STORY: Generate/verify README resource matrix (DOC-5)** — generated table or CI staleness gate per GUIDELINES §7.
