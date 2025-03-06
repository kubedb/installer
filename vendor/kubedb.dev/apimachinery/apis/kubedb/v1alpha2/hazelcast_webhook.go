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

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ofst "kmodules.xyz/offshoot-api/api/v2"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var hazelcastlog = logf.Log.WithName("hazelcast-resource")

var _ webhook.Defaulter = &Hazelcast{}

func (h *Hazelcast) Default() {
	if h == nil {
		return
	}
	hazelcastlog.Info("default", "name", h.Name)

	h.SetDefaults()
}

var hazelcastReservedVolumes = []string{
	kubedb.HazelcastConfigVolume,
	kubedb.HazelcastDefaultConfigVolume,
	kubedb.HazelcastCustomConfigVolume,
}

var hazelcastReservedVolumeMountPaths = []string{
	kubedb.HazelcastConfigDir,
	kubedb.HazelcastTempConfigDir,
	kubedb.HazelcastCustomConfigDir,
	kubedb.HazelcastDataDir,
}

var _ webhook.Validator = &Hazelcast{}

func (h *Hazelcast) ValidateCreate() (admission.Warnings, error) {
	hazelcastlog.Info("validate create", "name", h.Name)

	allErr := h.ValidateCreateOrUpdate()
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "hazelcast"}, h.Name, allErr)
}

func (h *Hazelcast) ValidateCreateOrUpdate() field.ErrorList {
	var allErr field.ErrorList

	if h.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			h.Name,
			"spec.version' is missing"))
	} else {
		err := hazelcastValidateVersion(h)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				h.Name,
				err.Error()))
		}
	}

	if h.Spec.Replicas != nil && *h.Spec.Replicas < 3 {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
			h.Name,
			"number of replicas can not be less than 3"))
	}
	err := hazelcastValidateVolumes(&h.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
			h.Name,
			err.Error()))
	}
	err = hazelcastValidateVolumesMountPaths(&h.Spec.PodTemplate)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("containers"),
			h.Name,
			err.Error()))
	}

	if h.Spec.StorageType == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
			h.Name,
			"StorageType can not be empty"))
	} else {
		if h.Spec.StorageType != StorageTypeDurable && h.Spec.StorageType != StorageTypeEphemeral {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("storageType"),
				h.Name,
				"StorageType should be either durable or ephemeral"))
		}
	}

	return allErr
}

func hazelcastValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range hazelcastReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserve volume name: " + rv)
			}
		}
	}

	return nil
}

func hazelcastValidateVersion(s *Hazelcast) error {
	hzVersion := catalog.HazelcastVersion{}
	err := DefaultClient.Get(context.TODO(), types.NamespacedName{Name: s.Spec.Version}, &hzVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

func hazelcastValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range hazelcastReservedVolumeMountPaths {
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

	for _, rvmp := range hazelcastReservedVolumeMountPaths {
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

func (h *Hazelcast) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	hazelcastlog.Info("validate update", "name", h.Name)

	_ = old.(*Hazelcast)
	allErr := h.ValidateCreateOrUpdate()

	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "hazelcast"}, h.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (h *Hazelcast) ValidateDelete() (admission.Warnings, error) {
	hazelcastlog.Info("validate delete", "name", h.Name)

	var allErr field.ErrorList
	if h.Spec.DeletionPolicy == DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("terminationPolicy"),
			h.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "hazelcast"}, h.Name, allErr)
	}
	return nil, nil
}
