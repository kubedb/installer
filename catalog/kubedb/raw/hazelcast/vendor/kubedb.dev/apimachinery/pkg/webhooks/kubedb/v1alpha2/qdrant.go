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

// SetupQdrantWebhookWithManager registers the webhook for Qdrant in the manager.
func SetupQdrantWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Qdrant{}).
		WithValidator(&QdrantCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&QdrantCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-kubedb-com-v1alpha2-qdrant,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=qdrants,verbs=create;update,versions=v1alpha2,name=qdrant.kb.io,admissionReviewVersions=v1

// +kubebuilder:object:generate=false
type QdrantCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &QdrantCustomWebhook{}

// log is for logging in this package.
var qdrantlog = logf.Log.WithName("qdrant-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *QdrantCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Qdrant)
	if !ok {
		return fmt.Errorf("expected a Qdrant object, got a %T", obj)
	}

	qdrantlog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &QdrantCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *QdrantCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Qdrant)
	if !ok {
		return nil, fmt.Errorf("expected a Qdrant object, got a %T", obj)
	}

	qdrantlog.Info("validate create", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindQdrant}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *QdrantCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Qdrant)
	if !ok {
		return nil, fmt.Errorf("expected a Qdrant object, got a %T", newObj)
	}

	qdrantlog.Info("validate update", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindQdrant}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *QdrantCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Qdrant)
	if !ok {
		return nil, fmt.Errorf("expected a Qdrant object, got a %T", obj)
	}

	qdrantlog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindQdrant}, db.Name, allErr)
	}
	return nil, nil
}

func (w *QdrantCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Qdrant) field.ErrorList {
	var allErr field.ErrorList

	err := w.qdrantValidateVersion(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			err.Error()))
	}

	err = qdrantValidateVolumes(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = qdrantValidateVolumesMountPaths(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("containers"),
			db.Name,
			err.Error()))
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

	if db.Spec.TLS != nil && db.Spec.TLS.P2P != nil && *db.Spec.TLS.P2P {
		if db.Spec.Mode == "" || db.Spec.Mode == olddbapi.QdrantStandalone {
			allErr = append(allErr,
				field.Invalid(
					field.NewPath("spec").Child("tls").Child("p2p"),
					db.Spec.TLS.P2P,
					"p2p TLS requires distributed mode",
				),
			)
		}
	}

	if len(allErr) == 0 {
		return nil
	}

	return allErr
}

// reserved volume and volumes mounts for qdrant
var qdrantReservedVolumes = []string{
	kubedb.QdrantDataVolName,
	kubedb.QdrantConfigVolName,
}

var qdrantReservedVolumesMountPaths = []string{
	kubedb.QdrantDataDir,
	kubedb.QdrantConfigDir,
}

func (w *QdrantCustomWebhook) qdrantValidateVersion(db *olddbapi.Qdrant) error {
	var qdrantVersion catalog.QdrantVersion

	return w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &qdrantVersion)
}

func qdrantValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range qdrantReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserved volume name: " + rv)
			}
		}
	}

	return nil
}

func qdrantValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}

	if podTemplate.Spec.Containers != nil {
		// Check container volume mounts
		for _, rvmp := range qdrantReservedVolumesMountPaths {
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
	}

	return nil
}
