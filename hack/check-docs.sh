#!/usr/bin/env bash
# check-docs.sh — staleness guard for the hand-maintained README resource matrix
# (DOC-5). GUIDELINES §7 says API tables should be generated; until the matrix is
# fully generated, this fails if the README's advertised resource/group counts
# drift from the actual generated CRDs under package/crds/.
set -euo pipefail

crd_dir="package/crds"
readme="README.md"

# Managed-resource CRDs = cluster-scoped ("*.platformrelay.io_*.yaml"), excluding
# the ".m." namespaced mirrors and the provider-config machinery.
shopt -s nullglob
managed=()
for path in "${crd_dir}"/*.platformrelay.io_*.yaml; do
  name="$(basename "${path}")"
  case "${name}" in
    *.m.platformrelay.io_*) continue ;;
    *providerconfig*|*storeconfig*) continue ;;
  esac
  managed+=("${name}")
done

resource_count="${#managed[@]}"
# API service groups = leading label of the CRD group (e.g. "gridscale" in
# gridscale.gridscale.platformrelay.io; "mysql8" in mysql8.gridscale...).
group_count="$(printf '%s\n' "${managed[@]}" | sed -E 's/^([a-z0-9]+)\..*/\1/' | sort -u | wc -l | tr -d ' ')"

echo "actual: ${resource_count} managed resources across ${group_count} API groups"

fail=0
if ! grep -qE "\b${resource_count}\b" "${readme}"; then
  echo "::error::README does not mention the current resource count (${resource_count}); resource matrix is stale."
  fail=1
fi
if ! grep -qE "\b${group_count}\b" "${readme}"; then
  echo "::error::README does not mention the current group count (${group_count}); resource matrix is stale."
  fail=1
fi
[ "${fail}" -eq 0 ] && echo "README resource matrix is in sync (${resource_count}/${group_count})."
exit "${fail}"
