# REQ-SONAR-PG-M01 — Empty register `init()` stubs are non-bare-empty

## Requirement
Package `init()` stubs in `apis/cluster/v1alpha1/register.go` and
`apis/namespaced/v1alpha1/register.go` MUST NOT be bare-empty `func init() {}`
(Sonar go:S1186). Each body MUST contain a nested comment explaining that types
register via `SchemeBuilder` / generated `zz_*.go`. Callers MUST NOT wire
`AddToScheme` from these stubs (double-register risk). **D-023 → A.**

## Scenario
Given the Upjet scaffold empty `init()` functions, when the fix lands, then each
`init` body includes an intentional nested comment and remains a no-op.

## Test / Verify
- **Level:** M
- **Test:** `hack/test/sonar_pg_m01_register_init_not_empty_test.sh`
- **Verify:** `bash hack/test/sonar_pg_m01_register_init_not_empty_test.sh`
