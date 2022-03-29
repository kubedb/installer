# KubeDB Opscenter

[KubeDB Opscenter by AppsCode](https://github.com/kubedb) - KubeDB Opscenter

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-opscenter --version=v2022.03.28
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2022.03.28
```

## Introduction

This chart deploys a KubeDB Opscenter on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-opscenter`:

```bash
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2022.03.28
```

The command deploys a KubeDB Opscenter on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-opscenter`:

```bash
$ helm uninstall kubedb-opscenter -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-opscenter` chart and their default values.

|             Parameter             |                       Description                        |      Default      |
|-----------------------------------|----------------------------------------------------------|-------------------|
| kubedb-metrics.enabled            | If enabled, installs the kubedb-metrics chart            | <code>true</code> |
| kubedb-ui-server.enabled          | If enabled, installs the kubedb-ui-server chart          | <code>true</code> |
| kubedb-grafana-dashboards.enabled | If enabled, installs the kubedb-grafana-dashboards chart | <code>true</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2022.03.28 --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2022.03.28 --values values.yaml
```
