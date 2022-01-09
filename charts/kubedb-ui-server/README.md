# KubeDB UI Server

[KubeDB UI Server by AppsCode](https://github.com/kubedb/ui-server) - KubeDB UI Server for ByteBuilders

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb-ui-server appscode/kubedb-ui-server -n kubeops
```

## Introduction

This chart deploys a KubeDB UI Server on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubedb-ui-server`:

```console
$ helm install kubedb-ui-server appscode/kubedb-ui-server -n kubeops
```

The command deploys a KubeDB UI Server on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb-ui-server`:

```console
$ helm delete kubedb-ui-server -n kubeops
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-ui-server` chart and their default values.

|              Parameter               |                                                                                                                                                                          Description                                                                                                                                                                           |       Default       |
|--------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------|
| nameOverride                         | Overrides name template                                                                                                                                                                                                                                                                                                                                        | `""`                |
| fullnameOverride                     | Overrides fullname template                                                                                                                                                                                                                                                                                                                                    | `""`                |
| replicaCount                         | Number of UI Server replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                   | `1`                 |
| image.registry                       | Docker registry used to pull operator image                                                                                                                                                                                                                                                                                                                    | `kubedb`            |
| image.repository                     | Name of operator container image                                                                                                                                                                                                                                                                                                                               | `kubedb-ui-server`  |
| image.tag                            | Operator container image tag                                                                                                                                                                                                                                                                                                                                   | `v0.0.1`            |
| image.resources                      | Compute Resources required by the operator container                                                                                                                                                                                                                                                                                                           | `{}`                |
| image.securityContext                | Security options the operator container should run with                                                                                                                                                                                                                                                                                                        | `{}`                |
| imagePullSecrets                     | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/stash \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                             | `[]`                |
| imagePullPolicy                      | Container image pull policy                                                                                                                                                                                                                                                                                                                                    | `Always`            |
| criticalAddon                        | If true, installs Stash operator as critical addon                                                                                                                                                                                                                                                                                                             | `false`             |
| logLevel                             | Log level for operator                                                                                                                                                                                                                                                                                                                                         | `3`                 |
| annotations                          | Annotations applied to operator deployment                                                                                                                                                                                                                                                                                                                     | `{}`                |
| podAnnotations                       | Annotations passed to operator pod(s).                                                                                                                                                                                                                                                                                                                         | `{}`                |
| nodeSelector                         | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                 | `{}`                |
| tolerations                          | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                 | `[]`                |
| affinity                             | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                              | `{}`                |
| podSecurityContext                   | Security options the operator pod should run with.                                                                                                                                                                                                                                                                                                             | `{"fsGroup":65535}` |
| serviceAccount.create                | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                          | `true`              |
| serviceAccount.annotations           | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                      | `{}`                |
| serviceAccount.name                  | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                         | ``                  |
| apiserver.groupPriorityMinimum       | The minimum priority the webhook api group should have at least. Please see https://github.com/kubernetes/kube-aggregator/blob/release-1.9/pkg/apis/apiregistration/v1beta1/types.go#L58-L64 for more information on proper values of this field.                                                                                                              | `10000`             |
| apiserver.versionPriority            | The ordering of the webhook api inside of the group. Please see https://github.com/kubernetes/kube-aggregator/blob/release-1.9/pkg/apis/apiregistration/v1beta1/types.go#L66-L70 for more information on proper values of this field                                                                                                                           | `15`                |
| apiserver.useKubeapiserverFqdnForAks | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)                                                                                                                                                                                                                                         | `true`              |
| apiserver.healthcheck.enabled        | If true, enables the readiness and liveliness probes for the operator pod.                                                                                                                                                                                                                                                                                     | `false`             |
| apiserver.servingCerts.generate      | If true, generates on install/upgrade the certs that allow the kube-apiserver (and potentially ServiceMonitor) to authenticate operators pods. Otherwise specify certs in `apiserver.servingCerts.{caCrt, serverCrt, serverKey}`. See also: [example terraform](https://github.com/kubeops/installer/blob/master/charts/kubedb-ui-server/example-terraform.tf) | `true`              |
| apiserver.servingCerts.caCrt         | CA certficate used by serving certificate of webhook server.                                                                                                                                                                                                                                                                                                   | `""`                |
| apiserver.servingCerts.serverCrt     | Serving certficate used by webhook server.                                                                                                                                                                                                                                                                                                                     | `""`                |
| apiserver.servingCerts.serverKey     | Private key for the serving certificate used by webhook server.                                                                                                                                                                                                                                                                                                | `""`                |
| monitoring.enabled                   | If true, enables monitoring KubeDB operator                                                                                                                                                                                                                                                                                                                    | `false`             |
| monitoring.agent                     | Name of monitoring agent (either "prometheus.io/operator" or "prometheus.io/builtin")                                                                                                                                                                                                                                                                          | `"none"`            |
| monitoring.serviceMonitor.labels     | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                            | `{}`                |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb-ui-server appscode/kubedb-ui-server -n kubeops --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb-ui-server appscode/kubedb-ui-server -n kubeops --values values.yaml
```
