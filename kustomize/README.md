## Deploy Kustomized KubeDB with ArgoCD

KubeDB officially maintains helm charts for installation purpose. You can simply deploy it using ArgoCD. [Here's](/kustomize/ArgoCD-Helm.md) a sample Application manifest. 

Due to Pre-installation requirements KubeDB can not be installed using kustomize. But, we can convert helm charts to deploy via kustomization and ArgoCD. Here's how to do it - 

### Converting to Kustomize

**Step 1:**

If you want to convert a release, just pull the helm chart package, open the tar file. you can simply do it with this `helm fetch [flags] [chart-uri]` command.

```bash
helm fetch \
   --untar \
   --untardir charts \
   --version v2025.1.9 \
    oci://ghcr.io/appscode-charts/kubedb
```

Alternatively, You can clone this repository, and directly convert from the charts directory.

We are going to generate a kustomize directory. The directory struct will be just as the following. The base directory will contain the generated template from kubedb helm chart. The deploy directory will contain required kustomization overlays.

```shell
kustomize/
├── base
└── deploy
```

**Step 2:**

Let's generate the template in kustimize/base directory. You can do it with a helm template command. Provide all the necessary flags and license file as per your requirement.

```bash
helm template kubedb ./charts/kubedb \
        --output-dir kustomize/base \
        --include-crds \
        --dependency-update \
        --namespace kubedb \
        --set global.featureGates.RabbitMQ=true \
        --set-file global.license=/path/to/license/file.txt
```

Let's shorten the directory depth of `kustomize/base` for making things simple.
```bash
mv ./kustomize/base/kubedb/charts/* ./kustomize/base/ && rm -rf ./kustomize/base/kubedb
```

Now, go to `kustomize/base` directory and generate the kustomization file. You need to have kustomization cli installed for this. you can install it from [here](https://kubectl.docs.kubernetes.io/installation/kustomize/). 

```bash
$ cd kustomize/base
$ kustomize create --autodetect --recursive
```

You will find all the resources to apply have been automatically added to the newly created `kustomization.yaml` file.

**Step 3:**

Let's move forward to the `kustomize/deploy` directory. KubeDB helm chart have several subcharts embedded withing. One of these subcharts is `kubedb-crd-manager` which runs a job and installs required crds before installing other resources. In helm it can be controlled using helm pre-install hooks. Since, Kustomize do not have any features for such types of hooking, we can use argocd `PreSync` hooks in this case. So that, when kustomized kubedb gets synchronized by argocd to the cluster, `kubedb-crd-manager` can be synced first. The job will run get succeeded/failed and then the other resources will be installed. You can use such annotation with any kubernetes manifest in it's `metadata.annotation` part.

```yaml
metadata:
  annotations:
    "argocd.argoproj.io/hook": "PreSync"
```

Now, create the deploy directory and create the overlays for `kubedb-crd-manager` clusterRole, clusterRoleBinding, ServiceAccount and the Job. You can just copy them from [here](/kustomize/deploy/kubedb-crd-manager).

Create a `kustomzation.yaml` file in `kustomize/deploy` directory. Refer it's base to `kustomize/base` and provide all the overlay patches. Here's how it should look like -

```yaml
resources:
  - ../base

patches:
  - path: kubedb-crd-manager/cluster-role.yaml
    target:
      group: rbac.authorization.k8s.io
      version: v1
      kind: ClusterRole
      name: kubedb-kubedb-crd-manager
  - path: kubedb-crd-manager/cluster-role-binding.yaml
    target:
      group: rbac.authorization.k8s.io
      version: v1
      kind: ClusterRoleBinding
      name: kubedb-kubedb-crd-manager
  - path: kubedb-crd-manager/serviceaccount.yaml
    target:
      group: ""
      version: v1
      kind: ServiceAccount
      name: kubedb-kubedb-crd-manager
  - path: kubedb-crd-manager/job.yaml
    target:
      group: batch
      version: v1
      kind: Job
      name: kubedb-kubedb-crd-manager
      namespace: kubedb
```

The final structure of this `kustomize` directory will look something like this - 

```bash
$ tree ./kustomize -L 2
./kustomize
├── base
│   ├── kubedb-autoscaler
│   ├── kubedb-catalog
│   ├── kubedb-crd-manager
│   ├── kubedb-kubestash-catalog
│   ├── kubedb-ops-manager
│   ├── kubedb-provisioner
│   ├── kubedb-webhook-server
│   ├── kustomization.yaml
│   ├── petset
│   └── sidekick
└── deploy
    ├── kubedb-crd-manager
    └── kustomization.yaml
```

**Step 4:**

Finally, we can now go for GitOps with ArgoCD. Push this `kustomize` directory in your git repo and refer it as source in argocd `Application`. Here's a sample manifest you can use to deploy the [kustomize](/kustomize) generated in this open source repository via ArgoCD to you cluster.

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: kubedb
  namespace: argocd
spec:
  project: default
  source:
      repoURL: https://github.com/kubedb/installer
      targetRevision: kustomize
      path: kustomize/deploy
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