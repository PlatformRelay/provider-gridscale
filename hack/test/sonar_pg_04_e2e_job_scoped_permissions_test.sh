#!/usr/bin/env bash
# REQ-SONAR-PG-04 (e2e half): no workflow-level contents: read; jobs that need
# it declare permissions: locally.
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

# Workflow-level permissions block must not grant contents: read.
if echo "${pre_jobs}" | grep -Eq '^[[:space:]]*permissions:[[:space:]]*$'; then
  # Collect indented keys under that permissions: until a non-indented / less-indented key.
  if echo "${pre_jobs}" | awk '
    /^[[:space:]]*permissions:[[:space:]]*$/ { in_perm=1; next }
    in_perm && /^[^[:space:]#]/ { exit }
    in_perm && /^[[:space:]]*contents:[[:space:]]*read[[:space:]]*$/ { found=1; exit }
    END { exit found ? 0 : 1 }
  '; then
    fail "${WF}: workflow-level 'contents: read' must be removed; declare permissions on jobs"
  fi
fi

# Also catch the one-liner form: permissions: contents: read (unusual but possible).
if echo "${pre_jobs}" | grep -Eq '^[[:space:]]*permissions:[[:space:]]*contents:[[:space:]]*read'; then
  fail "${WF}: workflow-level 'permissions: contents: read' must be removed"
fi

# Jobs that check out the repo / use gh pr checkout need a local permissions block
# with contents: read. get-example-list and uptest already elevate for statuses.
for job in get-example-list uptest; do
  if ! awk -v job="${job}" '
    $0 ~ "^[[:space:]]*" job ":[[:space:]]*$" { in_job=1; next }
    in_job && /^[[:space:]]+[a-zA-Z0-9_-]+:[[:space:]]*$/ && $0 !~ /^[[:space:]]{2}steps:/ {
      # next sibling job (2-space indent under jobs:)
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

echo "PASS: e2e.yaml has no workflow-level contents: read; jobs declare permissions locally"
