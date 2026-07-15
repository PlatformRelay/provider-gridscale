# Contributing to provider-gridscale

Thanks for your interest in improving **provider-gridscale** — a [Crossplane](https://crossplane.io)
provider for [gridscale](https://gridscale.io), generated with
[Upjet](https://github.com/crossplane/upjet) from `terraform-provider-gridscale`.

This guide covers how to set up a local environment, the rules that keep a *generated* provider
healthy, and how to get a change merged. It links to the authoritative documents rather than
restating them — please follow those links, they are the source of truth.

- **Engineering rules (binding):** [`agent-context/GUIDELINES.md`](agent-context/GUIDELINES.md)
- **Repo entry point & doc map:** [`AGENTS.md`](AGENTS.md)
- **Design decisions:** [`docs/adr/`](docs/adr/README.md) — start with
  [ADR-0001 (spec-driven workflow)](docs/adr/0001-spec-driven-workflow.md) and
  [ADR-0002 (testing strategy)](docs/adr/0002-testing-strategy.md).

By participating you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md). The project is
licensed under [Apache-2.0](LICENSE); contributions are accepted under the same license.

## Getting started

**Prerequisites**

- Go **1.24** (`GO_REQUIRED_VERSION` in the [`Makefile`](Makefile)).
- Docker (for building the provider image / XPKG).
- `git`, `make`, and a POSIX shell.

**Clone with submodules.** The build tooling lives in the `build/` git submodule (crossplane/build).
Clone recursively, or initialise the submodule after cloning:

```bash
git clone --recursive https://github.com/PlatformRelay/provider-gridscale.git
# or, in an existing checkout:
git submodule update --init --recursive   # (make submodules also does this)
```

The first `make` invocation bootstraps the submodule and tooling automatically, then re-runs.

**Build and test locally:**

```bash
make build     # build the provider binary + XPKG
make test      # Go unit tests (U) with coverage → cover.out
make run       # run the provider out-of-cluster (needs a kubeconfig + ProviderConfig)
```

Run `make help` to list all targets.

## Respect the generation boundary

This is a **thin, generated** provider: the controllers, API types, and CRDs are produced by Upjet.
**Never hand-edit generated output:**

- `apis/**/zz_*.go`
- `internal/controller/**`
- `package/crds/**`

To change behaviour, edit the *inputs* under [`config/`](config/) (external-name strategy, kind/field
overrides, provider wiring), then regenerate:

```bash
make generate      # regenerate the tree from config/ inputs
```

Commit the regenerated tree together with the `config/` change, in the *same* logical commit.
`make reviewable` and `make check-diff` fail any PR that hand-edits generated files. See
[GUIDELINES §1](agent-context/GUIDELINES.md#1-respect-the-generation-boundary) and
[AGENTS.md → "What this repo is"](AGENTS.md).

## Testing expectations

- **TDD is mandatory for hand-authored Go and `hack/` scripts** — write the failing test first.
  Generated code is exempt from unit TDD (that would be testing Upjet), but it must stay green.
- The hand-authored surface worth testing is small — assert on `config/` logic, not generated glue.
- The provider uses a four-tier, provider-shaped taxonomy — **U** (unit) · **M** (meta/CRD-contract) ·
  **E** (uptest example-as-test) · **D** (doc-sync). The rationale is in
  [ADR-0002](docs/adr/0002-testing-strategy.md); the rules are in
  [GUIDELINES §2–§4](agent-context/GUIDELINES.md#2-spec-driven-test-first).
- **Examples are tests.** Manifests under [`examples/`](examples/) double as uptest fixtures — keep
  them valid, minimal, and annotated. An example that can't apply cleanly is a defect.
- E-tier (uptest, real gridscale credentials) runs in CI via `/test-examples` and nightly, not on
  every push.

## Local preflight

Before you push, run:

```bash
make reviewable
```

`make reviewable` regenerates, lints, and tidies — it must be **clean** before you open a PR, and it
is what guards the generated tree in CI. It is also worth running `make test`, `make build`, and
`make check-diff` while iterating. (The structural spec gate is described in
[GUIDELINES §2](agent-context/GUIDELINES.md#2-spec-driven-test-first).)

## Commit convention

Commits follow **gitmoji + Conventional Commits**, using the **ASCII shortcode** form:

```
:gitmoji: type(scope): summary
```

Examples from this repo's history:

```
:sparkles: feat(config): add external-name override for object storage
:memo: docs(contributing): add CONTRIBUTING guide (E6-S01)
:white_check_mark: test(config): assert external-name strategy
:rotating_light: fix(clients): wire gridscale uuid/token credentials
```

Rules (see [AGENTS.md → Conventions](AGENTS.md#conventions)):

- Use the **ASCII shortcode** (`:sparkles:`) — **no Unicode emoji**.
- **One logical change per commit**; keep the tree green at each commit.
- No AI co-author trailers; never modify git config.

## Opening a pull request & merge policy

1. Branch from `main`; keep the change to one logical unit of work.
2. Make sure `make reviewable` is clean and tests pass, with no hand-edited generated files.
3. Open a PR describing the change and its motivation.

Merges use **rebase-and-merge only** — linear history, no squash, no merge commits.

A change is mergeable only when it meets the **Definition of Done** in
[GUIDELINES §8](agent-context/GUIDELINES.md#8-definition-of-done): failing test first (for
hand-authored code) · `make reviewable` clean · no hand-edited generated files · ADR added for any
real decision · examples/docs regenerated · one logical change per commit, tree green at each. A red
gate means not mergeable.

## Security

Please do **not** report security issues in public GitHub issues. gridscale tokens are supplied via a
`ProviderConfig` secretRef only — never commit secrets. Secret-scanning (`gitleaks`) and pre-commit
hooks are planned under E5-S03; until then, review your diffs carefully before pushing.
See [GUIDELINES §6](agent-context/GUIDELINES.md#6-security--supply-chain) for the supply-chain posture.

## What to work on next

The build order is public. To find "what to do next":

- Read [`docs/ROADMAP.md`](docs/ROADMAP.md) — epics **E1–E6**, build order, and exit criteria.
- Follow the **Next action** section in [`AGENTS.md`](AGENTS.md#next-action) (or run the
  `/pick-next-story` helper).

Every implementation lane maps to an OpenSpec change under
[`openspec/changes/`](openspec/) with `REQ ↔ Test ↔ Verify` traceability — see
[ADR-0001](docs/adr/0001-spec-driven-workflow.md).

Thanks for contributing!
