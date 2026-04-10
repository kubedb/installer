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

package yqstrings

import (
	"os"

	yqslib "kubedb.dev/apimachinery/pkg/yq3/lib"

	"github.com/mikefarah/yq/v3/pkg/yqlib"
	"gopkg.in/op/go-logging.v1"
	"gopkg.in/yaml.v3"
)

func StringToNode(yml string) (*yaml.Node, error) {
	var node yaml.Node
	err := yaml.Unmarshal([]byte(yml), &node)
	if err != nil {
		return nil, err
	}

	return &node, nil
}

func NodeToString(node *yaml.Node) (string, error) {
	b, err := yaml.Marshal(node)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func merge(arrayMergeStrategy yqlib.ArrayMergeStrategy, commentsMergeStrategy yqlib.CommentsMergeStrategy, overwrite bool, A, B string) (string, error) {
	// Hide debug message
	initializeLogging()

	nodeA, err := StringToNode(A)
	if err != nil {
		return "", err
	}

	nodeB, err := StringToNode(B)
	if err != nil {
		return "", err
	}
	lib := yqslib.NewYqStringLib()
	nodeCtxs, err := lib.NodeToNodeContexts(nodeB, arrayMergeStrategy)
	if err != nil {
		return "", err
	}
	for _, nctx := range nodeCtxs {
		err := lib.YqLib.Update(nodeA, yqlib.UpdateCommand{
			Command:               "merge",
			Path:                  lib.YqLib.MergePathStackToString(nctx.PathStack, arrayMergeStrategy),
			Value:                 nctx.Node,
			Overwrite:             overwrite,
			CommentsMergeStrategy: commentsMergeStrategy,
			// dont update the content for nodes midway, only leaf nodes
			DontUpdateNodeContent: nctx.IsMiddleNode && (arrayMergeStrategy != yqlib.OverwriteArrayMergeStrategy || nctx.Node.Kind != yaml.SequenceNode),
		}, true)
		if err != nil {
			return "", err
		}
	}

	output, err := NodeToString(nodeA)
	if err != nil {
		return "", err
	}

	return output, nil
}

func Merge(arrayMergeStrategy yqlib.ArrayMergeStrategy, commentsMergeStrategy yqlib.CommentsMergeStrategy, overwrite bool, inputs ...string) (string, error) {
	var output string
	var err error
	for _, s := range inputs {
		output, err = merge(arrayMergeStrategy, commentsMergeStrategy, overwrite, output, s)
		if err != nil {
			return "", err
		}
	}

	return output, nil
}

func initializeLogging() {
	backend := logging.AddModuleLevel(logging.NewLogBackend(os.Stderr, "", 0))
	backend.SetLevel(logging.ERROR, "")
	logging.SetBackend(backend)
}
