/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"flag"

	"kubedb.dev/hazelcast/pkg/controller"

	"github.com/spf13/pflag"
)

type WebhookOptions struct {
	QPS   float64
	Burst int

	MetricsAddr          string
	CertDir              string
	EnableLeaderElection bool
	ProbeAddr            string
	SecureMetrics        bool
	EnableHTTP2          bool
}

func NewWebhookOptions() *WebhookOptions {
	return &WebhookOptions{
		// ref: https://github.com/kubernetes/ingress-nginx/blob/e4d53786e771cc6bdd55f180674b79f5b692e552/pkg/ingress/controller/launch.go#L252-L259
		// High enough QPS to fit all expected use cases. QPS=0 is not set here, because client code is overriding it.
		QPS: 1e6,
		// High enough Burst to fit all expected use cases. Burst=0 is not set here, because client code is overriding it.
		Burst:                1e6,
		MetricsAddr:          "0",
		ProbeAddr:            ":8081",
		EnableLeaderElection: false,
		SecureMetrics:        true,
		CertDir:              "",
		EnableHTTP2:          false,
	}
}

func (s *WebhookOptions) AddGoFlags(fs *flag.FlagSet) {
	fs.Float64Var(&s.QPS, "qps", s.QPS, "The maximum QPS to the master from this client")
	fs.IntVar(&s.Burst, "burst", s.Burst, "The maximum burst for throttle")

	fs.StringVar(&s.MetricsAddr, "metrics-bind-address", s.MetricsAddr, "The address the metrics endpoint binds to. "+
		"Use :8443 for HTTPS or :8080 for HTTP, or leave as 0 to disable the metrics service.")
	fs.StringVar(&s.ProbeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	fs.BoolVar(&s.EnableLeaderElection, "leader-elect", s.EnableLeaderElection,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	fs.BoolVar(&s.SecureMetrics, "metrics-secure", s.SecureMetrics,
		"If set, the metrics endpoint is served securely via HTTPS. Use --metrics-secure=false to use HTTP instead.")
	fs.StringVar(&s.CertDir, "cert-dir", s.CertDir, "The directory that contains the webhook and metrics server certificate.")
	fs.BoolVar(&s.EnableHTTP2, "enable-http2", s.EnableHTTP2,
		"If set, HTTP/2 will be enabled for the metrics and webhook servers")
}

func (s *WebhookOptions) AddFlags(fs *pflag.FlagSet) {
	pfs := flag.NewFlagSet("extra-flags", flag.ExitOnError)
	s.AddGoFlags(pfs)
	fs.AddGoFlagSet(pfs)
}

func (s *WebhookOptions) ApplyTo(cfg *controller.WebhookConfig) error {
	cfg.ClientConfig.QPS = float32(s.QPS)
	cfg.ClientConfig.Burst = s.Burst

	cfg.MetricsAddr = s.MetricsAddr
	cfg.ProbeAddr = s.ProbeAddr
	cfg.EnableLeaderElection = s.EnableLeaderElection
	cfg.SecureMetrics = s.SecureMetrics
	cfg.CertDir = s.CertDir
	cfg.EnableHTTP2 = s.EnableHTTP2

	return nil
}

func (s *WebhookOptions) Validate() []error {
	return nil
}
