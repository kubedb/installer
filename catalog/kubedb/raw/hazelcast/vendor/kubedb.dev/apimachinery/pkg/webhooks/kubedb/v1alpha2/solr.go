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
	"slices"
	"strings"

	catalog "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	olddbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	"github.com/coreos/go-semver/semver"
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

// SetupSolrWebhookWithManager registers the webhook for Solr in the manager.
func SetupSolrWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&olddbapi.Solr{}).
		WithValidator(&SolrCustomWebhook{mgr.GetClient()}).
		WithDefaulter(&SolrCustomWebhook{mgr.GetClient()}).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-solr-kubedb-com-v1alpha1-solr,mutating=true,failurePolicy=fail,sideEffects=None,groups=kubedb.com,resources=solr,verbs=create;update,versions=v1alpha1,name=msolr.kb.io,admissionReviewVersions={v1,v1beta1}

// +kubebuilder:object:generate=false
type SolrCustomWebhook struct {
	DefaultClient client.Client
}

var _ webhook.CustomDefaulter = &SolrCustomWebhook{}

// log is for logging in this package.
var solrlog = logf.Log.WithName("solr-resource")

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (w *SolrCustomWebhook) Default(ctx context.Context, obj runtime.Object) error {
	db, ok := obj.(*olddbapi.Solr)
	if !ok {
		return fmt.Errorf("expected an Solr object but got %T", obj)
	}

	solrlog.Info("default", "name", db.Name)

	db.SetDefaults(w.DefaultClient)
	return nil
}

var _ webhook.CustomValidator = &SolrCustomWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (w *SolrCustomWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Solr)
	if !ok {
		return nil, fmt.Errorf("expected an Solr object but got %T", obj)
	}

	solrlog.Info("validate create", "name", db.Name)
	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}
	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Solr"}, db.Name, allErr)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (w *SolrCustomWebhook) ValidateUpdate(ctx context.Context, old, newObj runtime.Object) (admission.Warnings, error) {
	db, ok := newObj.(*olddbapi.Solr)
	if !ok {
		return nil, fmt.Errorf("expected an Solr object but got %T", newObj)
	}
	solrlog.Info("validate update", "name", db.Name)

	allErr := w.ValidateCreateOrUpdate(db)
	if len(allErr) == 0 {
		return nil, nil
	}

	return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Solr"}, db.Name, allErr)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (w *SolrCustomWebhook) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	db, ok := obj.(*olddbapi.Solr)
	if !ok {
		return nil, fmt.Errorf("expected an Solr object but got %T", obj)
	}
	solrlog.Info("validate delete", "name", db.Name)

	var allErr field.ErrorList
	if db.Spec.DeletionPolicy == olddbapi.DeletionPolicyDoNotTerminate {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("deletionPolicy"),
			db.Name,
			"Can not delete as terminationPolicy is set to \"DoNotTerminate\""))
		return nil, apierrors.NewInvalid(schema.GroupKind{Group: "kubedb.com", Kind: "Solr"}, db.Name, allErr)
	}
	return nil, nil
}

var solrReservedVolumes = []string{
	kubedb.SolrVolumeConfig,
	kubedb.SolrVolumeDefaultConfig,
	kubedb.SolrVolumeCustomConfig,
	kubedb.SolrVolumeAuthConfig,
}

var solrReservedVolumeMountPaths = []string{
	kubedb.SolrHomeDir,
	kubedb.SolrDataDir,
	kubedb.SolrCustomConfigDir,
	kubedb.SolrSecurityConfigDir,
	kubedb.SolrTempConfigDir,
}

var solrAvailableModules = []string{
	"analysis-extras", "extraction", "hdfs", "langid", "prometheus-exporter", "sql",
	"analytics", "gcs-repository", "jaegertracer-configurator", "ltr", "s3-repository",
	"clustering", "hadoop-auth", "jwt-auth", "opentelemetry", "scripting",
}

func (w *SolrCustomWebhook) ValidateCreateOrUpdate(db *olddbapi.Solr) field.ErrorList {
	var allErr field.ErrorList

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

	if db.Spec.Version == "" {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
			db.Name,
			"spec.version' is missing"))
	} else {
		err := w.solrValidateVersion(db)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("version"),
				db.Name,
				err.Error()))
		}
	}

	version := semver.New(db.Spec.Version)
	if version.Major == 8 && db.Spec.Topology != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("enableSSL"),
			db.Name,
			".spec.topology not supported for version 8"))
	}

	err := solrValidateModules(db)
	if err != nil {
		allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("solrmodules"),
			db.Name,
			err.Error()))
	}

	if db.Spec.Topology == nil {
		if db.Spec.Replicas != nil && *db.Spec.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("replicas"),
				db.Name,
				"number of replicas can not be less be 0 or less"))
		}
		err := solrValidateVolumes(&db.Spec.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = solrValidateVolumesMountPaths(&db.Spec.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("podTemplate").Child("spec").Child("containers"),
				db.Name,
				err.Error()))
		}

	} else {
		if db.Spec.Topology.Data == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("data"),
				db.Name,
				".spec.topology.data can't be empty in cluster mode"))
		}
		if db.Spec.Topology.Data.Replicas != nil && *db.Spec.Topology.Data.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("data").Child("replicas"),
				db.Name,
				"number of replicas can not be less be 0 or less"))
		}
		err := solrValidateVolumes(&db.Spec.Topology.Data.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("data").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = solrValidateVolumesMountPaths(&db.Spec.Topology.Data.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("data").Child("podTemplate").Child("spec").Child("containers"),
				db.Name,
				err.Error()))
		}

		if db.Spec.Topology.Overseer == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("overseer"),
				db.Name,
				".spec.topology.overseer can't be empty in cluster mode"))
		}
		if db.Spec.Topology.Overseer.Replicas != nil && *db.Spec.Topology.Overseer.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("overseer").Child("replicas"),
				db.Name,
				"number of replicas can not be less be 0 or less"))
		}
		err = solrValidateVolumes(&db.Spec.Topology.Overseer.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("overseer").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = solrValidateVolumesMountPaths(&db.Spec.Topology.Overseer.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("overseer").Child("podTemplate").Child("spec").Child("containers"),
				db.Name,
				err.Error()))
		}

		if db.Spec.Topology.Coordinator == nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("coordinator"),
				db.Name,
				".spec.topology.coordinator can't be empty in cluster mode"))
		}
		if db.Spec.Topology.Coordinator.Replicas != nil && *db.Spec.Topology.Coordinator.Replicas <= 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("coordinator").Child("replicas"),
				db.Name,
				"number of replicas can not be less be 0 or less"))
		}
		err = solrValidateVolumes(&db.Spec.Topology.Coordinator.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("coordinator").Child("podTemplate").Child("spec").Child("volumes"),
				db.Name,
				err.Error()))
		}
		err = solrValidateVolumesMountPaths(&db.Spec.Topology.Coordinator.PodTemplate)
		if err != nil {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("topology").Child("coordinator").Child("podTemplate").Child("spec").Child("containers"),
				db.Name,
				err.Error()))
		}
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

	for _, x := range db.Spec.SolrOpts {
		if strings.Count(x, " ") > 0 {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("solropts"),
				db.Name,
				"solropt jvm env variables must not contain space"))
		}
		if x[0] != '-' || x[1] != 'D' {
			allErr = append(allErr, field.Invalid(field.NewPath("spec").Child("solropts"),
				db.Name,
				"solropt jvm env variables must start with -D"))
		}
	}

	if len(allErr) == 0 {
		return nil
	}
	return allErr
}

func (w *SolrCustomWebhook) solrValidateVersion(s *olddbapi.Solr) error {
	slVersion := catalog.SolrVersion{}
	err := w.DefaultClient.Get(context.TODO(), types.NamespacedName{Name: s.Spec.Version}, &slVersion)
	if err != nil {
		return errors.New("version not supported")
	}
	return nil
}

func solrValidateModules(db *olddbapi.Solr) error {
	modules := db.Spec.SolrModules
	for _, mod := range modules {
		fl := slices.Contains(solrAvailableModules, mod)
		if !fl {
			return fmt.Errorf("%s does not exist in available modules", mod)
		}
	}
	return nil
}

func solrValidateVolumes(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Volumes == nil {
		return nil
	}

	for _, rv := range solrReservedVolumes {
		for _, ugv := range podTemplate.Spec.Volumes {
			if ugv.Name == rv {
				return errors.New("Can't use a reserve volume name: " + rv)
			}
		}
	}

	return nil
}

func solrValidateVolumesMountPaths(podTemplate *ofst.PodTemplateSpec) error {
	if podTemplate == nil {
		return nil
	}
	if podTemplate.Spec.Containers == nil {
		return nil
	}

	for _, rvmp := range solrReservedVolumeMountPaths {
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

	for _, rvmp := range solrReservedVolumeMountPaths {
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
