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
{{- list .Values.licenseSecretName .Values.global.licenseSecretName (printf "%s-license" (include "kubedb.fullname" .)) | compact | first }}
{{- end }}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.operator.registry .Values.global.registry) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for catalog docker images
*/}}
{{- define "catalog.registry" -}}
{{- list (default .registryFQDN .global.registryFQDN | default ._reg) (default .image.registry .global.registry | default ._repo) | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for webhook server docker image
*/}}
{{- define "server.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.server.registry .Values.global.registry) | compact | join "/" }}
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
{{- end }}
{{- end }}

{{/*
Returns the registry used for official docker images
*/}}
{{- define "official.registry" -}}
{{- if .image.overrideOfficialRegistry -}}
{{- list (default .registryFQDN .global.registryFQDN) (default .image.registry .global.registry) ._bin | compact | join "/" }}
{{- else -}}
{{- list (default .registryFQDN .global.registryFQDN) ._bin | compact | join "/" }}
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
Stash templates
*/}}

{{/*
Returns the registry used for cleaner docker image
*/}}
{{- define "cleaner.registry" -}}
{{- list (default .Values.registryFQDN .Values.global.registryFQDN) (default .Values.cleaner.registry .Values.global.registry) | compact | join "/" }}
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

{{- define "image-pull-secrets" -}}
{{- $secrets:= list -}}
{{- with .Values.global.imagePullSecrets -}}
{{- range $x := . -}}
{{- $secrets = append $secrets $x.name -}}
{{- end -}}
{{- else -}}
{{- range $x := $.Values.imagePullSecrets -}}
{{- $secrets = append $secrets $x.name -}}
{{- end -}}
{{- end }}
{{- $secrets | join "," | print -}}
{{- end -}}

