package mtls

import (
	"testing"
)

func TestLoadTLSConfig(t *testing.T) {
	// Skip TLS loading tests as they require valid certificates
	// In production, users will provide their own valid certificates from Evertec
	t.Skip("Skipping TLS tests - require valid certificates from Evertec")

}

func TestLoadTLSConfigFromBytes(t *testing.T) {
	// Skip TLS loading tests as they require valid certificates
	// In production, users will provide their own valid certificates from Evertec
	t.Skip("Skipping TLS tests - require valid certificates from Evertec")
}

func TestLoadTLSConfigErrors(t *testing.T) {
	t.Run("invalid cert file", func(t *testing.T) {
		_, err := LoadTLSConfig("nonexistent-cert.pem", "nonexistent-key.pem", "")
		if err == nil {
			t.Error("expected error for nonexistent cert file")
		}
	})

	t.Run("invalid cert bytes", func(t *testing.T) {
		_, err := LoadTLSConfigFromBytes([]byte("invalid"), []byte("invalid"), nil)
		if err == nil {
			t.Error("expected error for invalid cert bytes")
		}
	})
}
