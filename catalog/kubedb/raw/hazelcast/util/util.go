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

package util

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn/k8schain"
	"github.com/mikefarah/yq/v3/pkg/yqlib"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"kubedb.dev/apimachinery/pkg/yq3/yqstrings"
)

func K8sChainOpts(db *api.Hazelcast) *k8schain.Options {
	opts := &k8schain.Options{
		Namespace: db.Namespace,
	}
	if db.Spec.PodTemplate.Spec.ServiceAccountName == "" {
		opts.ServiceAccountName = db.OffshootName()
	} else {
		opts.ServiceAccountName = db.Spec.PodTemplate.Spec.ServiceAccountName
	}
	if db.Spec.PodTemplate.Spec.ImagePullSecrets != nil {
		for _, ims := range db.Spec.PodTemplate.Spec.ImagePullSecrets {
			opts.ImagePullSecrets = append(opts.ImagePullSecrets, ims.Name)
		}
	}
	return opts
}

func GetMergedConfig(file1, file2 map[string]string) (map[string]string, error) {
	MergedConfig := make(map[string]string)
	for fileName, newApplyConfig := range file2 {
		previousApplyConfig := file1[fileName]
		// We need to merge the new apply config with the previous applyConfig.
		mergedApplyConfig, err := yqstrings.Merge(yqlib.OverwriteArrayMergeStrategy, yqlib.OverwriteCommentsMergeStrategy, true, previousApplyConfig, newApplyConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to merge apply configs for file %s with error %v", fileName, err)
		}
		MergedConfig[fileName] = mergedApplyConfig
	}
	return MergedConfig, nil
}
