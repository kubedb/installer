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
	"time"

	"kubedb.dev/apimachinery/apis/kubedb"
	dbapi "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	opsapi "kubedb.dev/apimachinery/apis/ops/v1alpha1"

	pass "gomodules.xyz/password-generator"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	cu "kmodules.xyz/client-go/client"
	cutil "kmodules.xyz/client-go/conditions"
	policy_util "kmodules.xyz/client-go/policy"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *hzOpsReqController) RotateAuthentication() (time.Duration, error) {
	hz, err := c.UpdateHazelcastPhaseProgressingAndPauseReconcile(string(opsapi.HazelcastOpsRequestTypeRotateAuth), "hazelcast ops request has started to rotate auth for rmq nodes")
	if (err != nil) || (hz == RequeueDuration) {
		return hz, err
	}

	authConfig := c.req.Spec.Authentication
	dbcopy := c.db.DeepCopy()

	// Ensure AuthSecret is initialized before accessing its fields
	if dbcopy.Spec.AuthSecret == nil {
		dbcopy.Spec.AuthSecret = &dbapi.SecretReference{}
	}

	// Validate GetAuthSecretName method before calling it
	var authSecretName string
	if c.db != nil {
		authSecretName = c.db.GetAuthSecretName()
		if authSecretName == "" {
			c.log.Error(nil, "GetAuthSecretName returned empty string")
			return DefaultDuration, errors.New("GetAuthSecretName returned empty string")
		}
	} else {
		c.log.Error(nil, "c.db is nil when trying to get auth secret name")
		return DefaultDuration, errors.New("c.db is nil when trying to get auth secret name")
	}

	// Get old auth secret
	oldAuthSecret := &core.Secret{}
	err = c.KBClient.Get(context.TODO(), types.NamespacedName{
		Namespace: c.req.Namespace,
		Name:      authSecretName,
	}, oldAuthSecret)
	if err != nil {
		c.log.Error(err, "failed to get old auth secret")
		return DefaultDuration, err
	}

	var isUserProvided bool
	if authConfig != nil && authConfig.SecretRef != nil && authConfig.SecretRef.Name != "" {
		isUserProvided = true
		// Double-check that AuthSecret is still not nil before assignment
		if dbcopy.Spec.AuthSecret == nil {
			dbcopy.Spec.AuthSecret = &dbapi.SecretReference{}
		}
		dbcopy.Spec.AuthSecret.Name = authConfig.SecretRef.Name
		dbcopy.Spec.AuthSecret.ExternallyManaged = true
	}

	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateCredential) {
		var msg string
		if !isUserProvided {
			if err := c.generateAuthCredentials(c.db.GetAuthSecretName()); err != nil {
				return DefaultDuration, err
			}
			msg = "Successfully generated new credentials"
			c.log.Info(msg)
		} else {
			// Validate user provided authSecret
			if err := c.validateAuthSecret(authConfig.SecretRef.Name); err != nil {
				c.log.Error(err, "failed to validate user provided auth secret")
				return DefaultDuration, err
			}
			// Update user provided authSecret with prev secret credentials
			// e.g. username.prev, password.prev
			if err := c.updateWithOldCredentials(authConfig.SecretRef.Name, oldAuthSecret); err != nil {
				c.log.Error(err, "failed to update user provided auth secret")
				return DefaultDuration, err
			}
			msg = "Successfully referenced the user provided authSecret"
		}
		err := c.UpdateHazelcastOpsReqConditions(opsapi.UpdateCredential, msg)
		if err != nil {
			c.log.Error(err, "failed to update user provided auth secret")
			return DefaultDuration, err
		}
	}

	if cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateCredential) && !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
		c.log.Info("updating Petsets")
		rcl, err := c.NewHazelcastReconcile(dbcopy)
		if err != nil {
			return DefaultDuration, err
		}
		// Reconcile hazlecast resources
		if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
			c.RunParallel(opsapi.UpdateStatefulSets, "successfully reconciled the hazelcast with new auth credentials and configuration", c.NewReconcileFunc(dbcopy, rcl))
			return DefaultDuration, nil
		}
	}

	if cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.UpdateStatefulSets) {
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

			podNames := c.getPodsName()
			if len(podNames) == 0 {
				return DefaultDuration, errors.New("podlist is empty")
			}

			c.RunParallel(opsapi.RestartNodes, "Successfully restarted all nodes", c.newConcurrentRestartFunc(podNames, dbcopy))
			return DefaultDuration, nil
		}
	}

	if cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.RestartNodes) {
		// drop secret-already-updated key from the secret
		updatedSecret := &core.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: dbcopy.Namespace,
				Name:      dbcopy.GetAuthSecretName(),
			},
		}

		if _, err := cu.CreateOrPatch(context.TODO(), c.KBClient, updatedSecret, func(obj client.Object, createOp bool) client.Object {
			ret := obj.(*core.Secret)
			delete(ret.Annotations, opsapi.SecretAlreadyUpdatedAnnotation)
			return ret
		}); err != nil {
			c.log.Error(err, "failed to drop 'secret-already-updated' the secret")
			return DefaultDuration, err
		}

		activeFromTime, err := getActiveFromTimestamp(updatedSecret)
		if err != nil {
			c.log.Error(err, "failed to get activeFrom timestamp")
			return DefaultDuration, err
		}
		dbcopy.Spec.AuthSecret.ActiveFrom = activeFromTime

		// update hazelcast object with new authsecret
		_, err = cu.CreateOrPatch(context.TODO(), c.KBClient, c.db, func(obj client.Object, createOp bool) client.Object {
			ret := obj.(*dbapi.Hazelcast)
			ret.Spec.AuthSecret = dbcopy.Spec.AuthSecret
			return ret
		})
		if err != nil {
			c.log.Error(err, "failed to update hazelcast object with new auth secret")
			return DefaultDuration, err
		}
	}
	if !cutil.IsConditionTrue(c.req.Status.Conditions, opsapi.Successful) {
		if err := c.resumeHazelcast(); err != nil {
			return DefaultDuration, err
		}
		if err := c.updateHazelcastOpsRequestPhase(opsapi.Successful, "Successfully completed reconfigure Ignite", opsapi.OpsRequestPhaseSuccessful); err != nil {
			c.log.Error(err, "failed to update hazelcast ops request phase")
			return DefaultDuration, err
		}
	}
	return DefaultDuration, nil
}

func (c *hzOpsReqController) generateAuthCredentials(authSecretName string) error {
	_, err := cu.CreateOrPatch(context.TODO(), c.KBClient, &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      authSecretName,
			Namespace: c.req.Namespace,
		},
	}, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*core.Secret)
		if ret.Annotations == nil {
			ret.Annotations = make(map[string]string)
		}
		if _, exists := ret.Annotations[opsapi.SecretAlreadyUpdatedAnnotation]; exists {
			return ret
		}
		ret.Annotations[opsapi.SecretAlreadyUpdatedAnnotation] = "true"
		ret.Annotations[kubedb.AuthActiveFromAnnotation] = time.Now().Format(time.RFC3339)

		ret.Data[opsapi.BasicAuthPreviousUsernameKey] = ret.Data[core.BasicAuthUsernameKey]
		ret.Data[opsapi.BasicAuthPreviousPasswordKey] = ret.Data[core.BasicAuthPasswordKey]

		ret.Data[core.BasicAuthPasswordKey] = []byte(pass.Generate(kubedb.DefaultPasswordLength))
		return ret
	})
	if err != nil {
		c.log.Error(err, "failed to generate auth secret")
		return err
	}
	return nil
}

func (c *hzOpsReqController) validateAuthSecret(authSecretName string) error {
	secret := &core.Secret{}
	err := c.KBClient.Get(context.TODO(), types.NamespacedName{
		Namespace: c.db.Namespace,
		Name:      authSecretName,
	}, secret)
	if err != nil {
		if kerr.IsNotFound(err) {
			return errors.New("secret not found")
		}
		return err
	}
	if _, exists := secret.Data[core.BasicAuthUsernameKey]; !exists {
		return errors.New("username key not found in user provided secret")
	}
	if _, exists := secret.Data[core.BasicAuthPasswordKey]; !exists {
		return errors.New("password key not found in user provided secret")
	}
	return nil
}

func (c *hzOpsReqController) updateWithOldCredentials(authSecretName string, oldAuthSecret *core.Secret) error {
	_, err := cu.CreateOrPatch(context.TODO(), c.KBClient, &core.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      authSecretName,
			Namespace: c.req.Namespace,
		},
	}, func(obj client.Object, createOp bool) client.Object {
		ret := obj.(*core.Secret)
		if _, exists := ret.Annotations[opsapi.SecretAlreadyUpdatedAnnotation]; exists {
			return ret
		}
		if ret.Annotations == nil {
			ret.Annotations = make(map[string]string)
		}
		ret.Annotations[opsapi.SecretAlreadyUpdatedAnnotation] = "true"
		ret.Annotations[kubedb.AuthActiveFromAnnotation] = time.Now().Format(time.RFC3339)
		ret.Data[opsapi.BasicAuthPreviousUsernameKey] = oldAuthSecret.Data[core.BasicAuthUsernameKey]
		ret.Data[opsapi.BasicAuthPreviousPasswordKey] = oldAuthSecret.Data[core.BasicAuthPasswordKey]
		return ret
	})
	if err != nil {
		c.log.Error(err, "failed to update user provided auth secret")
		return err
	}
	return nil
}

func getActiveFromTimestamp(secretName *core.Secret) (*metav1.Time, error) {
	if val, exists := secretName.Annotations[kubedb.AuthActiveFromAnnotation]; exists {
		t, err := time.Parse(time.RFC3339, val)
		if err != nil {
			return nil, err
		}
		return &metav1.Time{Time: t}, nil
	}
	return nil, nil
}
