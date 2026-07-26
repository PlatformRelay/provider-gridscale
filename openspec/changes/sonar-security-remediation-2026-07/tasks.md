# Tasks — sonar-security-remediation-2026-07 (provider-gridscale)

Backlog stories: **E5-S11…E5-S17** (`agent-context/BACKLOG.md`). Claim one story per worktree.

| Story | REQ | File lock | Meta-test locks | Status |
| --- | --- | --- | --- | --- |
| E5-S11 | PG-01 + PG-06(publish) | `publish-provider-package.yml` | `sonar_pg_01_*`, `sonar_pg_06_publish_*` | ✅ |
| E5-S12 | PG-02 + PG-03 + PG-04(e2e) | `e2e.yaml` | `sonar_pg_02_*`, `sonar_pg_03_*`, `sonar_pg_04_e2e_*` | ✅ |
| E5-S13 | PG-04(ci) | `ci.yml` | `sonar_pg_04_ci_*` | ✅ |
| E5-S14 | PG-05 | `scorecard.yml` | `sonar_pg_05_*` | ✅ |
| E5-S15 | PG-06(gitleaks) | `gitleaks.yml` | `sonar_pg_06_gitleaks_*` | ✅ |
| E5-S16 | PG-07 | `scripts/version_diff.py` | `scripts/version_diff_test.py` | ✅ |
| E5-S17 | PG-08 | Dockerfile | `sonar_pg_08_*` | ✅ |

**Parallel batch:** E5-S11 + E5-S14 + E5-S16 + E5-S17 · **Solo first:** E5-S11

Per-story: failing **Test:** first → implement → **Verify:** → tech-review. Details in backlog.
PG-04 / PG-06 keep one REQ each but **split Test/Verify scripts** per site (see spec).

**Batch complete 2026-07-25** — all seven stories Integrated on `main`; meta tests `hack/test/sonar_pg_*` green (9/9). Change is implementation-complete and archive-ready.
