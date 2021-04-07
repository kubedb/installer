{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "kubedb-catalog.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kubedb-catalog.fullname" -}}
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
{{- define "kubedb-catalog.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kubedb-catalog.labels" -}}
helm.sh/chart: {{ include "kubedb-catalog.chart" . }}
{{ include "kubedb-catalog.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kubedb-catalog.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kubedb-catalog.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kubedb-catalog.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kubedb-catalog.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Returns the registry used for catalog docker images
*/}}
{{- define "catalog.registry" -}}
{{- list .Values.registryFQDN .Values.image.registry | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for official docker images
*/}}
{{- define "official.registry" -}}
{{- if .image.overrideOfficialRegistry -}}
{{- list .registryFQDN .image.registry (last .officialRegistry) | compact | join "/" }}
{{- else -}}
{{- prepend .officialRegistry .registryFQDN | compact | join "/" }}
{{- end }}
{{- end }}
