# KubeDB Provisioner

[KubeDB Provisioner by AppsCode](https://github.com/kubedb) - Community features for KubeDB by AppsCode

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubedb-webhook-server --version=v0.5.0-rc.0
$ helm upgrade -i kubedb-webhook-server appscode/kubedb-webhook-server -n kubedb --create-namespace --version=v0.5.0-rc.0
```

## Introduction

This chart deploys a KubeDB Provisioner operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubedb-webhook-server`:

```bash
$ helm upgrade -i kubedb-webhook-server appscode/kubedb-webhook-server -n kubedb --create-namespace --version=v0.5.0-rc.0
```

The command deploys a KubeDB Provisioner operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubedb-webhook-server`:

```bash
$ helm uninstall kubedb-webhook-server -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-webhook-server` chart and their default values.

|              Parameter               |                                                                                                                                                                                   Description                                                                                                                                                                                   |                  Default                  |
|--------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------|
| nameOverride                         | Overrides name template                                                                                                                                                                                                                                                                                                                                                         | <code>""</code>                           |
| fullnameOverride                     | Overrides fullname template                                                                                                                                                                                                                                                                                                                                                     | <code>""</code>                           |
| replicaCount                         | Number of KubeDB webhook server replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                        | <code>1</code>                            |
| license                              | License for the product. Get a license by following the steps from [here](https://kubedb.run/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubedb-ops-manager \` <br> `--set-file license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubedb-ops-manager \` <br> `--set license=<license file content>` | <code>""</code>                           |
| registryFQDN                         | Docker registry fqdn used to pull KubeDB related images Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                          | <code>""</code>                           |
| server.registry                      | Docker registry used to pull KubeDB webhook server image                                                                                                                                                                                                                                                                                                                        | <code>kubedb</code>                       |
| server.repository                    | KubeDB webhook server container image                                                                                                                                                                                                                                                                                                                                           | <code>kubedb-webhook-server</code>        |
| server.tag                           | KubeDB webhook server container image tag                                                                                                                                                                                                                                                                                                                                       | <code>""</code>                           |
| server.resources                     | Compute Resources required by the webhook server container                                                                                                                                                                                                                                                                                                                      | <code>{}</code>                           |
| server.securityContext               | requests: cpu: 100m memory: 128Mi Security options the webhook server container should run with                                                                                                                                                                                                                                                                                 | <code>{}</code>                           |
| imagePullSecrets                     | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb-webhook-server \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                              | <code>[]</code>                           |
| imagePullPolicy                      | Container image pull policy                                                                                                                                                                                                                                                                                                                                                     | <code>IfNotPresent</code>                 |
| criticalAddon                        | If true, installs KubeDB webhook server as critical addon                                                                                                                                                                                                                                                                                                                       | <code>false</code>                        |
| logLevel                             | Log level for webhook server                                                                                                                                                                                                                                                                                                                                                    | <code>3</code>                            |
| annotations                          | Annotations applied to webhook server deployment                                                                                                                                                                                                                                                                                                                                | <code>{}</code>                           |
| podAnnotations                       | Annotations passed to webhook server pod(s).                                                                                                                                                                                                                                                                                                                                    | <code>{}</code>                           |
| nodeSelector                         | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                                  | <code>{"kubernetes.io/os":"linux"}</code> |
| tolerations                          | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                                  | <code>[]</code>                           |
| affinity                             | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                                               | <code>{}</code>                           |
| podSecurityContext                   | Security options the webhook server pod should run with.                                                                                                                                                                                                                                                                                                                        | <code>{}</code>                           |
| serviceAccount.create                | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                                           | <code>true</code>                         |
| serviceAccount.annotations           | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                                       | <code>{}</code>                           |
| serviceAccount.name                  | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                                          | <code></code>                             |
| apiserver.groupPriorityMinimum       | The minimum priority the webhook api group should have at least. Please see https://github.com/kubernetes/kube-aggregator/blob/release-1.9/pkg/apis/apiregistration/v1beta1/types.go#L58-L64 for more information on proper values of this field.                                                                                                                               | <code>10000</code>                        |
| apiserver.versionPriority            | The ordering of the webhook api inside of the group. Please see https://github.com/kubernetes/kube-aggregator/blob/release-1.9/pkg/apis/apiregistration/v1beta1/types.go#L66-L70 for more information on proper values of this field                                                                                                                                            | <code>15</code>                           |
| apiserver.enableMutatingWebhook      | If true, mutating webhook is configured for KubeDB CRDss                                                                                                                                                                                                                                                                                                                        | <code>true</code>                         |
| apiserver.enableValidatingWebhook    | If true, validating webhook is configured for KubeDB CRDss                                                                                                                                                                                                                                                                                                                      | <code>true</code>                         |
| apiserver.ca                         | CA certificate used by the Kubernetes api server. This field is automatically assigned by the webhook server.                                                                                                                                                                                                                                                                   | <code>not-ca-cert</code>                  |
| apiserver.useKubeapiserverFqdnForAks | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)                                                                                                                                                                                                                                                          | <code>true</code>                         |
| apiserver.healthcheck.enabled        | healthcheck configures the readiness and liveliness probes for the webhook server pod.                                                                                                                                                                                                                                                                                          | <code>false</code>                        |
| apiserver.port                       | Port used to expose the webhook server apiserver                                                                                                                                                                                                                                                                                                                                | <code>8443</code>                         |
| apiserver.servingCerts.generate      | If true, generates on install/upgrade the certs that allow the kube-apiserver (and potentially ServiceMonitor) to authenticate webhook servers pods. Otherwise specify certs in `apiserver.servingCerts.{caCrt, serverCrt, serverKey}`.                                                                                                                                         | <code>true</code>                         |
| apiserver.servingCerts.caCrt         | CA certficate used by serving certificate of webhook server.                                                                                                                                                                                                                                                                                                                    | <code>""</code>                           |
| apiserver.servingCerts.serverCrt     | Serving certficate used by webhook server.                                                                                                                                                                                                                                                                                                                                      | <code>""</code>                           |
| apiserver.servingCerts.serverKey     | Private key for the serving certificate used by webhook server.                                                                                                                                                                                                                                                                                                                 | <code>""</code>                           |
| monitoring.agent                     | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                                                                                                                                            | <code>""</code>                           |
| monitoring.serviceMonitor.labels     | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/webhook server`.                                                                                                                                                                                                       | <code>{}</code>                           |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubedb-webhook-server appscode/kubedb-webhook-server -n kubedb --create-namespace --version=v0.5.0-rc.0 --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubedb-webhook-server appscode/kubedb-webhook-server -n kubedb --create-namespace --version=v0.5.0-rc.0 --values values.yaml
```
