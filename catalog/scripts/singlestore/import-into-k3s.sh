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

k3s ctr images import images/appscode-images-singlestore-node-alma-8.1.32-e3d3cde6da.tar
k3s ctr images import images/appscode-images-singlestore-node-alma-8.5.30-4f46ab16a5.tar
k3s ctr images import images/appscode-images-singlestore-node-alma-8.5.7-bf633c1a54.tar
k3s ctr images import images/appscode-images-singlestore-node-alma-8.7.10-95e2357384.tar
k3s ctr images import images/appscode-images-singlestore-node-alma-8.7.21-f0b8de04d5.tar
k3s ctr images import images/appscode-images-singlestore-node-alma-8.9.3-bfa36a984a.tar
k3s ctr images import images/kubedb-singlestore-coordinator-v0.7.0-rc.0.tar
k3s ctr images import images/kubedb-singlestore-init-8.1-v2.tar
k3s ctr images import images/kubedb-singlestore-init-8.5-v2.tar
k3s ctr images import images/kubedb-singlestore-init-8.7.10-v1.tar
k3s ctr images import images/singlestore-cluster-in-a-box-alma-8.1.32-e3d3cde6da-4.0.16-1.17.6.tar
k3s ctr images import images/singlestore-cluster-in-a-box-alma-8.5.22-fe61f40cd1-4.1.0-1.17.11.tar
k3s ctr images import images/singlestore-cluster-in-a-box-alma-8.5.7-bf633c1a54-4.0.17-1.17.8.tar
k3s ctr images import images/singlestore-cluster-in-a-box-alma-8.7.10-95e2357384-4.1.0-1.17.14.tar
