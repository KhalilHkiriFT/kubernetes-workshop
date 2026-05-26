# Cluster access

## TUM ASE CIT cluster (workshop default)

Workshop participants get a personal namespace on the shared TUM ASE CIT Kubernetes cluster (`k8s.ase.cit.tum.de`). The workshop facilitator will hand out a `kubeconfig` file per participant.

1. Save the file the facilitator gives you to `~/.kube/pedelec-workshop.yaml`.
2. Tell `kubectl` to use it for this session:

   ```bash
   export KUBECONFIG="$HOME/.kube/pedelec-workshop.yaml"
   ```

   Or merge it into your existing config:

   ```bash
   KUBECONFIG=~/.kube/config:~/.kube/pedelec-workshop.yaml \
     kubectl config view --flatten > ~/.kube/config.new
   mv ~/.kube/config.new ~/.kube/config
   ```

3. Confirm you can reach the cluster:

   ```bash
   kubectl cluster-info
   kubectl get pods
   ```

   The second command should succeed (possibly returning "No resources found") — that means RBAC is set up correctly.

### Your namespace and hostname

Each participant gets a unique namespace and a matching ingress hostname:

| Resource | Pattern | Example |
|---|---|---|
| Namespace | `<your-username>` | `mtze` |
| Ingress host | `<your-username>.pedelec.k8s.ase.cit.tum.de` | `mtze.pedelec.k8s.ase.cit.tum.de` |

Wherever the manifests in this repo say `example.pedelec.k8s.ase.cit.tum.de`, replace it with your hostname.

## Local cluster (self-study fallback)

If you're working through the exercises on your own machine, any local Kubernetes is fine. The fastest path:

```bash
# kind
kind create cluster --name pedelec
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

Then point your hostname at localhost in `/etc/hosts`:

```
127.0.0.1   pedelec.local
```

And use `pedelec.local` as the `host:` value in the ingress manifest.
