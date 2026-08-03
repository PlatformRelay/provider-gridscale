# REQ-SONAR-PG-M08 — Dockerfile quote expansions (docker:S6570)

Clear Sonar MAJOR Maintainability findings on unquoted shell expansions in the
provider image Setup Terraform block without regressing the E5-S17 local-`COPY`
contract.

**Sonar keys / rule**
- `cluster/images/provider-gridscale/Dockerfile:27–39` — docker:S6570 ×19
  keys `AZ-GeyaXbiMwwowBPDZ4` … `AZ-GeyaXbiMwwowBPDaK`

**Lane:** L-DOCKER / E5-S21

## Given / When / Then

**Given** the Setup Terraform block (~lines 27–39) expanding `${PLUGIN_DIR}`,
`${TERRAFORM_*}`, `${TARGETOS}`, `${TARGETARCH}`, `${USER_ID}`, zip paths, and
`${TF_CLI_CONFIG_FILE}`,
**when** the E5-S21 fix lands,
**then** each flagged expansion is double-quoted per docker:S6570
(`mkdir -p "${PLUGIN_DIR}"`, quoted `ADD` URLs, quoted zip paths in `RUN`,
`chmod +x "${PLUGIN_DIR}"/*`, `chown -R "${USER_ID}:${USER_ID}"`, etc.).

**Given** the E5-S17 / REQ-SONAR-PG-08 contract,
**when** this lane finishes,
**then** local artefacts still use **`COPY`** (not `ADD`) for
`bin/${TARGETOS}_${TARGETARCH}/provider` and `terraformrc.hcl`; remote URL
`ADD` may remain (quoted).

**Edge — Given** a quoted Dockerfile build,
**when** `docker build` / CI image build runs,
**then** terraform + provider plugin still land under `${PLUGIN_DIR}` and the
non-root `USER` still applies (no path-split regressions from quoting).

## Test / Verify

| | |
| --- | --- |
| **Level** | M (meta) |
| **Test** | `hack/test/sonar_pg_m08_dockerfile_quoting_test.sh` |
| **Verify** | `bash hack/test/sonar_pg_m08_dockerfile_quoting_test.sh` |
| **Regression** | `bash hack/test/sonar_pg_08_dockerfile_copy_test.sh` |

**Out of scope:** base-image bumps; `TERRAFORM_VERSION` pin changes; coverage CI.
