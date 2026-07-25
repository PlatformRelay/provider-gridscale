#!/usr/bin/env bash
# REQ-SONAR-PG-06: gitleaks installer curl must pin HTTPS + TLS 1.2+.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/gitleaks.yml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

# Installer download must use curl with HTTPS protocol lock and TLS 1.2 floor.
if ! grep -Eq "curl[[:space:]].*--proto[[:space:]]+'=https'" "${WF}"; then
  fail "${WF}: gitleaks installer curl missing --proto '=https'"
fi

if ! grep -Eq "curl[[:space:]].*--tlsv1\.2" "${WF}"; then
  fail "${WF}: gitleaks installer curl missing --tlsv1.2"
fi

echo "PASS: gitleaks.yml installer curl enforces HTTPS (--proto '=https' --tlsv1.2)"
