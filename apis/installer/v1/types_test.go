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

package v1_test

import (
	"os"
	"testing"

	v1 "kubedb.dev/installer/apis/installer/v1"

	schemachecker "kmodules.xyz/schema-checker"
)

func TestDefaultValues(t *testing.T) {
	checker := schemachecker.New(os.DirFS("../../.."),
		schemachecker.TestCase{Obj: v1.KubedbAutoscalerSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbCatalogSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbCrdManagerSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbDashboardSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbGitopsSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbKubestashCatalogSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbMigratorSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbOpsManagerSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbProviderAwsSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbProviderAzureSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbProviderGcpSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbProvisionerSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbSchemaManagerSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbUiServerSpec{}},
		schemachecker.TestCase{Obj: v1.KubedbWebhookServerSpec{}},
		schemachecker.TestCase{Obj: v1.PrepareClusterSpec{}},
	)
	checker.TestAll(t)
}
