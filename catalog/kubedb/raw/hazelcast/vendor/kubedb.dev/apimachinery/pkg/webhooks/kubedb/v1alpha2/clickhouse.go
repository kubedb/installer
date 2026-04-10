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

// SetupClickHouseWebhookWithManager registers the webhook for ClickHouse in the manager.
func SetupClickHouseWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.ClickHouse{}).
		WithValidator(&ClickHouseCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&ClickHouseCustomWebhook{mgr.GetClient()}).
		Complete()
}

type ClickHouseCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &ClickHouseCustomWebhook{}

// log is for logging in this package.
var clickhouselog = logf.Log.WithName("clickhouse-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *ClickHouseCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.ClickHouse)
	if !ok {
		return fmt.Errorf("expected an ClickHouse object but got %T", obj)
	}
	clickhouselog.Info("default", "name", db.Name)
	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &ClickHouseCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *ClickHouseCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.ClickHouse)
	if !ok {
		return nil, fmt.Errorf("expected an ClickHouse object but got %T", obj)
	}
	clickhouselog.Info("validate create", "name", db.Name)
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *ClickHouseCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.ClickHouse)
	if !ok {
		return nil, fmt.Errorf("expected an ClickHouse object but got %T", newObj)
	}
	clickhouselog.Info("validate update", "name", db.Name)
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *ClickHouseCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.ClickHouse)
	if !ok {
		return nil, fmt.Errorf("expected an ClickHouse object but got %T", obj)
	}
	clickhouselog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("teminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "ClickHouse.kubedb.com", Kind: "ClickHouse"}, db.Name, allErr)
	}
	return nil, nil
}

func (w *ClickHouseCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.ClickHouse) error {
	var allErr field.ErrorList

	if db.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			"spec.version' is missing"))
		return apierrors.NewInvalid(schema.GroupKind{Group: "ClickHouse.kubedb.com", Kind: "ClickHouse"}, db.Name, allErr)
	} else {
		err := w.ValidateVersion(db)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				db.Spec.Version,
				err.Error()))
			return apierrors.NewInvalid(schema.GroupKind{Group: "ClickHouse.kubedb.com", Kind: "ClickHouse"}, db.Name, allErr)
		}
	}

	if db.Spec.TLS != nil {
		if db.Spec.TLS.ClientCACertificateRefs != nil {
			for _, secret := range db.Spec.TLS.ClientCACertificateRefs {
				if secret.Name == "" {
					allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("tls").Child("clientCaCertificateRef").Child("name"),
						db.Name,
						"secret name is missing"))
					return apierrors.NewInvalid(schema.GroupKind{Group: "ClickHouse.kubedb.com", Kind: "ClickHouse"}, db.Name, allErr)
				}
			}
		}
		if db.Spec.SSLVerificationMode == "" {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("tls").Child("sslVerificationMode"),
				db.Name,
				"sslVerificationMode is missing"))
			return apierrors.NewInvalid(schema.GroupKind{Group: "ClickHouse.kubedb.com", Kind: "ClickHouse"}, db.Name, allErr)
		}
	}

	if db.Spec.DisableSecurity {
		if db.Spec.AuthSecret != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("authSecret"),
				db.Name,
				"authSecret should be nil when security is disabled"))
			return apierrors.NewInvalid(schema.GroupKind{Group: "ClickHouse.kubedb.com", Kind: "ClickHouse"}, db.Name, allErr)
		}
	}

	if db.Spec.ClusterTopology != nil {
		clusterName := map[string]bool{}
		cluster := db.Spec.ClusterTopology.Cluster
		if db.Spec.ClusterTopology.ClickHouseKeeper != nil {
			if !db.Spec.ClusterTopology.ClickHouseKeeper.ExternallyManaged {
				if db.Spec.ClusterTopology.ClickHouseKeeper.Spec == nil {
					allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("spec"),
						db.Name,
						"spec can't be nil when externally managed is false"))
				} else {
					if *db.Spec.ClusterTopology.ClickHouseKeeper.Spec.Replicas < 1 {
						allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("spec").Child("replica"),
							db.Name,
							"number of replica can not be 0 or less"))
					}
					allErr = w.validateClickHouseKeeperStorageType(db, allErr)
				}
				if db.Spec.ClusterTopology.ClickHouseKeeper.Node != nil {
					allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("node"),
						db.Name,
						"ClickHouse Keeper node should be empty when externally managed is false"))
				}
			} else {
				if db.Spec.ClusterTopology.ClickHouseKeeper.Node == nil {
					allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("node"),
						db.Name,
						"ClickHouse Keeper node can't be empty when externally managed is true"))
				} else {
					if db.Spec.ClusterTopology.ClickHouseKeeper.Node.Host == "" {
						allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("node").Child("host"),
							db.Name,
							"ClickHouse Keeper host can't be empty"))
					}
					if db.Spec.ClusterTopology.ClickHouseKeeper.Node.Port == nil {
						allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("node").Child("port"),
							db.Name,
							"ClickHouse Keeper port can't be empty"))
					}
				}
				if db.Spec.ClusterTopology.ClickHouseKeeper.Spec != nil {
					allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("spec"),
						db.Name,
						"ClickHouse Keeper spec should be empty when externally managed is true"))
				}
			}
		}
		if cluster.Shards != nil && *cluster.Shards <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("shards"),
				db.Name,
				"number of shards can not be 0 or less"))
		}
		if cluster.Replicas != nil && *cluster.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("replicas"),
				db.Name,
				"number of replicas can't be 0 or less"))
		}
		if clusterName[cluster.Name] {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child(cluster.Name),
				db.Name,
				"cluster name is already exists, use different cluster name"))
		}
		clusterName[cluster.Name] = true

		allErr = w.validateClusterStorageType(db, cluster, allErr)

		err := w.validateVolumes(cluster.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = w.validateVolumesMountPaths(cluster.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("podTemplate").Child("spec").Child("volumeMounts"),
				db.Name,
				err.Error()))
		}
		if db.Spec.PodTemplate != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate"),
				db.Name,
				"PodTemplate should be nil in clusterTopology"))
		}

		if db.Spec.Replicas != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replica"),
				db.Name,
				"replica should be nil in clusterTopology"))
		}

		if db.Spec.StorageType != "" {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				db.Name,
				"StorageType should be empty in clusterTopology"))
		}

		if db.Spec.Storage != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storage"),
				db.Name,
				"storage should be nil in clusterTopology"))
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

		allErr = w.validateStandaloneStorageType(db, allErr)
	}

	if len(allErr) == 0 {
		return nil
	}
	return apierrors.NewInvalid(schema.GroupKind{Group: "ClickHouse.kubedb.com", Kind: "ClickHouse"}, db.Name, allErr)
}

func (w *ClickHouseCustomWebhook) validateStandaloneStorageType(db *olddbapi.ClickHouse, allErr field.ErrorList) field.ErrorList {
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

	if db.Spec.Storage == nil && db.Spec.StorageType == olddbapi.StorageTypeDurable {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storage"),
			db.Name,
			"Storage can't be empty when StorageType is durable"))
	}

	return allErr
}

func (w *ClickHouseCustomWebhook) validateClusterStorageType(db *olddbapi.ClickHouse, cluster olddbapi.ClusterSpec, allErr field.ErrorList) field.ErrorList {
	if cluster.StorageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child(cluster.Name).Child("storageType"),
			db.Name,
			"StorageType can not be empty"))
	} else {
		if cluster.StorageType != olddbapi.StorageTypeDurable && cluster.StorageType != olddbapi.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child(cluster.Name).Child("storageType"),
				cluster.StorageType,
				"StorageType should be either durable or ephemeral"))
		}
	}
	if cluster.Storage == nil && cluster.StorageType == olddbapi.StorageTypeDurable {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child(cluster.Name).Child("storage"),
			db.Name,
			"Storage can't be empty when StorageType is durable"))
	}
	return allErr
}

func (w *ClickHouseCustomWebhook) validateClickHouseKeeperStorageType(db *olddbapi.ClickHouse, allErr field.ErrorList) field.ErrorList {
	if db.Spec.ClusterTopology.ClickHouseKeeper.Spec.StorageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("spec").Child("storageType"),
			db.Name,
			"StorageType can not be empty"))
	} else {
		if db.Spec.ClusterTopology.ClickHouseKeeper.Spec.StorageType != olddbapi.StorageTypeDurable && db.Spec.StorageType != olddbapi.StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("spec").Child("storageType"),
				db.Name,
				"StorageType should be either durable or ephemeral"))
		}
	}
	if db.Spec.ClusterTopology.ClickHouseKeeper.Spec.Storage == nil && db.Spec.StorageType == olddbapi.StorageTypeDurable {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("clusterTopology").Child("clickHouseKeeper").Child("spec").Child("storage"),
			db.Name,
			"Storage can't be empty when StorageType is durable"))
	}

	return allErr
}

func (w *ClickHouseCustomWebhook) ValidateVersion(db *olddbapi.ClickHouse) error {
	chVersion := catalog.ClickHouseVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: db.Spec.Version}, &chVersion)
	if err != nil {
		// fmt.Sprint(db.Spec.Version, "version not supported")
		return errors.New(fmt.Sprint("version ", db.Spec.Version, " not supported"))
	}
	return nil
}

var clickhouseReservedVolumes = []string{
	kubedb.ClickHouseVolumeData,
}

func (w *ClickHouseCustomWebhook) validateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate.Spec.Volumes == nil {
		return nil
	}
	rsv := make([]string, len(clickhouseReservedVolumes))
	copy(rsv, clickhouseReservedVolumes)
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

var clickhouseReservedVolumeMountPaths = []string{
	kubedb.ClickHouseDataDir,
}

func (w *ClickHouseCustomWebhook) validateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range clickhouseReservedVolumeMountPaths {
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

	for _, rvmp := range clickhouseReservedVolumeMountPaths {
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
