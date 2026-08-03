# REQ-SONAR-PG-M04 — check-docs case has default arm (S131)

`hack/check-docs.sh` case over CRD basenames MUST include a `*)` default arm.
**Test:** `hack/test/sonar_pg_m04_case_default_test.sh`
**Verify:** `bash hack/test/sonar_pg_m04_case_default_test.sh`
