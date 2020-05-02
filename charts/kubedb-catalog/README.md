# KubeDB Catalog

[KubeDB Catalog](https://github.com/kubedb) - Catalog of database versions supported by KubeDB

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb-catalog appscode/kubedb-catalog -n kube-system
```

## Introduction

This chart deploys KubeDB catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the chart with the release name `kubedb-catalog`:

```console
$ helm install kubedb-catalog appscode/kubedb-catalog -n kube-system
```

The command deploys KubeDB catalog on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb-catalog`:

```console
$ helm delete kubedb-catalog -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-catalog` chart and their default values.

|       Parameter       |                   Description                    | Default  |
|-----------------------|--------------------------------------------------|----------|
| nameOverride          | Overrides name template                          | `""`     |
| fullnameOverride      | Overrides fullname template                      | `""`     |
| image.registry        | Docker registry used to pull database image      | `kubedb` |
| catalog.elasticsearch | If true, installs Elasticsearch version catalog  | `true`   |
| catalog.etcd          | If true, installs Etcd version catalog           | `true`   |
| catalog.memcached     | If true, installs Memcached version catalog      | `true`   |
| catalog.mongo         | If true, installs MongoDB version catalog        | `true`   |
| catalog.mysql         | If true, installs MySQL version catalog          | `true`   |
| catalog.perconaxtradb | If true, installs Percona XtraDB version catalog | `true`   |
| catalog.pgbouncer     | If true, installs PgBouncer version catalog      | `true`   |
| catalog.postgres      | If true, installs PostgreSQL version catalog     | `true`   |
| catalog.proxysql      | If true, installs ProxySQL version catalog       | `true`   |
| catalog.redis         | If true, installs Redis version catalog          | `true`   |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb-catalog appscode/kubedb-catalog -n kube-system --set image.registry=kubedb
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb-catalog appscode/kubedb-catalog -n kube-system --values values.yaml
```
