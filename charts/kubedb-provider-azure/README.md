# KubeDB Azure Provider

[KubeDB Azure Provider for Crossplane](https://github.com/kubedb/provider-azure) - KubeDB Azure provider for Crossplane

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-provider-azure --version=v2025.7.30-rc.0
$ helm upgrade -i kubedb-provider-azure appscode/kubedb-provider-azure -n crossplane-system --create-namespace --version=v2025.7.30-rc.0
```

## Introduction

This chart deploys a KubeDB Azure provider on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-provider-azure`:

```bash
$ helm upgrade -i kubedb-provider-azure appscode/kubedb-provider-azure -n crossplane-system --create-namespace --version=v2025.7.30-rc.0
```

The command deploys a KubeDB Azure provider on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-provider-azure`:

```bash
$ helm uninstall kubedb-provider-azure -n crossplane-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-provider-azure` chart and their default values.

|            Parameter             |                                                                                                            Description                                                                                                             |                                                                                            Default                                                                                             |
|----------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                     | Overrides name template                                                                                                                                                                                                            | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                 | Overrides fullname template                                                                                                                                                                                                        | <code>""</code>                                                                                                                                                                                |
| replicaCount                     |                                                                                                                                                                                                                                    | <code>1</code>                                                                                                                                                                                 |
| registryFQDN                     | Docker registry fqdn used to pull docker images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                     | <code>ghcr.io</code>                                                                                                                                                                           |
| image.registry                   | Docker registry used to pull operator image                                                                                                                                                                                        | <code>kubedb</code>                                                                                                                                                                            |
| image.repository                 | Name of operator container image                                                                                                                                                                                                   | <code>provider-azure</code>                                                                                                                                                                    |
| image.tag                        | Overrides the image tag whose default is the chart appVersion.                                                                                                                                                                     | <code>""</code>                                                                                                                                                                                |
| image.resources                  | Compute Resources required by the operator container                                                                                                                                                                               | <code>{}</code>                                                                                                                                                                                |
| image.securityContext            | Security options the operator container should run with                                                                                                                                                                            | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/stash \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                                                                                                                                                                                |
| imagePullPolicy                  | Container image pull policy                                                                                                                                                                                                        | <code>Always</code>                                                                                                                                                                            |
| serviceAccount.create            | Specifies whether a service account should be created                                                                                                                                                                              | <code>true</code>                                                                                                                                                                              |
| serviceAccount.annotations       | Annotations to add to the service account                                                                                                                                                                                          | <code>{}</code>                                                                                                                                                                                |
| serviceAccount.name              | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                             | <code>""</code>                                                                                                                                                                                |
| podAnnotations                   |                                                                                                                                                                                                                                    | <code>{}</code>                                                                                                                                                                                |
| podSecurityContext               |                                                                                                                                                                                                                                    | <code>{}</code>                                                                                                                                                                                |
| nodeSelector                     |                                                                                                                                                                                                                                    | <code>{}</code>                                                                                                                                                                                |
| tolerations                      |                                                                                                                                                                                                                                    | <code>[]</code>                                                                                                                                                                                |
| affinity                         |                                                                                                                                                                                                                                    | <code>{}</code>                                                                                                                                                                                |
| monitoring.agent                 | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                               | <code>""</code>                                                                                                                                                                                |
| monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                | <code>{}</code>                                                                                                                                                                                |
| azure.secretName                 |                                                                                                                                                                                                                                    | <code>"azure-credential"</code>                                                                                                                                                                |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-provider-azure appscode/kubedb-provider-azure -n crossplane-system --create-namespace --version=v2025.7.30-rc.0 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-provider-azure appscode/kubedb-provider-azure -n crossplane-system --create-namespace --version=v2025.7.30-rc.0 --values values.yaml
```
