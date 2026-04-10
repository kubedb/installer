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
	"time"

	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"
	"kubedb.dev/apimachinery/pkg/lib"

	cm_api "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	cm_util "kmodules.xyz/cert-manager-util/certmanager/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	clientutil "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	core_util "kmodules.xyz/client-go/core/v1"
	policy_util "kmodules.xyz/client-go/policy"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type syncCertificatesRetries struct {
	getCertificateRetries   *Retries
	readyConditionRetries   *Retries
	issuingConditionRetries *Retries
}

type syncCertificates struct {
	*hzOpsReqController
	db      *dbapi.Hazelcast
	retries syncCertificatesRetries
}

func (c *hzOpsReqController) newSyncCertificates(db *dbapi.Hazelcast) func() (bool, error) {
	opts := syncCertificates{
		hzOpsReqController: c,
		db:                 db,
		retries:            syncCertificatesRetries{},
	}

	opts.retries.getCertificateRetries = c.newRetries("GetCertificateRetries")
	opts.retries.readyConditionRetries = c.newRetries("CheckReadyCondition")
	opts.retries.issuingConditionRetries = c.newRetries("IssuingCondition")

	return opts.sync
}

func (c *syncCertificates) sync() (bool, error) {
	// Add nil checks to prevent panic
	if c.db == nil {
		return false, errors.New("database object is nil")
	}

	if c.db.Spec.TLS == nil || c.db.Spec.TLS.Certificates == nil {
		return false, nil // No certificates to sync
	}

	for _, cert := range c.db.Spec.TLS.Certificates {
		certName := c.db.CertificateName(dbapi.HazelcastCertificateAlias(cert.Alias))
		cmCert, err := c.CertManagerClient.CertmanagerV1().Certificates(c.db.Namespace).Get(context.TODO(), certName, metav1.GetOptions{})
		if err != nil && kerr.IsNotFound(err) {
			return false, err
		} else if err != nil {
			return c.retries.getCertificateRetries.Wait(), err
		}
		c.retries.getCertificateRetries.Initialize()

		// A Certificate CR is synced when:
		//	- C.status.condition[type="Ready"].status == true
		//	- C.generation == C.status.condition[type="Ready"].observeGeneration
		// 	- C.status.conditions does not contain type="Issuing" (i.e. rotate)
		idx, cond := lib.GetCertificateCondition(cmCert.Status.Conditions, cm_api.CertificateConditionReady)
		if idx == -1 {
			return c.retries.getCertificateRetries.Wait(), errors.Errorf("Certificate: %s does not have Ready condition", certName)
		}

		if cmCert.Generation != cond.ObservedGeneration || cond.Status == cmmeta.ConditionFalse {
			return c.retries.readyConditionRetries.Wait(), errors.Errorf("Certificate: %s is not ready yet", certName)
		}
		c.retries.readyConditionRetries.Initialize()

		if lib.HasCertificateCondition(cmCert.Status.Conditions, cm_api.CertificateConditionIssuing) {
			return c.retries.issuingConditionRetries.Wait(), errors.Errorf("Certificate: %s is not issued yet", certName)
		}
		c.retries.issuingConditionRetries.Initialize()
	}
	return false, nil
}

// ReconfigureTLS Algorithm:
// - deepCopy DB object to "dbCopy"
// if req...Remove == true
//   - dbCopy.Spec.KeystoreCredSecret
//   - remove all certificate and keystoreCred if available
//
// else
//
//	# here, update, add, rotate certificates
//	- SetCertificate() from req object to dbCopy
//	- Set issuer to dbCopy if req object has any
//
//		- Call SetDefaultTLS(dbCopy)
//	else
//		# update & rotate
//		- Call SetDefaultTLS(dbCopy)
//	endif
//
//		- Call manageTLS(dbCopy) to update the Certificate CRs
//		- If req.rotateTLS==true
//			- Set Certificate condition["Issuing"]=true
//		- Wait for the Certificate object to sync
//			- check Certificate.generation == Certificate.status.Condition[type="Ready"--> true].observeGeneration
//			- also check, Condition[type="Issuing"] is removed.
//
// endif
// - Perform "ensureKeystoreCred(dbCopy) for keystore credentials secret
// - Perform "reconcile(dbCopy)" to update the PetSets
// - Restart the pods, to reflect the changes
// - Update original DB object with changes from dbCopy

func (c *hzOpsReqController) reconfigureTLS() (time.Duration, error) {
	rq, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeReconfigureTLS),
		"Hazelcast ops-request has started to reconfigure tls for Hazelcast nodes")

	if (err != nil) || (rq == RequeueDuration) {
		return rq, err
	}

	tlsConfig := c.req.Spec.TLS
	dbCopy := c.db.DeepCopy()

	if tlsConfig.Remove {
		c.log.Info("TLS removal: Starting cleanup process")

		if err := c.cleanUpCertificates(dbCopy); err != nil {
			return DefaultDuration, err
		}

		// Remove TLS and keystore configurations from the spec
		dbCopy.Spec.TLS = nil
		dbCopy.Spec.KeystoreSecret = nil
		dbCopy.Spec.EnableSSL = false

		// Update health probe schemes from HTTPS to HTTP
		c.updateProbeSchemes(dbCopy, core.URISchemeHTTP)
	} else {
		// For TLS addition/update
		reconcile, err := c.NewHazelcastReconcile(dbCopy)
		if err != nil {
			c.log.Error(err, "Failed to get Hazelcast Reconcile")
			return DefaultDuration, err
		}

		if tlsConfig.Certificates != nil {
			if dbCopy.Spec.TLS == nil {
				dbCopy.Spec.TLS = &kmapi.TLSConfig{
					Certificates: tlsConfig.Certificates,
				}
			} else {
				dbCopy.Spec.TLS.Certificates = tlsConfig.Certificates
			}
		}

		if tlsConfig.IssuerRef != nil {
			if dbCopy.Spec.TLS == nil {
				dbCopy.Spec.TLS = &kmapi.TLSConfig{
					IssuerRef: tlsConfig.IssuerRef,
				}
			} else {
				dbCopy.Spec.TLS.IssuerRef = tlsConfig.IssuerRef
			}
		}

		dbCopy.SetTLSDefaults()
		dbCopy.Spec.EnableSSL = true
		// Update health probe schemes from HTTP to HTTPS
		c.updateProbeSchemes(dbCopy, core.URISchemeHTTPS)

		// Ensure keystore secret only when adding/updating TLS
		_, err = reconcile.EnsureKeystoreSecret()
		if err != nil {
			c.log.Error(err, "Failed to ensure keystore secret")
			return DefaultDuration, err
		}

		// Set keystoreCredSecret name if not set(Rotate and Add TLS)
		if dbCopy.Spec.KeystoreSecret == nil {
			dbCopy.Spec.KeystoreSecret = &core.LocalObjectReference{
				Name: dbCopy.HazelcastSecretName("keystore-cred"),
			}
		} else {
			dbCopy.Spec.KeystoreSecret.Name = dbCopy.HazelcastSecretName("keystore-cred")
		}
		err = c.manageTLS(dbCopy)
		if err != nil {
			c.log.Error(err, "Failed to manage TLS for Hazelcast")
			return DefaultDuration, err
		}
	}

	// If req...RotateCertificate==true, add "Issuing" condition to all certificates
	if tlsConfig.RotateCertificates && !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.IssueCertificatesSucceeded) {
		// Add nil check to prevent panic
		if dbCopy.Spec.TLS != nil && dbCopy.Spec.TLS.Certificates != nil {
			for _, cert := range dbCopy.Spec.TLS.Certificates {
				// Add nil check for c.db before calling methods on it
				if c.db == nil {
					c.log.Error(nil, "database object is nil when getting certificate name")
					return DefaultDuration, errors.New("database object is nil when getting certificate name")
				}

				certName := c.db.CertificateName(dbapi.HazelcastCertificateAlias(cert.Alias))

				_, err := cm_util.UpdateCertificateStatus(
					context.TODO(), c.CertManagerClient.CertmanagerV1(), metav1.ObjectMeta{
						Name:      certName,
						Namespace: c.db.Namespace,
					}, func(status *cm_api.CertificateStatus) *cm_api.CertificateStatus {
						status.Conditions = lib.UpsertCertificateCondition(status.Conditions, cm_api.CertificateCondition{
							Type:    cm_api.CertificateConditionIssuing,
							Status:  cmmeta.ConditionTrue,
							Reason:  "RotateCertificate",
							Message: "Rotating Certificate for KubeDB",
						})
						return status
					}, metav1.UpdateOptions{})
				if err != nil {
					c.log.Error(err, "Failed to update certificate status")
					return DefaultDuration, err
				}
			}
		}
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.CertificateSynced) && !tlsConfig.Remove {
		c.RunParallel(opsapi.CertificateSynced, "Successfully synced TLS certificates", c.newSyncCertificates(dbCopy))
		return DefaultDuration, nil
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
		rcl, err := c.NewHazelcastReconcile(dbCopy)
		if err != nil {
			c.log.Error(err, "Failed to get Hazelcast Reconcile")
			return DefaultDuration, err
		}
		c.RunParallel(opsapi.UpdateStatefulSets, "Successfully updated pet sets", c.NewReconcileFunc(dbCopy, rcl))
		return DefaultDuration, nil
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.RestartNodes) {
		// Need to delete PDB, as need to delete all pods except the last one.
		err = policy_util.DeletePodDisruptionBudget(context.TODO(), c.Client, types.NamespacedName{
			Namespace: c.db.Namespace,
			Name:      c.db.StatefulSetName(),
		})
		if err != nil && !kerr.IsNotFound(err) {
			c.log.Error(err, "failed to delete pod disruption budget")
			return DefaultDuration, err
		}

		podlistName := c.getPodsName()
		if len(podlistName) == 0 {
			return DefaultDuration, errors.New("Pod list is empty")
		}
		// Use RestartNodes for simultaneous restart of all pods
		c.RunParallel(opsapi.RestartNodes, "Successfully restarted all nodes", c.newConcurrentRestartFunc(podlistName, dbCopy))
		return DefaultDuration, nil
	}
	_, err = clientutil.CreateOrPatch(context.TODO(), c.KBClient, c.db, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*dbapi.Hazelcast)
		ret.Spec = dbCopy.Spec
		c.log.Info("hazelcast spec overwite", "spec.enableSSL", ret.Spec.EnableSSL)
		return ret
	})
	if err != nil {
		c.log.Error(err, "failed to patch the hazelcast object")
		return DefaultDuration, err
	}

	// resume and Change the opsapi request phase to "Successful".
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		if err := c.resumeHazelcast(); err != nil {
			return DefaultDuration, err
		}
		if err := c.updateHazelcastOpsRequestPhase(opsapi.Successful, "Successfully completed reconfigureTLS for Hazelcast.", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update status")
			return DefaultDuration, err
		}
	}
	return DefaultDuration, nil
}

func (c *hzOpsReqController) cleanUpCertificates(db *dbapi.Hazelcast) error {
	// remove unwanted certificates
	certList, err := c.CertManagerClient.CertmanagerV1().Certificates(c.db.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.Set(c.db.OffshootSelectors()).String(),
	})
	if err != nil {
		return err
	}

	for _, cert := range certList.Items {
		if owned, _ := core_util.IsOwnedBy(&cert, &db.ObjectMeta); !owned {
			continue
		}
		err := c.CertManagerClient.CertmanagerV1().Certificates(c.db.Namespace).Delete(context.TODO(), cert.Name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	secretList, err := c.Client.CoreV1().Secrets(c.db.Namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.Set(c.db.OffshootSelectors()).String(),
	})
	if err != nil {
		return err
	}

	for _, secret := range secretList.Items {
		if secret.Type != core.SecretTypeTLS {
			continue
		}

		if owned, _ := core_util.IsOwnedBy(&secret, &db.ObjectMeta); !owned {
			continue
		}
		err := c.Client.CoreV1().Secrets(c.db.Namespace).Delete(context.TODO(), secret.Name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}

	// remove keystoreCredSecret
	if db.Spec.KeystoreSecret != nil && db.Spec.KeystoreSecret.Name != "" {
		if _, err := c.Client.CoreV1().Secrets(db.Namespace).Get(context.TODO(), db.Spec.KeystoreSecret.Name, metav1.GetOptions{}); err != nil {
			if kerr.IsNotFound(err) {
				return nil
			} else {
				c.log.Error(err, fmt.Sprintf("Failed to get KeystoreCredSecret %s/%s", db.Namespace, db.Spec.KeystoreSecret.Name))
				return err
			}
		}
		if err := c.Client.CoreV1().Secrets(db.Namespace).Delete(context.TODO(), db.Spec.KeystoreSecret.Name, metav1.DeleteOptions{}); err != nil {
			c.log.Error(err, fmt.Sprintf("Failed to delete KeystoreCredSecret %s/%s", db.Namespace, db.Spec.KeystoreSecret.Name))
			return err
		}
	}
	return nil
}

// updateProbeSchemes updates the health probe schemes (HTTP/HTTPS) in the database spec
func (c *hzOpsReqController) updateProbeSchemes(db *dbapi.Hazelcast, scheme core.URIScheme) {
	// Update probe schemes in the database spec instead of directly modifying StatefulSet
	// The reconciler will handle updating the actual StatefulSet based on the database spec

	if db.Spec.PodTemplate.Spec.Containers == nil {
		// Initialize containers if not present
		db.Spec.PodTemplate.Spec.Containers = []core.Container{}
	}

	container := core_util.GetContainerByName(c.getHazelcastContainers(db), kubedb.HazelcastContainerName)

	updated := false
	if container != nil {

		// Update liveness probe
		if container.LivenessProbe != nil && container.LivenessProbe.HTTPGet != nil {
			if container.LivenessProbe.HTTPGet.Scheme != scheme {
				container.LivenessProbe.HTTPGet.Scheme = scheme
				updated = true
			}
		}

		// Update readiness probe
		if container.ReadinessProbe != nil && container.ReadinessProbe.HTTPGet != nil {
			if container.ReadinessProbe.HTTPGet.Scheme != scheme {
				container.ReadinessProbe.HTTPGet.Scheme = scheme
				updated = true
			}
		}
	}

	if updated {
		c.log.Info(fmt.Sprintf("Updated probe schemes to %s in database spec", scheme))
	} else {
		c.log.Info(fmt.Sprintf("No probe scheme updates needed, already using %s", scheme))
	}
}
