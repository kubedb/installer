# phpMyAdmin

[phpMyAdmin](https://www.phpmyadmin.net) - An administration to for MySQL and MariaDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/phpmyadmin --version=v2024.3.16
$ helm upgrade -i phpmyadmin appscode/phpmyadmin -n demo --create-namespace --version=v2024.3.16
```

## Introduction

This chart deploys a phpMyAdmin deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `phpmyadmin`:

```bash
$ helm upgrade -i phpmyadmin appscode/phpmyadmin -n demo --create-namespace --version=v2024.3.16
```

The command deploys a phpMyAdmin deployment on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `phpmyadmin`:

```bash
$ helm uninstall phpmyadmin -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `phpmyadmin` chart and their default values.

|           Parameter            |                                                      Description                                                       |                      Default                       |
|--------------------------------|------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| replicaCount                   |                                                                                                                        | <code>1</code>                                     |
| image.repository               |                                                                                                                        | <code>"phpmyadmin"</code>                          |
| image.pullPolicy               |                                                                                                                        | <code>Always</code>                                |
| image.tag                      | Overrides the image tag whose default is the chart appVersion.                                                         | <code>"latest"</code>                              |
| imagePullSecrets               |                                                                                                                        | <code>[]</code>                                    |
| nameOverride                   |                                                                                                                        | <code>""</code>                                    |
| fullnameOverride               |                                                                                                                        | <code>""</code>                                    |
| serviceAccount.create          | Specifies whether a service account should be created                                                                  | <code>true</code>                                  |
| serviceAccount.annotations     | Annotations to add to the service account                                                                              | <code>{}</code>                                    |
| serviceAccount.name            | The name of the service account to use. If not set and create is true, a name is generated using the fullname template | <code>""</code>                                    |
| podAnnotations                 |                                                                                                                        | <code>{}</code>                                    |
| podSecurityContext             |                                                                                                                        | <code>{}</code>                                    |
| securityContext                |                                                                                                                        | <code>{}</code>                                    |
| service.type                   |                                                                                                                        | <code>ClusterIP</code>                             |
| service.port                   |                                                                                                                        | <code>80</code>                                    |
| resources                      |                                                                                                                        | <code>{}</code>                                    |
| nodeSelector                   |                                                                                                                        | <code>{}</code>                                    |
| tolerations                    |                                                                                                                        | <code>[]</code>                                    |
| affinity                       |                                                                                                                        | <code>{}</code>                                    |
| namespace.create               |                                                                                                                        | <code>false</code>                                 |
| gateway.className              |                                                                                                                        | <code>"ace"</code>                                 |
| gateway.port                   |                                                                                                                        | <code>8082</code>                                  |
| gateway.tlsSecretRef.name      |                                                                                                                        | <code>service-presets-cert</code>                  |
| gateway.tlsSecretRef.namespace |                                                                                                                        | <code>ace</code>                                   |
| gateway.referenceGrant.create  |                                                                                                                        | <code>true</code>                                  |
| keda.proxyService.namespace    |                                                                                                                        | <code>"keda"</code>                                |
| keda.proxyService.name         |                                                                                                                        | <code>"keda-add-ons-http-interceptor-proxy"</code> |
| keda.proxyService.port         |                                                                                                                        | <code>8080</code>                                  |
| targetPendingRequests          |                                                                                                                        | <code>200</code>                                   |
| autoscaling.http.minReplicas   |                                                                                                                        | <code>0</code>                                     |
| autoscaling.http.maxReplicas   |                                                                                                                        | <code>1</code>                                     |
| app.service.name               |                                                                                                                        | <code>""</code>                                    |
| app.service.namespace          |                                                                                                                        | <code>""</code>                                    |
| app.authSecret.name            |                                                                                                                        | <code>""</code>                                    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i phpmyadmin appscode/phpmyadmin -n demo --create-namespace --version=v2024.3.16 --set image.tag=latest
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i phpmyadmin appscode/phpmyadmin -n demo --create-namespace --version=v2024.3.16 --values values.yaml
```
