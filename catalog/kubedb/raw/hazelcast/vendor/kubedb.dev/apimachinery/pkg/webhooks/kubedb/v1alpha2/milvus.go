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

	"gomodules.xyz/x/arrays"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ofstv2 "kmodules.xyz/offshoot-api/api/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// SetupMilvusServerWebhookWithManager registers the webhook for Milvus in the manager.
func SetupMilvusWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Milvus{}).
		WithValidator(&MilvusCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&MilvusCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-kubedb-com-v1alpha2-milvus,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=milvuses,verbs=create;update,versions=v1alpha2,name=milvus.kb.io,admissionReviewVersions=v1

// +kubebuilder:object:generate=false
type MilvusCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &MilvusCustomWebhook{}

// log is for logging in this package.
var milvuslog = logf.Log.WithName("milvus-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (m *MilvusCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Milvus)
	if !ok {
		return fmt.Errorf("expected a Milvus object, got a %T", obj)
	}

	milvuslog.Info("default", "name", db.Name)

	db.SetDefaults(m.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &MilvusCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (m *MilvusCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Milvus)
	if !ok {
		return nil, fmt.Errorf("expected a Milvus object, got a %T", obj)
	}

	milvuslog.Info("validate create", "name", db.Name)

	allErr := m.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindMilvus}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (m *MilvusCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Milvus)
	if !ok {
		return nil, fmt.Errorf("expected a Milvus object, got a %T", newObj)
	}

	milvuslog.Info("validate update", "name", db.Name)

	allErr := m.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindMilvus}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (m *MilvusCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Milvus)
	if !ok {
		return nil, fmt.Errorf("expected a Milvus object, got a %T", obj)
	}

	milvuslog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindMilvus}, db.Name, allErr)
	}
	return nil, nil
}

func (m *MilvusCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Milvus) field.ErrorList {
	var allErr field.ErrorList

	err := m.milvusValidateVersion(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			err.Error()))
	}

	if db.Spec.PodTemplate != nil {
		if err = ValidateMilvusEnvVar(getMilvusContainerEnvs(db), forbiddenMilvusEnvVars, db.ResourceKind()); err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate"),
				db.Name,
				err.Error()))
		}
	}

	err = milvusValidateVolumes(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = milvusValidateVolumesMountPaths(db.Spec.PodTemplate)
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

	if len(allErr) == 0 {
		return nil
	}

	return allErr
}

// reserved volume and volumes mounts for milvus
var milvusReservedVolumes = []string{
	kubedb.MilvusVolumeNameData,
	kubedb.MilvusConfigVolName,
	kubedb.MilvusConfigFileName,
}

var milvusReservedVolumesMountPaths = []string{
	kubedb.MilvusDataDir,
	kubedb.MilvusConfigVolDir,
}

func (m *MilvusCustomWebhook) milvusValidateVersion(db *olddbapi.Milvus) error {
	var milvusVersion catalog.MilvusVersion

	return m.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &milvusVersion)
}

func milvusValidateVolumes(podTemplate *ofstv2.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range milvusReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserved volume name: " + rv)
			}
		}
	}

	return nil
}

func milvusValidateVolumesMountPaths(podTemplate *ofstv2.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}

	if podTemplate.Spec.Containers != nil {
		// Check container volume mounts
		for _, rvmp := range milvusReservedVolumesMountPaths {
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

var forbiddenMilvusEnvVars = []string{
	kubedb.MinioAddressName,
	kubedb.MinioAddressKey,
	kubedb.MinioAccessKeyName,
	kubedb.MinioAccessKey,
	kubedb.MinioSecretKeyName,
	kubedb.MinioSecretKey,
	kubedb.EtcdEndpointsName,
}

func getMilvusContainerEnvs(db *olddbapi.Milvus) []core.EnvVar {
	for _, container := range db.Spec.PodTemplate.Spec.Containers {
		if container.Name == kubedb.MilvusContainerName {
			return container.Env
		}
	}
	return []core.EnvVar{}
}

func ValidateMilvusEnvVar(envs []core.EnvVar, forbiddenEnvs []string, resourceType string) error {
	for _, env := range envs {
		present, _ := arrays.Contains(forbiddenEnvs, env.Name)
		if present {
			return fmt.Errorf("environment variable %s is forbidden to use in %s spec", env.Name, resourceType)
		}
	}
	return nil
}
