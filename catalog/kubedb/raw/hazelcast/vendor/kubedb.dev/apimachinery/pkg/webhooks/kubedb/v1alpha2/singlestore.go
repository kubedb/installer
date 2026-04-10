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
	"fmt"

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	olddbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	"github.com/pkg/errors"
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

// SetupSinglestoreWebhookWithManager registers the webhook for Singlestore in the manager.
func SetupSinglestoreWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Singlestore{}).
		WithValidator(&SinglestoreCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&SinglestoreCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-singlestore-kubedb-com-v1alpha1-singlestore,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=singlestores,verbs=create;update,versions=v1alpha1,name=msinglestore.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type SinglestoreCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &SinglestoreCustomWebhook{}

// log is for logging in this package.
var singlestorelog = logf.Log.WithName("singlestore-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *SinglestoreCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Singlestore)
	if !ok {
		return fmt.Errorf("expected an Singlestore object but got %T", obj)
	}

	singlestorelog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &SinglestoreCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *SinglestoreCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Singlestore)
	if !ok {
		return nil, fmt.Errorf("expected an Singlestore object but got %T", obj)
	}

	singlestorelog.Info("validate create", "name", db.Name)
	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Singlestore"}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *SinglestoreCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Singlestore)
	if !ok {
		return nil, fmt.Errorf("expected an Singlestore object but got %T", newObj)
	}
	singlestorelog.Info("validate update", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Singlestore"}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *SinglestoreCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Singlestore)
	if !ok {
		return nil, fmt.Errorf("expected an Singlestore object but got %T", obj)
	}
	singlestorelog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Singlestore"}, db.Name, allErr)
	}
	return nil, nil
}

func (w *SinglestoreCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Singlestore) field.ErrorList {
	var allErr field.ErrorList

	if db.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			"spec.version' is missing"))
	} else {
		err := w.sdbValidateVersion(db)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				db.Name,
				err.Error()))
		}
	}

	if db.Spec.Topology == nil {
		err := sdbValidateVolumes(db.Spec.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = sdbValidateVolumesMountPaths(db.Spec.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("containers"),
				db.Name,
				err.Error()))
		}

	} else {
		if db.Spec.Topology.Aggregator == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("aggregator"),
				db.Name,
				".spec.topology.aggregator can't be empty in cluster mode"))
		}
		if db.Spec.Topology.Leaf == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("leaf"),
				db.Name,
				".spec.topology.leaf can't be empty in cluster mode"))
		}

		if db.Spec.Topology.Aggregator.Replicas == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("aggregator").Child("replicas"),
				db.Name,
				"doesn't support spec.topology.aggregator.replicas is set"))
		}
		if db.Spec.Topology.Leaf.Replicas == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("leaf").Child("replicas"),
				db.Name,
				"doesn't support spec.topology.leaf.replicas is set"))
		}

		if *db.Spec.Topology.Aggregator.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("aggregator").Child("replicas"),
				db.Name,
				"number of replicas can not be less be 0 or less"))
		}

		if *db.Spec.Topology.Leaf.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("leaf").Child("replicas"),
				db.Name,
				"number of replicas can not be 0 or less"))
		}

		err := sdbValidateVolumes(db.Spec.Topology.Aggregator.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("aggregator").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = sdbValidateVolumes(db.Spec.Topology.Leaf.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("leaf").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}

		err = sdbValidateVolumesMountPaths(db.Spec.Topology.Aggregator.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("aggregator").Child("podTemplate").Child("spec").Child("containers"),
				db.Name,
				err.Error()))
		}
		err = sdbValidateVolumesMountPaths(db.Spec.Topology.Leaf.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("leaf").Child("podTemplate").Child("spec").Child("containers"),
				db.Name,
				err.Error()))
		}
	}

	if db.Spec.StorageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
			db.Name,
			"StorageType can not be empty"))
	} else {
		if db.Spec.StorageType != olddbapi.StorageTypeDurable && db.Spec.StorageType != olddbapi.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				db.Name,
				"StorageType should be either durable or ephemeral"))
		}
	}

	if len(allErr) == 0 {
		return nil
	}
	return allErr
}

// reserved volume and volumes mounts for singlestore
var sdbReservedVolumes = []string{
	kubedb.SinglestoreVolumeNameUserInitScript,
	kubedb.SinglestoreVolumeNameCustomConfig,
	kubedb.SinglestoreVolmeNameInitScript,
	kubedb.SinglestoreVolumeNameData,
	kubedb.SinglestoreVolumeNameTLS,
}

var sdbReservedVolumesMountPaths = []string{
	kubedb.SinglestoreVolumeMountPathData,
	kubedb.SinglestoreVolumeMountPathInitScript,
	kubedb.SinglestoreVolumeMountPathCustomConfig,
	kubedb.SinglestoreVolumeMountPathUserInitScript,
	kubedb.SinglestoreVolumeMountPathTLS,
}

func (w *SinglestoreCustomWebhook) sdbValidateVersion(db *olddbapi.Singlestore) error {
	var sdbVersion catalog.SinglestoreVersion
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &sdbVersion)
	if err != nil {
		return errors.New("version not supported")
	}

	return nil
}

func sdbValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range sdbReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserve volume name: " + rv)
			}
		}
	}

	return nil
}

func sdbValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range sdbReservedVolumesMountPaths {
		containerList := podTemplate.Spec.Containers
		for i := range containerList {
			mountPathList := containerList[i].VolumeMounts
			for j := range mountPathList {
				if mountPathList[j].MountPath == rvmp {
					return errors.New("Can't use a reserve volume mount path name: " + rvmp)
				}
			}
		}
	}

	if podTemplate.Spec.InitContainers == nil {
		return nil
	}

	for _, rvmp := range sdbReservedVolumesMountPaths {
		containerList := podTemplate.Spec.InitContainers
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
