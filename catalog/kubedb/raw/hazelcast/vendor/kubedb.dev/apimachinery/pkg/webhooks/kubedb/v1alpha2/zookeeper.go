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

package v1alpha2

import (
	"context"
	"errors"
	"fmt"

	"kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	olddbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// SetupZooKeeperWebhookWithManager registers the webhook for ZooKeeper in the manager.
func SetupZooKeeperWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.ZooKeeper{}).
		WithValidator(&ZooKeeperCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&ZooKeeperCustomWebhook{mgr.GetClient()}).
		Complete()
}

// log is for logging in this package.
var zookeeperlog = logf.Log.WithName("zookeeper-resource")

//+kubebuilder:webhook:path=/mutate-zookeeper-kubedb-com-v1alpha1-zookeeper,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=zookeepers,verbs=create;update,versions=v1alpha1,name=mzookeeper.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type ZooKeeperCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &ZooKeeperCustomWebhook{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *ZooKeeperCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.ZooKeeper)
	if !ok {
		return fmt.Errorf("expected an ZooKeeper object but got %T", obj)
	}
	zookeeperlog.Info("default", "name", db.Name)
	db.SetDefaults(w.DefaultClient)
	return nil
}

//+kubebuilder:webhook:path=/validate-zookeeper-kubedb-com-v1alpha1-zookeeper,mutating=false,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=zookeepers,verbs=create;update,versions=v1alpha1,name=vzookeeper.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.CustomValidator = &ZooKeeperCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *ZooKeeperCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.ZooKeeper)
	if !ok {
		return nil, fmt.Errorf("expected an ZooKeeper object but got %T", obj)
	}
	zookeeperlog.Info("validate create", "name", db.Name)
	return w.ValidateCreateOrUpdate(db)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *ZooKeeperCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.ZooKeeper)
	if !ok {
		return nil, fmt.Errorf("expected an ZooKeeper object but got %T", newObj)
	}
	zookeeperlog.Info("validate update", "name", db.Name)
	return w.ValidateCreateOrUpdate(db)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *ZooKeeperCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.ZooKeeper)
	if !ok {
		return nil, fmt.Errorf("expected an ZooKeeper object but got %T", obj)
	}
	zookeeperlog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("teminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "zookeeper.kubedb.com", Kind: "ZooKeeper"}, db.Name, allErr)
	}
	return nil, nil
}

func (w *ZooKeeperCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.ZooKeeper) (admission.Warnings, error) {
	var allErr field.ErrorList
	if db.Spec.Replicas != nil && *db.Spec.Replicas == 2 {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
			db.Name,
			"zookeeper ensemble should have 3 or more replicas"))
	}

	err := w.validateVersion(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			err.Error()))
	}

	err = w.validateVolumes(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = w.validateVolumesMountPaths(&db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumeMounts"),
			db.Name,
			err.Error()))
	}

	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "zookeeper.kubedb.com", Kind: "ZooKeeper"}, db.Name, allErr)
}

func (w *ZooKeeperCustomWebhook) validateVersion(db *olddbapi.ZooKeeper) error {
	zkVersion := v1alpha1.ZooKeeperVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: db.Spec.Version}, &zkVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

var zookeeperReservedVolumes = []string{
	kubedb.ZooKeeperDataVolumeName,
	kubedb.ZooKeeperScriptVolumeName,
	kubedb.ZooKeeperConfigVolumeName,
}

func (w *ZooKeeperCustomWebhook) validateVolumes(db *olddbapi.ZooKeeper) error {
	if db.Spec.PodTemplate.Spec.Volumes == nil {
		return nil
	}
	rsv := make([]string, len(zookeeperReservedVolumes))
	copy(rsv, zookeeperReservedVolumes)

	volumes := db.Spec.PodTemplate.Spec.Volumes
	for _, rv := range rsv {
		for _, ugv := range volumes {
			if ugv.Name == rv {
				return errors.New("Cannot use a reserve volume name: " + rv)
			}
		}
	}
	return nil
}

var zookeeperReservedVolumeMountPaths = []string{
	kubedb.ZooKeeperScriptVolumePath,
	kubedb.ZooKeeperConfigVolumePath,
	kubedb.ZooKeeperDataVolumePath,
}

func (w *ZooKeeperCustomWebhook) validateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	var containerList []core.Container
	if podTemplate.Spec.Containers != nil {
		containerList = append(containerList, podTemplate.Spec.Containers...)
	}
	if podTemplate.Spec.InitContainers != nil {
		containerList = append(containerList, podTemplate.Spec.InitContainers...)
	}

	for _, rvmp := range zookeeperReservedVolumeMountPaths {
		for i := range containerList {
			mountPathList := containerList[i].VolumeMounts
			for j := range mountPathList {
				if mountPathList[j].MountPath == rvmp {
					return errors.New("Can't use a reserve volume mount path name: " + rvmp)
				}
			}
		}
	}
	return nil
}
