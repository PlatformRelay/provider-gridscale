# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

---

## Decisions (open)

_(none — D-001…D-004 answered **A** and the E1–E6 plan approved by the operator 2026-07-15; see
[`decisions.md`](decisions.md).)_

---

## Operator tasks / reviews / PRs

- **CI: `local-deploy` red on `main`.** The `local-deploy` job fails in buildx `img.build.shared`
  with `failed to load cache key: invalid response status 404` (upstream `build/` submodule,
  `build/makelib/imagelight.mk`). Deterministic (both commits). Not hand-editable per **D-003**;
  needs an operator/infra look before `main` can go fully green and before any lane can auto-merge.
- **PR #1 — gofmt fix** (`worktree-fix-gofmt-config`): clears the `lint` job's gofmt failure on
  `main`. Operator-merge (its CI stays red on `local-deploy` above until that is resolved).
