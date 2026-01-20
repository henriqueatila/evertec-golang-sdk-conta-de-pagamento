package client_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/client"
)

// ExampleNew demonstrates creating a client with TLS configuration
func ExampleNew() {
	// Create TLS configuration with client certificates
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		// In production, load certificates using mtls.LoadTLSConfig()
	}

	// Create client with production environment
	c, err := client.New(
		"https://api.paysmart.com.br",
		"your-api-key",
		tlsConfig,
		client.WithProduction(),
		client.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	fmt.Println("Client created successfully")
	// Output: Client created successfully
}

// ExampleNewWithCertFiles demonstrates creating a client using certificate files
func ExampleNewWithCertFiles() {
	// Create client using certificate files
	c, err := client.NewWithCertFiles(
		"https://api.homolog.paysmart.com.br",
		"your-api-key",
		"path/to/client.crt",
		"path/to/client.key",
		"", // Optional CA certificate
		client.WithHomolog(),
	)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return
	}
	defer c.Close()

	fmt.Println("Client with cert files would be created")
	// Note: This example will fail without valid cert files
}

// LoggingHook is a custom hook for logging HTTP requests/responses
type LoggingHook struct{}

func (h *LoggingHook) BeforeRequest(ctx context.Context, method, path string, body any) {
	log.Printf("→ %s %s", method, path)
}

func (h *LoggingHook) AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error) {
	if err != nil {
		log.Printf("← %s %s [%d] %v (error: %v)", method, path, statusCode, duration, err)
	} else {
		log.Printf("← %s %s [%d] %v", method, path, statusCode, duration)
	}
}

// ExampleWithHooks demonstrates using hooks for observability
func ExampleWithHooks() {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	// Create client with logging hook
	c, err := client.New(
		"https://api.paysmart.com.br",
		"your-api-key",
		tlsConfig,
		client.WithHooks(&LoggingHook{}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	fmt.Println("Client with hooks created")
	// Output: Client with hooks created
}

// ExampleWithLogger demonstrates using a custom logger
func ExampleWithLogger() {
	// Create a custom structured logger
	logger := slog.New(slog.NewJSONHandler(log.Writer(), &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	// Create client with custom logger
	c, err := client.New(
		"https://api.paysmart.com.br",
		"your-api-key",
		tlsConfig,
		client.WithLogger(logger),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	fmt.Println("Client with custom logger created")
	// Output: Client with custom logger created
}

// ExampleWithProduction demonstrates using production environment
func ExampleWithProduction() {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	c, err := client.New(
		"https://api.example.com", // This will be overridden
		"your-api-key",
		tlsConfig,
		client.WithProduction(), // Sets base URL to production
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	config := c.Config()
	fmt.Println(config.BaseURL)
	// Output: https://api-v2.conta-digital.paysmart.com.br/conta-digital/api/v1
}

// ExampleWithHomolog demonstrates using homologation environment
func ExampleWithHomolog() {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	c, err := client.New(
		"https://api.example.com", // This will be overridden
		"your-api-key",
		tlsConfig,
		client.WithHomolog(), // Sets base URL to homologation
	)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	config := c.Config()
	fmt.Println(config.BaseURL)
	// Output: https://api-v2.homolog.conta-digital.paysmart.com.br/conta-digital/api/v1
}
