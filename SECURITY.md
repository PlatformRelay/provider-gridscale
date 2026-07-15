# Security Policy

`provider-gridscale` is a [Crossplane](https://crossplane.io/) provider for
[gridscale](https://gridscale.io/), generated with
[Upjet](https://github.com/crossplane/upjet) on top of the upstream
[gridscale Terraform provider](https://github.com/gridscale/terraform-provider-gridscale).

> **Pre-1.0 status.** This provider is under active development toward its first
> release (`v0.1.0` in progress). It has not yet reached a stable `1.x` line, and
> its security posture is still being hardened — see
> [Supply-chain posture](#supply-chain-posture) below for what is shipped today
> versus what is planned. Treat this document as evolving alongside the project
> [roadmap](docs/ROADMAP.md).

## Reporting a Vulnerability

**Please do not report security vulnerabilities through public GitHub issues,
pull requests, or discussions.**

We ask for **private, coordinated disclosure**. The preferred channel is
**GitHub Security Advisories**:

1. Go to the repository's
   [Security Advisories](https://github.com/PlatformRelay/provider-gridscale/security/advisories/new)
   page.
2. Open a new draft advisory describing the issue.

This keeps the report private between you and the maintainers until a fix is
available.

<!-- TODO: operator to confirm disclosure contact -->
> **Disclosure contact (provisional).** A dedicated security contact
> (for example, a `security@…` address) has **not yet been finalized** for this
> project. Until it is, please use GitHub Security Advisories as described above.
> `TODO: operator to confirm disclosure contact` — the canonical private
> disclosure address is an open maintainer decision and this section will be
> updated once it is set.

When reporting, please include as much of the following as you can:

- A description of the vulnerability and its potential impact.
- Steps to reproduce, or a proof-of-concept.
- The provider version (or commit) and the environment it was observed in.
- Any suggested remediation, if you have one.

### What to expect

As a pre-1.0 community project, we do not yet offer a formal response-time SLA.
We aim to acknowledge a report as quickly as we reasonably can, keep you updated
as we investigate, and credit reporters who wish to be acknowledged once a fix
ships. Please allow a reasonable period for a fix before any public disclosure.

## Supported Versions

> **Provisional.** The version-support policy below is **provisional** and will
> be formalized as the project approaches its `1.0` release. Because the provider
> is pre-1.0, **only the latest release line receives fixes**; there are no
> maintained back-branches yet.

| Version | Supported          |
| ------- | ------------------ |
| `0.1.x` (in progress) | ✅ Latest — security fixes land here |
| `< 0.1.0` (pre-release) | ❌ Not supported |

`TODO: operator to confirm supported-versions policy` once a versioning and
release cadence is established.

## Threat Model

This section describes, at a high level, the security-relevant boundaries of the
provider itself. It intentionally scopes to `provider-gridscale`; the security of
the upstream [gridscale Terraform provider](https://github.com/gridscale/terraform-provider-gridscale),
[Upjet](https://github.com/crossplane/upjet), Crossplane, and the underlying
Kubernetes cluster is governed by their respective projects.

### Credential handling

- **gridscale API credentials are never stored in this repository.** They are
  supplied at runtime through a Kubernetes `Secret` that a Crossplane
  `ProviderConfig` references via `secretRef`. See the
  [Credentials & ProviderConfig](README.md#credentials--providerconfig) section
  of the README.
- The secret is a JSON document carrying the gridscale `uuid` (User-UUID) and
  `token` (API token), optionally an `api_url`. The provider resolves the
  referenced `ProviderConfig`, extracts these credentials, and passes them to the
  underlying Terraform provider configuration
  (see [`internal/clients/gridscale.go`](internal/clients/gridscale.go)). Local
  development credentials (e.g. `.envrc`, lab secrets) are **gitignored** and
  never committed.
- Because credentials live in Kubernetes `Secret`s, their protection at rest and
  the RBAC that governs access to them is the responsibility of the cluster
  operator (secret encryption, least-privilege RBAC, namespace isolation).

### Network

- The provider communicates with the gridscale API over **HTTPS/TLS**
  (default `https://api.gridscale.io`, overridable via the `api_url` credential
  key). It does not open inbound network listeners for managed-resource traffic.

### Trust boundaries (out of scope for this document)

- The confidentiality and integrity of Kubernetes `Secret`s, cluster RBAC, and
  the control plane itself.
- Vulnerabilities originating in upstream dependencies (the gridscale Terraform
  provider, Upjet, Crossplane, Go toolchain) — please report those to the
  respective upstream projects, though we welcome a heads-up so we can pick up
  fixed versions.

## Supply-chain Posture

Supply-chain hardening for this provider is tracked as **Epic E5** in the
[roadmap](docs/ROADMAP.md#e5--cicd-hardening--supply-chain). We are deliberately
honest about what is shipped versus planned.

### Shipped today

- **No secrets in source control.** gridscale credentials are provided only via
  `ProviderConfig` `secretRef` at runtime (see [Threat Model](#threat-model)).
  Local development / lab credential files (`.envrc`, `*.env`, `.env.local`) are
  excluded by a `.gitignore` rule so they cannot be committed accidentally. This
  is a `.gitignore` rule only — automated secret scanning (`gitleaks`,
  pre-commit) is **planned**, not yet active (see below).
- **Deliberate Terraform version pin.** `TERRAFORM_VERSION` is pinned to `1.5.7`
  (the last BSL-boundary release) as a considered dependency decision, not an
  incidental one.
- **Apache-2.0 licensed**, generated tree kept in sync with its `config/` inputs.

### Planned / in progress (Epic E5 — not yet shipped)

The following controls are on the roadmap and **not yet implemented**. Do not
assume they are active until the corresponding workflow or config lands:

- **Secret scanning** — `gitleaks` plus a `pre-commit` hook set (E5-S03). *Planned.*
- **Dependency & code scanning** — `govulncheck`, GitHub CodeQL, and OpenSSF
  Scorecard workflows (E5-S01 / E5-S02). *Planned.*
- **Pinned GitHub Actions** — CI actions pinned to commit SHAs. *Planned.*
- **Automated dependency updates** — Dependabot/Renovate for Go modules and
  Actions (E5-S04). *Planned.*
- **Signed releases** — released XPKG signed with keyless **cosign** and shipped
  with an **SBOM** (E5-S06). *Planned.*

As these land, this section will be updated to move each item from *planned* to
*shipped*.

## References

- Engineering guidelines (security & supply chain rules):
  [`agent-context/GUIDELINES.md`](agent-context/GUIDELINES.md)
- Project roadmap: [`docs/ROADMAP.md`](docs/ROADMAP.md)
- Code of Conduct: [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md)
- License: [`LICENSE`](LICENSE)
