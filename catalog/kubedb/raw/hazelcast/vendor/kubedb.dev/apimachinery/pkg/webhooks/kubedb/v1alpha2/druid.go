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

// SetupDruidWebhookWithManager registers the webhook for Druid in the manager.
func SetupDruidWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Druid{}).
		WithValidator(&DruidCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&DruidCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-druid-kubedb-com-v1alpha1-druid,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=druids,verbs=create;update,versions=v1alpha1,name=mdruid.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type DruidCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &DruidCustomWebhook{}

// log is for logging in this package.
var druidlog = logf.Log.WithName("druid-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *DruidCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Druid)
	if !ok {
		return fmt.Errorf("expected an Druid object but got %T", obj)
	}

	druidlog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

//+kubebuilder:webhook:path=/validate-kubedb-com-v1alpha2-druid,mutating=false,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=druids,verbs=create;update,versions=v1alpha2,name=vdruid.kb.io,admissionReviewVersions=v1

var _ webhook.CustomValidator = &DruidCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *DruidCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Druid)
	if !ok {
		return nil, fmt.Errorf("expected an Druid object but got %T", obj)
	}

	allErr := w.validateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Druid"}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *DruidCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Druid)
	if !ok {
		return nil, fmt.Errorf("expected an Druid object but got %T", newObj)
	}

	druidlog.Info("validate update", "name", db.Name)
	_ = old.(*olddbapi.Druid)
	allErr := w.validateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Druid"}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *DruidCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Druid)
	if !ok {
		return nil, fmt.Errorf("expected an Druid object but got %T", obj)
	}

	druidlog.Info("validate delete", "name", db.Name)
	return nil, nil
}

var druidReservedVolumes = []string{
	kubedb.DruidVolumeOperatorConfig,
	kubedb.DruidVolumeMainConfig,
	kubedb.DruidVolumeCustomConfig,
	kubedb.DruidVolumeMySQLMetadataStorage,
}

var druidReservedVolumeMountPaths = []string{
	kubedb.DruidCConfigDirMySQLMetadata,
	kubedb.DruidOperatorConfigDir,
	kubedb.DruidMainConfigDir,
	kubedb.DruidCustomConfigDir,
}

func (w *DruidCustomWebhook) validateCreateOrUpdate(db *olddbapi.Druid) field.ErrorList {
	var allErr field.ErrorList

	if db.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			"spec.version is missing"))
	} else {
		err := w.druidValidateVersion(db)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				db.Name,
				err.Error()))
		}
	}

	if db.Spec.DeepStorage == nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deepStorage"),
			db.Name,
			"spec.deepStorage is missing"))
	} else {
		if db.Spec.DeepStorage.Type == "" {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deepStorage").Child("type"),
				db.Name,
				"spec.deepStorage.type is missing"))
		}
	}

	if db.Spec.MetadataStorage.Name == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("metadataStorage").Child("name"),
			db.Name,
			"spec.metadataStorage.name can not be empty"))
	}
	if db.Spec.MetadataStorage.Namespace == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("metadataStorage").Child("namespace"),
			db.Name,
			"spec.metadataStorage.namespace can not be empty"))
	}

	if db.Spec.ZookeeperRef.Name == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("zookeeperRef").Child("name"),
			db.Name,
			"spec.zookeeperRef.name can not be empty"))
	}
	if db.Spec.ZookeeperRef.Namespace == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("zookeeperRef").Child("namespace"),
			db.Name,
			"spec.zookeeperRef.namespace can not be empty"))
	}

	if db.Spec.Topology == nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology"),
			db.Name,
			"spec.topology can not be empty"))
	} else {
		// Required Nodes
		if db.Spec.Topology.Coordinators == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("coordinators"),
				db.Name,
				"spec.topology.coordinators can not be empty"))
		} else {
			w.validateDruidNode(db, olddbapi.DruidNodeRoleCoordinators, &allErr)
		}

		if db.Spec.Topology.Brokers == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("brokers"),
				db.Name,
				"spec.topology.brokers can not be empty"))
		} else {
			w.validateDruidNode(db, olddbapi.DruidNodeRoleBrokers, &allErr)
		}

		// Optional Nodes
		if db.Spec.Topology.Overlords != nil {
			w.validateDruidNode(db, olddbapi.DruidNodeRoleOverlords, &allErr)
		}
		if db.Spec.Topology.Routers != nil {
			w.validateDruidNode(db, olddbapi.DruidNodeRoleRouters, &allErr)
		}

		// Data Nodes
		if db.Spec.Topology.MiddleManagers == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("middleManagers"),
				db.Name,
				"spec.topology.middleManagers can not be empty"))
		} else {
			w.validateDruidDataNode(db, olddbapi.DruidNodeRoleMiddleManagers, &allErr)
		}
		if db.Spec.Topology.Historicals == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("historicals"),
				db.Name,
				"spec.topology.historicals can not be empty"))
		} else {
			w.validateDruidDataNode(db, olddbapi.DruidNodeRoleHistoricals, &allErr)
		}
	}
	if len(allErr) == 0 {
		return nil
	}
	return allErr
}

func (w *DruidCustomWebhook) validateDruidNode(db *olddbapi.Druid, nodeType olddbapi.DruidNodeRoleType, allErr *field.ErrorList) {
	node, dataNode := db.GetNodeSpec(nodeType)
	if dataNode != nil {
		node = &dataNode.DruidNode
	}

	if *node.Replicas <= 0 {
		*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("topology").Child(string(nodeType)).Child("replicas"),
			db.Name,
			"number of replicas can not be 0 or less"))
	}

	err := druidValidateVolumes(&node.PodTemplate, nodeType)
	if err != nil {
		*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("topology").Child(string(nodeType)).Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}
	err = druidValidateVolumesMountPaths(&node.PodTemplate, nodeType)
	if err != nil {
		*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("topology").Child(string(nodeType)).Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}
}

func (w *DruidCustomWebhook) validateDruidDataNode(db *olddbapi.Druid, nodeType olddbapi.DruidNodeRoleType, allErr *field.ErrorList) {
	w.validateDruidNode(db, nodeType, allErr)

	_, dataNode := db.GetNodeSpec(nodeType)
	if dataNode.StorageType == "" {
		*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("topology").Child(string(nodeType)).Child("storageType"),
			db.Name,
			fmt.Sprintf("spec.topology.%s.storageType can not be empty", string(nodeType))))
	} else {
		if dataNode.StorageType != olddbapi.StorageTypeDurable && dataNode.StorageType != olddbapi.StorageTypeEphemeral {
			*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				db.Name,
				fmt.Sprintf("spec.topology.%s.storageType should either be durable or ephemeral", string(nodeType))))
		}
	}
	if dataNode.StorageType == olddbapi.StorageTypeEphemeral && dataNode.Storage != nil {
		*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("topology").Child(string(nodeType)).Child("storage"),
			db.Name,
			fmt.Sprintf("spec.topology.%s.storage can not be set when db.Spec.topology.%s.storageType is Ephemeral", string(nodeType), string(nodeType))))
	}
	if dataNode.StorageType == olddbapi.StorageTypeDurable && dataNode.EphemeralStorage != nil {
		*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("topology").Child(string(nodeType)).Child("ephemeralStorage"),
			db.Name,
			fmt.Sprintf("spec.topology.%s.ephemeralStorage can not be set when d.spec.topology.%s.storageType is Durable", string(nodeType), string(nodeType))))
	}
	if dataNode.StorageType == olddbapi.StorageTypeDurable && dataNode.Storage == nil {
		*allErr = append(*allErr, field.Invalid(field.NewPath("spec").Child("topology").Child(string(nodeType)).Child("storage"),
			db.Name,
			fmt.Sprintf("spec.topology.%s.storage needs to be set when spec.topology.%s.storageType is Durable", string(nodeType), string(nodeType))))
	}
}

func (w *DruidCustomWebhook) druidValidateVersion(db *olddbapi.Druid) error {
	var druidVersion catalog.DruidVersion
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{
		Name: db.Spec.Version,
	}, &druidVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

func druidValidateVolumes(podTemplate *ofst.PodTemplateSpec, nodeType olddbapi.DruidNodeRoleType) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	switch nodeType {
	case olddbapi.DruidNodeRoleHistoricals:
		druidReservedVolumes = append(druidReservedVolumes, kubedb.DruidVolumeHistoricalsSegmentCache)
	case olddbapi.DruidNodeRoleMiddleManagers:
		druidReservedVolumes = append(druidReservedVolumes, kubedb.DruidVolumeMiddleManagersBaseTaskDir)
	}

	for _, rv := range druidReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserve volume name: " + rv)
			}
		}
	}

	return nil
}

func druidValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec, nodeType olddbapi.DruidNodeRoleType) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	if nodeType == olddbapi.DruidNodeRoleHistoricals {
		druidReservedVolumeMountPaths = append(druidReservedVolumeMountPaths, kubedb.DruidHistoricalsSegmentCacheDir)
	}
	if nodeType == olddbapi.DruidNodeRoleMiddleManagers {
		druidReservedVolumeMountPaths = append(druidReservedVolumeMountPaths, kubedb.DruidWorkerTaskBaseTaskDir)
	}

	for _, rvmp := range druidReservedVolumeMountPaths {
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

	for _, rvmp := range druidReservedVolumeMountPaths {
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
