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

// SetupRabbitMQWebhookWithManager registers the webhook for RabbitMQ in the manager.
func SetupRabbitMQWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.RabbitMQ{}).
		WithValidator(&RabbitMQCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&RabbitMQCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-rabbitmq-kubedb-com-v1alpha1-rabbitmq,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=rabbitmqs,verbs=create;update,versions=v1alpha1,name=mrabbitmq.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type RabbitMQCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &RabbitMQCustomWebhook{}

// log is for logging in this package.
var rabbitmqlog = logf.Log.WithName("rabbitmq-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *RabbitMQCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.RabbitMQ)
	if !ok {
		return fmt.Errorf("expected an RabbitMQ object but got %T", obj)
	}

	rabbitmqlog.Info("default", "name", db.GetName())
	db.SetDefaults(w.DefaultClient)
	return nil
}

//+kubebuilder:webhook:path=/validate-rabbitmq-kubedb-com-v1alpha1-rabbitmq,mutating=false,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=rabbitmqs,verbs=create;update,versions=v1alpha1,name=vrabbitmq.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.CustomValidator = &RabbitMQCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *RabbitMQCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.RabbitMQ)
	if !ok {
		return nil, fmt.Errorf("expected an RabbitMQ object but got %T", obj)
	}
	rabbitmqlog.Info("validate create", "name", db.GetName())
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *RabbitMQCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.RabbitMQ)
	if !ok {
		return nil, fmt.Errorf("expected an RabbitMQ object but got %T", newObj)
	}

	rabbitmqlog.Info("validate update", "name", db.GetName())
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *RabbitMQCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.RabbitMQ)
	if !ok {
		return nil, fmt.Errorf("expected an RabbitMQ object but got %T", obj)
	}
	rabbitmqlog.Info("validate delete", "name", db.GetName())

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deletionPolicy"),
			db.GetName(),
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "rabbitmq.kubedb.com", Kind: "RabbitMQ"}, db.GetName(), allErr)
	}
	return nil, nil
}

func (w *RabbitMQCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.RabbitMQ) error {
	var allErr field.ErrorList
	if db.Spec.EnableSSL {
		if db.Spec.TLS == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("enableSSL"),
				db.GetName(),
				".spec.tls can't be nil, if .spec.enableSSL is true"))
		}
	} else {
		if db.Spec.TLS != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("enableSSL"),
				db.GetName(),
				".spec.tls must be nil, if .spec.enableSSL is disabled"))
		}
	}

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
	return apierrors.NewInvalid(schema.GroupKind{Group: "rabbitmq.kubedb.com", Kind: "RabbitMQ"}, db.GetName(), allErr)
}

func (w *RabbitMQCustomWebhook) ValidateVersion(db *olddbapi.RabbitMQ) error {
	rmVersion := catalog.RabbitMQVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: db.Spec.Version}, &rmVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

var rabbitmqReservedVolumes = []string{
	kubedb.RabbitMQVolumeData,
	kubedb.RabbitMQVolumeConfig,
	kubedb.RabbitMQVolumeTempConfig,
}

func (w *RabbitMQCustomWebhook) validateVolumes(db *olddbapi.RabbitMQ) error {
	if db.Spec.PodTemplate.Spec.Volumes == nil {
		return nil
	}
	rsv := make([]string, len(rabbitmqReservedVolumes))
	copy(rsv, rabbitmqReservedVolumes)
	if db.Spec.TLS != nil && db.Spec.TLS.Certificates != nil {
		for _, c := range db.Spec.TLS.Certificates {
			rsv = append(rsv, db.CertSecretVolumeName(olddbapi.RabbitMQCertificateAlias(c.Alias)))
		}
	}
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

var rabbitmqReservedVolumeMountPaths = []string{
	kubedb.RabbitMQConfigDir,
	kubedb.RabbitMQTempConfigDir,
	kubedb.RabbitMQDataDir,
	kubedb.RabbitMQPluginsDir,
	kubedb.RabbitMQCertDir,
}

func (w *RabbitMQCustomWebhook) validateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range rabbitmqReservedVolumeMountPaths {
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

	for _, rvmp := range rabbitmqReservedVolumeMountPaths {
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
