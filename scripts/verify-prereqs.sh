#!/usr/bin/env bash
# Sanity-check the local toolchain before starting the workshop.
# Usage: ./scripts/verify-prereqs.sh

set -u

ok() { printf "  \033[32m✓\033[0m %s\n" "$1"; }
warn() { printf "  \033[33m!\033[0m %s\n" "$1"; }
fail() { printf "  \033[31m✗\033[0m %s\n" "$1"; FAILED=1; }

FAILED=0

echo "Checking workshop prerequisites..."
echo

# --- Required tools ---
echo "Required tools:"
for tool in docker kubectl helm; do
  if command -v "$tool" >/dev/null 2>&1; then
    ver=$("$tool" version --short 2>/dev/null || "$tool" --version 2>/dev/null | head -1)
    ok "$tool — $ver"
  else
    fail "$tool not found in PATH"
  fi
done

# Bruno (used for the API exercises) is optional but recommended
if command -v bruno >/dev/null 2>&1 || command -v bru >/dev/null 2>&1; then
  ok "bruno (API client) — installed"
else
  warn "bruno (API client) — not found. Install from https://www.usebruno.com/ to run the API exercises."
fi

echo
echo "Cluster connectivity:"
if kubectl cluster-info >/dev/null 2>&1; then
  ctx=$(kubectl config current-context 2>/dev/null || echo "?")
  ok "kubectl can reach a cluster (context: $ctx)"
  ns=$(kubectl config view --minify -o jsonpath='{..namespace}' 2>/dev/null)
  ns=${ns:-default}
  ok "default namespace in current context: $ns"
else
  fail "kubectl cannot reach a cluster — check your KUBECONFIG (see docs/cluster-access.md)"
fi

echo
echo "Docker daemon:"
if docker info >/dev/null 2>&1; then
  ok "docker daemon is reachable"
else
  fail "docker daemon is not reachable — start Docker Desktop / dockerd"
fi

echo
if [ "$FAILED" -eq 0 ]; then
  echo "All checks passed. You're ready to start the workshop."
  exit 0
else
  echo "Some checks failed — fix the items above before continuing."
  exit 1
fi
