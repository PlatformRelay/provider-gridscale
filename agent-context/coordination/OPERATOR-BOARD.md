# OPERATOR-BOARD — provider-gridscale (lanes + locks)

Coordinator: `/agent-loop-auto` session 2026-07-15. Base: `origin/main` @ `00e1c6c` (green).
Lanes own **disjoint paths**; auto-merge only on gates-green + `/tech-review` APPROVE + green CI +
MERGEABLE (security/API/release lanes are **surfaced, not auto-merged**).

### Batch 1 — landed on `main`

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| L1-CRED | CRED-1/2 + DOC-1 — wire credentials (`uuid`/`token`), fix secret templates + ProviderConfig group | `internal/clients/**`, `examples/**/providerconfig/**` | **NO — security, surfaced** | ✅ Integrated (`b16b5a2`; **live-validated 2026-07-15** — real VM create/delete) |
| L2-TEST | TEST-1 — real config unit tests (external-name + kind overrides); give the vacuous `make test` gate teeth | `config/*_test.go`, root `Makefile` (test wiring only, if needed) | yes | ✅ Integrated (`c53fded`) |
| L3-README | DOC-3 — rewrite README for gridscale (install + ProviderConfig/creds setup) | `README.md` | yes | ✅ Integrated |
| L0-AUDIT | Land baseline audit + tech-debt register + this board (docs only) | `agent-context/**` | yes | ✅ Integrated |

### Governance/docs/test batch — just merged to `main`

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| E6-S01 | `CONTRIBUTING.md` — standards map, merge policy, preflight, commit convention | `CONTRIBUTING.md` | yes | ✅ Integrated |
| E6-S02 | `SECURITY.md` + `.gitignore` | `SECURITY.md`, `.gitignore` | yes | ✅ Integrated |
| E6-S04 | CoC → Contributor Covenant 2.1 + issue/PR templates | `CODE_OF_CONDUCT.md`, `.github/*_TEMPLATE*` | yes | ✅ Integrated |
| E4-S03 | ADRs 0003/0004 (foundational decisions) | `docs/adr/**` | yes | ✅ Integrated |
| E2-S02 | CRD golden-contract tests | `test/**`, `config/*_test.go` | yes | ✅ Integrated |

### Batch 2 — test-hardening tooling (ported from kollect)

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| E2-S06…S10 | `make vuln` / `.go-arch-lint.yml`+`make arch-lint` / `make test.race` / `FuzzGetExternalName` / `make tidy-check` | `Makefile`, `config/*_test.go`, `.go-arch-lint.yml` | yes (local targets/tests only) | ✅ Integrated |
| SEC-govuln | Go 1.26.5 + `x/net` v0.55.0 (clears 3 called CVEs) | `go.mod`/`go.sum` | surfaced | ✅ Landed (`d75721e`); reconciled → decisions.md **D-007** |

### Batch 3 — final auto-mergeable docs lane

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| E3-S02b | README enrichment — 32-resource matrix (correct served `names.kind` casing) + community/contributing links | `README.md` | yes | ✅ Integrated — `/tech-review` REQUEST-CHANGES (P1: Kind casing) → **fixed** → re-verified, gates green |

> **Loop stopped here (2026-07-16).** Auto-mergeable backlog **exhausted**. All remaining work is
> operator-gated — surfaced as **D-008…D-012** in [`INBOX.md`](../INBOX.md): maintainer identity
> (D-008), real branding (D-009), Marketplace metadata (D-010), the E5 CI/supply-chain epic (D-011),
> and uptest lab creds (D-012). E3-S02's badges + logo halves are deferred inside D-009/D-011.

## Collision notes
- L1 and L3 both reference credential setup but touch **different files** — no path collision.
- `.golangci.yml` goimports-prefix fix (DIR-2) is deferred: it reformats imports repo-wide and would
  collide with every Go lane — run it **alone** in a later batch.
- Go-version-drift + CI-file edits (DIR-2 remainder) are CI/release-adjacent → surface, don't auto-merge.

## Deferred / surface-for-operator
- DOC-4 (real maintainer in CODEOWNERS/OWNERS.md) — needs an operator decision (who owns this).
- **CRED e2e now live-validated (2026-07-15):** an out-of-cluster smoke test authenticated to the live
  gridscale API and did a real Server create/delete (verified gone via API). This was a **single manual
  smoke test** — automated *uptest* coverage (E2-S04/S05) is still pending, gated on lab creds + CI.
- Wiring the new E2-S06…S10 quality tools (govulncheck / go-arch-lint / race / fuzz / tidy-check) into
  CI (`.github/workflows/`) — **deferred to Epic E5** (surfaced, not auto-merged). govulncheck job = E5-S01.
- **`govulncheck` findings (2026-07-16):** the first scan surfaced **3 called vulnerabilities** —
  GO-2026-5856 (`crypto/tls` ECH leak; fix go1.26.5), GO-2026-5026 + GO-2026-4918
  (`golang.org/x/net` idna Punycode + HTTP/2 loop; fix v0.55.0). Remediation (Go 1.26.5 + `x/net`
  bump) is a **surfaced operator decision** — see [`INBOX.md`](../INBOX.md).
