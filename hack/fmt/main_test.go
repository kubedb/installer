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

import "testing"

func TestParseImage(t *testing.T) {
	tests := []struct {
		name     string
		arg      string
		wantReg  string
		wantRepo string
		wantBin  string
		wantTag  string
	}{
		{
			name:     "kubedb/postgres:v1.2.3",
			arg:      "kubedb/postgres:v1.2.3",
			wantReg:  "",
			wantRepo: "kubedb",
			wantBin:  "postgres",
			wantTag:  "v1.2.3",
		},
		{
			name:     "postgres:v1.2.3",
			arg:      "postgres:v1.2.3",
			wantReg:  "",
			wantRepo: "",
			wantBin:  "postgres",
			wantTag:  "v1.2.3",
		},
		{
			name:     "ghcr.io/myorg/postgres:v1.2.3",
			arg:      "ghcr.io/myorg/postgres:v1.2.3",
			wantReg:  "ghcr.io",
			wantRepo: "myorg",
			wantBin:  "postgres",
			wantTag:  "v1.2.3",
		},
		{
			name:     "ghcr.io/sub/myorg/postgres:v1.2.3",
			arg:      "ghcr.io/sub/myorg/postgres:v1.2.3",
			wantReg:  "ghcr.io/sub",
			wantRepo: "myorg",
			wantBin:  "postgres",
			wantTag:  "v1.2.3",
		},
		{
			name:     "ghcr.io/sub/myorg/postgres",
			arg:      "ghcr.io/sub/myorg/postgres",
			wantReg:  "ghcr.io/sub",
			wantRepo: "myorg",
			wantBin:  "postgres",
			wantTag:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReg, gotRepo, gotBin, gotTag := ParseImage(tt.arg)
			if gotReg != tt.wantReg {
				t.Errorf("ParseImage() gotReg = %v, want %v", gotReg, tt.wantReg)
			}
			if gotRepo != tt.wantRepo {
				t.Errorf("ParseImage() gotRepo = %v, want %v", gotRepo, tt.wantRepo)
			}
			if gotBin != tt.wantBin {
				t.Errorf("ParseImage() gotBin = %v, want %v", gotBin, tt.wantBin)
			}
			if gotTag != tt.wantTag {
				t.Errorf("ParseImage() gotTag = %v, want %v", gotTag, tt.wantTag)
			}
		})
	}
}
