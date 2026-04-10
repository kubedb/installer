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
	"time"

	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"

	"github.com/pkg/errors"
	password "gomodules.xyz/password-generator"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kutil "kmodules.xyz/client-go"
	clientutil "kmodules.xyz/client-go/client"
	coreutil "kmodules.xyz/client-go/core/v1"
	metautil "kmodules.xyz/client-go/meta"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (rs *ReconcileState) EnsureSecrets() error {
	AuthSecret, err := rs.ensureAdminSecret()
	if err != nil {
		rs.log.Error(err, "Failed to ensure Hazelcast Auth Secret")
		return err
	}

	keystoreSecret, err := rs.EnsureKeystoreSecret()
	if err != nil {
		rs.log.Error(err, "Failed to ensure keystore Secret")
		return err
	}

	_, err = clientutil.CreateOrPatch(rs.ctx, rs.KBClient, &api.Hazelcast{
		ObjectMeta: meta.ObjectMeta{
			Name:      rs.db.Name,
			Namespace: rs.db.Namespace,
		},
	}, func(obj client.Object, createOp bool) client.Object {
		in := obj.(*api.Hazelcast)
		if keystoreSecret != nil {
			in.Spec.KeystoreSecret = &core.LocalObjectReference{
				Name: keystoreSecret.Name,
			}
		}
		if in.Spec.AuthSecret == nil {
			in.Spec.AuthSecret = &api.SecretReference{
				TypedLocalObjectReference: appcat.TypedLocalObjectReference{},
			}
		}
		in.Spec.AuthSecret.Name = AuthSecret.Name
		return in
	})
	if err != nil {
		rs.log.Error(err, "Failed to patch Hazelcast with KeystoreCred and AuthSecret")
		return err
	}

	_, err = rs.ensureConfigSecret()
	if err != nil {
		rs.log.Error(err, "Failed to create Hazelcast Config Secret")
		return err
	}

	return nil
}

func (rs *ReconcileState) ensureConfigSecret() (*core.Secret, error) {
	configSecret, err := rs.createConfigSecret()
	if err != nil {
		return nil, err
	}

	err = rs.validateAndSyncHazelcastConfigSecretLabels(configSecret, string(configSecret.Data[kubedb.HazelcastSecretKey]), string(configSecret.Data[kubedb.HazelcastClientSecretKey]))
	if err != nil {
		return nil, err
	}

	return configSecret, nil
}

func (rs *ReconcileState) createConfigSecret() (*core.Secret, error) {
	config, clientConfig, err := rs.getHazelcastConfig()
	if err != nil {
		rs.log.Error(err, "Failed to get config secret")
		return nil, err
	}
	data := map[string]string{
		kubedb.HazelcastSecretKey:       config,
		kubedb.HazelcastClientSecretKey: clientConfig,
	}

	secret := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Name:      rs.db.ConfigSecretName(),
			Namespace: rs.db.Namespace,
		},
	}

	v, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, secret, func(obj client.Object, createOp bool) client.Object {
		in := obj.(*core.Secret)
		if rs.db.Spec.Configuration != nil && rs.db.Spec.Configuration.Inline != nil {
			for key, value := range rs.db.Spec.Configuration.Inline {
				filename := kubedb.InlineConfigKeyPrefix + "-" + key
				data[filename] = string(value)
			}
		}

		in.StringData = data
		if createOp {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
		}
		return in
	})

	if v == kutil.VerbCreated {
		rs.log.Info(fmt.Sprintf("Config secret: %s/%s created", rs.db.Namespace, rs.db.ConfigSecretName()))
	}
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func (rs *ReconcileState) validateAndSyncHazelcastConfigSecretLabels(secret *core.Secret, config, clientConfig string) error {
	if value, exist := secret.Data[kubedb.HazelcastSecretKey]; !exist || len(value) == 0 {
		return errors.New("hazelcast.yaml is missing")
	} else if config != "" && string(value) != config {
		return errors.Errorf("hazelcast.yaml must be %s but it is %s", config, string(value))
	}

	if value, exist := secret.Data[kubedb.HazelcastClientSecretKey]; !exist || len(value) == 0 {
		return errors.New("hazelcast-client.yaml is missing")
	} else if clientConfig != "" && string(value) != clientConfig {
		return errors.Errorf("hazelcast-client.yaml must be %s but it is %s", clientConfig, string(value))
	}

	// If secret is owned by this Hazelcast object, update the labels.
	// Labels may hold important information, should be synced.
	owned, _ := coreutil.IsOwnedBy(secret, rs.db)
	if owned {
		// sync labels
		_, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, secret, func(obj client.Object, createOp bool) client.Object {
			in := obj.(*core.Secret)
			in.Labels = metautil.OverwriteKeys(in.Labels, rs.db.OffshootLabels())
			return in
		})
		if err != nil {
			rs.log.Error(err, fmt.Sprintf("Failed to sync Labels for secret %s/%s", rs.db.Namespace, secret.GetName()))
			return err
		}
	}
	return nil
}

func (rs *ReconcileState) ensureAdminSecret() (*core.Secret, error) {
	if rs.db.Spec.DisableSecurity {
		return nil, nil
	}
	if rs.db.Spec.AuthSecret != nil && rs.db.Spec.AuthSecret.ExternallyManaged {
		return rs.ensureExternalAdminSecret()
	}
	return rs.ensureInternalAdminSecret()
}

func (rs *ReconcileState) ensureInternalAdminSecret() (*core.Secret, error) {
	adminSecret := &core.Secret{}
	if err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
		Name:      rs.db.GetAuthSecretName(),
		Namespace: rs.db.Namespace,
	}, adminSecret); err != nil {
		if kerr.IsNotFound(err) {
			rs.log.Info("not found admin secret")
			adminSecret, err = rs.createAdminSecret()
			if err != nil {
				rs.log.Error(err, "Failed to create Admin Secret")
				return nil, err
			}
			return adminSecret, nil
		} else {
			rs.log.Error(err, fmt.Sprintf("Failed to get Admin Secret= %s/%s", rs.db.Namespace, rs.db.GetAuthSecretName()))
			return nil, err
		}
	}

	err := rs.validateAndSyncHazelcastAuthSecretLabels(adminSecret, string(adminSecret.Data[core.BasicAuthUsernameKey]))
	if err != nil {
		return nil, err
	}

	// If 'auth.activeFrom' annotation is not present that means user has provided the secret keeping the .authSecret.externallyManaged=false
	if _, exists := adminSecret.Annotations[kubedb.AuthActiveFromAnnotation]; !exists {
		_, err = clientutil.CreateOrPatch(rs.ctx, rs.KBClient, adminSecret, func(obj client.Object, createOp bool) client.Object {
			in := obj.(*core.Secret)
			if in.Annotations == nil {
				in.Annotations = make(map[string]string)
			}
			in.Annotations[kubedb.AuthActiveFromAnnotation] = in.CreationTimestamp.Format(time.RFC3339)
			return in
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to patch annotation for hazelcast authsecret")
		}
	}

	return adminSecret, nil
}

func (rs *ReconcileState) ensureExternalAdminSecret() (*core.Secret, error) {
	if rs.db.Spec.AuthSecret.Name == "" {
		return nil, fmt.Errorf("externally managed auth secret name is missing for hazelcast %s/%s", rs.db.Namespace, rs.db.Name)
	}

	adminSecret := &core.Secret{}
	if err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
		Name:      rs.db.GetAuthSecretName(),
		Namespace: rs.db.Namespace,
	}, adminSecret); err != nil {
		if kerr.IsNotFound(err) {
			return nil, fmt.Errorf("externally managed auth secret \"%s\" not found for Hazelcast %s/%s", rs.db.GetAuthSecretName(), rs.db.Namespace, rs.db.Name)
		}
		return nil, err
	}

	err := rs.validateAndSyncHazelcastAuthSecretLabels(adminSecret, string(adminSecret.Data[core.BasicAuthUsernameKey]))
	if err != nil {
		return nil, err
	}

	// If 'auth.activeFrom' annotation is not present that means user has provided the secret keeping the .authSecret.externallyManaged=false
	if _, exists := adminSecret.Annotations[kubedb.AuthActiveFromAnnotation]; !exists {
		_, err = clientutil.CreateOrPatch(rs.ctx, rs.KBClient, adminSecret, func(obj client.Object, createOp bool) client.Object {
			in := obj.(*core.Secret)
			if in.Annotations == nil {
				in.Annotations = make(map[string]string)
			}
			in.Annotations[kubedb.AuthActiveFromAnnotation] = time.Now().Format(time.RFC3339)
			return in
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to patch annotation for externally managed hazelcast authsecret")
		}
	}

	return adminSecret, nil
}

func (rs *ReconcileState) EnsureKeystoreSecret() (*core.Secret, error) {
	if rs.db.Spec.DisableSecurity || !rs.db.Spec.EnableSSL {
		return nil, nil
	}
	keystoreSecret := &core.Secret{}
	name := rs.db.HazelcastSecretName("keystore-cred")
	if rs.db.Spec.KeystoreSecret != nil {
		name = rs.db.Spec.KeystoreSecret.Name
	}

	if err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
		Name:      name,
		Namespace: rs.db.Namespace,
	}, keystoreSecret); err != nil {
		if kerr.IsNotFound(err) {

			keystoreSecret, err = rs.createKeystoreSecret()
			if err != nil {
				rs.log.Error(err, "Failed to create Keystore Secret")
				return nil, err
			}
		} else {
			rs.log.Error(err, fmt.Sprintf("Failed to get Keystore Secret %s/%s", rs.db.Namespace, name))
			return nil, err
		}
	}

	err := rs.validateAndSyncKeystoreSecretLabels(keystoreSecret)
	if err != nil {
		return nil, err
	}

	return keystoreSecret, nil
}

func (rs *ReconcileState) validateAndSyncKeystoreSecretLabels(secret *core.Secret) error {
	if value, exist := secret.Data[kubedb.HazelcastKeystorePassKey]; !exist || len(value) == 0 {
		return errors.New("password is missing")
	}

	// If secret is owned by this Hazelcast object, update the labels.
	// Labels may hold important information, should be synced.
	owned, _ := coreutil.IsOwnedBy(secret, rs.db)
	if owned {
		// sync labels
		_, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, secret, func(obj client.Object, createOp bool) client.Object {
			in := obj.(*core.Secret)
			in.Labels = metautil.OverwriteKeys(in.Labels, rs.db.OffshootLabels())
			return in
		})
		if err != nil {
			rs.log.Error(err, fmt.Sprintf("Failed to sync Labels for secret %s/%s", rs.db.Namespace, secret.GetName()))
			return err
		}
	}
	return nil
}

func (rs *ReconcileState) createKeystoreSecret() (*core.Secret, error) {
	name := rs.db.HazelcastSecretName("keystore-cred")
	if rs.db.Spec.KeystoreSecret != nil {
		name = rs.db.Spec.KeystoreSecret.Name
	}
	pass := password.Generate(kubedb.DefaultPasswordLength)
	data := map[string][]byte{
		kubedb.HazelcastKeystorePassKey: []byte(pass),
	}

	secret := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Name:      name,
			Namespace: rs.db.Namespace,
		},
	}

	v, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, secret, func(obj client.Object, createOp bool) client.Object {
		in := obj.(*core.Secret)
		in.Data = data
		if createOp {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
		}
		return in
	})

	if v == kutil.VerbCreated {
		rs.log.Info(fmt.Sprintf("Keystore secret: %s/%s created", rs.db.Namespace, name))
	}
	if err != nil {
		return nil, err
	}
	return secret, nil
}

func (rs *ReconcileState) createAdminSecret() (*core.Secret, error) {
	pass := password.Generate(kubedb.DefaultPasswordLength)
	data := map[string][]byte{
		core.BasicAuthUsernameKey: []byte(kubedb.HazelcastAdmin),
		core.BasicAuthPasswordKey: []byte(pass),
	}

	secret := &core.Secret{
		ObjectMeta: meta.ObjectMeta{
			Name:      rs.db.GetAuthSecretName(),
			Namespace: rs.db.Namespace,
		},
	}

	v, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, secret, func(obj client.Object, createOp bool) client.Object {
		in := obj.(*core.Secret)
		in.Data = data
		in.Type = core.SecretTypeBasicAuth
		if createOp {
			coreutil.EnsureOwnerReference(&in.ObjectMeta, rs.db.Owner())
		}
		return in
	})

	if v == kutil.VerbCreated {
		rs.log.Info(fmt.Sprintf("Auth secret: %s/%s created", rs.db.Namespace, rs.db.GetAuthSecretName()))
	}
	if err != nil {
		return nil, err
	}

	// If 'auth.activeFrom' annotation is not present that means user has provided the secret keeping the .authSecret.externallyManaged=false
	if _, exists := secret.Annotations[kubedb.AuthActiveFromAnnotation]; !exists {
		_, err = clientutil.CreateOrPatch(rs.ctx, rs.KBClient, secret, func(obj client.Object, createOp bool) client.Object {
			in := obj.(*core.Secret)
			if in.Annotations == nil {
				in.Annotations = make(map[string]string)
			}
			in.Annotations[kubedb.AuthActiveFromAnnotation] = in.CreationTimestamp.Format(time.RFC3339)
			return in
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to patch annotation for hazelcast authsecret")
		}
	}
	return secret, nil
}

func (rs *ReconcileState) validateAndSyncHazelcastAuthSecretLabels(secret *core.Secret, username string) error {
	if value, exist := secret.Data[core.BasicAuthUsernameKey]; !exist || len(value) == 0 {
		return errors.New("username is missing")
	} else if username != "" && string(value) != username {
		return errors.Errorf("username must be %s\n", username)
	}

	if value, exist := secret.Data[core.BasicAuthPasswordKey]; !exist || len(value) == 0 {
		return errors.New("password is missing")
	}

	// If secret is owned by this Hazelcast object, update the labels.
	// Labels may hold important information, should be synced.
	owned, _ := coreutil.IsOwnedBy(secret, rs.db)
	if owned {
		// sync labels
		_, err := clientutil.CreateOrPatch(rs.ctx, rs.KBClient, secret, func(obj client.Object, createOp bool) client.Object {
			in := obj.(*core.Secret)
			in.Labels = metautil.OverwriteKeys(in.Labels, rs.db.OffshootLabels())
			return in
		})
		if err != nil {
			rs.log.Error(err, fmt.Sprintf("Failed to sync Labels for secret %s/%s", rs.db.Namespace, secret.GetName()))
			return err
		}
	}
	return nil
}
