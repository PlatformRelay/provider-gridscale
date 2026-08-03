# REQ-SONAR-PG-M07 — merge nested ifs in sonar meta tests (S1066)

`sonar_pg_02` / `sonar_pg_04_{ci,e2e}` nested `if` pairs are merged with equivalent fail conditions.
**Test:** `hack/test/sonar_pg_m07_merge_nested_ifs_test.sh`
**Verify:** `bash hack/test/sonar_pg_m07_merge_nested_ifs_test.sh`
