# PlatformRelay Health & Direction Audit — provider-gridscale (re-run)

- **Date:** 2026-07-16 (run **b** — second audit of the day; supersedes the `6222151` snapshot)
- **Auditor:** read-only replayable audit (`/replayable-audit`)
- **Repo:** /Users/A242168/Projects/PlatformRelay/provider-gridscale
- **Commit audited:** `main` @ `a1748f3`
- **Prior baseline this run diffs against:** `HEALTH-AUDIT-2026-07-16.md` (`main` @ `6222151`)
- **Scope:** Upjet v2 Crossplane provider wrapping `gridscale/gridscale` TF provider. Generated `apis/`/`internal/controller/`/`package/crds/` boilerplate NOT audited line-by-line. Every prior "still open" finding was **re-verified at HEAD `a1748f3`** (not trusted from the prior report): 25 commits landed between `6222151` and `a1748f3`, several of which mitigate prior findings.

## Verdict: NEEDS-WORK — but only barely; effectively READY-for-v0.1.0 on code

Every P0/P1 is closed and verified at HEAD. All five task-flagged "confirm true state" items were re-checked from source: **ARCH-1, ARCH-2, DIR-2a, TEST-3 are genuinely FIXED at HEAD.** What remains is two low-severity hygiene items (a Go-version guard blindspot and an un-wired README staleness gate) plus one long-standing operator-blocked item (live e2e creds) and a definitional DOC-2 residue. No UNBLOCKED item is a functional or security defect. The operator's stated bar for cutting v0.1.0 is "zero UNBLOCKED backlog items"; two small UNBLOCKED items remain, so strictly **NEEDS-WORK**, not READY. Not AT-RISK.

---

## Gate matrix — re-run this session at `a1748f3` (nothing cached)

| Gate | Command | Result |
|------|---------|--------|
| Unit tests | `go test ./config/... ./internal/clients/...` | **PASS — 58 tests, 2 packages** |
| Coverage floor | `make coverage` | **PASS — total 93.0% / floor 70%** |
| Vet | `go vet ./config/... ./internal/clients/... ./cmd/...` | exit 0 (clean) |
| Arch boundary | `make arch-lint` | **PASS — "OK - No warnings found"** |
| Vulnerabilities | `make vuln` (govulncheck ./...) | **0 affecting** (1 in imports, not called) |
| Go-version guard | `make go-version-check` | PASS — but see DIR-2c (guard has a coverage hole) |
| README matrix | `make check-docs` | PASS — 32/8 in sync — but see DOC-5 (not CI-wired) |

---

## Per-dimension findings (verified at `a1748f3`)

### Architecture
| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| ARCH-1 | **FIXED** | Cross-resource references now fully wired. `config/references.go` (`configureReferences`) defines ~14 `Reference{}` edges (server→ipv4/ipv6/isoimage/network/firewall/storage, loadbalancer→ipv4/ipv6/ssl_certificate + backend host via custom `ExtractParamPath("ip")`, template→snapshot, storage_clone→storage, snapshot/schedule→storage, PaaS family→network). Applied by both cluster+namespaced constructors (`config/provider.go:34,59`). ADR-0005 recorded. **Symptom fully cleared: zero raw `${gridscale_...}` interpolations remain in `examples-generated/`** — re-emitted as `Ref:`/`Selector:` fields. | `config/references.go:53-135`; `config/provider.go:34,59` |
| ARCH-2 | **FIXED** | provider-metadata is fully `gridscale_*`-keyed. **32/32 resource keys** are `gridscale_<name>:`; **zero title-keyed leftovers** (the 8 prior `ISO Image:`/`marketplace application:`/… entries are gone). `hack/metadatafix` post-scrape step injects stubs (`2e7583e`); a unit test asserts keys match schema (`94c24bb`). The `capacity` field text (`.../ISO Image/template/snapshot...`) is upstream's real doc string, not corruption. | `config/provider-metadata.yaml` (32 `gridscale_*` keys) |
| ARCH-3 | closed-as-correct | Unchanged — `IdentifierFromProvider` verified correct via schema at prior run. | `config/external_name.go` |
| ARCH-4 | PASS | Generation boundary machine-gated; arch-lint green this run. | `.go-arch-lint.yml`; arch-lint exit 0 |

### Tests
| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| TEST-1 | **FIXED** (holds) | Coverage floor CI-enforced; fresh run **93.0% / floor 70%**. 58 unit tests pass. | `Makefile` coverage target; `coverage.yml` |
| TEST-2 | **P2 — OPERATOR-BLOCKED** | e2e (`uptest`) wired but non-blocking: `/test-examples` comment-gated + needs `UPTEST_CLOUD_CREDENTIALS`/`UPTEST_DATASOURCE` CI secrets. Cannot PR-gate today. Now that ARCH-1 references resolve example interpolations, the *only* remaining blocker is live lab creds — operator-provisioned. | `.github/workflows/e2e.yaml:40,123,142,148` |
| TEST-3 | **FIXED** | schema-version-diff gate revived. No longer guarded by the dead `inputs.upjet-based-provider`; now runs on `github.event_name == 'pull_request' && modified-schema (config/schema.json) any_changed` with a proper base-branch fetch. Skips cleanly on push / non-schema PRs. | `.github/workflows/ci.yml:64-68` |

### Security
| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| CRED-1/2, SEC-1, SEC-2 | **FIXED / PASS** (hold) | Credentials wired (uuid/token), gosec scoped, supply-chain (govulncheck/CodeQL/Scorecard/gitleaks, keyless cosign+SBOM) intact. govulncheck 0 affecting this run. | `internal/clients/gridscale.go`; `.golangci.yml`; `.github/workflows/*` |

### Docs
| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| DOC-2 | **P3 — UNBLOCKED (definitional residue)** | Curated `examples/` covers **all 8 API groups** (gridscale, marketplace, mysql8, object, paas, redis, ssl, storage — each with `example-id`-annotated manifests) → the user-facing/uptest surface is complete. **`examples-generated/cluster/` still emits only 6 of 8 groups** (gridscale, mysql8, paas, redis, ssl, storage; 23 files) — no generated `marketplace`/`object` examples. Downgraded P2→P3: curated coverage makes this cosmetic, not a functional gap. | `examples-generated/cluster/` (6 dirs) vs `examples/` (8 groups) vs `apis/cluster/` (8 groups) |
| DOC-5 | **P3 — UNBLOCKED** | `make check-docs` (README 32/8 matrix staleness guard, `hack/check-docs.sh`) **exists and passes**, but is **NOT wired into any CI gate.** Unlike `go-version-check` (pulled into CI via `check-diff: go-version-check`), there is no `check-diff: check-docs` / `reviewable: check-docs` edge and no workflow invokes it. So the README matrix guard is manual-only — a codegen change that adds/removes a resource still merges without running it. | `Makefile:322-326` (target, no RHS wiring); no `check-docs` in `.github/workflows/**` |

### Direction
| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| DIR-2a | **FIXED** | ChangeLogger label now `provider-gridscale:%s`. No `provider-upjet-aws` remains. | `cmd/provider/main.go:211` |
| DIR-2c | **P3 — UNBLOCKED (mutated / NEW shape)** | The 4 files the guard watches (Makefile `GO_REQUIRED_VERSION=1.26.5`, ci.yml `1.26.5`, e2e.yaml `1.26.5`, go.mod `toolchain go1.26.5`) are consistent and `make go-version-check` reports **OK**. **But 4 workflows the guard does not read still pin `setup-go '1.25'`: `codeql.yml:38`, `coverage.yml:32`, `govulncheck.yml:29`, `docs-sync.yml:40`.** The finding is no longer plain "3-way drift" — it is a **guard blindspot giving false assurance**: bump Go, update the 4 watched files, guard goes green, these 4 workflows silently stay on 1.25. Most concrete impact: govulncheck analyzes on a different toolchain than the one that ships. Fix = add these 4 files to `hack/check-go-version.sh` and bump them to 1.26.5. Severity P3 (reproducibility/consistency, no runtime break — GOTOOLCHAIN=auto). | `hack/check-go-version.sh:1-13` (file list); `codeql.yml:38`; `coverage.yml:32`; `govulncheck.yml:29`; `docs-sync.yml:40` |
| DIR-4 | Info | No upstream TF-provider bug markers in hand-authored config this run. | (searched) |

---

## Delta vs 2026-07-16 (`6222151` → `a1748f3`)

### Fixed since last run (verified from source at HEAD)
- **ARCH-1 (P1 → FIXED):** `config/references.go` wired + applied both constructors; generated examples now `Ref:`/`Selector:` (zero `${...}`).
- **ARCH-2 (P2 → FIXED):** 32/32 metadata keys `gridscale_*`; title-keyed entries gone; regen guard test added.
- **DIR-2a (P2 → FIXED):** changelog label `provider-gridscale:%s`.
- **TEST-3 (P3 → FIXED):** schema-version-diff PR-gated on `config/schema.json` change (dead `upjet-based-provider` guard removed).

### Regressed
- **None.** All prior FIXED items (CRED-1/2, DOC-1/3/4, SEC-1, DIR-2b, DIR-3, TEST-1) re-verified holding.

### New / mutated
- **DIR-2c mutated (P2 → P3):** narrowed from 3-way drift to a *guard coverage hole* — 4 workflows (codeql/coverage/govulncheck/docs-sync) still on Go 1.25, outside the `check-go-version.sh` watch list.
- **DOC-5 refined (P3):** the `make check-docs` guard now exists (prior recommendation delivered) but is **not CI-wired**, so it is only half-closed.

### Still open at HEAD
- **DIR-2c (P3, UNBLOCKED)** — Go-version guard blindspot.
- **DOC-5 (P3, UNBLOCKED)** — README-matrix guard not wired to CI.
- **DOC-2 (P3, UNBLOCKED)** — generated examples cover 6/8 groups (curated covers 8/8).
- **TEST-2 (P2, OPERATOR-BLOCKED)** — live e2e creds.

### Count reconciliation
- Newly FIXED this run: **4** (ARCH-1, ARCH-2, DIR-2a, TEST-3).
- Still OPEN: **4** = **3 UNBLOCKED** (DIR-2c P3, DOC-5 P3, DOC-2 P3) + **1 OPERATOR-BLOCKED** (TEST-2 P2).
- By severity: P0×0 · P1×0 · P2×1 (TEST-2, operator-blocked) · P3×3.

---

## Currently-open findings — full list with UNBLOCKED vs OPERATOR-BLOCKED

**UNBLOCKED (actionable now, no operator input):**
1. **DIR-2c (P3)** — add `codeql.yml`, `coverage.yml`, `govulncheck.yml`, `docs-sync.yml` to `hack/check-go-version.sh`'s watch list AND bump their `setup-go` from `1.25` to `1.26.5`. Closes the guard blindspot. `hack/check-go-version.sh`; the four workflow files.
2. **DOC-5 (P3)** — wire the existing `make check-docs` into CI (e.g. `check-diff: check-docs`, mirroring the `go-version-check` pattern) so the README 32/8 matrix guard actually runs on PRs. `Makefile:338` area.
3. **DOC-2 (P3)** — regenerate `examples-generated/` so `marketplace` + `object` groups get generated examples (curated `examples/` already covers all 8), or explicitly document that generated examples are limited to uptest-relevant groups. `examples-generated/cluster/`.

**OPERATOR-BLOCKED:**
4. **TEST-2 (P2)** — uptest e2e needs `UPTEST_CLOUD_CREDENTIALS` + `UPTEST_DATASOURCE` provisioned as CI secrets against a gridscale lab project. Only the operator can provision the lab + add repo secrets. `.github/workflows/e2e.yaml:142,148`.

Plus the non-code release action (not a defect): operator-dispatched `publish-provider-package` with an explicit version input so keyless cosign sign + SBOM attest run.

## Top 3 to fix next
1. **DIR-2c (P3)** — the guard blindspot is the only finding with a correctness edge (silent toolchain skew on the vuln/coverage jobs) and it is a ~10-line UNBLOCKED fix. → `/write-story`
2. **DOC-5 (P3)** — one-line Makefile wiring to make the README staleness guard real (GUIDELINES §7). → `/write-story`
3. **DOC-2 (P3)** — regenerate the two missing generated-example groups or annotate the scope. → `/write-story`

## Proposed backlog stories (hand to `/write-story`)
- **STORY: Close Go-version guard blindspot (DIR-2c)** — extend `check-go-version.sh` to codeql/coverage/govulncheck/docs-sync and bump them to 1.26.5; add a regression assertion.
- **STORY: Wire README-matrix staleness gate into CI (DOC-5)** — `check-diff: check-docs` (mirror go-version-check).
- **STORY: Complete generated examples (DOC-2)** — regenerate marketplace/object groups or document the limited scope.
