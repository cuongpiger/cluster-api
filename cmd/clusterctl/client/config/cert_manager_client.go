package config

import (
	"github.com/drone/envsubst/v2"
	"github.com/pkg/errors"
	"os"
	"time"
)

// ****************************************************** CONSTS *******************************************************

const (
	// CertManagerConfigKey defines the name of the top level config key for cert-manager configuration.
	CertManagerConfigKey = "cert-manager"

	// CertManagerDefaultVersion defines the default cert-manager version to be used by clusterctl.
	CertManagerDefaultVersion = "v1.13.2"

	// CertManagerDefaultURL defines the default cert-manager repository url to be used by clusterctl.
	// NOTE: At runtime CertManagerDefaultVersion may be replaced with the
	// version defined by the user in the clusterctl configuration file.
	CertManagerDefaultURL = "https://github.com/cert-manager/cert-manager/releases/" + CertManagerDefaultVersion + "/cert-manager.yaml"

	// CertManagerDefaultTimeout defines the default cert-manager timeout to be used by clusterctl.
	CertManagerDefaultTimeout = 10 * time.Minute
)

// CertManagerClient has methods to work with cert-manager configurations.
type CertManagerClient interface {
	// Get returns the cert-manager configuration.
	Get() (CertManager, error)
}

func newCertManagerClient(reader Reader) *certManagerClient {
	return &certManagerClient{
		reader: reader,
	}
}

// certManagerClient implements CertManagerClient.
type certManagerClient struct {
	reader Reader
}

func (p *certManagerClient) Get() (CertManager, error) {
	url := CertManagerDefaultURL
	version := CertManagerDefaultVersion
	timeout := CertManagerDefaultTimeout.String()

	userCertManager := &configCertManager{}
	if err := p.reader.UnmarshalKey(CertManagerConfigKey, &userCertManager); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal certManager from the clusterctl configuration file")
	}
	if userCertManager.URL != "" {
		url = userCertManager.URL
	}

	url, err := envsubst.Eval(url, os.Getenv)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to evaluate url: %q", url)
	}

	if userCertManager.Version != "" {
		version = userCertManager.Version
	}
	if userCertManager.Timeout != "" {
		timeout = userCertManager.Timeout
	}

	return NewCertManager(url, version, timeout), nil
}

// ___________________________________________________________________________________________________ configCertManager

// configCertManager mirrors config.CertManager interface and allows serialization of the corresponding info.
type configCertManager struct {
	URL     string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
	Timeout string `json:"timeout,omitempty"`
}
