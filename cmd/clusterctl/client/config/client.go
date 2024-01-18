package config

import (
	"context"

	"github.com/pkg/errors"
)

// *************************************************** PUBLIC METHODS **************************************************

// New returns a Client for interacting with the clusterctl configuration.
func New(ctx context.Context, path string, options ...Option) (Client, error) {
	return newConfigClient(ctx, path, options...)
}

// ************************************************** PRIVATE METHODS **************************************************

func newConfigClient(ctx context.Context, path string, options ...Option) (*configClient, error) {
	client := &configClient{}
	for _, o := range options {
		o(client)
	}

	// if there is an injected reader, use it, otherwise use a default one
	var err error
	if client.reader == nil {
		if client.reader, err = newViperReader(); err != nil {
			return nil, errors.Wrap(err, "failed to create the configuration reader")
		}
		if err = client.reader.Init(ctx, path); err != nil {
			return nil, errors.Wrap(err, "failed to initialize the configuration reader")
		}
	}

	return client, nil
}

// Reader define the behaviours of a configuration reader.
type Reader interface {
	// Init allows to initialize the configuration reader.
	Init(ctx context.Context, path string) error

	// Get returns a configuration value of type string.
	// In case the configuration value does not exists, it returns an error.
	Get(key string) (string, error)

	// Set allows to set an explicit override for a config value.
	// e.g. It is used to set an override from a flag value over environment/config file variables.
	Set(key, value string)

	// UnmarshalKey reads a configuration value and unmarshals it into the provided value object.
	UnmarshalKey(key string, value interface{}) error
}

// Option is a configuration option supplied to New.
type Option func(*configClient)

// configClient implements Client.
type configClient struct {
	reader Reader
}

// Client is used to interact with the clusterctl configurations.
// Clusterctl v2 handles the following configurations:
// 1. The cert manager configuration (URL of the repository)
// 2. The configuration of the providers (name, type and URL of the provider repository)
// 3. Variables used when installing providers/creating clusters. Variables can be read from the environment or from the config file
// 4. The configuration about image overrides.
type Client interface {
	// CertManager provide access to the cert-manager configurations.
	CertManager() CertManagerClient

	// Providers provide access to provider configurations.
	Providers() ProvidersClient

	// Variables provide access to environment variables and/or variables defined in the clusterctl configuration file.
	Variables() VariablesClient

	// ImageMeta provide access to image meta configurations.
	ImageMeta() ImageMetaClient
}

func (c *configClient) CertManager() CertManagerClient {
	return newCertManagerClient(c.reader)
}

func (c *configClient) Providers() ProvidersClient {
	return newProvidersClient(c.reader)
}

func (c *configClient) Variables() VariablesClient {
	return newVariablesClient(c.reader)
}

func (c *configClient) ImageMeta() ImageMetaClient {
	return newImageMetaClient(c.reader)
}
