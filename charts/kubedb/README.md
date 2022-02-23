# KubeDB

[KubeDB by AppsCode](https://github.com/kubedb) - Making running production-grade databases easy on Kubernetes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb appscode/kubedb -n kubedb
```

## Introduction

This chart deploys a KubeDB operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubedb`:

```console
$ helm install kubedb appscode/kubedb -n kubedb
```

The command deploys a KubeDB operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb`:

```console
$ helm delete kubedb -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb` chart and their default values.

|                Parameter                |                                                                                                                                                                              Description                                                                                                                                                                              | Default |
|-----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| global.license                          | License for the product. Get a license by following the steps from [here](https://kubedb.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubedb \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubedb \` <br> `--set global.license=<license file content>` | `""`    |
| global.registry                         | Docker registry used to pull KubeDB related images                                                                                                                                                                                                                                                                                                                    | `""`    |
| global.registryFQDN                     | Docker registry fqdn used to pull KubeDB related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                               | `""`    |
| global.imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                     | `[]`    |
| global.skipCleaner                      | Skip generating cleaner job YAML                                                                                                                                                                                                                                                                                                                                      | `false` |
| global.monitoring.enabled               | If true, enables monitoring KubeDB operator                                                                                                                                                                                                                                                                                                                           | `false` |
| global.monitoring.agent                 | Name of monitoring agent ("prometheus.io" or "prometheus.io/operator" or "prometheus.io/builtin")                                                                                                                                                                                                                                                                     | `""`    |
| global.monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                   | `{}`    |
| kubedb-provisioner.enabled              | If enabled, installs the kubedb-provisioner chart                                                                                                                                                                                                                                                                                                                     | `true`  |
| kubedb-catalog.enabled                  | If enabled, installs the kubedb-catalog chart                                                                                                                                                                                                                                                                                                                         | `true`  |
| kubedb-webhook-server.enabled           | If enabled, installs the kubedb-webhook-server chart                                                                                                                                                                                                                                                                                                                  | `true`  |
| kubedb-ops-manager.enabled              | If enabled, installs the kubedb-ops-manager chart                                                                                                                                                                                                                                                                                                                     | `false` |
| kubedb-autoscaler.enabled               | If enabled, installs the kubedb-autoscaler chart                                                                                                                                                                                                                                                                                                                      | `false` |
| kubedb-dashboard.enabled                | If enabled, installs the kubedb-dashboard chart                                                                                                                                                                                                                                                                                                                       | `false` |
| kubedb-schema-manager.enabled           | If enabled, installs the kubedb-schema-manager chart                                                                                                                                                                                                                                                                                                                  | `false` |
| kubedb-metrics.enabled                  | If enabled, installs the kubedb-metrics chart                                                                                                                                                                                                                                                                                                                         | `false` |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb appscode/kubedb -n kubedb --set global.registry=kubedb
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb appscode/kubedb -n kubedb --values values.yaml
```
