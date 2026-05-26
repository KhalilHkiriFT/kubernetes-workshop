# Kubernetes glossary

A pocket-sized reference for the terms used across the exercises. For the canonical definitions see the [Kubernetes Concepts docs](https://kubernetes.io/docs/concepts/).

| Term | What it is | In this workshop |
|---|---|---|
| **Pod** | The smallest deployable unit in Kubernetes. One or more containers sharing a network namespace and storage. | Each `Pod` runs one of our three Go services. |
| **Deployment** | Declares "I want N copies of this pod template running." Handles rollouts and self-healing. | One `Deployment` per service. |
| **ReplicaSet** | Owned by a Deployment; ensures the desired number of pod replicas exists. | We never write these by hand — Deployments make them for us. |
| **Service** | A stable virtual IP + DNS name that load-balances across a set of pods. | One `Service` per Deployment, exposing port 80 → the service's container port. |
| **Ingress** | HTTP(S) routing rules that an ingress controller turns into real layer-7 routing. | One `Ingress` for the whole app, fanning out paths to services. |
| **Namespace** | A scope for names + an RBAC boundary. | You get your own namespace; everyone's resources live in their own sandbox. |
| **ConfigMap / Secret** | Key-value config injected into pods. | Not used in this workshop (the services have no config). |
| **Label / Selector** | Labels are key=value tags; selectors are queries over them. | `app: damage` is how a Service finds its pods. |
| **Manifest** | A YAML file describing one or more Kubernetes resources. | Everything in `k8s/` and `helm/.../templates/` is a manifest. |
| **kubectl** | The CLI for talking to a cluster. | `kubectl apply -f`, `kubectl get`, `kubectl describe`, `kubectl logs`. |
| **Helm chart** | A bundle of templated manifests plus a `values.yaml`. | Exercise 03 converts the raw manifests into one. |
| **Release** | One installation of a chart on a cluster. | `helm install pedelec ./helm/pedelec` creates a release named `pedelec`. |
| **Image** | A built container snapshot referenced by `registry/name:tag`. | Ours live on `ghcr.io/<owner>/pedelec-<service>`. |
| **GHCR** | GitHub Container Registry. | Where the build workflow publishes our images. |

## Useful kubectl one-liners

```bash
# What's running in my namespace?
kubectl get all

# Why isn't this pod ready?
kubectl describe pod <pod-name>

# What did this pod print to stdout?
kubectl logs <pod-name>
kubectl logs -f <pod-name>           # follow
kubectl logs <pod-name> --previous   # if it crashed and restarted

# What's a service actually selecting?
kubectl get endpoints <service-name>

# What's an ingress controller seeing?
kubectl get ingress
kubectl describe ingress pedelec-ingress

# Quick port-forward for debugging without ingress
kubectl port-forward service/pedelec-reservation-service 8080:80
```
