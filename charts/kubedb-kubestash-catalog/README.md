# Stash Catalog

[Stash Catalog](https://github.com/stashed) - Catalog of Stash Addons

## TL;DR;

```bash
$ helm repo add appscode-testing https://charts.appscode.com/testing/
$ helm repo update
$ helm search repo appscode-testing/kubedb-kubestash-catalog --version=v2024.1.7-beta.0
$ helm upgrade -i kubedb-kubestash-catalog appscode-testing/kubedb-kubestash-catalog -n stash --create-namespace --version=v2024.1.7-beta.0
```

## Introduction

This chart deploys Stash catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-kubestash-catalog`:

```bash
$ helm upgrade -i kubedb-kubestash-catalog appscode-testing/kubedb-kubestash-catalog -n stash --create-namespace --version=v2024.1.7-beta.0
```

The command deploys Stash catalog on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-kubestash-catalog`:

```bash
$ helm uninstall kubedb-kubestash-catalog -n stash
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-kubestash-catalog` chart and their default values.

|            Parameter            |                                                                                                                                                                     Description                                                                                                                                                                     |           Default            |
|---------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------|
| proxies.dockerHub               |                                                                                                                                                                                                                                                                                                                                                     | <code>""</code>              |
| proxies.dockerLibrary           |                                                                                                                                                                                                                                                                                                                                                     | <code>""</code>              |
| proxies.ghcr                    |                                                                                                                                                                                                                                                                                                                                                     | <code>ghcr.io</code>         |
| proxies.kubernetes              |                                                                                                                                                                                                                                                                                                                                                     | <code>registry.k8s.io</code> |
| proxies.appscode                |                                                                                                                                                                                                                                                                                                                                                     | <code>r.appscode.com</code>  |
| waitTimeout                     | registryFQDN: harbor.appscode.ninja proxies: dockerHub: harbor.appscode.ninja/dockerhub dockerLibrary: "" ghcr: harbor.appscode.ninja/ghcr kubernetes: harbor.appscode.ninja/k8s appscode: harbor.appscode.ninja/ac proxies: ghcr: harbor.appscode.ninja/ghcr Number of seconds to wait for the database to be ready before backup/restore process. | <code>300</code>             |
| featureGates.Cassandra          |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.ClickHouse         |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.Druid              |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.Elasticsearch      |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.Etcd               |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.FerretDB           |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.Kafka              |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.MariaDB            |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.Memcached          |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.MicrosoftSQLServer |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.MongoDB            |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.MySQL              |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.PerconaXtraDB      |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.PgBouncer          |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.PgPool             |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.Postgres           |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.ProxySQL           |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.RabbitMQ           |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.Redis              |                                                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |
| featureGates.SingleStore        |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.Solr               |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| featureGates.ZooKeeper          |                                                                                                                                                                                                                                                                                                                                                     | <code>false</code>           |
| elasticsearch.backup.args       | Arguments to pass to `multielasticdump` command  during backup process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| elasticsearch.restore.args      | Arguments to pass to `multielasticdump` command during restore process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| opensearch.backup.args          | Arguments to pass to `multielasticdump` command  during backup process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| opensearch.restore.args         | Arguments to pass to `multielasticdump` command during restore process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| kubedbmanifest.enabled          | If true, deploys KubeDBManifest addon                                                                                                                                                                                                                                                                                                               | <code>true</code>            |
| mongodb.maxConcurrency          | Maximum concurrency to perform backup or restore tasks                                                                                                                                                                                                                                                                                              | <code>3</code>               |
| mongodb.backup.args             | Arguments to pass to `mongodump` command during backup process                                                                                                                                                                                                                                                                                      | <code>""</code>              |
| mongodb.restore.args            | Arguments to pass to `mongorestore` command during restore process                                                                                                                                                                                                                                                                                  | <code>""</code>              |
| postgres.backup.cmd             | Postgres dump command, can either be: pg_dumpall  or pg_dump                                                                                                                                                                                                                                                                                        | <code>"pg_dumpall"</code>    |
| postgres.backup.args            | Arguments to pass to `backup.cmd` command during backup process                                                                                                                                                                                                                                                                                     | <code>""</code>              |
| postgres.restore.args           | Arguments to pass to `psql` command during restore process                                                                                                                                                                                                                                                                                          | <code>""</code>              |
| mysql.backup.args               | Arguments to pass to `mysqldump` command  during bakcup process                                                                                                                                                                                                                                                                                     | <code>""</code>              |
| mysql.restore.args              | Arguments to pass to `mysql` command during restore process                                                                                                                                                                                                                                                                                         | <code>""</code>              |
| redis.backup.args               | Arguments to pass to `redis-dump` command  during bakcup process                                                                                                                                                                                                                                                                                    | <code>""</code>              |
| redis.restore.args              | Arguments to pass to `redis` command during restore process                                                                                                                                                                                                                                                                                         | <code>""</code>              |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-kubestash-catalog appscode-testing/kubedb-kubestash-catalog -n stash --create-namespace --version=v2024.1.7-beta.0 --set proxies.ghcr=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-kubestash-catalog appscode-testing/kubedb-kubestash-catalog -n stash --create-namespace --version=v2024.1.7-beta.0 --values values.yaml
```
