{{/*
Expand the name of the chart.
*/}}
{{- define "kubedb-opscenter.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kubedb-opscenter.fullname" -}}
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
{{- define "kubedb-opscenter.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kubedb-opscenter.labels" -}}
helm.sh/chart: {{ include "kubedb-opscenter.chart" . }}
{{ include "kubedb-opscenter.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kubedb-opscenter.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kubedb-opscenter.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kubedb-opscenter.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kubedb-opscenter.fullname" .) .Values.serviceAccount.name }}
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
Returns the registry used for image docker image
*/}}
{{- define "image.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.image.registry .Values.global.registry) | compact | join "/" }}
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
{{- range $key, $val := .Values.monitoring.serviceMonitor.labels }}
{{ $key }}: {{ $val }}
{{- else }}
{{- range $key, $val := .Values.global.monitoring.serviceMonitor.labels }}
{{ $key }}: {{ $val }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Returns whether the NetworkPolicy should be enabled
*/}}
{{- define "security.enableNetworkPolicy" -}}
{{- or .Values.global.networkPolicy.enabled (and .Values.networkPolicy .Values.networkPolicy.enabled) -}}
{{- end }}

{{/*
Returns whether the OpenShift distribution is used
*/}}
{{- define "distro.openshift" -}}
{{- or (.Capabilities.APIVersions.Has "project.openshift.io/v1/Project") .Values.global.distro.openshift (and .Values.distro .Values.distro.openshift) -}}
{{- end }}
