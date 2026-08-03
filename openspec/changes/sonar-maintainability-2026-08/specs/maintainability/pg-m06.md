# REQ-SONAR-PG-M06 — prefer `[[` over `[` (S7688)

Listed `hack/check-docs.sh` and `hack/test/e8_s0{1,2,3}_*` predicates use bash `[[`.
**Test:** `hack/test/sonar_pg_m06_prefer_double_bracket_test.sh`
**Verify:** `bash hack/test/sonar_pg_m06_prefer_double_bracket_test.sh`
