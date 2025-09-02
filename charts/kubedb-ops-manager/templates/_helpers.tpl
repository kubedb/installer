{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "kubedb-ops-manager.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kubedb-ops-manager.fullname" -}}
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
{{- define "kubedb-ops-manager.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kubedb-ops-manager.labels" -}}
helm.sh/chart: {{ include "kubedb-ops-manager.chart" . }}
{{ include "kubedb-ops-manager.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kubedb-ops-manager.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kubedb-ops-manager.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kubedb-ops-manager.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kubedb-ops-manager.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Returns the appscode license
*/}}
{{- define "appscode.license" -}}
{{- .Values.license }}
{{- end }}

{{/*
Returns the appscode license secret name
*/}}
{{- define "appscode.licenseSecretName" -}}
{{- if .Values.licenseSecretName }}
{{- .Values.licenseSecretName -}}
{{- else if .Values.license }}
{{- printf "%s-license" (include "kubedb-ops-manager.fullname" .) -}}
{{- end }}
{{- end }}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- list .Values.registryFQDN .Values.operator.registry | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for waitfor docker image
*/}}
{{- define "waitfor.registry" -}}
{{- list .Values.registryFQDN .Values.waitfor.registry | compact | join "/" }}
{{- end }}

{{- define "docker.imagePullSecrets" -}}
{{- with .Values.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{- define "docker.imagePullSecretFlags" -}}
{{- range .Values.imagePullSecrets }}
- --image-pull-secrets={{- .name -}}
{{- end }}
{{- end }}

{{/*
Returns the --insecure-registries flags
*/}}
{{- define "docker.insecureRegistries" -}}
{{- range (.Values.insecureRegistries | uniq | sortAlpha) }}
- --insecure-registries={{.}}
{{- end }}
{{- end }}

{{/*
Returns the enabled monitoring agent name
*/}}
{{- define "monitoring.agent" -}}
{{- .Values.monitoring.agent }}
{{- end }}

{{/*
Returns whether the ServiceMonitor will be labeled with custom label
*/}}
{{- define "monitoring.apply-servicemonitor-label" -}}
{{- ternary "false" "true" ( empty .Values.monitoring.serviceMonitor.labels ) -}}
{{- end }}

{{/*
Returns the ServiceMonitor labels
*/}}
{{- define "monitoring.servicemonitor-label" -}}
{{- range $key, $val := .Values.monitoring.serviceMonitor.labels }}
{{ $key }}: {{ $val }}
{{- end }}
{{- end }}

{{/*
Returns whether the OpenShift distribution is used
*/}}
{{- define "distro.openshift" -}}
{{- or (.Capabilities.APIVersions.Has "project.openshift.io/v1/Project") .Values.distro.openshift -}}
{{- end }}
