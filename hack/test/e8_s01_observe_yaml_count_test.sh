#!/usr/bin/env bash
set -euo pipefail
CLUSTER=$(find examples/cluster/observe -name '*.yaml' 2>/dev/null | wc -l | tr -d ' ')
NAMESPACED=$(find examples/namespaced/observe -name '*.yaml' 2>/dev/null | wc -l | tr -d ' ')
[ "$CLUSTER" -eq 19 ] || { echo "FAIL: cluster observe YAMLs: $CLUSTER (want 19)"; exit 1; }
[ "$NAMESPACED" -eq 19 ] || { echo "FAIL: namespaced observe YAMLs: $NAMESPACED (want 19)"; exit 1; }
echo "PASS: cluster=$CLUSTER namespaced=$NAMESPACED"
