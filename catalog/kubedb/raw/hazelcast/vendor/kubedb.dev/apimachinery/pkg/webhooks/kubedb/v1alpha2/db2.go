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

// SetupDb2WebhookWithManager registers the webhook for Db2 in the manager.
func SetupDb2WebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.DB2{}).
		WithValidator(&DB2CustomWebhook{mgr.GetClient()}).
		WithDefaulter(&DB2CustomWebhook{mgr.GetClient()}).
		Complete()
}

type DB2CustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &DB2CustomWebhook{}

// log is for logging in this package.
var db2log = logf.Log.WithName("DB2-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *DB2CustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.DB2)
	if !ok {
		return fmt.Errorf("expected an DB2 object but got %T", obj)
	}

	db2log.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &DB2CustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *DB2CustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.DB2)
	if !ok {
		return nil, fmt.Errorf("expected an DB2 object but got %T", obj)
	}

	allErr := w.validateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "DB2"}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *DB2CustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.DB2)
	if !ok {
		return nil, fmt.Errorf("expected an DB2 object but got %T", newObj)
	}

	db2log.Info("validate update", "name", db.Name)
	_ = old.(*olddbapi.DB2)
	allErr := w.validateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "DB2"}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *DB2CustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.DB2)
	if !ok {
		return nil, fmt.Errorf("expected an DB2 object but got %T", obj)
	}

	db2log.Info("validate delete", "name", db.Name)
	return nil, nil
}

func (w *DB2CustomWebhook) validateCreateOrUpdate(db *olddbapi.DB2) field.ErrorList {
	var allErr field.ErrorList
	// Check Version
	if db.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			"spec.version is required and cannot be empty"))
	} else {
		// Optional: version validation function (e.g., check if supported version)
		if err := w.DB2ValidateVersion(db); err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				db.Spec.Version,
				err.Error()))
		}
	}

	if db.Spec.StorageType == olddbapi.StorageTypeDurable {
		if db.Spec.Storage == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storage"),
				db.Name,
				"spec.storage is required when storageType is durable"))
		}
	}

	// Check AuthSecret (optional, but validate structure if provided)
	if db.Spec.AuthSecret != nil {
		if db.Spec.AuthSecret.Name == "" {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("authSecret").Child("name"),
				db.Name,
				"spec.authSecret.name cannot be empty"))
		}
	}

	// Validate HealthChecker (optional sanity check)
	if db.Spec.HealthChecker.PeriodSeconds == nil || *db.Spec.HealthChecker.PeriodSeconds <= 0 {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("healthChecker").Child("periodSeconds"),
			db.Spec.HealthChecker.PeriodSeconds,
			"periodSeconds must be greater than 0"))
	}

	if db.Spec.HealthChecker.TimeoutSeconds == nil || *db.Spec.HealthChecker.TimeoutSeconds <= 0 {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("healthChecker").Child("timeoutSeconds"),
			db.Spec.HealthChecker.TimeoutSeconds,
			"timeoutSeconds must be greater than 0"))
	}

	// Optional: Validate PodTemplate
	if db.Spec.PodTemplate != nil && db.Spec.PodTemplate.Spec.Containers == nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate"),
			db.Name,
			"spec.podTemplate.spec.containers cannot be empty"))
	}
	err := DB2ValidateVolumes(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = DB2ValidateVolumesMountPaths(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("containers"),
			db.Name,
			err.Error()))
	}

	// Optional: Validate ServiceTemplates
	for i, svc := range db.Spec.ServiceTemplates {
		if svc.Alias == "" {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("serviceTemplates").Index(i).Child("alias"),
				db.Name,
				"serviceTemplate.alias cannot be empty"))
		}
		if svc.Spec.Ports == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("serviceTemplates").Index(i).Child("spec").Child("ports"),
				db.Name,
				"serviceTemplate.spec.ports cannot be empty"))
		}
	}

	if len(allErr) == 0 {
		return nil
	}
	return allErr
}

func (w *DB2CustomWebhook) DB2ValidateVersion(db *olddbapi.DB2) error {
	var DB2Version catalog.DB2Version
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &DB2Version)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

// reserved volume and volumes mounts for DB2
var DB2ReservedVolumes = []string{
	kubedb.DB2DataVolume,
	kubedb.DB2VolumeScripts,
}

var DB2ReservedVolumesMountPaths = []string{
	kubedb.DB2DataDir,
	kubedb.DB2VolumeMountScripts,
}

func DB2ValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range DB2ReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserved volume name: " + rv)
			}
		}
	}

	return nil
}

func DB2ValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}

	if podTemplate.Spec.Containers != nil {
		// Check container volume mounts
		for _, rvmp := range DB2ReservedVolumesMountPaths {
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

	if podTemplate.Spec.InitContainers != nil {
		// Check init container volume mounts
		for _, rvmp := range DB2ReservedVolumesMountPaths {
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
	}

	return nil
}
