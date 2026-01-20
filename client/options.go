package client

import (
	"crypto/tls"
	"log/slog"
	"time"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// Option is a functional option for configuring the Client
type Option func(*Config)

// WithBaseURL sets a custom base URL for the API
func WithBaseURL(url string) Option {
	return func(c *Config) {
		c.BaseURL = url
	}
}

// WithProduction sets the base URL to the production environment
func WithProduction() Option {
	return func(c *Config) {
		c.BaseURL = ProductionBaseURL
	}
}

// WithHomolog sets the base URL to the homologation environment
func WithHomolog() Option {
	return func(c *Config) {
		c.BaseURL = HomologBaseURL
	}
}

// WithTimeout sets the HTTP client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithLogger sets a custom structured logger
func WithLogger(logger *slog.Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

// WithDefaultLogger creates and sets a default logger
func WithDefaultLogger() Option {
	return func(c *Config) {
		c.Logger = slog.Default()
	}
}

// WithTLSConfig sets the TLS configuration for mTLS
func WithTLSConfig(tlsConfig *tls.Config) Option {
	return func(c *Config) {
		c.TLSConfig = tlsConfig
	}
}

// WithUserAgent sets a custom User-Agent header
func WithUserAgent(userAgent string) Option {
	return func(c *Config) {
		c.UserAgent = userAgent
	}
}

// WithHooks adds observability hooks for request/response lifecycle
func WithHooks(hooks ...Hook) Option {
	return func(c *Config) {
		c.Hooks = append(c.Hooks, hooks...)
	}
}

// WithServerName sets a custom TLS ServerName for SNI.
// Use this when connecting to IP-based endpoints or when the server
// certificate's CN/SAN does not match the hostname in the base URL.
func WithServerName(serverName string) Option {
	return func(c *Config) {
		c.ServerName = serverName
	}
}

// WithTracerProvider sets a custom OpenTelemetry TracerProvider for distributed tracing
func WithTracerProvider(tp trace.TracerProvider) Option {
	return func(c *Config) {
		c.TracerProvider = tp
	}
}

// WithMeterProvider sets a custom OpenTelemetry MeterProvider for metrics
func WithMeterProvider(mp metric.MeterProvider) Option {
	return func(c *Config) {
		c.MeterProvider = mp
	}
}

// WithTracing enables OpenTelemetry tracing with default provider
func WithTracing() Option {
	return func(c *Config) {
		c.TracingEnabled = true
	}
}

// WithMetrics enables OpenTelemetry metrics with default provider
func WithMetrics() Option {
	return func(c *Config) {
		c.MetricsEnabled = true
	}
}

// WithAutoIdempotency enables automatic UUID-v4 idempotency key generation for POST/PUT/PATCH/DELETE requests
func WithAutoIdempotency() Option {
	return func(c *Config) {
		c.AutoIdempotency = true
	}
}
