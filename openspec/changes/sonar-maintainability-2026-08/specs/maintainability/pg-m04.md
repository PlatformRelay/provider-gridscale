# REQ-SONAR-PG-M04 — check-docs case has default arm (S131)

`hack/check-docs.sh` case over CRD basenames MUST include a `*)` default arm.

## Related: e2e half of REQ-SONAR-PG-04 (D-025)
`hack/test/sonar_pg_04_e2e_*` allows a workflow-level **read-only floor**
(`permissions: contents: read` per a311484) and MUST reject `permissions: write-all`
and any workflow-level non-read grant. Jobs `get-example-list` / `uptest` still
declare local `contents: read`.

**Test:** `hack/test/sonar_pg_m04_case_default_test.sh`
**Verify:** `bash hack/test/sonar_pg_m04_case_default_test.sh`
