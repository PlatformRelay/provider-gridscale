#!/usr/bin/env bash
# REQ-SONAR-PG-02: UPTEST_DATASOURCE reaches the shell only via env:, never
# via ${{ secrets.* }} interpolation inside a run: block.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
WF="${ROOT}/.github/workflows/e2e.yaml"

fail() {
  echo "FAIL: $*" >&2
  exit 1
}

[[ -f "${WF}" ]] || fail "missing ${WF}"

# Secret must not be expanded inside any run: script body.
# Allowed only under an env: binding (key: ${{ secrets.UPTEST_DATASOURCE }}).
# Disallow the echo/printf form that embeds the expression in run:.
if grep -Eq '\$\{\{[[:space:]]*secrets\.UPTEST_DATASOURCE[[:space:]]*\}\}' "${WF}" \
  && grep -Eq '(echo|printf).*\$\{\{[[:space:]]*secrets\.UPTEST_DATASOURCE[[:space:]]*\}\}' "${WF}"; then
  fail "${WF}: secrets.UPTEST_DATASOURCE must not appear inside run: (use env: + \"\$UPTEST_DATASOURCE\")"
fi

# Must bind the secret through the step env: block.
if ! grep -Eq '^[[:space:]]*UPTEST_DATASOURCE:[[:space:]]*\$\{\{[[:space:]]*secrets\.UPTEST_DATASOURCE[[:space:]]*\}\}' "${WF}"; then
  fail "${WF}: missing env.UPTEST_DATASOURCE: \${{ secrets.UPTEST_DATASOURCE }}"
fi

# run: must write the file from the env var (printf preferred; "$UPTEST_DATASOURCE" ok).
if ! grep -Eq '\$UPTEST_DATASOURCE' "${WF}"; then
  fail "${WF}: run: must reference \"\$UPTEST_DATASOURCE\" (not secrets expression)"
fi

# Accept either printf '%s\n' "$UPTEST_DATASOURCE" form from the spec.
if ! grep -Eq "printf[[:space:]]+'%s\\\\n'[[:space:]]+\"\\\$UPTEST_DATASOURCE\"" "${WF}" \
  && ! grep -Eq 'printf[[:space:]]+"%s\\n"[[:space:]]+"\$UPTEST_DATASOURCE"' "${WF}" \
  && ! grep -F 'printf' "${WF}" | grep -q 'UPTEST_DATASOURCE'; then
  fail "${WF}: run: should write datasource via printf '%s\\n' \"\$UPTEST_DATASOURCE\""
fi

echo "PASS: UPTEST_DATASOURCE is bound via env: and referenced as \"\$UPTEST_DATASOURCE\" in run:"
