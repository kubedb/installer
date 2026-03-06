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

# Build the manager binary
FROM quay.io/operator-framework/helm-operator:v1.42.0

ARG VERSION

LABEL org.opencontainers.image.source="https://github.com/kubedb/installer" \
    name="KubeDB Operator Installer" \
    maintainer=AppsCode \
    vendor=AppsCode \
    version=${VERSION} \
    release=${VERSION} \
    summary="KubeDB Operator Installer" \
    description="KubeDB Operator Installer"

USER root
RUN mkdir -p /licenses
COPY LICENSE.md /licenses/
USER ${USER_UID} # helm

ENV HOME=/opt/helm
COPY watches.yaml ${HOME}/watches.yaml
COPY charts/kubedb  ${HOME}/helm-charts/kubedb
WORKDIR ${HOME}
