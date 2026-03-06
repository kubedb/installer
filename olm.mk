# VERSION_NO_PREFIX defines the project version for the bundle.
# Update this value when you upgrade the version of your project.
# To re-generate a bundle for another specific version without changing the standard setup, you can:
# - use the VERSION_NO_PREFIX as arg of the bundle target (e.g make bundle VERSION_NO_PREFIX=0.0.2)
# - use environment variables to overwrite this value (e.g export VERSION_NO_PREFIX=0.0.2)
VERSION_NO_PREFIX = $(VERSION:v%=%)

# CHANNELS define the bundle channels used in the bundle.
# Add a new line here if you would like to change its default config. (E.g CHANNELS = "candidate,fast,stable")
# To re-generate a bundle for other specific channels without changing the standard setup, you can:
# - use the CHANNELS as arg of the bundle target (e.g make bundle CHANNELS=candidate,fast,stable)
# - use environment variables to overwrite this value (e.g export CHANNELS="candidate,fast,stable")
ifneq ($(origin CHANNELS), undefined)
BUNDLE_CHANNELS := --channels=$(CHANNELS)
endif

# DEFAULT_CHANNEL defines the default channel used in the bundle.
# Add a new line here if you would like to change its default config. (E.g DEFAULT_CHANNEL = "stable")
# To re-generate a bundle for any other default channel without changing the default setup, you can:
# - use the DEFAULT_CHANNEL as arg of the bundle target (e.g make bundle DEFAULT_CHANNEL=stable)
# - use environment variables to overwrite this value (e.g export DEFAULT_CHANNEL="stable")
ifneq ($(origin DEFAULT_CHANNEL), undefined)
BUNDLE_DEFAULT_CHANNEL := --default-channel=$(DEFAULT_CHANNEL)
endif
BUNDLE_METADATA_OPTS ?= $(BUNDLE_CHANNELS) $(BUNDLE_DEFAULT_CHANNEL)

# IMAGE_TAG_BASE defines the docker.io namespace and part of the image name for remote images.
# This variable is used to construct full image tags for bundle and catalog images.
#
# For example, running 'make bundle-build bundle-push catalog-build catalog-push' will build and push both
# ghcr.io/kubedb/installer-bundle:$VERSION and ghcr.io/kubedb/installer-catalog:$VERSION.
IMAGE_TAG_BASE ?= ghcr.io/kubedb/kubedb-operator

# BUNDLE_IMG defines the image:tag used for the bundle.
# You can use it as an arg. (E.g make bundle-build BUNDLE_IMG=<some-registry>/<project-name-bundle>:<tag>)
BUNDLE_IMG ?= $(IMAGE_TAG_BASE)-bundle:$(VERSION)

# BUNDLE_GEN_FLAGS are the flags passed to the operator-sdk generate bundle command
BUNDLE_GEN_FLAGS ?= -q --overwrite --version $(VERSION_NO_PREFIX) $(BUNDLE_METADATA_OPTS)

# USE_IMAGE_DIGESTS defines if images are resolved via tags or digests
# You can enable this value if you would like to use SHA Based Digests
# To enable set flag to true
USE_IMAGE_DIGESTS ?= false
ifeq ($(USE_IMAGE_DIGESTS), true)
	BUNDLE_GEN_FLAGS += --use-image-digests
endif

# Set the Operator SDK version to use. By default, what is installed on the system is used.
# This is useful for CI or a project to utilize a specific version of the operator-sdk toolkit.
OPERATOR_SDK_VERSION ?= v1.42.0

# Container tool to use for building and pushing images
CONTAINER_TOOL ?= docker

# Image URL to use all building/pushing image targets
IMG ?= $(IMAGE_TAG_BASE):$(VERSION)

.PHONY: all
all: docker-build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build

.PHONY: run
run: helm-operator ## Run against the configured Kubernetes cluster in ~/.kube/config
	$(HELM_OPERATOR) run

TEST_BUNDLE_TIMEOUT ?= 10m
TEST_BUNDLE_NAMESPACE ?= default
TEST_BUNDLE_SUBSCRIPTION ?= installer-v$(shell echo "$(VERSION_NO_PREFIX)" | tr '.+' '-')-sub

.PHONY: cleanup-bundle
cleanup-bundle: ## Cleanup OLM resources used by run-bundle in the target namespace
	@kubectl delete catalogsource installer-catalog -n $(TEST_BUNDLE_NAMESPACE) --ignore-not-found=true >/dev/null 2>&1 || true
	@kubectl wait --for=delete catalogsource/installer-catalog -n $(TEST_BUNDLE_NAMESPACE) --timeout=120s >/dev/null 2>&1 || true
	@kubectl delete subscription $(TEST_BUNDLE_SUBSCRIPTION) -n $(TEST_BUNDLE_NAMESPACE) --ignore-not-found=true >/dev/null 2>&1 || true
	@kubectl wait --for=delete subscription/$(TEST_BUNDLE_SUBSCRIPTION) -n $(TEST_BUNDLE_NAMESPACE) --timeout=120s >/dev/null 2>&1 || true
	@kubectl delete operatorgroup operator-sdk-og -n $(TEST_BUNDLE_NAMESPACE) --ignore-not-found=true >/dev/null 2>&1 || true
	@kubectl wait --for=delete operatorgroup/operator-sdk-og -n $(TEST_BUNDLE_NAMESPACE) --timeout=120s >/dev/null 2>&1 || true

.PHONY: run-bundle
run-bundle: operator-sdk cleanup-bundle ## Run against the configured Kubernetes cluster in ~/.kube/config
	- rm -rf bundle-* cache
	$(OPERATOR_SDK) run bundle $(BUNDLE_IMG) --timeout $(TEST_BUNDLE_TIMEOUT) -n $(TEST_BUNDLE_NAMESPACE)

.PHONY: olm-install
olm-install: operator-sdk ## Install OLM in the configured Kubernetes cluster
	$(OPERATOR_SDK) olm install

.PHONY: olm-uninstall
olm-uninstall: operator-sdk ## Uninstall OLM from the configured Kubernetes cluster
	$(OPERATOR_SDK) olm uninstall

.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	$(CONTAINER_TOOL) build -t ${IMG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	$(CONTAINER_TOOL) push ${IMG}

# PLATFORMS defines the target platforms for  the manager image be build to provide support to multiple
# architectures. (i.e. make docker-buildx IMG=myregistry/mypoperator:0.0.1). To use this option you need to:
# - able to use docker buildx . More info: https://docs.docker.com/build/buildx/
# - have enable BuildKit, More info: https://docs.docker.com/develop/develop-images/build_enhancements/
# - be able to push the image for your registry (i.e. if you do not inform a valid value via IMG=<myregistry/image:<tag>> than the export will fail)
# To properly provided solutions that supports more than one platform you should use this option.
PLATFORMS ?= linux/arm64,linux/amd64
.PHONY: docker-buildx
docker-buildx: ## Build and push docker image for the manager for cross-platform support
	- $(CONTAINER_TOOL) buildx create --name project-v3-builder
	$(CONTAINER_TOOL) buildx use project-v3-builder
	- $(CONTAINER_TOOL) buildx build --push --platform=$(PLATFORMS) --tag ${IMG} -f Dockerfile .
	- $(CONTAINER_TOOL) buildx rm project-v3-builder

.PHONY: docker-certify-redhat
docker-certify-redhat:
	@preflight check container ${IMG} \
		--submit \
		--certification-component-id=69aa458f8d2c13edb2ff0e74

##@ Deployment

.PHONY: install
install: kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

.PHONY: uninstall
uninstall: kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

.PHONY: deploy
deploy: kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

.PHONY: undeploy
undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | kubectl delete -f -

OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m | sed 's/x86_64/amd64/' | sed 's/aarch64/arm64/')

.PHONY: kustomize
KUSTOMIZE = $(shell pwd)/bin/kustomize
kustomize: ## Download kustomize locally if necessary.
ifeq (,$(wildcard $(KUSTOMIZE)))
ifeq (,$(shell which kustomize 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(KUSTOMIZE)) ;\
	curl -sSLo - https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v5.6.0/kustomize_v5.6.0_$(OS)_$(ARCH).tar.gz | \
	tar xzf - -C bin/ ;\
	}
else
KUSTOMIZE = $(shell which kustomize)
endif
endif

.PHONY: helm-operator
HELM_OPERATOR = $(shell pwd)/bin/helm-operator
helm-operator: ## Download helm-operator locally if necessary, preferring the $(pwd)/bin path over global if both exist.
ifeq (,$(wildcard $(HELM_OPERATOR)))
ifeq (,$(shell which helm-operator 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(HELM_OPERATOR)) ;\
	curl -sSLo $(HELM_OPERATOR) https://github.com/operator-framework/operator-sdk/releases/download/v1.42.0/helm-operator_$(OS)_$(ARCH) ;\
	chmod +x $(HELM_OPERATOR) ;\
	}
else
HELM_OPERATOR = $(shell which helm-operator)
endif
endif

.PHONY: operator-sdk
OPERATOR_SDK ?= $(LOCALBIN)/operator-sdk
operator-sdk: ## Download operator-sdk locally if necessary.
ifeq (,$(wildcard $(OPERATOR_SDK)))
ifeq (, $(shell which operator-sdk 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(OPERATOR_SDK)) ;\
	curl -sSLo $(OPERATOR_SDK) https://github.com/operator-framework/operator-sdk/releases/download/$(OPERATOR_SDK_VERSION)/operator-sdk_$(OS)_$(ARCH) ;\
	chmod +x $(OPERATOR_SDK) ;\
	}
else
OPERATOR_SDK = $(shell which operator-sdk)
endif
endif

.PHONY: role-aggregator
ROLE_AGGREGATOR = $(shell pwd)/bin/role-aggregator
role-aggregator: ## Download role-aggregator locally if necessary.
ifeq (,$(wildcard $(ROLE_AGGREGATOR)))
ifeq (, $(shell which role-aggregator 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(ROLE_AGGREGATOR)) ;\
	curl -sSLo $(ROLE_AGGREGATOR) https://github.com/kmodules/role-aggregator/releases/latest/download/role-aggregator-$(OS)-$(ARCH) ;\
	chmod +x $(ROLE_AGGREGATOR) ;\
	}
else
ROLE_AGGREGATOR = $(shell which role-aggregator)
endif
endif

.PHONY: gen-custom-role
gen-custom-role: role-aggregator ## Generate custom role for kubedb from selected charts.
	$(ROLE_AGGREGATOR) \
		--name kubedb-role \
		--dir config/rbac/role.yaml \
		--chart charts/kubedb-autoscaler \
		--chart charts/kubedb-crd-manager \
		--chart charts/kubedb-dashboard \
		--chart charts/kubedb-gitops \
		--chart charts/kubedb-migrator \
		--chart charts/kubedb-ops-manager \
		--chart charts/kubedb-provisioner \
		--chart charts/kubedb-schema-manager \
		--chart charts/kubedb-webhook-server \
		--output config/rbac/custom_role.yaml


.PHONY: bundle
bundle: kustomize operator-sdk ## Generate bundle manifests and metadata, then validate generated files.
	$(OPERATOR_SDK) generate kustomize manifests -q
	cd config/manager && $(KUSTOMIZE) edit set image controller=$(IMG)
	@$(MAKE) gen-custom-role --no-print-directory
	$(KUSTOMIZE) build config/manifests | $(OPERATOR_SDK) generate bundle $(BUNDLE_GEN_FLAGS)
	$(OPERATOR_SDK) bundle validate ./bundle
	@awk 'BEGIN{s=0} {if(!s && ($$0=="" || $$0=="---")){next} s=1; print}' config/crd/bases/installer.kubedb.com_kubedbs.yaml > bundle/manifests/installer.kubedb.com_kubedbs.yaml

.PHONY: bundle-build
bundle-build: ## Build the bundle image.
	$(CONTAINER_TOOL) build -f bundle.Dockerfile -t $(BUNDLE_IMG) .

.PHONY: bundle-push
bundle-push: ## Push the bundle image.
	$(MAKE) docker-push IMG=$(BUNDLE_IMG)

.PHONY: opm
OPM = $(LOCALBIN)/opm
opm: ## Download opm locally if necessary.
ifeq (,$(wildcard $(OPM)))
ifeq (,$(shell which opm 2>/dev/null))
	@{ \
	set -e ;\
	mkdir -p $(dir $(OPM)) ;\
	curl -sSLo $(OPM) https://github.com/operator-framework/operator-registry/releases/download/v1.55.0/$(OS)-$(ARCH)-opm ;\
	chmod +x $(OPM) ;\
	}
else
OPM = $(shell which opm)
endif
endif

# A comma-separated list of bundle images (e.g. make catalog-build BUNDLE_IMGS=example.com/operator-bundle:v0.1.0,example.com/operator-bundle:v0.2.0).
# These images MUST exist in a registry and be pull-able.
BUNDLE_IMGS ?= $(BUNDLE_IMG)

# The image tag given to the resulting catalog image (e.g. make catalog-build CATALOG_IMG=example.com/operator-catalog:v0.2.0).
CATALOG_IMG ?= $(IMAGE_TAG_BASE)-catalog:$(VERSION)

# Set CATALOG_BASE_IMG to an existing catalog image tag to add $BUNDLE_IMGS to that image.
ifneq ($(origin CATALOG_BASE_IMG), undefined)
FROM_INDEX_OPT := --from-index $(CATALOG_BASE_IMG)
endif

# Build a catalog image by adding bundle images to an empty catalog using the operator package manager tool, 'opm'.
# This recipe invokes 'opm' in 'semver' bundle add mode. For more information on add modes, see:
# https://github.com/operator-framework/community-operators/blob/7f1438c/docs/packaging-operator.md#updating-your-existing-operator
.PHONY: catalog-build
catalog-build: opm ## Build a catalog image.
	$(OPM) index add --container-tool $(CONTAINER_TOOL) --mode semver --tag $(CATALOG_IMG) --bundles $(BUNDLE_IMGS) $(FROM_INDEX_OPT)

# Push the catalog image.
.PHONY: catalog-push
catalog-push: ## Push a catalog image.
	$(MAKE) docker-push IMG=$(CATALOG_IMG)
