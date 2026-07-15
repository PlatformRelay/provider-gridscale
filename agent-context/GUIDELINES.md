# GUIDELINES — provider-gridscale (binding)

Engineering rules for this repo, adapted from the workspace guidelines for a **generated Upjet
provider**. Workspace-wide rules (commits, merge, adversarial collaboration) live in
[`../../AGENTS.md`](../../AGENTS.md) and are not restated here. A red gate ⇒ not mergeable.

## 1. Respect the generation boundary

The provider is machine-generated. **Never hand-edit** `apis/**/zz_*.go`, `internal/controller/**`,
or `package/crds/**`. To change behaviour, change the *inputs*:
- resource shape / external-name / references → `config/external_name.go`, `config/gridscale.go`
- provider wiring / group → `config/provider.go`
- then run `make generate` and commit the regenerated tree in the *same* logical change.
`make reviewable` and `make check-diff` guard this — a PR that hand-edits generated output fails CI.

## 2. Spec-driven, test-first

- Every implementation lane references an OpenSpec change (`openspec/changes/<slug>/`).
- **TDD is mandatory for hand-authored Go and `hack/` scripts**: write the failing **U**-level test
  first. Generated code is exempt from unit TDD (you'd be testing Upjet), but must stay green and is
  exercised at the **E** (uptest) tier.
- Every REQ carries one `**Test:**` (artifact path) + one `**Verify:**` (command) + a `**Level:**`
  (U/E/M/D). Gate: `make test.spec`. Epic exit: `STRICT_TEST_FILES=1 make test.spec` (all Test paths exist).

## 3. Test taxonomy (see ADR-0002)

| Level | Scope | Runs |
| --- | --- | --- |
| **U** | Go unit — `config/`, `hack/` helpers | every push (`make test`) |
| **E** | uptest example-as-test against the live gridscale API | `/test-examples`, nightly |
| **M** | meta/structural — file/CI/lint/diff invariant | every push |
| **D** | generated docs in sync with source | every push |

Assert on the **small hand-authored surface**. Reserve **E** for requirements that genuinely need a
real API round-trip; keep those examples minimal (cost + rate limits).

## 4. Examples are tests

`examples/` manifests double as uptest fixtures. Keep them valid, minimal, and annotated
(`meta.upbound.io/example-id`, uptest timeout). An example that can't apply cleanly is a defect.

## 5. Error handling & robustness (hand-authored Go)

- Wrap errors with context (`errors.Wrap`/`fmt.Errorf("%w")`); never discard.
- No `panic`/`os.Exit` outside `main`; respect `context` deadlines.
- `hack/` shell: `set -euo pipefail`; redact secrets; deterministic output.

## 6. Security & supply chain

- No secrets in git — gridscale tokens via `ProviderConfig` secretRef only; `.envrc`/lab creds are
  local and gitignored. `gitleaks` + pre-commit to block accidental commits are **planned** (E5-S03).
- **Planned** (E5, not yet shipped): pinned action SHAs; `govulncheck`, CodeQL, OpenSSF Scorecard on
  CI (E5-S01/S02); released XPKG signed with cosign + SBOM (E5-S06).
- Don't bump `TERRAFORM_VERSION` past `1.5.7` (BSL boundary — deliberate pin).

## 7. Documentation & decisions

- Every non-obvious decision → an ADR (`docs/adr/`) + a `decisions.md` entry with counterpoints.
- CRD API reference and examples are **generated** (`make docs`); never hand-maintain API tables.
- Public surface (README, Marketplace metadata) stays accurate to the shipped resource set.

## 8. Definition of Done

Failing test first (for hand-authored code) · `make reviewable` clean · REQ ↔ Test ↔ Verify complete
· no hand-edited generated files · ADR added for any real decision · examples/docs regenerated · one
logical change per commit, tree green at each.
