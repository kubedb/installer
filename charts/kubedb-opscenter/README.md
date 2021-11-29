# KubeDB Opscenter

[KubeDB Opscenter by AppsCode](https://github.com/kubedb) - KubeDB Opscenter

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb-opscenter appscode/kubedb-opscenter -n kubeops
```

## Introduction

This chart deploys a KubeDB Opscenter on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubedb-opscenter`:

```console
$ helm install kubedb-opscenter appscode/kubedb-opscenter -n kubeops
```

The command deploys a KubeDB Opscenter on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb-opscenter`:

```console
$ helm delete kubedb-opscenter -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-opscenter` chart and their default values.

|             Parameter             |                       Description                        | Default |
|-----------------------------------|----------------------------------------------------------|---------|
| kubedb-metrics.enabled            | If enabled, installs the kubedb-metrics chart            | `true`  |
| kubedb-ui-server.enabled          | If enabled, installs the kubedb-ui-server chart          | `true`  |
| kubedb-grafana-dashboards.enabled | If enabled, installs the kubedb-grafana-dashboards chart | `true`  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb-opscenter appscode/kubedb-opscenter -n kubeops --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb-opscenter appscode/kubedb-opscenter -n kubeops --values values.yaml
```
