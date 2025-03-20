# KubeDB GitOps operator

[KubeDB GitOps operator by AppsCode](https://github.com/appscode-cloud) - KubeDB GitOps operator

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-gitops --version=v0.1.0-rc.1
$ helm upgrade -i kubedb-gitops appscode/kubedb-gitops -n kubedb --create-namespace --version=v0.1.0-rc.1
```

## Introduction

This chart deploys a KubeDB GitOps operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-gitops`:

```bash
$ helm upgrade -i kubedb-gitops appscode/kubedb-gitops -n kubedb --create-namespace --version=v0.1.0-rc.1
```

The command deploys a KubeDB GitOps operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-gitops`:

```bash
$ helm uninstall kubedb-gitops -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-gitops` chart and their default values.

|             Parameter              |                                                             Description                                                              |                                                                                            Default                                                                                             |
|------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| replicaCount                       |                                                                                                                                      | <code>1</code>                                                                                                                                                                                 |
| registryFQDN                       | Docker registry fqdn used to pull app related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | <code>ghcr.io</code>                                                                                                                                                                           |
| image.registry                     | Docker registry used to pull app container image                                                                                     | <code>kubedb</code>                                                                                                                                                                            |
| image.repository                   |                                                                                                                                      | <code>kubedb-gitops</code>                                                                                                                                                                     |
| image.pullPolicy                   |                                                                                                                                      | <code>IfNotPresent</code>                                                                                                                                                                      |
| image.tag                          | Overrides the image tag whose default is the chart appVersion.                                                                       | <code>""</code>                                                                                                                                                                                |
| imagePullSecrets                   |                                                                                                                                      | <code>[]</code>                                                                                                                                                                                |
| nameOverride                       |                                                                                                                                      | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                   |                                                                                                                                      | <code>""</code>                                                                                                                                                                                |
| serviceAccount.create              | Specifies whether a service account should be created                                                                                | <code>true</code>                                                                                                                                                                              |
| serviceAccount.annotations         | Annotations to add to the service account                                                                                            | <code>{}</code>                                                                                                                                                                                |
| serviceAccount.name                | The name of the service account to use. If not set and create is true, a name is generated using the fullname template               | <code>""</code>                                                                                                                                                                                |
| podAnnotations                     |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| podLabels                          |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| podSecurityContext                 |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| securityContext                    |                                                                                                                                      | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| service.type                       |                                                                                                                                      | <code>ClusterIP</code>                                                                                                                                                                         |
| service.port                       |                                                                                                                                      | <code>8081</code>                                                                                                                                                                              |
| resources                          |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| livenessProbe.httpGet.path         |                                                                                                                                      | <code>/healthz</code>                                                                                                                                                                          |
| livenessProbe.httpGet.port         |                                                                                                                                      | <code>http</code>                                                                                                                                                                              |
| livenessProbe.initialDelaySeconds  |                                                                                                                                      | <code>15</code>                                                                                                                                                                                |
| livenessProbe.periodSeconds        |                                                                                                                                      | <code>20</code>                                                                                                                                                                                |
| readinessProbe.httpGet.path        |                                                                                                                                      | <code>/readyz</code>                                                                                                                                                                           |
| readinessProbe.httpGet.port        |                                                                                                                                      | <code>http</code>                                                                                                                                                                              |
| readinessProbe.initialDelaySeconds |                                                                                                                                      | <code>5</code>                                                                                                                                                                                 |
| readinessProbe.periodSeconds       |                                                                                                                                      | <code>10</code>                                                                                                                                                                                |
| volumes                            | Additional volumes on the output Deployment definition.                                                                              | <code>[]</code>                                                                                                                                                                                |
| volumeMounts                       | Additional volumeMounts on the output Deployment definition.                                                                         | <code>[]</code>                                                                                                                                                                                |
| nodeSelector                       |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| tolerations                        |                                                                                                                                      | <code>[]</code>                                                                                                                                                                                |
| affinity                           |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| featureGates.Cassandra             |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.ClickHouse            |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Druid                 |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Elasticsearch         |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.FerretDB              |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Kafka                 |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.MariaDB               |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Memcached             |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.MongoDB               |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.MSSQLServer           |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.MySQL                 |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.PerconaXtraDB         |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.PgBouncer             |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Pgpool                |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Postgres              |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.ProxySQL              |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.RabbitMQ              |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Redis                 |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Singlestore           |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.Solr                  |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |
| featureGates.ZooKeeper             |                                                                                                                                      | <code>true</code>                                                                                                                                                                              |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-gitops appscode/kubedb-gitops -n kubedb --create-namespace --version=v0.1.0-rc.1 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-gitops appscode/kubedb-gitops -n kubedb --create-namespace --version=v0.1.0-rc.1 --values values.yaml
```
