#!/usr/bin/env bash
set -euo pipefail
[ -f docs/observe-only.md ] || { echo "FAIL: docs/observe-only.md not found"; exit 1; }
grep -q "managementPolicies" docs/observe-only.md || { echo "FAIL: missing managementPolicies"; exit 1; }
grep -q "crossplane.io/external-name" docs/observe-only.md || { echo "FAIL: missing crossplane.io/external-name"; exit 1; }
echo "PASS: docs/observe-only.md is correct"
