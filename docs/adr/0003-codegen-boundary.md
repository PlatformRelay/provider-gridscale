# ADR-0003: Treat the provider as machine-generated (the codegen boundary)

**Theme:** 01 · Foundations · **Status:** Accepted · **Date:** 2026-07-15 · **Realises:** [GUIDELINES §1](../../agent-context/GUIDELINES.md) + BACKLOG "Provider-specific hazards"

## Context

`provider-gridscale` is an Upjet provider: the vast majority of the tree is **generated** from the
upstream terraform-provider-gridscale schema, not hand-authored. `apis/**/zz_*.go` (API types),
`internal/controller/**` (reconcilers), and `package/crds/**` (CRDs) are all Upjet output — roughly
95% of the repo. Only a thin `config/` layer and a few `hack/` helpers are ours.

The failure mode is obvious and tempting: a reviewer spots a wrong CRD field, a bad Kind, or a
missing reference and *fixes it in the generated file*. That edit works locally and passes review by
eye — then the next `make generate` silently reverts it, or worse, the tree diverges and the drift
is discovered much later. On a generated provider the generated files are **not** the source of
truth; the codegen *inputs* are. Editing output is editing a build artifact.

This is the single most important invariant for anyone (human or agent) touching this repo, which is
why it is stated as [GUIDELINES §1](../../agent-context/GUIDELINES.md) and carried into every lane
via the [BACKLOG "Provider-specific hazards"](../../agent-context/BACKLOG.md) list. This ADR records
the decision behind those rules.

## Decision

**Treat the generated tree as read-only build output. Route every behavioural change through the
`config/` inputs, then regenerate.**

- **Never hand-edit** `apis/**/zz_*.go`, `internal/controller/**`, or `package/crds/**`.
- To change behaviour, change the *input*:
  - resource shape / external-name / references → [`config/external_name.go`](../../config/external_name.go),
    [`config/gridscale.go`](../../config/gridscale.go)
  - provider wiring / API group → [`config/provider.go`](../../config/provider.go)
- Then run **`make generate`** and commit the regenerated tree in the **same logical change** as the
  config edit — inputs and outputs move together, tree green at each commit.
- The boundary is **CI-enforced**: `make reviewable` and `make check-diff` (both from the crossplane
  `build/` submodule) regenerate and fail the build on any diff, so a PR that hand-edits generated
  output — or forgets to regenerate after a config change — turns the gate red.

**Canonical example — the Kind overrides.** The clearest instance of legitimate config-layer
intervention lives in [`config/gridscale.go`](../../config/gridscale.go): two resources get their Go
Kind overridden because the default derived from the terraform name is invalid Go.

- `mysql8` → **`MySQL8`** — the default would derive Kind `"0"` (an invalid identifier).
- `storage_import` → **`StorageImport`** — the default collides with the Go keyword `import`.

Both are expressed as overrides in `config/`, not as post-hoc edits to the generated `apis/**` or
`package/crds/**`, and both are self-guarding: a broken override fails `make build`/`make generate`
at compile time, and the [CRD golden-contract tests (ADR-0002)](0002-testing-strategy.md) catch any
resulting schema drift. That is the pattern for *every* change: fix the input, regenerate, let the
gate prove the output matches.

## Consequences

- The repo has a crisp, mechanically-checkable contract: generated files are never authored by hand,
  so `make check-diff` staying green *is* the proof of correctness for the whole generated surface.
- Regeneration is safe and idempotent — an Upjet bump or a config edit reproduces the tree
  deterministically; there is no hand-edit to lose or reconcile.
- Reviewers can ignore the ~95% generated diff and focus attention on the small `config/`/`hack/`
  surface, which is where all real decisions and all [TDD (ADR-0002)](0002-testing-strategy.md) live.
- The cost: a fix that *looks* like a one-line edit to a generated file becomes "find the right
  config input, change it, regenerate, review the induced diff." Accepted — it is the only way the
  change survives the next regeneration.

## Counterpoints considered

- *"Allow a documented hand-edit for a field Upjet can't express."* Rejected: it defeats
  `check-diff`, rots on the next regen, and there is almost always a config-layer path (an override,
  a schema patch, or an upstream PR). If Upjet genuinely can't express something, the right move is
  an **upstream terraform-provider-gridscale PR**, not a local edit to generated output.
- *"Drop the `check-diff` gate; trust reviewers to notice hand-edits."* Rejected: hand-edits to
  generated files pass human review precisely because they look like ordinary code. The whole point
  is a machine check a human can't be relied on to perform.
- *"Vendor/fork the generated output and maintain it directly."* Rejected: that abandons Upjet's
  value entirely (free regeneration against upstream schema bumps) for a permanent maintenance
  burden over 30+ resources.
