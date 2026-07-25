#!/usr/bin/env bash
# REQ-SONAR-PG-06 (publish site): curl that installs `up` must pin TLS via
# --proto '=https' --tlsv1.2.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/publish-provider-package.yml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

# Locate the curl that fetches the Upbound `up` binary.
if ! grep -Eq 'curl[[:space:]].*upbound\.io|cli\.upbound\.io.*/up' "${WF}"; then
  fail "${WF}: no curl install of up (cli.upbound.io) found to verify"
fi

# Require both flags somewhere on the curl invocation (may span continued lines).
# Collapse backslash-continued lines around the upbound curl for a stable check.
collapsed="$(
  awk '
    BEGIN { buf = "" }
    {
      line = $0
      sub(/[[:space:]]*\\[[:space:]]*$/, "", line)
      if (buf != "") buf = buf " " line
      else buf = line
      if ($0 ~ /\\[[:space:]]*$/) next
      if (buf ~ /curl/ && buf ~ /upbound\.io/) print buf
      buf = ""
    }
  ' "${WF}"
)"

[[ -n "${collapsed}" ]] || fail "${WF}: could not isolate curl+upbound.io invocation"

proto_ok=0
tls_ok=0
while IFS= read -r inv; do
  [[ -z "${inv}" ]] && continue
  if [[ "${inv}" == *"--proto '=https'"* ]] || [[ "${inv}" == *'--proto "=https"'* ]]; then
    proto_ok=1
  fi
  if [[ "${inv}" == *"--tlsv1.2"* ]]; then
    tls_ok=1
  fi
done <<< "${collapsed}"

if [[ "${proto_ok}" -ne 1 ]]; then
  fail "${WF}: up install curl missing --proto '=https'"
fi
if [[ "${tls_ok}" -ne 1 ]]; then
  fail "${WF}: up install curl missing --tlsv1.2"
fi

echo "PASS: publish-provider-package.yml up install curl uses --proto '=https' --tlsv1.2"
