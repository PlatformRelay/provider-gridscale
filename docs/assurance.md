# Assurance case — provider-gridscale (E6-S05 stretch)

Lightweight, provider-scoped assurance notes aligned with the OpenSSF Scorecard
baseline. This is **not** a formal certification dossier; it records where the
repo already meets common supply-chain and community-health expectations, and
what remains operator-gated.

## Claims and evidence

| Claim | Evidence in this repo |
| --- | --- |
| Changes land through reviewed, linear history | Ruleset **`protect-main`** (1 approval + required checks + Admin PR-only bypass for solo); [`CONTRIBUTING.md`](../CONTRIBUTING.md) rebase-and-merge; [`GOVERNANCE.md`](../GOVERNANCE.md); [`AGENTS.md`](../AGENTS.md) |
| Maintainers are identifiable | [`OWNERS.md`](../OWNERS.md), [`CODEOWNERS`](../CODEOWNERS) |
| Vulnerability reporting path exists | [`SECURITY.md`](../SECURITY.md) |
| Code of Conduct is a recognised standard | [`CODE_OF_CONDUCT.md`](../CODE_OF_CONDUCT.md) (Contributor Covenant 2.1) |
| CI exercises build, lint, unit tests, and diff guards | [`.github/workflows/ci.yml`](../.github/workflows/ci.yml) |
| Dependency and Action updates are automated | [`.github/renovate.json5`](../.github/renovate.json5) |
| Secrets scanning on push/PR | [`.github/workflows/gitleaks.yml`](../.github/workflows/gitleaks.yml), [`.pre-commit-config.yaml`](../.pre-commit-config.yaml) |
| Vulnerability scanning of Go deps | [`.github/workflows/govulncheck.yml`](../.github/workflows/govulncheck.yml) (`make vuln`) |
| Static analysis | [`.github/workflows/codeql.yml`](../.github/workflows/codeql.yml), OpenSSF Scorecard workflow |
| Coverage floor on the hand-authored surface | [`codecov.yml`](../codecov.yml), `make coverage` (config + clients) |
| Released packages are signed with an SBOM | [`.github/workflows/publish-provider-package.yml`](../.github/workflows/publish-provider-package.yml) (keyless cosign + SPDX attest) |
| Generated API docs stay fresh | `make check-api-docs` + docs-sync workflow |

## SCA remediation posture

1. **Detect** — weekly `govulncheck` + Renovate PRs for Go modules and Actions pins.
2. **Triage** — called (reachable) findings block merge via `make vuln` / the Govulncheck job; uncalled findings are tracked but do not soft-fail the release.
3. **Remediate** — bump the pinned Go toolchain / module versions in `go.mod` and the Makefile/`ci.yml` pins together (`make go-version-check`).
4. **Verify** — re-run `make vuln`, unit coverage, and CI on `main` before tagging.

## Explicit non-claims

- **No live uptest in CI** — lab credentials stay operator-local (D-012); examples are curated for manual / `/test-examples` runs.
- **No Terraform datasource CRDs** — documented under README “Known limitations” (D-015).
- **No SonarCloud / OpenVEX** unless D-004 is revisited.

## OpenSSF Scorecard — known grade drivers

The README Scorecard badge tracks
[securityscorecards.dev](https://securityscorecards.dev/viewer/?uri=github.com/PlatformRelay/provider-gridscale).
As of the 2026-07-20 scan the overall grade was **6.7**. Drivers that are
*not* a red CI job, but still pull the badge into the amber band:

| Check | Why it is low | What moves it |
| --- | --- | --- |
| Maintained (0) | Repo younger than 90 days | Time + continued commits/issues |
| Code-Review (0) | Scorecard saw 0/27 human-approved merges (solo admin merges) | Land changes via reviewed PRs with a non-author approval |
| Token-Permissions | Some workflows lacked top-level `permissions:` | Declared deny-by-default roots on backport / tag / publish (this PR) |
| Packaging (−1) | Heuristic missed `workflow_dispatch`-only GHCR publish | Optional: also publish on `release` published events |
| Signed-Releases (−1) | Cosign signs the **xpkg/GHCR** digest, not GitHub Release assets | Optional: attach cosign bundles / checksums to GitHub Releases |
| CII-Best-Practices (0) | No OpenSSF Best Practices badge yet | Operator registers the project at bestpractices.dev |
| Vulnerabilities | OSV findings in the dependency graph | Keep `make vuln` / Govulncheck green (e.g. `golang.org/x/text` bump) |

Re-run Scorecard after the next scheduled workflow (or `workflow_dispatch`) to
pick up permission and vuln fixes.

## Review cadence

Re-read this page when cutting a minor release or after a Scorecard grade change.
Update the evidence table if a workflow is renamed or a gate is dropped.
