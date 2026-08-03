# REQ-SONAR-PG-M05 — doc-check errors on stderr (S7677)

Failure / `::error::` messages in `hack/check-docs.sh` and `hack/check-api-docs.sh`
MUST go to stderr (`>&2`).
**Test:** `hack/test/sonar_pg_m05_stderr_errors_test.sh`
**Verify:** `bash hack/test/sonar_pg_m05_stderr_errors_test.sh`
