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
	"errors"
	"fmt"
	"regexp"

	"kubedb.dev/apimachinery/apis/kubedb"

	"gopkg.in/yaml.v2"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type Kubernetes struct {
	Enabled     bool   `json:"enabled" yaml:"enabled"`
	ServiceName string `json:"service-name" yaml:"service-name"`
	Namespace   string `json:"namespace" yaml:"namespace"`
}

type EndpointGroups struct {
	ClusterRead  map[string]bool `json:"CLUSTER_READ" yaml:"CLUSTER_READ"`
	ClusterWrite map[string]bool `json:"CLUSTER_WRITE" yaml:"CLUSTER_WRITE"`
	HealthCheck  map[string]bool `json:"HEALTH_CHECK" yaml:"HEALTH_CHECK"`
	Persistence  map[string]bool `json:"PERSISTENCE" yaml:"PERSISTENCE"`
	Data         map[string]bool `json:"DATA" yaml:"DATA"`
}

type RestAPI struct {
	Enabled        bool           `json:"enabled" yaml:"enabled"`
	EndpointGroups EndpointGroups `json:"endpoint-groups" yaml:"endpoint-groups"`
}
type Join struct {
	Kubernetes Kubernetes `json:"kubernetes" yaml:"kubernetes"`
}

type Network struct {
	Join    Join    `json:"join" yaml:"join"`
	RestAPI RestAPI `json:"rest-api" yaml:"rest-api"`
	SSL     SSL     `json:"ssl,omitempty" yaml:"ssl,omitempty"`
}

type User struct {
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
	Roles    []string `json:"roles" yaml:"roles"`
}

type Map struct {
	Default Defaulted `json:"default" yaml:"default"`
}

type Cache struct {
	Default Defaulted `json:"default" yaml:"default"`
}

type Queue struct {
	Default Defaulted `json:"default" yaml:"default"`
}

type Defaulted struct {
	DataPersistence map[string]bool `json:"data-persistence" yaml:"data-persistence"`
}

type AuthenticationName struct {
	Realm string `json:"realm" yaml:"realm"`
}
type ALL struct {
	Principal string `json:"principal" yaml:"principal"`
}

type ClientPermission struct {
	All ALL `json:"all" yaml:"all"`
}

type Simple struct {
	Users []User `json:"users" yaml:"users"`
}

type Authentication struct {
	Simple Simple `json:"simple" yaml:"simple"`
}

type Identity struct {
	UsernamePassword UsernamePassword `json:"username-password" yaml:"username-password"`
}

type UsernamePassword struct {
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

type AccessControlService struct {
	FactoryClassName string `json:"factory-class-name" yaml:"factory-class-name"`
}

type Realm struct {
	Name                 string               `json:"name" yaml:"name"`
	Authentication       Authentication       `json:"authentication" yaml:"authentication"`
	Identity             Identity             `json:"identity" yaml:"identity"`
	AccessControlService AccessControlService `json:"access-control-service,omitempty" yaml:"access-control-service,omitempty"`
}

type Rest struct {
	Enabled              bool    `json:"enabled" yaml:"enabled"`
	Port                 int     `json:"port,omitempty" yaml:"port,omitempty"`
	SecurityRealm        string  `json:"security-realm,omitempty" yaml:"security-realm,omitempty"`
	TokenValiditySeconds int     `json:"token-validity-seconds,omitempty" yaml:"token-validity-seconds,omitempty"`
	SSL                  SSLREST `json:"ssl,omitempty" yaml:"ssl,omitempty"`
}

type Security struct {
	Enabled              bool               `json:"enabled" yaml:"enabled"`
	Realms               []Realm            `json:"realms" yaml:"realms"`
	ClientAuthentication AuthenticationName `json:"client-authentication" yaml:"client-authentication"`
	MemberAuthentication AuthenticationName `json:"member-authentication" yaml:"member-authentication"`
	ClientPermission     ClientPermission   `json:"client-permissions" yaml:"client-permissions"`
}

type Persistence struct {
	Enabled                   bool   `json:"enabled" yaml:"enabled"`
	BaseDIR                   string `json:"base-dir" yaml:"base-dir"`
	BackupDir                 string `json:"backup-dir,omitempty" yaml:"backup-dir,omitempty"`
	ValidationTimeoutSeconds  int    `json:"validation-timeout-seconds" yaml:"validation-timeout-seconds"`
	DataLoadTimeoutSeconds    int    `json:"data-load-timeout-seconds" yaml:"data-load-timeout-seconds"`
	AutoRemoveStaleData       bool   `json:"auto-remove-stale-data" yaml:"auto-remove-stale-data"`
	ClusterDataRecoveryPolicy string `json:"cluster-data-recovery-policy" yaml:"cluster-data-recovery-policy"`
}

type Jet struct {
	Enabled bool `json:"enabled" yaml:"enabled"`
}

type SQL struct {
	CatalogPersistenceEnabled bool `json:"catalog-persistence-enabled" yaml:"catalog-persistence-enabled"`
}

type Hazelcast struct {
	ClusterName string      `json:"cluster-name,omitempty" yaml:"cluster-name,omitempty"`
	Network     Network     `json:"network,omitempty" yaml:"network,omitempty"`
	Rest        Rest        `json:"rest,omitempty" yaml:"rest,omitempty"`
	Security    Security    `json:"security,omitempty" yaml:"security,omitempty"`
	Persistence Persistence `json:"persistence,omitempty" yaml:"persistence,omitempty"`
	Jet         Jet         `json:"jet,omitempty" yaml:"jet,omitempty"`
	SQL         SQL         `json:"sql,omitempty" yaml:"sql,omitempty"`
	Map         Map         `json:"map,omitempty" yaml:"map,omitempty"`
	Cache       Cache       `json:"cache,omitempty" yaml:"cache,omitempty"`
}

type SSLProperties struct {
	Protocol           string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	TrustStore         string `json:"trustStore,omitempty" yaml:"trustStore,omitempty"`
	TrustStorePassword string `json:"trustStorePassword,omitempty" yaml:"trustStorePassword,omitempty"`
	TrustStoreType     string `json:"trustStoreType,omitempty" yaml:"trustStoreType,omitempty"`
	KeyStore           string `json:"keyStore,omitempty" yaml:"keyStore,omitempty"`
	KeyStorePassword   string `json:"keyStorePassword,omitempty" yaml:"keyStorePassword,omitempty"`
	KeyStoreType       string `json:"keyStoreType,omitempty" yaml:"keyStoreType,omitempty"`
}

type SSL struct {
	Enabled    bool          `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Properties SSLProperties `json:"properties,omitempty" yaml:"properties,omitempty"`
}

type SSLREST struct {
	Enabled             bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ClientAuth          string `json:"client-auth,omitempty" yaml:"client-auth,omitempty"`
	Ciphers             string `json:"ciphers,omitempty" yaml:"ciphers,omitempty"`
	EnabledProtocols    string `json:"enabled-protocols,omitempty" yaml:"enabled-protocols,omitempty"`
	KeyAlias            string `json:"key-alias,omitempty" yaml:"key-alias,omitempty"`
	KeyPassword         string `json:"key-password,omitempty" yaml:"key-password,omitempty"`
	KeyStore            string `json:"key-store,omitempty" yaml:"key-store,omitempty"`
	KeyStorePassword    string `json:"key-store-password,omitempty" yaml:"key-store-password,omitempty"`
	KeyStoreType        string `json:"key-store-type,omitempty" yaml:"key-store-type,omitempty"`
	KeyStoreProvider    string `json:"key-store-provider,omitempty" yaml:"key-store-provider,omitempty"`
	TrustStore          string `json:"trust-store,omitempty" yaml:"trust-store,omitempty"`
	TrustStorePassword  string `json:"trust-store-password,omitempty" yaml:"trust-store-password,omitempty"`
	TrustStoreType      string `json:"trust-store-type,omitempty" yaml:"trust-store-type,omitempty"`
	TrustStoreProvider  string `json:"trust-store-provider,omitempty" yaml:"trust-store-provider,omitempty"`
	Protocol            string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Certificate         string `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	CertificateKey      string `json:"certificate-key,omitempty" yaml:"certificate-key,omitempty"`
	TrustCertificate    string `json:"trust-certificate,omitempty" yaml:"trust-certificate,omitempty"`
	TrustCertificateKey string `json:"trust-certificate-key,omitempty" yaml:"trust-certificate-key,omitempty"`
}
type HazelcastClient struct {
	ClusterName    string         `json:"cluster-name,omitempty" yaml:"cluster-name,omitempty"`
	ClientNetwork  ClientNetwork  `json:"network" yaml:"network"`
	ClientSecurity ClientSecurity `json:"security,omitempty" yaml:"security,omitempty"`
}

type ClientNetwork struct {
	Kubernetes Kubernetes `json:"kubernetes" yaml:"kubernetes"`
	SSL        SSL        `json:"ssl,omitempty" yaml:"ssl,omitempty"`
}

type ClientSecurity struct {
	UsernamePassword UsernamePassword `json:"username-password,omitempty" yaml:"username-password,omitempty"`
}

type HazelcastConfig struct {
	Hazelcast Hazelcast `json:"hazelcast" yaml:"hazelcast"`
}

type HazelcastClientConfig struct {
	HazelcastClient HazelcastClient `json:"hazelcast-client" yaml:"hazelcast-client"`
}

// addQuotesToCredentials adds single quotes around username and password values in YAML
func addQuotesToCredentials(yamlContent string) string {
	// Pattern to match username: value (with optional whitespace)
	usernamePattern := regexp.MustCompile(`(\s+username:\s+)([^\s'"]+)`)
	// Pattern to match password: value (with optional whitespace)
	passwordPattern := regexp.MustCompile(`(\s+password:\s+)([^\s'"]+)`)

	// Add quotes around username values
	result := usernamePattern.ReplaceAllString(yamlContent, "${1}'${2}'")
	// Add quotes around password values
	result = passwordPattern.ReplaceAllString(result, "${1}'${2}'")

	return result
}

func (rs *ReconcileState) getHazelcastConfig() (string, string, error) {
	config := HazelcastConfig{
		Hazelcast: Hazelcast{
			Map: Map{
				Default: Defaulted{
					DataPersistence: map[string]bool{"enabled": true},
				},
			},
			Cache: Cache{
				Default: Defaulted{
					DataPersistence: map[string]bool{"enabled": true},
				},
			},
			ClusterName: fmt.Sprintf("%s/%s", rs.db.Namespace, rs.db.Name),
			Network: Network{
				Join: Join{
					Kubernetes: Kubernetes{
						Enabled:     true,
						ServiceName: rs.db.GoverningServiceName(),
						Namespace:   rs.db.Namespace,
					},
				},
				RestAPI: RestAPI{
					Enabled: true,
					EndpointGroups: EndpointGroups{
						ClusterRead:  map[string]bool{"enabled": true},
						ClusterWrite: map[string]bool{"enabled": true},
						HealthCheck:  map[string]bool{"enabled": true},
						Persistence:  map[string]bool{"enabled": true},
						Data:         map[string]bool{"enabled": true},
					},
				},
			},
			Rest: Rest{
				Enabled:              true,
				Port:                 8443,  // Enterprise REST API port
				TokenValiditySeconds: 18000, // 5 hours
			},
			// Persistence section commented out as MapStore will be used instead
			Persistence: Persistence{
				Enabled: true,
				BaseDIR: kubedb.HazelcastDataDir,
				// BackupDir:                 kubedb.HazelcastBackupDir,
				ValidationTimeoutSeconds:  120,
				DataLoadTimeoutSeconds:    900,
				AutoRemoveStaleData:       true,
				ClusterDataRecoveryPolicy: "PARTIAL_RECOVERY_MOST_COMPLETE",
			},
			Jet: Jet{
				Enabled: true,
			},
			SQL: SQL{
				CatalogPersistenceEnabled: true,
			},
		},
	}

	clientConfig := HazelcastClientConfig{
		HazelcastClient: HazelcastClient{
			ClusterName: fmt.Sprintf("%s/%s", rs.db.Namespace, rs.db.Name),
			ClientNetwork: ClientNetwork{
				Kubernetes: Kubernetes{
					Enabled:     true,
					ServiceName: rs.db.GoverningServiceName(),
					Namespace:   rs.db.Namespace,
				},
			},
		},
	}

	if !rs.db.Spec.DisableSecurity {
		secret := &core.Secret{}
		err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
			Name:      rs.db.GetAuthSecretName(),
			Namespace: rs.db.Namespace,
		}, secret)
		if err != nil {
			return "", "", err
		}

		security := Security{
			Enabled: true,
			Realms: []Realm{
				{
					Name: "simpleRealm",
					Authentication: Authentication{
						Simple: Simple{
							Users: []User{
								{
									Username: string(secret.Data[core.BasicAuthUsernameKey]),
									Password: string(secret.Data[core.BasicAuthPasswordKey]),
									Roles:    []string{"admin"},
								},
							},
						},
					},
					Identity: Identity{
						UsernamePassword: UsernamePassword{
							Username: string(secret.Data[core.BasicAuthUsernameKey]),
							Password: string(secret.Data[core.BasicAuthPasswordKey]),
						},
					},
					AccessControlService: AccessControlService{
						FactoryClassName: "com.hazelcast.internal.rest.access.DefaultAccessControlServiceFactory",
					},
				},
			},
			ClientAuthentication: AuthenticationName{
				Realm: "simpleRealm",
			},
			MemberAuthentication: AuthenticationName{
				Realm: "simpleRealm",
			},
			ClientPermission: ClientPermission{
				All: ALL{
					Principal: "admin",
				},
			},
		}

		clientSecurity := ClientSecurity{
			UsernamePassword: UsernamePassword{
				Username: string(secret.Data[core.BasicAuthUsernameKey]),
				Password: string(secret.Data[core.BasicAuthPasswordKey]),
			},
		}

		config.Hazelcast.Security = security
		clientConfig.HazelcastClient.ClientSecurity = clientSecurity

		config.Hazelcast.Rest.SecurityRealm = "simpleRealm"
	}

	if rs.db.Spec.EnableSSL {
		if rs.db.Spec.KeystoreSecret == nil {
			return "", "", errors.New("keystore cred is not defined")
		}
		var secret core.Secret
		err := rs.KBClient.Get(rs.ctx, types.NamespacedName{
			Name:      rs.db.Spec.KeystoreSecret.Name,
			Namespace: rs.db.Namespace,
		}, &secret)
		if err != nil {
			rs.log.Error(err, "Failed to get keystore secret")
			return "", "", err
		}
		ssl := SSL{
			Enabled: true,
			Properties: SSLProperties{
				Protocol:           "TLSv1.2",
				TrustStore:         kubedb.HazelcastServerTruststoreFile,
				TrustStorePassword: string(secret.Data[kubedb.HazelcastKeystorePassKey]),
				TrustStoreType:     "PKCS12",
				KeyStore:           kubedb.HazelcastServerKeystoreFile,
				KeyStorePassword:   string(secret.Data[kubedb.HazelcastKeystorePassKey]),
				KeyStoreType:       "PKCS12",
				// MutualAuthentication: "OPTIONAL", // mtls not configured
				// KeyMaterialDuration:  "PT10M",
			},
		}

		sslClient := SSL{
			Enabled: true,
			Properties: SSLProperties{
				Protocol:           "TLSv1.2",
				TrustStore:         kubedb.HazelcastClientTruststoreFile,
				TrustStorePassword: string(secret.Data[kubedb.HazelcastKeystorePassKey]),
				TrustStoreType:     "PKCS12",
				KeyStore:           kubedb.HazelcastClientKeystoreFile,
				KeyStorePassword:   string(secret.Data[kubedb.HazelcastKeystorePassKey]),
				KeyStoreType:       "PKCS12",
			},
		}

		config.Hazelcast.Network.SSL = ssl
		clientConfig.HazelcastClient.ClientNetwork.SSL = sslClient

		// Configure REST API SSL
		// Note: ClientAuth is set to "NONE" to allow password-based authentication without client certificates
		// Set to "NEED" or "WANT" if mutual TLS authentication is required
		restSSL := SSLREST{
			Enabled:            true,
			ClientAuth:         "WANT",
			EnabledProtocols:   "TLSv1.2,TLSv1.3",
			KeyStore:           kubedb.HazelcastServerKeystoreFile,
			KeyStorePassword:   string(secret.Data[kubedb.HazelcastKeystorePassKey]),
			KeyStoreType:       "PKCS12",
			TrustStore:         kubedb.HazelcastServerTruststoreFile,
			TrustStorePassword: string(secret.Data[kubedb.HazelcastKeystorePassKey]),
			TrustStoreType:     "PKCS12",
			Protocol:           "TLS",
		}
		config.Hazelcast.Rest.SSL = restSSL
	}

	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal server config: %w", err)
	}

	yamlDataClient, err := yaml.Marshal(&clientConfig)
	return addQuotesToCredentials(string(yamlData)), addQuotesToCredentials(string(yamlDataClient)), err
}
