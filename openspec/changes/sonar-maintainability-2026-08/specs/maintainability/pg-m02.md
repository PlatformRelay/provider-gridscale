# REQ-SONAR-PG-M02: Config test cognitive complexity ≤ 15

**Priority:** must · **Story:** E5-S19 · **Findings:**
`AZ9nqLQ5RbNgprH5pBN9`, `AZ9nqLQ5RbNgprH5pBN-`, `AZ9n-G-75WRcdOwb-njc` (`go:S3776`)

**Given** `TestCRDGoldenContract`, `TestCRDStructuralInvariants`
(`config/crd_contract_test.go`), and `FuzzGetExternalName`
(`config/external_name_fuzz_test.go`)
**When** cognitive complexity is reduced via extracted helpers / early continue
(outcome-focused; goldens and stub contract unchanged)
**Then** each of the three functions reports cognitive complexity **≤ 15** under
Sonar (or an equivalent static check), and:
- `go test ./config/ -count=1` stays green without `UPDATE_GOLDEN=1`
- fuzz still builds: `go test ./config/ -run=FuzzGetExternalName -count=1`
- empty / non-string / missing `id` fuzz cases still enforce never-error + empty → `""`

**Test:** `go test ./config/ -count=1` (characterization; optional
`hack/test/sonar_pg_m02_*` structural meta not required)
**Verify:** `go test ./config/ -count=1 && go test ./config/ -run=FuzzGetExternalName -count=1`
