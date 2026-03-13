package mtls

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"testing"
	"time"
)

// generateTestCerts creates a self-signed CA, a client cert signed by CA, and returns PEM bytes
func generateTestCerts(t *testing.T) (caPEM, certPEM, keyPEM []byte) {
	t.Helper()

	// Generate CA key
	caKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate CA key: %v", err)
	}

	// Create CA certificate template
	caTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "Test CA"},
		NotBefore:             time.Now().Add(-1 * time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
	}

	caDER, err := x509.CreateCertificate(rand.Reader, caTemplate, caTemplate, &caKey.PublicKey, caKey)
	if err != nil {
		t.Fatalf("failed to create CA cert: %v", err)
	}

	caPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caDER})

	// Parse CA cert for signing
	caCert, err := x509.ParseCertificate(caDER)
	if err != nil {
		t.Fatalf("failed to parse CA cert: %v", err)
	}

	// Generate client key
	clientKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate client key: %v", err)
	}

	// Create client certificate template
	clientTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "Test Client"},
		NotBefore:    time.Now().Add(-1 * time.Hour),
		NotAfter:     time.Now().Add(24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	clientDER, err := x509.CreateCertificate(rand.Reader, clientTemplate, caCert, &clientKey.PublicKey, caKey)
	if err != nil {
		t.Fatalf("failed to create client cert: %v", err)
	}

	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: clientDER})

	keyDER, err := x509.MarshalECPrivateKey(clientKey)
	if err != nil {
		t.Fatalf("failed to marshal client key: %v", err)
	}
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER})

	return caPEM, certPEM, keyPEM
}

// writeTempFile writes bytes to a temp file and returns the path
func writeTempFile(t *testing.T, content []byte, pattern string) string {
	t.Helper()
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	t.Cleanup(func() { os.Remove(f.Name()) })
	if _, err := f.Write(content); err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}
	f.Close()
	return f.Name()
}

// TestLoadTLSConfigFromBytesValid tests loading TLS config from valid PEM bytes
func TestLoadTLSConfigFromBytesValid(t *testing.T) {
	caPEM, certPEM, keyPEM := generateTestCerts(t)

	t.Run("with CA", func(t *testing.T) {
		cfg, err := LoadTLSConfigFromBytes(certPEM, keyPEM, caPEM)
		if err != nil {
			t.Fatalf("LoadTLSConfigFromBytes failed: %v", err)
		}
		if cfg == nil {
			t.Fatal("expected non-nil TLS config")
		}
		if cfg.MinVersion != tls.VersionTLS12 {
			t.Errorf("MinVersion = %d, want %d", cfg.MinVersion, tls.VersionTLS12)
		}
		if len(cfg.Certificates) != 1 {
			t.Errorf("expected 1 certificate, got %d", len(cfg.Certificates))
		}
		if cfg.RootCAs == nil {
			t.Error("expected non-nil RootCAs pool")
		}
	})

	t.Run("without CA", func(t *testing.T) {
		cfg, err := LoadTLSConfigFromBytes(certPEM, keyPEM, nil)
		if err != nil {
			t.Fatalf("LoadTLSConfigFromBytes without CA failed: %v", err)
		}
		if cfg == nil {
			t.Fatal("expected non-nil TLS config")
		}
		if cfg.RootCAs != nil {
			t.Error("expected nil RootCAs when no CA provided")
		}
		if len(cfg.Certificates) != 1 {
			t.Errorf("expected 1 certificate, got %d", len(cfg.Certificates))
		}
	})

	t.Run("with empty CA bytes", func(t *testing.T) {
		cfg, err := LoadTLSConfigFromBytes(certPEM, keyPEM, []byte{})
		if err != nil {
			t.Fatalf("LoadTLSConfigFromBytes with empty CA failed: %v", err)
		}
		if cfg == nil {
			t.Fatal("expected non-nil TLS config")
		}
		if cfg.RootCAs != nil {
			t.Error("expected nil RootCAs when empty CA provided")
		}
	})
}

// TestLoadTLSConfigFromBytesErrors tests error cases for LoadTLSConfigFromBytes
func TestLoadTLSConfigFromBytesErrors(t *testing.T) {
	_, certPEM, keyPEM := generateTestCerts(t)

	t.Run("invalid cert bytes", func(t *testing.T) {
		_, err := LoadTLSConfigFromBytes([]byte("invalid cert"), []byte("invalid key"), nil)
		if err == nil {
			t.Error("expected error for invalid cert bytes")
		}
	})

	t.Run("invalid key bytes", func(t *testing.T) {
		_, err := LoadTLSConfigFromBytes(certPEM, []byte("invalid key"), nil)
		if err == nil {
			t.Error("expected error for invalid key bytes")
		}
	})

	t.Run("invalid CA bytes", func(t *testing.T) {
		_, err := LoadTLSConfigFromBytes(certPEM, keyPEM, []byte("not a valid pem cert"))
		if err == nil {
			t.Error("expected error for invalid CA PEM bytes")
		}
	})

	t.Run("mismatched cert and key", func(t *testing.T) {
		_, certPEM2, _ := generateTestCerts(t)
		_, _, keyPEM2 := generateTestCerts(t)
		_, err := LoadTLSConfigFromBytes(certPEM2, keyPEM2, nil)
		if err == nil {
			t.Error("expected error for mismatched cert and key")
		}
	})
}

// TestLoadTLSConfigValid tests loading TLS config from valid cert files
func TestLoadTLSConfigValid(t *testing.T) {
	caPEM, certPEM, keyPEM := generateTestCerts(t)

	certFile := writeTempFile(t, certPEM, "cert*.pem")
	keyFile := writeTempFile(t, keyPEM, "key*.pem")
	caFile := writeTempFile(t, caPEM, "ca*.pem")

	t.Run("with CA file", func(t *testing.T) {
		cfg, err := LoadTLSConfig(certFile, keyFile, caFile)
		if err != nil {
			t.Fatalf("LoadTLSConfig failed: %v", err)
		}
		if cfg == nil {
			t.Fatal("expected non-nil TLS config")
		}
		if cfg.MinVersion != tls.VersionTLS12 {
			t.Errorf("MinVersion = %d, want %d", cfg.MinVersion, tls.VersionTLS12)
		}
		if len(cfg.Certificates) != 1 {
			t.Errorf("expected 1 certificate, got %d", len(cfg.Certificates))
		}
		if cfg.RootCAs == nil {
			t.Error("expected non-nil RootCAs pool")
		}
	})

	t.Run("without CA file", func(t *testing.T) {
		cfg, err := LoadTLSConfig(certFile, keyFile, "")
		if err != nil {
			t.Fatalf("LoadTLSConfig without CA failed: %v", err)
		}
		if cfg == nil {
			t.Fatal("expected non-nil TLS config")
		}
		if cfg.RootCAs != nil {
			t.Error("expected nil RootCAs when no CA file provided")
		}
	})
}

// TestLoadTLSConfigFileErrors tests error cases for LoadTLSConfig
func TestLoadTLSConfigFileErrors(t *testing.T) {
	_, certPEM, keyPEM := generateTestCerts(t)
	certFile := writeTempFile(t, certPEM, "cert*.pem")
	keyFile := writeTempFile(t, keyPEM, "key*.pem")

	t.Run("nonexistent cert file", func(t *testing.T) {
		_, err := LoadTLSConfig("/nonexistent/cert.pem", "/nonexistent/key.pem", "")
		if err == nil {
			t.Error("expected error for nonexistent cert file")
		}
	})

	t.Run("nonexistent key file", func(t *testing.T) {
		_, err := LoadTLSConfig(certFile, "/nonexistent/key.pem", "")
		if err == nil {
			t.Error("expected error for nonexistent key file")
		}
	})

	t.Run("nonexistent CA file", func(t *testing.T) {
		_, err := LoadTLSConfig(certFile, keyFile, "/nonexistent/ca.pem")
		if err == nil {
			t.Error("expected error for nonexistent CA file")
		}
	})

	t.Run("invalid CA content", func(t *testing.T) {
		invalidCAFile := writeTempFile(t, []byte("not a pem cert"), "badca*.pem")
		_, err := LoadTLSConfig(certFile, keyFile, invalidCAFile)
		if err == nil {
			t.Error("expected error for invalid CA PEM content")
		}
	})
}
