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

// log is for logging in this package.
var hazelcastlog = logf.Log.WithName("hazelcast-resource")

// SetupHazelcastWebhookWithManager registers the webhook for Hazelcast in the manager.
func SetupHazelcastWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Hazelcast{}).
		WithValidator(&HazelcastCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&HazelcastCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-hazelcast-kubedb-com-v1alpha1-hazelcast,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=hazelcasts,verbs=create;update,versions=v1alpha1,name=mhazelcast.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type HazelcastCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &HazelcastCustomWebhook{}

func (w *HazelcastCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Hazelcast)
	if !ok {
		return fmt.Errorf("expected an Hazelcast object but got %T", obj)
	}

	hazelcastlog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var hazelcastReservedVolumes = []string{
	kubedb.HazelcastConfigVolume,
	kubedb.HazelcastDefaultConfigVolume,
	kubedb.HazelcastCustomConfigVolume,
}

var hazelcastReservedVolumeMountPaths = []string{
	kubedb.HazelcastConfigDir,
	kubedb.HazelcastTempConfigDir,
	kubedb.HazelcastCustomConfigDir,
	kubedb.HazelcastDataDir,
}

var _ webhook.CustomValidator = &HazelcastCustomWebhook{}

func (w *HazelcastCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Hazelcast)
	if !ok {
		return nil, fmt.Errorf("expected an Hazelcast object but got %T", obj)
	}

	hazelcastlog.Info("validate create", "name", db.Name)
	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindHazelcast}, db.Name, allErr)
}

func (w *HazelcastCustomWebhook) ValidateCreateOrUpdate(h *olddbapi.Hazelcast) field.ErrorList {
	var allErr field.ErrorList

	if h.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			h.Name,
			"spec.version' is missing"))
	} else {
		err := w.hazelcastValidateVersion(h)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				h.Name,
				err.Error()))
		}
	}

	err := w.hazelcastValidateVolumes(&h.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			h.Name,
			err.Error()))
	}
	err = w.hazelcastValidateVolumesMountPaths(&h.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("containers"),
			h.Name,
			err.Error()))
	}

	if h.Spec.StorageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
			h.Name,
			"StorageType can not be empty"))
	} else {
		if h.Spec.StorageType != olddbapi.StorageTypeDurable && h.Spec.StorageType != olddbapi.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				h.Name,
				"StorageType should be either durable or ephemeral"))
		}
	}

	return allErr
}

func (w *HazelcastCustomWebhook) hazelcastValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range hazelcastReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserve volume name: " + rv)
			}
		}
	}

	return nil
}

func (w *HazelcastCustomWebhook) hazelcastValidateVersion(h *olddbapi.Hazelcast) error {
	hzVersion := catalog.HazelcastVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: h.Spec.Version}, &hzVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

func (w *HazelcastCustomWebhook) hazelcastValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range hazelcastReservedVolumeMountPaths {
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

	for _, rvmp := range hazelcastReservedVolumeMountPaths {
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

func (w *HazelcastCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Hazelcast)
	if !ok {
		return nil, fmt.Errorf("expected an Hazelcast object but got %T", newObj)
	}
	hazelcastlog.Info("validate update", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindHazelcast}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *HazelcastCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Hazelcast)
	if !ok {
		return nil, fmt.Errorf("expected an Hazelcast object but got %T", obj)
	}
	hazelcastlog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindHazelcast}, db.Name, allErr)
	}
	return nil, nil
}
