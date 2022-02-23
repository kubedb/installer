# KubeDB Provisioner

[KubeDB Provisioner by AppsCode](https://github.com/kubedb) - Community features for KubeDB by AppsCode

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb-provisioner appscode/kubedb-provisioner -n kubedb
```

## Introduction

This chart deploys a KubeDB Provisioner operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.14+

## Installing the Chart

To install the chart with the release name `kubedb-provisioner`:

```console
$ helm install kubedb-provisioner appscode/kubedb-provisioner -n kubedb
```

The command deploys a KubeDB Provisioner operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb-provisioner`:

```console
$ helm delete kubedb-provisioner -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-provisioner` chart and their default values.

|              Parameter               |                                                                                                                                                                                   Description                                                                                                                                                                                   |            Default             |
|--------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------|
| nameOverride                         | Overrides name template                                                                                                                                                                                                                                                                                                                                                         | `""`                           |
| fullnameOverride                     | Overrides fullname template                                                                                                                                                                                                                                                                                                                                                     | `""`                           |
| replicaCount                         | Number of KubeDB operator replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                              | `1`                            |
| license                              | License for the product. Get a license by following the steps from [here](https://kubedb.run/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubedb-ops-manager \` <br> `--set-file license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubedb-ops-manager \` <br> `--set license=<license file content>` | `""`                           |
| registryFQDN                         | Docker registry fqdn used to pull KubeDB related images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                          | `""`                           |
| operator.registry                    | Docker registry used to pull KubeDB operator image                                                                                                                                                                                                                                                                                                                              | `kubedb`                       |
| operator.repository                  | KubeDB operator container image                                                                                                                                                                                                                                                                                                                                                 | `operator`                     |
| operator.tag                         | KubeDB operator container image tag                                                                                                                                                                                                                                                                                                                                             | `v0.25.0`                      |
| operator.resources                   | Compute Resources required by the operator container                                                                                                                                                                                                                                                                                                                            | `{}`                           |
| operator.securityContext             | requests: cpu: 100m memory: 128Mi Security options the operator container should run with                                                                                                                                                                                                                                                                                       | `{}`                           |
| imagePullSecrets                     | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb-provisioner \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                                 | `[]`                           |
| imagePullPolicy                      | Container image pull policy                                                                                                                                                                                                                                                                                                                                                     | `IfNotPresent`                 |
| criticalAddon                        | If true, installs KubeDB operator as critical addon                                                                                                                                                                                                                                                                                                                             | `false`                        |
| logLevel                             | Log level for operator                                                                                                                                                                                                                                                                                                                                                          | `3`                            |
| annotations                          | Annotations applied to operator deployment                                                                                                                                                                                                                                                                                                                                      | `{}`                           |
| podAnnotations                       | Annotations passed to operator pod(s).                                                                                                                                                                                                                                                                                                                                          | `{}`                           |
| nodeSelector                         | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                                  | `{"kubernetes.io/os":"linux"}` |
| tolerations                          | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                                  | `[]`                           |
| affinity                             | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                                               | `{}`                           |
| podSecurityContext                   | Security options the operator pod should run with.                                                                                                                                                                                                                                                                                                                              | `{}`                           |
| serviceAccount.create                | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                                           | `true`                         |
| serviceAccount.annotations           | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                                       | `{}`                           |
| serviceAccount.name                  | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                                          | ``                             |
| apiserver.useKubeapiserverFqdnForAks | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)                                                                                                                                                                                                                                                          | `true`                         |
| apiserver.healthcheck.enabled        | healthcheck configures the readiness and liveliness probes for the operator pod.                                                                                                                                                                                                                                                                                                | `true`                         |
| apiserver.healthcheck.probePort      | The port the probe endpoint binds to                                                                                                                                                                                                                                                                                                                                            | `8081`                         |
| enforceTerminationPolicy             | If true, namespace deletion will fail if it has a KubeDB resource with terminationPolicy DoNotTerminate                                                                                                                                                                                                                                                                         | `true`                         |
| monitoring.enabled                   | If true, enables monitoring KubeDB operator                                                                                                                                                                                                                                                                                                                                     | `false`                        |
| monitoring.bindPort                  | The port the metric endpoint binds to                                                                                                                                                                                                                                                                                                                                           | `8080`                         |
| monitoring.agent                     | Name of monitoring agent ("prometheus.io" or "prometheus.io/operator" or "prometheus.io/builtin")                                                                                                                                                                                                                                                                               | `""`                           |
| monitoring.serviceMonitor.labels     | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                             | `{}`                           |
| additionalPodSecurityPolicies        | Additional psp names passed to operator <br> Example: <br> `helm template ./chart/kubedb \` <br> `--set additionalPodSecurityPolicies[0]=abc \` <br> `--set additionalPodSecurityPolicies[1]=xyz`                                                                                                                                                                               | `[]`                           |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb-provisioner appscode/kubedb-provisioner -n kubedb --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb-provisioner appscode/kubedb-provisioner -n kubedb --values values.yaml
```
