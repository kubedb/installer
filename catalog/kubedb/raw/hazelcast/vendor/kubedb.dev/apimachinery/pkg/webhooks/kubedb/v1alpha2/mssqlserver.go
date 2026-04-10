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
	"k8s.io/utils/ptr"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// SetupMSSQLServerWebhookWithManager registers the webhook for MSSQLServer in the manager.
func SetupMSSQLServerWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.MSSQLServer{}).
		WithValidator(&MSSQLServerCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&MSSQLServerCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-kubedb-com-v1alpha2-mssqlserver,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=mssqlservers,verbs=create;update,versions=v1alpha2,name=mmssqlserver.kb.io,admissionReviewVersions=v1

// +kubebuilder:object:generate=false
type MSSQLServerCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &MSSQLServerCustomWebhook{}

// log is for logging in this package.
var mssqllog = logf.Log.WithName("mssql-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *MSSQLServerCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.MSSQLServer)
	if !ok {
		return fmt.Errorf("expected a MSSQLServer object, got a %T", obj)
	}

	mssqllog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &MSSQLServerCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *MSSQLServerCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.MSSQLServer)
	if !ok {
		return nil, fmt.Errorf("expected a MSSQLServer object, got a %T", obj)
	}

	mssqllog.Info("validate create", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindMSSQLServer}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *MSSQLServerCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.MSSQLServer)
	if !ok {
		return nil, fmt.Errorf("expected a MSSQLServer object, got a %T", newObj)
	}

	mssqllog.Info("validate update", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindMSSQLServer}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *MSSQLServerCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.MSSQLServer)
	if !ok {
		return nil, fmt.Errorf("expected a MSSQLServer object, got a %T", obj)
	}

	mssqllog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: kubedb.GroupName, Kind: olddbapi.ResourceKindMSSQLServer}, db.Name, allErr)
	}
	return nil, nil
}

func (w *MSSQLServerCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.MSSQLServer) field.ErrorList {
	var allErr field.ErrorList

	err := w.mssqlValidateVersion(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			err.Error()))
	}

	if db.IsStandalone() {
		if ptr.Deref(db.Spec.Replicas, 0) != 1 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas for standalone must be one "))
		}
	} else {
		if db.Spec.Topology.Mode == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("mode"),
				db.Name,
				".spec.topology.mode can't be empty in cluster mode"))
		}

		if ptr.Deref(db.Spec.Replicas, 0) <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas can not be nil and can not be less than or equal to 0"))
		}
	}

	if db.Spec.TLS == nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("tls"),
			db.Name, "spec.tls is missing"))
	} else {
		if db.Spec.TLS.IssuerRef == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("tls").Child("issuerRef"),
				db.Name, "spec.tls.issuerRef' is missing"))
		}
	}

	if db.Spec.PodTemplate != nil {
		if err = ValidateMSSQLServerEnvVar(getMSSQLServerContainerEnvs(db), forbiddenMSSQLServerEnvVars, db.ResourceKind()); err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate"),
				db.Name,
				err.Error()))
		}
	}

	err = mssqlValidateVolumes(db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = mssqlValidateVolumesMountPaths(db.Spec.PodTemplate)
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

// reserved volume and volumes mounts for mssql
var mssqlReservedVolumes = []string{
	kubedb.MSSQLVolumeNameData,
	kubedb.MSSQLVolumeNameConfig,
	kubedb.MSSQLVolumeNameInitScript,
	kubedb.MSSQLVolumeNameEndpointCert,
	kubedb.MSSQLVolumeNameCerts,
	kubedb.MSSQLVolumeNameTLS,
	kubedb.MSSQLVolumeNameSecurityCACertificates,
	kubedb.MSSQLVolumeNameCACerts,
}

var mssqlReservedVolumesMountPaths = []string{
	kubedb.MSSQLVolumeMountPathData,
	kubedb.MSSQLVolumeMountPathConfig,
	kubedb.MSSQLVolumeMountPathInitScript,
	kubedb.MSSQLVolumeMountPathEndpointCert,
	kubedb.MSSQLVolumeMountPathCerts,
	kubedb.MSSQLVolumeMountPathTLS,
	kubedb.MSSQLVolumeMountPathSecurityCACertificates,
	kubedb.MSSQLVolumeMountPathCACerts,
}

func (w *MSSQLServerCustomWebhook) mssqlValidateVersion(db *olddbapi.MSSQLServer) error {
	var mssqlVersion catalog.MSSQLServerVersion

	return w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &mssqlVersion)
}

func mssqlValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range mssqlReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserved volume name: " + rv)
			}
		}
	}

	return nil
}

func mssqlValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}

	if podTemplate.Spec.Containers != nil {
		// Check container volume mounts
		for _, rvmp := range mssqlReservedVolumesMountPaths {
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
		for _, rvmp := range mssqlReservedVolumesMountPaths {
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

var forbiddenMSSQLServerEnvVars = []string{
	kubedb.EnvMSSQLSAUsername,
	kubedb.EnvMSSQLSAPassword,
	kubedb.EnvMSSQLEnableHADR,
	kubedb.EnvMSSQLAgentEnabled,
	kubedb.EnvMSSQLVersion,
}

func getMSSQLServerContainerEnvs(db *olddbapi.MSSQLServer) []core.EnvVar {
	for _, container := range db.Spec.PodTemplate.Spec.Containers {
		if container.Name == kubedb.MSSQLContainerName {
			return container.Env
		}
	}
	return []core.EnvVar{}
}

func ValidateMSSQLServerEnvVar(envs []core.EnvVar, forbiddenEnvs []string, resourceType string) error {
	presentMSSQL_PID := false
	for _, env := range envs {
		present, _ := arrays.Contains(forbiddenEnvs, env.Name)
		if present {
			return fmt.Errorf("environment variable %s is forbidden to use in %s spec", env.Name, resourceType)
		}
		if env.Name == "MSSQL_PID" {
			presentMSSQL_PID = true
		}
	}
	if !presentMSSQL_PID {
		return fmt.Errorf("environment variable %s must be provided in %s spec", kubedb.EnvMSSQLPid, resourceType)
	}
	return nil
}
