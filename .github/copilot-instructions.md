# KubeDB Installer Development Guide

## Project Overview

This is the KubeDB installer repository containing Helm charts and deployment scripts for KubeDB - a production-ready database operator for Kubernetes. The repository manages 20+ charts for various databases (MongoDB, PostgreSQL, MySQL, Elasticsearch, Kafka, Redis, etc.) and operator components.

## Architecture

### Key Components

- **Helm Charts** (`charts/`): 24 charts organized as:
  - `kubedb`: Meta-chart that bundles all components via dependencies
  - Operator components: `kubedb-provisioner`, `kubedb-ops-manager`, `kubedb-autoscaler`, `kubedb-webhook-server`, `kubedb-schema-manager`, `kubedb-crd-manager`
  - Catalogs: `kubedb-catalog` (database versions), `kubedb-kubestash-catalog` (backup/restore tasks)
  - Provider charts: `kubedb-provider-aws`, `kubedb-provider-azure`, `kubedb-provider-gcp`
  - Additional: `kubedb-metrics`, `kubedb-dashboard`, `kubedb-opscenter`

- **Catalog System** (`catalog/`): Database version definitions and container images
  - `catalog/kubedb/raw/`: YAML definitions for each database version (e.g., `postgres/postgres-16.1.yaml`)
  - Embedded via Go's `//go:embed` in `catalog/kubedb/lib.go`
  - Generated files: `active_versions.json`, `backup_tasks.json`, `restore_tasks.json`

- **APIs** (`apis/installer/v1alpha1/`): Go types for chart values (used as CRDs for GitOps)
  - Each chart has a corresponding `*_types.go` file defining its Helm values schema

- **CRDs** (`crds/`): Generated Kubernetes CRD manifests
  - `kubedb-crds.yaml`: Core database CRDs
  - `kubedb-catalog-crds.yaml`: Version catalog CRDs

## Critical Workflows

### Building & Code Generation

```bash
# Full code generation pipeline
make gen          # Generates clientset, OpenAPI schema, CRDs, chart docs

# Individual generation steps
make clientset    # Generate deepcopy methods and typed clients
make openapi      # Generate OpenAPI v3 schemas from Go types
make gen-crds     # Generate CRD manifests from kubebuilder markers
make gen-values-schema  # Extract OpenAPI schema to values.openapiv3_schema.yaml
make gen-chart-doc      # Generate chart README from doc.yaml + values.yaml

# Format code and regenerate catalog metadata
make fmt          # Runs formatters + catalog/kubedb/fmt + gen-version-matrix
```

### Catalog Management

```bash
# Update catalog with new database versions
./hack/scripts/update-catalog.sh   # Uses image-packer to generate image lists and scripts

# This generates:
# - catalog/imagelist.yaml (all container images)
# - catalog/scripts/*/imagelist.yaml (per-database images)
# - catalog/export-images.sh, import-images.sh (image management scripts)
```

### Chart Development

```bash
# Update chart dependencies from upstream
./hack/scripts/update-chart-dependencies.sh

# Test charts with chart-testing (ct)
make ct CT_COMMAND=install TEST_CHARTS=charts/kubedb

# Update chart versions (typically done by CI)
make update-charts CHART_VERSION=v2026.1.19 APP_VERSION=v2026.1.19
```

### Testing & Validation

```bash
make lint         # golangci-lint
make verify       # Verify go modules and generated code
make ci           # Full CI suite: verify + check-license + lint + build + unit-tests

# Check image architecture compatibility
go test ./tests/  # Validates multi-arch image availability
```

## Project-Specific Conventions

### Version Management

- **Chart versioning**: Uses date-based versions like `v2026.1.19` across all charts
- **Dependencies**: `kubedb` chart references other charts via `file://../chart-name` paths
- **App versions**: Component charts have independent `appVersion` fields (e.g., `v0.60.0` for provisioner)

### Catalog Version Files

Database versions in `catalog/kubedb/raw/` follow this structure:

```yaml
apiVersion: catalog.kubedb.com/v1alpha1
kind: ClickHouseVersion  # Or PostgresVersion, MongoDBVersion, etc.
metadata:
  name: 24.4.1
spec:
  version: 24.4.1
  db:
    image: clickhouse/clickhouse-server:24.4.1
  initContainer:
    image: ghcr.io/kubedb/clickhouse-init:24.4.1-v3
  # ... additional images for coordinator, exporter, etc.
```

### Go Module Structure

- Uses `GO111MODULE=on` with vendor directory
- Docker-based builds via `BUILD_IMAGE=ghcr.io/appscode/golang-dev:1.25`
- All builds run in containers to ensure consistency

### Kubebuilder Markers

APIs use extensive kubebuilder markers for CRD generation:
- `+kubebuilder:object:root=true` for CRD root types
- `+kubebuilder:resource:path=...` for resource naming
- `+kubebuilder:printcolumn:...` for kubectl output columns
- Check vendor directory for examples: `vendor/kubedb.dev/apimachinery/apis/catalog/v1alpha1/`

### Docker-Based Tooling

All code generation and builds use Docker containers:
- **Code generation**: `CODE_GENERATOR_IMAGE=ghcr.io/appscode/gengo:release-1.32`
- **Go builds**: `BUILD_IMAGE=ghcr.io/appscode/golang-dev:1.25`
- **Chart testing**: `CHART_TEST_IMAGE=quay.io/helmpack/chart-testing:v3.13.0`

This ensures consistent tooling versions across environments.

### Feature Gates

Database support is controlled via feature gates in `charts/kubedb/values.yaml`:

```yaml
global:
  featureGates:
    Cassandra: false   # Disable Cassandra catalog entries
    ClickHouse: false
    Kafka: true        # Enable by default
```

Set in catalog chart: `charts/kubedb-catalog/values.yaml` → `featureGates: {}`

## File Organization Patterns

- **Per-database scripts**: `catalog/scripts/{postgres,mysql,kafka,...}/` contain image lists
- **Chart structure**: Each chart has `Chart.yaml`, `values.yaml`, `doc.yaml`, `templates/`, `ci/` directories
- **Generated files**: `.crds/*.yaml` are auto-generated from `apis/` - do not edit manually

## Common Pitfalls

1. **Modifying generated CRDs directly**: Always edit Go types in `apis/` and run `make gen-crds`
2. **Forgetting to run `make fmt`**: This updates catalog metadata and version matrices
3. **Chart dependency versions**: When bumping `kubedb` chart version, also update dependency versions in `Chart.yaml`
4. **Image architecture checks**: New images must support `linux/amd64`, `linux/arm64`, `linux/arm` (exceptions in `tests/check-charts_test.go`)

## Useful Commands

```bash
# Check what changed after generation
git diff

# Clean build artifacts
make clean

# Add license headers to new files
make add-license

# Verify everything is up-to-date
make verify-gen
```
