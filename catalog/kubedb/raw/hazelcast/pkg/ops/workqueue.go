/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hazelcast

import (
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"kubedb.dev/apimachinery/pkg/lib"

	cm_api "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/tools/queue"
)

func (c *Controller) Init() {
	c.dbInformer = c.KubedbInformerFactory.Kubedb().V1alpha2().Hazelcasts().Informer()
	c.dbQueue = queue.New[any]("Hazelcast", c.MaxNumRequeues, c.NumThreads, c.manageHazelcastEvent)
	c.dbLister = c.KubedbInformerFactory.Kubedb().V1alpha2().Hazelcasts().Lister()
	_, _ = c.dbInformer.AddEventHandler(queue.NewReconcilableHandler(c.dbQueue.GetQueue(), c.RestrictToNamespace))

	// initialize HazelcastOpsRequest watchers
	c.reqInformer = c.KubedbInformerFactory.Ops().V1alpha1().HazelcastOpsRequests().Informer()
	c.reqQueue = queue.New[any]("HazelcastOpsRequest", c.MaxNumRequeues, c.NumThreads, c.runHazelcastOpsRequest)
	c.reqLister = c.KubedbInformerFactory.Ops().V1alpha1().HazelcastOpsRequests().Lister()
	_, _ = c.reqInformer.AddEventHandler(queue.NewSpecStatusChangeHandler(c.reqQueue.GetQueue(), c.RestrictToNamespace))
}

func (c *Controller) NewSecretWatcher() cache.ResourceEventHandler {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			if secret, ok := obj.(*core.Secret); ok {
				if key, _ := lib.DBForSecret(c.CertManagerClient.CertmanagerV1(), dbapi.ResourceKindHazelcast, secret); key != "" {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		UpdateFunc: func(oldObj any, newObj any) {
			if secret, ok := newObj.(*core.Secret); ok {
				if key, _ := lib.DBForSecret(c.CertManagerClient.CertmanagerV1(), dbapi.ResourceKindHazelcast, secret); key != "" {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		DeleteFunc: func(obj any) {
		},
	}
}

func (c *Controller) NewServiceWatcher() cache.ResourceEventHandler {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			if svc, ok := obj.(*core.Service); ok {
				if key := lib.DBForService(dbapi.ResourceKindHazelcast, svc); key != "" {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		UpdateFunc: func(oldObj any, newObj any) {
			if svc, ok := newObj.(*core.Service); ok {
				if key := lib.DBForService(dbapi.ResourceKindHazelcast, svc); key != "" {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		DeleteFunc: func(obj any) {
			if svc, ok := obj.(*core.Service); ok {
				if key := lib.DBForService(dbapi.ResourceKindHazelcast, svc); key != "" {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
	}
}

func (c *Controller) NewIssuerWatcher() cache.ResourceEventHandler {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			if issuer, ok := obj.(*cm_api.Issuer); ok {
				keys, err := lib.DBsForIssuer(c.DynamicClient, dbapi.SchemeGroupVersion.WithResource(dbapi.ResourcePluralHazelcast), issuer)
				if err != nil {
					klog.Warningln(err)
				}
				for _, key := range keys {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		UpdateFunc: func(oldObj any, newObj any) {
			if issuer, ok := newObj.(*cm_api.Issuer); ok {
				keys, err := lib.DBsForIssuer(c.DynamicClient, dbapi.SchemeGroupVersion.WithResource(dbapi.ResourcePluralHazelcast), issuer)
				if err != nil {
					klog.Warningln(err)
				}
				for _, key := range keys {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		DeleteFunc: func(obj any) {
		},
	}
}

func (c *Controller) NewClusterIssuerWatcher() cache.ResourceEventHandler {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj any) {
			if issuer, ok := obj.(*cm_api.ClusterIssuer); ok {
				keys, err := lib.DBsForClusterIssuer(c.DynamicClient, dbapi.SchemeGroupVersion.WithResource(dbapi.ResourcePluralHazelcast), issuer)
				if err != nil {
					klog.Warningln(err)
				}
				for _, key := range keys {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		UpdateFunc: func(oldObj any, newObj any) {
			if issuer, ok := newObj.(*cm_api.ClusterIssuer); ok {
				keys, err := lib.DBsForClusterIssuer(c.DynamicClient, dbapi.SchemeGroupVersion.WithResource(dbapi.ResourcePluralHazelcast), issuer)
				if err != nil {
					klog.Warningln(err)
				}
				for _, key := range keys {
					queue.Enqueue(c.dbQueue.GetQueue(), key)
				}
			}
		},
		DeleteFunc: func(obj any) {
		},
	}
}
