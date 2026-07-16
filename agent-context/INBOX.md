# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Status 2026-07-16 evening:** decisions D-008…D-015 closed. v0.1.0 is on ghcr. Operator refreshed
> local creds in `.envrc` (classic `GITHUB_TOKEN` PAT + Upbound **robot** mirror token). Remaining:
> push those into **GitHub Actions secrets**, fix a likely typo in the robot JWT, then re-dispatch
> publish. `tag.yaml` pin is being fixed on `main` (self-contained workflow — upstream lost
> `workflow_call`).

---

## Decisions (open)

**None.**

---

## Operator tasks

### CRED — GHCR signing + Upbound mirror (in progress 2026-07-16)

1. **`GHCR_PAT` needs `read:packages` (or link the package).** The repo secret `GHCR_PAT` is what
   `publish-provider-package.yml` / `sign-and-sbom` uses — **not** the local `.envrc` `GITHUB_TOKEN`
   unless you copy it into that secret. See **How-to** below. Then re-dispatch:
   `gh workflow run publish-provider-package.yml -R PlatformRelay/provider-gridscale -f version=v0.1.0`
2. **Upbound mirror robot token regenerated (operator, 2026-07-16).** Local `.envrc` now has
   `XPKG_MIRROR_TOKEN` + `XPKG_MIRROR_ACCESS_ID` from an Upbound **robot** account (not a user PAT).
   **Must also update the GitHub Actions secrets** of the same names (still dated 2026-07-15 in the
   repo). **Check for typo:** the JWT in `.envrc` currently starts with `eeyJ…` — valid JWTs start
   with `eyJ…`. Drop the leading `e` if that was a paste error, then `gh secret set`.
3. **`tag.yaml` repin** — agent actioning: replace broken reusable-workflow call with a
   self-contained Tag workflow (upstream `tag.yml` no longer offers `workflow_call`).

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
