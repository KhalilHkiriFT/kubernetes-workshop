# Exercise 03 — Start here

> Templatize the raw manifests in `helm/pedelec/templates/` so the chart is configurable through `values.yaml`.

## What's given to you

```
helm/pedelec/
├── Chart.yaml             # chart metadata
├── values.yaml            # currently empty — you fill it in
├── .helmignore
└── templates/
    ├── _helpers.tpl       # stub — write reusable template snippets here
    ├── damage-deployment.yml
    ├── location-deployment.yml
    ├── reservation-deployment.yml
    ├── damage-service.yml
    ├── location-service.yml
    ├── reservation-service.yml
    └── pedelec-ingress.yml
```

The files in `templates/` are exact copies of the exercise 02 solution. Right now the chart works (`helm template ./helm/pedelec` will render valid YAML) — but nothing is configurable.

## Your task

Pull the values that would naturally differ between environments out of the templates and into `values.yaml`. At a minimum:

| What | Why |
|---|---|
| Image registry, owner, tag | So a fork can deploy its own builds without editing templates. |
| Ingress hostname | Every participant has their own host. |
| `enabled` flag for the Ingress | Some envs (CI, local) might not want one. |
| Resource requests/limits | Production might want different sizing. |

## Suggested order

### 1. Templatize the image reference

In each deployment, change the hardcoded image into a templated one. Define a helper in `_helpers.tpl` that builds the full image reference from `.Values.image.*`:

<details><summary>Hint — image helper</summary>

```gotemplate
{{/* Build a full image reference from .Values.image */}}
{{- define "pedelec.image" -}}
{{ .registry }}/{{ .owner }}/pedelec-{{ .service }}:{{ .tag }}
{{- end }}
```

And invoke it from a deployment:

```yaml
image: {{ include "pedelec.image" (dict "registry" .Values.image.registry "owner" .Values.image.owner "tag" .Values.image.tag "service" "damage") }}
```
</details>

### 2. Add a common labels helper

Every K8s resource should carry the standard `app.kubernetes.io/*` labels. Define it once, reuse everywhere:

```gotemplate
{{- define "pedelec.labels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}
```

Apply it via `{{- include "pedelec.labels" . | nindent 4 }}`.

### 3. Conditionally render the Ingress

Wrap the whole ingress manifest in `{{- if .Values.ingress.enabled }} … {{- end }}` so installs without an ingress controller don't fail.

### 4. Lift resources

Replace the hardcoded `requests`/`limits` blocks in each Deployment with `{{- toYaml .Values.resources | nindent 12 }}`.

## How to test before applying

`helm template` renders the chart locally without touching the cluster:

```bash
helm template pedelec ./helm/pedelec | less
```

Pipe through `kubectl apply --dry-run=client -f -` to validate the rendered YAML:

```bash
helm template pedelec ./helm/pedelec | kubectl apply --dry-run=client -f -
```

Or `helm lint` for schema checks:

```bash
helm lint ./helm/pedelec
```

## Install for real

```bash
helm install pedelec ./helm/pedelec \
  --namespace <your-namespace> \
  --set image.owner=mtze \
  --set ingress.host=<your-host>
```

Upgrade later with `helm upgrade pedelec ./helm/pedelec --reuse-values ...`.

## Uninstall

```bash
helm uninstall pedelec -n <your-namespace>
```

## Done?

Open [`../solution/README.md`](../solution/README.md) for a finished chart and a walkthrough of the choices.
