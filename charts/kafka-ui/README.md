# kafka-ui

[kafka-ui](https://docs.kafka-ui.provectus.io) - An administration tool for Kafka

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kafka-ui --version=v2024.3.9-rc.0
$ helm upgrade -i kafka-ui appscode/kafka-ui -n demo --create-namespace --version=v2024.3.9-rc.0
```

## Introduction

This chart deploys a kafka-ui deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.20+

## Installing the Chart

To install/upgrade the chart with the release name `kafka-ui`:

```bash
$ helm upgrade -i kafka-ui appscode/kafka-ui -n demo --create-namespace --version=v2024.3.9-rc.0
```

The command deploys a kafka-ui deployment on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kafka-ui`:

```bash
$ helm uninstall kafka-ui -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kafka-ui` chart and their default values.

|                 Parameter                  |                                                                        Description                                                                         |               Default               |
|--------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------|
| replicaCount                               |                                                                                                                                                            | <code>1</code>                      |
| image.registry                             |                                                                                                                                                            | <code>docker.io</code>              |
| image.repository                           |                                                                                                                                                            | <code>provectuslabs/kafka-ui</code> |
| image.pullPolicy                           |                                                                                                                                                            | <code>IfNotPresent</code>           |
| image.tag                                  | Overrides the image tag whose default is the chart appVersion.                                                                                             | <code>""</code>                     |
| imagePullSecrets                           |                                                                                                                                                            | <code>[]</code>                     |
| nameOverride                               |                                                                                                                                                            | <code>""</code>                     |
| fullnameOverride                           |                                                                                                                                                            | <code>""</code>                     |
| serviceAccount.create                      | Specifies whether a service account should be created                                                                                                      | <code>true</code>                   |
| serviceAccount.annotations                 | Annotations to add to the service account                                                                                                                  | <code>{}</code>                     |
| serviceAccount.name                        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                     | <code>""</code>                     |
| existingConfigMap                          |                                                                                                                                                            | <code>""</code>                     |
| yamlApplicationConfig                      |                                                                                                                                                            | <code>{}</code>                     |
| yamlApplicationConfigConfigMap             | kafka: clusters: - name: yaml bootstrapServers: kafka-service:9092 spring: security: oauth2: auth: type: disabled management: health: ldap: enabled: false | <code>{}</code>                     |
| existingSecret                             | keyName: config.yml name: configMapName                                                                                                                    | <code>""</code>                     |
| envs.secret                                |                                                                                                                                                            | <code>{}</code>                     |
| envs.config                                |                                                                                                                                                            | <code>{}</code>                     |
| networkPolicy.enabled                      |                                                                                                                                                            | <code>false</code>                  |
| networkPolicy.egressRules.customRules      | # Additional custom egress rules # e.g: # customRules: #   - to: #       - namespaceSelector: #           matchLabels: #             label: example        | <code>[]</code>                     |
| networkPolicy.ingressRules.customRules     | # Additional custom ingress rules # e.g: # customRules: #   - from: #       - namespaceSelector: #           matchLabels: #             label: example     | <code>[]</code>                     |
| podAnnotations                             |                                                                                                                                                            | <code>{}</code>                     |
| podLabels                                  |                                                                                                                                                            | <code>{}</code>                     |
| annotations                                | # Annotations to be added to kafka-ui Deployment #                                                                                                         | <code>{}</code>                     |
| probes.useHttpsScheme                      |                                                                                                                                                            | <code>false</code>                  |
| podSecurityContext                         |                                                                                                                                                            | <code>{}</code>                     |
| securityContext                            |                                                                                                                                                            | <code>{}</code>                     |
| service.type                               |                                                                                                                                                            | <code>ClusterIP</code>              |
| service.port                               |                                                                                                                                                            | <code>80</code>                     |
| ingress.enabled                            | Enable ingress resource                                                                                                                                    | <code>false</code>                  |
| ingress.annotations                        | Annotations for the Ingress                                                                                                                                | <code>{}</code>                     |
| ingress.ingressClassName                   | ingressClassName for the Ingress                                                                                                                           | <code>""</code>                     |
| ingress.path                               | The path for the Ingress                                                                                                                                   | <code>"/"</code>                    |
| ingress.pathType                           | The path type for the Ingress                                                                                                                              | <code>"Prefix"</code>               |
| ingress.host                               | The hostname for the Ingress                                                                                                                               | <code>""</code>                     |
| ingress.tls.enabled                        | Enable TLS termination for the Ingress                                                                                                                     | <code>false</code>                  |
| ingress.tls.secretName                     | the name of a pre-created Secret containing a TLS private key and certificate                                                                              | <code>""</code>                     |
| ingress.precedingPaths                     | HTTP paths to add to the Ingress before the default path                                                                                                   | <code>[]</code>                     |
| ingress.succeedingPaths                    | Http paths to add to the Ingress after the default path                                                                                                    | <code>[]</code>                     |
| resources                                  |                                                                                                                                                            | <code>{}</code>                     |
| autoscaling.enabled                        |                                                                                                                                                            | <code>false</code>                  |
| autoscaling.minReplicas                    |                                                                                                                                                            | <code>1</code>                      |
| autoscaling.maxReplicas                    |                                                                                                                                                            | <code>100</code>                    |
| autoscaling.targetCPUUtilizationPercentage |                                                                                                                                                            | <code>80</code>                     |
| nodeSelector                               |                                                                                                                                                            | <code>{}</code>                     |
| tolerations                                |                                                                                                                                                            | <code>[]</code>                     |
| affinity                                   |                                                                                                                                                            | <code>{}</code>                     |
| env                                        |                                                                                                                                                            | <code>{}</code>                     |
| initContainers                             |                                                                                                                                                            | <code>{}</code>                     |
| volumeMounts                               |                                                                                                                                                            | <code>{}</code>                     |
| volumes                                    |                                                                                                                                                            | <code>{}</code>                     |
| namespace.create                           |                                                                                                                                                            | <code>false</code>                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kafka-ui appscode/kafka-ui -n demo --create-namespace --version=v2024.3.9-rc.0 --set image.tag=latest
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kafka-ui appscode/kafka-ui -n demo --create-namespace --version=v2024.3.9-rc.0 --values values.yaml
```
