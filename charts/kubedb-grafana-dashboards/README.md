# KubeDB Grafana Dashboards

[KubeDB Grafana Dashboards by AppsCode](https://github.com/kubedb/installer) - KubeDB Grafana Dashboards for ByteBuilders

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-grafana-dashboards --version=v2024.9.30
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2024.9.30
```

## Introduction

This chart deploys a KubeDB Grafana Dashboards on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-grafana-dashboards`:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2024.9.30
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

|            Parameter            |                            Description                             |                                                                                                  Default                                                                                                   |
|---------------------------------|--------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                    | Overrides name template                                            | <code>""</code>                                                                                                                                                                                            |
| fullnameOverride                | Overrides fullname template                                        | <code>""</code>                                                                                                                                                                                            |
| resources                       | List of resources for which dashboards will be applied             | <code>["connectcluster","druid","elasticsearch","kafka","mariadb","memcached","mongodb","mysql","perconaxtradb","pgpool","postgres","proxysql","rabbitmq","redis","singlestore","solr","zookeeper"]</code> |
| dashboard.folderID              | ID of Grafana folder where these dashboards will be applied        | <code>0</code>                                                                                                                                                                                             |
| dashboard.overwrite             | If true, dashboard with matching uid will be overwritten           | <code>true</code>                                                                                                                                                                                          |
| dashboard.templatize.title      | If true, datasource will be prefixed to dashboard name             | <code>false</code>                                                                                                                                                                                         |
| dashboard.templatize.datasource | If true, datasource will be hardcoded in the dashboard             | <code>false</code>                                                                                                                                                                                         |
| dashboard.alerts                |                                                                    | <code>false</code>                                                                                                                                                                                         |
| dashboard.replacements          |                                                                    | <code></code>                                                                                                                                                                                              |
| grafana.name                    | Name of Grafana Appbinding where these dashboards are applied      | <code>""</code>                                                                                                                                                                                            |
| grafana.namespace               | Namespace of Grafana Appbinding where these dashboards are applied | <code>""</code>                                                                                                                                                                                            |
| grafana.version                 |                                                                    | <code>8.0.7</code>                                                                                                                                                                                         |
| grafana.url                     |                                                                    | <code>""</code>                                                                                                                                                                                            |
| grafana.apikey                  |                                                                    | <code>""</code>                                                                                                                                                                                            |
| app.name                        |                                                                    | <code>""</code>                                                                                                                                                                                            |
| app.namespace                   |                                                                    | <code>""</code>                                                                                                                                                                                            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2024.9.30 --set resources=["connectcluster","druid","elasticsearch","kafka","mariadb","memcached","mongodb","mysql","perconaxtradb","pgpool","postgres","proxysql","rabbitmq","redis","singlestore","solr","zookeeper"]
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-grafana-dashboards appscode/kubedb-grafana-dashboards -n kubeops --create-namespace --version=v2024.9.30 --values values.yaml
```
