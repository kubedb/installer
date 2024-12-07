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

package lib

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
	shell "gomodules.xyz/go-sh"
	"kubeops.dev/scanner/apis/trivy"
	"sigs.k8s.io/yaml"
)

// trivy image ubuntu --security-checks vuln --format json --quiet
func Scan(sh *shell.Session, img string) (*trivy.SingleReport, error) {
	args := []any{
		"image",
		img,
		"--scanners", "vuln",
		"--format", "json",
		"--ignore-unfixed",
		// "--quiet",
	}
	out, err := sh.Command("trivy", args...).Output()
	if err != nil {
		return nil, err
	}

	var r trivy.SingleReport
	err = trivy.JSON.Unmarshal(out, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func SummarizeReport(report *trivy.SingleReport) map[string]int {
	riskOccurrence := map[string]int{} // risk -> occurrence

	for _, rpt := range report.Results {
		for _, tv := range rpt.Vulnerabilities {
			riskOccurrence[tv.Severity]++
		}
	}

	return riskOccurrence
}

func ImageDigest(ref string) (string, bool, error) {
	digest, err := crane.Digest(ref, crane.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		if ImageNotFound(err) {
			return "", false, nil
		}
		return "", false, err
	}
	return digest, true, nil
}

func ImageManifest(ref string) (any, bool, error) {
	data, err := crane.Manifest(ref, crane.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		if ImageNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	var obj map[string]any
	if err = yaml.Unmarshal(data, &obj); err != nil {
		return nil, false, err
	}
	if _, ok := obj["manifests"]; ok {
		var mf v1.IndexManifest
		if err := yaml.Unmarshal(data, &mf); err != nil {
			return nil, false, err
		}
		return &mf, true, nil
	} else if _, ok := obj["layers"]; ok {
		var mf v1.Manifest
		if err := yaml.Unmarshal(data, &mf); err != nil {
			return nil, false, err
		}
		return &mf, true, nil
	}
	return nil, false, fmt.Errorf("unknown image manifest format")
}

func ImageNotFound(err error) bool {
	var terr *transport.Error
	if errors.As(err, &terr) {
		return terr.StatusCode == http.StatusNotFound //&& terr.StatusCode != http.StatusForbidden {
	}
	return false
}
