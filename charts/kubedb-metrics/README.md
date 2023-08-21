# KubeDB Metrics

[KubeDB Metrics](https://github.com/kubedb) - KubeDB State Metrics

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-metrics --version=v2023.08.18
$ helm upgrade -i kubedb-metrics appscode/kubedb-metrics -n kubedb --create-namespace --version=v2023.08.18
```

## Introduction

This chart deploys KubeDB metrics configurations on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-metrics`:

```bash
$ helm upgrade -i kubedb-metrics appscode/kubedb-metrics -n kubedb --create-namespace --version=v2023.08.18
```

The command deploys KubeDB metrics configurations on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-metrics`:

```bash
$ helm uninstall kubedb-metrics -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


