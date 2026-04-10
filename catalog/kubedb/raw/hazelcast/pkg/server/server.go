/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"context"
	"crypto/tls"
	"os"
	"path/filepath"
	"time"

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	kubedb "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"kubedb.dev/hazelcast/pkg/controller"

	admissionv1 "k8s.io/api/admission/v1"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	core "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
	cu "kmodules.xyz/client-go/client"
	appcatalog "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	scapi "kubeops.dev/operator-shard-manager/api/v1alpha1"
	psapi "kubeops.dev/petset/apis/apps/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/certwatcher"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/filters"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

var (
	Scheme   = runtime.NewScheme()
	setupLog = log.Log.WithName("setup")
)

func init() {
	utilruntime.Must(admissionv1.AddToScheme(Scheme))
	utilruntime.Must(admissionv1beta1.AddToScheme(Scheme))
	utilruntime.Must(clientgoscheme.AddToScheme(Scheme))
	utilruntime.Must(catalog.AddToScheme(Scheme))
	utilruntime.Must(kubedb.AddToScheme(Scheme))
	utilruntime.Must(appcatalog.AddToScheme(Scheme))
	utilruntime.Must(apiextensionsv1.AddToScheme(Scheme))
	utilruntime.Must(psapi.AddToScheme(Scheme))
	utilruntime.Must(netv1.AddToScheme(Scheme))
	utilruntime.Must(scapi.AddToScheme(Scheme))

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

type KubeDBWebhookConfig struct {
	WebhookConfig *controller.WebhookConfig
}

// KubeDBWebhookServer contains state for a Kubernetes cluster master/api server.
type KubeDBWebhookServer struct {
	Manager manager.Manager
}

func (op *KubeDBWebhookServer) Run(ctx context.Context) error {
	return op.Manager.Start(ctx)
}

type completedConfig struct {
	WebhookConfig *controller.WebhookConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *KubeDBWebhookConfig) Complete() CompletedConfig {
	completedCfg := completedConfig{
		c.WebhookConfig,
	}

	return CompletedConfig{&completedCfg}
}

// New returns a new instance of KubeDBWebhookServer from the given config.
func (c completedConfig) New() (*KubeDBWebhookServer, error) {
	var tlsOpts []func(*tls.Config)
	ctrl.SetLogger(klog.NewKlogr())
	syncPeriod := 1 * time.Hour

	// if the enable-http2 flag is false (the default), http/2 should be disabled
	// due to its vulnerabilities. More specifically, disabling http/2 will
	// prevent from being vulnerable to the HTTP/2 Stream Cancellation and
	// Rapid Reset CVEs. For more information see:
	// - https://github.com/advisories/GHSA-qppj-fm5r-hxr3
	// - https://github.com/advisories/GHSA-4374-p667-p6c8
	disableHTTP2 := func(c *tls.Config) {
		setupLog.Info("disabling http/2")
		c.NextProtos = []string{"http/1.1"}
	}

	if !c.WebhookConfig.EnableHTTP2 {
		tlsOpts = append(tlsOpts, disableHTTP2)
	}

	// Create watchers for metrics and webhooks certificates
	var certWatcher *certwatcher.CertWatcher

	// Initial webhook TLS options
	webhookTLSOpts := tlsOpts
	metricsTLSOpts := tlsOpts

	if len(c.WebhookConfig.CertDir) > 0 {
		setupLog.Info("Initializing certificate watcher using provided certificates",
			"cert-dir", c.WebhookConfig.CertDir, "cert-name", core.TLSCertKey, "cert-key", core.TLSPrivateKeyKey)

		var err error
		certWatcher, err = certwatcher.New(
			filepath.Join(c.WebhookConfig.CertDir, core.TLSCertKey),
			filepath.Join(c.WebhookConfig.CertDir, core.TLSPrivateKeyKey),
		)
		if err != nil {
			setupLog.Error(err, "Failed to initialize webhook certificate watcher")
			os.Exit(1)
		}

		webhookTLSOpts = append(webhookTLSOpts, func(config *tls.Config) {
			config.GetCertificate = certWatcher.GetCertificate
		})

		metricsTLSOpts = append(metricsTLSOpts, func(config *tls.Config) {
			config.GetCertificate = certWatcher.GetCertificate
		})
	}

	webhookServer := webhook.NewServer(webhook.Options{
		TLSOpts: webhookTLSOpts,
	})

	// Metrics endpoint is enabled in 'config/default/kustomization.yaml'. The Metrics options configure the server.
	// More info:
	// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.2/pkg/metrics/server
	// - https://book.kubebuilder.io/reference/metrics.html
	metricsServerOptions := metricsserver.Options{
		BindAddress:   c.WebhookConfig.MetricsAddr,
		SecureServing: c.WebhookConfig.SecureMetrics,
		TLSOpts:       metricsTLSOpts,
	}

	if c.WebhookConfig.SecureMetrics {
		// FilterProvider is used to protect the metrics endpoint with authn/authz.
		// These configurations ensure that only authorized users and service accounts
		// can access the metrics endpoint. The RBAC are configured in 'config/rbac/kustomization.yaml'. More info:
		// https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.2/pkg/metrics/filters#WithAuthenticationAndAuthorization
		metricsServerOptions.FilterProvider = filters.WithAuthenticationAndAuthorization
	}

	mgr, err := ctrl.NewManager(c.WebhookConfig.ClientConfig, ctrl.Options{
		Scheme:                 Scheme,
		Metrics:                metricsServerOptions,
		WebhookServer:          webhookServer,
		HealthProbeBindAddress: c.WebhookConfig.ProbeAddr,
		LeaderElection:         c.WebhookConfig.EnableLeaderElection,
		LeaderElectionID:       "5b87adeb.webhook.kubedb.com",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
		//ClientDisableCacheFor: []client.Object{
		//	&core.Pod{},
		//},
		NewClient: cu.NewClient,
		Cache: cache.Options{
			SyncPeriod: &syncPeriod, // Default SyncPeriod is 10 Hours
		},
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if certWatcher != nil {
		setupLog.Info("Adding certificate watcher to manager")
		if err := mgr.Add(certWatcher); err != nil {
			setupLog.Error(err, "unable to add certificate watcher to manager")
			os.Exit(1)
		}
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	s := &KubeDBWebhookServer{
		Manager: mgr,
	}
	return s, nil
}
