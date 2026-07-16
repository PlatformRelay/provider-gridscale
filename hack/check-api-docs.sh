#!/usr/bin/env bash
# check-api-docs.sh — D-tier sync guard for the generated CRD API reference
# (E4-S01). Regenerates docs/api/ via elastic/crd-ref-docs and fails if the
# committed tree drifts. Kept independent of the `build/` submodule so CI and
# local runs share one hermetic path. Pair with docs-sync.yml.
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "${root}"

CRD_REF_DOCS_VERSION="${CRD_REF_DOCS_VERSION:-v0.2.0}"

echo "regenerating docs/api/ with crd-ref-docs@${CRD_REF_DOCS_VERSION}…"
mkdir -p docs/api
go run "github.com/elastic/crd-ref-docs@${CRD_REF_DOCS_VERSION}" \
  --source-path=apis \
  --config=.crd-ref-docs.yaml \
  --renderer=markdown \
  --output-path=docs/api/

if ! git diff --quiet -- docs/api/; then
  echo "::error::docs/api/ is stale relative to apis/** — run \`make docs\` and commit."
  git --no-pager diff --stat -- docs/api/
  git --no-pager diff -- docs/api/ | head -n 200
  exit 1
fi

echo "docs/api/ is in sync with apis/**."
