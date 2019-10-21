# KubeDB
[KubeDB by AppsCode](https://github.com/kubedb/installer) - Making running production-grade databases easy on Kubernetes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install appscode/kubedb
```

## Introduction

This chart bootstraps a [KubeDB controller](https://github.com/kubedb/installer) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.9+

## Installing the Chart
To install the chart with the release name `my-release`:

```console
$ helm install appscode/kubedb --name my-release
```

The command deploys KubeDB operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release`:

```console
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the KubeDB chart and their default values.

| Parameter                               | Description                                                                                                                                                                | Default                                                   |
| --------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| `replicaCount`                          | Number of kubedb operator replicas to create (only 1 is supported)                                                                                                         | `1`                                                       |
| `kubedb.registry`                       | Docker registry used to pull KubeDB operator image                                                                                                                         | `kubedb`                                                  |
| `kubedb.repository`                     | KubeDB operator container image                                                                                                                                            | `operator`                                                |
| `kubedb.tag`                            | KubeDB operator container image tag                                                                                                                                        | `v0.13.0-rc.0`                                                  |
| `cleaner.registry`                      | Docker registry used to pull Webhook cleaner image                                                                                                                         | `appscode`                                                |
| `cleaner.repository`                    | Webhook cleaner container image                                                                                                                                            | `kubectl`                                                 |
| `cleaner.tag`                           | Webhook cleaner container image tag                                                                                                                                        | `v1.12`                                                   |
| `imagePullSecrets`                      | Specify image pull secrets                                                                                                                                                 | `nil` (does not add image pull secrets to deployed pods)  |
| `imagePullPolicy`                       | Image pull policy                                                                                                                                                          | `IfNotPresent`                                            |
| `criticalAddon`                         | If true, installs KubeDB operator as critical addon                                                                                                                        | `false`                                                   |
| `logLevel`                              | Log level for operator                                                                                                                                                     | `3`                                                       |
| `affinity`                              | Affinity rules for pod assignment                                                                                                                                          | `{}`                                                      |
| `annotations`                           | Annotations applied to operator pod(s)                                                                                                                                     | `{}`                                                      |
| `nodeSelector`                          | Node labels for pod assignment                                                                                                                                             | `{}`                                                      |
| `tolerations`                           | Tolerations used pod assignment                                                                                                                                            | `{}`                                                      |
| `resources.requests.cpu`                | CPU resources request                                                                                                                                                      | `100m`                                                    |
| `resources.requests.memory`             | Memory resources request                                                                                                                                                   | `60Mi`                                                    |
| `resources.limits.cpu`                  | CPU resources limit                                                                                                                                                        | `""`                                                      |
| `resources.limits.memory`               | Memory resources limit                                                                                                                                                     | `""`                                                      |
| `rbac.create`                           | If `true`, create and use RBAC resources                                                                                                                                   | `true`                                                    |
| `serviceAccount.create`                 | If `true`, create a new service account                                                                                                                                    | `true`                                                    |
| `serviceAccount.name`                   | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template                                              | ``                                                        |
| `apiserver.groupPriorityMinimum`        | The minimum priority the group should have.                                                                                                                                | 10000                                                     |
| `apiserver.versionPriority`             | The ordering of this API inside of the group.                                                                                                                              | 15                                                        |
| `apiserver.enableValidatingWebhook`     | Enable validating webhooks for KubeDB CRDs                                                                                                                                 | `true`                                                    |
| `apiserver.enableMutatingWebhook`       | Enable mutating webhooks for KubeDB CRDs                                                                                                                                   | `true`                                                    |
| `apiserver.ca`                          | CA certificate used by main Kubernetes api server                                                                                                                          | `not-ca-cert`                                             |
| `apiserver.bypassValidatingWebhookXray` | If true, bypasses validating webhook xray checks                                                                                                                           | `false`                                                   |
| `apiserver.useKubeapiserverFqdnForAks`  | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522                                                                    | `true`                                                    |
| `apiserver.healthcheck.enabled`         | Enable readiness and liveliness probes                                                                                                                                     | `true`                                                    |
| `enableAnalytics`                       | Send usage events to Google Analytics                                                                                                                                      | `true`                                                    |
| `monitoring.enabled`                    | Specify whether to monitor KubeDB operator.                                                                                                                                | `false`                                                   |
| `monitoring.agent`                      | Specify which monitoring agent to use for monitoring KubeDB operator. It accepts either `prometheus.io/builtin` or `prometheus.io/coreos-operator`.                        | `none`                                                    |
| `monitoring.prometheus.namespace`       | Specify the namespace where Prometheus server is running or will be deployed.                                                                                              | Release namespace                                         |
| `monitoring.serviceMonitor.labels`      | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/coreos-operator`. | `app: <generated app name>` and `release: <release name>` |
| `additionalPodSecurityPolicies`         | Additional psp names passed to operator                                                                                                                                    | `[]`                                                      |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install --name my-release --set image.tag=v0.2.1 appscode/kubedb
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install --name my-release --values values.yaml appscode/kubedb
```

## RBAC
By default the chart will not install the recommended RBAC roles and rolebindings.

You need to have the flag `--authorization-mode=RBAC` on the api server. See the following document for how to enable [RBAC](https://kubernetes.io/docs/admin/authorization/rbac/).

To determine if your cluster supports RBAC, run the following command:

```console
$ kubectl api-versions | grep rbac
```

If the output contains "beta", you may install the chart with RBAC enabled (see below).

### Enable RBAC role/rolebinding creation

To enable the creation of RBAC resources (On clusters with RBAC). Do the following:

```console
$ helm install --name my-release appscode/kubedb --set rbac.create=true
```
