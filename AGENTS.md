# AGENTS.md - KubeDB Installer

This file provides instructions for AI coding agents working in this repository.

## Project Overview

KubeDB installer repository: Helm charts, CRDs, catalog manifests, and deployment scripts for the KubeDB Kubernetes database operator platform. Hosts 23+ Helm charts (provisioner, ops-manager, autoscaler, dashboard, schema-manager, webhook-server, catalog, gitops, metrics, courier, certified, opscenter, ui-server, providers for AWS/Azure/GCP, etc.) and DBVersion catalog manifests for 30+ databases. Also packages an OLM bundle (`bundle/`) using `helm-operator` to deliver KubeDB on OpenShift.

Module: `kubedb.dev/installer` (Go 1.25). Apps under `apis/`, `catalog/`, and `tests/` are non-vendored source; everything else is config/manifests/generated code.

## Build & Development Commands

```bash
# Format, run codegen helpers, and lint (containerized via $BUILD_IMAGE)
make fmt

# Build the installer binary into bin/$OS_$ARCH/installer
make build

# Cross-compile all platforms (linux/amd64, linux/arm, linux/arm64)
make all-build

# Unit tests (runs against $SRC_PKGS = apis catalog tests)
make test
make unit-tests

# golangci-lint
make lint

# Full CI pipeline: verify, check-license, lint, build, unit-tests
make ci

# Code generation (deepcopy clients) via gengo image
make clientset

# Regenerate CRD manifests from apis/ (.crds/, bundle/, config/crd/)
make gen-crds

# Generate values.openapiv3_schema.yaml for each chart from corresponding CRD
make gen-values-schema

# Generate per-chart README.md via chart-doc-gen using doc.yaml + values.yaml
make gen-chart-doc

# Bundle of clientset + manifests + chart docs
make manifests
make gen           # clientset + manifests

# Refresh all generated artifacts: gen + update-catalog + fmt.
# ALWAYS run this before opening a PR so generated files are current.
make refresh

# Bump chart version and chart deps. Pass CHART_VERSION (and optionally APP_VERSION).
make update-charts CHART_VERSION=v0.64.0
# Or update one chart:
make chart-kubedb-provisioner CHART_VERSION=v0.64.0

# OpenAPI spec generation (see olm.mk)
make openapi

# Helm chart-testing (ct) lint+install in a cluster
make ct CT_COMMAND=lint TEST_CHARTS=kubedb-provisioner

# Add or check license headers
make add-license
make check-license

# Verify generated files / go modules are current
make verify
make verify-gen
make verify-modules

# OLM bundle (rules live in olm.mk)
# See olm.mk for bundle-build, bundle-push, catalog-build, etc.
```

Most targets run inside `$BUILD_IMAGE` (`ghcr.io/appscode/golang-dev:1.25`) or `$CODE_GENERATOR_IMAGE` (`ghcr.io/appscode/gengo:release-1.32`) so Docker is required for codegen.

## Project Structure

```
apis/installer/
  v1/                         # API types - one file per Helm chart (Kubedb, KubedbProvisioner, ...)
    types.go                  # Shared types: ImageRef, Container, Monitoring, etc.
    register.go               # Scheme registration of all 14 kinds
    zz_generated.deepcopy.go  # Generated via `make clientset`
  install/install.go          # Scheme installer
  fuzzer/fuzzer.go            # Fuzz funcs for roundtrip tests
  register.go                 # GroupName = "installer.kubedb.com"
catalog/
  kubedb/
    raw/{db}/*.yaml           # Source-of-truth DBVersion manifests per database
    active_versions.json      # Map of {Database: [versions]} that are currently supported
    backup_tasks.json
    restore_tasks.json
    lib.go                    # Embeds raw/* and *.json (`//go:embed`) for consumers
    fmt/main.go               # Regenerates raw/* YAMLs from templates
    gen-version-matrix/main.go    # Produces VersionMatrix.md
    gen-dbaas-license/main.go     # DBaaS license generation
  kubestash/{raw,fmt}/        # KubeStash addon catalog (backup/restore tasks per DB)
  scripts/{db}/imagelist.yaml # Per-component image lists used by image-packer
  imagelist.yaml              # Aggregated image catalog
  copy-images.sh / export-images.sh / import-images.sh / import-into-k3s.sh
  VersionMatrix.md            # Generated DB <-> operator version matrix
charts/                       # 23+ Helm charts (see list below)
  kubedb/                     # Umbrella chart with file:// deps on other charts
    Chart.yaml                # appVersion follows release tag (e.g. v2026.4.27)
    charts/*.tgz              # Vendored sub-chart archives
  kubedb-provisioner/
  kubedb-ops-manager/
  kubedb-autoscaler/
  kubedb-catalog/             # DBVersion CRs (rendered from catalog/kubedb/raw)
  kubedb-kubestash-catalog/
  kubedb-crd-manager/
  kubedb-crds/                # All KubeDB CRDs as a chart
  kubedb-certified/           # Red Hat certified variant
  kubedb-certified-crds/
  kubedb-dashboard/
  kubedb-gitops/
  kubedb-grafana-dashboards/
  kubedb-metrics/             # PrometheusRule + ServiceMonitor configs
  kubedb-courier/
  kubedb-opscenter/           # Umbrella for ops-only install (no provisioner)
  kubedb-perses-dashboards/
  kubedb-provider-aws/
  kubedb-provider-azure/
  kubedb-provider-gcp/
  kubedb-schema-manager/
  kubedb-ui-server/
  kubedb-webhook-server/
  prepare-cluster/
  <chart>/
    Chart.yaml
    values.yaml
    values.openapiv3_schema.yaml  # Generated from CRD spec
    values.schema.json            # Generated from openapiv3 schema
    doc.yaml                      # Input for chart-doc-gen → README.md
    templates/
    crds/                         # Chart-managed CRDs (if any)
    ci/*-values.yaml              # ct test value overrides
.crds/                        # Generated CRD manifests for installer.kubedb.com group
crds/
  kubedb-crds.yaml            # Concatenated CRDs imported from kubedb/apimachinery
  kubedb-catalog-crds.yaml
config/                       # Kubebuilder-style overlays (crd, default, manager, manifests, rbac, prometheus, network-policy, samples, scorecard)
bundle/                       # OLM bundle (manifests + metadata + tests)
bundle.Dockerfile
hack/
  build.sh                    # ldflags-driven `go install ./...`
  fmt.sh / test.sh / e2e.sh
  scripts/                    # update-catalog.sh, ct.sh, cleanup.sh, update-chart-dependencies.sh, open-pr.sh
  license/{bash,dockerfile,go,makefile}.txt
  config/                     # Reserved for build-time config
  kubernetes/kind.yaml
  crd-patch.json
  import_hacks.go
tests/
  check-charts_test.go        # image-packer based image existence/architecture checks
Dockerfile                    # FROM helm-operator; wraps charts/kubedb as a helm operator image
watches.yaml                  # helm-operator watches: Kind=Kubedb → chart=helm-charts/kubedb
Makefile / olm.mk             # Top-level + OLM-specific targets
```

## Key Packages / APIs

- `apis/installer/v1` (Group `installer.kubedb.com`, Version `v1`) - defines a Kubernetes-style spec object per chart, used as the Helm `values.yaml` schema source. Kinds registered in `register.go`:
  `Kubedb`, `KubedbAutoscaler`, `KubedbCatalog`, `KubedbCrdManager`, `KubedbDashboard`, `KubedbGitops`, `KubedbKubestashCatalog`, `KubedbCourier`, `KubedbOpsManager`, `KubedbProvisioner`, `KubedbSchemaManager`, `KubedbWebhookServer`, `KubedbUiServer`, `PrepareCluster` (plus `KubedbProviderAws/Azure/Gcp` types). Shared building blocks (`ImageRef`, `Container`, `ServiceAccountSpec`, `WebHookSpec`, `Monitoring`, `MonitoringAgent`, `ServingCerts`, `CertManagerCerts`, `NetworkPolicySpec`, ...) live in `types.go`.
- `catalog/kubedb` (package `catalog`) - embeds `raw/**` plus `active_versions.json`, `backup_tasks.json`, `restore_tasks.json` via `//go:embed`. `FS()` returns the embedded FS or an override via `--kubedb-catalog-dir` flag (use `AddFlags`/`AddGoFlags`). Helpers: `ActiveDBVersions()`, `BackupTasks()`, `RestoreTasks()`.
- `catalog/kubedb/fmt` - regenerates DBVersion YAMLs in `catalog/kubedb/raw/` from text templates (uses Masterminds sprig + semver).
- `catalog/kubedb/gen-version-matrix` - writes `catalog/VersionMatrix.md`.
- `catalog/kubedb/gen-dbaas-license` - generates DBaaS-specific license artifacts.
- `catalog/kubestash/fmt` - same idea for the KubeStash addon catalog.
- Top-level `watches.yaml` - drives the `helm-operator` runtime: each `Kubedb` CR triggers `helm install/upgrade` of `charts/kubedb` with the spec rendered as Helm values.

## Testing

- `make unit-tests` runs `./hack/test.sh apis catalog tests` inside `$BUILD_IMAGE`. The notable test is:
  - `tests/check-charts_test.go` - calls `image-packer` to validate every image referenced by the charts exists and supports the expected architectures. Skip lists for missing/single-arch images live in `ignoreMissingList` / `archSkipList`.
  - `apis/installer/install/roundtrip_test.go` - apimachinery roundtrip tests using `apis/installer/fuzzer`.
  - `apis/installer/v1/types_test.go`.
- `make ct` runs `helm/chart-testing` (`ct lint-and-install`) against `charts/*`. CI runs it across Kubernetes `v1.29.14`, `v1.31.14`, `v1.33.7`, `v1.35.0` via `kind`.
- CI also `helm template`s every Grafana DB dashboard and runs `metrics-configuration-checker` against `kubedb-metrics`.
- `make verify-gen` and `make verify-modules` are tripwires - PR fails if generated files or `go.mod`/`vendor/` are stale.

## Dependencies

Notable direct dependencies (see `go.mod` for full list):

- `k8s.io/api`, `k8s.io/apimachinery` @ v0.34.3 - core Kubernetes types.
- `kmodules.xyz/client-go`, `kmodules.xyz/resource-metadata`, `kmodules.xyz/schema-checker`, `kmodules.xyz/image-packer`, `kmodules.xyz/go-containerregistry` - AppsCode shared libs.
- `kubedb.dev/apimachinery` - KubeDB CRD types (imported transitively to keep `crds/kubedb-crds.yaml` aligned).
- `kubeops.dev/installer`, `stash.appscode.dev/installer` - cross-product installer types.
- `github.com/Masterminds/sprig/v3`, `github.com/Masterminds/semver/v3`, `gomodules.xyz/semvers` - template helpers for catalog generation.
- `github.com/yudai/gojsondiff`, `github.com/olekukonko/tablewriter`, `gomodules.xyz/go-sh` - used by catalog formatters.
- `sigs.k8s.io/yaml` (replaced by `github.com/kmodules/yaml`).

Active `replace` directives in `go.mod`:

- `github.com/Masterminds/sprig/v3` → `github.com/gomodules/sprig/v3 v3.2.3-0.20220405051441-0a8a99bac1b8`
- `github.com/imdario/mergo` → `v0.3.6`
- `sigs.k8s.io/yaml` → `github.com/kmodules/yaml v1.4.1-...`
- `sigs.k8s.io/controller-runtime` → `github.com/kmodules/controller-runtime ...`
- `k8s.io/apiserver` → `github.com/kmodules/apiserver ...`

Dependencies are vendored (`vendor/` is committed; build uses `GOFLAGS=-mod=vendor`).

External tooling (pulled in via `docker run` in Makefile targets):

| Tool | Image / Version | Purpose |
|------|-----------------|---------|
| golang-dev | `ghcr.io/appscode/golang-dev:1.25` | go build/test/lint/fmt |
| gengo | `ghcr.io/appscode/gengo:release-1.32` | deepcopy, openapi, controller-gen |
| helm-operator | `quay.io/operator-framework/helm-operator:v1.42.2` | runtime base image (Dockerfile) |
| chart-testing | `quay.io/helmpack/chart-testing:v3.13.0` | `make ct` |
| yq | `mikefarah/yq` v3.3.0 (`yqq`) + python yq | YAML editing in CI |
| image-packer | `kmodules.xyz/image-packer` | catalog image audits |

## Code Conventions

- **API group**: everything in `apis/installer` belongs to `installer.kubedb.com/v1`. New chart? Add a `*_types.go` under `apis/installer/v1/`, register the Kind in `register.go`, regenerate deepcopy + CRDs.
- **One spec type per chart**: each chart in `charts/<name>/` corresponds 1:1 with a `Kind<Name>` type. `values.openapiv3_schema.yaml` and `values.schema.json` are generated from that CRD by `make gen-values-schema`; do not hand-edit them.
- **Codegen pragmas**: `// +k8s:deepcopy-gen=package,register`, `// +k8s:openapi-gen=true`, `// +k8s:defaulter-gen=TypeMeta`, `// +groupName=installer.kubedb.com` (see `apis/installer/v1/doc.go`). CRD generation uses kubebuilder markers (`// +kubebuilder:validation:Enum=...`).
- **License headers**: required on all source files. Templates in `hack/license/`. `make add-license` applies them; `make check-license` is part of CI.
- **Lint config** (`.golangci.yml`): default linters + `unparam`; `gofmt` rewrites `interface{}` → `any`; excludes `generated.*\.go`, `client/`, `vendor/`.
- **Generated files**: `zz_generated.deepcopy.go` (apis), `.crds/*.yaml`, `bundle/manifests/*`, `config/crd/bases/*`, `charts/*/values.openapiv3_schema.yaml`, `charts/*/values.schema.json`, `charts/*/README.md`, `catalog/kubedb/raw/*.yaml`, `catalog/VersionMatrix.md` - never edit by hand; rerun the appropriate `make` target.
- **Chart conventions**: each `charts/<name>/` has `Chart.yaml`, `values.yaml`, `doc.yaml` (chart-doc-gen input), generated `README.md`, generated `values.openapiv3_schema.yaml`/`values.schema.json`, `ci/*-values.yaml` for chart-testing. The umbrella `charts/kubedb` and `charts/kubedb-opscenter` declare `file://` dependencies on sibling charts plus OCI deps on `petset`, `operator-shard-manager`, `sidekick`, `supervisor`, `ace-user-roles` from `oci://ghcr.io/appscode-charts`.
- **Chart version bumps**: always go through `make update-charts CHART_VERSION=...` (or per-chart `make chart-<name>`) so `Chart.yaml`, sibling chart deps in `charts/kubedb/Chart.yaml`, and `charts/kubedb-opscenter/Chart.yaml` stay in sync.
- **Certified charts are generated**: never edit `charts/kubedb-certified` or `charts/kubedb-certified-crds` directly. They are derived from `charts/kubedb` and its subcharts. Whenever `charts/kubedb` or a subchart changes, regenerate them:

  ```bash
  rm -rf charts/kubedb-certified charts/kubedb-certified-crds
  chart-packer crd-less --input charts/kubedb --output charts
  chart-packer crd-only --input charts/kubedb --output charts
  make gen-chart-doc
  ```

  - If any subcharts changed, also run `./hack/scripts/update-chart-dependencies.sh`.
  - If any chart changed, run `./hack/scripts/update-catalog.sh`.
- **Catalog edits**: never edit `catalog/kubedb/raw/*.yaml` directly. Update `catalog/kubedb/fmt/templates/` (or `active_versions.json`) and run `make fmt` - it executes `catalog/kubedb/fmt/main.go` and `catalog/kubedb/gen-version-matrix/main.go`. Same pattern for KubeStash via `catalog/kubestash/fmt/main.go`.
- **Image lists**: regenerate with `./hack/scripts/update-catalog.sh` (runs `image-packer list` over `charts/` plus per-component `image-packer generate-scripts`).
- **Build versioning**: `hack/build.sh` injects `main.Version`, `main.GitTag`, `main.CommitHash`, etc. via `-ldflags`. `VERSION` derives from `git describe --tags --always --dirty`, overridden when on a release branch or tag.
- **Vendoring**: keep `vendor/` in sync (`make verify-modules` checks). Run `go mod tidy && go mod vendor` after dependency changes.
- **Source dirs** (`SRC_PKGS` in Makefile): `apis catalog tests`. Build/test/format targets operate only on these.
- **OLM bundle**: `bundle/manifests/installer.kubedb.com_kubedbs.yaml` is regenerated by `make gen-crds`; the rest of `bundle/` is driven by `olm.mk` (`make bundle`, etc.).
