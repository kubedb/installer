# KubeDB

[KubeDB by AppsCode](https://github.com/kubedb) - Making running production-grade databases easy on Kubernetes

## TL;DR;

```bash
$ helm repo add appscode-testing https://charts.appscode.com/testing/
$ helm repo update
$ helm search repo appscode-testing/kubedb-one --version=v2024.1.19-beta.1
$ helm upgrade -i kubedb appscode-testing/kubedb-one -n kubedb --create-namespace --version=v2024.1.19-beta.1
```

## Introduction

This chart deploys a KubeDB operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb`:

```bash
$ helm upgrade -i kubedb appscode-testing/kubedb-one -n kubedb --create-namespace --version=v2024.1.19-beta.1
```

The command deploys a KubeDB operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb`:

```bash
$ helm uninstall kubedb -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-one` chart and their default values.

|                Parameter                |                                                                                                                                                                              Description                                                                                                                                                                              |      Default       |
|-----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| global.license                          | License for the product. Get a license by following the steps from [here](https://kubedb.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubedb \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubedb \` <br> `--set global.license=<license file content>` | <code>""</code>    |
| global.licenseSecretName                | Name of Secret with the license as key.txt key                                                                                                                                                                                                                                                                                                                        | <code>""</code>    |
| global.registry                         | Docker registry used to pull KubeDB related images                                                                                                                                                                                                                                                                                                                    | <code>""</code>    |
| global.registryFQDN                     | Docker registry fqdn used to pull KubeDB related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                               | <code>""</code>    |
| global.insecureRegistries               | Specify an array of insecure registries. <br> Example: <br> `helm template charts/kubedb-ops-manager \` <br> `--set global.insecureRegistries[0]=hub.company.com \` <br> `--set global.insecureRegistries[1]=reg.example.com`                                                                                                                                         | <code>[]</code>    |
| global.imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                     | <code>[]</code>    |
| global.featureGates.Cassandra           |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.ClickHouse          |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.Druid               |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.Elasticsearch       |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.Etcd                |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.FerretDB            |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.Kafka               |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.MariaDB             |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.Memcached           |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.MicrosoftSQLServer  |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.MongoDB             |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.MySQL               |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.PerconaXtraDB       |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.PgBouncer           |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.Pgpool              |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.Postgres            |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.ProxySQL            |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.RabbitMQ            |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.Redis               |                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| global.featureGates.Singlestore         |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.Solr                |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.featureGates.ZooKeeper           |                                                                                                                                                                                                                                                                                                                                                                       | <code>false</code> |
| global.monitoring.agent                 | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                                                                                                                                  | <code>""</code>    |
| global.monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                   | <code>{}</code>    |
| kubedb-crd-manager.enabled              | If enabled, installs the kubedb-crd-manager chart                                                                                                                                                                                                                                                                                                                     | <code>true</code>  |
| kubedb-provisioner.enabled              | If enabled, installs the kubedb-provisioner chart                                                                                                                                                                                                                                                                                                                     | <code>true</code>  |
| kubedb-webhook-server.enabled           | If enabled, installs the kubedb-webhook-server chart                                                                                                                                                                                                                                                                                                                  | <code>true</code>  |
| kubedb-ops-manager.enabled              | If enabled, installs the kubedb-ops-manager chart                                                                                                                                                                                                                                                                                                                     | <code>true</code>  |
| kubedb-autoscaler.enabled               | If enabled, installs the kubedb-autoscaler chart                                                                                                                                                                                                                                                                                                                      | <code>true</code>  |
| kubedb-schema-manager.enabled           | If enabled, installs the kubedb-schema-manager chart                                                                                                                                                                                                                                                                                                                  | <code>false</code> |
| kubedb-metrics.enabled                  | If enabled, installs the kubedb-metrics chart                                                                                                                                                                                                                                                                                                                         | <code>true</code>  |
| stash-enterprise.enabled                | If enabled, installs the stash-enterprise chart                                                                                                                                                                                                                                                                                                                       | <code>true</code>  |
| stash-catalog.enabled                   | If enabled, installs the stash-catalog chart                                                                                                                                                                                                                                                                                                                          | <code>true</code>  |
| stash-metrics.enabled                   | If enabled, installs the stash-metrics chart                                                                                                                                                                                                                                                                                                                          | <code>true</code>  |
| kubedb-catalog.enabled                  | If enabled, installs the kubedb-catalog chart                                                                                                                                                                                                                                                                                                                         | <code>true</code>  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb appscode-testing/kubedb-one -n kubedb --create-namespace --version=v2024.1.19-beta.1 --set global.registry=kubedb
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb appscode-testing/kubedb-one -n kubedb --create-namespace --version=v2024.1.19-beta.1 --values values.yaml
```
