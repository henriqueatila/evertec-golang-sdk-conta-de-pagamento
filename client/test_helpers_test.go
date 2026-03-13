package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Test helper functions for pointer types

func strPtr(s string) *string {
	return &s
}

func intPtr(i int64) *int64 {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}

// testEndpoint is a table-driven test case for API endpoint coverage
type testEndpoint struct {
	name     string
	method   string // expected HTTP method
	path     string // expected URL path
	status   int    // response status code
	response any    // response body to return as JSON
	call     func(ctx context.Context, c *Client) error
}

// runEndpointTests runs table-driven endpoint tests
func runEndpointTests(t *testing.T, tests []testEndpoint) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tt.method {
					t.Errorf("method = %s, want %s", r.Method, tt.method)
				}
				if r.URL.Path != tt.path {
					t.Errorf("path = %s, want %s", r.URL.Path, tt.path)
				}
				w.Header().Set("Content-Type", "application/json")
				status := tt.status
				if status == 0 {
					status = http.StatusOK
				}
				w.WriteHeader(status)
				if tt.response != nil {
					_ = json.NewEncoder(w).Encode(tt.response)
				}
			}))
			defer server.Close()

			c, err := New(server.URL, "test-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
			if err != nil {
				t.Fatalf("failed to create client: %v", err)
			}
			defer c.Close()

			err = tt.call(context.Background(), c)
			if tt.status >= 400 && err == nil {
				t.Errorf("expected error for status %d, got nil", tt.status)
			}
			if (tt.status == 0 || tt.status < 400) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
