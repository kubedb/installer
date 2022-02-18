# KubeDB Grafana Dashboards

[KubeDB Grafana Dashboards by AppsCode](https://github.com/kubedb/installer) - KubeDB Grafana Dashboards for ByteBuilders

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-grafana-dashboards --version=v2022.05.24
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2022.05.24
```

## Introduction

This chart deploys a KubeDB Grafana Dashboards on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-grafana-dashboards`:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2022.05.24
```

The command deploys a KubeDB Grafana Dashboards on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-grafana-dashboards`:

```bash
$ helm uninstall kubedb-grafana-dashboards -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-grafana-dashboards` chart and their default values.

|            Parameter            |                            Description                             |                              Default                               |
|---------------------------------|--------------------------------------------------------------------|--------------------------------------------------------------------|
| nameOverride                    | Overrides name template                                            | `""`                                                               |
| fullnameOverride                | Overrides fullname template                                        | `""`                                                               |
| resources                       | List of resources for which dashboards will be applied             | `["elasticsearch","mariadb","mongodb","mysql","postgres","redis"]` |
| dashboard.folderID              | ID of Grafana folder where these dashboards will be applied        | `0`                                                                |
| dashboard.overwrite             | If true, dashboard with matching uid will be overwritten           | `true`                                                             |
| dashboard.templatize.title      | If true, datasource will be prefixed to dashboard name             | `false`                                                            |
| dashboard.templatize.datasource | If true, datasource will be hardcoded in the dashboard             | `true`                                                             |
| grafana.name                    | Name of Grafana Appbinding where these dashboards are applied      | `""`                                                               |
| grafana.namespace               | Namespace of Grafana Appbinding where these dashboards are applied | `""`                                                               |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2022.05.24 --set resources=["elasticsearch","mariadb","mongodb","mysql","postgres","redis"]
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2022.05.24 --values values.yaml
```
