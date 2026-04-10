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

	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// SetupNeo4jWebhookWithManager registers the webhook for Neo4j in the manager.
func SetupNeo4jWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Neo4j{}).
		WithValidator(&Neo4jCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&Neo4jCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-Neo4j-kubedb-com-v1alpha1-Neo4j,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=Neo4js,verbs=create;update,versions=v1alpha1,name=mNeo4j.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type Neo4jCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &Neo4jCustomWebhook{}

// log is for logging in this package.
var Neo4jlog = logf.Log.WithName("Neo4j-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *Neo4jCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Neo4j)
	if !ok {
		return fmt.Errorf("expected an Neo4j object but got %T", obj)
	}

	Neo4jlog.Info("default", "name", db.GetName())
	db.SetDefaults(w.DefaultClient)
	return nil
}

//+kubebuilder:webhook:path=/validate-Neo4j-kubedb-com-v1alpha1-Neo4j,mutating=false,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=Neo4js,verbs=create;update,versions=v1alpha1,name=vNeo4j.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.CustomValidator = &Neo4jCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *Neo4jCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Neo4j)
	if !ok {
		return nil, fmt.Errorf("expected an Neo4j object but got %T", obj)
	}
	Neo4jlog.Info("validate create", "name", db.GetName())
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *Neo4jCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Neo4j)
	if !ok {
		return nil, fmt.Errorf("expected an Neo4j object but got %T", newObj)
	}

	Neo4jlog.Info("validate update", "name", db.GetName())
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *Neo4jCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Neo4j)
	if !ok {
		return nil, fmt.Errorf("expected an Neo4j object but got %T", obj)
	}
	Neo4jlog.Info("validate delete", "name", db.GetName())

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deletionPolicy"),
			db.GetName(),
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "Neo4j.kubedb.com", Kind: "Neo4j"}, db.GetName(), allErr)
	}
	return nil, nil
}

func (w *Neo4jCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Neo4j) error {
	var allErr field.ErrorList

	if db.Spec.TLS != nil && db.Spec.TLS.IssuerRef == nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("tls").Child("issuerRef"),
			db.Name, "spec.tls.issuerRef' is missing"))
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

	if db.Spec.Monitor != nil {
		if db.Spec.Monitor.Agent == "" {
			return fmt.Errorf("monitor.agent is missing")
		}
		if !mona.IsKnownAgentType(db.Spec.Monitor.Agent) {
			return fmt.Errorf("monitor.agent '%v' is not known", db.Spec.Monitor.Agent)
		}
	}

	if db.Spec.Configuration != nil && db.Spec.Configuration.SecretName != "" {
		configSecret := &core.Secret{}
		err := w.DefaultClient.Get(context.TODO(), client.ObjectKey{
			Name:      db.Spec.Configuration.SecretName,
			Namespace: db.Namespace,
		}, configSecret)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("configuration").Child("secretName"),
				db.GetName(),
				"failed to get ConfigSecret"))
		}
	}

	if len(allErr) == 0 {
		return nil
	}
	return apierrors.NewInvalid(schema.GroupKind{Group: "Neo4j.kubedb.com", Kind: "Neo4j"}, db.GetName(), allErr)
}

func (w *Neo4jCustomWebhook) ValidateVersion(db *olddbapi.Neo4j) error {
	rmVersion := catalog.Neo4jVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: db.Spec.Version}, &rmVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

var Neo4jReservedVolumes = []string{
	kubedb.Neo4jVolumeData,
}

func (w *Neo4jCustomWebhook) validateVolumes(db *olddbapi.Neo4j) error {
	if db.Spec.PodTemplate.Spec.Volumes == nil {
		return nil
	}
	rsv := make([]string, len(Neo4jReservedVolumes))
	copy(rsv, Neo4jReservedVolumes)
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

var Neo4jReservedVolumeMountPaths = []string{
	kubedb.Neo4jDataDir,
}

func (w *Neo4jCustomWebhook) validateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range Neo4jReservedVolumeMountPaths {
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

	for _, rvmp := range Neo4jReservedVolumeMountPaths {
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
