/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package authn

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/awslabs/amazon-ecr-credential-helper/ecr-login"
	"github.com/chrismellard/docker-credential-acr-env/pkg/credhelper"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/authn/github"
	"github.com/google/go-containerregistry/pkg/authn/k8schain"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/spf13/pflag"
	"golang.org/x/net/publicsuffix"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

var (
	insecureRegistries  []string
	once                sync.Once
	insecureRegistrySet sets.Set[string]

	SkipImageDigest string
	amazonKeychain  = authn.NewKeychainFromHelper(ecr.NewECRHelper(ecr.WithLogger(io.Discard)))
	azureKeychain   = authn.NewKeychainFromHelper(credhelper.NewACRCredentialsHelper())
)

// AddIsecureRegsitryFlag is for explicitly initializing the flags
func AddInsecureRegistriesFlag(fs *pflag.FlagSet) {
	if fs == nil {
		fs = pflag.CommandLine
	}
	fs.StringSliceVar(&insecureRegistries, "insecure-registries", insecureRegistries, "List of registries to be used without TLS verification")
}

func ImageWithDigest(kc kubernetes.Interface, image string, k8sOpts *k8schain.Options) (string, error) {
	// Drop the "@sha256:hash_string" part, if any
	image, err := ImageWithoutDigest(image)
	if err != nil {
		return "", err
	}
	if SkipImageDigest == "true" {
		return image, nil
	}

	keyChain, err := CreateKeyChain(context.TODO(), kc, k8sOpts)
	if err != nil {
		return "", err
	}

	digest, err := crane.Digest(image, crane.WithAuthFromKeychain(keyChain), WithTLSSkipVerify(image))
	if err != nil {
		var ce *tls.CertificateVerificationError
		if errors.As(err, &ce) {
			klog.ErrorS(err, "UnverifiedCertificates")
			for _, cert := range ce.UnverifiedCertificates {
				klog.Errorln(string(encodeCertPEM(cert)))
			}
		}
		return "", err
	}

	return image + "@" + digest, nil
}

func encodeCertPEM(cert *x509.Certificate) []byte {
	block := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}
	return pem.EncodeToMemory(&block)
}

// CreateKeyChain a multi keychain based in input arguments
func CreateKeyChain(ctx context.Context, client kubernetes.Interface, k8sOpts *k8schain.Options) (authn.Keychain, error) {
	// xref: https://github.com/google/k8s-digester/blob/v0.1.9/pkg/keychain/keychain.go#L42-L64
	if k8sOpts != nil || isOptionsSet() {
		keychains := make([]authn.Keychain, 0, 3)
		if k8sOpts != nil {
			kChain, err := k8schain.New(ctx, client, *k8sOpts)
			if err != nil {
				return nil, err
			}
			keychains = append(keychains, kChain)
		}
		if isOptionsSet() {
			kChain, err := k8schain.New(ctx, client, optsFromFlags)
			if err != nil {
				return nil, err
			}
			keychains = append(keychains, kChain)
		}
		keychains = append(keychains, authn.DefaultKeychain)
		return authn.NewMultiKeychain(keychains...), nil
	}

	return authn.NewMultiKeychain(
		google.Keychain,
		authn.DefaultKeychain,
		github.Keychain,
		amazonKeychain,
		azureKeychain,
	), nil
}

// CreateInClusterKeyChain a multi keychain based using in cluster client
func CreateInClusterKeyChain(ctx context.Context, k8sOpts *k8schain.Options) (authn.Keychain, error) {
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		return nil, err
	}

	return CreateKeyChain(ctx, client, k8sOpts)
}

// ImageWithoutDigest takes image as input, return image without the digest value
func ImageWithoutDigest(image string) (string, error) {
	if before, _, found := strings.Cut(image, "@"); found {
		if len(before) > 0 {
			return before, nil
		}
		return "", fmt.Errorf("invalid image: %s", image)
	}
	return image, nil
}

func WithTLSSkipVerify(s string) crane.Option {
	// xref: https://github.com/google/go-containerregistry/pull/1054
	rt := remote.DefaultTransport.(*http.Transport).Clone()
	rt.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: probablyInsecureRegistry(s), //nolint: gosec
	}
	return crane.WithTransport(rt)
}

func probablyInsecureRegistry(s string) bool {
	parts := strings.Split(s, "/")
	if len(parts) > 1 && strings.ContainsRune(parts[0], '.') {
		once.Do(func() {
			insecureRegistrySet = sets.New[string](insecureRegistries...)
		})
		if insecureRegistrySet.Has(parts[0]) {
			return true
		}
		if _, icann := publicsuffix.PublicSuffix(parts[0]); !icann {
			return true
		}
	}
	return false
}
