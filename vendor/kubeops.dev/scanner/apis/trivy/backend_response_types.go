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

type BackendResponse struct {
	Report       SingleReport `json:"report"`
	TrivyVersion Version      `json:"trivyVersion"`
	ImageDetails ImageDetails `json:"image_details"`
	ErrorMessage string       `json:"error_message"`
}

type ImageDetails struct {
	Name string `json:"name,omitempty"`
	// Tag & Digest is optional field. One of these fields may not present
	// +optional
	Tag string `json:"tag,omitempty"`
	// +optional
	Digest string `json:"digest,omitempty"`
	// +kubebuilder:default="Public"
	Visibility ImageVisibility `json:"visibility,omitempty"`
}

// +kubebuilder:validation:Enum=Public;Private;Unknown
type ImageVisibility string

const (
	ImageVisibilityPublic  ImageVisibility = "Public"
	ImageVisibilityPrivate ImageVisibility = "Private"
	ImageVisibilityUnknown ImageVisibility = "Unknown"
)
