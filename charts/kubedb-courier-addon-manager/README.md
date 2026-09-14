# KubeDB Courier Addon Manager

[KubeDB Courier Addon Manager by AppsCode](https://github.com/kubedb/courier) - Hub-side KubeDB Courier addon manager

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-courier-addon-manager --version=v0.7.0-rc.2
$ helm upgrade -i kubedb-courier-addon-manager appscode/kubedb-courier-addon-manager -n kubedb --create-namespace --version=v0.7.0-rc.2
```

## Introduction

This chart deploys a KubeDB Courier addon manager operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+
- Open Cluster Management (registration and work) installed on the hub

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-courier-addon-manager`:

```bash
$ helm upgrade -i kubedb-courier-addon-manager appscode/kubedb-courier-addon-manager -n kubedb --create-namespace --version=v0.7.0-rc.2
```

The command deploys a KubeDB Courier addon manager operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-courier-addon-manager`:

```bash
$ helm uninstall kubedb-courier-addon-manager -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-courier-addon-manager` chart and their default values.

|                  Parameter                  |                                                          Description                                                          |                     Default                     |
|---------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------|
| nameOverride                                | nameOverride / fullnameOverride let you rename the release-derived resources.                                                 | <code>""</code>                                 |
| fullnameOverride                            |                                                                                                                               | <code>""</code>                                 |
| image.registry                              |                                                                                                                               | <code>ghcr.io/kubedb</code>                     |
| image.repository                            |                                                                                                                               | <code>kubedb-courier</code>                     |
| image.tag                                   |                                                                                                                               | <code>"" # defaults to .Chart.AppVersion</code> |
| image.pullPolicy                            |                                                                                                                               | <code>IfNotPresent</code>                       |
| imagePullSecrets                            |                                                                                                                               | <code>[]</code>                                 |
| manager.replicaCount                        |                                                                                                                               | <code>1</code>                                  |
| manager.maxFeedbackWatch                    | maxFeedbackWatch caps concurrent watch-based status feedbacks; branches past the cap fall back to Poll (design.md Section 7). | <code>100</code>                                |
| manager.leaderElection                      |                                                                                                                               | <code>true</code>                               |
| manager.resources.requests.cpu              |                                                                                                                               | <code>100m</code>                               |
| manager.resources.requests.memory           |                                                                                                                               | <code>128Mi</code>                              |
| manager.resources.limits.cpu                |                                                                                                                               | <code>"1"</code>                                |
| manager.resources.limits.memory             |                                                                                                                               | <code>512Mi</code>                              |
| manager.nodeSelector                        |                                                                                                                               | <code>{}</code>                                 |
| manager.tolerations                         |                                                                                                                               | <code>[]</code>                                 |
| manager.affinity                            |                                                                                                                               | <code>{}</code>                                 |
| manager.podAnnotations                      |                                                                                                                               | <code>{}</code>                                 |
| manager.securityContext.runAsNonRoot        |                                                                                                                               | <code>true</code>                               |
| manager.securityContext.seccompProfile.type |                                                                                                                               | <code>RuntimeDefault</code>                     |
| agent.image.registry                        |                                                                                                                               | <code>""</code>                                 |
| agent.image.repository                      |                                                                                                                               | <code>""</code>                                 |
| agent.image.tag                             |                                                                                                                               | <code>""</code>                                 |
| agent.image.pullPolicy                      | falls back to .Values.image.pullPolicy when empty                                                                             | <code>""</code>                                 |
| agent.installNamespace                      | installNamespace is where the addon agent is installed on each managed cluster.                                               | <code>kubedb</code>                             |
| agent.placement.name                        | name of an existing OCM Placement in the placementNamespace.                                                                  | <code>""</code>                                 |
| agent.placement.namespace                   |                                                                                                                               | <code>open-cluster-management</code>            |
| ocm.addonManager.namespace                  |                                                                                                                               | <code>open-cluster-management-hub</code>        |
| ocm.addonManager.serviceAccountName         |                                                                                                                               | <code>addon-manager-controller-sa</code>        |
| ocm.workAgent.namespace                     |                                                                                                                               | <code>open-cluster-management-agent</code>      |
| ocm.workAgent.serviceAccountName            |                                                                                                                               | <code>klusterlet-work-sa</code>                 |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-courier-addon-manager appscode/kubedb-courier-addon-manager -n kubedb --create-namespace --version=v0.7.0-rc.2 --set image.registry=ghcr.io/kubedb
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-courier-addon-manager appscode/kubedb-courier-addon-manager -n kubedb --create-namespace --version=v0.7.0-rc.2 --values values.yaml
```
