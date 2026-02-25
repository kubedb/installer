# Database Migrator

[Database Migrator by AppsCodes](https://github.com/kubedb) - Community features for KubeDB by AppsCode

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-migrator --version=v0.1.0-rc.1
$ helm upgrade -i kubedb-migrator appscode/kubedb-migrator -n kubedb --create-namespace --version=v0.1.0-rc.1
```

## Introduction

This chart deploys a Database Migrator operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-migrator`:

```bash
$ helm upgrade -i kubedb-migrator appscode/kubedb-migrator -n kubedb --create-namespace --version=v0.1.0-rc.1
```

The command deploys a Database Migrator operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-migrator`:

```bash
$ helm uninstall kubedb-migrator -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-migrator` chart and their default values.

|             Parameter              |                                                             Description                                                              |                                                                                            Default                                                                                             |
|------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| replicaCount                       |                                                                                                                                      | <code>1</code>                                                                                                                                                                                 |
| registryFQDN                       | Docker registry fqdn used to pull app related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image} | <code>ghcr.io</code>                                                                                                                                                                           |
| image.registry                     | Docker registry used to pull app container image                                                                                     | <code>kubedb</code>                                                                                                                                                                            |
| image.repository                   |                                                                                                                                      | <code>migrator-operator</code>                                                                                                                                                                 |
| image.pullPolicy                   |                                                                                                                                      | <code>IfNotPresent</code>                                                                                                                                                                      |
| image.tag                          | Overrides the image tag whose default is the chart appVersion.                                                                       | <code>""</code>                                                                                                                                                                                |
| imagePullSecrets                   |                                                                                                                                      | <code>[]</code>                                                                                                                                                                                |
| nameOverride                       |                                                                                                                                      | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                   |                                                                                                                                      | <code>""</code>                                                                                                                                                                                |
| serviceAccount.create              | Specifies whether a service account should be created                                                                                | <code>true</code>                                                                                                                                                                              |
| serviceAccount.annotations         | Annotations to add to the service account                                                                                            | <code>{}</code>                                                                                                                                                                                |
| serviceAccount.name                | The name of the service account to use. If not set and create is true, a name is generated using the fullname template               | <code>""</code>                                                                                                                                                                                |
| podAnnotations                     |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| podLabels                          |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| podSecurityContext                 |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| securityContext                    |                                                                                                                                      | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| service.type                       |                                                                                                                                      | <code>ClusterIP</code>                                                                                                                                                                         |
| service.port                       |                                                                                                                                      | <code>8081</code>                                                                                                                                                                              |
| resources                          |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| livenessProbe.httpGet.path         |                                                                                                                                      | <code>/healthz</code>                                                                                                                                                                          |
| livenessProbe.httpGet.port         |                                                                                                                                      | <code>8081</code>                                                                                                                                                                              |
| livenessProbe.httpGet.scheme       |                                                                                                                                      | <code>HTTP</code>                                                                                                                                                                              |
| livenessProbe.initialDelaySeconds  |                                                                                                                                      | <code>15</code>                                                                                                                                                                                |
| livenessProbe.periodSeconds        |                                                                                                                                      | <code>20</code>                                                                                                                                                                                |
| readinessProbe.httpGet.path        |                                                                                                                                      | <code>/readyz</code>                                                                                                                                                                           |
| readinessProbe.httpGet.port        |                                                                                                                                      | <code>8081</code>                                                                                                                                                                              |
| readinessProbe.httpGet.scheme      |                                                                                                                                      | <code>HTTP</code>                                                                                                                                                                              |
| readinessProbe.initialDelaySeconds |                                                                                                                                      | <code>5</code>                                                                                                                                                                                 |
| readinessProbe.periodSeconds       |                                                                                                                                      | <code>10</code>                                                                                                                                                                                |
| volumes                            |                                                                                                                                      | <code>[]</code>                                                                                                                                                                                |
| volumeMounts                       |                                                                                                                                      | <code>[]</code>                                                                                                                                                                                |
| nodeSelector                       |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| tolerations                        |                                                                                                                                      | <code>[]</code>                                                                                                                                                                                |
| affinity                           |                                                                                                                                      | <code>{}</code>                                                                                                                                                                                |
| distro.openshift                   | Set true, if installed in OpenShift                                                                                                  | <code>false</code>                                                                                                                                                                             |
| distro.ubi                         | Set operator or all to use ubi images                                                                                                | <code>""</code>                                                                                                                                                                                |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-migrator appscode/kubedb-migrator -n kubedb --create-namespace --version=v0.1.0-rc.1 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-migrator appscode/kubedb-migrator -n kubedb --create-namespace --version=v0.1.0-rc.1 --values values.yaml
```
