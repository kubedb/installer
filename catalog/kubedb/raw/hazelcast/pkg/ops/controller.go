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

	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	api_listers "kubedb.dev/apimachinery/client/listers/kubedb/v1alpha2"
	ops_listers "kubedb.dev/apimachinery/client/listers/ops/v1alpha1"
	amc "kubedb.dev/apimachinery/pkg/controller"
	"kubedb.dev/apimachinery/pkg/eventer"
	"kubedb.dev/apimachinery/pkg/lib"

	cm "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/apis/core"
	"kmodules.xyz/client-go/tools/queue"
)

type Controller struct {
	amc.Config
	*amc.Controller
	*amc.OpsRequestController

	// CertManagerClient for cert-manger
	CertManagerClient cm.Interface

	// Hazelcast
	dbQueue    *queue.Worker[any]
	dbInformer cache.SharedIndexInformer
	dbLister   api_listers.HazelcastLister

	// HazelcastOpsRequest
	reqQueue    *queue.Worker[any]
	reqInformer cache.SharedIndexInformer
	reqLister   ops_listers.HazelcastOpsRequestLister

	pemEncodeCert bool
	log           logr.Logger
}

func New(
	opt amc.Config,
	ctrl *amc.Controller,
	certManagerClient cm.Interface,
	pemEncodeCert bool,
	verbosity int,
) *Controller {
	return &Controller{
		Controller:           ctrl,
		Config:               opt,
		CertManagerClient:    certManagerClient,
		OpsRequestController: amc.NewOpsRequestController(ctrl.KBClient, opsapi.ResourceKindHazelcastOpsRequest),
		pemEncodeCert:        pemEncodeCert,
		log:                  lib.NewLogger(verbosity).WithName(dbapi.ResourceSingularHazelcast),
	}
}

func (c *Controller) RunControllers(stopCh <-chan struct{}) {
	c.dbQueue.Run(stopCh)
	c.reqQueue.Run(stopCh)
}

func (c *Controller) pushFailureEvent(db *dbapi.Hazelcast, reason string) {
	c.Recorder.Eventf(
		db,
		core.EventTypeWarning,
		eventer.EventReasonFailedToStart,
		`Fail to be ready database: "%v". Reason: %v`,
		db.Name,
		reason,
	)
}

func (c *Controller) getObjectFromKey(key string) (*opsapi.HazelcastOpsRequest, error) {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return nil, err
	}

	var object opsapi.HazelcastOpsRequest
	err = c.KBClient.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, &object)
	return &object, err
}
