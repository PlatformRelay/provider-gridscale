#!/usr/bin/env bash
# REQ-SONAR-PG-05: Scorecard workflow uses explicit minimal permissions
# (not top-level permissions: read-all).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/scorecard.yml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

# Top-level must not grant the broad read-all shortcut.
if grep -Eq '^[[:space:]]*permissions:[[:space:]]*read-all[[:space:]]*$' "${WF}"; then
  fail "${WF}: top-level 'permissions: read-all' must be replaced with explicit scopes"
fi

# Required minimal scopes (top-level or job-level).
required=(
  'security-events:[[:space:]]*write'
  'id-token:[[:space:]]*write'
  'contents:[[:space:]]*read'
  'actions:[[:space:]]*read'
)

for pat in "${required[@]}"; do
  if ! grep -Eq "${pat}" "${WF}"; then
    fail "${WF}: missing required permission scope matching /${pat}/"
  fi
done

# SARIF upload to code scanning needs security-events: write.
if ! grep -Eq 'security-events:[[:space:]]*write' "${WF}"; then
  fail "${WF}: security-events: write required for SARIF upload"
fi

echo "PASS: scorecard.yml uses explicit Scorecard permissions (no read-all)"
