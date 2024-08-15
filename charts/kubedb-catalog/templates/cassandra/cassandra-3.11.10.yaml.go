{{ $featureGates := .Values.featureGates }}
{{- if .Values.global }}
  {{ $featureGates = mergeOverwrite dict .Values.featureGates .Values.global.featureGates }}
{{- end }}

{{ if $featureGates.Druid }}

apiVersion: catalog.kubedb.com/v1alpha1
kind: DruidVersion
metadata:
  name: '3.11.10'
  labels:
    {{- include "kubedb-catalog.labels" . | nindent 4 }}
spec:
  db:
    image: '{{ include "image.dockerHub" (merge (dict "_repo" "apache/cassandra") $) }}:3.11.10'
  initContainer:
    image: '{{ include "image.ghcr" (merge (dict "_repo" "kubedb/cassandra-init") $) }}'
  securityContext:
    runAsUser: 999
  version: 3.11.10
{{ end }}
