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
	"errors"
	"fmt"

	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"kubedb.dev/apimachinery/pkg/lib"

	cm_api "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	cutil "kmodules.xyz/client-go/conditions"
)

func (c *Controller) manageHazelcastEvent(k any) error {
	key := k.(string)
	c.log.Info("started processing", "key", key)
	obj, exists, err := c.dbInformer.GetIndexer().GetByKey(key)
	if err != nil {
		c.log.Error(err, fmt.Sprintf("Fetching object with key %s from store failed", key))
		return err
	}

	if !exists {
		c.log.Info(fmt.Sprintf("Hazelcast %s does not exist anymore\n", key))
	} else {
		if obj == nil {
			c.log.Error(nil, "Fetched object is nil", "key", key)
			return fmt.Errorf("fetched object is nil for key %s", key)
		}
		db, ok := obj.(*dbapi.Hazelcast)
		if !ok {
			c.log.Error(nil, "Object is not of type *dbapi.Hazelcast", "key", key)
			return fmt.Errorf("object is not of type *dbapi.Hazelcast for key %s", key)
		}

		if db.DeletionTimestamp == nil {
			db = db.DeepCopy()
			if cutil.IsConditionTrue(db.Status.Conditions, kubedb.DatabasePaused) {
				return nil
			}
			if err := c.manageTLS(db); err != nil {
				c.log.Error(err, "Failed to manage TLS for Hazelcast")
				c.pushFailureEvent(db, err.Error())
				return err
			}
		}
	}
	return nil
}

func (c *Controller) manageTLS(db *dbapi.Hazelcast) error {
	if db.Spec.DisableSecurity {
		c.log.Info("Hazelcast security is disabled")
		return nil
	}
	if db.Spec.TLS == nil || db.Spec.TLS.IssuerRef == nil {
		c.log.V(1).Info("Hazelcast TLS/IssuerRef is not configured")
		return nil
	}

	var svc core.Service
	err := c.KBClient.Get(context.TODO(), types.NamespacedName{
		Namespace: db.Namespace,
		Name:      db.GoverningServiceName(),
	}, &svc)
	if err != nil {
		c.log.Info("Failed to get database service for Hazelcast")
		return err
	}

	if !lib.IsServiceReady(c.Client.CoreV1(), svc.ObjectMeta) {
		c.log.Info("Hazelcast database service is not ready")
		return nil
	}

	if db.Spec.KeystoreSecret == nil {
		c.log.Info("KeystoreCredSecret is not set")
		return nil
	}

	var keystoreCredSecret core.Secret
	if err := c.KBClient.Get(context.TODO(), types.NamespacedName{
		Name:      db.Spec.KeystoreSecret.Name,
		Namespace: db.Namespace,
	}, &keystoreCredSecret); err != nil {
		if kerr.IsNotFound(err) {
			c.log.Info("KeystoreCredSecret is not found")
			return nil
		} else {
			c.log.Error(err, fmt.Sprintf("Failed to get KeystoreCredSecret %s/%s", db.Namespace, db.Spec.KeystoreSecret.Name))
			return err
		}
	}

	switch db.Spec.TLS.IssuerRef.Kind {
	case cm_api.IssuerKind:
		_, err := c.CertManagerClient.CertmanagerV1().Issuers(db.Namespace).Get(context.TODO(), db.Spec.TLS.IssuerRef.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
	case cm_api.ClusterIssuerKind:
		_, err := c.CertManagerClient.CertmanagerV1().ClusterIssuers().Get(context.TODO(), db.Spec.TLS.IssuerRef.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
	default:
		return errors.New("db.Spec.TLS.Client.IssuerRef.Kind is not either Issuer or ClusterIssuer")
	}

	if err := c.manageServerCert(db); err != nil {
		return err
	}

	if err := c.manageClientCert(db); err != nil {
		return err
	}

	return nil
}
