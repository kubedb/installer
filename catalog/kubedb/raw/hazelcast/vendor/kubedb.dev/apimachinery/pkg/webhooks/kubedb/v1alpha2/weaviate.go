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

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	olddbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

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

// SetupWeaviateWebhookWithManager registers the webhook for Weaviate in the manager.
func SetupWeaviateWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Weaviate{}).
		WithValidator(&WeaviateCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&WeaviateCustomWebhook{mgr.GetClient()}).
		Complete()
}

type WeaviateCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &WeaviateCustomWebhook{}

// log is for logging in this package.
var weaviatelog = logf.Log.WithName("weaviate-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *WeaviateCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Weaviate)
	if !ok {
		return fmt.Errorf("expected an Weaviate object but got %T", obj)
	}

	weaviatelog.Info("default", "name", db.GetName())
	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &WeaviateCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *WeaviateCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Weaviate)
	if !ok {
		return nil, fmt.Errorf("expected an Weaviate object but got %T", obj)
	}
	weaviatelog.Info("validate create", "name", db.GetName())
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *WeaviateCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Weaviate)
	if !ok {
		return nil, fmt.Errorf("expected an Weaviate object but got %T", newObj)
	}

	weaviatelog.Info("validate update", "name", db.GetName())
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *WeaviateCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Weaviate)
	if !ok {
		return nil, fmt.Errorf("expected an Weaviate object but got %T", obj)
	}
	weaviatelog.Info("validate delete", "name", db.GetName())

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deletionPolicy"),
			db.GetName(),
			"Can not delete as deletionPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "weaviate.kubedb.com", Kind: "Weaviate"}, db.GetName(), allErr)
	}
	return nil, nil
}

func (w *WeaviateCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Weaviate) error {
	var allErr field.ErrorList
	// number of replicas can not be 0 or less
	if db.Spec.Replicas != nil && *db.Spec.Replicas <= 0 {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
			db.GetName(),
			"number of replicas can not be 0 or less"))
	}

	if db.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.GetName(),
			"spec.version' is missing"))
	} else {
		err := w.ValidateVersion(db)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				db.GetName(),
				err.Error()))
		}
	}

	err := w.validateVolumes(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.GetName(),
			err.Error()))
	}

	err = w.validateVolumesMountPaths(&db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumeMounts"),
			db.GetName(),
			err.Error()))
	}

	if db.Spec.StorageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
			db.GetName(),
			"StorageType can not be empty"))
	} else {
		if db.Spec.StorageType != olddbapi.StorageTypeDurable && db.Spec.StorageType != olddbapi.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				db.GetName(),
				"StorageType should be either durable or ephemeral"))
		}
	}

	if len(allErr) == 0 {
		return nil
	}
	return apierrors.NewInvalid(schema.GroupKind{Group: "weaviate.kubedb.com", Kind: "Weaviate"}, db.GetName(), allErr)
}

func (w *WeaviateCustomWebhook) ValidateVersion(db *olddbapi.Weaviate) error {
	wvVersion := catalog.WeaviateVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: db.Spec.Version}, &wvVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

var weaviateReservedVolumes = []string{
	kubedb.WeaviateVolumeData,
}

func (w *WeaviateCustomWebhook) validateVolumes(db *olddbapi.Weaviate) error {
	if db.Spec.PodTemplate.Spec.Volumes == nil {
		return nil
	}
	rsv := make([]string, len(weaviateReservedVolumes))
	copy(rsv, weaviateReservedVolumes)

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

var weaviateReservedVolumeMountPaths = []string{
	kubedb.WeaviateDataDir,
}

func (w *WeaviateCustomWebhook) validateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range weaviateReservedVolumeMountPaths {
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

	return nil
}
