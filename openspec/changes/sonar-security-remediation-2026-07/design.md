# Design — sonar-security-remediation-2026-07

## Dispositions

Every open SonarCloud SECURITY finding on `PlatformRelay_provider-gridscale`
(scan 2026-07-25) is dispositioned in `proposal.md` as **FIX** or **SAFE→harden**.
No DEFER. PG-03 (`md5sum` for a status-check key) is not a crypto misuse, but we
still switch to SHA-256 so the CRITICAL clears without arguing with the analyzer.

## Lane split

Findings span disjoint files, so remediation is seven backlog stories (E5-S11…S17)
with exclusive file locks. Multi-site REQs keep one REQ id but **two Test/Verify
artifacts** so parallel lanes do not share a script:

| REQ | Sites | Meta-test artifacts |
| --- | --- | --- |
| PG-04 | `e2e.yaml` · `ci.yml` | `sonar_pg_04_e2e_*` · `sonar_pg_04_ci_*` |
| PG-06 | publish · gitleaks | `sonar_pg_06_publish_*` · `sonar_pg_06_gitleaks_*` |

Filenames in `specs/security/spec.md` match `agent-context/BACKLOG.md` locks.

## Why meta-tests grep workflows

These are structural / YAML / Dockerfile fixes, not runtime Go behavior. Failing
`hack/test/sonar_pg_*.sh` scripts (and `scripts/version_diff_test.py`) assert the
unsafe pattern is gone before a lane claims Verify. Same contract as kaddy
OpenSpec security baselines: Test first, then change the locked file.
