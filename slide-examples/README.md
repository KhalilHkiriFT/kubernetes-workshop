# Slide examples

Self-contained YAML snippets that match the code shown on the workshop slides. Use these during the live walkthrough to demonstrate one concept at a time without having to type from scratch.

For full end-to-end exercises that put these concepts together, see [`../exercises/`](../exercises/).

## Numbering scheme

- **`01–09`** — Docker Compose concepts
- **`10–19`** — Kubernetes concepts
- (Future: `20+` reserved for Helm and other advanced topics)

Gaps in the numbering are intentional; they leave room to insert new examples without renumbering the existing ones.

## What's here

| # | Concept | Slide | File |
|---|---|---|---|
| 01 | Single-service Docker Compose | 8 | [01-docker-compose-single-service/compose.yml](01-docker-compose-single-service/compose.yml) |
| 02 | Multiple services in Compose | 9 | [02-docker-compose-multi-service/compose.yml](02-docker-compose-multi-service/compose.yml) |
| 03 | Publishing container ports | 11 | [03-docker-compose-ports/compose.yml](03-docker-compose-ports/compose.yml) |
| 04 | Environment variables | 14 | [04-docker-compose-environment/compose.yml](04-docker-compose-environment/compose.yml) |
| 05 | Custom networks + `expose` | 24 | [05-docker-compose-networks/compose.yml](05-docker-compose-networks/compose.yml) |
| 10 | Kubernetes `Namespace` | 50 | [10-kubernetes-namespace/namespace.yml](10-kubernetes-namespace/namespace.yml) |
| 11 | Kubernetes `Pod` | 54 | [11-kubernetes-pod/pod.yml](11-kubernetes-pod/pod.yml) — and a deliberately broken [`incorrect-pod.yml`](11-kubernetes-pod/incorrect-pod.yml) for spot-the-bug discussion |
| 12 | Kubernetes `Deployment` (generic) | 60 | [12-kubernetes-deployment/deployment.yml](12-kubernetes-deployment/deployment.yml) |
| 13 | Kubernetes `Deployment` (Pedelec reservation) | 68 | [13-kubernetes-deployment-reservation/reservation-deployment.yml](13-kubernetes-deployment-reservation/reservation-deployment.yml) |
| 14 | Kubernetes `Service` (generic) | 75 | [14-kubernetes-service/service.yml](14-kubernetes-service/service.yml) |
| 15 | Kubernetes `Service` (Pedelec reservation) | 81 | [15-kubernetes-service-reservation/reservation-service.yml](15-kubernetes-service-reservation/reservation-service.yml) |
| 16 | Kubernetes `Ingress` | 90 | [16-kubernetes-ingress/ingress.yml](16-kubernetes-ingress/ingress.yml) |

## How to use

Each subfolder has just the manifest. For the Docker Compose examples:

```bash
cd slide-examples/01-docker-compose-single-service
docker compose up
```

For the Kubernetes examples (substitute `my-namespace` for your own):

```bash
kubectl apply -f slide-examples/11-kubernetes-pod/pod.yml
```

## Heads up — image references vary

The image refs in the slide examples reflect what's *on the slides*, which historically come from a few different registries:

- Some slide examples use `nginxdemos/hello` from Docker Hub (always-available demo image)
- Example 13 (Reservation Deployment) uses `ghcr.io/ls1intum/pedelec-app/reservation:latest`
- The Pedelec build workflow in this repo publishes to `ghcr.io/<owner>/pedelec-<service>:latest` (different naming)

This is documented but not yet unified; the canonical example-exercise mapping is in [`../exercises/`](../exercises/).
