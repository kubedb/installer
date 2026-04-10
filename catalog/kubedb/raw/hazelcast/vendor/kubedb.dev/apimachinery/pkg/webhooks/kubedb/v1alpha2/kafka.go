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

	errors2 "github.com/pkg/errors"
	"gomodules.xyz/pointer"
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

// SetupKafkaWebhookWithManager registers the webhook for Kafka in the manager.
func SetupKafkaWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Kafka{}).
		WithValidator(&KafkaCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&KafkaCustomWebhook{mgr.GetClient()}).
		Complete()
}

type KafkaCustomWebhook struct {
	DefaultClient client.Client
}

// log is for logging in this package.
var kafkalog = logf.Log.WithName("kafka-resource")

var _ webhook.CustomDefaulter = &KafkaCustomWebhook{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *KafkaCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Kafka)
	if !ok {
		return fmt.Errorf("expected an Kafka object but got %T", obj)
	}
	kafkalog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &KafkaCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *KafkaCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Kafka)
	if !ok {
		return nil, fmt.Errorf("expected an Kafka object but got %T", obj)
	}
	kafkalog.Info("validate create", "name", db.Name)
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *KafkaCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Kafka)
	if !ok {
		return nil, fmt.Errorf("expected an Kafka object but got %T", newObj)
	}
	kafkalog.Info("validate update", "name", db.Name)
	return nil, w.ValidateCreateOrUpdate(db)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *KafkaCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Kafka)
	if !ok {
		return nil, fmt.Errorf("expected an Kafka object but got %T", obj)
	}
	kafkalog.Info("validate delete", "name", db.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("teminationPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kafka.kubedb.com", Kind: "Kafka"}, db.Name, allErr)
	}
	return nil, nil
}

func (w *KafkaCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Kafka) error {
	var allErr field.ErrorList
	// TODO(user): fill in your validation logic upon object creation.
	if db.Spec.EnableSSL {
		if db.Spec.TLS == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("enableSSL"),
				db.Name,
				".spec.tls can't be nil, if .spec.enableSSL is true"))
		}
	} else {
		if db.Spec.TLS != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("enableSSL"),
				db.Name,
				".spec.tls must be nil, if .spec.enableSSL is disabled"))
		}
	}
	if db.Spec.Topology != nil {
		if db.Spec.Topology.Controller == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("controller"),
				db.Name,
				".spec.topology.controller can't be empty in topology cluster"))
		}
		if db.Spec.Topology.Broker == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("broker"),
				db.Name,
				".spec.topology.broker can't be empty in topology cluster"))
		}

		if db.Spec.Replicas != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"doesn't support spec.replicas when spec.topology is set"))
		}
		if db.Spec.Storage != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("broker"),
				db.Name,
				"doesn't support spec.storage when spec.topology is set"))
		}
		if db.Spec.PodTemplate.Spec.Containers[0].Resources.Size() != 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("resources"),
				db.Name,
				"doesn't support spec.podTemplate.spec.resources when spec.topology is set"))
		}

		if *db.Spec.Topology.Controller.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("controller").Child("replicas"),
				db.Name,
				"number of replicas can not be less be 0 or less"))
		}

		if *db.Spec.Topology.Broker.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("broker").Child("replicas"),
				db.Name,
				"number of replicas can not be 0 or less"))
		}

		// validate that multiple nodes don't have same suffixes
		err := w.validateNodeSuffix(db.Spec.Topology)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology"),
				db.Name,
				err.Error()))
		}

		// validate that node replicas are not 0 or negative
		err = w.validateNodeReplicas(db.Spec.Topology)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology"),
				db.Name,
				err.Error()))
		}
	} else {
		// number of replicas can not be 0 or less
		if db.Spec.Replicas != nil && *db.Spec.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas can not be 0 or less"))
		}
	}

	if db.Spec.Halted && db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("halted"),
			db.Name,
			`can't halt if deletionPolicy is set to "DoNotTerminate"`))
	}

	err := w.validateVersion(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			err.Error()))
	}

	err = w.validateVolumes(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			db.Name,
			err.Error()))
	}

	err = w.validateVolumesMountPaths(&db.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumeMounts"),
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
		if db.Spec.StorageType == olddbapi.StorageTypeEphemeral && db.Spec.DeletionPolicy == olddbapi.DeletionPolicyHalt {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deletionPolicy"),
				db.Name,
				`'spec.deletionPolicy: Halt' can not be used for 'Ephemeral' storage`))
		}
	}

	if len(allErr) == 0 {
		return nil
	}
	return apierrors.NewInvalid(schema.GroupKind{Group: "kafka.kubedb.com", Kind: "Kafka"}, db.Name, allErr)
}

func (w *KafkaCustomWebhook) validateVersion(db *olddbapi.Kafka) error {
	kfVersion := &catalog.KafkaVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: db.Spec.Version}, kfVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

func (w *KafkaCustomWebhook) validateNodeSuffix(topology *olddbapi.KafkaClusterTopology) error {
	tMap := topology.ToMap()
	names := make(map[string]bool)
	for _, value := range tMap {
		names[value.Suffix] = true
	}
	if len(tMap) != len(names) {
		return errors.New("two or more node cannot have same suffix")
	}
	return nil
}

func (w *KafkaCustomWebhook) validateNodeReplicas(topology *olddbapi.KafkaClusterTopology) error {
	tMap := topology.ToMap()
	for key, node := range tMap {
		if pointer.Int32(node.Replicas) <= 0 {
			return errors2.Errorf("replicas for node role %s must be alteast 1", string(key))
		}
	}
	return nil
}

var kafkaReservedVolumes = []string{
	kubedb.KafkaVolumeData,
	kubedb.KafkaVolumeConfig,
	kubedb.KafkaVolumeTempConfig,
}

func (w *KafkaCustomWebhook) validateVolumes(db *olddbapi.Kafka) error {
	if db.Spec.PodTemplate.Spec.Volumes == nil {
		return nil
	}
	rsv := make([]string, len(kafkaReservedVolumes))
	copy(rsv, kafkaReservedVolumes)
	if db.Spec.TLS != nil && db.Spec.TLS.Certificates != nil {
		for _, c := range db.Spec.TLS.Certificates {
			rsv = append(rsv, db.CertSecretVolumeName(olddbapi.KafkaCertificateAlias(c.Alias)))
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

var kafkaReservedVolumeMountPaths = []string{
	kubedb.KafkaConfigDir,
	kubedb.KafkaTempConfigDir,
	kubedb.KafkaDataDir,
	kubedb.KafkaMetaDataDir,
	kubedb.KafkaCertDir,
}

func (w *KafkaCustomWebhook) validateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range kafkaReservedVolumeMountPaths {
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

	for _, rvmp := range kafkaReservedVolumeMountPaths {
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
