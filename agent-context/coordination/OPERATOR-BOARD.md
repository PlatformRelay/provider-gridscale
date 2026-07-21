# OPERATOR-BOARD — provider-gridscale (lanes + locks)

Coordinator: session wrap **2026-07-21** (v0.2.1 security patch + docs). Base: `origin/main` @ `c53fc82`.
**No lanes In-flight.** **v0.2.1 published & signed** — see [`SESSION-HANDOFF.md`](SESSION-HANDOFF.md).

### Batch 10 — vuln + README/docs + logo candidates (2026-07-21) — landed on `main`

| Lane | Story | State |
| --- | --- | --- |
| L-DOCS-BRAND | x/text GO-2026-5970; Scorecard permissions; kollect badges; README/docs; logo candidates; BRAND-2 INBOX | ✅ Merged PR #15 → tag **v0.2.1** published ([run](https://github.com/PlatformRelay/provider-gridscale/actions/runs/29868142402)) |

### Batch 9 — badges + v0.2.0 release (2026-07-17c) — landed on `main`

| Lane | Story | State |
| --- | --- | --- |
| L-BADGES | OpenSSF Scorecard badge; drop Go Report; GitHub Releases for release badge | ✅ Integrated (`1799637`) |
| L-REL020 | Tag + GitHub Release v0.2.0; dispatch signed publish | ✅ Published ([run](https://github.com/PlatformRelay/provider-gridscale/actions/runs/29574409336)) |
| L-VERIFY | D-020-FU verify accepts `up` v0.51 per-extension annotations | ✅ Integrated (`43294ed`) |

### Batch 8 — audit-mitigation + upstream fixes (2026-07-17) — landed on `main`

| Lane | Story | State |
| --- | --- | --- |
| L-DISPO | Audit re-run dispositions (DOC-5 CI-wired; DIR-2c/DOC-2 ACCEPTED) + ROADMAP refresh | ✅ Integrated (`ca9d3aa`) |
| L-SECHARD | SEC-H1/2/3 least-privilege `GITHUB_TOKEN` (ci/e2e) + keyword-gate | ✅ Integrated (`68f8d66`) |
| L-DECIDE | Apply operator D-018/019/020 (workspace relay) + track sweep findings | ✅ Integrated (`949d1de`) |
| L-U1 | **U-1** credential-sensitivity fix (`config/sensitive.go`) + regen | ✅ Integrated (`8ae7376`+`9e6ca3c` docs) |
| L-LB1 | **LB-1** loadbalancer `status` Computed (`config/loadbalancer.go`) + regen | ✅ Integrated (`c38e52b`) |
| Upstream PRs | #509 (U-1) / #510 (LB-1) / #511 (objectstorage) on gridscale:master | 🔵 Opened — awaiting gridscale review |

Lanes own **disjoint paths**; auto-merge only on gates-green + `/tech-review` APPROVE + green CI
(security/API/release lanes are **surfaced, not auto-merged**). Merge model: PRs into `main`
(`protect-main` ruleset; solo via `gh pr merge --rebase --admin`).

### Batch 1 — landed on `main`

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| L1-CRED | CRED-1/2 + DOC-1 — wire credentials (`uuid`/`token`), fix secret templates + ProviderConfig group | `internal/clients/**`, `examples/**/providerconfig/**` | **NO — security, surfaced** | ✅ Integrated (`b16b5a2`; **live-validated 2026-07-15** — real VM create/delete) |
| L2-TEST | TEST-1 — real config unit tests (external-name + kind overrides); give the vacuous `make test` gate teeth | `config/*_test.go`, root `Makefile` (test wiring only, if needed) | yes | ✅ Integrated (`c53fded`) |
| L3-README | DOC-3 — rewrite README for gridscale (install + ProviderConfig/creds setup) | `README.md` | yes | ✅ Integrated |
| L0-AUDIT | Land baseline audit + tech-debt register + this board (docs only) | `agent-context/**` | yes | ✅ Integrated |

### Governance/docs/test batch — landed on `main`

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| E6-S01 | `CONTRIBUTING.md` | `CONTRIBUTING.md` | yes | ✅ Integrated |
| E6-S02 | `SECURITY.md` + `.gitignore` | `SECURITY.md`, `.gitignore` | yes | ✅ Integrated |
| E6-S04 | CoC → Contributor Covenant 2.1 + issue/PR templates | `CODE_OF_CONDUCT.md`, `.github/*_TEMPLATE*` | yes | ✅ Integrated |
| E4-S03 | ADRs 0003/0004 | `docs/adr/**` | yes | ✅ Integrated |
| E2-S02 | CRD golden-contract tests | `test/**`, `config/*_test.go` | yes | ✅ Integrated |

### Batch 2 — test-hardening tooling (ported from kollect)

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| E2-S06…S10 | `make vuln` / arch-lint / race / fuzz / tidy-check | `Makefile`, `config/*_test.go`, `.go-arch-lint.yml` | yes (local targets) | ✅ Integrated |
| SEC-govuln | Go 1.26.5 + `x/net` v0.55.0 | `go.mod`/`go.sum` | surfaced | ✅ Landed (`d75721e`); **D-007** |

### Batch 3 — docs

| Lane | Story | State |
| --- | --- | --- |
| E3-S02b | README enrichment — resource matrix + community links | ✅ Integrated |

### Batch 4 — audit gap-stories (`/agent-loop-auto` 2026-07-16)

IDs renumbered per **D-014** (gap E2-S06/07/08 → **E7-S01/02/03**, research E2-S09 → **E2-S11**).

| Lane | Story | State |
| --- | --- | --- |
| L-REFS | **E7-S01** cross-resource `config.Reference{}` | ✅ Integrated (`9ebca1d`+`b8fe841`) |
| L-LINT | **E5-S07** gosec G101/G104 scoping | ✅ Integrated (`4438b32`) |
| L-CI | **E5-S08 + E5-S10** Go-version pin + schema-version-diff | ✅ Integrated (`c44e178`+`a58b174`) |
| L-CREDGUARD | **E2-S11** non-live credential-wiring tests | ✅ Integrated (`be73bd8`) |
| L-EXTNAME | **E7-S03** external-name rationale + import-format docs | ✅ Integrated (`921ac95`) |
| L-E8SCOPE | **E8** scope → document-omit + track upstream (**D-015**) | ✅ Integrated (`6622354`) |

### Batch 5 — after batch 4

| Lane | Story | State |
| --- | --- | --- |
| L-META | **E7-S02** metadata key tests | ✅ Integrated |
| L-DEPREC | **E4-S05** deprecation hygiene | ✅ Integrated |
| L-UPSTREAM | **E6-S06** upstream TF triage (**D-016**) | ✅ Integrated |

### Batch 6 — goimports (alone)

| Lane | Story | State |
| --- | --- | --- |
| L-GOIMPORTS | **E5-S09** goimports local-prefix | ✅ Integrated (`1e9cfe1`) |

### Batch 7 — demo polish (2026-07-16 late)

| Lane | Story | State |
| --- | --- | --- |
| L-ICON | Marketplace icon (operator interrupt) | ✅ Integrated (`64fc363`) + appended on `v0.1.1` |
| L-UNAFFIL | Unaffiliation disclaimer | ✅ Integrated (`b5acc41`); listing via **v0.1.1** |
| L-README | **E3-S04** badges + **E8** limitations | ✅ Integrated (`c15a328`…`bcc0ba0`) |
| L-DOCSYNC | **E4-S01** `check-api-docs` sync gate | ✅ Integrated (`58f1103`+`2384fe4`) |
| L-EXAMPLES | **E4-S02** curated examples index | ✅ Integrated (`fdbf300`) |
| L-ASSURE | **E6-S05** stretch assurance docs | ✅ Integrated (`3edc5f0`) |

### Batch 8 — CI restore (2026-07-16 wrap)

| Lane | Story | State |
| --- | --- | --- |
| L-CIRED | Restore green `main` CI — filesystem CRD generate drift + `hack/metadatafix` lint + `docs/api` sync | ✅ Integrated (`d8433f5`+`d275897`) |

> **Session wrap complete** — prefer Marketplace/install tag **v0.1.1**.
> Primary CI green on `d8433f5`; Docs Sync green on `d275897`. See SESSION-HANDOFF.

## Deferred / blocked (not board lanes)

- **E2-S04/S05 uptest** — **D-012 → B** (manual smoke only; do not wire CI lab creds).
- **E8 datasource codegen** — unsupported in upjet today (**D-015**); document-only done; FR do-not-auto-file.
- **Stale ROADMAP/BACKLOG status prose** — refresh when convenient (docs only).
- Optional: delete superseded remote branches; close upstream TF #188; nudge #467/#468.
