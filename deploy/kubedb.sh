#!/bin/bash
set -eou pipefail

crds=(
  dormantdatabases.kubedb.com
  elasticsearches.kubedb.com
  etcds.kubedb.com
  memcacheds.kubedb.com
  mongodbs.kubedb.com
  mysqls.kubedb.com
  postgreses.kubedb.com
  redises.kubedb.com
  snapshots.kubedb.com
  elasticsearchversions.catalog.kubedb.com
  etcdversions.catalog.kubedb.com
  memcachedversions.catalog.kubedb.com
  mongodbversions.catalog.kubedb.com
  mysqlversions.catalog.kubedb.com
  postgresversions.catalog.kubedb.com
  redisversions.catalog.kubedb.com
  appbindings.appcatalog.appscode.com
)
apiServices=(v1alpha1.validators v1alpha1.mutators)

echo "checking kubeconfig context"
kubectl config current-context || {
  echo "Set a context (kubectl use-context <context>) out of the following:"
  echo
  kubectl config get-contexts
  exit 1
}
echo ""

# http://redsymbol.net/articles/bash-exit-traps/
function cleanup() {
  rm -rf $ONESSL ca.crt ca.key server.crt server.key
}
trap cleanup EXIT

onessl_found() {
  # https://stackoverflow.com/a/677212/244009
  if [ -x "$(command -v onessl)" ]; then
    onessl wait-until-has -h >/dev/null 2>&1 || {
      # old version of onessl found
      echo "Found outdated onessl"
      return 1
    }
    export ONESSL=onessl
    return 0
  fi
  return 1
}

onessl_found || {
  echo "Downloading onessl ..."
  if [[ "$(uname -m)" == "aarch64" ]]; then
    curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.10.0/onessl-linux-arm64
    chmod +x onessl
    export ONESSL=./onessl
  else
    # ref: https://stackoverflow.com/a/27776822/244009
    case "$(uname -s)" in
      Darwin)
        curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.10.0/onessl-darwin-amd64
        chmod +x onessl
        export ONESSL=./onessl
        ;;

      Linux)
        curl -fsSL -o onessl https://github.com/kubepack/onessl/releases/download/0.10.0/onessl-linux-amd64
        chmod +x onessl
        export ONESSL=./onessl
        ;;

      CYGWIN* | MINGW* | MSYS*)
        curl -fsSL -o onessl.exe https://github.com/kubepack/onessl/releases/download/0.10.0/onessl-windows-amd64.exe
        chmod +x onessl.exe
        export ONESSL=./onessl.exe
        ;;
      *)
        echo 'other OS'
        ;;
    esac
  fi
}

# ref: https://stackoverflow.com/a/7069755/244009
# ref: https://jonalmeida.com/posts/2013/05/26/different-ways-to-implement-flags-in-bash/
# ref: http://tldp.org/LDP/abs/html/comparison-ops.html

export KUBEDB_NAMESPACE=kube-system
export KUBEDB_SERVICE_ACCOUNT=kubedb-operator
export KUBEDB_ENABLE_RBAC=true
export KUBEDB_RUN_ON_MASTER=0
export KUBEDB_ENABLE_VALIDATING_WEBHOOK=false
export KUBEDB_ENABLE_MUTATING_WEBHOOK=false
export KUBEDB_CATALOG=${KUBEDB_CATALOG:-all}
export KUBEDB_DOCKER_REGISTRY=${KUBEDB_DOCKER_REGISTRY:-kubedb}
export KUBEDB_OPERATOR_TAG=${KUBEDB_OPERATOR_TAG:-v0.13.0-rc.0}
export KUBEDB_OPERATOR_NAME=operator
export KUBEDB_IMAGE_PULL_SECRET=
export KUBEDB_IMAGE_PULL_POLICY=IfNotPresent
export KUBEDB_ENABLE_ANALYTICS=true
export KUBEDB_UNINSTALL=0
export KUBEDB_PURGE=0
export KUBEDB_ENABLE_STATUS_SUBRESOURCE=false
export KUBEDB_BYPASS_VALIDATING_WEBHOOK_XRAY=false
export KUBEDB_USE_KUBEAPISERVER_FQDN_FOR_AKS=true
export KUBEDB_PRIORITY_CLASS=system-cluster-critical

export APPSCODE_ENV=${APPSCODE_ENV:-prod}
export SCRIPT_LOCATION="curl -fsSL https://raw.githubusercontent.com/kubedb/installer/v0.13.0-rc.0/"
if [ "$APPSCODE_ENV" = "dev" ]; then
  export SCRIPT_LOCATION="cat "
  export KUBEDB_IMAGE_PULL_POLICY=Always
fi

KUBE_APISERVER_VERSION=$(kubectl version -o=json | $ONESSL jsonpath '{.serverVersion.gitVersion}')
$ONESSL semver --check='<1.9.0' $KUBE_APISERVER_VERSION || {
  export KUBEDB_ENABLE_VALIDATING_WEBHOOK=true
  export KUBEDB_ENABLE_MUTATING_WEBHOOK=true
}
$ONESSL semver --check='<1.11.0' $KUBE_APISERVER_VERSION || { export KUBEDB_ENABLE_STATUS_SUBRESOURCE=true; }

export KUBEDB_WEBHOOK_SIDE_EFFECTS=
$ONESSL semver --check='<1.12.0' $KUBE_APISERVER_VERSION || { export KUBEDB_WEBHOOK_SIDE_EFFECTS='sideEffects: None'; }

MONITORING_AGENT_NONE="none"
MONITORING_AGENT_BUILTIN="prometheus.io/builtin"
MONITORING_AGENT_COREOS_OPERATOR="prometheus.io/coreos-operator"

export MONITORING_ENABLE=${MONITORING_ENABLE:-false}
export MONITORING_AGENT=${MONITORING_AGENT:-$MONITORING_AGENT_NONE}
export SERVICE_MONITOR_LABEL_KEY="app"
export SERVICE_MONITOR_LABEL_VALUE="kubedb"

show_help() {
  echo "kubedb.sh - install kubedb operator"
  echo " "
  echo "kubedb.sh [options]"
  echo " "
  echo "options:"
  echo "-h, --help                             show brief help"
  echo "-n, --namespace=NAMESPACE              specify namespace (default: kube-system)"
  echo "    --rbac                             create RBAC roles and bindings (default: true)"
  echo "    --docker-registry                  docker registry used to pull KubeDB images (default: appscode)"
  echo "    --image-pull-secret                name of secret used to pull KubeDB operator images"
  echo "    --run-on-master                    run KubeDB operator on master"
  echo "    --enable-validating-webhook        enable/disable validating webhooks for KubeDB CRDs"
  echo "    --enable-mutating-webhook          enable/disable mutating webhooks for KubeDB CRDs"
  echo "    --bypass-validating-webhook-xray   if true, bypasses validating webhook xray checks"
  echo "    --enable-status-subresource        if enabled, uses status sub resource for KubeDB crds"
  echo "    --use-kubeapiserver-fqdn-for-aks   if true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)"
  echo "    --enable-analytics                 send usage events to Google Analytics (default: true)"
  echo "    --install-catalog                  installs KubeDB database version catalog (default: all)"
  echo "    --operator-name                    specify which KubeDB operator to deploy (default: operator)"
  echo "    --uninstall                        uninstall KubeDB"
  echo "    --purge                            purges KubeDB crd objects and crds"
  echo "    --monitoring-enable                specify whether to monitor KubeDB operator (default: false)"
  echo "    --monitoring-agent                 specify which monitoring agent to use (default: none)"
  echo "    --prometheus-namespace             specify the namespace where Prometheus server is running or will be deployed (default: same namespace as kubedb-operator)"
  echo "    --servicemonitor-label             specify the label for ServiceMonitor crd. Prometheus crd will use this label to select the ServiceMonitor. (default: 'app: kubedb')"
}

while test $# -gt 0; do
  case "$1" in
    -h | --help)
      show_help
      exit 0
      ;;
    -n)
      shift
      if test $# -gt 0; then
        export KUBEDB_NAMESPACE=$1
      else
        echo "no namespace specified"
        exit 1
      fi
      shift
      ;;
    --namespace*)
      export KUBEDB_NAMESPACE=$(echo $1 | sed -e 's/^[^=]*=//g')
      shift
      ;;
    --docker-registry*)
      export KUBEDB_DOCKER_REGISTRY=$(echo $1 | sed -e 's/^[^=]*=//g')
      shift
      ;;
    --image-pull-secret*)
      secret=$(echo $1 | sed -e 's/^[^=]*=//g')
      export KUBEDB_IMAGE_PULL_SECRET="name: '$secret'"
      shift
      ;;
    --enable-validating-webhook*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_ENABLE_VALIDATING_WEBHOOK=false
      fi
      shift
      ;;
    --enable-mutating-webhook*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_ENABLE_MUTATING_WEBHOOK=false
      fi
      shift
      ;;
    --bypass-validating-webhook-xray*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_BYPASS_VALIDATING_WEBHOOK_XRAY=false
      else
        export KUBEDB_BYPASS_VALIDATING_WEBHOOK_XRAY=true
      fi
      shift
      ;;
    --enable-status-subresource*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_ENABLE_STATUS_SUBRESOURCE=false
      fi
      shift
      ;;
    --use-kubeapiserver-fqdn-for-aks*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_USE_KUBEAPISERVER_FQDN_FOR_AKS=false
      else
        export KUBEDB_USE_KUBEAPISERVER_FQDN_FOR_AKS=true
      fi
      shift
      ;;
    --enable-analytics*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_ENABLE_ANALYTICS=false
      fi
      shift
      ;;
    --install-catalog*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_CATALOG=false
      fi
      shift
      ;;
    --rbac*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "false" ]; then
        export KUBEDB_SERVICE_ACCOUNT=default
        export KUBEDB_ENABLE_RBAC=false
      fi
      shift
      ;;
    --run-on-master)
      export KUBEDB_RUN_ON_MASTER=1
      shift
      ;;
    --operator-name*)
      export KUBEDB_OPERATOR_NAME=$(echo $1 | sed -e 's/^[^=]*=//g')
      shift
      ;;
    --uninstall)
      export KUBEDB_UNINSTALL=1
      shift
      ;;
    --purge)
      export KUBEDB_PURGE=1
      shift
      ;;
    --monitoring-enable*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" = "true" ]; then
        export MONITORING_ENABLE="$val"
      fi
      shift
      ;;
    --monitoring-agent*)
      val=$(echo $1 | sed -e 's/^[^=]*=//g')
      if [ "$val" != "$MONITORING_AGENT_BUILTIN" ] && [ "$val" != "$MONITORING_AGENT_COREOS_OPERATOR" ]; then
        echo 'Invalid monitoring agent. Use "builtin" or "coreos-operator"'
        exit 1
      else
        export MONITORING_AGENT="$val"
      fi
      shift
      ;;
    --prometheus-namespace*)
      export PROMETHEUS_NAMESPACE=$(echo $1 | sed -e 's/^[^=]*=//g')
      shift
      ;;
    --servicemonitor-label*)
      label=$(echo $1 | sed -e 's/^[^=]*=//g')
      # split label into key value pair
      IFS='='
      pair=($label)
      unset IFS
      # check if the label is valid
      if [ ! ${#pair[@]} = 2 ]; then
        echo "Invalid ServiceMonitor label format. Use '--servicemonitor-label=key=value'"
        exit 1
      fi
      export SERVICE_MONITOR_LABEL_KEY="${pair[0]}"
      export SERVICE_MONITOR_LABEL_VALUE="${pair[1]}"
      shift
      ;;
    *)
      echo "Error: unknown flag:" $1
      show_help
      exit 1
      ;;
  esac
done

export PROMETHEUS_NAMESPACE=${PROMETHEUS_NAMESPACE:-$KUBEDB_NAMESPACE}

if [ "$KUBEDB_NAMESPACE" != "kube-system" ]; then
    export KUBEDB_PRIORITY_CLASS=""
fi

if [ "$KUBEDB_UNINSTALL" -eq 1 ]; then
  # delete webhooks and apiservices
  kubectl delete validatingwebhookconfiguration -l app=kubedb || true
  kubectl delete mutatingwebhookconfiguration -l app=kubedb || true
  kubectl delete apiservice -l app=kubedb
  # delete kubedb operator
  kubectl delete deployment -l app=kubedb --namespace $KUBEDB_NAMESPACE
  kubectl delete service -l app=kubedb --namespace $KUBEDB_NAMESPACE
  kubectl delete secret -l app=kubedb --namespace $KUBEDB_NAMESPACE
  # delete RBAC objects, if --rbac flag was used.
  kubectl delete serviceaccount -l app=kubedb --namespace $KUBEDB_NAMESPACE
  kubectl delete clusterrolebindings -l app=kubedb
  kubectl delete clusterrole -l app=kubedb
  kubectl delete rolebindings -l app=kubedb --namespace $KUBEDB_NAMESPACE
  kubectl delete role -l app=kubedb --namespace $KUBEDB_NAMESPACE
  kubectl delete psp -l app=kubedb

  # delete servicemonitor and kubedb-operator-apiserver-cert secret. ignore error as they might not exist
  kubectl delete servicemonitor kubedb-${KUBEDB_OPERATOR_NAME}-servicemonitor --namespace $PROMETHEUS_NAMESPACE || true
  kubectl delete secret kubedb-${KUBEDB_OPERATOR_NAME}-apiserver-cert --namespace $PROMETHEUS_NAMESPACE || true

  echo "waiting for kubedb operator pod to stop running"
  for (( ; ; )); do
    pods=($(kubectl get pods --namespace $KUBEDB_NAMESPACE -l app=kubedb -o jsonpath='{range .items[*]}{.metadata.name} {end}'))
    total=${#pods[*]}
    if [ $total -eq 0 ]; then
      break
    fi
    sleep 2
  done

  # https://github.com/kubernetes/kubernetes/issues/60538
  if [ "$KUBEDB_PURGE" -eq 1 ]; then
    for crd in "${crds[@]}"; do
      pairs=($(kubectl get ${crd} --all-namespaces -o jsonpath='{range .items[*]}{.metadata.name} {.metadata.namespace} {end}' || true))
      total=${#pairs[*]}

      # save objects
      if [ $total -gt 0 ]; then
        echo "dumping ${crd} objects into ${crd}.yaml"
        kubectl get ${crd} --all-namespaces -o yaml >${crd}.yaml
      fi

      for ((i = 0; i < $total; i++)); do
        name=${pairs[$i]}
        namespace="default"
        if [[ $crd != *"catalog.kubedb.com" ]]; then
          namespace=${pairs[$i + 1]}
          i=$((i + 1))
        fi
        # remove finalizers
        kubectl patch ${crd} $name -n $namespace -p '{"metadata":{"finalizers":[]}}' --type=merge || true
        # delete crd object
        echo "deleting ${crd} $namespace/$name"
        kubectl delete ${crd} $name -n $namespace --ignore-not-found=true
      done

      # delete crd
      kubectl delete crd ${crd} --ignore-not-found=true
    done

    # delete user roles
    kubectl delete clusterroles kubedb:core:admin kubedb:core:edit kubedb:core:view --ignore-not-found=true
  fi

  echo
  echo "Successfully uninstalled KubeDB!"
  exit 0
fi

echo "checking whether extended apiserver feature is enabled"
$ONESSL has-keys configmap --namespace=kube-system --keys=requestheader-client-ca-file extension-apiserver-authentication || {
  echo "Set --requestheader-client-ca-file flag on Kubernetes apiserver"
  exit 1
}
echo ""

export KUBE_CA=
export KUBEDB_ENABLE_APISERVER=false
if [ "$KUBEDB_ENABLE_VALIDATING_WEBHOOK" = true ] || [ "$KUBEDB_ENABLE_MUTATING_WEBHOOK" = true ]; then
  $ONESSL get kube-ca >/dev/null 2>&1 || {
    echo "Admission webhooks can't be used when kube apiserver is accesible without verifying its TLS certificate (insecure-skip-tls-verify : true)."
    echo
    exit 1
  }
  export KUBE_CA=$($ONESSL get kube-ca | $ONESSL base64)
  export KUBEDB_ENABLE_APISERVER=true
fi

env | sort | grep KUBEDB*
echo ""

# create necessary TLS certificates:
# - a local CA key and cert
# - a webhook server key and cert signed by the local CA
$ONESSL create ca-cert
$ONESSL create server-cert server --domains=kubedb-$KUBEDB_OPERATOR_NAME.$KUBEDB_NAMESPACE.svc
export SERVICE_SERVING_CERT_CA=$(cat ca.crt | $ONESSL base64)
export TLS_SERVING_CERT=$(cat server.crt | $ONESSL base64)
export TLS_SERVING_KEY=$(cat server.key | $ONESSL base64)

${SCRIPT_LOCATION}deploy/operator.yaml | $ONESSL envsubst | kubectl apply -f -

if [ "$KUBEDB_ENABLE_RBAC" = true ]; then
  ${SCRIPT_LOCATION}deploy/service-account.yaml | $ONESSL envsubst | kubectl apply -f -
  ${SCRIPT_LOCATION}deploy/rbac-list.yaml | $ONESSL envsubst | kubectl auth reconcile -f -
  ${SCRIPT_LOCATION}deploy/user-roles.yaml | $ONESSL envsubst | kubectl auth reconcile -f -
  ${SCRIPT_LOCATION}deploy/appcatalog-user-roles.yaml | $ONESSL envsubst | kubectl auth reconcile -f -
fi

echo "Applying Pod Sucurity Policies"
${SCRIPT_LOCATION}deploy/psp/operator.yaml | $ONESSL envsubst | kubectl apply -f -
${SCRIPT_LOCATION}deploy/psp/elasticsearch.yaml | $ONESSL envsubst | kubectl apply -f -
${SCRIPT_LOCATION}deploy/psp/memcached.yaml | $ONESSL envsubst | kubectl apply -f -
${SCRIPT_LOCATION}deploy/psp/mongodb.yaml | $ONESSL envsubst | kubectl apply -f -
${SCRIPT_LOCATION}deploy/psp/mysql.yaml | $ONESSL envsubst | kubectl apply -f -
${SCRIPT_LOCATION}deploy/psp/postgres.yaml | $ONESSL envsubst | kubectl apply -f -
${SCRIPT_LOCATION}deploy/psp/redis.yaml | $ONESSL envsubst | kubectl apply -f -

if [ "$KUBEDB_RUN_ON_MASTER" -eq 1 ]; then
  kubectl patch deploy kubedb-$KUBEDB_OPERATOR_NAME -n $KUBEDB_NAMESPACE \
    --patch="$(${SCRIPT_LOCATION}deploy/run-on-master.yaml)"
fi

if [ "$KUBEDB_ENABLE_VALIDATING_WEBHOOK" = true ]; then
  ${SCRIPT_LOCATION}deploy/validating-webhook.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_ENABLE_MUTATING_WEBHOOK" = true ]; then
  ${SCRIPT_LOCATION}deploy/mutating-webhook.yaml | $ONESSL envsubst | kubectl apply -f -
fi

echo
echo "waiting until kubedb operator deployment is ready"
$ONESSL wait-until-ready deployment kubedb-$KUBEDB_OPERATOR_NAME --namespace $KUBEDB_NAMESPACE || {
  echo "KubeDB operator deployment failed to be ready"
  exit 1
}

if [ "$KUBEDB_ENABLE_APISERVER" = true ]; then
  echo "waiting until kubedb apiservice is available"
  for api in "${apiServices[@]}"; do
    $ONESSL wait-until-ready apiservice ${api}.kubedb.com || {
      echo "KubeDB apiservice $api failed to be ready"
      exit 1
    }
  done
fi

if [ "$KUBEDB_OPERATOR_NAME" = "operator" ]; then
  echo "waiting until kubedb crds are ready"
  for crd in "${crds[@]}"; do
    $ONESSL wait-until-ready crd ${crd} || {
      echo "$crd crd failed to be ready"
      exit 1
    }
  done
fi

if [ "$KUBEDB_CATALOG" = "all" ] || [ "$KUBEDB_CATALOG" = "elasticsearch" ]; then
  echo
  echo "installing KubeDB Elasticsearch catalog"
  ${SCRIPT_LOCATION}deploy/kubedb-catalog/elasticsearch.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_CATALOG" = "all" ] || [ "$KUBEDB_CATALOG" = "etcd" ]; then
  echo "installing KubeDB Etcd catalog"
  ${SCRIPT_LOCATION}deploy/kubedb-catalog/etcd.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_CATALOG" = "all" ] || [ "$KUBEDB_CATALOG" = "memcached" ]; then
  echo "installing KubeDB Memcached catalog"
  ${SCRIPT_LOCATION}deploy/kubedb-catalog/memcached.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_CATALOG" = "all" ] || [ "$KUBEDB_CATALOG" = "mongo" ]; then
  echo "installing KubeDB MongoDB catalog"
  ${SCRIPT_LOCATION}deploy/kubedb-catalog/mongodb.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_CATALOG" = "all" ] || [ "$KUBEDB_CATALOG" = "mysql" ]; then
  echo "installing KubeDB MySQL catalog"
  ${SCRIPT_LOCATION}deploy/kubedb-catalog/mysql.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_CATALOG" = "all" ] || [ "$KUBEDB_CATALOG" = "postgres" ]; then
  echo "installing KubeDB Postgres catalog"
  ${SCRIPT_LOCATION}deploy/kubedb-catalog/postgres.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_CATALOG" = "all" ] || [ "$KUBEDB_CATALOG" = "redis" ]; then
  echo "installing KubeDB Redis catalog"
  ${SCRIPT_LOCATION}deploy/kubedb-catalog/redis.yaml | $ONESSL envsubst | kubectl apply -f -
fi

if [ "$KUBEDB_ENABLE_VALIDATING_WEBHOOK" = true ]; then
  echo "checking whether admission webhook(s) are activated or not"
  active=$($ONESSL wait-until-has annotation \
    --apiVersion=apiregistration.k8s.io/v1beta1 \
    --kind=APIService \
    --name=v1alpha1.validators.kubedb.com \
    --key=admission-webhook.appscode.com/active \
    --timeout=5m || {
    echo
    echo "Failed to check if admission webhook(s) are activated or not. Please check operator logs to debug further."
    exit 1
  })
  if [ "$active" = false ]; then
    echo
    echo "Admission webhooks are not activated."
    echo "Enable it by configuring --enable-admission-plugins flag of kube-apiserver."
    echo "For details, visit: https://appsco.de/kube-apiserver-webhooks ."
    echo "After admission webhooks are activated, please uninstall and then reinstall Voyager operator."
    # uninstall misconfigured webhooks to avoid failures
    kubectl delete validatingwebhookconfiguration -l app=kubedb || true
    kubectl delete mutatingwebhookconfiguration -l app=kubedb || true
    exit 1
  fi
fi

# configure prometheus monitoring
 if [ "$MONITORING_ENABLE" = "true" ] && [ "$MONITORING_AGENT" != "$MONITORING_AGENT_NONE" ]; then
   case "$MONITORING_AGENT" in
     "$MONITORING_AGENT_BUILTIN")
        kubectl annotate service kubedb-${KUBEDB_OPERATOR_NAME} -n "$KUBEDB_NAMESPACE" --overwrite \
          prometheus.io/scrape="true" \
          prometheus.io/path="/metrics" \
          prometheus.io/port="8443" \
          prometheus.io/scheme="https"
       ;;
     "$MONITORING_AGENT_COREOS_OPERATOR")
       ${SCRIPT_LOCATION}deploy/monitoring/servicemonitor.yaml | $ONESSL envsubst | kubectl apply -f -
       ;;
   esac

    # if operator monitoring is enabled and prometheus-namespace is provided,
   # create kubedb-operator-apiserver-cert there. this will be mounted on prometheus pod.
   if [ "$PROMETHEUS_NAMESPACE" != "$KUBEDB_NAMESPACE" ]; then
     ${SCRIPT_LOCATION}deploy/monitoring/apiserver-cert.yaml | $ONESSL envsubst | kubectl apply -f -
   fi
 fi

echo
echo "Successfully installed KubeDB operator in $KUBEDB_NAMESPACE namespace!"
