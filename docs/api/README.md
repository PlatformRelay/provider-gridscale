# CRD API Reference

This directory contains the generated API reference for the provider's Custom
Resource Definitions (CRDs).

**Do not hand-edit these files.** They are generated from the Go API types under
[`apis/`](../../apis) by [`elastic/crd-ref-docs`](https://github.com/elastic/crd-ref-docs).

To regenerate, run from the repository root:

```sh
make docs
```

The generator, its pinned version, and the exact invocation are defined by the
`docs` target in the root `Makefile`; rendering options live in
`.crd-ref-docs.yaml`. Generation is deterministic — re-running `make docs`
produces no diff — so the committed output can be diffed against a fresh render
in CI to catch drift.

The reference itself is in [`out.md`](out.md).
