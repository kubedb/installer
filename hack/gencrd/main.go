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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"kubedb.dev/installer/apis/installer/install"
	api "kubedb.dev/installer/apis/installer/v1alpha1"

	"github.com/go-openapi/spec"
	"github.com/golang/glog"
	gort "gomodules.xyz/runtime"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kube-openapi/pkg/common"
	"kmodules.xyz/client-go/openapi"
)

func generateSwaggerJson() {
	var (
		Scheme = runtime.NewScheme()
		Codecs = serializer.NewCodecFactory(Scheme)
	)

	install.Install(Scheme)

	apispec, err := openapi.RenderOpenAPISpec(openapi.Config{
		Scheme: Scheme,
		Codecs: Codecs,
		Info: spec.InfoProps{
			Title:   "KubeDB",
			Version: "v0",
			Contact: &spec.ContactInfo{
				Name:  "AppsCode Inc.",
				URL:   "https://appscode.com",
				Email: "hello@appscode.com",
			},
			License: &spec.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
			api.GetOpenAPIDefinitions,
		},
		//nolint:govet
		Resources: []openapi.TypeInfo{
			{api.SchemeGroupVersion, api.ResourceKubeDBOperators, api.ResourceKindKubeDBOperator, false},
		},
	})
	if err != nil {
		glog.Fatal(err)
	}

	filename := gort.GOPath() + "/src/kubedb.dev/installer/api/openapi-spec/swagger.json"
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		glog.Fatal(err)
	}
	err = ioutil.WriteFile(filename, []byte(apispec), 0644)
	if err != nil {
		glog.Fatal(err)
	}
}

func main() {
	generateSwaggerJson()
}
