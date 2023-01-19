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
)

type ImageRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type Container struct {
	ImageRef `json:",inline"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type WebHookSpec struct {
	UseKubeapiserverFqdnForAks bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                HealthcheckSpec `json:"healthcheck"`
}

type ServingCerts struct {
	Generate bool `json:"generate"`
	// +optional
	CaCrt string `json:"caCrt"`
	// +optional
	ServerCrt string `json:"serverCrt"`
	// +optional
	ServerKey string `json:"serverKey"`
}

type HealthcheckSpec struct {
	// +optional
	Enabled   bool `json:"enabled"`
	ProbePort int  `json:"probePort"`
}

// +kubebuilder:validation:Enum=prometheus.io;prometheus.io/operator;prometheus.io/builtin
type MonitoringAgent string

type Monitoring struct {
	Agent          MonitoringAgent       `json:"agent"`
	BindPort       int                   `json:"bindPort"`
	ServiceMonitor *ServiceMonitorLabels `json:"serviceMonitor"`
}

type ServiceMonitorLabels struct {
	// +optional
	Labels map[string]string `json:"labels"`
}

type EASSpec struct {
	GroupPriorityMinimum       int32              `json:"groupPriorityMinimum"`
	VersionPriority            int32              `json:"versionPriority"`
	UseKubeapiserverFqdnForAks bool               `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                EASHealthcheckSpec `json:"healthcheck"`
	ServingCerts               ServingCerts       `json:"servingCerts"`
}

type EASHealthcheckSpec struct {
	// +optional
	Enabled bool `json:"enabled"`
}

type EASMonitoring struct {
	Agent          MonitoringAgent      `json:"agent"`
	ServiceMonitor ServiceMonitorLabels `json:"serviceMonitor"`
}
