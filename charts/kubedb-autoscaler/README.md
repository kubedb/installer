# KubeDB Autoscaler

[KubeDB Autoscaler by AppsCode](https://github.com/kubedb) - Autoscale KubeDB operated Databases

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-autoscaler --version=v0.12.0
$ helm upgrade -i kubedb-autoscaler appscode/kubedb-autoscaler -n kubedb --create-namespace --version=v0.12.0
```

## Introduction

This chart deploys a KubeDB Autoscaler operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.12+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-autoscaler`:

```bash
$ helm upgrade -i kubedb-autoscaler appscode/kubedb-autoscaler -n kubedb --create-namespace --version=v0.12.0
```

The command deploys a KubeDB Autoscaler operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-autoscaler`:

```bash
$ helm uninstall kubedb-autoscaler -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-autoscaler` chart and their default values.

|                Parameter                 |                                                                                                                                                                                 Description                                                                                                                                                                                  |                           Default                           |
|------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| nameOverride                             | Overrides name template                                                                                                                                                                                                                                                                                                                                                      | <code>""</code>                                             |
| fullnameOverride                         | Overrides fullname template                                                                                                                                                                                                                                                                                                                                                  | <code>""</code>                                             |
| replicaCount                             | Number of KubeDB operator replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                           | <code>1</code>                                              |
| license                                  | License for the product. Get a license by following the steps from [here](https://stash.run/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubedb-autoscaler \` <br> `--set-file license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubedb-autoscaler \` <br> `--set license=<license file content>` | <code>""</code>                                             |
| updateInterval                           | Interval between each autoscaler loop                                                                                                                                                                                                                                                                                                                                        | <code>1m</code>                                             |
| registryFQDN                             | Docker registry fqdn used to pull KubeDB related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                      | <code>""</code>                                             |
| operator.registry                        | Docker registry used to pull KubeDB enterprise operator image                                                                                                                                                                                                                                                                                                                | <code>kubedb</code>                                         |
| operator.repository                      | KubeDB enterprise operator container image                                                                                                                                                                                                                                                                                                                                   | <code>kubedb-autoscaler</code>                              |
| operator.tag                             | KubeDB enterprise operator container image tag                                                                                                                                                                                                                                                                                                                               | <code>v0.12.0</code>                                        |
| operator.resources                       | Compute Resources required by the enterprise operator container                                                                                                                                                                                                                                                                                                              | <code>{}</code>                                             |
| operator.securityContext                 | requests: cpu: 100m memory: 128Mi Security options the enterprise operator container should run with                                                                                                                                                                                                                                                                         | <code>{}</code>                                             |
| imagePullSecrets                         | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb-autoscaler \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                               | <code>[]</code>                                             |
| imagePullPolicy                          | Container image pull policy                                                                                                                                                                                                                                                                                                                                                  | <code>IfNotPresent</code>                                   |
| criticalAddon                            | If true, installs KubeDB operator as critical addon                                                                                                                                                                                                                                                                                                                          | <code>false</code>                                          |
| logLevel                                 | Log level for operator                                                                                                                                                                                                                                                                                                                                                       | <code>3</code>                                              |
| annotations                              | Annotations applied to operator deployment                                                                                                                                                                                                                                                                                                                                   | <code>{}</code>                                             |
| podAnnotations                           | Annotations passed to operator pod(s).                                                                                                                                                                                                                                                                                                                                       | <code>{}</code>                                             |
| nodeSelector                             | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                               | <code>{"kubernetes.io/os":"linux"}</code>                   |
| tolerations                              | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                               | <code>[]</code>                                             |
| affinity                                 | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                                            | <code>{}</code>                                             |
| podSecurityContext                       | Security options the operator pod should run with.                                                                                                                                                                                                                                                                                                                           | <code>{}</code>                                             |
| serviceAccount.create                    | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                                        | <code>true</code>                                           |
| serviceAccount.annotations               | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                                    | <code>{}</code>                                             |
| serviceAccount.name                      | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                                       | <code></code>                                               |
| apiserver.useKubeapiserverFqdnForAks     | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)                                                                                                                                                                                                                                                       | <code>true</code>                                           |
| apiserver.healthcheck.enabled            | healthcheck configures the readiness and liveliness probes for the operator pod.                                                                                                                                                                                                                                                                                             | <code>true</code>                                           |
| apiserver.healthcheck.probePort          | The port the probe endpoint binds to                                                                                                                                                                                                                                                                                                                                         | <code>8081</code>                                           |
| monitoring.bindPort                      | The port the metric endpoint binds to                                                                                                                                                                                                                                                                                                                                        | <code>8080</code>                                           |
| monitoring.agent                         | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                                                                                                                                         | <code>""</code>                                             |
| monitoring.serviceMonitor.labels         | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                          | <code>{}</code>                                             |
| storageAutoscaler.prometheus.address     | Prometheus address for storage metrics                                                                                                                                                                                                                                                                                                                                       | <code>http://prometheus-operated.monitoring.svc:9090</code> |
| storageAutoscaler.prometheus.bearerToken | Bearer token for prometheus server                                                                                                                                                                                                                                                                                                                                           | <code>""</code>                                             |
| storageAutoscaler.prometheus.caCert      | CA cert for prometheus server TLS connections                                                                                                                                                                                                                                                                                                                                | <code>""</code>                                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-autoscaler appscode/kubedb-autoscaler -n kubedb --create-namespace --version=v0.12.0 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-autoscaler appscode/kubedb-autoscaler -n kubedb --create-namespace --version=v0.12.0 --values values.yaml
```
