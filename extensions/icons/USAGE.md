# Usage

`icon.svg` is the official **gridscale figurative mark** (Bildmarke) from the
gridscale press Logo-Package (synced from the operator-provided asset under
`.cache/gridscale.svg`). Keep square dimensions; do not recolour — swap to
`docs/assets/branding/icon-{light,dark}.svg` for contrast instead.

`icon.png` is the same mark as a 512×512 raster for consumers that prefer PNG.
Marketplace listing icons are attached to a published version with:

```sh
up alpha xpkg append --extensions-root=./extensions \
  xpkg.upbound.io/platformrelay/provider-gridscale:<version>
```
