# OPERATOR-BOARD — provider-gridscale (lanes + locks)

Coordinator: `/agent-loop-auto` session 2026-07-15. Base: `origin/main` @ `00e1c6c` (green).
Lanes own **disjoint paths**; auto-merge only on gates-green + `/tech-review` APPROVE + green CI +
MERGEABLE (security/API/release lanes are **surfaced, not auto-merged**).

| Lane | Story | Owns (disjoint paths) | Auto-merge? | State |
| --- | --- | --- | --- | --- |
| L1-CRED | CRED-1/2 + DOC-1 — wire credentials (`uuid`/`token`), fix secret templates + ProviderConfig group | `internal/clients/**`, `examples/**/providerconfig/**` | **NO — security, surface** | 🔶 In flight |
| L2-TEST | TEST-1 — real config unit tests (external-name + kind overrides); give the vacuous `make test` gate teeth | `config/*_test.go`, root `Makefile` (test wiring only, if needed) | yes | 🔶 In flight |
| L3-README | DOC-3 — rewrite README for gridscale (install + ProviderConfig/creds setup) | `README.md` | yes | 🔶 In flight |
| L0-AUDIT | Land baseline audit + tech-debt register + this board (docs only) | `agent-context/**` | yes | 🔶 In flight (coordinator-pushed) |

## Collision notes
- L1 and L3 both reference credential setup but touch **different files** — no path collision.
- `.golangci.yml` goimports-prefix fix (DIR-2) is deferred: it reformats imports repo-wide and would
  collide with every Go lane — run it **alone** in a later batch.
- Go-version-drift + CI-file edits (DIR-2 remainder) are CI/release-adjacent → surface, don't auto-merge.

## Deferred / surface-for-operator
- DOC-4 (real maintainer in CODEOWNERS/OWNERS.md) — needs an operator decision (who owns this).
- CRED e2e proof needs live gridscale creds (`/test-examples`) — gates only prove the mapping compiles.
