# Change: sonar-security-remediation-2026-07 — SonarCloud SECURITY findings

## Story map (parallel sessions)

Backlog IDs (claim one per worktree — see `agent-context/BACKLOG.md` / OPERATOR-BOARD):

| Story | REQs |
| --- | --- |
| E5-S11 | PG-01, PG-06 (publish) |
| E5-S12 | PG-02, PG-03, PG-04 (e2e) |
| E5-S13 | PG-04 (ci) |
| E5-S14 | PG-05 |
| E5-S15 | PG-06 (gitleaks) |
| E5-S16 | PG-07 |
| E5-S17 | PG-08 |

## Why

SonarCloud project `PlatformRelay_provider-gridscale` reports **13 open SECURITY
issues** (2 BLOCKER · 1 CRITICAL · 10 MAJOR). Zero Security Hotspots `TO_REVIEW`.
This change is the authoritative, carefully analyzed remediation backlog: every
finding is dispositioned **FIX / SAFE / DEFER**, with a REQ that carries
**Test:** + **Verify:** so implementation lanes cannot claim done without a
checkable gate (same contract as kaddy OpenSpec).

Source scan: SonarCloud main analysis, 2026-07-25 (issue keys cited in the spec).

## Scope

- Disposition all 13 SECURITY issues.
- Spec REQs for every **FIX** (and for **SAFE** items that need an explicit
  Sonar accept / NOSONAR rationale in-tree).
- Meta tests under `hack/test/` that grep workflows/scripts for the unsafe
  patterns (TDD: failing test first). Multi-site REQs PG-04 and PG-06 use
  **per-lane Test/Verify scripts** aligned with BACKLOG file locks
  (`sonar_pg_04_{e2e,ci}_*`, `sonar_pg_06_{publish,gitleaks}_*`).

## Non-goals

- Closing Sonar issues in the UI before the matching code lands.
- Regenerating Upjet `zz_*.go` / CRDs.
- Expanding Scorecard beyond the concrete permission findings.

## Analysis summary (operator-facing)

| ID | Rule | Sev | File | Disposition | Why |
|----|------|-----|------|-------------|-----|
| PG-01 | `githubactions:S7630` | BLOCKER ×2 | `publish-provider-package.yml` | **FIX** | `${{ inputs.version }}` expanded inside `run:` via `format()` — classic script injection |
| PG-02 | `githubactions:S7636` | MAJOR | `e2e.yaml:166` | **FIX** | `echo "${{ secrets.UPTEST_DATASOURCE }}"` expands a secret in the script body (leak + injection surface) |
| PG-03 | `githubactions:S4790` | CRITICAL | `e2e.yaml:88` | **SAFE→harden** | `md5sum` hashes an example-list string for a status-check context key — **not** a crypto/password use. Still switch to `sha256sum` so Sonar clears and the key stays collision-resistant |
| PG-04 | `githubactions:S8264` | MAJOR ×2 | `ci.yml`, `e2e.yaml` | **FIX** | Workflow-level `contents: read` → move to per-job `permissions` |
| PG-05 | `githubactions:S8234` | MAJOR | `scorecard.yml` | **FIX** | Replace `permissions: read-all` with explicit minimal scopes |
| PG-06 | `githubactions:S6506` | MAJOR ×2 | publish + gitleaks workflows | **FIX** | `curl` without `--proto '=https' --tlsv1.2` (URLs are already https; enforce protocol) |
| PG-07 | `pythonsecurity:S8707` | MAJOR ×3 | `scripts/version_diff.py` | **FIX** | CLI path args opened without bounding to the repo root — LLM/tooling risk class Sonar flags; validate with `Path.resolve()` under `ROOT` |
| PG-08 | `docker:S7029` | MAJOR | Dockerfile | **FIX** | Local `ADD` of `bin/.../provider` and `terraformrc.hcl` → `COPY`; keep remote `ADD` for Hashicorp/provider zips (or replace with `curl`) |

## Counterpoints considered

- *"Accept the BLOCKER — only maintainers can `workflow_dispatch`."* Rejected: any
  collaborator with `workflow_dispatch` on that workflow can inject via `version`;
  env-var indirection is a one-line fix.
- *"Exclude `scripts/` from Sonar."* Rejected: path validation is cheap and correct.
- *"Leave md5 — it's not crypto."* Partially accepted as risk assessment; still
  harden to sha256 so the CRITICAL clears without arguing with the analyzer.

## Links

- SonarCloud: https://sonarcloud.io/project/issues?id=PlatformRelay_provider-gridscale&impactSoftwareQualities=SECURITY
- Sibling pattern: kaddy `openspec/changes/e1c-security-baseline/` (Verify + Test per REQ)
