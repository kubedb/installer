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

package hazelcast

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	apiutils "kubedb.dev/apimachinery/pkg/utils"

	"github.com/go-logr/logr"
	"github.com/go-resty/resty/v2"
	hazelcast "github.com/hazelcast/hazelcast-go-client"
	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeDBClientBuilder struct {
	kc      client.Client
	db      *api.Hazelcast
	log     logr.Logger
	url     string
	podName string
	ctx     context.Context
}

func NewKubeDBClientBuilder(kc client.Client, db *api.Hazelcast) *KubeDBClientBuilder {
	return &KubeDBClientBuilder{
		kc: kc,
		db: db,
	}
}

func (o *KubeDBClientBuilder) WithPod(podName string) *KubeDBClientBuilder {
	o.podName = podName
	return o
}

func (o *KubeDBClientBuilder) WithLog(log logr.Logger) *KubeDBClientBuilder {
	o.log = log
	return o
}

func (o *KubeDBClientBuilder) WithURL(url string) *KubeDBClientBuilder {
	o.url = url
	return o
}

func (o *KubeDBClientBuilder) WithContext(ctx context.Context) *KubeDBClientBuilder {
	o.ctx = ctx
	return o
}

func (o *KubeDBClientBuilder) GetAuthCredentials() (string, string, error) {
	var authSecret core.Secret
	var username, password string
	err := o.kc.Get(o.ctx, client.ObjectKey{Namespace: o.db.Namespace, Name: o.db.GetAuthSecretName()}, &authSecret)
	if err != nil {
		return "", "", errors.Errorf("Failed to get auth secret with %s", err)
	}

	if value, ok := authSecret.Data[core.BasicAuthUsernameKey]; ok {
		username = string(value)
	} else {
		klog.Errorf("Failed for secret: %s/%s, username is missing", authSecret.Namespace, authSecret.Name)
		return "", "", errors.New("username is missing")
	}

	if value, ok := authSecret.Data[core.BasicAuthPasswordKey]; ok {
		password = string(value)
	} else {
		klog.Errorf("Failed for secret: %s/%s, password is missing", authSecret.Namespace, authSecret.Name)
		return "", "", errors.New("password is missing")
	}

	return username, password, nil
}

func (o *KubeDBClientBuilder) GetTLSConfig() (*tls.Config, error) {
	var certSecret core.Secret
	err := o.kc.Get(o.ctx, types.NamespacedName{
		Namespace: o.db.Namespace,
		Name:      o.db.GetCertSecretName(api.HazelcastClientCert),
	}, &certSecret)
	if err != nil {
		klog.Error(err, "failed to get clientCert secret")
		return nil, err
	}

	// get tls cert, clientCA and rootCA for tls config
	// use server cert ca for rootca as issuer ref is not taken into account
	clientCA := x509.NewCertPool()
	rootCA := x509.NewCertPool()

	crt, err := tls.X509KeyPair(certSecret.Data[core.TLSCertKey], certSecret.Data[core.TLSPrivateKeyKey])
	if err != nil {
		klog.Error(err, "failed to create certificate for TLS config")
		return nil, err
	}
	clientCA.AppendCertsFromPEM(certSecret.Data[kubedb.CACert])
	rootCA.AppendCertsFromPEM(certSecret.Data[kubedb.CACert])

	tlsConfig := &tls.Config{
		ServerName:   o.db.ServiceName(),
		Certificates: []tls.Certificate{crt},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCA,
		RootCAs:      rootCA,
		MaxVersion:   tls.VersionTLS13,
	}
	return tlsConfig, nil
}

func (o *KubeDBClientBuilder) GetHazelcastClient() (*Client, error) {
	if o.podName != "" {
		o.url = fmt.Sprintf("%s.%s.%s.svc.%s:%d", o.podName, o.db.GoverningServiceName(), o.db.GetNamespace(), apiutils.FindDomain(), kubedb.HazelcastRestPort)
	}

	if o.url == "" {
		o.url = o.ServiceURL()
	}

	if o.db == nil {
		return nil, errors.New("db is empty")
	}

	config := hazelcast.Config{}
	config.Cluster.Name = fmt.Sprintf("%s/%s", o.db.Namespace, o.db.Name)
	config.Cluster.Network.SetAddresses(o.url)

	if !o.db.Spec.DisableSecurity {
		username, password, err := o.GetAuthCredentials()
		if err != nil {
			return nil, err
		}
		config.Cluster.Security.Credentials.Username = username
		config.Cluster.Security.Credentials.Password = password
	}

	// If EnableSSL is true set tls config,
	// provide client certs and root CA
	if o.db.Spec.EnableSSL {
		tlsConfig, err := o.GetTLSConfig()
		if err != nil {
			return nil, err
		}
		config.Cluster.Network.SSL.Enabled = true
		config.Cluster.Network.SSL.ServerName = fmt.Sprintf("%s.%s.svc", o.db.Name, o.db.Namespace)
		config.Cluster.Network.SSL.SetTLSConfig(tlsConfig)
	}

	hzClient, err := hazelcast.StartNewClientWithConfig(o.ctx, config)
	if err != nil {
		klog.Errorf("Failed to create HTTP client for Hazelcast: %s/%s with: %s", o.db.Namespace, o.db.Name, err)
		return nil, err
	}

	isRunning := hzClient.Running()
	if !isRunning {
		return nil, errors.New("Hazelcast client is not running")
	}

	return &Client{
		Client: hzClient,
	}, nil
}

type RestyConfig struct {
	host      string
	transport *http.Transport
}

func (o *KubeDBClientBuilder) GetHazelcastRestyClient() (*HZRestyClient, error) {
	if o.url == "" {
		o.url = fmt.Sprintf("%s://%s.%s.svc:%d", o.db.GetConnectionScheme(), o.db.ServiceName(), o.db.GetNamespace(), kubedb.HazelcastRestPort)
	}

	config := RestyConfig{
		host: o.url,
		transport: &http.Transport{
			IdleConnTimeout: time.Second * 3,
			DialContext: (&net.Dialer{
				Timeout: time.Second * 30,
			}).DialContext,
		},
	}

	var username, password string
	if !o.db.Spec.DisableSecurity {
		user, pass, err := o.GetAuthCredentials()
		if err != nil {
			return nil, err
		}
		username = user
		password = pass
	}

	defaultTlsConfig, err := o.GetTLSConfig()
	if err != nil {
		klog.Errorf("Failed to get default tls config: %v", err)
	}
	config.transport.TLSClientConfig = defaultTlsConfig
	newClient := resty.New()
	newClient.SetTransport(config.transport).SetScheme(o.db.GetConnectionScheme()).SetBaseURL(config.host)
	newClient.SetHeader("Accept", "application/json")
	newClient.SetBasicAuth(username, password)
	newClient.SetTimeout(time.Second * 30)

	return &HZRestyClient{
		Client:   newClient,
		Config:   &config,
		password: password,
	}, nil
}

func (client *HZRestyClient) ChangeClusterState(state string) (string, error) {
	req := client.Client.R().SetDoNotParseResponse(true)
	param := fmt.Sprintf("admin&%s&%s", client.password, state)
	req.SetHeader("Content-Type", "application/json")
	req.SetBody(param)
	res, err := req.Post("/hazelcast/rest/management/cluster/changeState")
	if err != nil {
		klog.Error(err, "Failed to send http request")
		return "", err
	}
	if res != nil {
		if res.IsError() {
			klog.Error(res.Error())
			return "", errors.New(fmt.Sprintf("HTTP request failed: %v, StatusCode: %v", res.Error(), res.StatusCode()))
		}
	} else {
		return "", errors.New("response can not be nil")
	}
	body := res.RawBody()
	responseBody := make(map[string]any)
	if err := json.NewDecoder(body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to deserialize the response: %v", err)
	}
	if val, ok := responseBody["status"]; ok {
		if strValue, ok := val.(string); ok {
			return strValue, nil
		}
		return "", errors.New("failed to convert response to string")
	}
	return "", errors.New("status is missing")
}

func (client *HZRestyClient) GetClusterState() (string, error) {
	req := client.Client.R().SetDoNotParseResponse(true)

	res, err := req.Get("/hazelcast/health")
	if err != nil {
		klog.Error(err, "Failed to send http request")
		return "", err
	}
	if res != nil {
		if res.IsError() {
			klog.Error(res.Error())
			return "", errors.New(fmt.Sprintf("HTTP request failed: %v, StatusCode: %v", res.Error(), res.StatusCode()))
		}
	} else {
		return "", errors.New("response can not be nil")
	}
	body := res.RawBody()
	responseBody := make(map[string]any)
	if err := json.NewDecoder(body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to deserialize the response: %v", err)
	}
	if val, ok := responseBody["clusterState"]; ok {
		if strValue, ok := val.(string); ok {
			return strValue, nil
		}
		return "", errors.New("failed to convert response to string")
	}
	return "", errors.New("status is missing")
}

func (o *KubeDBClientBuilder) ServiceURL() string {
	return fmt.Sprintf("%s.%s.svc:%d", o.db.ServiceName(), o.db.GetNamespace(), kubedb.HazelcastRestPort)
}
