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

k3s ctr images import images/tianon-toybox-0.8.11.tar
k3s ctr images import images/appscode-kubectl-nonroot-1.34.tar
k3s ctr images import images/appscode-petset-v0.1.0.tar
k3s ctr images import images/appscode-sidekick-v0.0.15.tar
k3s ctr images import images/kubedb-cassandra-medusa-plugin-v0.13.0.tar
k3s ctr images import images/kubedb-clickhouse-backup-plugin-v0.3.0.tar
k3s ctr images import images/kubedb-dashboard-restic-plugin-v0.24.0.tar
k3s ctr images import images/kubedb-elasticsearch-restic-plugin-v0.29.0.tar
k3s ctr images import images/kubedb-kubedb-autoscaler-v0.51.0.tar
k3s ctr images import images/kubedb-kubedb-courier-v0.6.0.tar
k3s ctr images import images/kubedb-kubedb-crd-manager-v0.21.0.tar
k3s ctr images import images/kubedb-kubedb-gitops-v0.14.0.tar
k3s ctr images import images/kubedb-kubedb-kibana-v0.42.0.tar
k3s ctr images import images/kubedb-kubedb-manifest-plugin-v0.29.0.tar
k3s ctr images import images/kubedb-kubedb-ops-manager-v0.53.0.tar
k3s ctr images import images/kubedb-kubedb-provisioner-v0.66.0.tar
k3s ctr images import images/kubedb-kubedb-schema-manager-v0.42.0.tar
k3s ctr images import images/kubedb-kubedb-ui-server-v0.42.0.tar
k3s ctr images import images/kubedb-kubedb-verifier-v0.17.0.tar
k3s ctr images import images/kubedb-kubedb-webhook-server-v0.42.0.tar
k3s ctr images import images/kubedb-mariadb-csi-snapshotter-plugin-v0.26.0.tar
k3s ctr images import images/kubedb-mariadb-restic-plugin-v0.24.0_10.11.6-jammy.tar
k3s ctr images import images/kubedb-mariadb-restic-plugin-v0.24.0_10.4.32-focal.tar
k3s ctr images import images/kubedb-mariadb-restic-plugin-v0.24.0_10.6.16-focal.tar
k3s ctr images import images/kubedb-mariadb-restic-plugin-v0.24.0_11.1.3-jammy.tar
k3s ctr images import images/kubedb-mongodb-csi-snapshotter-plugin-v0.27.0.tar
k3s ctr images import images/kubedb-mongodb-restic-plugin-v0.29.0_4.2.3.tar
k3s ctr images import images/kubedb-mongodb-restic-plugin-v0.29.0_4.4.6.tar
k3s ctr images import images/kubedb-mongodb-restic-plugin-v0.29.0_5.0.15.tar
k3s ctr images import images/kubedb-mongodb-restic-plugin-v0.29.0_5.0.3.tar
k3s ctr images import images/kubedb-mongodb-restic-plugin-v0.29.0_6.0.5.tar
k3s ctr images import images/kubedb-mongodb-restic-plugin-v0.29.0_8.0.3.tar
k3s ctr images import images/kubedb-mssqlserver-walg-plugin-v0.20.0.tar
k3s ctr images import images/kubedb-mysql-csi-snapshotter-plugin-v0.27.0.tar
k3s ctr images import images/kubedb-mysql-restic-plugin-v0.29.0_5.7.25.tar
k3s ctr images import images/kubedb-mysql-restic-plugin-v0.29.0_8.0.21.tar
k3s ctr images import images/kubedb-mysql-restic-plugin-v0.29.0_8.0.3.tar
k3s ctr images import images/kubedb-mysql-restic-plugin-v0.29.0_8.4.2.tar
k3s ctr images import images/kubedb-mysql-restic-plugin-v0.29.0_9.0.1.tar
k3s ctr images import images/kubedb-neo4j-backup-plugin-v0.2.0.tar
k3s ctr images import images/kubedb-postgres-csi-snapshotter-plugin-v0.27.0.tar
k3s ctr images import images/kubedb-postgres-restic-plugin-v0.29.0_12.17.tar
k3s ctr images import images/kubedb-postgres-restic-plugin-v0.29.0_14.10.tar
k3s ctr images import images/kubedb-postgres-restic-plugin-v0.29.0_16.4.tar
k3s ctr images import images/kubedb-postgres-restic-plugin-v0.29.0_17.2.tar
k3s ctr images import images/kubedb-postgres-restic-plugin-v0.29.0_18.2.tar
k3s ctr images import images/kubedb-provider-aws-v0.27.0.tar
k3s ctr images import images/kubedb-provider-azure-v0.27.0.tar
k3s ctr images import images/kubedb-provider-gcp-v0.27.0.tar
k3s ctr images import images/kubedb-qdrant-restic-plugin-v0.2.0.tar
k3s ctr images import images/kubedb-redis-restic-plugin-v0.29.0.tar
k3s ctr images import images/kubedb-singlestore-restic-plugin-v0.24.0_alma-8.1.32-e3d3cde6da.tar
k3s ctr images import images/kubedb-singlestore-restic-plugin-v0.24.0_alma-8.5.7-bf633c1a54.tar
k3s ctr images import images/kubedb-xtrabackup-restic-plugin-v0.14.0_2.4.29.tar
k3s ctr images import images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.0.35.tar
k3s ctr images import images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.1.0.tar
k3s ctr images import images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.2.0.tar
k3s ctr images import images/kubedb-xtrabackup-restic-plugin-v0.14.0_8.4.0.tar
k3s ctr images import images/kubedb-zookeeper-restic-plugin-v0.21.0.tar
