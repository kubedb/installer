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
	ofst "kmodules.xyz/offshoot-api/api/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// SetupCassandraWebhookWithManager registers the webhook for Cassandra in the manager.
func SetupCassandraWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Cassandra{}).
		WithValidator(&CassandraCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&CassandraCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-cassandra-kubedb-com-v1alpha1-cassandra,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=cassandras,verbs=create;update,versions=v1alpha1,name=mcassandra.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type CassandraCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &CassandraCustomWebhook{}

// log is for logging in this package.
var cassandralog = logf.Log.WithName("cassandra-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *CassandraCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Cassandra)
	if !ok {
		return fmt.Errorf("expected an Cassandra object but got %T", obj)
	}
	cassandralog.Info("default", "name", db.GetName())
	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &CassandraCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *CassandraCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Cassandra)
	if !ok {
		return nil, fmt.Errorf("expected an Cassandra object but got %T", obj)
	}
	cassandralog.Info("validate create", "name", db.Name)
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *CassandraCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Cassandra)
	if !ok {
		return nil, fmt.Errorf("expected an Cassandra object but got %T", newObj)
	}
	cassandralog.Info("validate update", "name", db.Name)
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *CassandraCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Cassandra)
	if !ok {
		return nil, fmt.Errorf("expected an Cassandra object but got %T", obj)
	}

	cassandralog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deletionPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "Cassandra.kubedb.com", Kind: "Cassandra"}, db.Name, allErr)
	}
	return nil, nil
}

func (w *CassandraCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Cassandra) error {
	var allErr field.ErrorList

	if db.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			"spec.version' is missing"))
		return apierrors.NewInvalid(schema.GroupKind{Group: "Cassandra.kubedb.com", Kind: "Cassandra"}, db.Name, allErr)
	} else {
		err := w.ValidateVersion(db)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				db.Spec.Version,
				err.Error()))
			return apierrors.NewInvalid(schema.GroupKind{Group: "cassandra.kubedb.com", Kind: "cassandra"}, db.Name, allErr)
		}
	}

	if db.Spec.Topology != nil {
		rackName := map[string]bool{}
		racks := db.Spec.Topology.Rack
		for _, rack := range racks {
			if rack.Replicas != nil && *rack.Replicas <= 0 {
				allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("Topology").Child("replicas"),
					db.Name,
					"number of replicas can't be 0 or less"))
			}
			if rackName[rack.Name] {
				allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("Topology").Child(rack.Name),
					db.Name,
					"rack name is duplicated, use different rack name"))
			}
			rackName[rack.Name] = true

			allErr = w.validateClusterStorageType(db, rack, allErr)

			err := w.validateVolumes(rack.PodTemplate)
			if err != nil {
				allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("Topology").Child("podTemplate").Child("spec").Child("volumes"),
					db.Name,
					err.Error()))
			}
			err = w.validateVolumesMountPaths(rack.PodTemplate)
			if err != nil {
				allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("Topology").Child("podTemplate").Child("spec").Child("volumeMounts"),
					db.Name,
					err.Error()))
			}
		}
		if db.Spec.PodTemplate != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate"),
				db.Name,
				"PodTemplate should be nil in Topology"))
		}

		if db.Spec.Replicas != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replica"),
				db.Name,
				"replica should be nil in Topology"))
		}

		if db.Spec.StorageType != "" {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				db.Name,
				"StorageType should be empty in Topology"))
		}

		if db.Spec.Storage != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storage"),
				db.Name,
				"storage should be nil in Topology"))
		}

	} else {
		// number of replicas can not be 0 or less
		if db.Spec.Replicas != nil && *db.Spec.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas can't be 0 or less"))
		}

		// number of replicas can not be greater than 1
		if db.Spec.Replicas != nil && *db.Spec.Replicas > 1 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas can't be greater than 1 in standalone mode"))
		}
		err := w.validateVolumes(db.Spec.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = w.validateVolumesMountPaths(db.Spec.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumeMounts"),
				db.Name,
				err.Error()))
		}

		allErr = w.validateStandaloneStorageType(db, db.Spec.StorageType, db.Spec.Storage, allErr)
	}

	if len(allErr) == 0 {
		return nil
	}
	return apierrors.NewInvalid(schema.GroupKind{Group: "Cassandra.kubedb.com", Kind: "Cassandra"}, db.Name, allErr)
}

func (w *CassandraCustomWebhook) validateStandaloneStorageType(db *olddbapi.Cassandra, storageType olddbapi.StorageType, storage *core.PersistentVolumeClaimSpec, allErr field.ErrorList) field.ErrorList {
	if storageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
			db.Name,
			"StorageType can not be empty"))
	} else {
		if storageType != olddbapi.StorageTypeDurable && db.Spec.StorageType != olddbapi.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				db.Name,
				"StorageType should be either durable or ephemeral"))
		}
	}

	if storage == nil && db.Spec.StorageType == olddbapi.StorageTypeDurable {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storage"),
			db.Name,
			"Storage can't be empty when StorageType is durable"))
	}

	return allErr
}

func (w *CassandraCustomWebhook) validateClusterStorageType(db *olddbapi.Cassandra, rack olddbapi.RackSpec, allErr field.ErrorList) field.ErrorList {
	if rack.StorageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("Topology").Child(rack.Name).Child("storageType"),
			db.Name,
			"StorageType can not be empty"))
	} else {
		if rack.StorageType != olddbapi.StorageTypeDurable && rack.StorageType != olddbapi.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("Topology").Child(rack.Name).Child("storageType"),
				rack.StorageType,
				"StorageType should be either durable or ephemeral"))
		}
	}
	if rack.Storage == nil && rack.StorageType == olddbapi.StorageTypeDurable {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("Topology").Child(rack.Name).Child("storage"),
			db.Name,
			"Storage can't be empty when StorageType is durable"))
	}
	return allErr
}

func (w *CassandraCustomWebhook) ValidateVersion(db *olddbapi.Cassandra) error {
	casVersion := catalog.CassandraVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: db.Spec.Version}, &casVersion)
	if err != nil {
		return errors.New(fmt.Sprint("version ", db.Spec.Version, " not supported"))
	}
	return nil
}

var cassandraReservedVolumes = []string{
	kubedb.CassandraVolumeData,
}

func (w *CassandraCustomWebhook) validateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate.Spec.Volumes == nil {
		return nil
	}
	rsv := make([]string, len(cassandraReservedVolumes))
	copy(rsv, cassandraReservedVolumes)
	volumes := podTemplate.Spec.Volumes
	for _, rv := range rsv {
		for _, ugv := range volumes {
			if ugv.Name == rv {
				return errors.New("Cannot use a reserve volume name: " + rv)
			}
		}
	}
	return nil
}

var cassandraReservedVolumeMountPaths = []string{
	kubedb.CassandraDataDir,
}

func (w *CassandraCustomWebhook) validateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range cassandraReservedVolumeMountPaths {
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

	for _, rvmp := range cassandraReservedVolumeMountPaths {
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
