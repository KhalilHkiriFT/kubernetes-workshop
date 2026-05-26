# Prerequisites

Before the workshop, install the following on your machine:

| Tool | Minimum version | Used for |
|---|---|---|
| [Docker](https://docs.docker.com/get-docker/) | 24.x | Building and running the microservices in exercise 01. |
| [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) | 1.28+ | Talking to the cluster in exercises 02 and 03. |
| [Helm](https://helm.sh/docs/intro/install/) | 3.13+ | The chart exercise in exercise 03. |
| [Bruno](https://www.usebruno.com/) (optional) | latest | Running the prepared API requests against the services. |
| Go (optional) | 1.22+ | Only if you want to read or modify the service code locally. The build runs inside Docker. |

> [!NOTE]
> You can verify your setup with `./scripts/verify-prereqs.sh` from the repo root.

## Cluster access

You'll also need access to a Kubernetes cluster. See [`cluster-access.md`](cluster-access.md) for the TUM ASE CIT cluster setup. If you're following this workshop on your own and don't have access to that cluster, any local cluster will do — [kind](https://kind.sigs.k8s.io/), [minikube](https://minikube.sigs.k8s.io/), [k3d](https://k3d.io/), or Docker Desktop's built-in Kubernetes all work fine.

If you're using a local cluster, you'll also need an ingress controller. The quickest path:

```bash
# kind / k3d / generic
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
```

## A namespace to play in

Most workshop exercises assume you have a namespace you can write to. On the TUM cluster you'll already have one — see `cluster-access.md`. On a local cluster, create one:

```bash
kubectl create namespace pedelec
kubectl config set-context --current --namespace=pedelec
```

After that, every `kubectl apply` lands in your namespace without you having to pass `-n` every time.
