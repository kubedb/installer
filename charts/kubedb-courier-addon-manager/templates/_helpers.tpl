{{/* Expand the name of the chart. */}}
{{- define "courier-addon-manager.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/* Fully qualified app name. */}}
{{- define "courier-addon-manager.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/* Common labels. */}}
{{- define "courier-addon-manager.labels" -}}
app.kubernetes.io/name: {{ include "courier-addon-manager.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" }}
{{- end -}}

{{/* Selector labels. */}}
{{- define "courier-addon-manager.selectorLabels" -}}
app.kubernetes.io/name: {{ include "courier-addon-manager.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{/* Manager image reference. */}}
{{- define "courier-addon-manager.image" -}}
{{- $tag := .Values.image.tag | default .Chart.AppVersion -}}
{{- printf "%s/%s:%s" .Values.image.registry .Values.image.repository $tag -}}
{{- end -}}

{{/* Agent image reference, defaulting to the shared manager image. */}}
{{- define "courier-addon-manager.agentImage" -}}
{{- $reg := .Values.agent.image.registry | default .Values.image.registry -}}
{{- $repo := .Values.agent.image.repository | default .Values.image.repository -}}
{{- $tag := .Values.agent.image.tag | default .Values.image.tag | default .Chart.AppVersion -}}
{{- printf "%s/%s:%s" $reg $repo $tag -}}
{{- end -}}
