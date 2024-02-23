/*
Copyright AppsCode Inc. and Contributors.

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

package trivy

type SingleReport struct {
	SchemaVersion int           `json:"schemaVersion" tv:"SchemaVersion"`
	ArtifactName  string        `json:"artifactName" tv:"ArtifactName"`
	ArtifactType  string        `json:"artifactType" tv:"ArtifactType"`
	Metadata      ImageMetadata `json:"metadata" tv:"Metadata"`
	Results       []Result      `json:"results" tv:"Results"`
}

type ImageMetadata struct {
	Os          ImageOS     `json:"os" tv:"OS"`
	ImageID     string      `json:"imageID" tv:"ImageID"`
	DiffIDs     []string    `json:"diffIDs" tv:"DiffIDs"`
	RepoTags    []string    `json:"repoTags" tv:"RepoTags"`
	RepoDigests []string    `json:"repoDigests" tv:"RepoDigests"`
	ImageConfig ImageConfig `json:"imageConfig" tv:"ImageConfig"`
}

type ImageOS struct {
	Family string `json:"family" tv:"Family"`
	Name   string `json:"name" tv:"Name"`
}

type ImageConfig struct {
	Architecture  string             `json:"architecture" tv:"architecture"`
	Author        string             `json:"author,omitempty" tv:"author,omitempty"`
	Container     string             `json:"container,omitempty" tv:"container,omitempty"`
	Created       Time               `json:"created" tv:"created"`
	DockerVersion string             `json:"dockerVersion,omitempty" tv:"docker_version,omitempty"`
	History       []ImageHistory     `json:"history" tv:"history"`
	Os            string             `json:"os" tv:"os"`
	Rootfs        ImageRootfs        `json:"rootfs" tv:"rootfs"`
	Config        ImageRuntimeConfig `json:"config" tv:"config"`
}

type ImageHistory struct {
	Created    Time   `json:"created" tv:"created"`
	CreatedBy  string `json:"createdBy" tv:"created_by"`
	EmptyLayer bool   `json:"emptyLayer,omitempty" tv:"empty_layer,omitempty"`
	Comment    string `json:"comment,omitempty" tv:"comment,omitempty"`
}

type ImageRootfs struct {
	Type    string   `json:"type" tv:"type"`
	DiffIds []string `json:"diffIDs" tv:"diff_ids"`
}

type ImageRuntimeConfig struct {
	Cmd         []string          `json:"cmd" tv:"Cmd"`
	Env         []string          `json:"env,omitempty" tv:"Env,omitempty"`
	Image       string            `json:"image,omitempty" tv:"Image,omitempty"`
	Entrypoint  []string          `json:"entrypoint,omitempty" tv:"Entrypoint,omitempty"`
	Labels      map[string]string `json:"labels,omitempty" tv:"Labels,omitempty"`
	ArgsEscaped bool              `json:"argsEscaped,omitempty" tv:"ArgsEscaped,omitempty"`
	StopSignal  string            `json:"stopSignal,omitempty" tv:"StopSignal,omitempty"`
}

type VulnerabilityLayer struct {
	Digest string `json:"digest,omitempty" tv:"Digest,omitempty"`
	DiffID string `json:"diffID" tv:"DiffID"`
}

type VulnerabilityDataSource struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
	URL  string `json:"URL"`
}

type CVSSScore struct {
	V2Vector string  `json:"v2Vector,omitempty" tv:"V2Vector,omitempty"`
	V3Vector string  `json:"v3Vector,omitempty" tv:"V3Vector,omitempty"`
	V2Score  float64 `json:"v2Score,omitempty" tv:"V2Score,omitempty"`
	V3Score  float64 `json:"v3Score,omitempty" tv:"V3Score,omitempty"`
}

type Vulnerability struct {
	VulnerabilityID  string                  `json:"vulnerabilityID" tv:"VulnerabilityID"`
	PkgName          string                  `json:"pkgName" tv:"PkgName"`
	PkgID            string                  `json:"pkgID,omitempty" tv:"PkgID,omitempty"`
	InstalledVersion string                  `json:"-" tv:"InstalledVersion"`
	Layer            VulnerabilityLayer      `json:"-" tv:"Layer"`
	SeveritySource   string                  `json:"severitySource" tv:"SeveritySource"`
	PrimaryURL       string                  `json:"primaryURL" tv:"PrimaryURL"`
	DataSource       VulnerabilityDataSource `json:"dataSource" tv:"DataSource"`
	Title            string                  `json:"title,omitempty" tv:"Title,omitempty"`
	Description      string                  `json:"description" tv:"Description"`
	Severity         string                  `json:"severity" tv:"Severity"`
	CweIDs           []string                `json:"cweIDs,omitempty" tv:"CweIDs,omitempty"`
	Cvss             map[string]CVSSScore    `json:"cvss,omitempty" tv:"CVSS,omitempty"`
	References       []string                `json:"references" tv:"References"`
	PublishedDate    *Time                   `json:"publishedDate,omitempty" tv:"PublishedDate,omitempty"`
	LastModifiedDate *Time                   `json:"lastModifiedDate,omitempty" tv:"LastModifiedDate,omitempty"`
	FixedVersion     string                  `json:"fixedVersion,omitempty" tv:"FixedVersion,omitempty"`
}

type VulnerabilityInfo struct {
	VulnerabilityID string `json:"vulnerabilityID"`
	Title           string `json:"title,omitempty"`
	Severity        string `json:"severity"`
	PrimaryURL      string `json:"primaryURL"`
	Occurrence      int    `json:"occurrence"`

	// +optional
	Results []ImageResult          `json:"results,omitempty"  tv:"-"`
	R       map[string]ImageResult `json:"-" tv:"-"`
}

type ImageResult struct {
	Image   string   `json:"image,omitempty"`
	Targets []Target `json:"targets,omitempty"`
}

type Target struct {
	Layer            *VulnerabilityLayer `json:"layer,omitempty"`
	InstalledVersion string              `json:"installedVersion,omitempty"`
	Target           string              `json:"target"`
	Class            string              `json:"class"`
	Type             string              `json:"type"`
}

type Result struct {
	Target          string          `json:"target" tv:"Target"`
	Class           string          `json:"class" tv:"Class"`
	Type            string          `json:"type" tv:"Type"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities,omitempty" tv:"Vulnerabilities,omitempty"`
}
