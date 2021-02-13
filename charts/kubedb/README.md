# KubeDB

[KubeDB by AppsCode](https://github.com/kubedb) - Making running production-grade databases easy on Kubernetes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb appscode/kubedb -n kube-system
```

## Introduction

This chart deploys a KubeDB operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubedb`:

```console
$ helm install kubedb appscode/kubedb -n kube-system
```

The command deploys a KubeDB operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb`:

```console
$ helm delete kubedb -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb` chart and their default values.

|         Parameter         | Description | Default |
|---------------------------|-------------|---------|
| kubedb-community.enabled  |             | `true`  |
| kubedb-community.license  |             | `""`    |
| kubedb-catalog.enabled    |             | `false` |
| kubedb-enterprise.enabled |             | `true`  |
| kubedb-enterprise.license |             | `""`    |
| kubedb-autoscaler.enabled |             | `true`  |
| kubedb-autoscaler.license |             | `""`    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb appscode/kubedb -n kube-system --set -- generate from values file --
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb appscode/kubedb -n kube-system --values values.yaml
```
