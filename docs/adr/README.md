# Architecture Decision Records — provider-gridscale

Short, dated records of non-obvious decisions. Each ADR: **Context → Decision → Consequences →
Counterpoints considered**. Status is `Proposed`, `Accepted`, `Superseded`, or `Deprecated`.

Theme-numbered `0Txx` (kaddy convention): `01` foundations · `02` testing & quality · `03` release &
supply chain.

## Index

| ADR | Title | Status |
| --- | --- | --- |
| [0001](0001-spec-driven-workflow.md) | Adopt a committed spec-driven (OpenSpec) workflow | Accepted |
| [0002](0002-testing-strategy.md) | Testing strategy for a generated Upjet provider | Accepted |

## Conventions

- One decision per file; filename `NNNN-kebab-title.md`.
- Link the OpenSpec change / `decisions.md` entry (`D-nnn`) that the ADR realises.
- Supersede rather than delete: mark the old ADR `Superseded by NNNN` and add a new one.
