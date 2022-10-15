# KubeDB Catalog

[KubeDB Catalog](https://github.com/kubedb) - Catalog of database versions supported by KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-catalog --version=v2022.10.18
$ helm upgrade -i kubedb-catalog appscode/kubedb-catalog -n kubedb --create-namespace --version=v2022.10.18
```

## Introduction

This chart deploys KubeDB catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-catalog`:

```bash
$ helm upgrade -i kubedb-catalog appscode/kubedb-catalog -n kubedb --create-namespace --version=v2022.10.18
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

|                 Parameter                  |                                                              Description                                                               |      Default       |
|--------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| nameOverride                               | Overrides name template                                                                                                                | <code>""</code>    |
| fullnameOverride                           | Overrides fullname template                                                                                                            | <code>""</code>    |
| registryFQDN                               | Docker registry fqdn used to pull KubeDB related images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | <code>""</code>    |
| image.registry                             | Docker registry used to pull database image                                                                                            | <code>""</code>    |
| image.overrideOfficialRegistry             | If true, uses image registry for pulling official docker images. This can be used to pull images from a private registry               | <code>false</code> |
| catalog.elasticsearch                      | If true, deploys Elasticsearch version catalog                                                                                         | <code>true</code>  |
| catalog.etcd                               | If true, deploys Etcd version catalog                                                                                                  | <code>true</code>  |
| catalog.memcached                          | If true, deploys Memcached version catalog                                                                                             | <code>true</code>  |
| catalog.mongodb                            | If true, deploys MongoDB version catalog                                                                                               | <code>true</code>  |
| catalog.mysql                              | If true, deploys MySQL version catalog                                                                                                 | <code>true</code>  |
| catalog.mariadb                            | If true, deploys MariaDB version catalog                                                                                               | <code>true</code>  |
| catalog.perconaxtradb                      | If true, deploys Percona XtraDB version catalog                                                                                        | <code>true</code>  |
| catalog.pgbouncer                          | If true, deploys PgBouncer version catalog                                                                                             | <code>true</code>  |
| catalog.postgres                           | If true, deploys PostgreSQL version catalog                                                                                            | <code>true</code>  |
| catalog.proxysql                           | If true, deploys ProxySQL version catalog                                                                                              | <code>true</code>  |
| catalog.redis                              | If true, deploys Redis version catalog                                                                                                 | <code>true</code>  |
| psp.elasticsearch.allowPrivilegeEscalation |                                                                                                                                        | <code>true</code>  |
| psp.elasticsearch.privileged               |                                                                                                                                        | <code>true</code>  |
| psp.mariadb.allowPrivilegeEscalation       |                                                                                                                                        | <code>false</code> |
| psp.mariadb.privileged                     |                                                                                                                                        | <code>false</code> |
| psp.memcached.allowPrivilegeEscalation     |                                                                                                                                        | <code>false</code> |
| psp.memcached.privileged                   |                                                                                                                                        | <code>false</code> |
| psp.mongodb.allowPrivilegeEscalation       |                                                                                                                                        | <code>false</code> |
| psp.mongodb.privileged                     |                                                                                                                                        | <code>false</code> |
| psp.mysql.allowPrivilegeEscalation         |                                                                                                                                        | <code>false</code> |
| psp.mysql.privileged                       |                                                                                                                                        | <code>false</code> |
| psp.perconaxtradb.allowPrivilegeEscalation |                                                                                                                                        | <code>false</code> |
| psp.perconaxtradb.privileged               |                                                                                                                                        | <code>false</code> |
| psp.postgres.allowPrivilegeEscalation      |                                                                                                                                        | <code>false</code> |
| psp.postgres.privileged                    |                                                                                                                                        | <code>false</code> |
| psp.proxysql.allowPrivilegeEscalation      |                                                                                                                                        | <code>false</code> |
| psp.proxysql.privileged                    |                                                                                                                                        | <code>false</code> |
| psp.redis.allowPrivilegeEscalation         |                                                                                                                                        | <code>false</code> |
| psp.redis.privileged                       |                                                                                                                                        | <code>false</code> |
| skipDeprecated                             | Set true to avoid deploying deprecated versions                                                                                        | <code>true</code>  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-catalog appscode/kubedb-catalog -n kubedb --create-namespace --version=v2022.10.18 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-catalog appscode/kubedb-catalog -n kubedb --create-namespace --version=v2022.10.18 --values values.yaml
```
