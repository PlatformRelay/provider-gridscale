# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Handoff 2026-07-16 evening:** coding loop idle — see
> [`coordination/SESSION-HANDOFF.md`](coordination/SESSION-HANDOFF.md). No open decisions.

---

## Decisions (open)

**None.**

---

## Operator tasks

1. **Revoke the old classic PAT** that was briefly stored as `GHCR_PAT` (new local PAT is in
   `.envrc` only — do **not** put it in Actions secrets).
2. **Watch / finish `v0.1.0` publish** — inlined workflow run:
   https://github.com/PlatformRelay/provider-gridscale/actions/runs/29510417338  
   Re-dispatch if needed:
   `gh workflow run publish-provider-package.yml -R PlatformRelay/provider-gridscale -f version=v0.1.0`
3. **Publish Upbound Marketplace listing** (repo public, still `publishPolicy: draft`):
   ```bash
   up repository update provider-gridscale --private=false --publish --force --organization=platformrelay
   ```
   First public listing may need Upbound review via Crossplane Slack `#upbound`.

### Non-blocking

- Upjet `DataSourceSchemas` feature request (D-015) — outward-facing; do not auto-file.
- Optional: close upstream #188; nudge doc drafts #467/#468 (D-016).

---

## Reference — resolved / no action

- D-007…D-017b, audit-gap E7 + E5-S07…S10 + E4-S05 + E6-S06, CRED live smoke, v0.1.0 on public
  linked GHCR — see `decisions.md` / SESSION-HANDOFF.
