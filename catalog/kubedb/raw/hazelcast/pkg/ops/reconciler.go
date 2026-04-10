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
	"context"
	"errors"
	"os"

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	amc "kubedb.dev/apimachinery/pkg/controller"
	"kubedb.dev/apimachinery/pkg/lib"
	hzapi "kubedb.dev/hazelcast/pkg/controller"

	prom "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	health "kmodules.xyz/client-go/tools/healthchecker"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
	ctrl "sigs.k8s.io/controller-runtime"
)

type reconcileRetries struct {
	reconcileRetries *Retries
}
type reconcile struct {
	*hzapi.ReconcileState
	retries reconcileRetries
	db      *dbapi.Hazelcast
}

func (c *hzOpsReqController) NewReconcileFunc(hz *dbapi.Hazelcast, rcl *hzapi.ReconcileState) func() (bool, error) {
	r := &reconcile{
		ReconcileState: rcl,
		retries:        reconcileRetries{},
		db:             hz,
	}
	r.retries.reconcileRetries = c.newRetries("Reconcile")
	return r.reconcile
}

func (r *reconcile) reconcile() (bool, error) {
	log := lib.NewLogger(3).WithName(dbapi.ResourceSingularHazelcast)
	if r.ReconcileState == nil {
		log.Error(errors.New("reconciler nil"), "reconciler is nil")
		return false, nil
	}

	err := r.EnsureServices()
	if err != nil {
		log.Error(err, "failed to ensure services")
		return r.retries.reconcileRetries.Wait(), err
	}

	err = r.EnsureSecrets()
	if err != nil {
		log.Error(err, "failed to ensure secrets")
		return r.retries.reconcileRetries.Wait(), err
	}

	err = r.EnsureStatefulSet()
	if err != nil {
		return r.retries.reconcileRetries.Wait(), err
	}
	return false, nil
}

func (c *hzOpsReqController) NewHazelcastReconcile(db *dbapi.Hazelcast) (*hzapi.ReconcileState, error) {
	appCatalogClient, err := appcat_cs.NewForConfig(c.ClientConfig)
	if err != nil {
		c.log.Error(err, "failed to create app catalog client")
		return nil, err
	}
	promClient, err := prom.NewForConfig(c.ClientConfig)
	if err != nil {
		c.log.Error(err, "failed to create prometheus client")
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(c.ClientConfig)
	if err != nil {
		c.log.Error(err, "failed to create dynamic client")
		os.Exit(1)
	}

	version := &catalog.HazelcastVersion{}
	err = c.KBClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, version)
	if err != nil {
		return nil, err
	}

	rs := &hzapi.ReconcileState{
		HazelcastReconciler: &hzapi.HazelcastReconciler{
			Controller: &amc.Controller{
				ClientConfig:     nil,
				KBClient:         c.KBClient,
				Client:           c.Client,
				DBClient:         nil,
				DynamicClient:    dynamicClient,
				AppCatalogClient: appCatalogClient,
			},
			Scheme:        runtime.NewScheme(),
			PromClient:    promClient,
			HealthChecker: health.NewHealthChecker(),
		},
	}
	rs.SetLoggerWithReq(ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      db.Name,
			Namespace: db.Namespace,
		},
	})
	rs.WithContext(context.TODO())
	rs.WithHazelcast(db)
	rs.WithVersion(version)
	return rs, nil
}
