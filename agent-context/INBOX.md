# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Status 2026-07-16:** D-017 accepted — public Upbound repo created; GHCR package still needs a
> UI link to this repo; PAT needs `read:packages`; then re-dispatch publish.

---

## Decisions (open)

**None.**

---

## Operator tasks

### CRED / publish unblock (D-017) — do these once

1. **Add `read:packages` to the classic PAT** used as `GHCR_PAT` (current scopes have
   `write:packages` + `delete:packages` but **not** `read:packages` — that breaks `crane digest` /
   signing). Re-`gh secret set GHCR_PAT` after. Also **Authorize SSO** for the PlatformRelay org on
   that PAT if prompted.
2. **Link the org GHCR package to this repo** (API link 404'd — UI required):
   https://github.com/orgs/PlatformRelay/packages/container/package/provider-gridscale
   → Package settings → **Manage Actions access** → add `PlatformRelay/provider-gridscale` with
   **Read and write**. (Package is already `visibility: public`; `repository` is still `null`.)
3. After (1)+(2), re-dispatch:
   ```bash
   gh workflow run publish-provider-package.yml -R PlatformRelay/provider-gridscale -f version=v0.1.0
   ```
4. **Upbound Marketplace listing** — repo `platformrelay/provider-gridscale` is **public** but
   `publishPolicy: draft`. When the create rate-limit clears (~minutes), run:
   ```bash
   up repository update provider-gridscale --private=false --publish --force --organization=platformrelay
   ```
   First public listing may still need Upbound review via Crossplane Slack `#upbound`
   (git repo + account + repo name). Target URL:
   https://marketplace.upbound.io/providers/platformrelay/provider-gridscale/

### Other (non-blocking)

- **Upjet feature request — `DataSourceSchemas` (from D-015; outward-facing, do not auto-file).**
- **Optional — close stale upstream #188 / nudge doc drafts (from D-016 / E6-S06).**
- **`tag.yaml`** — already being fixed on `main` if not merged; self-contained Tag workflow.

---

## Reference — resolved / no action

- D-007…D-017 (see `decisions.md`). v0.1.0 artifact already on public ghcr; mirror + sign pending
  cred/link steps above. Actions secrets were refreshed 2026-07-16; robot write granted via team
  `providers`.
