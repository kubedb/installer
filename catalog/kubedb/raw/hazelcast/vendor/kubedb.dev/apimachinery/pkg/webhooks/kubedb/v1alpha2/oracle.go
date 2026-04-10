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

// SetupOracleWebhookWithManager registers the webhook for Oracle in the manager.
func SetupOracleWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Oracle{}).
		WithValidator(&OracleCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&OracleCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-kubedb-com-v1alpha2-oracle,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=oracles,verbs=create;update,versions=v1alpha2,name=moracle.kb.io,admissionReviewVersions=v1

// +kubebuilder:object:generate=false
type OracleCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &OracleCustomWebhook{}

// log is for logging in this package.
var oraLog = logf.Log.WithName("oracle-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *OracleCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Oracle)
	if !ok {
		return fmt.Errorf("expected a Oracle object, got a %T", obj)
	}

	oraLog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &OracleCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *OracleCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Oracle)
	if !ok {
		return nil, fmt.Errorf("expected a Oracle object, got a %T", obj)
	}

	oraLog.Info("validate create", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindOracle}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *OracleCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Oracle)
	if !ok {
		return nil, fmt.Errorf("expected a Oracle object, got a %T", newObj)
	}
	olddb, ok := old.(*olddbapi.Oracle)
	if !ok {
		return nil, fmt.Errorf("expected a Oracle object, got a %T", old)
	}
	if ptr.Deref(olddb.Spec.Replicas, 0) > 1 && ptr.Deref(db.Spec.Replicas, 0) == 1 {
		return nil, fmt.Errorf("can't scale down to 1 replica")
	}
	oraLog.Info("validate update", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindOracle}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *OracleCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Oracle)
	if !ok {
		return nil, fmt.Errorf("expected a Oracle object, got a %T", obj)
	}

	oraLog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindOracle}, db.Name, allErr)
	}
	return nil, nil
}

func (w *OracleCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Oracle) field.ErrorList {
	var allErr field.ErrorList

	err := w.oracleValidateVersion(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			err.Error()))
	}
	if db.Spec.Edition != kubedb.OracleEditionEnterprise {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("edition"), db.Name, "only enterprise edition is supported for now"))
	}
	if db.Spec.Listener != nil && db.Spec.Listener.Service != nil && len(*db.Spec.Listener.Service) > 12 {
		// TODO: research if we can have more than 12 characters following some other way
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("listener").Child("service"), db.Name, "maximum 12 characters supported for now"))
	}

	if db.Spec.Mode == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("mode"), db.Name, "db.spec.mode has to be defined"))
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
	} else if db.IsDataGuardEnabled() {
		if ptr.Deref(db.Spec.Replicas, 0) <= 1 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas can not be nil and can not be less than or equal to 2"))
		}
	}

	if db.Spec.TCPSConfig != nil && db.Spec.TCPSConfig.TLS == nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("tcpsConfig").Child("tls"),
			db.Name,
			"spec.tcpsConfig.tls is missing"))
	}

	if db.Spec.PodTemplate != nil {
		if err = w.validateEnvsForAllContainers(db); err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate"),
				db.Name,
				err.Error()))
		}
	}

	err = oracleValidateVolumes(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = oracleValidateVolumesMountPaths(db.Spec.PodTemplate)
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

	if err := amv.ValidateStorage(w.DefaultClient, db.Spec.StorageType, db.Spec.Storage); err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storage"), db.Name, err.Error()))
	}

	if len(allErr) == 0 {
		return nil
	}

	return allErr
}

// reserved volume and volumes mounts for oracle
var oracleReservedVolumes = []string{
	kubedb.OracleDataVolume,
	kubedb.OracleVolumeScripts,
}

var oracleReservedVolumesMountPaths = []string{
	kubedb.OracleDataDir,
	kubedb.OracleVolumeMountScripts,
}

func (w *OracleCustomWebhook) oracleValidateVersion(db *olddbapi.Oracle) error {
	var oracleVersion catalog.OracleVersion

	return w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &oracleVersion)
}

func oracleValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range oracleReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserved volume name: " + rv)
			}
		}
	}

	return nil
}

func oracleValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}

	if podTemplate.Spec.Containers != nil {
		// Check container volume mounts
		for _, rvmp := range oracleReservedVolumesMountPaths {
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
		for _, rvmp := range oracleReservedVolumesMountPaths {
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

var forbiddenOracleEnvVars = []string{
	kubedb.OracleEnvUserName,
	kubedb.OracleEnvPassword,
	kubedb.OracleEnvOracleSID,
	kubedb.OracleEnvDataDir,
}

func (w *OracleCustomWebhook) validateEnvsForAllContainers(oracle *olddbapi.Oracle) error {
	var err error
	for _, container := range oracle.Spec.PodTemplate.Spec.Containers {
		if errC := amv.ValidateEnvVar(container.Env, forbiddenOracleEnvVars, olddbapi.ResourceKindOracle); errC != nil {
			if err == nil {
				err = errC
			} else {
				err = errors.Wrap(err, errC.Error())
			}
		}
	}
	return err
}
