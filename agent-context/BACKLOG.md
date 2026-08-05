# BACKLOG — provider-gridscale (private lane detail)

Public build order: [docs/ROADMAP.md](../docs/ROADMAP.md). OpenSpec changes: `openspec/changes/`.

## Status

| Phase | State |
| --- | --- |
| Codegen (E6g-S01 in kaddy's terms) | **complete** — 32 resources generated & building |
| Design (openspec, ADRs, ROADMAP, backlog) | **complete** — E1 landed |
| Implementation E1–E7 + Batch 7 polish | **complete** on `main` (see OPERATOR-BOARD; tip after wrap) |
| E4-S04 (docs site) | **dropped** (D-002→A, provider-native docs) |
| E2-S04/S05 uptest | **intentionally skipped** (D-012→B — manual smoke only) |
| E6-S05 (assurance) | **landed** (`3edc5f0`); was stretch |
| L-CIRED (CI restore) | **landed** (`d8433f5`+`d275897`) — primary CI green |
| E8 datasources | **rescoped** (D-015) — document omission + track upstream; no codegen |
| U-1 + LB-1 overrides | **landed** — S3/console creds sensitive (`8ae7376`, `config/sensitive.go`); loadbalancer `status` Computed (`c38e52b`, `config/loadbalancer.go`); D-021 |
| D-020-FU (extensions in signed release) | **landed** (`43294ed`+`60b9e8f`) — `make xpkg.append.extensions` append-then-sign + fail-closed verify in publish workflow; v0.2.2 signed w/ extensions verified |
| E5 SonarCloud SECURITY (E5-S11…S17) | **landed 2026-07-25** — 7 lanes Integrated; 9/9 `hack/test/sonar_pg_*` green; SonarCloud SECURITY open: 0 |
| E5 SonarCloud MAINT/RELIABILITY (E5-S18…S21) + coverage CI | **landed 2026-08-04** — PR #41 → `f9f4628`; OpenSpec `sonar-maintainability-2026-08` |
| Backlog | **exhausted** for implementable lanes. Operator: wait on upstream TF #509/#510/#511 (on merge, re-vendor + drop U-1/LB-1). Optional Scorecard polish (D-016) non-blocking. No package release for Batch 11 alone (CI/maintainability). |

## E2 test-hardening batch (S06–S10) — ported from kollect's test tooling

Appended to the E2 story table in [docs/ROADMAP.md](../docs/ROADMAP.md); levels per ADR-0002 (U/E/M/D):

| Story | Adds (local, auto-mergeable now) | Level |
| --- | --- | --- |
| E2-S06 | `make vuln` → `go run golang.org/x/vuln/cmd/govulncheck@v1.1.4 ./...` | M |
| E2-S07 | `.go-arch-lint.yml` + `make arch-lint` — generation boundary: `config`/`internal/clients` must **not** import generated `internal/controller/**` | M |
| E2-S08 | `make test.race` → `go test -race -count=1` | U |
| E2-S09 | `FuzzGetExternalName` native fuzz test on the `config` external-name extraction | U |
| E2-S10 | `make tidy-check` → `go mod tidy` + `git diff --exit-code go.{mod,sum}` | M |

> **CI split (harness guardrail):** the `make` targets + tests above land **now** (auto-mergeable —
> paths `Makefile`, `config/*_test.go`, `.go-arch-lint.yml`). **Wiring each tool into CI
> (`.github/workflows/`) is deferred to Epic E5** (CI/release = surfaced-not-auto-merged). The
> govulncheck CI job in particular is the **existing E5-S01 follow-up**, not new scope; arch-lint /
> race / fuzz / tidy-check gating likewise become E5 CI jobs.

## Lane dependency graph

```
E1 (foundation)
 ├─ E2 tests        (parallel after E1)
 ├─ E3 branding     (parallel after E1)
 ├─ E4 docs         (parallel after E1; E4-S01/02 need no assets; E4-S04 dropped per D-002→A)
 ├─ E6 governance   (parallel after E1)
 └─ E5 CI/supply    (after E2 coverage lands + E3 metadata/badges; E5-S04 anytime)
```

Disjoint path sets (safe to run as parallel worktrees):
- **E2** → `config/*_test.go`, `examples/**`, `test/**`, Makefile test targets, `.go-arch-lint.yml`
- **E3** → `extensions/**`, `docs/assets/**`, `README.md`, `package/crossplane.yaml`
- **E4** → `docs/api/**`, `docs/adr/**`, `.crd-ref-docs.yaml`, `mkdocs.yml`
- **E6** → `CONTRIBUTING.md`, `SECURITY.md`, `GOVERNANCE.md`, `CODEOWNERS`, `OWNERS.md`, `.github/*_TEMPLATE*`
- **E5** → `.github/workflows/**`, `.pre-commit-config.yaml`, `codecov.yml`, `cliff.toml`, `.github/gitleaks.toml`

Collision watch: **E3-S02 (README)** and **E3-S04 (badges)** both touch `README.md` — sequence them.
**E5** edits `.github/workflows/` which the stock upjet template owns; prefer *adding* workflow files
over editing `ci.yml` where possible.

## Story ID scheme

`E<n>-S<nn>` — e.g. `E2-S01`, `E5-S06`. REQ IDs: `REQ-E<n>-S<nn>-<mm>` (see `openspec/config.yaml`).
OpenSpec change slug: `e2-test-foundation`, `e5-ci-supplychain`, … (lowercase, hyphenated).

> **Numbering note (D-001):** this repo uses a *fresh, repo-local* E1–E6 scheme. Kaddy tracks this
> whole provider as its epic **E6g** — that is kaddy's external handle for us and does not appear in
> this repo's IDs (workspace rule: repos never cross-reference). The one legacy touchpoint is the
> existing commit `…(E6g-S01)` for the initial generation; we do not renumber history.

## Provider-specific hazards (carry into every lane)

- **Never hand-edit generated code**: `apis/**/zz_*.go`, `internal/controller/**`, `package/crds/**`.
  Route all changes through `config/` or codegen inputs. `make reviewable` / `make check-diff` guard this.
- **BSL terraform pin**: `TERRAFORM_VERSION=1.5.7` is deliberate (BSL after that). Don't bump it.
- **Kind overrides exist**: `config/gridscale.go` overrides mysql8 (avoids Go kind "0") and
  storageimport (avoids Go keyword `Import`) — E2-S02 tests must assert these specifically.
- **golangci local-prefix is stale**: points at `upjet-provider-template`, not
  `PlatformRelay/provider-gridscale` — worth a fix (belongs in E5-S03 or a tidy commit).
- **uptest needs real creds**: E-level tests hit the live gridscale API; keep them behind
  `/test-examples` + nightly, never on every push. Document the creds contract in E2-S05.

## E7 — Config correctness (audit gap-stories, renumbered per D-014)

The 2026-07-15 audit gap-stories proposed IDs `E2-S06…S08`, which clash with the integrated
test-hardening batch above — renumbered 2026-07-16 (D-014). Story bodies live in
[AUDIT-GAP-STORIES-2026-07-15.md](AUDIT-GAP-STORIES-2026-07-15.md) (epic-agnostic, as written there):

| Story | Was (gap doc) | Finding | Path-set |
| --- | --- | --- | --- |
| **E7-S01** | "E2-S06" | ARCH-1 — wire cross-resource `config.Reference{}` (edge list in [PROVIDER-DOCS-RESEARCH-2026-07-15.md](PROVIDER-DOCS-RESEARCH-2026-07-15.md) Q3) | `config/*.go` (non-test), `examples-generated/**` |
| **E7-S02** | "E2-S07" | ARCH-2 — `provider-metadata.yaml` title-keyed entries | `config/provider-metadata.yaml`, `apis/generate.go`, `config/metadata_test.go` |
| **E7-S03** | "E2-S08" | ARCH-3 — external-name rationale comment + import-format docs | `config/external_name.go` (comment), `docs/adr/` |

**E2-S11** (was research "E2-S09", rescoped per D-012 — credential-free parts only): non-live
credential-wiring regression test (`internal/clients/*_test.go`) + creds-contract doc. E5-S07…S10
keep their gap-doc IDs. Research stories E4-S05 / E6-S06 / epic E8 as proposed in the research doc.

## E6-S06 — Upstream TF-provider triage (D-016, 2026-07-16)

Research Q5 closed in [`decisions.md`](decisions.md) **D-016** — summary for lane planners:

| Upstream | Action for *this* provider |
| --- | --- |
| Doc #200 / #194 | Track open drafts [#467](https://github.com/gridscale/terraform-provider-gridscale/pull/467) / [#468](https://github.com/gridscale/terraform-provider-gridscale/pull/468); no in-repo code work. File-rename follow-up only if drafts stall (recipe in D-016). |
| Feature #187 location | **track-upstream** — no `gridscale_location` in schema; do not hand-implement. |
| Feature #188 backup location | **done upstream** (PR #193 merged); already in our `BackupSchedule` CRDs. No lane. |

No new stories spawned from this triage.

## E5 — SonarCloud SECURITY remediation (2026-07)

OpenSpec (worktree / upcoming PR): `openspec/changes/sonar-security-remediation-2026-07/`
(REQs PG-01…PG-08). **Lane rule:** file locks below are exclusive; meta tests under
`hack/test/sonar_pg_*` / `scripts/version_diff_test.py` owned by the same lane.

| Story | Closes | File lock | P | Status |
| --- | --- | --- | --- | --- |
| E5-S11 | PG-01, PG-06 (publish curl) | `.github/workflows/publish-provider-package.yml`, `hack/test/sonar_pg_01_*`, `hack/test/sonar_pg_06_publish_*` | P0 | ✅ Integrated (a9c6efe) |
| E5-S12 | PG-02, PG-03, PG-04 (e2e) | `.github/workflows/e2e.yaml`, `hack/test/sonar_pg_02_*`, `sonar_pg_03_*`, `sonar_pg_04_e2e_*` | P0 | ✅ Integrated (6a6aa5f) |
| E5-S13 | PG-04 (ci) | `.github/workflows/ci.yml`, `hack/test/sonar_pg_04_ci_*` | P1 | ✅ Integrated (21dfb81) |
| E5-S14 | PG-05 | `.github/workflows/scorecard.yml`, `hack/test/sonar_pg_05_*` | P1 | ✅ Integrated (b55ae7a) |
| E5-S15 | PG-06 (gitleaks) | `.github/workflows/gitleaks.yml`, `hack/test/sonar_pg_06_gitleaks_*` | P1 | ✅ Integrated (1da4080) |
| E5-S16 | PG-07 | `scripts/version_diff.py`, `scripts/version_diff_test.py` | P1 | ✅ Integrated (48aacef) |
| E5-S17 | PG-08 | `cluster/images/provider-gridscale/Dockerfile`, `hack/test/sonar_pg_08_*` | P1 | ✅ Integrated (ef6053a) |

**Batch landed 2026-07-25** (E5-S11…S17 all Integrated on `main`; 9/9 `hack/test/sonar_pg_*` meta tests green; SonarCloud SECURITY open: 0).

### E5-S11 — Publish workflow: no input injection + HTTPS curl (P0)
**Outcome:** `workflow_dispatch` version and curl installs cannot inject shell or downgrade TLS.

**Acceptance**
- Given Build/Publish steps, when `inputs.version` is set, then it reaches the shell only via `env:` / `$VERSION` — never `${{ inputs.* }}` / `format(...)` inside `run:`.
- Given the Upbound/`up` curl step, when it runs, then curl uses `--proto '=https' --tlsv1.2`.
- **Edge:** given a metacharacter version (`v1"; echo pwned; "`), when dispatched, then no injected command runs.

**Done when:** failing meta tests first → green Verify; closes **REQ-SONAR-PG-01** + **REQ-SONAR-PG-06** (publish site); gates green.
**File lock:** `.github/workflows/publish-provider-package.yml`, `hack/test/sonar_pg_01_*`, `hack/test/sonar_pg_06_publish_*`
**Depends on:** none · **Parallel-safe-with:** E5-S12…S17 · **Not in scope:** e2e/gitleaks curls

### E5-S12 — e2e.yaml: secret via env, sha256 hash, job-scoped perms (P0)
**Outcome:** e2e stops expanding secrets unsafely; hashes with sha256; scopes `contents: read` to jobs.

**Acceptance**
- Datasource secret only via `env:` + `"$UPTEST_DATASOURCE"` in `run:`.
- Example-list hash uses `sha256sum`, not `md5sum`.
- No workflow-level `contents: read`; jobs declare needs.
- **Edge:** empty/unset datasource → explicit failure, not silent empty success.

**Done when:** Verify `sonar_pg_02/03/04_e2e`; closes **REQ-SONAR-PG-02/03/04** (e2e half).
**File lock:** `.github/workflows/e2e.yaml` + listed meta tests
**Depends on:** none · **Parallel-safe-with:** E5-S11, S13–S17

### E5-S13 — ci.yml job-scoped permissions (P1)
**Outcome:** drop workflow-level `contents: read`; per-job `permissions`.
**Edge:** jobs that need read still declare it so checkout does not 403.
**Verify:** `bash hack/test/sonar_pg_04_ci_job_scoped_permissions_test.sh` · closes **REQ-SONAR-PG-04** (ci half)
**File lock:** `.github/workflows/ci.yml`, `hack/test/sonar_pg_04_ci_*` · **Parallel-safe-with:** all except another `ci.yml` lane

### E5-S14 — Scorecard explicit permissions (P1)
**Outcome:** replace `permissions: read-all` with minimal scopes (`security-events: write`, `id-token: write`, `contents: read`, `actions: read`).
**Edge:** SARIF upload still has `security-events: write`.
**Verify:** `bash hack/test/sonar_pg_05_scorecard_permissions_test.sh` · closes **REQ-SONAR-PG-05**
**File lock:** `.github/workflows/scorecard.yml`, `hack/test/sonar_pg_05_*`

### E5-S15 — gitleaks curl HTTPS enforce (P1)
**Outcome:** gitleaks installer curl uses `--proto '=https' --tlsv1.2`.
**Verify:** `bash hack/test/sonar_pg_06_gitleaks_curl_https_test.sh` · closes **REQ-SONAR-PG-06** (gitleaks)
**File lock:** `.github/workflows/gitleaks.yml`, `hack/test/sonar_pg_06_gitleaks_*`

### E5-S16 — version_diff.py path bound to repo root (P1)
**Outcome:** CLI paths `resolve()`d; reject escapes outside repo root.
**Edge:** `../../etc/passwd`-shaped path → non-zero, no open outside root.
**Verify:** `python3 -m pytest scripts/version_diff_test.py -q` · closes **REQ-SONAR-PG-07**
**File lock:** `scripts/version_diff.py`, `scripts/version_diff_test.py`

### E5-S17 — Dockerfile local ADD → COPY (P1)
**Outcome:** local `bin/…/provider` + `terraformrc.hcl` use `COPY`; remote URL ADD/curl allowed.
**Verify:** `bash hack/test/sonar_pg_08_dockerfile_copy_test.sh` · closes **REQ-SONAR-PG-08**
**File lock:** `cluster/images/provider-gridscale/Dockerfile`, `hack/test/sonar_pg_08_*`

---

## E5 — SonarCloud MAINTAINABILITY + RELIABILITY remediation (2026-08)

OpenSpec (create when first lane starts): `openspec/changes/sonar-maintainability-2026-08/`
(REQs **PG-M01…PG-M08**). Source: SonarCloud project `PlatformRelay_provider-gridscale`,
**46 open CODE_SMELL** (Maintainability 38 · Reliability 8). Zero SECURITY / BUG.

**Out of scope for this batch:** coverage CI wiring (`sonar-coverage-ci` @ `d98f242`);
generated `zz_*.go` smells; SECURITY (already 0).

**Lane rule:** file locks below are exclusive. Clusters that share a file are **one story /
one lane** (never parallel). Meta tests under `hack/test/sonar_pg_m_*` owned by the same lane
as the production path they guard.

### agent-loop-local lane plan (N=4, file-disjoint)

| Lane | Story | Closes (Sonar) | File lock | P | Parallel? |
| --- | --- | --- | --- | --- | --- |
| **L-REG** | E5-S18 | go:S1186 ×2 CRITICAL | `apis/cluster/v1alpha1/register.go`, `apis/namespaced/v1alpha1/register.go`, `hack/test/sonar_pg_m01_*` | P0 | ✅ vs all |
| **L-CFGTEST** | E5-S19 | go:S3776 ×3 CRITICAL + godre:S8205 ×6 MINOR | `config/crd_contract_test.go`, `config/external_name_fuzz_test.go`, `hack/test/sonar_pg_m02_*` (optional structural) | P0 | ✅ vs all **except** any other `config/*_test.go` lane |
| **L-SHELL** | E5-S20 | shelldre:S131 ×1 + S7677 ×3 + S7688 ×8 + S1066 ×4 | `hack/check-docs.sh`, `hack/check-api-docs.sh`, `hack/test/e8_s01_*`, `hack/test/e8_s02_*`, `hack/test/e8_s03_*`, `hack/test/sonar_pg_02_*`, `hack/test/sonar_pg_04_ci_*`, `hack/test/sonar_pg_04_e2e_*`, `hack/test/sonar_pg_m04_*`…`m07_*` | P0/P1 | ✅ vs L-REG / L-CFGTEST / L-DOCKER |
| **L-DOCKER** | E5-S21 | docker:S6570 ×19 MAJOR | `cluster/images/provider-gridscale/Dockerfile`, `hack/test/sonar_pg_m08_*` (extend or sibling of `sonar_pg_08_*`) | P1 | ✅ vs all **except** another Dockerfile lane |

**Serialization notes**
- Clusters **2 + 8** both touch `config/crd_contract_test.go` → **only E5-S19** (L-CFGTEST). Do not split.
- Clusters **3 + 5 + part of 6** all touch `hack/check-docs.sh` → **only E5-S20** (L-SHELL).
- Cluster **6** e8 scripts + cluster **7** sonar_pg nested-ifs share the `hack/test/` tree but
  **disjoint filenames** with each other; kept in **one lane** so shell style stays consistent and
  one agent owns all Sonar shell hygiene.
- **L-DOCKER** re-touches the Dockerfile previously owned by landed E5-S17 — preserve `COPY` for
  local artefacts; only add quoting for expansions (lines 27–39).
- Base worktrees on a tip that includes the E8 observe meta tests
  (`hack/test/e8_s01_*`, `e8_s03_*` — present on `origin/main`; may be missing on stale local
  checkouts). Independent of `sonar-coverage-ci`.

| Story | REQs | Status |
| --- | --- | --- |
| E5-S18 | PG-M01 | ✅ Integrated (PR #41) |
| E5-S19 | PG-M02, PG-M03 | ✅ Integrated (PR #41) |
| E5-S20 | PG-M04, PG-M05, PG-M06, PG-M07 | ✅ Integrated (PR #41) |
| E5-S21 | PG-M08 | ✅ Integrated (PR #41) |
| coverage CI | sonar_pg_09 | ✅ Integrated (PR #41) |

### E5-S18 — Empty `init()` register funcs: nested comments (P0)
**As a** maintainer **I want** the package `init()` stubs in cluster/namespaced `register.go` to
satisfy go:S1186 **so that** Sonar CRITICAL empty-function smells clear without changing scheme
registration behaviour.

**Sonar keys / rule**
- `apis/cluster/v1alpha1/register.go:22` — `AZ9magFP_3fq4FMeL0wZ` (go:S1186 CRITICAL)
- `apis/namespaced/v1alpha1/register.go:22` — `AZ9magJP_3fq4FMeL0xy` (go:S1186 CRITICAL)

**Acceptance criteria**
- Given the Upjet scaffold `func init() {}` at both paths, when the fix lands, then each empty
  function body contains a **nested comment** explaining why it is intentionally empty (types
  register via `SchemeBuilder` / generated `zz_*.go`), meeting Sonar S1186 — **or** an equivalent
  no-op that Sonar accepts without registering types twice.
- Given `make generate` / `make reviewable`, when run after the edit, then neither `register.go` is
  overwritten back to a bare empty body (if codegen regenerates them, the intentional comment is
  restored in the same change or the comment lives in a non-generated sibling — document which).
- **Edge:** given a mistaken “fix” that calls `SchemeBuilder.AddToScheme` from both package
  `init()`s, when controllers start, then we must **not** double-register; prefer nested comment
  over wiring registration here.

**Done when**
- [ ] Meta test asserts both files’ `init` bodies are non-bare-empty (comment or documented pattern). **Level:** M · **Test:** `hack/test/sonar_pg_m01_register_init_not_empty_test.sh` · **Verify:** `bash hack/test/sonar_pg_m01_register_init_not_empty_test.sh`
- [ ] `go test ./apis/...` (or package compile) still green; `make reviewable` clean
- [ ] Closes Sonar keys above (or marked Fixed on next scan)
- [ ] OpenSpec REQ-SONAR-PG-M01 filled under `sonar-maintainability-2026-08`

**Touches:** `apis/cluster/v1alpha1/register.go`, `apis/namespaced/v1alpha1/register.go`,
`hack/test/sonar_pg_m01_*`  
**Depends on:** none · **Parallel-safe-with:** E5-S19, E5-S20, E5-S21  
**Not in scope:** any `zz_*.go`; changing Group/Version; coverage CI

### E5-S19 — Config test cognitive complexity + named CRD view types (P0)
**As a** maintainer **I want** CRD-contract and external-name fuzz tests under the cognitive-complexity
threshold and anonymous nested structs replaced with named types **so that** Sonar CRITICAL/MINOR
smells in `config/*_test.go` clear without weakening golden/structural assertions.

**Sonar keys / rules**
- Cognitive complexity (go:S3776 CRITICAL):  
  - `config/crd_contract_test.go:164` (21→≤15) `AZ9nqLQ5RbNgprH5pBN9` — `TestCRDGoldenContract`  
  - `config/crd_contract_test.go:201` (28→≤15) `AZ9nqLQ5RbNgprH5pBN-` — `TestCRDStructuralInvariants`  
  - `config/external_name_fuzz_test.go:18` (20→≤15) `AZ9n-G-75WRcdOwb-njc` — `FuzzGetExternalName`
- Named types (godre:S8205 MINOR) — anonymous structs at `config/crd_contract_test.go:41–53`
  keys `AZ9nqLQ5RbNgprH5pBN_` … `AZ9nqLQ5RbNgprH5pBOE` (6 issues)

> **Must stay one story:** complexity + named types both edit `crd_contract_test.go`.

**Acceptance criteria**
- Given the three functions above, when complexity is reduced (extract helpers / table helpers /
  early continue — outcome-focused), then each reports cognitive complexity **≤ 15** under Sonar
  (or golangci equivalent if configured).
- Given the nested anonymous structs inside `type crd`, when refactored, then each flagged nesting
  level uses a **named type** (exported not required) with the same JSON tags and behaviour.
- Given existing goldens under `config/testdata/crd-contract/`, when `go test ./config/ -count=1`
  runs without `UPDATE_GOLDEN=1`, then all golden + structural invariants still pass (no silent
  contract relaxation).
- **Edge:** given `UPDATE_GOLDEN=1`, when regenerating, then goldens remain byte-stable for an
  unchanged CRD tree (refactor touches tests only).
- **Edge (fuzz):** given empty / non-string / missing `id` cases already in the fuzzer, when
  complexity is reduced, then the stub contract (never error; empty id → `""`) still holds.

**Done when**
- [ ] `go test ./config/ -count=1` green; fuzz still builds (`go test ./config/ -run=FuzzGetExternalName -count=1`)
- [ ] Closes the 3× S3776 + 6× S8205 keys above
- [ ] REQ-SONAR-PG-M02 (complexity) + REQ-SONAR-PG-M03 (named types) in OpenSpec
- [ ] Gates green (`make test` / reviewable as applicable)

**Touches:** `config/crd_contract_test.go`, `config/external_name_fuzz_test.go`  
**Depends on:** none · **Parallel-safe-with:** E5-S18, E5-S20, E5-S21  
**Not in scope:** changing CRD goldens for product reasons; production `config/*.go`; coverage CI

### E5-S20 — Shell hygiene: case default, stderr, `[[`, merge nested ifs (P0/P1)
**As a** maintainer **I want** hand-authored `hack/` scripts and meta tests to satisfy Sonar shell
rules **so that** CRITICAL/MAJOR Maintainability and Reliability smells clear without changing
gate pass/fail semantics.

**Sonar keys / rules (grouped)**
1. **shelldre:S131 CRITICAL** — `hack/check-docs.sh:17` `AZ9oaSNAnB7bKDM9yqGC` (case missing `*)`)
2. **shelldre:S7677 MAJOR** — error messages on stdout → stderr (`>&2`):  
   - `hack/check-api-docs.sh:22` `AZ9rwTiXTAot3Azg_Ks5`  
   - `hack/check-docs.sh:33,37` `AZ9oaSNAnB7bKDM9yqGD`, `AZ9oaSNAnB7bKDM9yqGE`
3. **shelldre:S7688 MAJOR RELIABILITY** — prefer `[[` over `[`:  
   - `hack/check-docs.sh:40`  
   - `hack/test/e8_s01_observe_docs_test.sh:3`  
   - `hack/test/e8_s01_observe_yaml_count_test.sh:5,6`  
   - `hack/test/e8_s02_backuplist_crd_exists_test.sh:5,6`  
   - `hack/test/e8_s03_publicnetwork_crd_exists_test.sh:5,6`
4. **shelldre:S1066 MAJOR** — merge nested `if`s:  
   - `hack/test/sonar_pg_02_no_secret_in_run_test.sh:20,38`  
   - `hack/test/sonar_pg_04_ci_job_scoped_permissions_test.sh:24`  
   - `hack/test/sonar_pg_04_e2e_job_scoped_permissions_test.sh:25`

**Acceptance criteria**
- Given `hack/check-docs.sh`’s `case`, when a CRD basename matches none of the skip patterns, then a
  `*)` default arm handles it (continue / append — same as today’s implicit fall-through) so S131
  clears.
- Given `::error::` / failure echoes in `check-docs.sh` and `check-api-docs.sh`, when printed, then
  they go to **stderr** (`>&2`).
- Given the listed `[ … ]` tests, when updated, then they use bash `[[ … ]]` with equivalent
  predicates (counts, `-f`, `-eq`).
- Given nested `if` pairs in the three `sonar_pg_0{2,4}_*` scripts, when merged, then the combined
  condition preserves the **same fail messages and fail conditions** (no weaker security grep).
- **Edge:** given `check-docs.sh` with a stale README, when run, then exit non-zero and error text
  still appears on stderr; success path still prints the sync confirmation.
- **Edge:** given e8 meta tests after `[[` migration, when CRDs/docs are present, then scripts still
  exit 0; when a path is missing, then still exit 1 with FAIL text.

**Done when**
- [ ] `bash hack/check-docs.sh` behaviour preserved (exit code + stderr on failure)
- [ ] `bash hack/test/e8_s01_*.sh` `e8_s02_*.sh` `e8_s03_*.sh` green on a tip that has those files
- [ ] `bash hack/test/sonar_pg_02_*.sh` `sonar_pg_04_ci_*.sh` `sonar_pg_04_e2e_*.sh` green
- [ ] Meta guards for S131/S7677/S7688/S1066 patterns as needed (`hack/test/sonar_pg_m04_*` …)
- [ ] Closes all keys in the four clusters above · REQ-SONAR-PG-M04…M07 in OpenSpec

**Touches:** `hack/check-docs.sh`, `hack/check-api-docs.sh`, `hack/test/e8_s01_*`,
`hack/test/e8_s02_*`, `hack/test/e8_s03_*`, `hack/test/sonar_pg_02_*`,
`hack/test/sonar_pg_04_ci_*`, `hack/test/sonar_pg_04_e2e_*`, `hack/test/sonar_pg_m0{4,5,6,7}_*`  
**Depends on:** none (rebase onto tip with E8 scripts) · **Parallel-safe-with:** E5-S18, S19, S21  
**Not in scope:** rewriting check-docs counting logic; Dockerfile; Go tests; coverage CI

### E5-S21 — Dockerfile: quote expansions (P1)
**As a** release engineer **I want** Dockerfile `RUN`/`ADD`/`ENV` expansions quoted per docker:S6570
**so that** 19 MAJOR Maintainability issues clear without breaking the terraform/plugin install layer.

**Sonar keys / rule**
- `cluster/images/provider-gridscale/Dockerfile:27–39` — docker:S6570 ×19
  keys `AZ-GeyaXbiMwwowBPDZ4` … `AZ-GeyaXbiMwwowBPDaK`

**Acceptance criteria**
- Given lines that expand `${PLUGIN_DIR}`, `${TERRAFORM_*}`, `${TARGETOS}`, `${TARGETARCH}`,
  `${USER_ID}`, zip paths, etc. in the Setup Terraform block (≈27–39), when fixed, then each
  flagged expansion is properly quoted per S6570.
- Given the E5-S17 contract, when this lane finishes, then local artefacts still use **`COPY`**
  (not `ADD`) for `bin/…/provider` and `terraformrc.hcl`; remote URL `ADD` may remain.
- **Edge:** given a build with the quoted Dockerfile, when `docker build` / CI image build runs,
  then terraform + provider plugin still land under `${PLUGIN_DIR}` and the non-root `USER` still
  applies (no path-split regressions from quoting).

**Done when**
- [ ] Meta test (extend `sonar_pg_08_*` or add `sonar_pg_m08_dockerfile_quoting_test.sh`) asserts
  quoted forms for the flagged lines · **Level:** M
- [ ] Existing `bash hack/test/sonar_pg_08_dockerfile_copy_test.sh` still green
- [ ] Closes the 19× S6570 keys · REQ-SONAR-PG-M08 in OpenSpec

**Touches:** `cluster/images/provider-gridscale/Dockerfile`, `hack/test/sonar_pg_m08_*`
(and/or `hack/test/sonar_pg_08_*` if extended in-lane)  
**Depends on:** none · **Parallel-safe-with:** E5-S18, S19, S20  
**Not in scope:** base-image bumps; TERRAFORM_VERSION pin changes; coverage CI

### First slice recommendation

Start **E5-S18 (L-REG)** first — two CRITICAL fixes, tiny diff, zero product behaviour risk — then
fan out **L-CFGTEST + L-SHELL + L-DOCKER** in parallel under `/agent-loop-local`. If only one
implementer: S18 → S19 (remaining CRITICAL complexity) → S20 → S21.
