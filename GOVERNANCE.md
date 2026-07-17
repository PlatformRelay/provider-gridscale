# Governance — provider-gridscale

This document describes how **provider-gridscale** is governed: the roles people play, how
decisions get made and recorded, how a change reaches `main`, and how someone becomes a maintainer.
It is deliberately **lightweight** — sized for a small, single-provider open-source project, not an
enterprise foundation. As the community grows, this model can grow with it.

It complements, and does not restate, the operational documents:

- **How to contribute:** [`CONTRIBUTING.md`](CONTRIBUTING.md) — workflow, local preflight, commit convention.
- **Engineering rules (binding):** [`agent-context/GUIDELINES.md`](agent-context/GUIDELINES.md) — a red gate ⇒ not mergeable.
- **Definition of Done:** [GUIDELINES §8](agent-context/GUIDELINES.md#8-definition-of-done).
- **Code of Conduct:** [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) (Contributor Covenant 2.1).
- **Roadmap & build order:** [`docs/ROADMAP.md`](docs/ROADMAP.md).

## Roles & responsibilities

We keep the role set small. In practice there are three:

### Contributors

Anyone who opens an issue or a pull request. Contributors:

- Follow [`CONTRIBUTING.md`](CONTRIBUTING.md) and the [engineering guidelines](agent-context/GUIDELINES.md).
- Respect the **generation boundary** — change `config/` inputs and regenerate, never hand-edit
  generated output (see [ADR-0003](docs/adr/0003-codegen-boundary.md)).
- Write the failing test first for hand-authored Go / `hack/` scripts, and keep the tree green.
- Abide by the [Code of Conduct](CODE_OF_CONDUCT.md).

No prior standing is required to contribute; every merged change starts as a contributor's PR.

### Reviewers

Contributors with a track record who review PRs and mentor newcomers. Reviewing is a
responsibility any trusted contributor can take on — it is **not a separate formal tier** with its
own permissions. A reviewer's approval is valuable, but the **binding** approval for a merge comes
from a maintainer (see [Merge policy](#merge-policy)). Consistent, high-quality review is the main
path toward maintainership.

### Maintainers

The people accountable for the health and direction of the repository. Maintainers:

- Have write/merge access and are the required approvers for every PR.
- Are listed authoritatively in [`CODEOWNERS`](CODEOWNERS) and [`OWNERS.md`](OWNERS.md); those two
  files are the single source of truth for *who* the maintainers are.
- Triage issues, uphold the [Definition of Done](agent-context/GUIDELINES.md#8-definition-of-done)
  and merge policy, cut releases, and record decisions (ADRs / `decisions.md`).
- Handle Code of Conduct reports and coordinated security disclosure
  (see [`SECURITY.md`](SECURITY.md)).
- Steward the roadmap and are the operator(s) for the open decisions this project tracks.

**Current maintainers:** Konrad Heimel (@konih)

> The authoritative maintainer list lives in [`CODEOWNERS`](CODEOWNERS) and [`OWNERS.md`](OWNERS.md);
> this document intentionally does not duplicate the names. (Maintainer identity was resolved by
> **D-008** — see [`agent-context/decisions.md`](agent-context/decisions.md).)

## How decisions are made

The project runs a **spec-driven, decision-logged** workflow (see
[ADR-0001](docs/adr/0001-spec-driven-workflow.md)). Two lightweight artifacts carry the record:

1. **ADRs — [`docs/adr/`](docs/adr/README.md).** Every non-obvious or architectural decision is
   captured as a short Architecture Decision Record (*Context → Decision → Consequences →
   Counterpoints considered*). ADRs are how the project explains **why** it is shaped the way it is.

2. **Decision log — [`agent-context/decisions.md`](agent-context/decisions.md).** Scope-shaping and
   operator decisions (the `D-nnn` series, e.g. **D-008** for maintainer identity) are
   logged here with their counterpoints, then realised through ADRs and implementation lanes.

Implementation lanes are traceable back to a spec under [`openspec/`](openspec/config.yaml) with
`REQ ↔ Test ↔ Verify` links, so a decision, its rationale, and the code that realises it stay
connected. Day-to-day decisions are made by **lazy consensus** among maintainers on the relevant
issue or PR; a decision with lasting consequences is written down (ADR + `decisions.md`) rather than
left implicit. There is no heavier voting machinery at this maturity — if the project ever needs it,
that change would itself be recorded as an ADR.

## Merge policy

The authoritative rules live in
[CONTRIBUTING.md → merge policy](CONTRIBUTING.md#opening-a-pull-request--merge-policy) and
[GUIDELINES §8](agent-context/GUIDELINES.md#8-definition-of-done). In summary:

- **Rebase-and-merge only** — linear history, no squash, no merge commits.
- **Green gate is required.** A red gate (`make reviewable`, tests, or CI failing, or hand-edited
  generated files) means the change is **not mergeable**.
- **At least one maintainer approval** is required before any PR merges; [`CODEOWNERS`](CODEOWNERS)
  assigns that reviewer automatically.

## The role of CODEOWNERS

[`CODEOWNERS`](CODEOWNERS) is the mechanism that operationalises the review requirement: it
automatically requests a maintainer as reviewer on each PR and, together with branch protection,
enforces that **no change merges without maintainer approval**. It is intentionally the same list of
people as [`OWNERS.md`](OWNERS.md) — `OWNERS.md` is the human-readable roster, `CODEOWNERS` is the
machine-enforced one. Keeping the two in sync (and in sync with the maintainer list above) is a
maintainer responsibility.

## Becoming a maintainer

Maintainership is earned through **sustained, high-quality contribution**, not appointed on request.
A typical path:

1. Contribute merged PRs and help review others' work over a period of time.
2. Demonstrate good judgement about the generation boundary, tests, and the Definition of Done.
3. An existing maintainer nominates you; the current maintainers reach consensus.
4. On agreement, you are added to [`CODEOWNERS`](CODEOWNERS) and [`OWNERS.md`](OWNERS.md) — that
   addition **is** the promotion, and the change is recorded like any other decision.

Maintainers who step back are moved to an emeritus note (or simply removed from the owner files);
inactive maintainers may be retired by consensus of the active ones. The initial maintainer roster is set (@konih, per **D-008**); this process governs future
changes to it.

## Code of Conduct enforcement

The project follows the [Contributor Covenant 2.1](CODE_OF_CONDUCT.md). Enforcement is a maintainer
responsibility, following the escalation ladder in that document. Reports are made through the
contact channel defined in [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md#enforcement) — that contact is
itself an open operator item, so please refer to the Code of Conduct for the current channel rather
than a value duplicated here.

## Relationship to the OpenSSF baseline

Community-health and governance transparency are tracked as **Epic E6** in the
[roadmap](docs/ROADMAP.md#e6--community-health--governance), with the explicit goal of meeting the
**OpenSSF** baseline for an open-source project — a newcomer should be able to find how to
contribute, how to report a vulnerability, and who owns the repository.

This `GOVERNANCE.md`, alongside [`CONTRIBUTING.md`](CONTRIBUTING.md),
[`SECURITY.md`](SECURITY.md), and the [Code of Conduct](CODE_OF_CONDUCT.md), exists to make the
project's governance **transparent and discoverable** — publishing the model *is* the maturity
signal. Supply-chain automation from [Epic E5](docs/ROADMAP.md#e5--cicd-hardening--supply-chain) is
**running** (CodeQL, OpenSSF Scorecard, govulncheck, gitleaks, coverage). Branch protection on
`main` is enforced by ruleset **`protect-main`**: one approval for external contributors, required
CI/SAST checks, and an Admin **pull-request-only** bypass so the solo maintainer can merge without
a second human (`gh pr merge --rebase --admin`). Scorecard **Code-Review** (second-person approvals)
remains a multi-maintainer goal.
