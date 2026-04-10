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

package server

import (
	"context"

	webhooks "kubedb.dev/apimachinery/pkg/webhooks/kubedb/v1alpha2"
	"kubedb.dev/hazelcast/pkg/controller"
	"kubedb.dev/hazelcast/pkg/server"

	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

type KubeDBWebhookOptions struct {
	WebhookOptions *WebhookOptions
}

func NewKubeDBWebhookOptions() *KubeDBWebhookOptions {
	o := &KubeDBWebhookOptions{
		WebhookOptions: NewWebhookOptions(),
	}

	return o
}

func (o *KubeDBWebhookOptions) AddFlags(fs *pflag.FlagSet) {
	o.WebhookOptions.AddFlags(fs)
}

func (o KubeDBWebhookOptions) Validate() error {
	errs := o.WebhookOptions.Validate()
	return utilerrors.NewAggregate(errs)
}

func (o *KubeDBWebhookOptions) Complete() error {
	return nil
}

func (o KubeDBWebhookOptions) Config() (*server.KubeDBWebhookConfig, error) {
	webhookConfig := &controller.WebhookConfig{
		ClientConfig: ctrl.GetConfigOrDie(),
	}
	if err := o.WebhookOptions.ApplyTo(webhookConfig); err != nil {
		return nil, err
	}

	config := &server.KubeDBWebhookConfig{
		WebhookConfig: webhookConfig,
	}
	return config, nil
}

func (o KubeDBWebhookOptions) Run(ctx context.Context) error {
	cfg, err := o.Config()
	if err != nil {
		return err
	}

	s, err := cfg.Complete().New()
	if err != nil {
		return err
	}

	if err = webhooks.SetupHazelcastWebhookWithManager(s.Manager); err != nil {
		return err
	}
	return s.Run(ctx)
}
