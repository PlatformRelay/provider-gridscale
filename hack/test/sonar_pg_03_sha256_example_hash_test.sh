#!/usr/bin/env bash
# REQ-SONAR-PG-03: multi-example status-check context hash uses SHA-256, not MD5.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/e2e.yaml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

if grep -Eq 'md5sum' "${WF}"; then
  fail "${WF}: md5sum must not be used for example-list hash (use sha256sum or shasum -a 256)"
fi

if ! grep -Eq 'sha256sum|shasum[[:space:]]+-a[[:space:]]*256' "${WF}"; then
  fail "${WF}: example-list hash must use sha256sum (or shasum -a 256)"
fi

echo "PASS: example-list hash uses SHA-256"
