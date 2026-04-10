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
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	amv "kubedb.dev/apimachinery/pkg/validator"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// SetupHanaDBWebhookWithManager registers the webhook for HanaDB in the manager.
func SetupHanaDBWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&api.HanaDB{}).
		WithValidator(&HanaDBCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&HanaDBCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-kubedb-com-v1alpha2-hanadb,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=HanaDBs,verbs=create;update,versions=v1alpha2,name=mHanaDB.kb.io,admissionReviewVersions=v1

// +kubebuilder:object:generate=false
type HanaDBCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &HanaDBCustomWebhook{}

// log is for logging in this package.
var hanaLog = logf.Log.WithName("hanadb-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *HanaDBCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*api.HanaDB)
	if !ok {
		return fmt.Errorf("expected a HanaDB object, got a %T", obj)
	}

	hanaLog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &HanaDBCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *HanaDBCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*api.HanaDB)
	if !ok {
		return nil, fmt.Errorf("expected a HanaDB object, got a %T", obj)
	}

	hanaLog.Info("validate create", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: api.ResourceKindHanaDB}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *HanaDBCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*api.HanaDB)
	if !ok {
		return nil, fmt.Errorf("expected a HanaDB object, got a %T", newObj)
	}
	olddb, ok := old.(*api.HanaDB)
	if !ok {
		return nil, fmt.Errorf("expected a HanaDB object, got a %T", old)
	}
	if ptr.Deref(olddb.Spec.Replicas, 0) > 1 && ptr.Deref(db.Spec.Replicas, 0) == 1 {
		return nil, fmt.Errorf("can't scale down to 1 replica")
	}
	hanaLog.Info("validate update", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: api.ResourceKindHanaDB}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *HanaDBCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*api.HanaDB)
	if !ok {
		return nil, fmt.Errorf("expected a HanaDB object, got a %T", obj)
	}

	hanaLog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == api.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: api.ResourceKindHanaDB}, db.Name, allErr)
	}
	return nil, nil
}

func (w *HanaDBCustomWebhook) ValidateCreateOrUpdate(db *api.HanaDB) field.ErrorList {
	var allErr field.ErrorList

	err := w.hanadbValidateVersion(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			err.Error()))
	}

	if db.Spec.Replicas == nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"), db.Name, "db.spec.replicas has to be defined"))
	}

	if db.IsStandalone() {
		if ptr.Deref(db.Spec.Replicas, 0) != 1 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas for standalone must be one "))
		}
	}

	if db.Spec.PodTemplate != nil {
		if err = w.validateEnvsForAllContainers(db); err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate"),
				db.Name,
				err.Error()))
		}
	}

	err = hanadbValidateVolumes(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = hanadbValidateVolumesMountPaths(db.Spec.PodTemplate)
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
		if db.Spec.StorageType != api.StorageTypeDurable && db.Spec.StorageType != api.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				db.Name,
				"StorageType should be either durable or ephemeral"))
		}
	}

	if err := amv.ValidateStorage(w.DefaultClient, db.Spec.StorageType, db.Spec.Storage); err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storage"), db.Name, err.Error()))
	}

	if len(allErr) == 0 {
		return nil
	}

	return allErr
}

// reserved volume and volumes mounts for HanaDB
var hanadbReservedVolumes = []string{
	kubedb.HanaDBDataVolume,
	kubedb.HanaDBVolumeScripts,
	kubedb.HanaDBVolumePasswordSecret,
}

var hanadbReservedVolumesMountPaths = []string{
	kubedb.HanaDBDataDir,
	kubedb.HanaDBVolumeMountScripts,
}

func (w *HanaDBCustomWebhook) hanadbValidateVersion(db *api.HanaDB) error {
	hanadbVersion := catalog.HanaDBVersion{}

	return w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &hanadbVersion)
}

func hanadbValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range hanadbReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserved volume name: " + rv)
			}
		}
	}

	return nil
}

func hanadbValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}

	if podTemplate.Spec.Containers != nil {
		// Check container volume mounts
		for _, rvmp := range hanadbReservedVolumesMountPaths {
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

var forbiddenHanaDBEnvVars = []string{}

func (w *HanaDBCustomWebhook) validateEnvsForAllContainers(hdb *api.HanaDB) error {
	var err error
	for _, container := range hdb.Spec.PodTemplate.Spec.Containers {
		if errC := amv.ValidateEnvVar(container.Env, forbiddenHanaDBEnvVars, api.ResourceKindHanaDB); errC != nil {
			if err == nil {
				err = errC
			} else {
				err = errors.Wrap(err, errC.Error())
			}
		}
	}
	return err
}
