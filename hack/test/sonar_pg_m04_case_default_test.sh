#!/usr/bin/env bash
# REQ-SONAR-PG-M04: hack/check-docs.sh case must include a *) default arm (shelldre:S131).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SCRIPT="${ROOT}/hack/check-docs.sh"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${SCRIPT}" ]] || fail "missing ${SCRIPT}"

# Extract the case…esac that filters CRD basenames (must have an explicit default).
case_block="$(awk '
  /case[[:space:]]+"/ { in_case=1 }
  in_case { print }
  in_case && /^[[:space:]]*esac[[:space:]]*$/ { exit }
' "${SCRIPT}")"

[[ -n "${case_block}" ]] || fail "${SCRIPT}: no case/esac block found"

if ! printf '%s\n' "${case_block}" | grep -Eq '^[[:space:]]*\*\)'; then
  fail "${SCRIPT}: case missing *) default arm (REQ-SONAR-PG-M04 / shelldre:S131)"
fi

echo "PASS: REQ-SONAR-PG-M04 ${SCRIPT#$ROOT/} case has *) default"
