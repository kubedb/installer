# KubeDB CRD Manager

[KubeDB CRD Manager by AppsCode](https://github.com/kubedb) - KubeDB CRD Installer

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-crd-manager --version=v0.7.0
$ helm upgrade -i kubedb-ops-manager appscode/kubedb-crd-manager -n kubedb --create-namespace --version=v0.7.0
```

## Introduction

This chart deploys a KubeDB CRD Manager operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-ops-manager`:

```bash
$ helm upgrade -i kubedb-ops-manager appscode/kubedb-crd-manager -n kubedb --create-namespace --version=v0.7.0
```

The command deploys a KubeDB CRD Manager operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-ops-manager`:

```bash
$ helm uninstall kubedb-ops-manager -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-crd-manager` chart and their default values.

|         Parameter          |                                                                                                                   Description                                                                                                                   |                                                                    Default                                                                     |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| registryFQDN               | Docker registry fqdn used to pull app related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                            | <code>ghcr.io</code>                                                                                                                           |
| image.registry             | Docker registry used to pull app container image                                                                                                                                                                                                | <code>kubedb</code>                                                                                                                            |
| image.repository           | App container image                                                                                                                                                                                                                             | <code>kubedb-crd-manager</code>                                                                                                                |
| image.tag                  | Overrides the image tag whose default is the chart appVersion.                                                                                                                                                                                  | <code>""</code>                                                                                                                                |
| imagePullSecrets           | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb-ops-manager \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1` | <code>[]</code>                                                                                                                                |
| imagePullPolicy            | Container image pull policy                                                                                                                                                                                                                     | <code>IfNotPresent</code>                                                                                                                      |
| nameOverride               |                                                                                                                                                                                                                                                 | <code>""</code>                                                                                                                                |
| fullnameOverride           |                                                                                                                                                                                                                                                 | <code>""</code>                                                                                                                                |
| podAnnotations             |                                                                                                                                                                                                                                                 | <code>{}</code>                                                                                                                                |
| podSecurityContext         |                                                                                                                                                                                                                                                 | <code>{}</code>                                                                                                                                |
| securityContext            | Security options this container should run with                                                                                                                                                                                                 | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"runAsNonRoot":true,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| resources                  |                                                                                                                                                                                                                                                 | <code>{}</code>                                                                                                                                |
| nodeSelector               |                                                                                                                                                                                                                                                 | <code>{}</code>                                                                                                                                |
| tolerations                |                                                                                                                                                                                                                                                 | <code>[]</code>                                                                                                                                |
| affinity                   |                                                                                                                                                                                                                                                 | <code>{}</code>                                                                                                                                |
| ttlSecondsAfterFinished    |                                                                                                                                                                                                                                                 | <code>120</code>                                                                                                                               |
| serviceAccount.create      | Specifies whether a service account should be created                                                                                                                                                                                           | <code>true</code>                                                                                                                              |
| serviceAccount.annotations | Annotations to add to the service account                                                                                                                                                                                                       | <code>{}</code>                                                                                                                                |
| serviceAccount.name        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                          | <code></code>                                                                                                                                  |
| featureGates.Cassandra     |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.ClickHouse    |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Druid         |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Elasticsearch |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.FerretDB      |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Kafka         |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.MariaDB       |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Memcached     |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.MongoDB       |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.MSSQLServer   |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.MySQL         |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.PerconaXtraDB |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.PgBouncer     |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Pgpool        |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Postgres      |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.ProxySQL      |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.RabbitMQ      |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Redis         |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Singlestore   |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.Solr          |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| featureGates.ZooKeeper     |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |
| removeUnusedCRDs           |                                                                                                                                                                                                                                                 | <code>false</code>                                                                                                                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-ops-manager appscode/kubedb-crd-manager -n kubedb --create-namespace --version=v0.7.0 --set registryFQDN=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-ops-manager appscode/kubedb-crd-manager -n kubedb --create-namespace --version=v0.7.0 --values values.yaml
```
