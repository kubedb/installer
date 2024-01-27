# KubeDB Catalog CRDs

[KubeDB Catalog CRDs](https://github.com/kubedb) - KubeDB Custom Resource Definitions

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-catalog-crds --version=v2024.1.26-rc.0
$ helm upgrade -i kubedb-catalog-crds appscode/kubedb-catalog-crds -n kubedb --create-namespace --version=v2024.1.26-rc.0
```

## Introduction

This chart deploys KubeDB catalog crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-catalog-crds`:

```bash
$ helm upgrade -i kubedb-catalog-crds appscode/kubedb-catalog-crds -n kubedb --create-namespace --version=v2024.1.26-rc.0
```

The command deploys KubeDB catalog crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-catalog-crds`:

```bash
$ helm uninstall kubedb-catalog-crds -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


