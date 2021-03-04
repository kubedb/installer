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

SHELL=/bin/bash -o pipefail

GO_PKG   := stash.appscode.dev
REPO     := $(notdir $(shell pwd))
BIN      := catalog

# This version-strategy uses git tags to set the version string
git_branch       := $(shell git rev-parse --abbrev-ref HEAD)
git_tag          := $(shell git describe --exact-match --abbrev=0 2>/dev/null || echo "")
commit_hash      := $(shell git rev-parse --verify HEAD)
commit_timestamp := $(shell date --date="@$$(git show -s --format=%ct)" --utc +%FT%T)

VERSION          := $(shell git describe --tags --always --dirty)
version_strategy := commit_hash
ifdef git_tag
	VERSION := $(git_tag)
	version_strategy := tag
else
	ifeq (,$(findstring $(git_branch),master HEAD))
		ifneq (,$(patsubst release-%,,$(git_branch)))
			VERSION := $(git_branch)
			version_strategy := branch
		endif
	endif
endif

###
### These variables should not need tweaking.
###

SRC_DIRS := # directories which hold app source (not vendored)

DOCKER_PLATFORMS := linux/amd64 linux/arm linux/arm64
BIN_PLATFORMS    := $(DOCKER_PLATFORMS) windows/amd64 darwin/amd64

# Used internally.  Users should pass GOOS and/or GOARCH.
OS   := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

BASEIMAGE_PROD   ?= gcr.io/distroless/static-debian10
BASEIMAGE_DBG    ?= debian:buster

GO_VERSION       ?= 1.16
BUILD_IMAGE      ?= appscode/golang-dev:$(GO_VERSION)

OUTBIN = bin/$(OS)_$(ARCH)/$(BIN)
ifeq ($(OS),windows)
  OUTBIN = bin/$(OS)_$(ARCH)/$(BIN).exe
endif

# Directories that we need created to build/test.
BUILD_DIRS  := bin/$(OS)_$(ARCH)     \
               .go/bin/$(OS)_$(ARCH) \
               .go/cache             \
               hack/config           \
               $(HOME)/.credentials  \
               $(HOME)/.kube         \
               $(HOME)/.minikube

# If you want to build all binaries, see the 'all-build' rule.
# If you want to build all containers, see the 'all-container' rule.
# If you want to build AND push all containers, see the 'all-push' rule.
all: build fmt

# For the following OS/ARCH expansions, we transform OS/ARCH into OS_ARCH
# because make pattern rules don't match with embedded '/' characters.

build: gen
	@$(MAKE) add-license --no-print-directory

version:
	@echo version=$(VERSION)
	@echo version_strategy=$(version_strategy)
	@echo git_tag=$(git_tag)
	@echo git_branch=$(git_branch)
	@echo commit_hash=$(commit_hash)
	@echo commit_timestamp=$(commit_timestamp)

DOCKER_REPO_ROOT := /go/src/$(GO_PKG)/$(REPO)

.PHONY: gen
gen:
	render-gotpl --template=templates/readme.txt --data=catalog.json > README.md
	render-gotpl --template=helm3.sh --pattern=templates/* --data=catalog.json > deploy/helm3.sh
	render-gotpl --template=helm2.sh --pattern=templates/* --data=catalog.json > deploy/helm2.sh
	render-gotpl --template=script.sh --pattern=templates/* --data=catalog.json > deploy/script.sh

fmt: $(BUILD_DIRS)
	@docker run                                                 \
	    -i                                                      \
	    --rm                                                    \
	    -u $$(id -u):$$(id -g)                                  \
	    -v $$(pwd):/src                                         \
	    -w /src                                                 \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
	    -v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
	    -v $$(pwd)/.go/cache:/.cache                            \
	    --env HTTP_PROXY=$(HTTP_PROXY)                          \
	    --env HTTPS_PROXY=$(HTTPS_PROXY)                        \
	    $(BUILD_IMAGE)                                          \
	    /bin/bash -c "                                          \
	        REPO_PKG=$(GO_PKG)                                  \
	        ./hack/fmt.sh $(SRC_DIRS)                           \
	    "

$(BUILD_DIRS):
	@mkdir -p $@

.PHONY: dev
dev: gen fmt

.PHONY: verify
verify: verify-gen

.PHONY: verify-gen
verify-gen: gen fmt
	@if !(git diff --exit-code HEAD); then \
		echo "generated files are out of date, run make gen"; exit 1; \
	fi

.PHONY: add-license
add-license:
	@echo "Adding license header"
	@docker run --rm 	                                 \
		-u $$(id -u):$$(id -g)                           \
		-v /tmp:/.cache                                  \
		-v $$(pwd):$(DOCKER_REPO_ROOT)                   \
		-w $(DOCKER_REPO_ROOT)                           \
		--env HTTP_PROXY=$(HTTP_PROXY)                   \
		--env HTTPS_PROXY=$(HTTPS_PROXY)                 \
		$(BUILD_IMAGE)                                   \
		ltag -t "./hack/license" --excludes "vendor contrib" -v

.PHONY: check-license
check-license:
	@echo "Checking files for license header"
	@docker run --rm 	                                 \
		-u $$(id -u):$$(id -g)                           \
		-v /tmp:/.cache                                  \
		-v $$(pwd):$(DOCKER_REPO_ROOT)                   \
		-w $(DOCKER_REPO_ROOT)                           \
		--env HTTP_PROXY=$(HTTP_PROXY)                   \
		--env HTTPS_PROXY=$(HTTPS_PROXY)                 \
		$(BUILD_IMAGE)                                   \
		ltag -t "./hack/license" --excludes "vendor contrib" --check -v

.PHONY: ci
ci: verify check-license

.PHONY: clean
clean:
	rm -rf .go bin
