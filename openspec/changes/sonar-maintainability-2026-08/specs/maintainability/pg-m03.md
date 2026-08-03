# REQ-SONAR-PG-M03: Named CRD view types (no anonymous nesting)

**Priority:** must · **Story:** E5-S19 · **Findings:**
`AZ9nqLQ5RbNgprH5pBN_` … `AZ9nqLQ5RbNgprH5pBOE` (`godre:S8205`, 6×)

**Given** the nested anonymous structs inside `type crd` in
`config/crd_contract_test.go` (formerly ~lines 41–53)
**When** each flagged nesting level is replaced with a **named type**
(unexported OK)
**Then** JSON tags and unmarshalling behaviour are identical, and
`go test ./config/ -count=1` without `UPDATE_GOLDEN=1` still passes all golden +
structural invariants (byte-stable goldens for an unchanged CRD tree)

**Test:** `go test ./config/ -count=1`
**Verify:** `go test ./config/ -count=1`
