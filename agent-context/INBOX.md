# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Handoff 2026-07-16 wrap:** `main` @ `3edc5f0`; Marketplace **v0.1.1**. See
> [`coordination/SESSION-HANDOFF.md`](coordination/SESSION-HANDOFF.md).

---

## Decisions (open) — confirmations (non-blocking)

1. **Marketplace icon source** — used operator `.cache/gridscale.svg` (+ PNG) for
   `extensions/icons/icon.svg` and branding light/dark. **Recommended:** keep.
   **Answer / instructions:** _(operator)_

2. **Unaffiliation wording** — shipped as:
   > community PlatformRelay project … **not** affiliated with, endorsed by, or an official
   > product of gridscale GmbH.
   **Recommended:** keep. **Answer / instructions:** _(operator)_

3. **Post-append cosign** — each `up alpha xpkg append` invalidates prior keyless signatures.
   `v0.1.1` was signed in publish then re-appended for icon/readme extensions — re-sign if
   Scorecard/supply-chain demo needs a valid signature on the current digest.
   **Recommended:** re-dispatch sign when convenient (not blocking demo copy).
   **Answer / instructions:** _(operator)_

---

## Operator tasks

1. **Revoke the old classic PAT** that was briefly stored as `GHCR_PAT` (local PAT in
   `.envrc` only — do **not** put it in Actions secrets).
2. **Diagnose red CI on `main` @ `3edc5f0`** — primary **CI** workflow failed
   ([run 29515604942](https://github.com/PlatformRelay/provider-gridscale/actions/runs/29515604942));
   Gitleaks / Scorecard / Govulncheck / Coverage / Docs Sync were green. Next coding loop should
   not assume `main` is fully green until this is fixed or explained.

### Non-blocking / optional

- Delete superseded remote branches: `add-config-unit-tests`, `feat/official-gridscale-logo`,
  `fix-credential-wiring`, `rewrite-readme-gridscale`, `worktree-gridscale-audit-2026-07-15`
  (content already on `main`).
- Upjet `DataSourceSchemas` feature request (D-015) — do not auto-file.
- Optional: close upstream TF #188; nudge doc drafts #467/#468 (D-016).

---

## Reference — resolved / no action

- D-007…D-017b, audit-gap E7 + E5-S07…S10, CRED live smoke, BRAND-1 / D-009b, Batch 7 polish —
  see `decisions.md` / SESSION-HANDOFF.
- E2-S04/S05 uptest — **intentionally skipped** (D-012 → B).
