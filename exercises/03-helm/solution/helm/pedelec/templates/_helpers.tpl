{{/* Common labels applied to all resources */}}
{{- define "pedelec.labels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/* Build a full image reference from values.image */}}
{{- define "pedelec.image" -}}
{{ .registry }}/{{ .owner }}/pedelec-{{ .service }}:{{ .tag }}
{{- end }}
