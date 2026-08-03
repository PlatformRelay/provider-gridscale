#!/usr/bin/env bash
# REQ-SONAR-PG-M07: sonar_pg_02 / sonar_pg_04_* must not keep Sonar-flagged nested if
# pairs (shelldre:S1066). Merge with &&; preserve fail messages/conditions.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

# Detect a classic nested-if smell: an `if …; then` whose next non-blank/non-comment
# line is another `if …; then` at deeper indent (the S1066 pattern Sonar flagged).
has_nested_if_pair() {
  local file="$1"
  awk '
    /^[[:space:]]*#/ { next }
    /^[[:space:]]*$/ { next }
    {
      line = $0
      match(line, /^[[:space:]]*/)
      indent = RLENGTH
      is_if = (line ~ /^[[:space:]]*if[[:space:]]/)
      if (prev_is_if && is_if && indent > prev_indent) {
        found = 1
        exit
      }
      prev_is_if = is_if
      prev_indent = indent
    }
    END { exit found ? 0 : 1 }
  ' "${file}"
}

FILES=(
  "${ROOT}/hack/test/sonar_pg_02_no_secret_in_run_test.sh"
  "${ROOT}/hack/test/sonar_pg_04_ci_job_scoped_permissions_test.sh"
  "${ROOT}/hack/test/sonar_pg_04_e2e_job_scoped_permissions_test.sh"
)

for file in "${FILES[@]}"; do
  [[ -f "${file}" ]] || fail "missing ${file}"
  if has_nested_if_pair "${file}"; then
    fail "${file}: nested if pair still present (REQ-SONAR-PG-M07 / shelldre:S1066); merge with &&"
  fi

  # Security / permission fail strings must remain (do not weaken greps).
  case "${file}" in
    *sonar_pg_02*)
      grep -Fq 'secrets.UPTEST_DATASOURCE must not appear inside run:' "${file}" \
        || fail "${file}: missing secret-in-run fail message"
      grep -Eq '\$\{\{[[:space:]]*secrets\.UPTEST_DATASOURCE' "${file}" \
        || fail "${file}: missing secrets.UPTEST_DATASOURCE grep"
      ;;
    *sonar_pg_04_ci*)
      grep -Fq "workflow-level 'contents: read' must be removed" "${file}" \
        || fail "${file}: missing workflow-level contents: read fail message"
      ;;
    *sonar_pg_04_e2e*)
      # e2e allows read-only workflow floor (a311484); still forbid write/admin.
      grep -Eq "workflow-level permissions must not grant write/admin scopes" "${file}" \
        || fail "${file}: missing workflow write/admin forbid message"
      ;;
  esac

  echo "PASS: ${file#$ROOT/} nested ifs merged; fail semantics retained"
done

echo "PASS: REQ-SONAR-PG-M07 merge nested ifs"
