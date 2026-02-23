# Build the manager binary
FROM quay.io/operator-framework/helm-operator:v1.42.0

ENV HOME=/opt/helm
COPY watches.yaml ${HOME}/watches.yaml
COPY charts/kubedb  ${HOME}/helm-charts/kubedb
WORKDIR ${HOME}
