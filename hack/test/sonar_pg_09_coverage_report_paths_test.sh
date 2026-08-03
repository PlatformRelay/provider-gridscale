#!/usr/bin/env bash
# REQ-SONAR-PG-09: CI-based SonarCloud analysis must import the Go coverprofile
# produced by `make coverage` (cover.out). Automatic Analysis cannot ingest Go
# coverage; sonar-project.properties + a sonarcloud CI job are required (kollect /
# mkurator pattern).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
PROPS="${ROOT}/sonar-project.properties"
COV_WF="${ROOT}/.github/workflows/coverage.yml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${PROPS}" ]] || fail "missing ${PROPS} (CI scanner config; Autoscan cannot import Go coverage)"

# Project identity must match the SonarCloud project used by SECURITY remediation.
grep -Eq '^sonar\.projectKey=PlatformRelay_provider-gridscale$' "${PROPS}" \
  || fail "${PROPS}: sonar.projectKey must be PlatformRelay_provider-gridscale"
grep -Eq '^sonar\.organization=platformrelay$' "${PROPS}" \
  || fail "${PROPS}: sonar.organization must be platformrelay"

# Coverage report path must match Makefile `coverage` target output.
grep -Eq '^sonar\.go\.coverage\.reportPaths=cover\.out$' "${PROPS}" \
  || fail "${PROPS}: sonar.go.coverage.reportPaths must be cover.out (make coverage artifact)"

# Preserve CPD exclusions from the Autoscan-era .sonarcloud.properties so the
# new_duplicated_lines_density gate does not regress when switching to CI analysis.
grep -Eq '^sonar\.cpd\.exclusions=.*internal/controller/\*\*/\*\.go' "${PROPS}" \
  || fail "${PROPS}: sonar.cpd.exclusions must keep internal/controller/**/*.go"
grep -Eq '^sonar\.exclusions=.*\*\*/zz_\*\.go' "${PROPS}" \
  || fail "${PROPS}: sonar.exclusions must keep **/zz_*.go"

[[ -f "${COV_WF}" ]] || fail "missing ${COV_WF}"

# coverage.yml must upload cover.out and run a sonarcloud job that consumes it.
grep -Eq 'files:[[:space:]]*cover\.out|name:[[:space:]]*coverage' "${COV_WF}" \
  || fail "${COV_WF}: must reference cover.out / coverage artifact"
grep -Eq 'sonarcloud:' "${COV_WF}" \
  || fail "${COV_WF}: must define a sonarcloud job"
grep -Eq 'sonarqube-scan-action|sonar-scanner' "${COV_WF}" \
  || fail "${COV_WF}: sonarcloud job must invoke the Sonar scanner action"

echo "PASS: sonar-project.properties points at cover.out; coverage.yml has sonarcloud job"
