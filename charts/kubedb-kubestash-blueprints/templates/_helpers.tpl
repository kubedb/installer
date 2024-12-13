{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "kubedb-kubestash-blueprints.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kubedb-kubestash-blueprints.fullname" -}}
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
{{- define "kubedb-kubestash-blueprints.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kubedb-kubestash-blueprints.labels" -}}
helm.sh/chart: {{ include "kubedb-kubestash-blueprints.chart" . }}
{{ include "kubedb-kubestash-blueprints.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kubedb-kubestash-blueprints.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kubedb-kubestash-blueprints.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kubedb-kubestash-blueprints.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kubedb-kubestash-blueprints.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "image.dockerHub" -}}
{{ list .Values.proxies.dockerHub ._repo | compact | join "/" }}
{{- end }}

{{- define "image.dockerLibrary" -}}
{{ prepend (list ._repo) (list .Values.proxies.dockerLibrary .Values.proxies.dockerHub | compact | first) | compact | join "/" }}
{{- end }}

{{- define "image.ghcr" -}}
{{ list .Values.proxies.ghcr ._repo | compact | join "/" }}
{{- end }}

{{- define "image.kubernetes" -}}
{{ list .Values.proxies.kubernetes ._repo | compact | join "/" }}
{{- end }}

{{- define "image.microsoft" -}}
{{ list .Values.proxies.microsoft ._repo | compact | join "/" }}
{{- end }}

{{- define "image.appscode" -}}
{{ list .Values.proxies.appscode ._repo | compact | join "/" }}
{{- end }}
