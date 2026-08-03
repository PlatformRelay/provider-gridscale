#!/usr/bin/env bash
# REQ-SONAR-PG-04 (e2e half): workflow may set a read-only floor
# (`permissions: contents: read` — see a311484), but must not grant broader
# write scopes at workflow level. Jobs that checkout / post statuses still
# declare local permissions including contents: read.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/e2e.yaml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

# Extract everything before the top-level `jobs:` key.
pre_jobs="$(awk '
  /^[[:space:]]*jobs:[[:space:]]*$/ { exit }
  { print }
' "${WF}")"

# Workflow-level permissions must not grant write-class scopes.
if echo "${pre_jobs}" | grep -Eq '^[[:space:]]*(contents|pull-requests|issues|packages|id-token):[[:space:]]*(write|admin)'; then
  fail "${WF}: workflow-level permissions must not grant write/admin scopes"
fi
if echo "${pre_jobs}" | awk '
  /^[[:space:]]*permissions:[[:space:]]*$/ { in_perm=1; next }
  in_perm && /^[^[:space:]#]/ { exit }
  in_perm && /^[[:space:]]*(contents|pull-requests|issues|packages|id-token):[[:space:]]*(write|admin)[[:space:]]*$/ { found=1; exit }
  END { exit found ? 0 : 1 }
'; then
  fail "${WF}: workflow-level permissions must not grant write/admin scopes"
fi

# Jobs that check out the repo / use gh pr checkout need a local permissions block
# with contents: read. get-example-list and uptest already elevate for statuses.
for job in get-example-list uptest; do
  if ! awk -v job="${job}" '
    $0 ~ "^[[:space:]]*" job ":[[:space:]]*$" { in_job=1; next }
    in_job && /^[[:space:]]+[a-zA-Z0-9_-]+:[[:space:]]*$/ && $0 !~ /^[[:space:]]{2}steps:/ {
      if (match($0, /^  [a-zA-Z0-9_-]+:[[:space:]]*$/)) { exit }
    }
    in_job && /^  [a-zA-Z0-9_-]+:[[:space:]]*$/ && $0 !~ "^  " job ":" { exit }
    in_job && /^[[:space:]]+permissions:[[:space:]]*$/ { has_perm=1 }
    in_job && /contents:[[:space:]]*read/ { has_contents=1 }
    END { exit (has_perm && has_contents) ? 0 : 1 }
  ' "${WF}"; then
    fail "${WF}: job '${job}' must declare local permissions including contents: read"
  fi
done

echo "PASS: e2e.yaml workflow floor is read-only; jobs declare permissions locally"
