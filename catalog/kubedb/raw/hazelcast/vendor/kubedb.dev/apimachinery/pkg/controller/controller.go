/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"time"

	configapi "kubedb.dev/apimachinery/apis/config/v1alpha1"
	cs "kubedb.dev/apimachinery/client/clientset/versioned"
	kubedbinformers "kubedb.dev/apimachinery/client/informers/externalversions"

	cmInformers "github.com/cert-manager/cert-manager/pkg/client/informers/externalversions"
	auditlib "go.bytebuilders.dev/audit/lib"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	core_util "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/client-go/tools/queue"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
	appcat_in "kmodules.xyz/custom-resources/client/informers/externalversions"
	petsetcs "kubeops.dev/petset/client/clientset/versioned"
	psinformer "kubeops.dev/petset/client/informers/externalversions"
	pslister "kubeops.dev/petset/client/listers/apps/v1"
	ocmclient "open-cluster-management.io/api/client/work/clientset/versioned"
	manifestinformers "open-cluster-management.io/api/client/work/informers/externalversions"
	manifestlisters "open-cluster-management.io/api/client/work/listers/work/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	scs "stash.appscode.dev/apimachinery/client/clientset/versioned"
	stashinformer "stash.appscode.dev/apimachinery/client/informers/externalversions"
	lister "stash.appscode.dev/apimachinery/client/listers/stash/v1beta1"
)

type Controller struct {
	ClientConfig *rest.Config
	// KubeBuilder Client
	KBClient client.Client
	// Kubernetes client
	Client kubernetes.Interface
	// KubeDB client
	DBClient cs.Interface
	// PetSet client
	PSClient petsetcs.Interface
	// ManifestWork client
	OCMClient ocmclient.Interface
	// Dynamic client
	DynamicClient dynamic.Interface
	// AppCatalog client
	AppCatalogClient appcat_cs.Interface
	// Cluster topology when the operator started
	ClusterTopology *core_util.Topology
	// Event Recorder
	Recorder record.EventRecorder
	// Audit Event Publisher
	Auditor *auditlib.EventPublisher
}

type Config struct {
	// Informer factory
	KubeInformerFactory        informers.SharedInformerFactory
	KubedbInformerFactory      kubedbinformers.SharedInformerFactory
	AppCatInformerFactory      appcat_in.SharedInformerFactory
	CertManagerInformerFactory cmInformers.SharedInformerFactory
	PetSetInformerFactory      psinformer.SharedInformerFactory
	ManifestInformerFactory    manifestinformers.SharedInformerFactory

	// External tool to initialize the database
	Initializers Initializers

	// Secret
	SecretInformer cache.SharedIndexInformer
	SecretLister   corelisters.SecretLister

	// PetSet Watcher
	PSQueue    *queue.Worker[any]
	PSInformer cache.SharedIndexInformer
	PSLister   pslister.PetSetLister

	MWInformer cache.SharedIndexInformer
	// manifestWorkerLister is able to list/get manifestWork from a shared informer's store
	MWLister manifestlisters.ManifestWorkLister

	// Only watch or reconcile objects in this namespace (usually for license reasons)
	RestrictToNamespace    string
	LicenseRestrictions    configapi.LicenseRestrictions
	ResyncPeriod           time.Duration
	ReadinessProbeInterval time.Duration
	MaxNumRequeues         int
	NumThreads             int
	NetworkPolicyEnabled   bool
	ShardConfig            string
}

type Initializers struct {
	Stash StashInitializer
}

type StashInitializer struct {
	StashClient          scs.Interface
	StashInformerFactory stashinformer.SharedInformerFactory
	// StashInitializer RestoreSession
	RSQueue    *queue.Worker[any]
	RSInformer cache.SharedIndexInformer
	RSLister   lister.RestoreSessionLister

	// StashInitializer RestoreBatch
	RBQueue    *queue.Worker[any]
	RBInformer cache.SharedIndexInformer
	RBLister   lister.RestoreBatchLister
}
