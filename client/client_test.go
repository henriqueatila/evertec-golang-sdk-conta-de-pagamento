package client

import (
	"crypto/tls"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	tests := []struct {
		name      string
		baseURL   string
		apiKey    string
		tlsConfig *tls.Config
		opts      []Option
		wantErr   bool
	}{
		{
			name:      "valid configuration",
			baseURL:   "https://api.example.com",
			apiKey:    "test-key",
			tlsConfig: tlsConfig,
			opts:      nil,
			wantErr:   false,
		},
		{
			name:      "missing base URL",
			baseURL:   "",
			apiKey:    "test-key",
			tlsConfig: tlsConfig,
			opts:      nil,
			wantErr:   true,
		},
		{
			name:      "missing API key",
			baseURL:   "https://api.example.com",
			apiKey:    "",
			tlsConfig: tlsConfig,
			opts:      nil,
			wantErr:   true,
		},
		{
			name:      "missing TLS config",
			baseURL:   "https://api.example.com",
			apiKey:    "test-key",
			tlsConfig: nil,
			opts:      nil,
			wantErr:   true,
		},
		{
			name:      "with custom timeout",
			baseURL:   "https://api.example.com",
			apiKey:    "test-key",
			tlsConfig: tlsConfig,
			opts:      []Option{WithTimeout(60 * time.Second)},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := New(tt.baseURL, tt.apiKey, tt.tlsConfig, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && client == nil {
				t.Error("New() returned nil client without error")
			}

			if client != nil {
				client.Close()
			}
		})
	}
}

func TestConfigValidation(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				BaseURL: "https://api.example.com",
				APIKey:  "test-key",
				TLSConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
			wantErr: false,
		},
		{
			name: "missing base URL",
			config: &Config{
				BaseURL: "",
				APIKey:  "test-key",
				TLSConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
			wantErr: true,
		},
		{
			name: "TLS version too low",
			config: &Config{
				BaseURL: "https://api.example.com",
				APIKey:  "test-key",
				TLSConfig: &tls.Config{
					MinVersion: tls.VersionTLS10,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigApplyDefaults(t *testing.T) {
	config := &Config{
		BaseURL: "https://api.example.com",
		APIKey:  "test-key",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	config.applyDefaults()

	if config.Timeout != DefaultTimeout {
		t.Errorf("expected timeout %v, got %v", DefaultTimeout, config.Timeout)
	}

	if config.Logger == nil {
		t.Error("expected Logger to be set")
	}

	if config.UserAgent != DefaultUserAgent {
		t.Errorf("expected UserAgent %s, got %s", DefaultUserAgent, config.UserAgent)
	}

	if config.Hooks == nil {
		t.Error("expected Hooks to be initialized")
	}
}

func TestOptions(t *testing.T) {
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	t.Run("WithProduction", func(t *testing.T) {
		client, err := New("https://api.example.com", "test-key", tlsConfig, WithProduction())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer client.Close()

		if client.config.BaseURL != ProductionBaseURL {
			t.Errorf("expected BaseURL %s, got %s", ProductionBaseURL, client.config.BaseURL)
		}
	})

	t.Run("WithHomolog", func(t *testing.T) {
		client, err := New("https://api.example.com", "test-key", tlsConfig, WithHomolog())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer client.Close()

		if client.config.BaseURL != HomologBaseURL {
			t.Errorf("expected BaseURL %s, got %s", HomologBaseURL, client.config.BaseURL)
		}
	})

	t.Run("WithTimeout", func(t *testing.T) {
		customTimeout := 60 * time.Second
		client, err := New("https://api.example.com", "test-key", tlsConfig, WithTimeout(customTimeout))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer client.Close()

		if client.config.Timeout != customTimeout {
			t.Errorf("expected Timeout %v, got %v", customTimeout, client.config.Timeout)
		}
	})

	t.Run("WithUserAgent", func(t *testing.T) {
		customAgent := "CustomApp/1.0"
		client, err := New("https://api.example.com", "test-key", tlsConfig, WithUserAgent(customAgent))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer client.Close()

		if client.config.UserAgent != customAgent {
			t.Errorf("expected UserAgent %s, got %s", customAgent, client.config.UserAgent)
		}
	})

	t.Run("WithHooks", func(t *testing.T) {
		hook := &NoOpHook{}
		client, err := New("https://api.example.com", "test-key", tlsConfig, WithHooks(hook))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer client.Close()

		if len(client.config.Hooks) != 1 {
			t.Errorf("expected 1 hook, got %d", len(client.config.Hooks))
		}
	})
}
