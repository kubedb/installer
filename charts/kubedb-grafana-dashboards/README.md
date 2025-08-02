# KubeDB Grafana Dashboards

[KubeDB Grafana Dashboards by AppsCode](https://github.com/kubedb/installer) - KubeDB Grafana Dashboards for ByteBuilders

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-grafana-dashboards --version=v2025.7.30-rc.0
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2025.7.30-rc.0
```

## Introduction

This chart deploys a KubeDB Grafana Dashboards on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-grafana-dashboards`:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2025.7.30-rc.0
```

The command deploys a KubeDB Grafana Dashboards on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-grafana-dashboards`:

```bash
$ helm uninstall kubedb-grafana-dashboards -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-grafana-dashboards` chart and their default values.

|            Parameter            |                                                                                                            Description                                                                                                             |                                                                                            Default                                                                                             |
|---------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                    | Overrides name template                                                                                                                                                                                                            | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                | Overrides fullname template                                                                                                                                                                                                        | <code>""</code>                                                                                                                                                                                |
| resources                       | List of resources for which dashboards will be applied                                                                                                                                                                             | <code>["elasticsearch","kafka","mariadb","mongodb","mysql","postgres","redis"]</code>                                                                                                          |
| dashboard.folderID              | ID of Grafana folder where these dashboards will be applied                                                                                                                                                                        | <code>0</code>                                                                                                                                                                                 |
| dashboard.overwrite             | If true, dashboard with matching uid will be overwritten                                                                                                                                                                           | <code>true</code>                                                                                                                                                                              |
| dashboard.templatize.title      | If true, datasource will be prefixed to dashboard name                                                                                                                                                                             | <code>false</code>                                                                                                                                                                             |
| dashboard.templatize.datasource | If true, datasource will be hardcoded in the dashboard                                                                                                                                                                             | <code>false</code>                                                                                                                                                                             |
| dashboard.alerts                |                                                                                                                                                                                                                                    | <code>false</code>                                                                                                                                                                             |
| dashboard.replacements          |                                                                                                                                                                                                                                    | <code></code>                                                                                                                                                                                  |
| grafana.name                    | Name of Grafana Appbinding where these dashboards are applied                                                                                                                                                                      | <code>""</code>                                                                                                                                                                                |
| grafana.namespace               | Namespace of Grafana Appbinding where these dashboards are applied                                                                                                                                                                 | <code>""</code>                                                                                                                                                                                |
| grafana.version                 |                                                                                                                                                                                                                                    | <code>8.0.7</code>                                                                                                                                                                             |
| grafana.url                     |                                                                                                                                                                                                                                    | <code>""</code>                                                                                                                                                                                |
| grafana.apikey                  |                                                                                                                                                                                                                                    | <code>""</code>                                                                                                                                                                                |
| app.name                        |                                                                                                                                                                                                                                    | <code>""</code>                                                                                                                                                                                |
| app.namespace                   |                                                                                                                                                                                                                                    | <code>""</code>                                                                                                                                                                                |
| registryFQDN                    | Docker registry fqdn used to pull KubeDB related images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                             | <code>""</code>                                                                                                                                                                                |
| image.registry                  | Docker registry used to pull operator image                                                                                                                                                                                        | <code>curlimages</code>                                                                                                                                                                        |
| image.repository                | Name of operator container image                                                                                                                                                                                                   | <code>curl</code>                                                                                                                                                                              |
| image.tag                       | Operator container image tag                                                                                                                                                                                                       | <code>"latest"</code>                                                                                                                                                                          |
| image.securityContext           | Security options this container should run with                                                                                                                                                                                    | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| image.resources                 | Compute Resources required by this container                                                                                                                                                                                       | <code>{}</code>                                                                                                                                                                                |
| imagePullSecrets                | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/stash \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                                                                                                                                                                                |
| imagePullPolicy                 | Container image pull policy                                                                                                                                                                                                        | <code>Always</code>                                                                                                                                                                            |
| annotations                     | Annotations applied to operator deployment                                                                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| podAnnotations                  | Annotations passed to operator pod(s).                                                                                                                                                                                             | <code>{}</code>                                                                                                                                                                                |
| nodeSelector                    | Node labels for pod assignment                                                                                                                                                                                                     | <code>{}</code>                                                                                                                                                                                |
| tolerations                     | Tolerations for pod assignment                                                                                                                                                                                                     | <code>[]</code>                                                                                                                                                                                |
| affinity                        | Affinity rules for pod assignment                                                                                                                                                                                                  | <code>{}</code>                                                                                                                                                                                |
| podSecurityContext              | Security options the operator pod should run with.                                                                                                                                                                                 | <code>{"fsGroup":65534}</code>                                                                                                                                                                 |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2025.7.30-rc.0 --set resources=["elasticsearch","kafka","mariadb","mongodb","mysql","postgres","redis"]
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2025.7.30-rc.0 --values values.yaml
```
