#!/usr/bin/env bash
# REQ-SONAR-PG-05: Scorecard workflow uses explicit minimal permissions
# (not top-level permissions: read-all). Write scopes for Scorecard publish
# must be job-scoped — top-level write clears read-all but breaks
# ossf/scorecard-action workflow verification.
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

# Extract the top-level permissions block (before jobs:).
toplevel_perms="$(awk '
  /^permissions:/ { in_perm=1; print; next }
  in_perm && /^[^[:space:]#]/ { exit }
  in_perm { print }
' "${WF}")"

[[ -n "${toplevel_perms}" ]] || fail "${WF}: missing top-level permissions: block"

# Top-level must be read-only for Scorecard publish verification.
if echo "${toplevel_perms}" | grep -Eq 'security-events:[[:space:]]*write'; then
  fail "${WF}: top-level 'security-events: write' must be job-scoped (not workflow-global)"
fi
if echo "${toplevel_perms}" | grep -Eq 'id-token:[[:space:]]*write'; then
  fail "${WF}: top-level 'id-token: write' must be job-scoped (not workflow-global)"
fi

# Job analysis must keep the write scopes needed for SARIF + publish.
analysis_perms="$(awk '
  /^[[:space:]]*analysis:/ { in_job=1; next }
  in_job && /^[[:space:]]*permissions:/ { in_perm=1; print; next }
  in_job && in_perm && /^[[:space:]]{4}[^[:space:]#]/ { exit }
  in_job && in_perm { print }
  in_job && /^[[:space:]]{2}[a-zA-Z0-9_-]+:/ && !/^[[:space:]]*analysis:/ { exit }
' "${WF}")"

[[ -n "${analysis_perms}" ]] || fail "${WF}: missing job analysis permissions: block"

for pat in \
  'security-events:[[:space:]]*write' \
  'id-token:[[:space:]]*write' \
  'contents:[[:space:]]*read' \
  'actions:[[:space:]]*read'
do
  if ! echo "${analysis_perms}" | grep -Eq "${pat}"; then
    fail "${WF}: job analysis missing required permission scope matching /${pat}/"
  fi
done

# Top-level should still declare the read scopes Scorecard needs as baseline.
if ! echo "${toplevel_perms}" | grep -Eq 'contents:[[:space:]]*read'; then
  fail "${WF}: top-level must declare contents: read"
fi
if ! echo "${toplevel_perms}" | grep -Eq 'actions:[[:space:]]*read'; then
  fail "${WF}: top-level must declare actions: read"
fi

echo "PASS: scorecard.yml uses read-only top-level + job-scoped write permissions (no read-all)"
