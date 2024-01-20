# KubeDB Catalog

[KubeDB Catalog](https://github.com/kubedb) - Catalog of database versions supported by KubeDB

## TL;DR;

```bash
$ helm repo add appscode-testing https://charts.appscode.com/testing/
$ helm repo update
$ helm search repo appscode-testing/kubedb-catalog --version=v2024.1.19-beta.1
$ helm upgrade -i kubedb-catalog appscode-testing/kubedb-catalog -n kubedb --create-namespace --version=v2024.1.19-beta.1
```

## Introduction

This chart deploys KubeDB catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-catalog`:

```bash
$ helm upgrade -i kubedb-catalog appscode-testing/kubedb-catalog -n kubedb --create-namespace --version=v2024.1.19-beta.1
```

The command deploys KubeDB catalog on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-catalog`:

```bash
$ helm uninstall kubedb-catalog -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-catalog` chart and their default values.

|                 Parameter                  |                   Description                   |           Default            |
|--------------------------------------------|-------------------------------------------------|------------------------------|
| nameOverride                               | Overrides name template                         | <code>""</code>              |
| fullnameOverride                           | Overrides fullname template                     | <code>""</code>              |
| proxies.dockerHub                          |                                                 | <code>""</code>              |
| proxies.dockerLibrary                      |                                                 | <code>""</code>              |
| proxies.ghcr                               |                                                 | <code>ghcr.io</code>         |
| proxies.kubernetes                         |                                                 | <code>registry.k8s.io</code> |
| proxies.appscode                           |                                                 | <code>r.appscode.com</code>  |
| featureGates.Cassandra                     |                                                 | <code>false</code>           |
| featureGates.ClickHouse                    |                                                 | <code>false</code>           |
| featureGates.Druid                         |                                                 | <code>true</code>            |
| featureGates.Elasticsearch                 |                                                 | <code>true</code>            |
| featureGates.Etcd                          |                                                 | <code>false</code>           |
| featureGates.FerretDB                      |                                                 | <code>true</code>            |
| featureGates.Kafka                         |                                                 | <code>true</code>            |
| featureGates.MariaDB                       |                                                 | <code>true</code>            |
| featureGates.Memcached                     |                                                 | <code>true</code>            |
| featureGates.MicrosoftSQLServer            |                                                 | <code>false</code>           |
| featureGates.MongoDB                       |                                                 | <code>true</code>            |
| featureGates.MySQL                         |                                                 | <code>true</code>            |
| featureGates.PerconaXtraDB                 |                                                 | <code>true</code>            |
| featureGates.PgBouncer                     |                                                 | <code>true</code>            |
| featureGates.Pgpool                        |                                                 | <code>true</code>            |
| featureGates.Postgres                      |                                                 | <code>true</code>            |
| featureGates.ProxySQL                      |                                                 | <code>true</code>            |
| featureGates.RabbitMQ                      |                                                 | <code>true</code>            |
| featureGates.Redis                         |                                                 | <code>true</code>            |
| featureGates.Singlestore                   |                                                 | <code>true</code>            |
| featureGates.Solr                          |                                                 | <code>true</code>            |
| featureGates.ZooKeeper                     |                                                 | <code>true</code>            |
| psp.enabled                                |                                                 | <code>false</code>           |
| psp.elasticsearch.allowPrivilegeEscalation |                                                 | <code>true</code>            |
| psp.elasticsearch.privileged               |                                                 | <code>true</code>            |
| psp.mariadb.allowPrivilegeEscalation       |                                                 | <code>false</code>           |
| psp.mariadb.privileged                     |                                                 | <code>false</code>           |
| psp.memcached.allowPrivilegeEscalation     |                                                 | <code>false</code>           |
| psp.memcached.privileged                   |                                                 | <code>false</code>           |
| psp.mongodb.allowPrivilegeEscalation       |                                                 | <code>false</code>           |
| psp.mongodb.privileged                     |                                                 | <code>false</code>           |
| psp.mysql.allowPrivilegeEscalation         |                                                 | <code>false</code>           |
| psp.mysql.privileged                       |                                                 | <code>false</code>           |
| psp.perconaxtradb.allowPrivilegeEscalation |                                                 | <code>false</code>           |
| psp.perconaxtradb.privileged               |                                                 | <code>false</code>           |
| psp.postgres.allowPrivilegeEscalation      |                                                 | <code>false</code>           |
| psp.postgres.privileged                    |                                                 | <code>false</code>           |
| psp.proxysql.allowPrivilegeEscalation      |                                                 | <code>false</code>           |
| psp.proxysql.privileged                    |                                                 | <code>false</code>           |
| psp.redis.allowPrivilegeEscalation         |                                                 | <code>false</code>           |
| psp.redis.privileged                       |                                                 | <code>false</code>           |
| psp.kafka.allowPrivilegeEscalation         |                                                 | <code>false</code>           |
| psp.kafka.privileged                       |                                                 | <code>false</code>           |
| skipDeprecated                             | Set true to avoid deploying deprecated versions | <code>true</code>            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-catalog appscode-testing/kubedb-catalog -n kubedb --create-namespace --version=v2024.1.19-beta.1 --set proxies.ghcr=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-catalog appscode-testing/kubedb-catalog -n kubedb --create-namespace --version=v2024.1.19-beta.1 --values values.yaml
```
