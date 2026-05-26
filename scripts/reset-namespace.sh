#!/usr/bin/env bash
# Remove everything the workshop has deployed into a namespace.
#
# Usage:
#   ./scripts/reset-namespace.sh                  # uses the current kubectl context's default namespace
#   ./scripts/reset-namespace.sh my-namespace     # explicit namespace

set -euo pipefail

NS="${1:-$(kubectl config view --minify -o jsonpath='{..namespace}')}"
NS="${NS:-default}"

read -r -p "About to delete all pedelec workshop resources in namespace '$NS'. Continue? [y/N] " ans
case "$ans" in
  [yY]|[yY][eE][sS]) ;;
  *) echo "Aborted."; exit 1 ;;
esac

echo "Deleting Helm release 'pedelec' (if present)..."
helm uninstall pedelec -n "$NS" 2>/dev/null || true

echo "Deleting raw K8s resources tagged for the workshop..."
kubectl -n "$NS" delete ingress  pedelec-ingress              --ignore-not-found
kubectl -n "$NS" delete service  pedelec-damage-service       --ignore-not-found
kubectl -n "$NS" delete service  pedelec-location-service     --ignore-not-found
kubectl -n "$NS" delete service  pedelec-reservation-service  --ignore-not-found
kubectl -n "$NS" delete deployment damage location reservation --ignore-not-found

echo
echo "Done. Remaining workload in namespace '$NS':"
kubectl -n "$NS" get all
