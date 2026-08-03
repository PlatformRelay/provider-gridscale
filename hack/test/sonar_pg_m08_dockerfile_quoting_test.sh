#!/usr/bin/env bash
# REQ-SONAR-PG-M08: Dockerfile Setup Terraform expansions must be quoted (docker:S6570).
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
DOCKERFILE="${ROOT}/cluster/images/provider-gridscale/Dockerfile"

if [[ ! -f "${DOCKERFILE}" ]]; then
  echo "FAIL: Dockerfile not found at ${DOCKERFILE}" >&2
  exit 1
fi

failures=0

assert_contains() {
  local label="$1"
  local pattern="$2"

  if grep -F -- "${pattern}" "${DOCKERFILE}" >/dev/null; then
    echo "PASS: ${label}"
  else
    echo "FAIL: missing quoted form — ${label}" >&2
    echo "       expected substring: ${pattern}" >&2
    failures=$((failures + 1))
  fi
}

assert_absent() {
  local label="$1"
  local pattern="$2"

  if grep -F -- "${pattern}" "${DOCKERFILE}" >/dev/null; then
    echo "FAIL: unquoted expansion still present — ${label}" >&2
    echo "       forbidden substring: ${pattern}" >&2
    failures=$((failures + 1))
  else
    echo "PASS: no unquoted ${label}"
  fi
}

# --- Required quoted forms (Setup Terraform block ≈ lines 27–39) ---
assert_contains 'mkdir PLUGIN_DIR' 'mkdir -p "${PLUGIN_DIR}"'

assert_contains 'ADD terraform zip URL' \
  'ADD "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip" /tmp'

assert_contains 'ADD provider zip URL' \
  'ADD "${TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip" /tmp'

assert_contains 'COPY terraformrc dest' 'COPY terraformrc.hcl "${TF_CLI_CONFIG_FILE}"'

assert_contains 'unzip terraform zip' \
  'unzip "/tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip"'

assert_contains 'rm terraform zip' \
  'rm "/tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip"'

assert_contains 'unzip provider zip' \
  'unzip "/tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip"'

assert_contains 'unzip -d PLUGIN_DIR' '-d "${PLUGIN_DIR}"'

assert_contains 'chmod PLUGIN_DIR glob' 'chmod +x "${PLUGIN_DIR}"/*'

assert_contains 'rm provider zip' \
  'rm "/tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip"'

assert_contains 'chown USER_ID' 'chown -R "${USER_ID}:${USER_ID}" /terraform'

# --- Forbidden unquoted forms that Sonar flags ---
assert_absent 'mkdir PLUGIN_DIR' 'mkdir -p ${PLUGIN_DIR}'
assert_absent 'ADD terraform zip (unquoted URL)' \
  'ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp'
assert_absent 'COPY terraformrc unquoted dest' 'COPY terraformrc.hcl ${TF_CLI_CONFIG_FILE}'
assert_absent 'unzip terraform unquoted' \
  'unzip /tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip'
assert_absent 'rm terraform unquoted' \
  'rm /tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip'
assert_absent 'unzip -d PLUGIN_DIR unquoted' '-d ${PLUGIN_DIR}'
assert_absent 'chmod PLUGIN_DIR unquoted' 'chmod +x ${PLUGIN_DIR}/*'
assert_absent 'chown USER_ID unquoted' 'chown -R ${USER_ID}:${USER_ID} /terraform'

if [[ "${failures}" -ne 0 ]]; then
  echo "REQ-SONAR-PG-M08: ${failures} failure(s)" >&2
  exit 1
fi

echo "REQ-SONAR-PG-M08: all checks passed"
