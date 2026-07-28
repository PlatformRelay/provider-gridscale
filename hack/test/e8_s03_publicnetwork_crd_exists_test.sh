#!/usr/bin/env bash
set -euo pipefail
CLUSTER=package/crds/gridscale.gridscale.platformrelay.io_publicnetworks.yaml
NAMESPACED=package/crds/gridscale.gridscale.m.platformrelay.io_publicnetworks.yaml
[ -f "$CLUSTER" ] || { echo "FAIL: missing $CLUSTER"; exit 1; }
[ -f "$NAMESPACED" ] || { echo "FAIL: missing $NAMESPACED"; exit 1; }
grep -q "kind: CustomResourceDefinition" "$CLUSTER" || { echo "FAIL: not a CRD: $CLUSTER"; exit 1; }
grep -q "kind: CustomResourceDefinition" "$NAMESPACED" || { echo "FAIL: not a CRD: $NAMESPACED"; exit 1; }
echo "PASS: PublicNetwork CRD YAMLs exist"
