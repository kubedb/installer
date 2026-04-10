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

package cmds

import (
	"context"

	"kubedb.dev/hazelcast/pkg/cmds/server"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"kmodules.xyz/client-go/meta"
)

func NewCmdWebhook(ctx context.Context) *cobra.Command {
	o := server.NewKubeDBWebhookOptions()

	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch KubeDB Webhook Server",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			klog.Infoln("Starting kubedb-webhook-server...")

			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(ctx); err != nil {
				return err
			}
			return nil
		},
	}

	o.AddFlags(cmd.Flags())
	meta.AddLabelBlacklistFlag(cmd.Flags())

	return cmd
}
