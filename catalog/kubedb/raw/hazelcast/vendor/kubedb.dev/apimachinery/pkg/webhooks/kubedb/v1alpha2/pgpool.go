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
	"unsafe"

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1"
	olddbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	amv "kubedb.dev/apimachinery/pkg/validator"

	"github.com/pkg/errors"
	"gomodules.xyz/x/arrays"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	kmapi "kmodules.xyz/client-go/api/v1"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// SetupPgpoolWebhookWithManager registers the webhook for Pgpool in the manager.
func SetupPgpoolWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Pgpool{}).
		WithValidator(&PgpoolCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&PgpoolCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-pgpool-kubedb-com-v1alpha1-pgpool,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=pgpools,verbs=create;update,versions=v1alpha1,name=mpgpool.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type PgpoolCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &PgpoolCustomWebhook{}

// log is for logging in this package.
var pgpoollog = logf.Log.WithName("pgpool-resource")

var _ webhook.CustomDefaulter = &PgpoolCustomWebhook{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *PgpoolCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	pp, ok := obj.(*olddbapi.Pgpool)
	if !ok {
		return fmt.Errorf("expected an pgpool object but got %T", obj)
	}
	pgpoollog.Info("default", "name", pp.Name)
	pp.SetDefaults(w.DefaultClient)
	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-kubedb-com-v1alpha2-pgpool,mutating=false,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=pgpools,verbs=create;update;delete,versions=v1alpha2,name=vpgpool.kb.io,admissionReviewVersions=v1

var _ webhook.CustomValidator = &PgpoolCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *PgpoolCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	pp, ok := obj.(*olddbapi.Pgpool)
	if !ok {
		return nil, fmt.Errorf("expected an pgpool object but got %T", obj)
	}
	pgpoollog.Info("validate create", "name", pp.Name)
	errorList := w.ValidateCreateOrUpdate(pp)
	if len(errorList) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Pgpool"}, pp.Name, errorList)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *PgpoolCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	pp, ok := newObj.(*olddbapi.Pgpool)
	if !ok {
		return nil, fmt.Errorf("expected an pgpool object but got %T", pp)
	}
	pgpoollog.Info("validate update", "name", pp.Name)

	errorList := w.ValidateCreateOrUpdate(pp)
	if len(errorList) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Pgpool"}, pp.Name, errorList)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *PgpoolCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	pp, ok := obj.(*olddbapi.Pgpool)
	if !ok {
		return nil, fmt.Errorf("expected an pgpool object but got %T", pp)
	}
	pgpoollog.Info("validate delete", "name", pp.Name)

	var errorList field.ErrorList
	if pp.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			pp.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Pgpool"}, pp.Name, errorList)
	}
	return nil, nil
}

func (w *PgpoolCustomWebhook) ValidateCreateOrUpdate(pp *olddbapi.Pgpool) field.ErrorList {
	var errorList field.ErrorList
	if pp.Spec.Version == "" {
		errorList = append(errorList, field.Required(field.NewPath("spec").Child("version"),
			"`spec.version` is missing",
		))
	} else {
		err := w.PgpoolValidateVersion(pp)
		if err != nil {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("version"),
				pp.Name,
				err.Error()))
		}
	}

	if pp.Spec.PostgresRef == nil {
		errorList = append(errorList, field.Required(field.NewPath("spec").Child("postgresRef"),
			"`spec.postgresRef` is missing",
		))
	}

	if pp.DeletionTimestamp == nil {
		apb := appcat.AppBinding{}
		err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{
			Name:      pp.Spec.PostgresRef.Name,
			Namespace: pp.Spec.PostgresRef.Namespace,
		}, &apb)
		if err != nil {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("postgresRef"),
				pp.Name,
				err.Error(),
			))
		}

		backendSSL, err := pp.IsBackendTLSEnabled(w.DefaultClient)
		if err != nil {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("postgresRef"),
				pp.Name,
				err.Error(),
			))
		}

		if pp.Spec.TLS == nil && backendSSL {
			errorList = append(errorList, field.Required(field.NewPath("spec").Child("tls"),
				"`spec.tls` must be set because backend postgres is tls enabled",
			))
		}
	}

	if pp.Spec.TLS == nil {
		if pp.Spec.SSLMode != "disable" {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("sslMode"),
				pp.Name,
				"Tls is not enabled, enable it to use this sslMode",
			))
		}

		if pp.Spec.ClientAuthMode == "cert" {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("clientAuthMode"),
				pp.Name,
				"Tls is not enabled, enable it to use this clientAuthMode",
			))
		}
	}

	if pp.Spec.Replicas != nil {
		if *pp.Spec.Replicas <= 0 {
			errorList = append(errorList, field.Required(field.NewPath("spec").Child("replicas"),
				"`spec.replica` must be greater than 0",
			))
		}
		if *pp.Spec.Replicas > 9 {
			errorList = append(errorList, field.Required(field.NewPath("spec").Child("replicas"),
				"`spec.replica` must be less than 10",
			))
		}
	}

	if pp.Spec.PodTemplate != nil {
		if err := w.ValidateEnvVar(PgpoolGetMainContainerEnvs(pp), PgpoolForbiddenEnvVars, pp.ResourceKind()); err != nil {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("containers").Child("env"),
				pp.Name,
				err.Error(),
			))
		}
		err := PgpoolValidateVolumes(pp)
		if err != nil {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
				pp.Name,
				err.Error(),
			))
		}

		err = PgpoolValidateVolumesMountPaths(pp)
		if err != nil {
			errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumeMounts"),
				pp.Name,
				err.Error()))
		}
	}

	if err := w.ValidateHealth(&pp.Spec.HealthChecker); err != nil {
		errorList = append(errorList, field.Invalid(field.NewPath("spec").Child("healthChecker"),
			pp.Name,
			err.Error(),
		))
	}

	if len(errorList) == 0 {
		return nil
	}
	return errorList
}

func (w *PgpoolCustomWebhook) ValidateEnvVar(envs []core.EnvVar, forbiddenEnvs []string, resourceType string) error {
	for _, env := range envs {
		present, _ := arrays.Contains(forbiddenEnvs, env.Name)
		if present {
			return fmt.Errorf("environment variable %s is forbidden to use in %s spec", env.Name, resourceType)
		}
	}
	return nil
}

func (w *PgpoolCustomWebhook) ValidateHealth(health *kmapi.HealthCheckSpec) error {
	if health.PeriodSeconds != nil && *health.PeriodSeconds <= 0 {
		return fmt.Errorf(`spec.healthCheck.periodSeconds: can not be less than 1`)
	}

	if health.TimeoutSeconds != nil && *health.TimeoutSeconds <= 0 {
		return fmt.Errorf(`spec.healthCheck.timeoutSeconds: can not be less than 1`)
	}

	if health.FailureThreshold != nil && *health.FailureThreshold <= 0 {
		return fmt.Errorf(`spec.healthCheck.failureThreshold: can not be less than 1`)
	}
	return nil
}

func (w *PgpoolCustomWebhook) PgpoolValidateVersion(pp *olddbapi.Pgpool) error {
	ppVersion := catalog.PgpoolVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: pp.Spec.Version,
	}, &ppVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

var PgpoolReservedVolumes = []string{
	kubedb.PgpoolConfigVolumeName,
	kubedb.PgpoolTlsVolumeName,
}

func PgpoolValidateVolumes(pp *olddbapi.Pgpool) error {
	if pp.Spec.PodTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range PgpoolReservedVolumes {
		for _, ugv := range pp.Spec.PodTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Cannot use a reserve volume name: " + rv)
			}
		}
	}
	return nil
}

var PgpoolForbiddenEnvVars = []string{
	kubedb.EnvPostgresUsername, kubedb.EnvPostgresPassword, kubedb.EnvPgpoolPcpUser, kubedb.EnvPgpoolPcpPassword,
	kubedb.EnvPgpoolPasswordEncryptionMethod, kubedb.EnvEnablePoolPasswd, kubedb.EnvSkipPasswdEncryption,
}

func PgpoolGetMainContainerEnvs(pp *olddbapi.Pgpool) []core.EnvVar {
	for _, container := range pp.Spec.PodTemplate.Spec.Containers {
		if container.Name == kubedb.PgpoolContainerName {
			return container.Env
		}
	}
	return []core.EnvVar{}
}

func PgpoolValidateVolumesMountPaths(pgpool *olddbapi.Pgpool) error {
	podTemplate := pgpool.Spec.PodTemplate
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range PgpoolReservedVolumesMountPaths {
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

	for _, rvmp := range PgpoolReservedVolumesMountPaths {
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

	init := (*dbapi.InitSpec)(unsafe.Pointer(pgpool.Spec.Init))
	if err := amv.ValidateGitInitRootPath(init, PgpoolReservedVolumesMountPaths); err != nil {
		return err
	}
	return nil
}

var PgpoolReservedVolumesMountPaths = []string{
	kubedb.PgpoolConfigSecretMountPath,
	kubedb.PgpoolTlsVolumeMountPath,
}
