#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Installs charts/kubedb-courier-addon-manager on an Open Cluster Management
# topology of 3 kind clusters (1 hub + 2 spokes) and verifies that the courier
# addon agent rolls out to every spoke.
#
# Requires: docker, kind, kubectl, helm, clusteradm, jq
#
# Env:
#   K8S_VERSION    kindest/node tag (default v1.35.0)
#   KEEP_CLUSTERS  keep the kind clusters after the test (default false)
#   TIMEOUT        wait timeout for each check (default 5m)

set -eou pipefail

K8S_VERSION=${K8S_VERSION:-v1.35.0}
KEEP_CLUSTERS=${KEEP_CLUSTERS:-false}
TIMEOUT=${TIMEOUT:-5m}

CHART=charts/kubedb-courier-addon-manager
RELEASE=courier
RELEASE_NS=kubedb
ADDON=courier
AGENT_NS=kubedb
PLACEMENT=courier
PLACEMENT_NS=open-cluster-management

HUB=hub
SPOKES=(spoke1 spoke2)
CLUSTERS=("$HUB" "${SPOKES[@]}")
HUB_CTX=kind-$HUB

log() {
    echo
    echo "==> $*"
}

dump() {
    set +e
    log "Collecting diagnostics"
    for c in "${CLUSTERS[@]}"; do
        local ctx=kind-$c
        echo "---------------- $ctx ----------------"
        kubectl --context "$ctx" get pods -A -o wide
        kubectl --context "$ctx" get events -A --sort-by=.lastTimestamp | tail -n 50
        kubectl --context "$ctx" get klusterlet -o yaml 2>/dev/null
        kubectl --context "$ctx" -n "$AGENT_NS" logs deploy/courier-agent --tail=200 2>/dev/null
    done
    echo "---------------- $HUB_CTX (ocm) ----------------"
    kubectl --context "$HUB_CTX" get managedclusters -o wide
    kubectl --context "$HUB_CTX" get clustermanagementaddon,addontemplate -o yaml
    kubectl --context "$HUB_CTX" get managedclusteraddon -A -o yaml
    kubectl --context "$HUB_CTX" get manifestwork -A -o yaml
    kubectl --context "$HUB_CTX" get placement,placementdecision -A -o yaml
    kubectl --context "$HUB_CTX" -n "$RELEASE_NS" logs -l app.kubernetes.io/instance="$RELEASE" --tail=200
    kubectl --context "$HUB_CTX" -n open-cluster-management-hub logs deploy/cluster-manager-addon-manager-controller --tail=200
}

cleanup() {
    if [ "$KEEP_CLUSTERS" != "true" ]; then
        for c in "${CLUSTERS[@]}"; do
            kind delete cluster --name "$c" || true
        done
    fi
}

on_exit() {
    local rc=$?
    if [ $rc -ne 0 ]; then
        dump
    fi
    cleanup
    exit $rc
}
trap on_exit EXIT

# Runs "$@" once per spoke in parallel (with {} replaced by the spoke name)
# and fails if any of them fails.
for_spokes() {
    local pids=()
    for s in "${SPOKES[@]}"; do
        "${@//\{\}/$s}" &
        pids+=($!)
    done
    for p in "${pids[@]}"; do
        wait "$p"
    done
}

# Waits until the given command succeeds.
retry() {
    local end=$((SECONDS + 300))
    until "$@" >/dev/null 2>&1; do
        if [ $SECONDS -ge $end ]; then
            echo "timed out waiting for: $*"
            return 1
        fi
        sleep 5
    done
}

log "Raising inotify limits for multiple kind clusters"
sudo -n sysctl -w fs.inotify.max_user_watches=524288 fs.inotify.max_user_instances=512 || true

log "Creating kind clusters: ${CLUSTERS[*]}"
pids=()
for c in "${CLUSTERS[@]}"; do
    kind create cluster --name "$c" --image "kindest/node:$K8S_VERSION" --wait 5m &
    pids+=($!)
done
for p in "${pids[@]}"; do
    wait "$p"
done

log "Initializing OCM hub"
clusteradm init --wait --context "$HUB_CTX"
HUB_TOKEN=$(clusteradm get token --context "$HUB_CTX" -o json | jq -r '."hub-token"')
# clusteradm runs on the host, so it needs the host-reachable hub address. With
# --force-internal-endpoint-lookup the klusterlets instead use the in-network
# endpoint (https://hub-control-plane:6443) published in the hub's cluster-info.
HUB_APISERVER=$(kubectl config view --raw -o jsonpath="{.clusters[?(@.name==\"$HUB_CTX\")].cluster.server}")

log "Joining spokes: ${SPOKES[*]}"
for_spokes clusteradm join \
    --hub-token "$HUB_TOKEN" \
    --hub-apiserver "$HUB_APISERVER" \
    --cluster-name {} \
    --force-internal-endpoint-lookup \
    --wait \
    --context kind-{}

log "Accepting spokes"
spoke_list=$(
    IFS=,
    echo "${SPOKES[*]}"
)
retry clusteradm accept --clusters "$spoke_list" --context "$HUB_CTX"
kubectl --context "$HUB_CTX" wait managedcluster --all \
    --for=condition=ManagedClusterConditionAvailable --timeout="$TIMEOUT"

# A spoke runs KubeDB, so it already has the KubeDB and courier CRDs (the latter
# from charts/kubedb-courier). The addon only delivers the agent.
log "Preparing spokes"
for s in "${SPOKES[@]}"; do
    kubectl --context "kind-$s" create namespace "$AGENT_NS"
    kubectl --context "kind-$s" apply --server-side -f charts/kubedb-courier/crds
    kubectl --context "kind-$s" apply --server-side -f crds/kubedb-crds.yaml
done

log "Creating placement $PLACEMENT_NS/$PLACEMENT"
kubectl --context "$HUB_CTX" apply -f - <<EOF
apiVersion: cluster.open-cluster-management.io/v1beta2
kind: ManagedClusterSetBinding
metadata:
  name: global
  namespace: $PLACEMENT_NS
spec:
  clusterSet: global
---
apiVersion: cluster.open-cluster-management.io/v1beta1
kind: Placement
metadata:
  name: $PLACEMENT
  namespace: $PLACEMENT_NS
spec:
  clusterSets:
    - global
EOF

log "Installing $CHART on the hub"
helm install "$RELEASE" "$CHART" \
    --kube-context "$HUB_CTX" \
    --namespace "$RELEASE_NS" --create-namespace \
    --set agent.placement.name="$PLACEMENT" \
    --set agent.placement.namespace="$PLACEMENT_NS" \
    --wait --timeout "$TIMEOUT"

log "Checking hub"
kubectl --context "$HUB_CTX" -n "$RELEASE_NS" wait deploy \
    -l app.kubernetes.io/instance="$RELEASE" --for=condition=Available --timeout="$TIMEOUT"
kubectl --context "$HUB_CTX" get clustermanagementaddon "$ADDON"

for s in "${SPOKES[@]}"; do
    log "Checking addon registration for $s"
    retry kubectl --context "$HUB_CTX" -n "$s" get managedclusteraddon "$ADDON"
    # RegistrationApplied=False (SetPermissionFailed) is silent otherwise: the addon
    # still reports Available while the spoke never gets a usable hub kubeconfig.
    kubectl --context "$HUB_CTX" -n "$s" wait managedclusteraddon "$ADDON" \
        --for=condition=RegistrationApplied --timeout="$TIMEOUT"
    kubectl --context "$HUB_CTX" -n "$s" wait managedclusteraddon "$ADDON" \
        --for=condition=Available --timeout="$TIMEOUT"
    # the addon-manager binds courier-branchwork in the spoke's own hub namespace
    retry bash -c "kubectl --context $HUB_CTX -n $s get rolebinding -o json |
        jq -e '[.items[] | select(.roleRef.kind == \"ClusterRole\" and .roleRef.name == \"courier-branchwork\")] | length > 0'"

    log "Checking courier agent on $s"
    retry kubectl --context "kind-$s" -n "$AGENT_NS" get deploy courier-agent
    kubectl --context "kind-$s" -n "$AGENT_NS" rollout status deploy/courier-agent --timeout="$TIMEOUT"
    retry kubectl --context "kind-$s" -n "$AGENT_NS" get secret "$ADDON-hub-kubeconfig"
    # the agent's controllers only start after it wins leader election
    retry bash -c "kubectl --context kind-$s -n $AGENT_NS logs deploy/courier-agent | grep -q 'successfully acquired lease'"
    # a watched kind without its CRD makes the agent exit once the cache sync times out
    if kubectl --context "kind-$s" -n "$AGENT_NS" logs deploy/courier-agent | grep 'if kind is a CRD'; then
        echo "courier agent on $s is watching kinds whose CRDs are missing"
        exit 1
    fi
done

log "Uninstalling $CHART from the hub"
helm uninstall "$RELEASE" --kube-context "$HUB_CTX" --namespace "$RELEASE_NS" --wait --timeout "$TIMEOUT"
for s in "${SPOKES[@]}"; do
    log "Checking courier agent is removed from $s"
    retry bash -c "! kubectl --context kind-$s -n $AGENT_NS get deploy courier-agent"
done

log "OCM courier addon test passed"
