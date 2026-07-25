# Spec — SonarCloud SECURITY remediation (2026-07)

Epic: sonar-security-remediation-2026-07 · **Level:** M (meta / structural)
Sonar project: `PlatformRelay_provider-gridscale`

Every REQ below maps 1:1 to a dispositioned finding. Implementation **must**
add the failing **Test:** artifact first.

---

## REQ-SONAR-PG-01: No workflow_dispatch input expansion inside `run:`

**Priority:** must · **Story:** E5-S11 · **Finding:** `AZ9rgOSpjx1_oJFvbrV5`, `AZ9rgOSpjx1_oJFvbrV6` (`githubactions:S7630`)
**Given** `.github/workflows/publish-provider-package.yml` Build/Publish Artifacts steps
**When** an optional `inputs.version` is supplied
**Then** the value is assigned via step `env:` (e.g. `VERSION: ${{ inputs.version }}`) and the
`run:` block references only `$VERSION` / `"${VERSION}"` — never `${{ github.event.inputs.version }}`
or `format('VERSION={0}', …)` inside the script body
**Test:** `hack/test/sonar_pg_01_no_input_interpolation_test.sh`

**Verify:** `bash hack/test/sonar_pg_01_no_input_interpolation_test.sh`

---

## REQ-SONAR-PG-02: Secrets reach the shell only through `env:`

**Priority:** must · **Story:** E5-S12 · **Finding:** `AZ9magKA_3fq4FMeL0xz` (`githubactions:S7636`)
**Given** the e2e uptest step that writes the datasource file
**When** `UPTEST_DATASOURCE` is needed on disk
**Then** the secret is bound as `env.UPTEST_DATASOURCE: ${{ secrets.UPTEST_DATASOURCE }}` and the
`run:` block uses `printf '%s\n' "$UPTEST_DATASOURCE" > …` (no `${{ secrets.* }}` inside `run:`)
**Test:** `hack/test/sonar_pg_02_no_secret_in_run_test.sh`

**Verify:** `bash hack/test/sonar_pg_02_no_secret_in_run_test.sh`

---

## REQ-SONAR-PG-03: Example-list context hash uses SHA-256

**Priority:** should · **Story:** E5-S12 · **Finding:** `AZ-GeyXvbiMwwowBPDZ3` (`githubactions:S4790`)
**Given** `.github/workflows/e2e.yaml` Prepare The Example List step
**When** a multi-example hash is computed for the status-check context
**Then** hashing uses `sha256sum` (or `shasum -a 256`), not `md5sum`
**Test:** `hack/test/sonar_pg_03_sha256_example_hash_test.sh`

**Verify:** `bash hack/test/sonar_pg_03_sha256_example_hash_test.sh`

---

## REQ-SONAR-PG-04: Job-scoped read permissions on ci + e2e

**Priority:** must · **Stories:** E5-S12 (e2e), E5-S13 (ci) · **Finding:** `AZ9s1CAq4e9EUCVmkTAf`, `AZ9s1B-Z4e9EUCVmkTAe` (`githubactions:S8264`)
**Given** `.github/workflows/ci.yml` and `.github/workflows/e2e.yaml`
**When** permissions are inspected
**Then** there is no workflow-level `contents: read`; each job that needs it declares
`permissions:` locally (jobs that need write scopes keep their elevated block)

Coverage is split per file lock so lanes do not fight over one script:

**Test (e2e / E5-S12):** `hack/test/sonar_pg_04_e2e_job_scoped_permissions_test.sh`

**Verify (e2e / E5-S12):** `bash hack/test/sonar_pg_04_e2e_job_scoped_permissions_test.sh`

**Test (ci / E5-S13):** `hack/test/sonar_pg_04_ci_job_scoped_permissions_test.sh`

**Verify (ci / E5-S13):** `bash hack/test/sonar_pg_04_ci_job_scoped_permissions_test.sh`

---

## REQ-SONAR-PG-05: Scorecard workflow uses explicit permissions

**Priority:** must · **Story:** E5-S14 · **Finding:** `AZ9oSbjca22a70H8_MAQ` (`githubactions:S8234`)
**Given** `.github/workflows/scorecard.yml`
**When** top-level `permissions` is read
**Then** it is not `read-all`; it enumerates the minimum scopes Scorecard needs
(typically `security-events: write`, `id-token: write`, `contents: read`, `actions: read`)
**Test:** `hack/test/sonar_pg_05_scorecard_permissions_test.sh`

**Verify:** `bash hack/test/sonar_pg_05_scorecard_permissions_test.sh`

---

## REQ-SONAR-PG-06: curl enforces HTTPS protocol

**Priority:** must · **Stories:** E5-S11 (publish), E5-S15 (gitleaks) · **Finding:** `AZ-GeyUvbiMwwowBPDZ1`, `AZ-GeyVebiMwwowBPDZ2` (`githubactions:S6506`)
**Given** workflow steps that `curl` remote installers/binaries
**When** curl is invoked
**Then** every invocation includes `--proto '=https' --tlsv1.2` (or equivalent curl config)
in addition to `-f` / `-S`

Coverage is split per file lock so lanes do not fight over one script:

**Test (publish / E5-S11):** `hack/test/sonar_pg_06_publish_curl_https_test.sh`

**Verify (publish / E5-S11):** `bash hack/test/sonar_pg_06_publish_curl_https_test.sh`

**Test (gitleaks / E5-S15):** `hack/test/sonar_pg_06_gitleaks_curl_https_test.sh`

**Verify (gitleaks / E5-S15):** `bash hack/test/sonar_pg_06_gitleaks_curl_https_test.sh`

---

## REQ-SONAR-PG-07: version_diff.py bounds CLI paths to the repo

**Priority:** must · **Story:** E5-S16 · **Finding:** `AZ9magKT_3fq4FMeL0x1`…`L0x3` (`pythonsecurity:S8707`)
**Given** `scripts/version_diff.py`
**When** argv path arguments are opened
**Then** each path is resolved and rejected unless it is under the repository root
(no `../` escape); unit coverage asserts rejection of out-of-root paths
**Test:** `scripts/version_diff_test.py`

**Verify:** `python3 -m pytest scripts/version_diff_test.py -q`

---

## REQ-SONAR-PG-08: Dockerfile uses COPY for local artifacts

**Priority:** must · **Story:** E5-S17 · **Finding:** `AZ9magKK_3fq4FMeL0x0` (`docker:S7029`)
**Given** `cluster/images/provider-gridscale/Dockerfile`
**When** local files (`bin/…/provider`, `terraformrc.hcl`) are added to the image
**Then** the instruction is `COPY`, not `ADD`. Remote zip fetches may remain `ADD <url>`
or be replaced with `curl --proto '=https'`
**Test:** `hack/test/sonar_pg_08_dockerfile_copy_test.sh`

**Verify:** `bash hack/test/sonar_pg_08_dockerfile_copy_test.sh`
