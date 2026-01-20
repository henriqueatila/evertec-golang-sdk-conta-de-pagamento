package client

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	// ProductionBaseURL is the base URL for the Evertec production environment
	// v2 endpoint per official documentation
	ProductionBaseURL = "https://api-v2.conta-digital.paysmart.com.br/conta-digital/api/v1"

	// HomologBaseURL is the base URL for the Evertec homologation API
	HomologBaseURL = "https://api-v2.homolog.conta-digital.paysmart.com.br/conta-digital/api/v1"

	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout = 30 * time.Second

	// DefaultUserAgent is the default User-Agent header
	DefaultUserAgent = "evertec-golang-sdk-conta-de-pagamento/1.0"

	// APIKeyHeader is the header name for the API key
	APIKeyHeader = "X-API-KEY"
)

// Config holds the configuration for the Evertec API client
type Config struct {
	// BaseURL is the base URL for the API (defaults to ProductionBaseURL)
	BaseURL string

	// APIKey is the API key for authentication
	APIKey string

	// Timeout is the HTTP client timeout (defaults to DefaultTimeout)
	Timeout time.Duration

	// Logger is the structured logger (defaults to slog.Default())
	Logger *slog.Logger

	// TLSConfig is the TLS configuration for mTLS authentication
	TLSConfig *tls.Config

	// ServerName overrides the TLS ServerName for SNI (useful for IP-based endpoints)
	ServerName string

	// UserAgent is the User-Agent header value
	UserAgent string

	// Hooks are called before/after each request for observability
	Hooks []Hook

	// TracerProvider is the OpenTelemetry tracer provider for distributed tracing
	TracerProvider trace.TracerProvider

	// MeterProvider is the OpenTelemetry meter provider for metrics
	MeterProvider metric.MeterProvider

	// TracingEnabled enables OpenTelemetry tracing (uses default provider if TracerProvider is nil)
	TracingEnabled bool

	// MetricsEnabled enables OpenTelemetry metrics (uses default provider if MeterProvider is nil)
	MetricsEnabled bool

	// AutoIdempotency enables automatic UUID-v4 idempotency key generation for mutating requests
	AutoIdempotency bool
}

// validate checks if the configuration is valid
func (c *Config) validate() error {
	if c.BaseURL == "" {
		return fmt.Errorf("base URL is required")
	}

	if c.APIKey == "" {
		return fmt.Errorf("API key is required")
	}

	if c.TLSConfig == nil {
		return fmt.Errorf("TLS config is required for mTLS authentication")
	}

	// Enforce TLS 1.2 minimum
	if c.TLSConfig.MinVersion == 0 {
		c.TLSConfig.MinVersion = tls.VersionTLS12
	} else if c.TLSConfig.MinVersion < tls.VersionTLS12 {
		return fmt.Errorf("TLS version must be 1.2 or higher")
	}

	// Apply ServerName override if provided
	if c.ServerName != "" {
		c.TLSConfig.ServerName = c.ServerName
	}

	return nil
}

// applyDefaults applies default values to the configuration
func (c *Config) applyDefaults() {
	if c.Timeout == 0 {
		c.Timeout = DefaultTimeout
	}

	if c.Logger == nil {
		c.Logger = slog.Default()
	}

	if c.UserAgent == "" {
		c.UserAgent = DefaultUserAgent
	}

	if c.Hooks == nil {
		c.Hooks = []Hook{}
	}
}
