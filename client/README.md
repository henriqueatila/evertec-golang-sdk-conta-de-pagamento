# Client Package

The `client` package provides the core SDK foundation for the Evertec Conta de Pagamento API.

## Features

- **Mutual TLS (mTLS) Authentication**: Secure authentication using client certificates
- **X-API-KEY Header**: API key authentication on every request
- **Functional Options Pattern**: Flexible configuration via options
- **Observability Hooks**: Request/response lifecycle hooks for logging and metrics
- **Structured Logging**: Built-in support for `slog`
- **Error Handling**: Comprehensive error types matching API spec
- **Context Support**: All requests support context cancellation
- **TLS 1.2+ Enforcement**: Minimum TLS version requirement

## Quick Start

### Basic Client Setup

```go
package main

import (
	"log"
	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/client"
)

func main() {
	// Create client using certificate files
	c, err := client.NewWithCertFiles(
		"https://api.homolog.paysmart.com.br",
		"your-api-key",
		"client.crt",
		"client.key",
		"",
		client.WithHomolog(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Use the client...
}
```

### With TLS Config

```go
import (
	"crypto/tls"
	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/client"
	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/internal/mtls"
)

// Load TLS configuration
tlsConfig, err := mtls.LoadTLSConfig("client.crt", "client.key", "ca.crt")
if err != nil {
	log.Fatal(err)
}

// Create client
c, err := client.New(
	"https://api.paysmart.com.br",
	"your-api-key",
	tlsConfig,
	client.WithProduction(),
)
if err != nil {
	log.Fatal(err)
}
defer c.Close()
```

## Configuration Options

### Environment Options

```go
// Production environment
client.WithProduction()  // https://api.paysmart.com.br

// Homologation environment
client.WithHomolog()  // https://api.homolog.paysmart.com.br

// Custom base URL
client.WithBaseURL("https://custom.api.com")
```

### Network Options

```go
// Custom timeout
client.WithTimeout(60 * time.Second)

// Custom User-Agent
client.WithUserAgent("MyApp/1.0")
```

### Observability Options

```go
// Custom logger
client.WithLogger(myLogger)

// Default logger
client.WithDefaultLogger()

// Add hooks
client.WithHooks(hook1, hook2)
```

### TLS Options

```go
// Custom TLS configuration
client.WithTLSConfig(tlsConfig)
```

## Hooks

Hooks provide observability into HTTP client operations. Implement the `Hook` interface:

```go
type Hook interface {
	BeforeRequest(ctx context.Context, method, path string, body any)
	AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error)
}
```

### Example: Logging Hook

```go
type LoggingHook struct{}

func (h *LoggingHook) BeforeRequest(ctx context.Context, method, path string, body any) {
	log.Printf("→ %s %s", method, path)
}

func (h *LoggingHook) AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error) {
	log.Printf("← %s %s [%d] %v", method, path, statusCode, duration)
}

// Use the hook
c, err := client.New(baseURL, apiKey, tlsConfig,
	client.WithHooks(&LoggingHook{}),
)
```

### Example: Metrics Hook

```go
type MetricsHook struct {
	metrics *prometheus.Registry
}

func (h *MetricsHook) BeforeRequest(ctx context.Context, method, path string, body any) {
	// Record request count
}

func (h *MetricsHook) AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error) {
	// Record latency histogram
	// Record error rate
}
```

## Error Handling

The client defines specific error types based on HTTP status codes:

### Error Types

- **ValidationError (400)**: Field-level validation errors
  ```go
  type ValidationError struct {
      StatusCode int
      Errors     []ValidationDetail
  }

  type ValidationDetail struct {
      Code    string
      Field   string
      Message string
  }
  ```

- **BusinessRuleError (409)**: Business rule violations
  ```go
  type BusinessRuleError struct {
      StatusCode int
      Code       string
      Message    string
  }
  ```

- **ExceptionError (500)**: Internal server errors
  ```go
  type ExceptionError struct {
      StatusCode int
      Message    string
  }
  ```

- **IntegrationError (503)**: Service unavailable
  ```go
  type IntegrationError struct {
      StatusCode int
      Message    string
  }
  ```

- **APIError**: Generic API errors
  ```go
  type APIError struct {
      StatusCode int
      Code       string
      Message    string
      Details    interface{}
  }
  ```

### Error Handling Example

```go
import "errors"

account, err := c.GetAccount(ctx, accountID)
if err != nil {
	var validErr *client.ValidationError
	if errors.As(err, &validErr) {
		for _, detail := range validErr.Errors {
			log.Printf("Field %s: %s", detail.Field, detail.Message)
		}
		return
	}

	var businessErr *client.BusinessRuleError
	if errors.As(err, &businessErr) {
		log.Printf("Business rule: %s (code: %s)", businessErr.Message, businessErr.Code)
		return
	}

	log.Fatal(err)
}
```

## HTTP Methods

The client provides internal HTTP helpers (not exported):

- `do(ctx, method, path, body, response)` - Generic HTTP request
- `get(ctx, path, response)` - GET request
- `post(ctx, path, body, response)` - POST request
- `put(ctx, path, body, response)` - PUT request
- `patch(ctx, path, body, response)` - PATCH request
- `delete(ctx, path, response)` - DELETE request

These methods:
- Set required headers (`Content-Type`, `Accept`, `User-Agent`, `X-API-KEY`)
- Execute hooks before/after requests
- Parse error responses
- Support context cancellation
- Log requests at debug level

## Authentication

All requests include:

1. **X-API-KEY Header**: API key provided during client initialization
2. **mTLS**: Client certificate configured in TLS config

Example request headers:
```
Content-Type: application/json
Accept: application/json
User-Agent: evertec-golang-sdk-conta-de-pagamento/1.0
X-API-KEY: your-api-key-here
```

## Context Support

All client methods accept a `context.Context`:

```go
import "context"

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

account, err := c.GetAccount(ctx, accountID)

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
	// Cancel on signal
	<-sigChan
	cancel()
}()

balance, err := c.GetBalance(ctx, accountID)
```

## Thread Safety

The `Client` is safe for concurrent use by multiple goroutines. The underlying `http.Client` manages connection pooling automatically.

## Best Practices

1. **Reuse Clients**: Create one client and reuse it across requests
2. **Close on Exit**: Call `Close()` to clean up idle connections
3. **Use Contexts**: Always pass contexts for timeout/cancellation
4. **Handle Errors**: Check for specific error types
5. **Add Hooks**: Use hooks for observability
6. **Configure TLS**: Ensure TLS 1.2+ is enforced
7. **Secure Credentials**: Never hardcode API keys or certificates

## Testing

The package includes comprehensive tests:

```bash
# Run all tests
go test ./client/...

# With coverage
go test ./client/... -cover

# With verbose output
go test ./client/... -v
```

## Package Structure

```
client/
├── client.go       # Main client implementation
├── config.go       # Configuration struct
├── options.go      # Functional options
├── hooks.go        # Hook interface
├── http.go         # HTTP helpers
├── errors.go       # Error types
├── client_test.go  # Client tests
├── errors_test.go  # Error tests
├── example_test.go # Usage examples
└── README.md       # This file
```

## See Also

- [internal/mtls](../internal/mtls/README.md) - TLS configuration utilities
- [Root README](../README.md) - SDK overview
- [API Documentation](https://docs.paysmart.com.br) - Evertec API reference
