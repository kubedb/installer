# KubeDB Catalog

[KubeDB Catalog](https://github.com/kubedb) - Catalog of database versions supported by KubeDB

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb-catalog appscode/kubedb-catalog -n kubedb
```

## Introduction

This chart deploys KubeDB catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install the chart with the release name `kubedb-catalog`:

```console
$ helm install kubedb-catalog appscode/kubedb-catalog -n kubedb
```

The command deploys KubeDB catalog on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb-catalog`:

```console
$ helm delete kubedb-catalog -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-catalog` chart and their default values.

|           Parameter            |                                                              Description                                                               | Default  |
|--------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|----------|
| nameOverride                   | Overrides name template                                                                                                                | `""`     |
| fullnameOverride               | Overrides fullname template                                                                                                            | `""`     |
| registryFQDN                   | Docker registry fqdn used to pull KubeDB related images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | `""`     |
| image.registry                 | Docker registry used to pull database image                                                                                            | `kubedb` |
| image.overrideOfficialRegistry | If true, uses image registry for pulling official docker images. This can be used to pull images from a private registry               | `false`  |
| catalog.elasticsearch          | If true, deploys Elasticsearch version catalog                                                                                         | `true`   |
| catalog.etcd                   | If true, deploys Etcd version catalog                                                                                                  | `true`   |
| catalog.memcached              | If true, deploys Memcached version catalog                                                                                             | `true`   |
| catalog.mongodb                | If true, deploys MongoDB version catalog                                                                                               | `true`   |
| catalog.mysql                  | If true, deploys MySQL version catalog                                                                                                 | `true`   |
| catalog.mariadb                | If true, deploys MariaDB version catalog                                                                                               | `true`   |
| catalog.perconaxtradb          | If true, deploys Percona XtraDB version catalog                                                                                        | `true`   |
| catalog.pgbouncer              | If true, deploys PgBouncer version catalog                                                                                             | `true`   |
| catalog.postgres               | If true, deploys PostgreSQL version catalog                                                                                            | `true`   |
| catalog.proxysql               | If true, deploys ProxySQL version catalog                                                                                              | `true`   |
| catalog.redis                  | If true, deploys Redis version catalog                                                                                                 | `true`   |
| skipDeprecated                 | Set true to avoid deploying deprecated versions                                                                                        | `true`   |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb-catalog appscode/kubedb-catalog -n kubedb --set image.registry=kubedb
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb-catalog appscode/kubedb-catalog -n kubedb --values values.yaml
```
