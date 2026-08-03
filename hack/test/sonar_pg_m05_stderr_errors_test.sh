#!/usr/bin/env bash
# REQ-SONAR-PG-M05: ::error:: / failure echoes in check-docs.sh and check-api-docs.sh
# must go to stderr (shelldre:S7677).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

check_error_stderr() {
  local file="$1"
  [[ -f "${file}" ]] || fail "missing ${file}"

  # Every echo that emits ::error:: (or similarly labeled error text) must redirect >&2.
  local line
  while IFS= read -r line; do
    [[ -z "${line}" ]] && continue
    if ! grep -Eq '>&2[[:space:]]*$|&>\&2' <<<"${line}"; then
      # Also accept >&2 before the string: echo >&2 "::error::..."
      if ! grep -Eq 'echo[[:space:]]+>&2' <<<"${line}"; then
        fail "${file}: error echo missing >&2: ${line}"
      fi
    fi
  done < <(grep -E 'echo[[:space:]].*::error::' "${file}" || true)

  if ! grep -Eq 'echo[[:space:]].*::error::' "${file}"; then
    fail "${file}: expected at least one ::error:: echo to guard (REQ-SONAR-PG-M05)"
  fi

  echo "PASS: ${file#$ROOT/} ::error:: echoes redirect to stderr"
}

check_error_stderr "${ROOT}/hack/check-docs.sh"
check_error_stderr "${ROOT}/hack/check-api-docs.sh"

echo "PASS: REQ-SONAR-PG-M05 stderr error redirects"
