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
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- default .Values.operator.registry .Values.global.registry }}
{{- end }}

{{/*
Returns the registry used for catalog docker images
*/}}
{{- define "catalog.registry" -}}
{{- default .Values.image.registry .Values.global.registry }}
{{- end }}

{{/*
Returns the registry used for cleaner docker image
*/}}
{{- define "cleaner.registry" -}}
{{- default .Values.cleaner.registry .Values.global.registry }}
{{- end }}

{{/*
Returns whether the cleaner job YAML will be generated or not
*/}}
{{- define "cleaner.generate" -}}
{{- ternary "false" "true" (or .Values.global.skipCleaner .Values.cleaner.skip) -}}
{{- end }}

{{/*
Returns the appscode image pull secrets
*/}}
{{- define "appscode.imagePullSecrets" -}}
{{- with .Values.global.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- else -}}
imagePullSecrets:
{{- toYaml $.Values.imagePullSecrets | nindent 2 }}
{{- end }}
{{- end }}

{{/*
Returns the registry used for official docker images
*/}}
{{- define "official.registry" -}}
{{- if .image.overrideOfficialRegistry -}}
{{- $reg := default .image.registry .global.registry -}}
{{- list $reg (last .officialRegistry) | join "/" }}
{{- else -}}
{{- .officialRegistry | join "/" }}
{{- end }}
{{- end }}
