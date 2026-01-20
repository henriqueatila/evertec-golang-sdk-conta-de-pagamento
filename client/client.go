package client

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/internal/mtls"
)

// Client is the Evertec API client
type Client struct {
	config *Config
	http   *http.Client
}

// New creates a new Evertec API client with the provided configuration
func New(baseURL, apiKey string, tlsConfig *tls.Config, opts ...Option) (*Client, error) {
	// Initialize config with required parameters
	config := &Config{
		BaseURL:   baseURL,
		APIKey:    apiKey,
		TLSConfig: tlsConfig,
	}

	// Apply functional options
	for _, opt := range opts {
		opt(config)
	}

	// Apply defaults
	config.applyDefaults()

	// Validate configuration
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Create HTTP client with TLS config
	httpClient := &http.Client{
		Timeout: config.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: config.TLSConfig,
		},
	}

	client := &Client{
		config: config,
		http:   httpClient,
	}

	config.Logger.Info("Evertec API client initialized",
		"base_url", config.BaseURL,
		"timeout", config.Timeout,
	)

	return client, nil
}

// NewWithCertFiles creates a new Evertec API client using certificate files for mTLS
func NewWithCertFiles(baseURL, apiKey, certFile, keyFile, caFile string, opts ...Option) (*Client, error) {
	// Load TLS configuration from files
	tlsConfig, err := mtls.LoadTLSConfig(certFile, keyFile, caFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load TLS configuration: %w", err)
	}

	// Create client with the loaded TLS config
	return New(baseURL, apiKey, tlsConfig, opts...)
}

// Config returns a copy of the client configuration
func (c *Client) Config() Config {
	return *c.config
}

// Close closes the HTTP client's idle connections
func (c *Client) Close() {
	c.http.CloseIdleConnections()
	c.config.Logger.Info("Evertec API client closed")
}
