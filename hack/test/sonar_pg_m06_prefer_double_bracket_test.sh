#!/usr/bin/env bash
# REQ-SONAR-PG-M06: listed hack scripts must prefer [[ over [ (shelldre:S7688).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

# Files Sonar flagged for S7688 in E5-S20 (single-bracket tests).
FILES=(
  "${ROOT}/hack/check-docs.sh"
  "${ROOT}/hack/test/e8_s01_observe_docs_test.sh"
  "${ROOT}/hack/test/e8_s01_observe_yaml_count_test.sh"
  "${ROOT}/hack/test/e8_s02_backuplist_crd_exists_test.sh"
  "${ROOT}/hack/test/e8_s03_publicnetwork_crd_exists_test.sh"
)

# Match a test/command form of `[` that is not part of `[[` or array assignment.
# Heuristic: line starts (after indent) with `[` then space/flag — classic `[ -f` / `[ "$x"`.
single_bracket_re='^[[:space:]]*\[[[:space:]]'

for file in "${FILES[@]}"; do
  [[ -f "${file}" ]] || fail "missing ${file}"
  if grep -Eq "${single_bracket_re}" "${file}"; then
    fail "${file}: still uses single-bracket [ test (REQ-SONAR-PG-M06 / shelldre:S7688); prefer [["
  fi
  echo "PASS: ${file#$ROOT/} has no single-bracket [ tests"
done

echo "PASS: REQ-SONAR-PG-M06 prefer [["
