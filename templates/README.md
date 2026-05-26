# Kubernetes manifest templates

Generic boilerplate for the four core Kubernetes resources, with `<placeholder> #FIXME` markers everywhere a value needs to be supplied. Copy a file into your own working area, then search for `#FIXME` and fill in the blanks.

For ready-to-apply manifests against the Pedelec example, see [`../slide-examples/`](../slide-examples/) (already filled in) or [`../exercises/02-kubernetes/solution/`](../exercises/02-kubernetes/solution/) (the worked exercise).

## What's here

| File | Resource | Use when… |
|---|---|---|
| [`namespace.yaml`](namespace.yaml) | `Namespace` | Creating a new isolated environment for your resources. |
| [`deployment.yaml`](deployment.yaml) | `Deployment` | Running one or more replicas of a container with rollout / self-healing. |
| [`service.yaml`](service.yaml) | `Service` | Giving a Deployment a stable in-cluster DNS name and load-balanced port. |
| [`ingress.yaml`](ingress.yaml) | `Ingress` | Exposing a Service to the outside world via HTTP(S). |

## Workflow

1. Pick the template you need.
2. Copy it into your workspace (e.g. `cp templates/deployment.yaml my-app-deployment.yml`).
3. Open in your editor and resolve every `#FIXME` marker.
4. Apply: `kubectl apply -f my-app-deployment.yml`.
5. Verify: `kubectl describe deployment <name>`.

## Provenance

These templates originated in the standalone [`Mtze/kubernetes-templates`](https://github.com/Mtze/kubernetes-templates) repository (Matthias Linhuber & Markus Braunbeck, *Container based Software Engineering: The Cloud at your Fingertips* workshop). They were copied into this repo at the time of the snapshot so the workshop is self-contained — if the upstream repo evolves, this copy will not pick up those changes automatically.
