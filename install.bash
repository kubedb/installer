helm upgrade -i kubedb charts/kubedb \
  --version v2026.5.18-rc.0 \
  --namespace kubedb --create-namespace \
  --wait --burst-limit=10000 --debug
