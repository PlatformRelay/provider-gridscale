# Change: sonar-maintainability-2026-08 — SonarCloud MAINTAINABILITY + RELIABILITY

## Why
SonarCloud reports **46 open CODE_SMELL** (Maintainability 38 · Reliability 8) on
`PlatformRelay_provider-gridscale` with zero SECURITY/BUG. Batch 11 (E5-S18…S21) clears the
hand-authored clusters: empty `init()`, config test complexity/named types, shell hygiene, and
Dockerfile expansion quoting. Generated `zz_*.go` smells stay deferred.

## What Changes
- Nested comments in empty package `init()` stubs (`apis/{cluster,namespaced}/v1alpha1/register.go`).
- Refactor `config/*_test.go` for cognitive complexity ≤15 and named CRD view types.
- Shell hygiene in `hack/check-*.sh` and selected `hack/test/*` meta scripts.
- Quote Dockerfile expansions (preserve E5-S17 `COPY` for local artefacts).
- Meta tests `hack/test/sonar_pg_m0{1..8}_*` + OpenSpec REQs **PG-M01…PG-M08**.

## Impact
- Affected: hand-authored register stubs, config tests, hack scripts, Dockerfile, OpenSpec.
- No product API/CRD behaviour change intended.
- Out of scope: coverage CI wiring (`sonar-coverage-ci`), generated code smells.
