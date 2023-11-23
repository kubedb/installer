# Stash Catalog

[Stash Catalog](https://github.com/stashed) - Catalog of Stash Addons

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-kubestash-catalog --version=v2023.11.2
$ helm upgrade -i kubedb-kubestash-catalog appscode/kubedb-kubestash-catalog -n stash --create-namespace --version=v2023.11.2
```

## Introduction

This chart deploys Stash catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-kubestash-catalog`:

```bash
$ helm upgrade -i kubedb-kubestash-catalog appscode/kubedb-kubestash-catalog -n stash --create-namespace --version=v2023.11.2
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

|             Parameter             |                                                                                                                                                                     Description                                                                                                                                                                     |           Default            |
|-----------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------|
| proxies.dockerHub                 |                                                                                                                                                                                                                                                                                                                                                     | <code>""</code>              |
| proxies.dockerLibrary             |                                                                                                                                                                                                                                                                                                                                                     | <code>""</code>              |
| proxies.ghcr                      |                                                                                                                                                                                                                                                                                                                                                     | <code>ghcr.io</code>         |
| proxies.kubernetes                |                                                                                                                                                                                                                                                                                                                                                     | <code>registry.k8s.io</code> |
| proxies.appscode                  |                                                                                                                                                                                                                                                                                                                                                     | <code>r.appscode.com</code>  |
| waitTimeout                       | registryFQDN: harbor.appscode.ninja proxies: dockerHub: harbor.appscode.ninja/dockerhub dockerLibrary: "" ghcr: harbor.appscode.ninja/ghcr kubernetes: harbor.appscode.ninja/k8s appscode: harbor.appscode.ninja/ac proxies: ghcr: harbor.appscode.ninja/ghcr Number of seconds to wait for the database to be ready before backup/restore process. | <code>300</code>             |
| elasticsearch.enabled             | If true, deploys Elasticsearch addon                                                                                                                                                                                                                                                                                                                | <code>true</code>            |
| elasticsearch.backup.args         | Arguments to pass to `multielasticdump` command  during backup process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| elasticsearch.restore.args        | Arguments to pass to `multielasticdump` command during restore process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| opensearch.enabled                | If true, deploys Opensearch addon                                                                                                                                                                                                                                                                                                                   | <code>true</code>            |
| opensearch.backup.args            | Arguments to pass to `multielasticdump` command  during backup process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| opensearch.restore.args           | Arguments to pass to `multielasticdump` command during restore process                                                                                                                                                                                                                                                                              | <code>""</code>              |
| kubedbmanifest.enabled            | If true, deploys KubeDBManifest addon                                                                                                                                                                                                                                                                                                               | <code>true</code>            |
| kubedump.enabled                  | If true, deploy kubedump addon                                                                                                                                                                                                                                                                                                                      | <code>true</code>            |
| kubedump.backup.sanitize          | Specify whether to remove the decorator                                                                                                                                                                                                                                                                                                             | <code>true</code>            |
| kubedump.backup.labelSelector     | Specify label selector to filter resources                                                                                                                                                                                                                                                                                                          | <code>""</code>              |
| kubedump.backup.includeDependants | Specify whether to include the dependants resources along with it's parent                                                                                                                                                                                                                                                                          | <code>false</code>           |
| mongodb.enabled                   | If true, deploys MongoDB addon                                                                                                                                                                                                                                                                                                                      | <code>true</code>            |
| mongodb.maxConcurrency            | Maximum concurrency to perform backup or restore tasks                                                                                                                                                                                                                                                                                              | <code>3</code>               |
| mongodb.backup.args               | Arguments to pass to `mongodump` command during backup process                                                                                                                                                                                                                                                                                      | <code>""</code>              |
| mongodb.restore.args              | Arguments to pass to `mongorestore` command during restore process                                                                                                                                                                                                                                                                                  | <code>""</code>              |
| mysql.enabled                     | If true, deploys MySQL addon                                                                                                                                                                                                                                                                                                                        | <code>true</code>            |
| mysql.backup.args                 | Arguments to pass to `mysqldump` command  during bakcup process                                                                                                                                                                                                                                                                                     | <code>""</code>              |
| mysql.restore.args                | Arguments to pass to `mysql` command during restore process                                                                                                                                                                                                                                                                                         | <code>""</code>              |
| pvc.enabled                       | If true, deploys PVC addon                                                                                                                                                                                                                                                                                                                          | <code>true</code>            |
| redis.enabled                     | If true, deploys Redis addon                                                                                                                                                                                                                                                                                                                        | <code>true</code>            |
| redis.backup.args                 | Arguments to pass to `redis-dump` command  during bakcup process                                                                                                                                                                                                                                                                                    | <code>""</code>              |
| redis.restore.args                | Arguments to pass to `redis` command during restore process                                                                                                                                                                                                                                                                                         | <code>""</code>              |
| volumesnapshot.enabled            | If true, deploys VolumeSnapshot addon                                                                                                                                                                                                                                                                                                               | <code>true</code>            |
| workload.enabled                  | If true, deploys Workload addon                                                                                                                                                                                                                                                                                                                     | <code>true</code>            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-kubestash-catalog appscode/kubedb-kubestash-catalog -n stash --create-namespace --version=v2023.11.2 --set proxies.ghcr=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-kubestash-catalog appscode/kubedb-kubestash-catalog -n stash --create-namespace --version=v2023.11.2 --values values.yaml
```
