# Exercise 03 — Helm

> Convert the raw Kubernetes manifests from exercise 02 into a reusable Helm chart, then install it with a single command.

This exercise has two folders:

- [`start/`](start/README.md) — a chart skeleton (`Chart.yaml`, empty `values.yaml`, raw manifests copied into `templates/`). **You'll templatize them.**
- [`solution/`](solution/README.md) — the finished chart with a walkthrough of the templating decisions.

## Why Helm?

By exercise 02 every team member has nearly-identical manifests with one small difference — typically the namespace, the ingress hostname, or the image tag. Helm replaces "edit YAML for each environment" with "pass different `--set` flags."

## Learning goals

By the end of this exercise you'll be able to:

- Read and write Go-template syntax in a Helm chart.
- Pull configurable bits (image, hostname, replica count) into `values.yaml`.
- Install, upgrade, and uninstall a Helm release on a cluster.
- Use `helm template` to debug before applying anything to a cluster.

## Prerequisites

- Helm 3.13+ installed locally
- Exercise 02 completed (or at least studied — the chart is just the manifests with template directives)

## Where to start

```bash
cd exercises/03-helm/start
cat README.md
```
