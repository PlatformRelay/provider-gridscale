# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Status 2026-07-16 evening:** decisions D-008…D-015 closed. v0.1.0 is on ghcr. Operator set
> Actions secrets `XPKG_MIRROR_TOKEN` + `XPKG_MIRROR_ACCESS_ID` (robot) and linked the GHCR package
> to this repo. Workflow updated to use `github.token` for GHCR push/sign (old `GHCR_PAT` was
> denied at login). Re-dispatch publish after that lands on `main`.

---

## Decisions (open)

**None.**

---

## Operator tasks

### CRED — GHCR signing + Upbound mirror (2026-07-16)

- [x] Upbound robot → Actions secrets `XPKG_MIRROR_TOKEN` / `XPKG_MIRROR_ACCESS_ID` (operator).
- [x] GHCR package linked to repo (operator).
- [x] `tag.yaml` self-contained (agent).
- [x] **sign-and-sbom green** on run [29509707113](https://github.com/PlatformRelay/provider-gridscale/actions/runs/29509707113)
      — `github.token` + linked package signed + attested existing `v0.1.0`.
- [ ] **Still need:** refresh repo secret `GHCR_PAT` with a classic PAT that has
      `write:packages` **and** `read:packages`. Upstream publish login uses
      `username=PlatformRelay` + that secret; `github.token` is denied there (same run).
      Mirror job never ran because publish-artifacts failed first.
      ```bash
      # from provider-gridscale/ after direnv allow — PAT must include packages scopes
      gh secret set GHCR_PAT -R PlatformRelay/provider-gridscale < <(printenv GITHUB_TOKEN)
      gh workflow run publish-provider-package.yml -R PlatformRelay/provider-gridscale -f version=v0.1.0
      ```

### Other (non-blocking)

- **Upjet feature request — `DataSourceSchemas` (from D-015; outward-facing, do not auto-file).**
- **Optional — close stale upstream #188 / nudge doc drafts (from D-016 / E6-S06).**

---

## How-to — grant GHCR read for signing (pick one)

### Option A — regenerate / edit the PAT used as `GHCR_PAT` (simplest)

1. GitHub → Settings → Developer settings → Personal access tokens (classic).
2. Create or edit a token with **both** `write:packages` **and** `read:packages` (and `repo` if you
   also use it for private ops). `write:packages` alone is not enough for `crane digest` on an
   org package the token just pushed in some org visibility setups.
3. Set the **repo secret** (this is what CI reads):
   ```bash
   gh secret set GHCR_PAT -R PlatformRelay/provider-gridscale
   # paste the PAT when prompted
   ```
4. Re-dispatch publish with `version=v0.1.0` so `sign-and-sbom` can `crane digest` + cosign.

If your local `.envrc` `GITHUB_TOKEN` already has those scopes, you can pipe it into the secret
without printing it:
```bash
# from provider-gridscale/, after direnv allow — do NOT echo the value
gh secret set GHCR_PAT -R PlatformRelay/provider-gridscale < <(printenv GITHUB_TOKEN)
```

### Option B — link the org GHCR package to this repository

So the default `GITHUB_TOKEN` in Actions can read/write the package (still requires workflow change
if the sign job keeps using `secrets.GHCR_PAT`):

1. Open https://github.com/orgs/PlatformRelay/packages?repo_name=provider-gridscale
   (or Packages → `provider-gridscale` → Package settings).
2. **Manage Actions access** → add `PlatformRelay/provider-gridscale` with **Read and write**.
3. Optionally tighten package visibility. Then either keep using an updated `GHCR_PAT` (A) **or**
   change `sign-and-sbom` login to `${{ secrets.GITHUB_TOKEN }}` / `github.actor` after linking.

**Agent cannot complete A/B without a packages-capable token in this shell** (org package API was
Forbidden with current auth). Operator must set the secret or link the package once.

### Upbound robot → Actions secrets

```bash
# Fix JWT typo in .envrc first if it starts with eeyJ (should be eyJ), then:
gh secret set XPKG_MIRROR_TOKEN -R PlatformRelay/provider-gridscale < <(printenv XPKG_MIRROR_TOKEN)
gh secret set XPKG_MIRROR_ACCESS_ID -R PlatformRelay/provider-gridscale < <(printenv XPKG_MIRROR_ACCESS_ID)
```

---

## Reference — resolved / no action

- D-007…D-015, CRED-1/2 live smoke, AUDIT 2026-07-15, PR #7, RELEASE v0.1.0 published to ghcr —
  see `decisions.md` / prior INBOX history.
