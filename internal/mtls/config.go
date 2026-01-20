package mtls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

// LoadTLSConfig loads TLS configuration from certificate files.
// When a custom CA is provided, it is appended to system root CAs.
func LoadTLSConfig(certFile, keyFile, caFile string) (*tls.Config, error) {
	// Load client certificate and key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate: %w", err)
	}

	// Create TLS config with the certificate
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	// Load CA certificate if provided
	if caFile != "" {
		caCert, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA certificate: %w", err)
		}

		// Start with system root CAs, fall back to empty pool
		caCertPool, sysErr := x509.SystemCertPool()
		if sysErr != nil || caCertPool == nil {
			caCertPool = x509.NewCertPool()
		}

		if !caCertPool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to append CA certificate to pool")
		}

		tlsConfig.RootCAs = caCertPool
	}

	return tlsConfig, nil
}

// LoadTLSConfigFromBytes loads TLS configuration from certificate bytes.
// When a custom CA is provided, it is appended to system root CAs.
func LoadTLSConfigFromBytes(certPEM, keyPEM, caPEM []byte) (*tls.Config, error) {
	// Load client certificate and key
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate: %w", err)
	}

	// Create TLS config with the certificate
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	// Load CA certificate if provided
	if len(caPEM) > 0 {
		// Start with system root CAs, fall back to empty pool
		caCertPool, sysErr := x509.SystemCertPool()
		if sysErr != nil || caCertPool == nil {
			caCertPool = x509.NewCertPool()
		}

		if !caCertPool.AppendCertsFromPEM(caPEM) {
			return nil, fmt.Errorf("failed to append CA certificate to pool")
		}

		tlsConfig.RootCAs = caCertPool
	}

	return tlsConfig, nil
}
