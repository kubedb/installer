module kubedb.dev/installer

go 1.23.6

toolchain go1.24.5

require (
	github.com/Masterminds/semver/v3 v3.3.1
	github.com/Masterminds/sprig/v3 v3.3.0
	github.com/google/gofuzz v1.2.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/pflag v1.0.6
	github.com/yudai/gojsondiff v1.0.0
	gomodules.xyz/go-sh v0.1.0
	gomodules.xyz/semvers v0.0.2
	k8s.io/api v0.32.3
	k8s.io/apimachinery v0.32.3
	kmodules.xyz/client-go v0.32.7
	kmodules.xyz/go-containerregistry v0.0.14
	kmodules.xyz/image-packer v0.0.0-20250709183414-f93633723666
	kmodules.xyz/resource-metadata v0.32.1
	kmodules.xyz/schema-checker v0.4.2
	kubedb.dev/apimachinery v0.57.0-rc.0
	kubeops.dev/installer v0.0.0-20250630172252-60882a8ed9ab
	sigs.k8s.io/yaml v1.4.0
	stash.appscode.dev/installer v0.12.2-0.20250630171625-d3c674690e22
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cert-manager/cert-manager v1.18.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.16.3 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/docker/cli v27.5.0+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.2 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.12.1 // indirect
	github.com/evanphx/json-patch v5.9.11+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.9.11 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-sql-driver/mysql v1.9.0 // indirect
	github.com/gobeam/stringy v0.0.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/gnostic-models v0.6.9 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-containerregistry v0.20.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.81.0 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.61.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cast v1.7.0 // indirect
	github.com/vbatts/tar-split v0.11.6 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.virtual-secrets.dev/apimachinery v0.0.1 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/oauth2 v0.28.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/term v0.32.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	golang.org/x/time v0.10.0 // indirect
	gomodules.xyz/encoding v0.0.8 // indirect
	gomodules.xyz/jsonpatch/v2 v2.5.0 // indirect
	gomodules.xyz/jsonpath v0.0.2 // indirect
	gomodules.xyz/mergo v0.3.13 // indirect
	gomodules.xyz/pointer v0.1.0 // indirect
	google.golang.org/protobuf v1.36.3 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.32.3 // indirect
	k8s.io/client-go v0.32.3 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20250318190949-c8a335a9a2ff // indirect
	k8s.io/utils v0.0.0-20241210054802-24370beab758 // indirect
	kmodules.xyz/apiversion v0.2.0 // indirect
	kmodules.xyz/custom-resources v0.32.0 // indirect
	kmodules.xyz/monitoring-agent-api v0.32.1 // indirect
	kmodules.xyz/offshoot-api v0.32.0 // indirect
	kubeops.dev/petset v0.0.11 // indirect
	kubeops.dev/scanner v0.0.19 // indirect
	kubeops.dev/sidekick v0.0.11 // indirect
	open-cluster-management.io/api v1.0.0 // indirect
	sigs.k8s.io/controller-runtime v0.20.4 // indirect
	sigs.k8s.io/gateway-api v1.1.0 // indirect
	sigs.k8s.io/json v0.0.0-20241014173422-cfa47c3a1cc8 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.6.0 // indirect
	x-helm.dev/apimachinery v0.0.17 // indirect
)

replace (
	github.com/Masterminds/sprig/v3 => github.com/gomodules/sprig/v3 v3.2.3-0.20220405051441-0a8a99bac1b8
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.6
	sigs.k8s.io/yaml => github.com/kmodules/yaml v1.4.1-0.20231224133800-a4e3f1abb174
)

replace sigs.k8s.io/controller-runtime => github.com/kmodules/controller-runtime v0.20.3-0.20250221050548-8eabe54e7dda

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.32.3-0.20250221062720-35dc674c7dd6
