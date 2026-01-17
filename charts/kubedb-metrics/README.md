# KubeDB Metrics

[KubeDB Metrics](https://github.com/kubedb) - KubeDB State Metrics

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-metrics --version=v2025.12.31-rc.1
$ helm upgrade -i kubedb-metrics appscode/kubedb-metrics -n kubedb --create-namespace --version=v2025.12.31-rc.1
```

## Introduction

This chart deploys KubeDB metrics configurations on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-metrics`:

```bash
$ helm upgrade -i kubedb-metrics appscode/kubedb-metrics -n kubedb --create-namespace --version=v2025.12.31-rc.1
```

The command deploys KubeDB metrics configurations on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-metrics`:

```bash
$ helm uninstall kubedb-metrics -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-metrics` chart and their default values.

|         Parameter          |         Description         |      Default      |
|----------------------------|-----------------------------|-------------------|
| nameOverride               | Overrides name template     | <code>""</code>   |
| fullnameOverride           | Overrides fullname template | <code>""</code>   |
| featureGates.Cassandra     |                             | <code>true</code> |
| featureGates.ClickHouse    |                             | <code>true</code> |
| featureGates.DB2           |                             | <code>true</code> |
| featureGates.Druid         |                             | <code>true</code> |
| featureGates.Elasticsearch |                             | <code>true</code> |
| featureGates.FerretDB      |                             | <code>true</code> |
| featureGates.HanaDB        |                             | <code>true</code> |
| featureGates.Hazelcast     |                             | <code>true</code> |
| featureGates.Ignite        |                             | <code>true</code> |
| featureGates.Kafka         |                             | <code>true</code> |
| featureGates.MariaDB       |                             | <code>true</code> |
| featureGates.Memcached     |                             | <code>true</code> |
| featureGates.Milvus        |                             | <code>true</code> |
| featureGates.MongoDB       |                             | <code>true</code> |
| featureGates.MSSQLServer   |                             | <code>true</code> |
| featureGates.MySQL         |                             | <code>true</code> |
| featureGates.Neo4j         |                             | <code>true</code> |
| featureGates.Oracle        |                             | <code>true</code> |
| featureGates.PerconaXtraDB |                             | <code>true</code> |
| featureGates.PgBouncer     |                             | <code>true</code> |
| featureGates.Pgpool        |                             | <code>true</code> |
| featureGates.Postgres      |                             | <code>true</code> |
| featureGates.ProxySQL      |                             | <code>true</code> |
| featureGates.Qdrant        |                             | <code>true</code> |
| featureGates.RabbitMQ      |                             | <code>true</code> |
| featureGates.Redis         |                             | <code>true</code> |
| featureGates.Singlestore   |                             | <code>true</code> |
| featureGates.Solr          |                             | <code>true</code> |
| featureGates.Weaviate      |                             | <code>true</code> |
| featureGates.ZooKeeper     |                             | <code>true</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-metrics appscode/kubedb-metrics -n kubedb --create-namespace --version=v2025.12.31-rc.1 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-metrics appscode/kubedb-metrics -n kubedb --create-namespace --version=v2025.12.31-rc.1 --values values.yaml
```
