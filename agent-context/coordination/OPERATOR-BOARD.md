# OPERATOR-BOARD — provider-gridscale (lanes + locks)

Coordinator: session wrap 2026-07-16 (CI-red fix landed). Base: `origin/main` @ `d275897`.
**No lanes In-flight.** Auto-mergeable backlog exhausted. See
[`SESSION-HANDOFF.md`](SESSION-HANDOFF.md) for next entry points.

Lanes own **disjoint paths**; auto-merge only on gates-green + `/tech-review` APPROVE + green CI
(security/API/release lanes are **surfaced, not auto-merged**). Merge model this session: local
ff-merge → push `main` (no PRs).

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

> **Session wrap complete** — HEAD `d275897`. Prefer Marketplace/install tag **v0.1.1**.
> Primary CI green on `d8433f5`; Docs Sync green on `d275897`. See SESSION-HANDOFF.

## Deferred / blocked (not board lanes)

- **E2-S04/S05 uptest** — **D-012 → B** (manual smoke only; do not wire CI lab creds).
- **E8 datasource codegen** — unsupported in upjet today (**D-015**); document-only done; FR do-not-auto-file.
- **Stale ROADMAP/BACKLOG status prose** — refresh when convenient (docs only).
- Optional: delete superseded remote branches; close upstream TF #188; nudge #467/#468.
