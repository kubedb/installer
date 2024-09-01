{{/*
Expand the name of the chart.
*/}}
{{- define "kubedb.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kubedb.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "kubedb.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kubedb.labels" -}}
helm.sh/chart: {{ include "kubedb.chart" . }}
{{ include "kubedb.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kubedb.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kubedb.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kubedb.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kubedb.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Returns the appscode license
*/}}
{{- define "appscode.license" -}}
{{- default .Values.global.license .Values.license }}
{{- end }}

{{/*
Returns the appscode license secret name
*/}}
{{- define "appscode.licenseSecretName" -}}
{{- if .Values.licenseSecretName }}
{{- .Values.licenseSecretName -}}
{{- else if .Values.global.licenseSecretName }}
{{- .Values.global.licenseSecretName -}}
{{- else if (default .Values.global.license .Values.license) }}
{{- printf "%s-license" (include "kubedb.fullname" .) -}}
{{- end }}
{{- end }}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.operator.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for webhook server docker image
*/}}
{{- define "server.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.server.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for cleaner docker image
*/}}
{{- define "cleaner.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.cleaner.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for image docker image
*/}}
{{- define "image.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.image.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for rbacproxy docker rbacproxy
*/}}
{{- define "rbacproxy.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.rbacproxy.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for k8s-wait-for docker image
*/}}
{{- define "waitfor.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.waitfor.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the appscode image pull secrets
*/}}
{{- define "docker.imagePullSecrets" -}}
{{- with .Values.global.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- else -}}
imagePullSecrets:
{{- toYaml $.Values.imagePullSecrets | nindent 2 }}
{{- end }}
{{- end }}

{{- define "docker.imagePullSecretFlags" -}}
{{- range .Values.global.imagePullSecrets }}
- --image-pull-secrets={{- .name -}}
{{- else -}}
{{- range $.Values.imagePullSecrets }}
- --image-pull-secrets={{- .name -}}
{{- end }}
{{- end }}
{{- end }}

{{/*
Returns the --insecure-registries flags
*/}}
{{- define "docker.insecureRegistries" -}}
{{- range (concat .Values.global.insecureRegistries .Values.insecureRegistries | uniq | sortAlpha) }}
- --insecure-registries={{.}}
{{- end -}}
{{- end }}

{{/*
Returns the enabled monitoring agent name
*/}}
{{- define "monitoring.agent" -}}
{{- default .Values.monitoring.agent .Values.global.monitoring.agent }}
{{- end }}

{{/*
Returns whether the ServiceMonitor will be labeled with custom label
*/}}
{{- define "monitoring.apply-servicemonitor-label" -}}
{{- ternary "false" "true" (and (empty .Values.global.monitoring.serviceMonitor.labels) (empty .Values.monitoring.serviceMonitor.labels) ) -}}
{{- end }}

{{/*
Returns the ServiceMonitor labels
*/}}
{{- define "monitoring.servicemonitor-label" -}}
{{- range $key, $val := mergeOverwrite dict .Values.monitoring.serviceMonitor.labels .Values.global.monitoring.serviceMonitor.labels }}
{{ $key }}: {{ $val }}
{{- end }}
{{- end }}
