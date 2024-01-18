package config

// CertManager defines cert-manager configuration.
type CertManager interface {
	// URL returns the name of the cert-manager repository.
	// If empty, "https://github.com/cert-manager/cert-manager/releases/{DefaultVersion}/cert-manager.yaml" will be used.
	URL() string

	// Version returns the cert-manager version to install.
	// If empty, a default version will be used.
	Version() string

	// Timeout returns the timeout for cert-manager to start.
	// If empty, 10m will be used.
	Timeout() string
}

// ****************************************************** OBJECTS ******************************************************

// certManager implements CertManager.
type certManager struct {
	url     string
	version string
	timeout string
}

func (p *certManager) URL() string {
	return p.url
}

func (p *certManager) Version() string {
	return p.version
}

func (p *certManager) Timeout() string {
	return p.timeout
}

// ************************************************** PUBLIC METHODS ***************************************************

// NewCertManager creates a new CertManager with the given configuration.
func NewCertManager(url, version, timeout string) CertManager {
	return &certManager{
		url:     url,
		version: version,
		timeout: timeout,
	}
}
