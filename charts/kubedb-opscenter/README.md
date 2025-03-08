# KubeDB Opscenter

[KubeDB Opscenter by AppsCode](https://github.com/kubedb) - KubeDB Opscenter

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-opscenter --version=v2025.2.19
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2025.2.19
```

## Introduction

This chart deploys a KubeDB Opscenter on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-opscenter`:

```bash
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2025.2.19
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

|                       Parameter                       |                                                                                                                                                                              Description                                                                                                                                                                              |                          Default                           |
|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|
| global.license                                        | License for the product. Get a license by following the steps from [here](https://kubedb.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubedb \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubedb \` <br> `--set global.license=<license file content>` | <code>""</code>                                            |
| global.licenseSecretName                              | Name of Secret with the license as key.txt key                                                                                                                                                                                                                                                                                                                        | <code>""</code>                                            |
| global.registry                                       | Docker registry used to pull KubeDB related images                                                                                                                                                                                                                                                                                                                    | <code>""</code>                                            |
| global.registryFQDN                                   | Docker registry fqdn used to pull KubeDB related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                               | <code>ghcr.io</code>                                       |
| global.imagePullSecrets                               | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                     | <code>[]</code>                                            |
| global.monitoring.agent                               | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                                                                                                                                  | <code>""</code>                                            |
| global.monitoring.serviceMonitor.labels               | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                   | <code>{"monitoring.appscode.com/prometheus":"auto"}</code> |
| global.networkPolicy.enabled                          |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| kubedb-metrics.enabled                                | If enabled, installs the kubedb-metrics chart                                                                                                                                                                                                                                                                                                                         | <code>true</code>                                          |
| kubedb-ui-server.enabled                              | If enabled, installs the kubedb-ui-server chart                                                                                                                                                                                                                                                                                                                       | <code>true</code>                                          |
| kubedb-grafana-dashboards.enabled                     | If enabled, installs the kubedb-grafana-dashboards chart                                                                                                                                                                                                                                                                                                              | <code>true</code>                                          |
| ace-user-roles.enabled                                | If enabled, installs the ace-user-roles chart                                                                                                                                                                                                                                                                                                                         | <code>true</code>                                          |
| ace-user-roles.enableClusterRoles.ace                 |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.appcatalog          |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.catalog             |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.cert-manager        |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.kubedb-ui           |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.kubedb              |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.kubestash           |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.kubevault           |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.license-proxyserver |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.metrics             |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>                                          |
| ace-user-roles.enableClusterRoles.prometheus          |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.secrets-store       |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.stash               |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |
| ace-user-roles.enableClusterRoles.virtual-secrets     |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code>                                         |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2025.2.19 --set global.registryFQDN=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-opscenter appscode/kubedb-opscenter -n kubedb --create-namespace --version=v2025.2.19 --values values.yaml
```
