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

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	ResourceKindPgoutbox = "Pgoutbox"
	ResourcePgoutbox     = "pgoutbox"
	ResourcePgoutboxs    = "pgoutboxs"
)

// Pgoutbox defines the schama for ui server installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=pgoutboxs,singular=pgoutbox,categories={kubeops,appscode}
type Pgoutbox struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PgoutboxSpec `json:"spec,omitempty"`
}

// PgoutboxSpec is the schema for Identity Server values file
type PgoutboxSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string             `json:"fullnameOverride"`
	ReplicaCount     int32              `json:"replicaCount"`
	RegistryFQDN     string             `json:"registryFQDN"`
	Image            HelmImageReference `json:"image"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	Resources   core.ResourceRequirements `json:"resources"`
	Autoscaling AutoscalingSpec           `json:"autoscaling"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodLabels map[string]string `json:"podLabels"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	//+optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	//+optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
	ServiceAccount  HelmServiceAccount    `json:"serviceAccount"`
	Service         HelmServiceSpec       `json:"service"`
	// +optional
	LivenessProbe *core.Probe `json:"livenessProbe"`
	// +optional
	ReadinessProbe *core.Probe        `json:"readinessProbe"`
	Volumes        []core.Volume      `json:"volumes"`
	VolumeMounts   []core.VolumeMount `json:"volumeMounts"`
	Ingress        AppIngress         `json:"ingress"`
	App            PgOutboxConfig     `json:"app"`
}

type HelmImageReference struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy"`
	Tag        string `json:"tag"`
}

type HelmServiceSpec struct {
	Type string `json:"type"`
	Port int    `json:"port"`
}

type AutoscalingSpec struct {
	Enabled     bool `json:"enabled"`
	MinReplicas int  `json:"minReplicas"`
	MaxReplicas int  `json:"maxReplicas"`
	// +optional
	TargetCPUUtilizationPercentage int `json:"targetCPUUtilizationPercentage,omitempty"`
	// +optional
	TargetMemoryUtilizationPercentage int `json:"targetMemoryUtilizationPercentage,omitempty"`
}

type AppIngress struct {
	Enabled     bool              `json:"enabled"`
	ClassName   string            `json:"className"`
	Annotations map[string]string `json:"annotations"`
	Hosts       []IngressHost     `json:"hosts"`
	TLS         []IngressTLS      `json:"tls"`
}

type IngressHost struct {
	Host  string     `json:"host"`
	Paths []HostPath `json:"paths"`
}

type HostPath struct {
	Path     string `json:"path"`
	PathType string `json:"pathType"`
}

type IngressTLS struct {
	SecretName string   `json:"secretName"`
	Hosts      []string `json:"hosts"`
}
type HelmServiceAccount struct {
	Create      bool              `json:"create"`
	Automount   bool              `json:"automount"`
	Annotations map[string]string `json:"annotations"`
	Name        string            `json:"name"`
}

type PgOutboxConfig struct {
	Config           runtime.RawExtension `json:"config"`
	ConfigSecretName string               `json:"configSecretName"`
	NatsSecretName   string               `json:"natsSecretName"`
	NatsMountPath    string               `json:"natsMountPath"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PgoutboxList is a list of Pgoutboxs
type PgoutboxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Pgoutbox CRD objects
	Items []Pgoutbox `json:"items,omitempty"`
}
