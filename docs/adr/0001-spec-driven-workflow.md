# ADR-0001: Adopt a committed spec-driven (OpenSpec) workflow

**Theme:** 01 · Foundations · **Status:** Accepted · **Date:** 2026-07-15 · **Realises:** D-000

## Context

`provider-gridscale` was scaffolded from the upjet-provider-template with no development harness:
no specs, backlog, decision log, or contribution guidance. The sibling repos (`kaddy`, `kollect`,
`mkurator`) all run a spec-driven, TDD, ADR-backed workflow with per-lane traceability. To raise
this provider to a professional OSS standard we need the same discipline, and we need the *maturity
of the process* to be visible to external contributors and adopters.

## Decision

**Adopt kaddy's OpenSpec workflow and commit it to the repo.**

- `openspec/` — `config.yaml` (conventions + provider-shaped test taxonomy), `changes/<slug>/`
  (proposal + design + tasks + specs) per lane, `specs/` (accepted capability specs).
- `agent-context/` — `AGENTS.md` map, `GUIDELINES.md`, `BACKLOG.md`, `INBOX.md`, `decisions.md`.
- Root `AGENTS.md` entry point; `docs/adr/` for decisions.
- **Committed, not gitignored** — the divergence from kaddy (which gitignores its harness) is
  deliberate: for a public provider, a visible spec-driven workflow is itself a trust/maturity
  signal. Only `.claude/` (local coordination) stays gitignored.

## Consequences

- Every implementation lane maps to an OpenSpec change with `REQ ↔ Test ↔ Verify` traceability,
  gated by `make test.spec`.
- Contributors get a single entry point (`AGENTS.md`) and a discoverable "what next" (`/pick-next-story`).
- The repo carries process docs a casual reader may skip — mitigated by keeping `AGENTS.md` a thin
  doc-map with detail behind links.

## Counterpoints considered

- *Keep the harness gitignored like kaddy.* Rejected: kaddy is a private exercise; a public provider
  benefits from showing its rigour, and committed specs let CI enforce `test.spec`.
- *Lighter weight (a CONTRIBUTING.md and issues only).* Rejected: loses per-lane spec traceability
  and the TDD gate that keeps the small hand-authored surface honest.
