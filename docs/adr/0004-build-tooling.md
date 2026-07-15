# ADR-0004: Keep the upstream Makefile + `build/` submodule; add thin `hack/` scripts

**Theme:** 01 · Foundations · **Status:** Accepted · **Date:** 2026-07-15 · **Realises:** [D-003](../../agent-context/decisions.md)

> **Approved** by the operator on 2026-07-15 as part of the E1–E6 epic-shape sign-off (D-003, option **A**).

## Context

The sibling app repos (`kaddy`, `kollect`, `mkurator`) drive their build and CI from a **Taskfile**
(`task …`) — a consistent, discoverable runner across those projects. `provider-gridscale`, by
contrast, was scaffolded from the upjet-provider-template, which ships the **crossplane `build/`
submodule** (`makelib/`) driven by a root `Makefile`. That Makefile is where the whole Upjet
toolchain lives: `make generate` (codegen), `make reviewable` / `make check-diff` (the
[codegen-boundary gate, ADR-0003](0003-codegen-boundary.md)), `make build` / `make test`,
`make uptest` / `make e2e`, and the terraform-provider fetch used to build the provider image.

The tension: adopting a Taskfile for cross-repo consistency would mean either **reimplementing** the
submodule's targets (fragile, and it drifts on every `build/` bump) or **wrapping** them (a Taskfile
that only shells out to `make`, which adds a layer without adding value). Either way we would be
fighting the upstream tooling the template — and every other Upjet provider — is built around.

## Decision

**Keep the upstream `Makefile` + crossplane `build/` submodule as the build runner. Do not introduce
a Taskfile. Extend the existing tooling with thin `hack/` scripts and a small set of new make
targets that wrap them.**

- The [`build/` submodule](../../.gitmodules) (`github.com/crossplane/build`, `makelib/`) stays the
  source of the core targets (`generate`, `reviewable`, `check-diff`, `build`, `test`, …). We do not
  fork or vendor it; we track upstream.
- Repo-specific automation lives in [`hack/`](../../hack/) as small, `set -euo pipefail` shell
  scripts, invoked from **new make targets** layered on top of the upstream Makefile — e.g.
  `make test.spec` (structural REQ↔Test↔Verify coverage, per
  [ADR-0002](0002-testing-strategy.md)) and `make docs` (generated CRD reference + examples). These
  targets are additions this ADR calls for; they wrap `hack/` helpers rather than duplicating
  makelib logic.
- **`TERRAFORM_VERSION` stays pinned at `1.5.7`** — a deliberate BSL boundary. Terraform 1.6
  introduced the Business Source License; the Makefile encodes this as a hard guard
  (`must be less than 1.6.0 since that version introduced a not permitted BSL license`) and the
  provider-build fetch uses the pin. **Do not bump it** ([GUIDELINES §6](../../agent-context/GUIDELINES.md)).

## Consequences

- We stay aligned with the Upjet/crossplane ecosystem: `build/` bumps, template updates, and
  upstream docs all apply directly, with no translation layer to maintain.
- New automation has a single, obvious home (`hack/` + a make target) that any contributor familiar
  with the template will recognise — no bespoke runner to learn.
- The BSL pin keeps the shipped provider on a permissively-licensed Terraform; it is guarded in code,
  so an accidental bump fails the build rather than silently crossing the licence boundary.
- The cost: build invocation diverges from the sibling app repos (`make …` here vs `task …` there),
  so cross-repo muscle memory doesn't fully transfer. Accepted — matching Upjet tooling is worth more
  than cross-repo runner uniformity, and the target *names* stay conventional.

## Counterpoints considered

- *"Adopt a Taskfile for consistency with `kaddy`/`kollect`/`mkurator`."* Rejected (the core of
  [D-003](../../agent-context/decisions.md)): it would either reimplement makelib targets (fragile,
  drifts on every submodule bump) or merely wrap `make` (a layer with no benefit). The runner
  differing from the app repos is an accepted, deliberate divergence — it matches upjet tooling and
  avoids fighting upstream.
- *"Fork/vendor the `build/` submodule so all targets live in-repo."* Rejected: it trades free
  upstream maintenance for a permanent burden, the same anti-pattern rejected for generated code in
  [ADR-0003](0003-codegen-boundary.md).
- *"Bump `TERRAFORM_VERSION` to a current release for newer features."* Rejected: 1.6+ is BSL-licensed
  and out of bounds for what we ship; the pin is a licence decision, not a staleness bug.
