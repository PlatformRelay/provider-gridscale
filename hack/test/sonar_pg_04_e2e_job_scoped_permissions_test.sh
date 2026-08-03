#!/usr/bin/env bash
# REQ-SONAR-PG-04 (e2e half): workflow may set a read-only floor
# (`permissions: contents: read` — a311484), but must not use write-all or any
# non-read permission at workflow level. Jobs that checkout / post statuses
# still declare local permissions including contents: read.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/e2e.yaml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

pre_jobs="$(awk '
  /^[[:space:]]*jobs:[[:space:]]*$/ { exit }
  { print }
' "${WF}")"

# Canonical broad grant.
if echo "${pre_jobs}" | grep -Eq '^[[:space:]]*permissions:[[:space:]]*write-all[[:space:]]*$'; then
  fail "${WF}: workflow-level permissions must not grant write/admin scopes (write-all)"
fi

# Nested permissions block: every granted key must be `: read` (allowlist).
# Reject write/admin and unknown write-class values; allow `contents: read` only.
if echo "${pre_jobs}" | grep -Eq '^[[:space:]]*permissions:[[:space:]]*$'; then
  bad="$(echo "${pre_jobs}" | awk '
    /^[[:space:]]*permissions:[[:space:]]*$/ { in_perm=1; next }
    in_perm && /^[^[:space:]#]/ { exit }
    in_perm && /^[[:space:]]*[a-z0-9_-]+:[[:space:]]*/ {
      line=$0
      sub(/^[[:space:]]*/, "", line)
      split(line, a, /:[[:space:]]*/)
      key=a[1]; val=a[2]
      gsub(/[[:space:]]/, "", val)
      if (key == "") next
      if (!(key == "contents" && val == "read")) {
        print key ": " val
      }
    }
  ')"
  if [[ -n "${bad}" ]]; then
    fail "${WF}: workflow-level permissions must not grant write/admin scopes (only contents: read allowed); got: ${bad}"
  fi
fi

# One-liner non-read forms.
if echo "${pre_jobs}" | grep -Eq '^[[:space:]]*permissions:[[:space:]]+' \
  && ! echo "${pre_jobs}" | grep -Eq '^[[:space:]]*permissions:[[:space:]]*contents:[[:space:]]*read[[:space:]]*$'; then
  # permissions: <something> that is not contents: read
  if echo "${pre_jobs}" | grep -Eq '^[[:space:]]*permissions:[[:space:]]*(write-all|.+:[[:space:]]*(write|admin))'; then
    fail "${WF}: workflow-level permissions must not grant write/admin scopes"
  fi
fi

for job in get-example-list uptest; do
  if ! awk -v job="${job}" '
    $0 ~ "^[[:space:]]*" job ":[[:space:]]*$" { in_job=1; next }
    in_job && /^  [a-zA-Z0-9_-]+:[[:space:]]*$/ && $0 !~ "^  " job ":" { exit }
    in_job && /^[[:space:]]+permissions:[[:space:]]*$/ { has_perm=1 }
    in_job && /contents:[[:space:]]*read/ { has_contents=1 }
    END { exit (has_perm && has_contents) ? 0 : 1 }
  ' "${WF}"; then
    fail "${WF}: job '${job}' must declare local permissions including contents: read"
  fi
done

echo "PASS: e2e.yaml workflow floor is read-only; jobs declare permissions locally"
