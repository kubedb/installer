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
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"kubedb.dev/apimachinery/apis/config/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	olddbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"kubedb.dev/apimachinery/client/clientset/versioned"
	amc "kubedb.dev/apimachinery/pkg/controller"
	"kubedb.dev/apimachinery/pkg/utils"
	hazelcastcontrollers "kubedb.dev/hazelcast/pkg/controller"
	"kubedb.dev/hazelcast/pkg/server"

	"github.com/pkg/errors"
	pcm "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	"github.com/spf13/pflag"
	auditlib "go.bytebuilders.dev/audit/lib"
	proxyserver "go.bytebuilders.dev/license-proxyserver/apis/proxyserver/v1alpha1"
	licenseapi "go.bytebuilders.dev/license-verifier/apis/licenses/v1alpha1"
	license "go.bytebuilders.dev/license-verifier/kubernetes"
	v "gomodules.xyz/x/version"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/apiextensions"
	cu "kmodules.xyz/client-go/client"
	clustermeta "kmodules.xyz/client-go/cluster"
	"kmodules.xyz/client-go/discovery"
	"kmodules.xyz/client-go/tools/clientcmd"
	health "kmodules.xyz/client-go/tools/healthchecker"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

var setupLog = log.Log.WithName("setup")

type OperatorOptions struct {
	MasterURL            string
	KubeconfigPath       string
	LicenseFile          string
	QPS                  float64
	Burst                int
	ResyncPeriod         time.Duration
	MaxNumRequeues       int
	NumThreads           int
	shardConfig          string
	NetworkPolicyEnabled bool
	metricsAddr          string
	enableLeaderElection bool
	probeAddr            string
}

func NewOperatorOptions() *OperatorOptions {
	return &OperatorOptions{
		ResyncPeriod:   10 * time.Minute,
		MaxNumRequeues: 5,
		NumThreads:     2,
		// ref: https://github.com/kubernetes/ingress-nginx/blob/e4d53786e771cc6bdd55f180674b79f5b692e552/pkg/ingress/controller/launch.go#L252-L259
		// High enough QPS to fit all expected use cases. QPS=0 is not set here, because client code is overriding it.
		QPS: 1e6,
		// High enough Burst to fit all expected use cases. Burst=0 is not set here, because client code is overriding it.
		Burst:                1e6,
		metricsAddr:          ":8080",
		enableLeaderElection: false,
		probeAddr:            ":8081",
	}
}

func (s *OperatorOptions) AddGoFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.MasterURL, "master", s.MasterURL, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	fs.StringVar(&s.KubeconfigPath, "kubeconfig", s.KubeconfigPath, "Path to kubeconfig file with authorization information (the master location is set by the master flag).")

	fs.StringVar(&s.LicenseFile, "license-file", s.LicenseFile, "Path to license file")
	fs.Float64Var(&s.QPS, "qps", s.QPS, "The maximum QPS to the master from this client")
	fs.IntVar(&s.Burst, "burst", s.Burst, "The maximum burst for throttle")
	fs.DurationVar(&s.ResyncPeriod, "resync-period", s.ResyncPeriod, "If non-zero, will re-list this often. Otherwise, re-list will be delayed aslong as possible (until the upstream source closes the watch or times out.")
	fs.BoolVar(&s.NetworkPolicyEnabled, "enable-network-policy", s.NetworkPolicyEnabled, "Controls the network policy creation")
	fs.StringVar(&s.metricsAddr, "metrics-bind-address", s.metricsAddr, "The address the metric endpoint binds to.")
	fs.StringVar(&s.probeAddr, "health-probe-bind-address", s.probeAddr, "The address the probe endpoint binds to.")
	fs.BoolVar(&s.enableLeaderElection, "leader-elect", s.enableLeaderElection,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	fs.StringVar(&s.shardConfig, "shard-config", s.shardConfig, "Shard configuration that will this operator use")
	fs.IntVar(&s.NumThreads, "max-concurrent-reconciles", s.NumThreads, "The maximum number of concurrent reconciles which can be run")
}

func (s *OperatorOptions) AddFlags(fs *pflag.FlagSet) {
	pfs := flag.NewFlagSet("extra-flags", flag.ExitOnError)
	s.AddGoFlags(pfs)
	fs.AddGoFlagSet(pfs)
}

func (s *OperatorOptions) Validate() []error {
	return nil
}

func (s *OperatorOptions) Complete() error {
	return nil
}

func (s OperatorOptions) Run(ctx context.Context) error {
	klog.Infof("Starting binary version %s+%s ...", v.Version.Version, v.Version.CommitHash)

	log.SetLogger(klog.NewKlogr())

	cfg, err := clientcmd.BuildConfigFromFlags(s.MasterURL, s.KubeconfigPath)
	if err != nil {
		klog.Fatalf("Could not get Kubernetes config: %s", err)
	}

	// Fixes https://github.com/Azure/AKS/issues/522
	clientcmd.Fix(cfg)

	cfg.QPS = float32(s.QPS)
	cfg.Burst = s.Burst

	mgr, err := manager.New(cfg, manager.Options{
		Scheme:  server.Scheme,
		Metrics: metricsserver.Options{BindAddress: ""},
		Cache: cache.Options{
			SyncPeriod: &s.ResyncPeriod,
		},
		HealthProbeBindAddress: s.probeAddr,
		LeaderElection:         s.enableLeaderElection,
		LeaderElectionID:       "8d2935c9.kubedb.com",
		NewClient:              cu.NewClient,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	LicenseProvided := func() bool {
		if s.LicenseFile != "" {
			return true
		}

		ok, _ := discovery.HasGVK(
			kubernetes.NewForConfigOrDie(cfg).Discovery(),
			proxyserver.SchemeGroupVersion.String(),
			proxyserver.ResourceKindLicenseRequest)
		return ok
	}

	// audit event publisher
	// WARNING: https://stackoverflow.com/a/46275411/244009
	var auditor *auditlib.EventPublisher
	var restrictions v1alpha1.LicenseRestrictions
	if LicenseProvided() {
		info, _ := license.MustLicenseEnforcer(cfg, s.LicenseFile).LoadLicense()
		if info.Status != licenseapi.LicenseActive {
			return fmt.Errorf("license status %s", info.Status)
		}
		if !sets.NewString(info.Features...).Has("kubedb-enterprise") {
			return fmt.Errorf("not a valid license for this product")
		}
		if data := strings.TrimSpace(info.FeatureFlags[licenseapi.FeatureRestrictions]); len(data) > 0 {
			err = json.Unmarshal([]byte(data), &restrictions)
			if err != nil {
				return err
			}
		}
		if !info.DisableAnalytics() {
			cmeta, err := clustermeta.ClusterMetadata(mgr.GetAPIReader())
			if err != nil {
				return fmt.Errorf("failed to extract cluster metadata, reason: %v", err)
			}
			nc := auditlib.NewNatsClient(mgr.GetConfig(), cmeta.UID, s.LicenseFile)
			mapper := discovery.NewResourceMapper(mgr.GetRESTMapper())
			fn := auditlib.BillingEventCreator{
				Mapper:          mapper,
				ClusterMetadata: cmeta,
				ClientBilling:   info.EnableClientBilling(),
				PodLister:       mgr.GetClient(),
				PVCLister:       mgr.GetClient(),
				NamespaceLister: mgr.GetClient(),
			}
			auditor = auditlib.NewEventPublisher(nc, mapper, fn.CreateEvent)
			err = auditor.SetupSiteInfoPublisherWithManager(mgr)
			if err != nil {
				return fmt.Errorf("failed to setup site info publisher, reason: %v", err)
			}
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
	if err := apiextensions.NewReconciler(ctx, mgr).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "unable to create controller controller CustomResourceReconciler")
	}

	apiextensions.RegisterSetup(schema.GroupKind{
		Group: kubedb.GroupName,
		Kind:  olddbapi.ResourceKindHazelcast,
	}, func(ctx context.Context, mgr manager.Manager) {
		if s.shardConfig != "" {
			utils.WaitForShardIdUpdate(mgr.GetClient(), s.shardConfig)
		}
		SetupHazelcastControllers(ctx, mgr, auditor, s.NetworkPolicyEnabled, restrictions, s.shardConfig)
	})

	// Start periodic license verification
	//nolint:errcheck
	//go license.VerifyLicensePeriodically(mgr.GetConfig(), s.LicenseFile, ctx.Done())
	setupLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
	return nil
}

func SetupHazelcastControllers(ctx context.Context, mgr manager.Manager, auditor *auditlib.EventPublisher, networkPolicyEnabled bool, restrictions v1alpha1.LicenseRestrictions, shardConfig string) {
	appCatalogClient, err := appcat_cs.NewForConfig(mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "failed to create appCatalog client")
		os.Exit(1)
	}
	k8sClient, err := kubernetes.NewForConfig(mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "failed to create kubernetes client")
		os.Exit(1)
	}
	dbClient, err := versioned.NewForConfig(mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "failed to create kubedb client")
		os.Exit(1)
	}
	dynamicClient, err := dynamic.NewForConfig(mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "failed to create dynamic client")
		os.Exit(1)
	}
	promClient, err := pcm.NewForConfig(mgr.GetConfig())
	if err != nil {
		setupLog.Error(err, "failed to create prometheus client")
		os.Exit(1)
	}
	if err := (&hazelcastcontrollers.HazelcastReconciler{
		Config: &amc.Config{
			ShardConfig: shardConfig,
		},
		Controller: &amc.Controller{
			KBClient:         mgr.GetClient(),
			AppCatalogClient: appCatalogClient,
			DynamicClient:    dynamicClient,
			DBClient:         dbClient,
			Client:           k8sClient,
		},
		Scheme:               mgr.GetScheme(),
		PromClient:           promClient,
		HealthChecker:        health.NewHealthChecker(),
		NetworkPolicyEnabled: networkPolicyEnabled,
		LicenseRestrictions:  restrictions,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Hazelcast")
		os.Exit(1)
	}
	if err := auditor.SetupWithManager(ctx, mgr, &olddbapi.Hazelcast{}); err != nil {
		setupLog.Error(err, "unable to set up auditor", "kind", "Hazelcast")
		os.Exit(1)
	}
}
