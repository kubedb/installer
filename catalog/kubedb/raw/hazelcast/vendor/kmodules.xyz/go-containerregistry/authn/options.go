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

package authn

import (
	"kmodules.xyz/client-go/meta"

	"github.com/google/go-containerregistry/pkg/authn/k8schain"
	"github.com/spf13/pflag"
)

var optsFromFlags = k8schain.Options{
	Namespace:          meta.PodNamespace(),
	ServiceAccountName: meta.PodServiceAccount(),
}

func AddKubeChainOptionsFlags(fs *pflag.FlagSet) {
	if fs == nil {
		fs = pflag.CommandLine
	}
	fs.StringSliceVar(&optsFromFlags.ImagePullSecrets, "image-pull-secrets", optsFromFlags.ImagePullSecrets, "Name of image pull secret")
}

func isOptionsSet() bool {
	return optsFromFlags.ServiceAccountName != "" || len(optsFromFlags.ImagePullSecrets) > 0
}
