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

package controller

import (
	"fmt"

	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	kutil "kmodules.xyz/client-go"
	clientutil "kmodules.xyz/client-go/client"
	coreutil "kmodules.xyz/client-go/core/v1"
	meta_util "kmodules.xyz/client-go/meta"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
	ofst "kmodules.xyz/offshoot-api/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type serviceOptions struct {
	isHeadless  bool
	svcName     string
	svcTemplate ofst.ServiceTemplateSpec
	labels      map[string]string
	selectors   map[string]string
	svcPort     []core.ServicePort
}

func (rs *ReconcileState) EnsureServices() error {
	svcFunc := func(opts serviceOptions) error {
		var err error
		var svc core.Service

		svcTemplate := opts.svcTemplate

		if err = rs.KBClient.Get(rs.ctx, types.NamespacedName{
			Name:      opts.svcName,
			Namespace: rs.db.Namespace,
		}, &svc); err != nil {
			if !kerr.IsNotFound(err) {
				rs.log.Error(err, fmt.Sprintf("Failed to get Service %s/%s", rs.db.Namespace, opts.svcName))
				return err
			}
		}

		v, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, &core.Service{
			ObjectMeta: meta.ObjectMeta{
				Name:      opts.svcName,
				Namespace: rs.db.Namespace,
			},
		}, func(obj client.Object, createOp bool) client.Object {
			in := obj.(*core.Service)
			in.Labels = opts.labels
			in.Spec.Type = core.ServiceTypeClusterIP
			in.Spec.PublishNotReadyAddresses = true
			if createOp {
				coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
			}
			in.Spec.Selector = opts.selectors
			in.Spec.Ports = coreutil.MergeServicePorts(in.Spec.Ports, opts.svcPort)
			if !opts.isHeadless {
				in.Annotations = svcTemplate.Annotations
				in.Spec.Ports = ofst.PatchServicePorts(
					in.Spec.Ports,
					svcTemplate.Spec.Ports,
				)
				if svcTemplate.Spec.ClusterIP != "" {
					in.Spec.ClusterIP = svcTemplate.Spec.ClusterIP
				}
				if svcTemplate.Spec.Type != "" {
					in.Spec.Type = svcTemplate.Spec.Type
				}
				in.Spec.ExternalIPs = svcTemplate.Spec.ExternalIPs
				in.Spec.LoadBalancerIP = svcTemplate.Spec.LoadBalancerIP
				in.Spec.LoadBalancerSourceRanges = svcTemplate.Spec.LoadBalancerSourceRanges
				in.Spec.ExternalTrafficPolicy = svcTemplate.Spec.ExternalTrafficPolicy
				if svcTemplate.Spec.HealthCheckNodePort > 0 {
					in.Spec.HealthCheckNodePort = svcTemplate.Spec.HealthCheckNodePort
				}
			} else {
				// create headless service
				in.Spec.ClusterIP = core.ClusterIPNone
			}

			return in
		})

		if v == kutil.VerbCreated {
			rs.log.Info(fmt.Sprintf("Service: %s/%s created", rs.db.Namespace, opts.svcName))
		}
		if err != nil {
			return err
		}
		return nil
	}

	if err := rs.ensureGoverningService(svcFunc); err != nil {
		rs.log.Error(err, "Failed to ensure governing service")
		return err
	}

	if err := rs.ensurePrimaryService(svcFunc); err != nil {
		rs.log.Error(err, "Failed to ensure governing service")
		return err
	}

	if err := rs.ensureStatsService(); err != nil {
		rs.log.Error(err, "Failed to ensure stats service")
		return err
	}

	return nil
}

func (rs *ReconcileState) ensureGoverningService(svcFunc func(opts serviceOptions) error) error {
	svcOptsGoverning := serviceOptions{
		isHeadless: true,
		svcName:    rs.db.GoverningServiceName(),
		labels:     rs.db.OffshootLabels(),
		selectors:  rs.db.OffshootSelectors(),
		svcPort: []core.ServicePort{
			{
				Name:       kubedb.HazelcastPortName,
				Protocol:   core.ProtocolTCP,
				TargetPort: intstr.FromString(kubedb.HazelcastPortName),
				Port:       kubedb.HazelcastRestPort,
			},
		},
	}

	if err := svcFunc(svcOptsGoverning); err != nil {
		rs.log.Error(err, fmt.Sprintf("Failed to ensure Service: %s/%s", rs.db.Namespace, rs.db.GoverningServiceName()))
		return err
	}

	return nil
}

func (rs *ReconcileState) ensurePrimaryService(svcFunc func(opts serviceOptions) error) error {
	svcSelectors := rs.db.OffshootSelectors()

	svcOptsPrimary := serviceOptions{
		svcName:     rs.db.ServiceName(),
		svcTemplate: api.GetServiceTemplate(rs.db.Spec.ServiceTemplates, api.PrimaryServiceAlias),
		labels:      rs.db.OffshootLabels(),
		selectors:   svcSelectors,
		svcPort: []core.ServicePort{
			{
				Name:       kubedb.HazelcastPortName,
				Protocol:   core.ProtocolTCP,
				TargetPort: intstr.FromString(kubedb.HazelcastPortName),
				Port:       kubedb.HazelcastRestPort,
			},
			{
				Name:       "ui-port",
				Protocol:   core.ProtocolTCP,
				TargetPort: intstr.FromString("ui-port"),
				Port:       kubedb.HazelcastUIPort,
			},
		},
	}

	if err := svcFunc(svcOptsPrimary); err != nil {
		rs.log.Error(err, fmt.Sprintf("Failed to ensure Service: %s/%s", rs.db.Namespace, rs.db.ServiceName()))
		return err
	}

	return nil
}

func (rs *ReconcileState) ensureStatsService() error {
	// return if monitoring is not prometheus
	if rs.db.Spec.Monitor == nil || rs.db.Spec.Monitor.Agent.Vendor() != mona.VendorPrometheus {
		return nil
	}

	var svc core.Service
	svcName := rs.db.StatsService().ServiceName()
	svcTemplate := api.GetServiceTemplate(rs.db.Spec.ServiceTemplates, api.StatsServiceAlias)
	if err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
		Name:      svcName,
		Namespace: rs.db.Namespace,
	}, &svc); err != nil {
		if kerr.IsNotFound(err) {
			rs.log.Error(err, fmt.Sprintf("Service %s/%s not found", rs.db.Namespace, svcName))
		} else {
			rs.log.Error(err, fmt.Sprintf("Failed to get Service %s/%s", rs.db.Namespace, svcName))
			return err
		}
	}

	v, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, &core.Service{
		ObjectMeta: meta.ObjectMeta{
			Name:      svcName,
			Namespace: rs.db.Namespace,
		},
	}, func(obj client.Object, createOp bool) client.Object {
		in := obj.(*core.Service)
		in.Labels = rs.db.StatsServiceLabels()
		in.Annotations = meta_util.OverwriteKeys(in.Annotations, svcTemplate.Annotations)
		in.Spec.Selector = rs.db.OffshootSelectors()
		in.Spec.Ports = ofst.PatchServicePorts(
			coreutil.MergeServicePorts(in.Spec.Ports, []core.ServicePort{
				{
					Name:       mona.PrometheusExporterPortName,
					Protocol:   core.ProtocolTCP,
					TargetPort: intstr.FromString(mona.PrometheusExporterPortName),
					Port:       rs.db.Spec.Monitor.Prometheus.Exporter.Port,
				},
			}),
			svcTemplate.Spec.Ports,
		)
		if svcTemplate.Spec.ClusterIP != "" {
			in.Spec.ClusterIP = svcTemplate.Spec.ClusterIP
		}
		if svcTemplate.Spec.Type != "" {
			in.Spec.Type = svcTemplate.Spec.Type
		}
		in.Spec.ExternalIPs = svcTemplate.Spec.ExternalIPs
		in.Spec.LoadBalancerIP = svcTemplate.Spec.LoadBalancerIP
		in.Spec.LoadBalancerSourceRanges = svcTemplate.Spec.LoadBalancerSourceRanges
		in.Spec.ExternalTrafficPolicy = svcTemplate.Spec.ExternalTrafficPolicy
		if svcTemplate.Spec.HealthCheckNodePort > 0 {
			in.Spec.HealthCheckNodePort = svcTemplate.Spec.HealthCheckNodePort
		}
		if createOp {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
		}
		return in
	})

	if v == kutil.VerbCreated {
		rs.log.Info(fmt.Sprintf("Service: %s/%s created", rs.db.Namespace, svcName))
	}
	if err != nil {
		return err
	}
	return nil
}
