# REQ-SONAR-PG-M01 — Empty `register.go` `init()` nested comments (go:S1186)

Clear Sonar CRITICAL empty-function smells on the Upjet scaffold `func init() {}`
stubs without changing scheme registration behaviour.

**Sonar keys**
- `apis/cluster/v1alpha1/register.go` — `AZ9magFP_3fq4FMeL0wZ` (go:S1186 CRITICAL)
- `apis/namespaced/v1alpha1/register.go` — `AZ9magJP_3fq4FMeL0xy` (go:S1186 CRITICAL)

**Decision:** D-023 → A (nested comments; do not delete `init`; do not call `AddToScheme`).

## Given / When / Then

**Given** the Upjet scaffold `func init() {}` in
`apis/cluster/v1alpha1/register.go` and `apis/namespaced/v1alpha1/register.go`,
**when** the E5-S18 fix lands,
**then** each `init` body contains a **nested comment** explaining that CRD types
register via `SchemeBuilder` / generated `zz_*.go`, satisfying go:S1186, and
neither stub calls `AddToScheme` (double-register risk).

**Given** `make generate` / `make reviewable` after the edit,
**when** codegen runs,
**then** neither `register.go` is overwritten back to a bare-empty body; if
codegen regenerates them, the intentional nested comment is restored in the same
change.

**Edge — Given** a mistaken fix that calls `SchemeBuilder.AddToScheme` from both
package `init()`s,
**when** controllers start,
**then** types must not double-register; prefer the nested-comment pattern over
wiring registration in these stubs.

## Test / Verify

| | |
| --- | --- |
| **Level** | M (meta) |
| **Test** | `hack/test/sonar_pg_m01_register_init_not_empty_test.sh` |
| **Verify** | `bash hack/test/sonar_pg_m01_register_init_not_empty_test.sh` |
| **Compile** | `go test ./apis/... -count=1` |
