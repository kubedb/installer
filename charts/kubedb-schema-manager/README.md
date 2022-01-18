# KubeDB Schema Manager

[KubeDB Schema Manager by AppsCode](https://github.com/kubedb) - Database Schema Manager for KubeDB

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubedb-schema-manager appscode/kubedb-schema-manager -n kubedb
```

## Introduction

This chart deploys a KubeDB schema manager operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubedb-schema-manager`:

```console
$ helm install kubedb-schema-manager appscode/kubedb-schema-manager -n kubedb
```

The command deploys a KubeDB schema manager operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubedb-schema-manager`:

```console
$ helm delete kubedb-schema-manager -n kubedb
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubedb-schema-manager` chart and their default values.

|           Parameter            |                                                                                                                                                                                             Description                                                                                                                                                                                              |              Default               |
|--------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| nameOverride                   | Overrides name template                                                                                                                                                                                                                                                                                                                                                                              | <code>""</code>                    |
| fullnameOverride               | Overrides fullname template                                                                                                                                                                                                                                                                                                                                                                          | <code>""</code>                    |
| replicaCount                   | Number of Kubeform operator replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                                                 | <code>1</code>                     |
| license                        | License for the product. Get a license by following the steps from [here](https://kubedb.com/docs/latest/setup/install/overview/#get-a-license). <br> Example: <br> `helm install appscode/kubedb-schema-manager-enterprise \` <br> `--set-file license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubedb-schema-manager-enterprise \` <br> `--set license=<license file content>` | <code>""</code>                    |
| registryFQDN                   | Docker registry fqdn used to pull Kubeform related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                                            | <code>""</code>                    |
| operator.registry              | Docker registry used to pull operator image                                                                                                                                                                                                                                                                                                                                                          | <code>kubedb</code>                |
| operator.repository            | Operator container docker image                                                                                                                                                                                                                                                                                                                                                                      | <code>kubedb-schema-manager</code> |
| operator.tag                   | Operator container image tag                                                                                                                                                                                                                                                                                                                                                                         | <code>v0.0.1</code>                |
| operator.resources             | Compute Resources required by the operator container                                                                                                                                                                                                                                                                                                                                                 | <code>{}</code>                    |
| operator.securityContext       | Security options the operator container should run with                                                                                                                                                                                                                                                                                                                                              | <code>{}</code>                    |
| cleaner.registry               | Docker registry used to pull Webhook cleaner image                                                                                                                                                                                                                                                                                                                                                   | <code>appscode</code>              |
| cleaner.repository             | Webhook cleaner container image                                                                                                                                                                                                                                                                                                                                                                      | <code>kubectl</code>               |
| cleaner.tag                    | Webhook cleaner container image tag                                                                                                                                                                                                                                                                                                                                                                  | <code>v1.22</code>                 |
| cleaner.skip                   | Skip generating cleaner YAML                                                                                                                                                                                                                                                                                                                                                                         | <code>false</code>                 |
| imagePullSecrets               | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubedb-schema-manager \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                                                   | <code>[]</code>                    |
| imagePullPolicy                | Container image pull policy                                                                                                                                                                                                                                                                                                                                                                          | <code>IfNotPresent</code>          |
| criticalAddon                  | If true, installs kubedb-schema-manager operator as critical addon                                                                                                                                                                                                                                                                                                                                   | <code>false</code>                 |
| logLevel                       | Log level for operator                                                                                                                                                                                                                                                                                                                                                                               | <code>3</code>                     |
| annotations                    | Annotations applied to operator deployment                                                                                                                                                                                                                                                                                                                                                           | <code>{}</code>                    |
| podAnnotations                 | Annotations passed to operator pod(s).                                                                                                                                                                                                                                                                                                                                                               | <code>{}</code>                    |
| nodeSelector                   | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                                                       | <code></code>                      |
| tolerations                    | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                                                       | <code>[]</code>                    |
| affinity                       | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                                                                    | <code>{}</code>                    |
| podSecurityContext             | Security options the operator pod should run with.                                                                                                                                                                                                                                                                                                                                                   | <code>{}</code>                    |
| serviceAccount.create          | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                                                                | <code>true</code>                  |
| serviceAccount.annotations     | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                                                            | <code>{}</code>                    |
| serviceAccount.name            | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                                                               | <code></code>                      |
| secretKey                      | Specifies a base64-encoded key, of length 32 bytes when decoded. It is used to encrypt the state file.                                                                                                                                                                                                                                                                                               | <code></code>                      |
| enableAnalytics                | If true, sends usage analytics                                                                                                                                                                                                                                                                                                                                                                       | <code>true</code>                  |
| proxy.https                    | To configure HTTPS_PROXY environment variable specify <ip_address>:<port>                                                                                                                                                                                                                                                                                                                            | <code>''</code>                    |
| proxy.http                     | To configure HTTP_PROXY environment variable specify <ip_address>:<port>                                                                                                                                                                                                                                                                                                                             | <code>''</code>                    |
| proxy.no                       | To configure NO_PROXY environment variable specify <ip_address>:<port> By default exclude Kubernetes apiserver internal IP.                                                                                                                                                                                                                                                                          | <code>'10.43.0.1'</code>           |
| webhook.servingCerts.generate  | If true, generates on install/upgrade the certs that is required for the webhook server.                                                                                                                                                                                                                                                                                                             | <code>true</code>                  |
| webhook.servingCerts.caCrt     | CA certificate used by serving certificate of webhook server.                                                                                                                                                                                                                                                                                                                                        | <code>""</code>                    |
| webhook.servingCerts.serverCrt | Serving certificate used by webhook server.                                                                                                                                                                                                                                                                                                                                                          | <code>""</code>                    |
| webhook.servingCerts.serverKey | Private key for the serving certificate used by webhook server.                                                                                                                                                                                                                                                                                                                                      | <code>""</code>                    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubedb-schema-manager appscode/kubedb-schema-manager -n kubedb --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubedb-schema-manager appscode/kubedb-schema-manager -n kubedb --values values.yaml
```
