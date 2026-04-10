/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"fmt"
	"strings"

	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	certlib "kubedb.dev/hazelcast/pkg/lib/cert"
	"kubedb.dev/hazelcast/util"

	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kutil "kmodules.xyz/client-go"
	app_util "kmodules.xyz/client-go/apps/v1"
	coreutil "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/go-containerregistry/authn"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

func (rs *ReconcileState) EnsureStatefulSet() error {
	volumes := rs.getVolumes()
	pvc := rs.getPVC()
	podTemplate := rs.db.Spec.PodTemplate

	containers := rs.getContainers()
	initContainers := rs.getInitContainers()

	// Create of Patch the sts with given opts
	sts := &apps.StatefulSet{
		ObjectMeta: meta.ObjectMeta{
			Name:      rs.db.StatefulSetName(),
			Namespace: rs.db.Namespace,
		},
	}
	sts, v, err := app_util.CreateOrPatchStatefulSet(rs.ctx, rs.Client, sts.ObjectMeta, func(in *apps.StatefulSet) *apps.StatefulSet {
		in.Labels = rs.db.OffshootLabels()
		in.Annotations = podTemplate.Controller.Annotations
		in.Spec.Template.Labels = rs.db.PodLabels()
		in.Spec.Selector = &meta.LabelSelector{
			MatchLabels: rs.db.OffshootSelectors(),
		}
		coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())

		in.Spec.Replicas = rs.db.Spec.Replicas
		in.Spec.ServiceName = rs.db.GoverningServiceName()
		in.Spec.Template.Spec.ServiceAccountName = rs.db.Name
		if rs.db.Spec.PodTemplate.Spec.ServiceAccountName != "" {
			in.Spec.Template.Spec.ServiceAccountName = rs.db.Spec.PodTemplate.Spec.ServiceAccountName
		}
		in.Spec.Template.Spec.InitContainers = coreutil.UpsertContainers(in.Spec.Template.Spec.InitContainers, initContainers)
		in.Spec.Template.Spec.Containers = coreutil.UpsertContainers(in.Spec.Template.Spec.Containers, containers)
		in.Spec.Template.Spec.Volumes = coreutil.UpsertVolume(in.Spec.Template.Spec.Volumes, volumes...)
		if pvc != nil {
			in.Spec.VolumeClaimTemplates = coreutil.UpsertVolumeClaim(in.Spec.VolumeClaimTemplates, *pvc)
		}
		in.Spec.Template.Spec.NodeSelector = rs.db.Spec.PodTemplate.Spec.NodeSelector
		if rs.db.Spec.PodTemplate.Spec.SchedulerName != "" {
			in.Spec.Template.Spec.SchedulerName = podTemplate.Spec.SchedulerName
		}
		in.Spec.Template.Spec.Tolerations = podTemplate.Spec.Tolerations
		in.Spec.Template.Spec.ImagePullSecrets = podTemplate.Spec.ImagePullSecrets
		in.Spec.Template.Spec.PriorityClassName = podTemplate.Spec.PriorityClassName
		in.Spec.Template.Spec.Priority = podTemplate.Spec.Priority
		in.Spec.Template.Spec.HostNetwork = podTemplate.Spec.HostNetwork
		in.Spec.Template.Spec.HostPID = podTemplate.Spec.HostPID
		in.Spec.Template.Spec.HostIPC = podTemplate.Spec.HostIPC
		in.Spec.Template.Spec.SecurityContext = podTemplate.Spec.SecurityContext
		if rs.db.Spec.PodTemplate.Spec.DNSPolicy != "" {
			in.Spec.Template.Spec.DNSPolicy = podTemplate.Spec.DNSPolicy
		}
		if rs.db.Spec.PodTemplate.Spec.TerminationGracePeriodSeconds != nil {
			in.Spec.Template.Spec.TerminationGracePeriodSeconds = podTemplate.Spec.TerminationGracePeriodSeconds
		}
		if rs.db.Spec.PodTemplate.Spec.RuntimeClassName != nil {
			in.Spec.Template.Spec.RuntimeClassName = podTemplate.Spec.RuntimeClassName
		}
		if rs.db.Spec.PodTemplate.Spec.EnableServiceLinks != nil {
			in.Spec.Template.Spec.EnableServiceLinks = podTemplate.Spec.EnableServiceLinks
		}
		// PetSet update strategy is set default to "OnDelete"
		in.Spec.UpdateStrategy = apps.StatefulSetUpdateStrategy{
			Type: apps.OnDeleteStatefulSetStrategyType,
		}
		in = rs.ensureContainers(in)

		return in
	}, meta.PatchOptions{})
	if err != nil {
		return err
	}

	if v == kutil.VerbCreated {
		rs.log.Info(fmt.Sprintf("StatefulSet %s/%s created", sts.Namespace, sts.Name))
	}

	// ensure pdb
	if err := rs.SyncStatefulSetSetPodDisruptionBudget(sts); err != nil {
		rs.log.Error(err, "Failed to create/patch PodDisruptionBudget")
		return err
	}

	return nil
}

func (rs *ReconcileState) getEnv() []core.EnvVar {
	var envList []core.EnvVar
	opts := strings.Join(rs.db.Spec.JavaOpts, " ")

	envList = coreutil.UpsertEnvVars(envList, []core.EnvVar{
		{
			Name: "HZ_LICENSEKEY",
			ValueFrom: &core.EnvVarSource{
				SecretKeyRef: &core.SecretKeySelector{
					LocalObjectReference: core.LocalObjectReference{
						Name: rs.db.Spec.LicenseSecret.Name,
					},
					Key: kubedb.HazelcastLicenseKey,
				},
			},
		},
		{
			Name: "POD_NAME",
			ValueFrom: &core.EnvVarSource{
				FieldRef: &core.ObjectFieldSelector{
					FieldPath:  "metadata.name",
					APIVersion: "v1",
				},
			},
		},
		{
			Name:  "SERVICE_NAME",
			Value: rs.db.Name,
		},
		{
			Name: "POD_NAMESPACE",
			ValueFrom: &core.EnvVarSource{
				FieldRef: &core.ObjectFieldSelector{
					FieldPath: "metadata.namespace",
				},
			},
		},
		{
			Name: "JAVA_OPTS",
			Value: fmt.Sprintf(
				"-Dhazelcast.config=/data/hazelcast/hazelcast.yaml "+
					"-Dhazelcast.client.config=/data/hazelcast/hazelcast-client.yaml "+
					"-DserviceName=%s "+
					"-Dnamespace=%s "+
					"-Dhazelcast.persistence=true "+
					"-Dhazelcast.stale.join.prevention.duration.seconds=5 "+
					"-Dhz.jet.enabled=true "+
					"-Dhazelcast.shutdownhook.policy=GRACEFUL "+
					"-Dhazelcast.shutdownhook.enabled=true "+
					"-Dhazelcast.graceful.shutdown.max.wait=600 "+
					"-Dhazelcast.cluster.version.auto.upgrade.enabled=true "+
					"-Dhazelcast.network.join.kubernetes.use-node-name-as-hostname=true "+
					"%s",
				rs.db.GoverningServiceName(),
				rs.db.Namespace,
				opts,
			),
		},

		{
			Name:  "LOGGING_LEVEL",
			Value: "INFO",
		},
		{
			Name:  "LOGGING_LEVEL_ROOT",
			Value: "INFO",
		},
		{
			Name:  "PROMETHEUS_PORT",
			Value: "56790",
		},
		{
			Name:  "SCHEME",
			Value: rs.db.GetConnectionScheme(),
		},
	}...)

	if !rs.db.Spec.DisableSecurity {
		secret := &core.Secret{}
		err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
			Name:      rs.db.GetAuthSecretName(),
			Namespace: rs.db.Namespace,
		}, secret)
		if err != nil {
			return envList
		}
		envList = coreutil.UpsertEnvVars(envList, []core.EnvVar{
			{
				Name:  "USERNAME",
				Value: string(secret.Data[core.BasicAuthUsernameKey]),
			},
			{
				Name:  "PASSWORD",
				Value: string(secret.Data[core.BasicAuthPasswordKey]),
			},
		}...)

	}

	return envList
}

func (rs *ReconcileState) ensureContainers(ps *apps.StatefulSet) *apps.StatefulSet {
	dbContainer := coreutil.GetContainerByName(ps.Spec.Template.Spec.Containers, kubedb.HazelcastContainerName)
	if rs.db.Spec.Configuration != nil && rs.db.Spec.Configuration.SecretName == "" {
		ps.Spec.Template.Spec.Volumes = coreutil.EnsureVolumeDeleted(ps.Spec.Template.Spec.Volumes, kubedb.HazelcastCustomConfigVolume)
		if dbContainer != nil {
			dbContainer.VolumeMounts = coreutil.EnsureVolumeMountDeleted(dbContainer.VolumeMounts, kubedb.HazelcastCustomConfigVolume)
		}
	}
	if !rs.db.Spec.EnableSSL {
		ps.Spec.Template.Spec.Volumes = coreutil.EnsureVolumeDeleted(ps.Spec.Template.Spec.Volumes, rs.db.CertSecretVolumeName(api.HazelcastServerCert))
		if dbContainer != nil {
			dbContainer.VolumeMounts = coreutil.EnsureVolumeMountDeleted(dbContainer.VolumeMounts, rs.db.CertSecretVolumeName(api.HazelcastServerCert))
		}
		ps.Spec.Template.Spec.Volumes = coreutil.EnsureVolumeDeleted(ps.Spec.Template.Spec.Volumes, rs.db.CertSecretVolumeName(api.HazelcastClientCert))
		if dbContainer != nil {
			dbContainer.VolumeMounts = coreutil.EnsureVolumeMountDeleted(dbContainer.VolumeMounts, rs.db.CertSecretVolumeName(api.HazelcastClientCert))
		}
	}
	return ps
}

func (rs *ReconcileState) getContainers() []core.Container {
	image, err := authn.ImageWithDigest(rs.Client, rs.version.Spec.DB.Image, util.K8sChainOpts(rs.db))
	if err != nil {
		rs.log.Error(err, "Failed to get image with digest")
		return nil
	}
	containerTemplate := coreutil.GetContainerByName(rs.db.Spec.PodTemplate.Spec.Containers, kubedb.HazelcastContainerName)
	dbContainer := core.Container{
		Name:         kubedb.HazelcastContainerName,
		Image:        image,
		Ports:        rs.getDBContainerPorts(),
		Env:          rs.getEnv(),
		VolumeMounts: rs.getDBVolumeMounts(),
	}
	if containerTemplate != nil {
		dbContainer = coreutil.MergeContainer(dbContainer, *containerTemplate)
	}
	containers := coreutil.UpsertContainers(rs.db.Spec.PodTemplate.Spec.Containers, []core.Container{dbContainer})
	return containers
}

func (rs *ReconcileState) getInitContainers() []core.Container {
	initImage, err := authn.ImageWithDigest(rs.Client, rs.version.Spec.InitContainer.Image, util.K8sChainOpts(rs.db))
	if err != nil {
		rs.log.Error(err, "Failed to get image with digest")
		return nil
	}
	containerTemplate := coreutil.GetContainerByName(rs.db.Spec.PodTemplate.Spec.InitContainers, kubedb.HazelcastInitContainerName)
	initContainer := core.Container{
		Name:         kubedb.HazelcastInitContainerName,
		Image:        initImage,
		VolumeMounts: rs.getInitContainerVolumeMounts(),
		Env:          rs.getEnv(),
		Resources:    core.ResourceRequirements{},
	}
	if containerTemplate != nil {
		initContainer = coreutil.MergeContainer(initContainer, *containerTemplate)
	}
	containers := coreutil.UpsertContainers(rs.db.Spec.PodTemplate.Spec.InitContainers, []core.Container{initContainer})
	return containers
}

func (rs *ReconcileState) getVolumes() []core.Volume {
	// User provided custom volume (if any)
	volumes := ofst.ConvertVolumes(rs.db.Spec.PodTemplate.Spec.Volumes)

	volumes = coreutil.UpsertVolume(volumes, []core.Volume{
		{
			Name: kubedb.HazelcastDefaultConfigVolume,
			VolumeSource: core.VolumeSource{
				Secret: &core.SecretVolumeSource{
					SecretName: rs.db.ConfigSecretName(),
				},
			},
		},
		{
			Name: kubedb.HazelcastConfigVolume,
			VolumeSource: core.VolumeSource{
				EmptyDir: &core.EmptyDirVolumeSource{},
			},
		},
	}...)

	if rs.db.Spec.Configuration != nil && rs.db.Spec.Configuration.SecretName != "" {
		configSecret := &core.Secret{}
		// Get the configSecret,
		// If not found return error
		if err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
			Name:      rs.db.Spec.Configuration.SecretName,
			Namespace: rs.db.Namespace,
		}, configSecret); err != nil {
			if kerr.IsNotFound(err) {
				rs.log.Error(err, fmt.Sprintf("Secret: %s/%s not found", rs.db.Namespace, rs.db.Spec.Configuration.SecretName))
				return volumes
			} else {
				rs.log.Error(err, fmt.Sprintf("Failed to get configSecret %s/%s", rs.db.Namespace, rs.db.Spec.Configuration.SecretName))
				return volumes
			}
		}

		volumes = coreutil.UpsertVolume(volumes, core.Volume{
			Name: kubedb.HazelcastCustomConfigVolume,
			VolumeSource: core.VolumeSource{
				Secret: &core.SecretVolumeSource{
					SecretName: rs.db.Spec.Configuration.SecretName,
				},
			},
		})
	}

	if rs.db.Spec.EnableSSL {
		volumes = coreutil.UpsertVolume(volumes, core.Volume{
			Name: rs.db.CertSecretVolumeName(api.HazelcastServerCert),
			VolumeSource: core.VolumeSource{
				Secret: &core.SecretVolumeSource{
					SecretName: rs.db.GetCertSecretName(api.HazelcastServerCert),
					Items: []core.KeyToPath{
						{
							Key:  certlib.CACert,
							Path: certlib.CACert,
						},
						{
							Key:  certlib.TLSCert,
							Path: certlib.TLSCert,
						},
						{
							Key:  certlib.TLSKey,
							Path: certlib.TLSKey,
						},
						{
							Key:  certlib.KeystoreKey,
							Path: certlib.KeystoreKey,
						},
						{
							Key:  certlib.TruststoreKey,
							Path: certlib.TruststoreKey,
						},
					},
				},
			},
		})

		volumes = coreutil.UpsertVolume(volumes, core.Volume{
			Name: rs.db.CertSecretVolumeName(api.HazelcastClientCert),
			VolumeSource: core.VolumeSource{
				Secret: &core.SecretVolumeSource{
					SecretName: rs.db.GetCertSecretName(api.HazelcastClientCert),
					Items: []core.KeyToPath{
						{
							Key:  certlib.CACert,
							Path: certlib.CACert,
						},
						{
							Key:  certlib.TLSCert,
							Path: certlib.TLSCert,
						},
						{
							Key:  certlib.TLSKey,
							Path: certlib.TLSKey,
						},
						{
							Key:  certlib.KeystoreKey,
							Path: certlib.KeystoreKey,
						},
						{
							Key:  certlib.TruststoreKey,
							Path: certlib.TruststoreKey,
						},
					},
				},
			},
		})

	}

	if rs.db.Spec.StorageType == api.StorageTypeEphemeral {
		ed := core.EmptyDirVolumeSource{}
		if rs.db.Spec.Storage != nil {
			if sz, found := rs.db.Spec.Storage.Resources.Requests[core.ResourceStorage]; found {
				ed.SizeLimit = &sz
			}
		}
		volumes = coreutil.UpsertVolume(volumes, core.Volume{
			Name: rs.db.PVCName(kubedb.HazelcastVolumeData),
			VolumeSource: core.VolumeSource{
				EmptyDir: &ed,
			},
		})
	}

	return volumes
}

func (rs *ReconcileState) getPVC() *core.PersistentVolumeClaim {
	if rs.db.Spec.StorageType == api.StorageTypeEphemeral {
		return nil
	}
	pvc := &core.PersistentVolumeClaim{
		ObjectMeta: meta.ObjectMeta{
			Name: rs.db.PVCName("data"),
		},
		Spec: *rs.db.Spec.Storage,
	}

	if len(rs.db.Spec.Storage.AccessModes) == 0 {
		pvc.Spec.AccessModes = []core.PersistentVolumeAccessMode{
			core.ReadWriteOnce,
		}
	} else {
		pvc.Spec.AccessModes = rs.db.Spec.Storage.AccessModes
	}

	if rs.db.Spec.Storage.StorageClassName != nil {
		pvc.Spec.StorageClassName = rs.db.Spec.Storage.StorageClassName
		pvc.Annotations = map[string]string{
			"volume.beta.kubernetes.io/storage-class": *rs.db.Spec.Storage.StorageClassName,
		}
	}

	if rs.db.Spec.Storage.Resources.Requests != nil {
		pvc.Spec.Resources.Requests = rs.db.Spec.Storage.Resources.Requests
	}

	return pvc
}

func (rs *ReconcileState) getDBVolumeMounts() []core.VolumeMount {
	var volumeMounts []core.VolumeMount

	container := coreutil.GetContainerByName(rs.db.Spec.PodTemplate.Spec.Containers, kubedb.HazelcastContainerName)
	if container != nil {
		volumeMounts = container.VolumeMounts
	}

	volumeMounts = coreutil.UpsertVolumeMount(volumeMounts, []core.VolumeMount{
		{
			Name:      rs.db.PVCName("data"),
			MountPath: kubedb.HazelcastDataDir,
			ReadOnly:  false,
		},
		{
			Name:      kubedb.HazelcastConfigVolume,
			MountPath: kubedb.HazelcastConfigDir,
		},
	}...)

	if rs.db.Spec.EnableSSL {
		volumeMounts = coreutil.UpsertVolumeMount(volumeMounts, core.VolumeMount{
			Name:      rs.db.CertSecretVolumeName(api.HazelcastServerCert),
			MountPath: kubedb.HazelcastTLSServerMountPath,
		})
		volumeMounts = coreutil.UpsertVolumeMount(volumeMounts, core.VolumeMount{
			Name:      rs.db.CertSecretVolumeName(api.HazelcastClientCert),
			MountPath: kubedb.HazelcastTLSClientMountPath,
		})
	}

	return volumeMounts
}

func (rs *ReconcileState) getInitContainerVolumeMounts() []core.VolumeMount {
	var volumeMounts []core.VolumeMount

	container := coreutil.GetContainerByName(rs.db.Spec.PodTemplate.Spec.InitContainers, kubedb.HazelcastInitContainerName)
	if container != nil {
		volumeMounts = container.VolumeMounts
	}

	volumeMounts = coreutil.UpsertVolumeMount(volumeMounts, []core.VolumeMount{
		{
			Name:      kubedb.HazelcastDefaultConfigVolume,
			MountPath: kubedb.HazelcastTempConfigDir,
		},
		{
			Name:      kubedb.HazelcastConfigVolume,
			MountPath: kubedb.HazelcastConfigDir,
		},
	}...)

	if rs.db.Spec.Configuration != nil && rs.db.Spec.Configuration.SecretName != "" {
		volumeMounts = coreutil.UpsertVolumeMount(volumeMounts, core.VolumeMount{
			Name:      kubedb.HazelcastCustomConfigVolume,
			MountPath: kubedb.HazelcastCustomConfigDir,
		})
	}

	return volumeMounts
}

func (rs *ReconcileState) getDBContainerPorts() []core.ContainerPort {
	dbContainerPorts := []core.ContainerPort{
		{
			Name:          kubedb.HazelcastPortName,
			ContainerPort: kubedb.HazelcastRestPort,
			Protocol:      core.ProtocolTCP,
		},
		{
			Name:          "ui-port",
			Protocol:      core.ProtocolTCP,
			ContainerPort: kubedb.HazelcastUIPort,
		},
	}

	if rs.db.Spec.Monitor != nil {
		dbContainerPorts = append(dbContainerPorts, core.ContainerPort{
			Name:          mona.PrometheusExporterPortName,
			ContainerPort: mona.PrometheusExporterPortNumber,
			Protocol:      core.ProtocolTCP,
		})
	}

	return dbContainerPorts
}
