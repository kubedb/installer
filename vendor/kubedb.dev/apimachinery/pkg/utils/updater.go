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

package utils

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"slices"
	"time"

	v1 "kubedb.dev/apimachinery/apis/kubedb/v1"

	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	core_util "kmodules.xyz/client-go/core/v1"
	health "kmodules.xyz/client-go/tools/healthchecker"
	scutil "kubeops.dev/operator-shard-manager/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	ReadinessGateType = "kubedb.com/conversion"
)

func UpdateReadinessGateCondition(ctx context.Context, kc client.Client) error {
	namespace := os.Getenv("POD_NAMESPACE")
	name := os.Getenv("POD_NAME")
	var pod core.Pod
	err := kc.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, &pod)
	if err != nil {
		return err
	}

	foundCondition := false
	for i := range pod.Status.Conditions {
		if pod.Status.Conditions[i].Type == ReadinessGateType {
			pod.Status.Conditions[i].Status = core.ConditionTrue
			foundCondition = true
			break
		}
	}

	if !foundCondition { // Add a new condition if not found
		pod.Status.Conditions = append(pod.Status.Conditions, core.PodCondition{
			Type:   ReadinessGateType,
			Status: core.ConditionTrue,
		})
	}

	err = kc.Status().Update(context.TODO(), &pod)
	if err != nil {
		return err
	}

	klog.Infoln("Successfully updated the readiness gate condition to True")
	return nil
}

func WaitForShardIdUpdate(kc client.Client, shardConfigName string) {
	hostName := os.Getenv("HOSTNAME")
	head, err := scutil.FindHeadOfLineage(kc)
	if err != nil {
		panic(fmt.Sprintf("failed to find the head of the lineage for %v, err: %v", hostName, err))
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	timeout := time.After(5 * time.Minute)
	klog.Infof("Waiting for the shard-id to be updated for %v in shardConfig %v \n", hostName, shardConfigName)
	for {
		select {
		case <-timeout:
			panic("shardConfig flag provided but no shard object is found with that name")
		case <-ticker.C:
			pods, err := scutil.GetPodListsFromShardConfig(kc, *head, shardConfigName)
			if err != nil {
				klog.V(6).Infoln(err.Error())
				continue
			}
			if slices.Contains(pods, hostName) {
				return
			}
		}
	}
}

type Predicator interface {
	GetOwnerObject(obj client.Object) (*unstructured.Unstructured, error)
	GetPredicateFuncsForDatabase() predicate.Funcs
	GetPredicateFuncsForOwnerObjects() predicate.Funcs
	GetArchiverToDatabasesMappingFunc(ctx context.Context, obj client.Object) []reconcile.Request
}
type dbPredicate struct {
	kc            client.Client
	shardConfig   string
	healthChecker *health.HealthChecker
	gvk           schema.GroupVersionKind
}

func NewPredicator(kc client.Client, gvk schema.GroupVersionKind, shardConfig string, healthChecker *health.HealthChecker) Predicator {
	return &dbPredicate{
		kc:            kc,
		shardConfig:   shardConfig,
		healthChecker: healthChecker,
		gvk:           gvk,
	}
}

func (p *dbPredicate) GetOwnerObject(obj client.Object) (*unstructured.Unstructured, error) {
	ctrl := metav1.GetControllerOf(obj)
	if ctrl == nil {
		return nil, nil
	}

	ok, err := core_util.IsOwnerOfGroupKind(ctrl, p.gvk.Group, p.gvk.Kind)
	if err != nil || !ok {
		return nil, errors.Wrap(err, fmt.Sprintf("%v/%v is not controlled by %v ", obj.GetNamespace(), obj.GetName(), p.gvk))
	}

	var un unstructured.Unstructured
	un.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   p.gvk.Group,
		Version: p.gvk.Version,
		Kind:    p.gvk.Kind,
	})

	err = p.kc.Get(context.TODO(), types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      ctrl.Name,
	}, &un)
	if err != nil {
		return nil, err
	}

	return &un, err
}

func (p *dbPredicate) GetPredicateFuncsForDatabase() predicate.Funcs {
	return predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			obj := e.Object
			rq := scutil.ShouldEnqueueObjectForShard(p.kc, p.shardConfig, obj.GetLabels())
			if !rq && p.healthChecker != nil {
				p.healthChecker.Stop(obj.GetNamespace() + "/" + obj.GetName())
			}
			return rq
		},

		UpdateFunc: func(e event.UpdateEvent) bool {
			newObj := e.ObjectNew
			rq := scutil.ShouldEnqueueObjectForShard(p.kc, p.shardConfig, newObj.GetLabels())
			if !rq && p.healthChecker != nil {
				p.healthChecker.Stop(newObj.GetNamespace() + "/" + newObj.GetName())
			}
			return rq
		},

		DeleteFunc: func(e event.DeleteEvent) bool {
			obj := e.Object
			rq := scutil.ShouldEnqueueObjectForShard(p.kc, p.shardConfig, obj.GetLabels())
			if !rq && p.healthChecker != nil {
				p.healthChecker.Stop(obj.GetNamespace() + "/" + obj.GetName())
			}
			return rq
		},
	}
}

func (p *dbPredicate) GetPredicateFuncsForOwnerObjects() predicate.Funcs {
	return predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			dbObj, err := p.GetOwnerObject(e.Object)
			if err != nil && !kerr.IsNotFound(err) {
				klog.Errorln(err)
				return false
			}
			if dbObj == nil {
				return false
			}
			rq := scutil.ShouldEnqueueObjectForShard(p.kc, p.shardConfig, dbObj.GetLabels())
			if !rq && p.healthChecker != nil {
				p.healthChecker.Stop(dbObj.GetNamespace() + "/" + dbObj.GetName())
			}
			return rq
		},

		UpdateFunc: func(e event.UpdateEvent) bool {
			dbObj, err := p.GetOwnerObject(e.ObjectNew)
			if err != nil && !kerr.IsNotFound(err) {
				klog.Errorln(err)
				return false
			}
			if dbObj == nil {
				return false
			}
			rq := scutil.ShouldEnqueueObjectForShard(p.kc, p.shardConfig, dbObj.GetLabels())
			if !rq && p.healthChecker != nil {
				p.healthChecker.Stop(dbObj.GetNamespace() + "/" + dbObj.GetName())
			}
			return rq
		},

		DeleteFunc: func(e event.DeleteEvent) bool {
			dbObj, err := p.GetOwnerObject(e.Object)
			if err != nil && !kerr.IsNotFound(err) {
				klog.Errorln(err)
				return false
			}
			if dbObj == nil {
				return false
			}
			rq := scutil.ShouldEnqueueObjectForShard(p.kc, p.shardConfig, dbObj.GetLabels())
			if !rq && p.healthChecker != nil {
				p.healthChecker.Stop(dbObj.GetNamespace() + "/" + dbObj.GetName())
			}
			return rq
		},
	}
}

func (p *dbPredicate) GetArchiverToDatabasesMappingFunc(ctx context.Context, obj client.Object) (matched []reconcile.Request) {
	archiverNS, archiverName := obj.GetNamespace(), obj.GetName()
	consumers, err := getAllowedConsumers(obj)
	if err != nil {
		klog.Warningf("failed to get databases field as consumer for archiver: %s/%s. Reason: %v", archiverNS, archiverName, err)
		return
	}

	namespaceAllowlist, err := getAllowedNamespaceList(ctx, p.kc, consumers)
	if err != nil {
		klog.Warningf("failed to get allowed namespace list for archiver: %s/%s. Reason: %v", archiverNS, archiverName, err)
		return
	}

	dbs, err := p.listDatabasesForArchiver(ctx, consumers)
	if err != nil {
		klog.Warningf("failed to list dbs for archiver: %s/%s. Reason: %v", archiverNS, archiverName, err)
		return
	}

	for _, db := range dbs.Items {
		dbNS, dbName := db.GetNamespace(), db.GetName()
		if !isDatabaseNamespaceAllowed(dbNS, archiverNS, *consumers.Namespaces.From, namespaceAllowlist) {
			continue
		}

		key := dbNS + "/" + dbName
		if scutil.ShouldEnqueueObjectForShard(p.kc, p.shardConfig, db.GetLabels()) {
			matched = append(matched, reconcile.Request{
				NamespacedName: types.NamespacedName{Namespace: dbNS, Name: dbName},
			})
		} else if p.healthChecker != nil {
			p.healthChecker.Stop(key)
		}
	}
	return
}

func getAllowedConsumers(obj client.Object) (*v1.AllowedConsumers, error) {
	v := reflect.ValueOf(obj).Elem() // get struct value
	spec := v.FieldByName("Spec")
	if !spec.IsValid() {
		return nil, fmt.Errorf("failed to get databases field from archiver")
	}
	databases := spec.FieldByName("Databases")
	if !databases.IsValid() {
		return nil, fmt.Errorf("failed to get databases field from archiver")
	}
	if databases.IsNil() {
		return nil, fmt.Errorf("databases field is nil ")
	}
	return databases.Interface().(*v1.AllowedConsumers), nil
}

func getAllowedNamespaceList(ctx context.Context, kc client.Client, consumers *v1.AllowedConsumers) (map[string]struct{}, error) {
	if *consumers.Namespaces.From != v1.NamespacesFromSelector {
		return nil, nil
	}
	nsSelector, err := metav1.LabelSelectorAsSelector(consumers.Namespaces.Selector)
	if err != nil {
		return nil, fmt.Errorf("failed to converting namespace selector. Reason: %v", err)
	}

	nsList := &core.NamespaceList{}
	err = kc.List(ctx, nsList, client.MatchingLabelsSelector{Selector: nsSelector})
	if err != nil {
		return nil, fmt.Errorf("failed to listing namespaces. Reason: %v", err)
	}

	allowlist := make(map[string]struct{}, len(nsList.Items))
	for _, ns := range nsList.Items {
		allowlist[ns.Name] = struct{}{}
	}

	return allowlist, nil
}

func (p *dbPredicate) listDatabasesForArchiver(ctx context.Context, consumers *v1.AllowedConsumers) (*unstructured.UnstructuredList, error) {
	dbSelector, err := metav1.LabelSelectorAsSelector(consumers.Selector)
	if err != nil {
		return nil, fmt.Errorf("failed to converting namespace selector. Reason: %v", err)
	}

	dbs, err := listByGVK(ctx, p.kc, p.gvk, []client.ListOption{
		client.MatchingLabelsSelector{Selector: dbSelector},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to listing databases. Reason: %v", err)
	}
	return dbs, nil
}

func isDatabaseNamespaceAllowed(dbNamespace, archiverNamespace string, from v1.FromNamespaces, namespaceAllowlist map[string]struct{}) bool {
	if namespaceAllowlist != nil {
		_, ok := namespaceAllowlist[dbNamespace]
		return ok
	}
	return from != v1.NamespacesFromSame || dbNamespace == archiverNamespace
}

func listByGVK(ctx context.Context, kc client.Client, gvk schema.GroupVersionKind, opts []client.ListOption) (*unstructured.UnstructuredList, error) {
	list := &unstructured.UnstructuredList{}
	list.SetGroupVersionKind(gvk)
	if err := kc.List(ctx, list, opts...); err != nil {
		return nil, err
	}

	return list, nil
}
