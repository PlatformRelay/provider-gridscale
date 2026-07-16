# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Status 2026-07-16:** D-017 — public Upbound + GHCR package linked to this repo. **No personal
> PAT in repo secrets** (kollect/mkurator model: `github.token` + `packages: write`). `GHCR_PAT`
> secret **deleted**. Only Upbound robot mirror creds remain as secrets.

---

## Decisions (open)

**None.**

---

## Operator tasks

### Publish unblock (remaining)

1. **Revoke the classic PAT** that was briefly stored as `GHCR_PAT` (even though the secret is
   deleted — rotate/revoke at GitHub → Settings → Developer settings → PATs). Keep a local PAT in
   `.envrc` only if you need it for `gh` CLI; CI must not use it.
2. **Re-dispatch publish** (after the workflow change lands on `main`):
   ```bash
   gh workflow run publish-provider-package.yml -R PlatformRelay/provider-gridscale -f version=v0.1.0
   ```
3. **Upbound Marketplace listing** — repo is public but `publishPolicy: draft`. When rate-limit
   clears:
   ```bash
   up repository update provider-gridscale --private=false --publish --force --organization=platformrelay
   ```
   First public listing may need Upbound review via Crossplane Slack `#upbound`.

### Other (non-blocking)

- **Upjet feature request — `DataSourceSchemas` (from D-015; outward-facing, do not auto-file).**
- **Optional — close stale upstream #188 / nudge doc drafts (from D-016 / E6-S06).**

---

## Reference — resolved / no action

- D-007…D-017 (see `decisions.md`). GHCR package linked to `PlatformRelay/provider-gridscale` and
  public. Publish/sign uses `github.token` (no PAT). Upbound robot secrets retained for mirror only.
