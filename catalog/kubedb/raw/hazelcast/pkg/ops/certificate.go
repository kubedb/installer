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

package hazelcast

import (
	"context"
	"fmt"

	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"kubedb.dev/apimachinery/pkg/eventer"
	"kubedb.dev/apimachinery/pkg/lib"
	apiutils "kubedb.dev/apimachinery/pkg/utils"

	cm_api "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	v12 "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	cm_util "kmodules.xyz/cert-manager-util/certmanager/v1"
	kutil "kmodules.xyz/client-go"
	kmapi "kmodules.xyz/client-go/api/v1"
	core_util "kmodules.xyz/client-go/core/v1"
)

func (c *Controller) manageServerCert(hz *dbapi.Hazelcast) error {
	certVerb, err := c.ensureServerCert(hz, dbapi.HazelcastServerCert)
	if err != nil {
		return err
	}

	switch certVerb {
	case kutil.VerbCreated:
		c.Recorder.Event(
			hz,
			core.EventTypeNormal,
			eventer.EventReasonSuccessful,
			"Successfully created Hazelcast server certificates",
		)
	case kutil.VerbPatched:
		c.Recorder.Event(
			hz,
			core.EventTypeNormal,
			eventer.EventReasonSuccessful,
			"Successfully patched Hazelcast server certificates",
		)
	}
	if certVerb != kutil.VerbUnchanged {
		c.log.Info(fmt.Sprintf("server-certificates %s", certVerb))
	}

	// wait for certificate secret to be created
	ref := metav1.ObjectMeta{
		Name:      hz.GetCertSecretName(dbapi.HazelcastServerCert),
		Namespace: hz.Namespace,
	}
	if lib.SecretExists(c.Client, ref) {
		// set hz as the owner of the certificate-secret
		if err := lib.AddOwnerReferenceToSecret(c.Client, dbapi.SchemeGroupVersion.WithKind(dbapi.ResourceKindHazelcast), hz, ref); err != nil {
			c.log.Error(err, "failed to set owner reference to server certificate secret")
			return err
		}
	}

	return err
}

func (c *Controller) ensureServerCert(db *dbapi.Hazelcast, alias dbapi.HazelcastCertificateAlias) (kutil.VerbType, error) {
	var duration *metav1.Duration
	var subject *cm_api.X509Subject
	renewBefore := metav1.Duration{Duration: lib.DefaultCertRenewBefore}
	var uriSANs, emailSANs []string

	dnsNames, ipAddresses, err := c.upsertServiceHosts(db)
	if err != nil {
		return kutil.VerbUnchanged, err
	}

	if _, cert := kmapi.GetCertificate(db.Spec.TLS.Certificates, string(dbapi.HazelcastServerCert)); cert != nil {
		dnsNames.Insert(cert.DNSNames...)
		ipAddresses.Insert(cert.IPAddresses...)
		duration = cert.Duration
		if cert.Subject != nil {
			subject = &cm_api.X509Subject{
				Organizations:       cert.Subject.Organizations,
				Countries:           cert.Subject.Countries,
				OrganizationalUnits: cert.Subject.OrganizationalUnits,
				Localities:          cert.Subject.Localities,
				Provinces:           cert.Subject.Provinces,
				StreetAddresses:     cert.Subject.StreetAddresses,
				PostalCodes:         cert.Subject.PostalCodes,
				SerialNumber:        cert.Subject.SerialNumber,
			}
		}
		uriSANs = cert.URIs
		emailSANs = cert.EmailAddresses
	}

	setDnsNames(&dnsNames, db)
	ipAddresses.Insert(kubedb.LocalHostIP)
	ref := metav1.NewControllerRef(db, dbapi.SchemeGroupVersion.WithKind(dbapi.ResourceKindHazelcast))

	_, vt, err := cm_util.CreateOrPatchCertificate(context.TODO(),
		c.CertManagerClient.CertmanagerV1(),
		metav1.ObjectMeta{
			Name:      db.CertificateName(dbapi.HazelcastServerCert),
			Namespace: db.GetNamespace(),
		},
		func(in *cm_api.Certificate) *cm_api.Certificate {
			in.Labels = db.OffshootLabels()
			core_util.EnsureOwnerReference(in, ref)

			in.Spec.Subject = subject
			in.Spec.CommonName = db.ServiceName()
			in.Spec.Duration = duration // Default
			in.Spec.RenewBefore = &renewBefore
			in.Spec.DNSNames = sets.List[string](dnsNames)       // including Service URL and localhost
			in.Spec.IPAddresses = sets.List[string](ipAddresses) // including 127.0.0.1
			in.Spec.URIs = sets.NewString(uriSANs...).List()
			in.Spec.EmailAddresses = sets.NewString(emailSANs...).List()
			in.Spec.SecretName = db.GetCertSecretName(alias)
			in.Spec.IssuerRef = lib.GetIssuerObjectRef(db.Spec.TLS, string(alias))
			in.Spec.Usages = []cm_api.KeyUsage{
				cm_api.UsageDigitalSignature,
				cm_api.UsageKeyEncipherment,
				cm_api.UsageServerAuth,
				cm_api.UsageClientAuth,
			}
			in.Spec.Keystores = &cm_api.CertificateKeystores{
				PKCS12: &cm_api.PKCS12Keystore{
					Create: true,
					PasswordSecretRef: v12.SecretKeySelector{
						LocalObjectReference: v12.LocalObjectReference{
							Name: db.Spec.KeystoreSecret.Name,
						},
						Key: kubedb.HazelcastKeystorePassKey,
					},
				},
			}

			return in
		}, metav1.PatchOptions{},
	)
	if err != nil {
		return kutil.VerbUnchanged, err
	}

	return vt, nil
}

func (c *Controller) manageClientCert(hz *dbapi.Hazelcast) error {
	certVerb, err := c.ensureClientCert(hz, dbapi.HazelcastClientCert)
	if err != nil {
		return err
	}

	switch certVerb {
	case kutil.VerbCreated:
		c.Recorder.Event(
			hz,
			core.EventTypeNormal,
			eventer.EventReasonSuccessful,
			"Successfully created Hazelcast client-certificates",
		)
	case kutil.VerbPatched:
		c.Recorder.Event(
			hz,
			core.EventTypeNormal,
			eventer.EventReasonSuccessful,
			"Successfully patched Hazelcast client-certificates",
		)
	}
	if certVerb != kutil.VerbUnchanged {
		c.log.Info(fmt.Sprintf("client-certificates %s", certVerb))
	}

	// wait for certificate secret to be created
	ref := metav1.ObjectMeta{
		Name:      hz.GetCertSecretName(dbapi.HazelcastClientCert),
		Namespace: hz.Namespace,
	}
	if lib.SecretExists(c.Client, ref) {
		// set hz as the owner of the certificate-ref
		if err := lib.AddOwnerReferenceToSecret(c.Client, dbapi.SchemeGroupVersion.WithKind(dbapi.ResourceKindHazelcast), hz, ref); err != nil {
			c.log.Error(err, "failed to set owner reference to client certificate secret")
			return err
		}
	}
	return nil
}

func (c *Controller) ensureClientCert(db *dbapi.Hazelcast, alias dbapi.HazelcastCertificateAlias) (kutil.VerbType, error) {
	var duration *metav1.Duration
	var subject *cm_api.X509Subject
	renewBefore := metav1.Duration{Duration: lib.DefaultCertRenewBefore}
	var uriSANs, emailSANs []string
	dnsNames, ipAddresses, err := c.upsertServiceHosts(db)
	if err != nil {
		return kutil.VerbUnchanged, err
	}
	if _, cert := kmapi.GetCertificate(db.Spec.TLS.Certificates, string(alias)); cert != nil {
		dnsNames.Insert(cert.DNSNames...)
		ipAddresses.Insert(cert.IPAddresses...)
		duration = cert.Duration
		if cert.Subject != nil {
			subject = &cm_api.X509Subject{
				Organizations:       cert.Subject.Organizations,
				Countries:           cert.Subject.Countries,
				OrganizationalUnits: cert.Subject.OrganizationalUnits,
				Localities:          cert.Subject.Localities,
				Provinces:           cert.Subject.Provinces,
				StreetAddresses:     cert.Subject.StreetAddresses,
				PostalCodes:         cert.Subject.PostalCodes,
				SerialNumber:        cert.Subject.SerialNumber,
			}
		}
		uriSANs = cert.URIs
		emailSANs = cert.EmailAddresses
	}

	setDnsNames(&dnsNames, db)
	ipAddresses.Insert(kubedb.LocalHostIP)
	ref := metav1.NewControllerRef(db, dbapi.SchemeGroupVersion.WithKind(dbapi.ResourceKindHazelcast))

	_, vt, err := cm_util.CreateOrPatchCertificate(context.TODO(),
		c.CertManagerClient.CertmanagerV1(),
		metav1.ObjectMeta{
			Name:      db.CertificateName(alias),
			Namespace: db.GetNamespace(),
		},
		func(in *cm_api.Certificate) *cm_api.Certificate {
			in.Labels = db.OffshootLabels()
			core_util.EnsureOwnerReference(in, ref)

			in.Spec.CommonName = db.ServiceName()
			in.Spec.Subject = subject
			in.Spec.Duration = duration
			in.Spec.RenewBefore = &renewBefore
			in.Spec.DNSNames = sets.List[string](dnsNames)
			in.Spec.IPAddresses = sets.List[string](ipAddresses)
			in.Spec.URIs = sets.NewString(uriSANs...).List()
			in.Spec.EmailAddresses = sets.NewString(emailSANs...).List()
			in.Spec.SecretName = db.GetCertSecretName(alias) // Secret where issued certificates will be saved
			in.Spec.IssuerRef = lib.GetIssuerObjectRef(db.Spec.TLS, string(alias))
			in.Spec.Usages = []cm_api.KeyUsage{
				cm_api.UsageDigitalSignature,
				cm_api.UsageKeyEncipherment,
				cm_api.UsageClientAuth,
			}

			in.Spec.Keystores = &cm_api.CertificateKeystores{
				PKCS12: &cm_api.PKCS12Keystore{
					Create: true,
					PasswordSecretRef: v12.SecretKeySelector{
						LocalObjectReference: v12.LocalObjectReference{
							Name: db.Spec.KeystoreSecret.Name,
						},
						Key: kubedb.HazelcastKeystorePassKey,
					},
				},
			}

			return in
		}, metav1.PatchOptions{})
	if err != nil {
		return kutil.VerbUnchanged, err
	}

	return vt, nil
}

func (c *Controller) upsertServiceHosts(db *dbapi.Hazelcast) (sets.Set[string], sets.Set[string], error) {
	ctx := context.Background()
	var svc core.Service
	err := c.KBClient.Get(ctx, types.NamespacedName{
		Namespace: db.Namespace,
		Name:      db.GoverningServiceName(),
	}, &svc)
	if err != nil {
		return nil, nil, err
	}
	dnsNames, ipAddresses, err := lib.ServiceHosts(c.Client.CoreV1(), svc.ObjectMeta)
	if err != nil {
		return nil, nil, err
	}

	return dnsNames, ipAddresses, nil
}

func setDnsNames(dnsNames *sets.Set[string], db *dbapi.Hazelcast) {
	dnsNames.Insert(fmt.Sprintf("%v.%v.svc.%v", db.GoverningServiceName(), db.Namespace, apiutils.FindDomain()))
	dnsNames.Insert(fmt.Sprintf("*.%v.%v.svc.%v", db.GoverningServiceName(), db.Namespace, apiutils.FindDomain()))
	dnsNames.Insert(fmt.Sprintf("*.%v.%v", db.GoverningServiceName(), db.Namespace))
	dnsNames.Insert(fmt.Sprintf("*.%v.%v.svc", db.GoverningServiceName(), db.Namespace))
	dnsNames.Insert(fmt.Sprintf("*.%v.%v.svc.%v", db.GoverningServiceName(), db.Namespace, apiutils.FindDomain()))
	dnsNames.Insert(ServiceDNS(metav1.ObjectMeta{
		Name:      db.ServiceName(),
		Namespace: db.GetNamespace(),
	})...)
	dnsNames.Insert(kubedb.LocalHost)
}

func ServiceDNS(svc metav1.ObjectMeta) []string {
	return []string{
		svc.Name + "." + svc.Namespace + ".svc",
		svc.Name + "." + svc.Namespace + ".svc." + apiutils.FindDomain(),
		svc.Name,
	}
}
