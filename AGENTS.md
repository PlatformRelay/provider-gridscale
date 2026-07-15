# AGENTS.md — provider-gridscale

Entry point for humans and AI assistants working on **provider-gridscale**: a
[Crossplane](https://crossplane.io) provider for [gridscale](https://gridscale.io), generated with
[Upjet](https://github.com/crossplane/upjet) from `terraform-provider-gridscale` v2.3.0. It exposes
XRM-conformant managed resources (32 resources, 8 API groups) for the gridscale IaaS/PaaS API.

## Documentation map

| | Document | What it covers |
| --- | --- | --- |
| 🎯 | [README.md](README.md) | Pitch, install, resource matrix, status |
| 🎯 | [docs/ROADMAP.md](docs/ROADMAP.md) | Epics E1–E6, build order, exit criteria |
| 🏗️ | [docs/adr/README.md](docs/adr/README.md) | Architecture decision records |
| 📋 | [openspec/config.yaml](openspec/config.yaml) | Spec-driven (OPSX) conventions & test taxonomy |
| 📋 | [agent-context/GUIDELINES.md](agent-context/GUIDELINES.md) | Binding engineering rules for this repo |
| 📋 | [agent-context/BACKLOG.md](agent-context/BACKLOG.md) | Lane deps, story IDs, hazards |
| 🗳️ | [agent-context/INBOX.md](agent-context/INBOX.md) | Open decisions & operator items |
| 🧾 | [agent-context/decisions.md](agent-context/decisions.md) | Append-only decision log |

Workspace-level working contract (commits, merge policy, adversarial collaboration):
[`../AGENTS.md`](../AGENTS.md). This file owns *provider-gridscale specifics*; it does not restate it.

## What this repo is

A **thin, generated** provider. The valuable, hand-authored surface is small — treat the generated
tree as read-only build output:

| Surface | Who writes it | Edit? |
| --- | --- | --- |
| `apis/**/zz_*.go`, `internal/controller/**`, `package/crds/**` | Upjet codegen | ❌ never — change `config/` instead |
| `config/` (external-name, kind/field overrides, wiring) | humans | ✅ the main lever |
| `examples/` (also uptest e2e fixtures) | humans | ✅ |
| `package/crossplane.yaml`, `extensions/` (icon, readme, notes) | humans | ✅ Marketplace surface |
| `docs/`, `.github/`, governance files | humans | ✅ maturity surface |

`make reviewable` / `make check-diff` will fail a PR that hand-edits generated code.

## Conventions

- **Commits:** `:gitmoji: type(scope): summary` — ASCII shortcode only; no Unicode emoji; no AI
  co-author trailers; never modify git config. One logical change per commit.
- **Merge:** rebase-and-merge only — linear history, no squash, no merge commits.
- **Spec-driven:** every implementation lane references an OpenSpec change under `openspec/changes/`.
  A change is `proposal.md` + `design.md` + `tasks.md` + `specs/<capability>/spec.md`.
- **TDD:** mandatory for hand-authored Go and `hack/` scripts — failing test first. Generated code
  is exempt but must stay green and is covered at the **E** (uptest) tier.
- **Test taxonomy:** **U** unit · **E** uptest e2e · **M** meta/structural · **D** doc-sync. Every
  REQ has one `**Test:**` (artifact) + one `**Verify:**` (command). Details in `openspec/config.yaml`.
- **Committed harness:** `AGENTS.md`, `openspec/`, `agent-context/`, `docs/` are committed here (a
  public provider showing its workflow). Only `.claude/` is local/gitignored.

## Gate commands

```bash
make reviewable      # generate + lint + tidy — must be clean before PR (guards generated tree)
make test            # Go unit tests (U) with coverage → cover.out
make test.spec       # structural: every REQ has a Test: path + Verify: command  (added in E1-S01)
make check-diff      # CRD/schema diff — fails on hand-edited generated output
make build           # build provider + xpkg
# E-level (uptest, real gridscale creds) runs in CI via /test-examples + nightly — not on push.
```

## Next action

Run `/pick-next-story` (or read [docs/ROADMAP.md](docs/ROADMAP.md) + [agent-context/BACKLOG.md](agent-context/BACKLOG.md)):
start at **E1** (workflow foundation); E2·E3·E4·E6 parallelise after it; E5 integrates last. Resolve
open decisions **D-001…D-004** ([INBOX.md](agent-context/INBOX.md)) before fanning out.
