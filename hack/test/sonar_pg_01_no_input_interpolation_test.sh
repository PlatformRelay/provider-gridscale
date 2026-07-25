#!/usr/bin/env bash
# REQ-SONAR-PG-01: publish workflow must not expand github.event.inputs.version
# (or format('VERSION=...')) inside run: script bodies. Assign via step env: and
# use $VERSION / "${VERSION}" in the shell instead.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/publish-provider-package.yml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

# Extract run: script bodies (heuristic) and flag forbidden input interpolation.
# if: / with: / env: lines outside run bodies are intentionally ignored.
awk -v fmt_needle="format('VERSION=" '
  BEGIN { in_run = 0; run_indent = -1; bad = 0 }

  # Single-line run: <cmd>
  /^[[:space:]]*run:[[:space:]]+[^|>[:space:]]/ {
    body = $0
    sub(/^[[:space:]]*run:[[:space:]]+/, "", body)
    check(body, NR)
    in_run = 0
    next
  }

  # Block scalar run: | or run: >
  /^[[:space:]]*run:[[:space:]]*[|>]/ {
    match($0, /^[[:space:]]*/)
    run_indent = RLENGTH
    in_run = 1
    next
  }

  in_run {
    if ($0 ~ /^[[:space:]]*$/) next
    match($0, /^[[:space:]]*/)
    ind = RLENGTH
    if (ind <= run_indent) {
      in_run = 0
      next
    }
    check($0, NR)
    next
  }

  function check(line, n) {
    if (line ~ /[$][{][{][[:space:]]*github[.]event[.]inputs[.]version/) {
      printf "line %d: github.event.inputs.version inside run body\n", n > "/dev/stderr"
      printf "  %s\n", line > "/dev/stderr"
      bad = 1
    }
    if (index(line, fmt_needle) > 0) {
      printf "line %d: format(VERSION=...) inside run body\n", n > "/dev/stderr"
      printf "  %s\n", line > "/dev/stderr"
      bad = 1
    }
  }

  END { exit bad ? 1 : 0 }
' "${WF}" && status=0 || status=$?

if [[ "${status}" -ne 0 ]]; then
  fail "${WF}: forbidden input interpolation inside run: script body (REQ-SONAR-PG-01)"
fi

echo "PASS: publish-provider-package.yml has no inputs.version interpolation inside run: bodies"
