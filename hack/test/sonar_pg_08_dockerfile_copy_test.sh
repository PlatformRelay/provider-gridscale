#!/usr/bin/env bash
# REQ-SONAR-PG-08: local Dockerfile artifacts must use COPY, not ADD.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
DOCKERFILE="${ROOT}/cluster/images/provider-gridscale/Dockerfile"

if [[ ! -f "${DOCKERFILE}" ]]; then
  echo "FAIL: Dockerfile not found at ${DOCKERFILE}" >&2
  exit 1
fi

failures=0

assert_local_copy() {
  local label="$1"
  local pattern="$2"

  if grep -E "^[[:space:]]*ADD[[:space:]]+${pattern}" "${DOCKERFILE}" >/dev/null; then
    echo "FAIL: ${label} uses ADD; must use COPY (docker:S7029 / REQ-SONAR-PG-08)" >&2
    failures=$((failures + 1))
    return
  fi

  if ! grep -E "^[[:space:]]*COPY[[:space:]]+${pattern}" "${DOCKERFILE}" >/dev/null; then
    echo "FAIL: ${label} missing COPY instruction matching ${pattern}" >&2
    failures=$((failures + 1))
    return
  fi

  echo "PASS: ${label} uses COPY"
}

# Quoted/unquoted local provider binary path.
assert_local_copy "bin/\${TARGETOS}_\${TARGETARCH}/provider" \
  '"?bin/\$\{TARGETOS\}_\$\{TARGETARCH\}/provider"?'

# Local terraform CLI config.
assert_local_copy "terraformrc.hcl" 'terraformrc\.hcl'

if [[ "${failures}" -ne 0 ]]; then
  echo "REQ-SONAR-PG-08: ${failures} failure(s)" >&2
  exit 1
fi

echo "REQ-SONAR-PG-08: all checks passed"
