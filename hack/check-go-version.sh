#!/usr/bin/env bash
# check-go-version.sh (E5-S08 / DIR-2): guard against Go toolchain version drift.
#
# Asserts that the pinned Go version is byte-identical across
#   - Makefile                    (GO_REQUIRED_VERSION ?= X)
#   - .github/workflows/ci.yml    (env GO_VERSION: 'X')
#   - .github/workflows/e2e.yaml  (env GO_VERSION: "X")
# and that it is compatible with go.mod:
#   - if go.mod has a `toolchain goX` directive, the pin must equal X
#     (that toolchain is what actually builds the module);
#   - otherwise the pin must equal the `go` directive.
#
# Deterministic: reads only the four files, no network, no environment input.
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

fail() {
  echo "check-go-version: FAIL: $*" >&2
  exit 1
}

extract_one() {
  # extract_one <label> <file> <sed-expr> -> exactly one match on stdout
  local label="$1" file="$2" expr="$3" out
  [[ -f "$file" ]] || fail "$label: file not found: $file"
  out="$(sed -nE "$expr" "$file")"
  [[ -n "$out" ]] || fail "$label: no Go version found in $file"
  [[ "$(wc -l <<<"$out")" -eq 1 ]] || fail "$label: multiple Go version lines in $file"
  printf '%s' "$out"
}

makefile_v="$(extract_one Makefile "$root/Makefile" \
  's/^GO_REQUIRED_VERSION[[:space:]]*\?=[[:space:]]*([0-9][0-9.]*)[[:space:]]*$/\1/p')"
ci_v="$(extract_one ci.yml "$root/.github/workflows/ci.yml" \
  "s/^[[:space:]]*GO_VERSION:[[:space:]]*['\"]?([0-9][0-9.]*)['\"]?[[:space:]]*$/\1/p")"
e2e_v="$(extract_one e2e.yaml "$root/.github/workflows/e2e.yaml" \
  "s/^[[:space:]]*GO_VERSION:[[:space:]]*['\"]?([0-9][0-9.]*)['\"]?[[:space:]]*$/\1/p")"

gomod_go="$(sed -nE 's/^go[[:space:]]+([0-9][0-9.]*)[[:space:]]*$/\1/p' "$root/go.mod")"
gomod_toolchain="$(sed -nE 's/^toolchain[[:space:]]+go([0-9][0-9.]*)[[:space:]]*$/\1/p' "$root/go.mod")"
[[ -n "$gomod_go" ]] || fail "go.mod: no go directive found"

echo "check-go-version: Makefile=$makefile_v ci.yml=$ci_v e2e.yaml=$e2e_v" \
  "go.mod(go)=$gomod_go go.mod(toolchain)=${gomod_toolchain:-<none>}"

[[ "$ci_v" == "$makefile_v" ]] ||
  fail "drift: ci.yml GO_VERSION=$ci_v != Makefile GO_REQUIRED_VERSION=$makefile_v"
[[ "$e2e_v" == "$makefile_v" ]] ||
  fail "drift: e2e.yaml GO_VERSION=$e2e_v != Makefile GO_REQUIRED_VERSION=$makefile_v"

if [[ -n "$gomod_toolchain" ]]; then
  [[ "$makefile_v" == "$gomod_toolchain" ]] ||
    fail "drift: pinned Go $makefile_v != go.mod toolchain go$gomod_toolchain"
else
  [[ "$makefile_v" == "$gomod_go" ]] ||
    fail "drift: pinned Go $makefile_v != go.mod go directive $gomod_go"
fi

echo "check-go-version: OK — Go $makefile_v pinned consistently"
