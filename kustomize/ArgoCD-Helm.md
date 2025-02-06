Here's a sample manifest that can be used to deploy KubeDB helm chart via ArgoCD. The source here have been reffered to the kubedb oci chart and destination has been set to an Incluster kubernetes platform. Automatic sync policy has been enabled in this ArgoCD Application manifest. 

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: kubedb
  namespace: argocd
spec:
  project: default
  source:
    chart: kubedb
    repoURL: ghcr.io/appscode-charts      # note: the oci:// syntax is not included.
    targetRevision: v2024.6.4             # supports wildcards for automatic chart upgrade (eg. v2024.*.* or *.*.*)
    helm:
      parameters:
      - name: global.featureGates.RabbitMQ
        value: "true"
      - name: supervisor.enabled
        value: "true"
      - name: global.licenseSecretName
        value: kubedb-license
  destination:
    server: "https://kubernetes.default.svc"
    namespace: kubedb
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true            # creates target namespace if not created
      - PruneLast=true                  # ensures that resources are not deleted prematurely during a sync.
      - ApplyOutOfSyncOnly=true         # only apply resources that are out of sync with the desired state
      - RespectIgnoreDifferences=true   # ArgoCD will respect the ignoreDifferences field defined in the Application manifest
      - ServerSideApply=true
    retry:
      limit: 5
      backoff:
        duration: 5s
        factor: 2
        maxDuration: 3m

  # ignoreDifferences:                  # ensures that dynamically managed fields (e.g., replicas, annotations, labels) are not overridden by ArgoCD
  #   - group: ""
  #     kind: Service
  #     name: my-service
  #     namespace: my-namespace
  #     jsonPointers:
  #       - /metadata/annotations
  #       - /metadata/labels
```