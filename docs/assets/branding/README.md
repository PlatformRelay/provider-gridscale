# Branding assets (placeholder)

> **Status: PLACEHOLDER — pending operator/designer sign-off.**
> These are neutral, original geometric marks provided so the provider renders
> recognisably on the Upbound Marketplace and in docs. They are **not** an official
> brand and are **not** gridscale's trademarked logo. Replace them with an
> operator/designer-approved brand identity before treating this as final.

## The mark

A rounded square framing a 3x3 grid, with two accent nodes highlighted on the
diagonal — an abstract "grid of scale" motif. It is intentionally simple so it
stays legible from 32px up to 256px, uses flat fills only, and has no external
references, gradients, fonts, or scripts.

## Files

| File | Background | Use |
| --- | --- | --- |
| `../../../extensions/icons/icon.svg` | brand navy | Provider package icon (Marketplace). Source of truth for the mark. |
| `icon-light.svg` | light (`#F4F7FB`) | Docs / sites with light backgrounds. |
| `icon-dark.svg` | dark (`#0E1729`) | Docs / sites with dark backgrounds. |

## Usage guidance

- Keep the icon **square**; do not stretch or crop.
- Preserve clear space around the mark equal to roughly the corner radius.
- Prefer the variant that matches the surrounding background for contrast.
- Do not recolour arbitrarily; if a new palette is needed, take it to the
  operator/designer as part of the brand sign-off.

## Scope note

Wiring the icon into `package/crossplane.yaml` marketplace metadata is handled
separately (release-adjacent). This directory only provides the asset files.
