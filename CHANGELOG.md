# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
Release notes are generated from [Conventional Commits](https://www.conventionalcommits.org/)
(with a leading gitmoji token) on the default branch using
[git-cliff](https://git-cliff.org/).

## [Unreleased]

### Bug Fixes

- **ci:** Accept up v0.51 per-extension annotations in D-020-FU verify [43294ed](https://github.com/platformrelay/provider-gridscale/commit/43294edb7cb5e32de9e309e17e6d07d3fbc0cee4)


### Documentation

- Document protect-main solo Branch-Protection [843ea82](https://github.com/platformrelay/provider-gridscale/commit/843ea823ef54d8643337315b9f1a5c887b10108a)

- **coordination:** Close v0.2.0 release — signed publish green [77be16d](https://github.com/platformrelay/provider-gridscale/commit/77be16dbd60906933c1bdcf715d8de1611e6c9ae)

## [0.2.0](https://github.com/platformrelay/provider-gridscale/compare/v0.1.1..v0.2.0) - 2026-07-17

### Bug Fixes

- **config:** Make loadbalancer status Computed to stop perpetual diff (LB-1) [c38e52b](https://github.com/platformrelay/provider-gridscale/commit/c38e52baabb9f1a8c37142d26e8df475d1c77924)

- **config:** Mark upstream-missed S3/console credentials sensitive (U-1) [8ae7376](https://github.com/platformrelay/provider-gridscale/commit/8ae7376bf34b5a862c57e342ee8100dc01aad0fb)

- **ci:** Restore green main — filesystem CRD generate + metadatafix lint [d8433f5](https://github.com/platformrelay/provider-gridscale/commit/d8433f5942b909575c552d08be057a78fb410938)


### CI/CD

- **publish:** Self-diagnosing D-020-FU verify + name dispatch watch-points [60b9e8f](https://github.com/platformrelay/provider-gridscale/commit/60b9e8f659a4c1f65339e955173c4f140f798d95)

- **publish:** Append extensions before cosign sign (D-020-FU) [2e15c3e](https://github.com/platformrelay/provider-gridscale/commit/2e15c3e916094ca6ac08bf8a1c44e3d6d1a1398d)

- **security:** Least-privilege GITHUB_TOKEN on ci.yml + e2e.yaml (SEC-H1/2/3) [68f8d66](https://github.com/platformrelay/provider-gridscale/commit/68f8d6683f6178d4ba420e820b4d711ea77d01c3)

- **checks:** Wire check-docs into check-diff; document GOTOOLCHAIN base [d21087f](https://github.com/platformrelay/provider-gridscale/commit/d21087ff2afa14f33a0d6519bd9f5b5bc9906477)

- **docs:** Add check-api-docs sync gate (E4-S01) [58f1103](https://github.com/platformrelay/provider-gridscale/commit/58f11036155c0749dd9c3bcaaf22381568954a7e)


### Documentation

- **readme:** OpenSSF Scorecard badge, drop Go Report, prep v0.2.0 [1799637](https://github.com/platformrelay/provider-gridscale/commit/1799637fcc822baeb5bb5e78322a8531d152ef2b)

- **coordination:** D-020-FU fixed (2e15c3e) — ready for signed v0.2.0 dispatch [acbfb48](https://github.com/platformrelay/provider-gridscale/commit/acbfb4878517b6928770eb314b483345527b8b93)

- **coordination:** 2026-07-17 handover — U-1/LB-1 landed, upstream PRs, D-020-FU [b4782dd](https://github.com/platformrelay/provider-gridscale/commit/b4782dd306f7fe0822c39da6ba11f3461e35fa7c)

- **api:** Sync CRD reference after U-1 sensitivity fix [9e6ca3c](https://github.com/platformrelay/provider-gridscale/commit/9e6ca3c0b52d0183a46c3050d458537941e653e1)

- **decisions:** Apply operator answers D-018/019/020 + track sweep findings [949d1de](https://github.com/platformrelay/provider-gridscale/commit/949d1de4d16d5100b1cf612bf5a3562275bf656d)

- **inbox:** Surface upstream U-1 sensitive-flag security finding [f9895b8](https://github.com/platformrelay/provider-gridscale/commit/f9895b82b1d26d0d3cdd59a30b0b5d6d8705168d)

- **audit:** 07-16b re-audit disposition + ROADMAP refresh [ca9d3aa](https://github.com/platformrelay/provider-gridscale/commit/ca9d3aac966e9bed31f0a17370d519513f52f278)

- **coordination:** Finalize wrap tip references [a1748f3](https://github.com/platformrelay/provider-gridscale/commit/a1748f370a251056590e57ba0856981e8e01ce67)

- **coordination:** Wrap — CI green, handoff, board, inbox [8976f55](https://github.com/platformrelay/provider-gridscale/commit/8976f55ca1c0a8c34b6d6c9b6a458f7c6e67c00f)

- **api:** Sync Filesystem description after metadata stub [d275897](https://github.com/platformrelay/provider-gridscale/commit/d2758973534d2d10446fd258dcfd394a18805254)

- **coordination:** Session wrap — handoff, board, backlog, inbox [c9eff5b](https://github.com/platformrelay/provider-gridscale/commit/c9eff5bddb4cd9f9de2d68e767e3fbebb7eaaf5a)

- **security:** Assurance case + refresh supply-chain posture (E6-S05) [3edc5f0](https://github.com/platformrelay/provider-gridscale/commit/3edc5f0eeb57909bb698026720f663d0616d94b7)

- **examples:** Sync curated index + ApplicationImport (E4-S02) [fdbf300](https://github.com/platformrelay/provider-gridscale/commit/fdbf3002388dc162124e81362b48d7bdaaf582e8)

- **api:** Regenerate CRD reference for sync gate (E4-S01) [2384fe4](https://github.com/platformrelay/provider-gridscale/commit/2384fe44993874e476f6df1b70e8ebafdeef6f14)

- **readme:** Prefer v0.1.1 package tag in install hint [bcc0ba0](https://github.com/platformrelay/provider-gridscale/commit/bcc0ba0ff8f10c06bbdd226293bc319d2dc75667)

- **readme:** Point Marketplace badge at v0.1.1 [f4cfe59](https://github.com/platformrelay/provider-gridscale/commit/f4cfe59ac26f3ddcddb62baa741aec0e6814b618)

- **readme:** Marketplace badge + datasource limitations (E3-S04/E8) [c15a328](https://github.com/platformrelay/provider-gridscale/commit/c15a328b0c43356c26b5da5e8246fd968ee8ee11)

## [0.1.1](https://github.com/platformrelay/provider-gridscale/compare/v0.1.0..v0.1.1) - 2026-07-16

### Bug Fixes

- **ci:** Use linked-package GITHUB_TOKEN for GHCR publish/sign [dda3fe6](https://github.com/platformrelay/provider-gridscale/commit/dda3fe6b5271d92b41a660cd04d9d68423ad59e9)

- **ci:** Self-contained Tag workflow; record supply-chain follow-ups [62da618](https://github.com/platformrelay/provider-gridscale/commit/62da6185834b9f80fd693859dcf1ab9f22323cc6)

- **hack/metadatafix:** Inject stubs for schema resources missing from scrape (E7-S02) [2e7583e](https://github.com/platformrelay/provider-gridscale/commit/2e7583ef2f0d784641bd9ae9395fafd83ae1ad61)


### Build System

- **toolchain:** Pin Go 1.26.5 everywhere + go-version drift guard (E5-S08) [c44e178](https://github.com/platformrelay/provider-gridscale/commit/c44e17817eb5ef018eabb610d0cca0c9eafeecd9)


### CI/CD

- **release:** Inline publish with github.token (drop upstream reusable) [1f8a20b](https://github.com/platformrelay/provider-gridscale/commit/1f8a20b2d167e604806570b1027c791ea031bce2)

- **release:** Publish/sign with github.token — drop GHCR_PAT secret [a3e3f54](https://github.com/platformrelay/provider-gridscale/commit/a3e3f5454f85ea273bc9b77d5eb6f0a892197c21)

- **workflows:** Activate schema-version-diff gate on PR schema changes (E5-S10) [a58b174](https://github.com/platformrelay/provider-gridscale/commit/a58b17442ecab10255e2444d930a6c8763518812)

- **lint:** Scope gosec G101/G104 to generated+test code (E5-S07) [4438b32](https://github.com/platformrelay/provider-gridscale/commit/4438b32f10b5174438482fa5213f96b73b41969f)

- **release:** Sign published package even if the xpkg.upbound.io mirror fails (decouple sign-and-sbom) [65b56ec](https://github.com/platformrelay/provider-gridscale/commit/65b56ec4ccbb6701616a3f97053b7dc67acc8b7c)


### Documentation

- Add unaffiliation disclaimer for gridscale GmbH [b5acc41](https://github.com/platformrelay/provider-gridscale/commit/b5acc418a62185dfa0c13843f3bc5e2602032c77)

- **coordination:** Session handoff — audit-gap loop done, next entry points [ff5a5f1](https://github.com/platformrelay/provider-gridscale/commit/ff5a5f1c67776da3a2d6803b354c0962d8dcd9e5)

- **decisions,inbox:** D-017 public Upbound + link GHCR package to repo [4cd7e94](https://github.com/platformrelay/provider-gridscale/commit/4cd7e941a8f08c164c9cafb3f2c7fb89b9dbc857)

- **ci:** Keep GHCR_PAT for upstream publish; token works for sign [c3a5b89](https://github.com/platformrelay/provider-gridscale/commit/c3a5b897733748b654db146af27f49a4203b8f81)

- **coordination:** Mark batch-5/6 complete; audit-gap loop exhausted [0f3f7b8](https://github.com/platformrelay/provider-gridscale/commit/0f3f7b805d95cfe72360d23aff2d3d9a717c786e)

- **decisions:** E6-S06 upstream triage for gridscale TF provider (#194/#200/#187/#188) [222f75c](https://github.com/platformrelay/provider-gridscale/commit/222f75c2caa0b04c04a253084f14dd8692d43665)

- **readme:** Flag upstream-deprecated resources and fields (E4-S05) [0f5fdaa](https://github.com/platformrelay/provider-gridscale/commit/0f5fdaa34fab5ec36397c5896b464410e1a26586)

- **coordination:** Mark batch-4 integrated; claim batch-5 lanes [752b764](https://github.com/platformrelay/provider-gridscale/commit/752b76448037dd93c09ec9509aec26c043758909)

- **adr:** Record cross-resource reference decision (ADR-0005) [b8fe841](https://github.com/platformrelay/provider-gridscale/commit/b8fe8414795e9ecb04ac0138763db62e56222279)

- **decisions:** D-015 — E8 datasources unsupported by upjet; rescope to document+track (L-E8SCOPE) [6622354](https://github.com/platformrelay/provider-gridscale/commit/6622354c5f07fa1f61118b7376bab70588e7679c)

- **decisions:** D-009b — BRAND-1 closed via official gridscale Bildmarke [10e107d](https://github.com/platformrelay/provider-gridscale/commit/10e107d1d4f9ecbf0696abcfd6da700e8b179f4a)

- **config:** Correct external-name rationale + document import formats (E7-S03) [921ac95](https://github.com/platformrelay/provider-gridscale/commit/921ac958c20590a542d9719ad1b0fd29a895214b)

- **coordination:** Claim batch-4 lanes; register E7 epic + E2-S11 renumbering (D-014) [ac858b1](https://github.com/platformrelay/provider-gridscale/commit/ac858b17da23404d2d60b80acdbbe746b8593329)

- **backlog:** Reconcile gap-stories vs D-012/D-013 — flag uptest-dependent criteria as superseded [784055c](https://github.com/platformrelay/provider-gridscale/commit/784055cf2ca3005030b74e6ab67123daffa69196)

- **backlog:** Audit gap-stories + gridscale docs research (for review) [693fee4](https://github.com/platformrelay/provider-gridscale/commit/693fee4146bc073641b475bc7944a370390ee8fe)

- **decisions,inbox:** Log operator answers — D-009 confirmed B (closed), D-012→B, D-013 PR#7 reconcile-then-merge [dade776](https://github.com/platformrelay/provider-gridscale/commit/dade776a48d955f2441a03aa263b9ba13daced3a)

- **inbox:** V0.1.0 published to ghcr; record signing (GHCR_PAT read) + Upbound-mirror cred follow-ups [4f96887](https://github.com/platformrelay/provider-gridscale/commit/4f968877d79e01f032876b64b2f43195b35e6fb4)


### Features

- **branding:** Sync official gridscale icon for Marketplace append [64fc363](https://github.com/platformrelay/provider-gridscale/commit/64fc363db50c783d8d5d50a0704c74eb0fb514cd)

- **config:** Wire full cross-resource reference set (E7-S01) [9ebca1d](https://github.com/platformrelay/provider-gridscale/commit/9ebca1d5389171477af12a58624383806a298ebc)

- **branding:** Use official gridscale Bildmarke for provider icon [9c342cb](https://github.com/platformrelay/provider-gridscale/commit/9c342cb7705c642579ea448bf90d19d3d70b5d60)


### Tests

- **config:** Assert provider-metadata keys match schema (E7-S02) [94c24bb](https://github.com/platformrelay/provider-gridscale/commit/94c24bb199929f25cbf1e50e38e4ec5e3dc6aaaa)

- **clients:** Non-live credential-wiring regression tests (E2-S11) [be73bd8](https://github.com/platformrelay/provider-gridscale/commit/be73bd8be680098c7dfd613ca0369a2e933c9dd5)

## [0.1.0](https://github.com/platformrelay/provider-gridscale/compare/v0.1.0-alpha.1..v0.1.0) - 2026-07-16

### Bug Fixes

- **cleanup:** Gridscale changelog label, go1.25 unify, drop dead schema gate, README matrix guard (DIR-2a/2c,TEST-3,DOC-5) [f9bf6f7](https://github.com/platformrelay/provider-gridscale/commit/f9bf6f73ddb5a4196fe666ea6686f622c2a3c543)

- **integration:** Cite D-012 for coverage-widen; switch Codecov to tokenless OIDC [05114ed](https://github.com/platformrelay/provider-gridscale/commit/05114edea04414d41b2f7aa7d07e72c187c49dd9)

- **lint:** Repair dead zz_ generated-file glob + correct goimports local-prefix (LINT-1/DIR-2) [1e9cfe1](https://github.com/platformrelay/provider-gridscale/commit/1e9cfe107307761deac32d024ea217e99a33c696)

- **lint:** Scope gosec G101/G104 suppression to tests only (SEC-1) [04f4527](https://github.com/platformrelay/provider-gridscale/commit/04f452789477c641a84e59802db8563fa4096d1b)

- **deps:** Bump go 1.26.5 + x/net v0.55.0 (govulncheck GO-2026-5856/5026/4918) ([#10](https://github.com/platformrelay/provider-gridscale/pull/10))[d75721e](https://github.com/platformrelay/provider-gridscale/commit/d75721e674b5ec130f9e2af62b335d8ecedefa03)

- **quality:** Exclude .claude/.work from arch-lint; gofmt + ascii-escape fuzz seed [0051304](https://github.com/platformrelay/provider-gridscale/commit/0051304473d62d529adf1283aab49f60bfd9c82d)


### Build System

- **quality:** Add vuln/race/tidy/arch-lint make targets (E2-S06/S07/S08/S10) [418957f](https://github.com/platformrelay/provider-gridscale/commit/418957fa7fb9e78719fb879eea3bd9edf56b886b)


### CI/CD

- Renovate go-version matcher; visible unsigned-publish warning; cliff skips changelog commits (CI-02/03) [d936e5a](https://github.com/platformrelay/provider-gridscale/commit/d936e5a70d0174881aaea41f03a71f848c6c95c8)

- **release:** Keyless cosign signing + SBOM on publish (E5-S06) [88f5326](https://github.com/platformrelay/provider-gridscale/commit/88f5326ab27ba8d19ad34b3a871be57c1ed61d75)

- **release:** Renovate tuning + git-cliff changelog + workflow (E5-S04/S05) [467b919](https://github.com/platformrelay/provider-gridscale/commit/467b919551c7de1c810b7cad03cb618ee3953d94)

- **security:** Add gitleaks config/workflow + pre-commit hooks (E5-S03) [6a4d384](https://github.com/platformrelay/provider-gridscale/commit/6a4d3841cb3a2791800f2a014e079af6c659005e)

- **supply-chain:** Add coverage/govulncheck/scorecard/codeql workflows (E5-S01/S02) [119e2de](https://github.com/platformrelay/provider-gridscale/commit/119e2de656473658decf8b15e179f651b7e80144)


### Documentation

- **governance:** Drop stale DOC-4 open-decision refs (maintainer set); track BRAND-1 icon polish [a3d7131](https://github.com/platformrelay/provider-gridscale/commit/a3d7131cb570e15e0c46aa6b3d023e45676c2d25)

- **decisions:** Reconcile D-008..D-011 as resolved; D-012 stays operator-blocked [23275ae](https://github.com/platformrelay/provider-gridscale/commit/23275ae6435e4485a19100c0453ab1692c3c3966)

- **audit,inbox:** Commit 2026-07-16 health audit + record v0.1.0 release-ready/held for operator go-ahead [6f254da](https://github.com/platformrelay/provider-gridscale/commit/6f254da24bf28e23329a2a03e6c9c4e9955a1092)

- **audit:** Mark ARCH-1/2, DIR-2, TEST-3, DOC-2 fixed; only operator-blocked TEST-2 remains [950f064](https://github.com/platformrelay/provider-gridscale/commit/950f0649e0a8414b35ae3cbb7a780711a6463c0d)

- **examples:** Add curated marketplace + ipv6 examples (DOC-2) [bdcf722](https://github.com/platformrelay/provider-gridscale/commit/bdcf72226747a2d1fd23c8c097acec0cadcbaa1a)

- **branding,testing:** Image-gen prompt (D-009) + uptest smoke runbook (D-012) [6222151](https://github.com/platformrelay/provider-gridscale/commit/6222151b2ab980d393733ffa796135ac0f1904f0)

- **readme:** Add logo + badge row + docs/examples pointers (E3-S02/S04) [eec5d77](https://github.com/platformrelay/provider-gridscale/commit/eec5d77751b1957bd19cf60c3be591116fea53b6)

- **governance:** Name @konih as maintainer (DOC-4) [f79d1a0](https://github.com/platformrelay/provider-gridscale/commit/f79d1a0807139a948ab16af5369552ce460b5d4e)

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

- **config:** Cross-resource references + provider-metadata fixes (ARCH-1/ARCH-2) [5959886](https://github.com/platformrelay/provider-gridscale/commit/595988691768cb33657a464a28b0e7af0ab8c167)

- **marketplace:** Add crossplane.yaml metadata + marketplace readme (E3-S03) [8ed8ec5](https://github.com/platformrelay/provider-gridscale/commit/8ed8ec5ff2200cd360a65bdc39e7e6716ee67d56)

- **branding:** Add provider icon + light/dark variants (E3-S01, placeholder) [df4a3ec](https://github.com/platformrelay/provider-gridscale/commit/df4a3ecf86f6440e5dbb9d8d7534a5129d166c35)


### Miscellaneous

- **deps:** Promote sigs.k8s.io/yaml to direct (go mod tidy; E2-S10) [a75053c](https://github.com/platformrelay/provider-gridscale/commit/a75053c174bb1ee0e6f10589c7798c22ec1b8a6b)


### Tests

- **metadatafix:** Table test for re-key rules + idempotency (F1) [3f3508b](https://github.com/platformrelay/provider-gridscale/commit/3f3508bf32c801aa3787c4e97f67f30e32636cda)

- **clients:** Cover credential resolution; widen coverage floor to internal/clients (E2-S03/D-007) [5d054e5](https://github.com/platformrelay/provider-gridscale/commit/5d054e55c3463d4f48d33a1ea9076c5093524abf)

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

