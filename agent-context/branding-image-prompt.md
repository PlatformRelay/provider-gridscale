# Branding image prompt — provider-gridscale icon (D-009)

Hand this to an image-generation agent to produce the **real** provider icon that replaces the current
placeholder (`extensions/icons/icon.svg` + `docs/assets/branding/icon-{light,dark}.svg`). Decision
**D-009 → A**: ship v0.1.0 with the placeholder / an interim gridscale web logo; commission the real
mark from this prompt afterwards. Do **not** block the release on it.

## What this icon is for

The Crossplane/Upbound **Marketplace listing tile** and the README logo for `provider-gridscale` — a
Crossplane provider that manages [gridscale](https://gridscale.io) cloud infrastructure (servers,
storage, networks, managed databases, Kubernetes) as Kubernetes custom resources. It sits next to
other provider tiles (provider-aws, provider-gcp, provider-azure), so it must read as **"gridscale,
on Kubernetes/Crossplane"** at a glance and hold up at small sizes.

## Primary prompt

> A flat, modern, minimal **square** app icon for a Kubernetes/Crossplane infrastructure provider.
> Subject: the **gridscale** cloud brand fused with the Crossplane ecosystem — a clean geometric
> **grid** motif (nod to "grid-scale") resolving into interconnected infrastructure nodes, subtly
> echoing the Crossplane hexagon/orbit language without copying any existing provider logo. Centered
> single mark, no text, no wordmark. Bold, high-contrast, confident negative space; crisp geometry
> that stays legible from 512px down to a 24px favicon. Vector-style flat design, no photorealism, no
> gradients heavier than a single subtle two-stop shade, no drop shadows, no 3D bevels.

## Style constraints (hard requirements)

- **Format:** delivered as clean **SVG** (simple paths, no embedded raster, no filters that don't
  survive SVG export). Square 1:1 aspect, designed on a 512×512 grid with ~10% safe margin.
- **No text / no letters** in the mark.
- **Palette:** lead with gridscale's brand teal/green family; pair with a neutral dark
  (`#0B1F2A`-ish) and clean white. Keep to **≤3 core colors** so it prints and scales cleanly.
- **Two variants required:**
  - `icon-light.svg` — mark on transparent/white background (for light UIs).
  - `icon-dark.svg` — same mark tuned for dark backgrounds (raise contrast; don't just invert).
- **Contrast:** must pass a legibility check on both `#FFFFFF` and `#0D1117` (GitHub dark) backgrounds.
- **Original artwork only** — do not trace or reproduce gridscale's exact registered logo or any other
  provider's logo; this is an ecosystem mark *inspired by* the gridscale brand, safe to ship under the
  repo's Apache-2.0 license.

## Deliverables

1. `extensions/icons/icon.svg` — the canonical square Marketplace icon.
2. `docs/assets/branding/icon-light.svg` and `icon-dark.svg` — README / docs variants.
3. `docs/assets/branding/social-card.svg` (optional, stretch) — a 1280×640 social/OpenGraph card
   using the mark + "provider-gridscale" wordmark for GitHub link previews.

## After generation

- Replace the placeholder assets, then set/update the Marketplace `iconURI` annotation in
  `package/crossplane.yaml` (added by E3-S03; point it at the final icon) and the README logo `<img>`.
- Sanity-check the SVG renders in GitHub markdown (some filters/`<style>` blocks are stripped).

> **Interim (v0.1.0):** until this mark exists, use gridscale's public web logo as the icon and keep
> the placeholder in `docs/assets/branding/`. Tracked under D-009 / E3-S01.
