<!--
Thanks for contributing to provider-gridscale!

This is a generated Upjet provider. Please read the engineering rules in
agent-context/GUIDELINES.md before opening a PR — especially §1 (generation
boundary) and §8 (Definition of Done). Find the Crossplane community in
https://slack.crossplane.io/messages/dev if you need help.
-->

### Description of your changes

<!--
Briefly describe what this pull request does and why. Direct reviewers'
attention to anything that needs special consideration.
-->

Fixes #

### How has this code been tested

<!--
Describe the testing done or planned. For hand-authored Go/hack, note the
failing-test-first (TDD) artifact. For resources, note whether an uptest (E)
example exercised the change.
-->

### Definition of Done checklist

See `agent-context/GUIDELINES.md` §8.

- [ ] **Test-first**: for hand-authored Go / `hack/` scripts, a failing test was written before the fix (TDD; §2). Generated code is exempt but stays green.
- [ ] **Generation boundary respected** (§1): no hand-edited `apis/**/zz_*.go`, `internal/controller/**`, or `package/crds/**`. Behavior changes went through `config/*` inputs + `make generate`.
- [ ] `make reviewable` is clean.
- [ ] `make check-diff` is clean — the regenerated tree is committed in this same change (no drift).
- [ ] REQ ↔ Test ↔ Verify are complete for any spec'd requirement; `make test.spec` passes.
- [ ] Examples/docs regenerated where the resource surface changed (`make docs`); `examples/` still apply cleanly.
- [ ] An ADR (`docs/adr/`) + `decisions.md` entry was added for any non-obvious decision (§7).
- [ ] One logical change per commit, tree green at each; commits follow the gitmoji-conventional format (e.g. `:sparkles: feat(config): ...`).
