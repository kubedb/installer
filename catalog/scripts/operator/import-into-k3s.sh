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

set -x

if [ -z "${IMAGE_REGISTRY}" ]; then
    echo "IMAGE_REGISTRY is not set"
    exit 1
fi

TARBALL=${1:-}
tar -zxvf $TARBALL

k3s ctr images import images/appscode-kubectl-nonroot-1.31.tar
k3s ctr images import images/appscode-petset-v0.0.10.tar
k3s ctr images import images/appscode-sidekick-v0.0.11.tar
k3s ctr images import images/kubedb-dashboard-restic-plugin-v0.1.0.tar
k3s ctr images import images/kubedb-dashboard-restic-plugin-v0.11.0.tar
k3s ctr images import images/kubedb-elasticsearch-restic-plugin-v0.16.0.tar
k3s ctr images import images/kubedb-kubedb-autoscaler-v0.38.0.tar
k3s ctr images import images/kubedb-kubedb-crd-manager-v0.8.0.tar
k3s ctr images import images/kubedb-kubedb-gitops-v0.1.0.tar
k3s ctr images import images/kubedb-kubedb-kibana-v0.29.0.tar
k3s ctr images import images/kubedb-kubedb-manifest-plugin-v0.16.0.tar
k3s ctr images import images/kubedb-kubedb-ops-manager-v0.40.0.tar
k3s ctr images import images/kubedb-kubedb-provisioner-v0.53.0.tar
k3s ctr images import images/kubedb-kubedb-schema-manager-v0.29.0.tar
k3s ctr images import images/kubedb-kubedb-ui-server-v0.29.0.tar
k3s ctr images import images/kubedb-kubedb-verifier-v0.1.0.tar
k3s ctr images import images/kubedb-kubedb-webhook-server-v0.29.0.tar
k3s ctr images import images/kubedb-mariadb-csi-snapshotter-plugin-v0.13.0.tar
k3s ctr images import images/kubedb-mariadb-restic-plugin-v0.11.0.tar
k3s ctr images import images/kubedb-mongodb-csi-snapshotter-plugin-v0.14.0.tar
k3s ctr images import images/kubedb-mssqlserver-walg-plugin-v0.0.1.tar
k3s ctr images import images/kubedb-mysql-csi-snapshotter-plugin-v0.14.0.tar
k3s ctr images import images/kubedb-postgres-csi-snapshotter-plugin-v0.14.0.tar
k3s ctr images import images/kubedb-postgres-restic-plugin-v0.16.0_16.1.tar
k3s ctr images import images/kubedb-provider-aws-v0.14.0.tar
k3s ctr images import images/kubedb-provider-azure-v0.14.0.tar
k3s ctr images import images/kubedb-provider-gcp-v0.14.0.tar
k3s ctr images import images/kubedb-redis-restic-plugin-v0.16.0.tar
k3s ctr images import images/kubedb-zookeeper-restic-plugin-v0.9.0.tar
k3s ctr images import images/tianon-toybox-0.8.11.tar
