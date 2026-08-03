# Design: sonar-maintainability-2026-08

## Approach
File-disjoint lanes (L-REG / L-CFGTEST / L-SHELL / L-DOCKER) land independently under
agent-loop-local. Each REQ has a meta **Test:** under `hack/test/sonar_pg_m*` and a **Verify:**
bash command. Prefer Sonar-accepted minimal fixes (nested comments, `[[`, `>&2`, quoted expansions)
over behaviour-changing refactors.

## Decisions
- **D-023 → A:** nested comments in empty `init()` (not AddToScheme).
- Preserve E5-S17 Dockerfile `COPY` for local provider/terraformrc artefacts.
