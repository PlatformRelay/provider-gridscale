# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
Release notes are generated from [Conventional Commits](https://www.conventionalcommits.org/)
(with a leading gitmoji token) on the default branch using
[git-cliff](https://git-cliff.org/).

## [Unreleased]

### Bug Fixes

- **lint:** Scope gosec G101/G104 suppression to tests only (SEC-1) [04f4527](https://github.com/platformrelay/provider-gridscale/commit/04f452789477c641a84e59802db8563fa4096d1b)

- **deps:** Bump go 1.26.5 + x/net v0.55.0 (govulncheck GO-2026-5856/5026/4918) ([#10](https://github.com/platformrelay/provider-gridscale/pull/10))[d75721e](https://github.com/platformrelay/provider-gridscale/commit/d75721e674b5ec130f9e2af62b335d8ecedefa03)

- **quality:** Exclude .claude/.work from arch-lint; gofmt + ascii-escape fuzz seed [0051304](https://github.com/platformrelay/provider-gridscale/commit/0051304473d62d529adf1283aab49f60bfd9c82d)


### Build System

- **quality:** Add vuln/race/tidy/arch-lint make targets (E2-S06/S07/S08/S10) [418957f](https://github.com/platformrelay/provider-gridscale/commit/418957fa7fb9e78719fb879eea3bd9edf56b886b)


### Documentation

- **coordination:** Consolidate operator decision set + reconcile govulncheck (D-007..D-012) [ab8a86e](https://github.com/platformrelay/provider-gridscale/commit/ab8a86e53b4913349f8946d8bec545c661b3f818)

- **readme:** Add supported-resource matrix + community links (E3-S02) [ac550e0](https://github.com/platformrelay/provider-gridscale/commit/ac550e0db53f8cacab677704c0a62aa97e51bf8e)

- **api:** Generate CRD API reference via crd-ref-docs (E4-S01) [942b123](https://github.com/platformrelay/provider-gridscale/commit/942b123934a2e4f0093f80883e32b523a24d59cc)

- **governance:** Add GOVERNANCE model (E6-S03; maintainers TBD) [e486acc](https://github.com/platformrelay/provider-gridscale/commit/e486acce58a83fc1661786d946b787d39a1945d5)

- **examples:** Curated apply-able examples per core resource (E4-S02) [f0120ef](https://github.com/platformrelay/provider-gridscale/commit/f0120ef8924daec54ac2f0324a0f06eaaa80b66f)

- **inbox:** Record 3 govulncheck CVEs as surfaced security decision (B2-01) [b8fa9e9](https://github.com/platformrelay/provider-gridscale/commit/b8fa9e99a792d9af992c6b1b5a900006c78299bd)

- **backlog:** Add E2-S06..S10 test-hardening stories; reconcile board + register (batch-1 integrated, CRED live-validated) [0f04e25](https://github.com/platformrelay/provider-gridscale/commit/0f04e258d309468590be3167c476ae25d091dc0f)

- **adr:** Add codegen-boundary + build-tooling ADRs (E4-S03) [5804810](https://github.com/platformrelay/provider-gridscale/commit/5804810b8496be1c1bcba69a3b1cf2056c81ddde)

- **community:** Contributor Covenant 2.1 + issue/PR templates (E6-S04) [7c5aacf](https://github.com/platformrelay/provider-gridscale/commit/7c5aacfaccf84fee2085cf89868a12c33fc54dd6)

- **security:** Add SECURITY policy + gitignore local creds (E6-S02) [8e9f682](https://github.com/platformrelay/provider-gridscale/commit/8e9f68272832547a5df34ad254939b50e211b47c)

- **contributing:** Add CONTRIBUTING guide (E6-S01) [694e3ba](https://github.com/platformrelay/provider-gridscale/commit/694e3ba3c5336fca279b839126ecda6f8034dc09)


### Features

- **branding:** Add provider icon + light/dark variants (E3-S01, placeholder) [df4a3ec](https://github.com/platformrelay/provider-gridscale/commit/df4a3ecf86f6440e5dbb9d8d7534a5129d166c35)


### Miscellaneous

- **deps:** Promote sigs.k8s.io/yaml to direct (go mod tidy; E2-S10) [a75053c](https://github.com/platformrelay/provider-gridscale/commit/a75053c174bb1ee0e6f10589c7798c22ec1b8a6b)


### Tests

- **config:** Coverage floor target scoped to ./config (E2-S03) [2faf26c](https://github.com/platformrelay/provider-gridscale/commit/2faf26c2e62d7a5f1ede507dbfaecfa65f26cdab)

- **config:** Fuzz external-name extraction (E2-S09) [d2ecd45](https://github.com/platformrelay/provider-gridscale/commit/d2ecd45c7640048d0e52f5ba81eaae637a7f772f)

- **config:** CRD golden-contract tests for scoped subset (E2-S02) [73a3410](https://github.com/platformrelay/provider-gridscale/commit/73a3410a4f08166fa0204b913176afca4e66cd40)

## [0.1.0-alpha.1] - 2026-07-15

### Bug Fixes

- **clients:** Wire gridscale uuid/token credentials [b16b5a2](https://github.com/platformrelay/provider-gridscale/commit/b16b5a21dcd108d486987fb9fdebfed7bcd0b4cc)

- **build:** Fetch gridscale provider from GitHub releases, not releases.hashicorp.com ([#1](https://github.com/platformrelay/provider-gridscale/pull/1))[797adad](https://github.com/platformrelay/provider-gridscale/commit/797adad4de6e6dcf5126e0c33c2bba5062cb3791)

- **config:** Gofmt provider.go to unblock lint on main [6904277](https://github.com/platformrelay/provider-gridscale/commit/6904277370dabd15ef89a53938ed55a6d236c342)


### CI/CD

- **publish:** Use GHCR_PAT to create the org-level ghcr package (#9) ([#9](https://github.com/platformrelay/provider-gridscale/pull/9))[ac5fe76](https://github.com/platformrelay/provider-gridscale/commit/ac5fe7630e03129111f0440ab77767221bad7734)

- **publish:** Target platformrelay org for ghcr + Upbound mirror (#8) ([#8](https://github.com/platformrelay/provider-gridscale/pull/8))[9a36da9](https://github.com/platformrelay/provider-gridscale/commit/9a36da91d029eca6448b031fee74e7a27cc51876)


### Documentation

- **examples:** Correct ProviderConfig API group to platformrelay.io ([#5](https://github.com/platformrelay/provider-gridscale/pull/5))[64b0d13](https://github.com/platformrelay/provider-gridscale/commit/64b0d13315760a2c1a5e1daa47a4f2d93e5e71ec)

- **examples:** Use uuid/token keys in providerconfig secret templates [447dbf6](https://github.com/platformrelay/provider-gridscale/commit/447dbf647f53dedab02f28c71fcbf0bcb0890795)

- **readme:** Replace upjet template with gridscale provider README ([#6](https://github.com/platformrelay/provider-gridscale/pull/6))[51e4cad](https://github.com/platformrelay/provider-gridscale/commit/51e4cad1e5e01711db15d675be93dc16534a7a4b)

- **coordination:** Add OPERATOR-BOARD for E-lane auto-loop batch ([#3](https://github.com/platformrelay/provider-gridscale/pull/3))[3ada70c](https://github.com/platformrelay/provider-gridscale/commit/3ada70c38608c0814192f52d2aab6b5f53e8bd36)

- **audit:** Baseline health & direction audit (2026-07-15) [2d80152](https://github.com/platformrelay/provider-gridscale/commit/2d801527e33a2f50b6594e57ea88f652391f8183)

- Record D-006 (accepted A); resolve stale decision-state (review F1/F2) ([#2](https://github.com/platformrelay/provider-gridscale/pull/2))[00e1c6c](https://github.com/platformrelay/provider-gridscale/commit/00e1c6cda900175e89ecaf1bf2655af4ecb3b6a7)


### Features

- Generate all 32 gridscale resources (E6g-S01) [b95c469](https://github.com/platformrelay/provider-gridscale/commit/b95c4693b4e63d0a0253c8afd37b253365769455)


### Miscellaneous

- **e1:** Port committed spec-driven harness (E1 foundation) [946f99f](https://github.com/platformrelay/provider-gridscale/commit/946f99fe066c884353de3a874399f6938ca188eb)

- Scaffold Upjet provider-gridscale (gridscale/gridscale v2.3.0) [30d59e0](https://github.com/platformrelay/provider-gridscale/commit/30d59e0056866cbed5bfd2b013440e0e3e96bdaa)


### Tests

- **config:** Assert external-name strategy + mysql8/storageimport kind overrides ([#4](https://github.com/platformrelay/provider-gridscale/pull/4))[c53fded](https://github.com/platformrelay/provider-gridscale/commit/c53fded7bc095b09b634a3cd279c3af19ed58971)

