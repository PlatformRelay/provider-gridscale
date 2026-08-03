#!/usr/bin/env bash
# REQ-SONAR-PG-M01: apis/{cluster,namespaced}/v1alpha1/register.go must not keep a
# bare-empty `func init() {}` (go:S1186). Require a nested comment or other
# non-empty body. Do not call AddToScheme from these stubs.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

check_register_init() {
  local file="$1"
  [[ -f "${file}" ]] || fail "missing ${file}"

  # Fail when `func init()` is declared with a bare-empty body: `func init() {}`
  # (optional whitespace inside braces). Nested comments or statements pass.
  if grep -Eq 'func[[:space:]]+init[[:space:]]*\([[:space:]]*\)[[:space:]]*\{[[:space:]]*\}' "${file}"; then
    fail "${file}: bare-empty func init() {} (REQ-SONAR-PG-M01 / go:S1186)"
  fi

  if ! grep -Eq 'func[[:space:]]+init[[:space:]]*\(' "${file}"; then
    fail "${file}: missing func init() (do not delete; nest a comment instead)"
  fi

  echo "PASS: ${file#$ROOT/} init body is not bare-empty"
}

check_register_init "${ROOT}/apis/cluster/v1alpha1/register.go"
check_register_init "${ROOT}/apis/namespaced/v1alpha1/register.go"

echo "PASS: REQ-SONAR-PG-M01 register init stubs are non-empty"
