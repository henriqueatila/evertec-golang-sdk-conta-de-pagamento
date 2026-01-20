package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestJoinURL tests the URL joining logic
func TestJoinURL(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		path     string
		expected string
	}{
		{
			name:     "base with trailing slash and path with leading slash",
			base:     "https://api.example.com/",
			path:     "/v1/users",
			expected: "https://api.example.com/v1/users",
		},
		{
			name:     "base without trailing slash and path with leading slash",
			base:     "https://api.example.com",
			path:     "/v1/users",
			expected: "https://api.example.com/v1/users",
		},
		{
			name:     "base with trailing slash and path without leading slash",
			base:     "https://api.example.com/",
			path:     "v1/users",
			expected: "https://api.example.com/v1/users",
		},
		{
			name:     "base without trailing slash and path without leading slash",
			base:     "https://api.example.com",
			path:     "v1/users",
			expected: "https://api.example.com/v1/users",
		},
		{
			name:     "empty base",
			base:     "",
			path:     "/v1/users",
			expected: "/v1/users",
		},
		{
			name:     "base with multiple trailing slashes",
			base:     "https://api.example.com///",
			path:     "/v1/users",
			expected: "https://api.example.com/v1/users",
		},
		{
			name:     "path with multiple leading slashes",
			base:     "https://api.example.com",
			path:     "///v1/users",
			expected: "https://api.example.com/v1/users",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := joinURL(tt.base, tt.path)
			if result != tt.expected {
				t.Errorf("joinURL(%q, %q) = %q; want %q", tt.base, tt.path, result, tt.expected)
			}
		})
	}
}

// TestIsRetryableStatus tests the retry status detection
func TestIsRetryableStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   int
		expected bool
	}{
		{
			name:     "429 Too Many Requests - retryable",
			status:   http.StatusTooManyRequests,
			expected: true,
		},
		{
			name:     "500 Internal Server Error - retryable",
			status:   http.StatusInternalServerError,
			expected: true,
		},
		{
			name:     "502 Bad Gateway - retryable",
			status:   http.StatusBadGateway,
			expected: true,
		},
		{
			name:     "503 Service Unavailable - retryable",
			status:   http.StatusServiceUnavailable,
			expected: true,
		},
		{
			name:     "504 Gateway Timeout - retryable",
			status:   http.StatusGatewayTimeout,
			expected: true,
		},
		{
			name:     "200 OK - not retryable",
			status:   http.StatusOK,
			expected: false,
		},
		{
			name:     "400 Bad Request - not retryable",
			status:   http.StatusBadRequest,
			expected: false,
		},
		{
			name:     "401 Unauthorized - not retryable",
			status:   http.StatusUnauthorized,
			expected: false,
		},
		{
			name:     "404 Not Found - not retryable",
			status:   http.StatusNotFound,
			expected: false,
		},
		{
			name:     "409 Conflict - not retryable",
			status:   http.StatusConflict,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isRetryableStatus(tt.status)
			if result != tt.expected {
				t.Errorf("isRetryableStatus(%d) = %v; want %v", tt.status, result, tt.expected)
			}
		})
	}
}

// TestIsMutatingMethod tests the mutating method detection
func TestIsMutatingMethod(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		expected bool
	}{
		{
			name:     "POST is mutating",
			method:   http.MethodPost,
			expected: true,
		},
		{
			name:     "PUT is mutating",
			method:   http.MethodPut,
			expected: true,
		},
		{
			name:     "PATCH is mutating",
			method:   http.MethodPatch,
			expected: true,
		},
		{
			name:     "DELETE is mutating",
			method:   http.MethodDelete,
			expected: true,
		},
		{
			name:     "GET is not mutating",
			method:   http.MethodGet,
			expected: false,
		},
		{
			name:     "HEAD is not mutating",
			method:   http.MethodHead,
			expected: false,
		},
		{
			name:     "OPTIONS is not mutating",
			method:   http.MethodOptions,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMutatingMethod(tt.method)
			if result != tt.expected {
				t.Errorf("isMutatingMethod(%q) = %v; want %v", tt.method, result, tt.expected)
			}
		})
	}
}

// TestWithIdempotencyKey tests adding idempotency key to context
func TestWithIdempotencyKey(t *testing.T) {
	tests := []struct {
		name string
		key  string
	}{
		{
			name: "add valid idempotency key",
			key:  "550e8400-e29b-41d4-a716-446655440000",
		},
		{
			name: "add empty idempotency key",
			key:  "",
		},
		{
			name: "add alphanumeric idempotency key",
			key:  "abc123def456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			ctxWithKey := WithIdempotencyKey(ctx, tt.key)

			// Verify context is not nil
			if ctxWithKey == nil {
				t.Fatal("WithIdempotencyKey returned nil context")
			}

			// Verify the key can be retrieved
			retrievedKey, ok := getIdempotencyKey(ctxWithKey)

			if tt.key == "" {
				// Empty key should return ok=false
				if ok {
					t.Errorf("getIdempotencyKey with empty key returned ok=true; want false")
				}
			} else {
				// Non-empty key should return ok=true and match
				if !ok {
					t.Errorf("getIdempotencyKey returned ok=false; want true")
				}
				if retrievedKey != tt.key {
					t.Errorf("getIdempotencyKey returned %q; want %q", retrievedKey, tt.key)
				}
			}
		})
	}
}

// TestGetIdempotencyKey tests extracting idempotency key from context
func TestGetIdempotencyKey(t *testing.T) {
	tests := []struct {
		name      string
		setupCtx  func() context.Context
		expectKey string
		expectOk  bool
	}{
		{
			name: "context with valid key",
			setupCtx: func() context.Context {
				return WithIdempotencyKey(context.Background(), "test-key-123")
			},
			expectKey: "test-key-123",
			expectOk:  true,
		},
		{
			name: "context without key",
			setupCtx: func() context.Context {
				return context.Background()
			},
			expectKey: "",
			expectOk:  false,
		},
		{
			name: "context with empty key",
			setupCtx: func() context.Context {
				return WithIdempotencyKey(context.Background(), "")
			},
			expectKey: "",
			expectOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.setupCtx()
			key, ok := getIdempotencyKey(ctx)

			if ok != tt.expectOk {
				t.Errorf("getIdempotencyKey ok = %v; want %v", ok, tt.expectOk)
			}
			if key != tt.expectKey {
				t.Errorf("getIdempotencyKey key = %q; want %q", key, tt.expectKey)
			}
		})
	}
}

// mockHook is a test hook that records calls
type mockHook struct {
	beforeRequestCalls []mockBeforeRequest
	afterResponseCalls []mockAfterResponse
}

type mockBeforeRequest struct {
	method string
	path   string
	body   any
}

type mockAfterResponse struct {
	method     string
	path       string
	statusCode int
	duration   time.Duration
	err        error
}

func (h *mockHook) BeforeRequest(ctx context.Context, method, path string, body any) {
	h.beforeRequestCalls = append(h.beforeRequestCalls, mockBeforeRequest{
		method: method,
		path:   path,
		body:   body,
	})
}

func (h *mockHook) AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error) {
	h.afterResponseCalls = append(h.afterResponseCalls, mockAfterResponse{
		method:     method,
		path:       path,
		statusCode: statusCode,
		duration:   duration,
		err:        err,
	})
}

// TestNotifyHooks tests the hook notification function
func TestNotifyHooks(t *testing.T) {
	tests := []struct {
		name       string
		hooks      []Hook
		method     string
		path       string
		status     int
		duration   time.Duration
		err        error
		expectCall bool
	}{
		{
			name:       "single hook called",
			hooks:      []Hook{&mockHook{}},
			method:     "GET",
			path:       "/users",
			status:     200,
			duration:   100 * time.Millisecond,
			err:        nil,
			expectCall: true,
		},
		{
			name:       "multiple hooks called",
			hooks:      []Hook{&mockHook{}, &mockHook{}},
			method:     "POST",
			path:       "/accounts",
			status:     201,
			duration:   200 * time.Millisecond,
			err:        nil,
			expectCall: true,
		},
		{
			name:       "no hooks",
			hooks:      []Hook{},
			method:     "DELETE",
			path:       "/resource",
			status:     204,
			duration:   50 * time.Millisecond,
			err:        nil,
			expectCall: false,
		},
		{
			name:       "hooks called with error",
			hooks:      []Hook{&mockHook{}},
			method:     "GET",
			path:       "/fail",
			status:     500,
			duration:   150 * time.Millisecond,
			err:        fmt.Errorf("test error"),
			expectCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			notifyHooks(ctx, tt.hooks, tt.method, tt.path, tt.status, tt.duration, tt.err)

			if tt.expectCall {
				for _, hook := range tt.hooks {
					mh := hook.(*mockHook)
					if len(mh.afterResponseCalls) != 1 {
						t.Errorf("expected 1 AfterResponse call, got %d", len(mh.afterResponseCalls))
						continue
					}

					call := mh.afterResponseCalls[0]
					if call.method != tt.method {
						t.Errorf("method = %q; want %q", call.method, tt.method)
					}
					if call.path != tt.path {
						t.Errorf("path = %q; want %q", call.path, tt.path)
					}
					if call.statusCode != tt.status {
						t.Errorf("statusCode = %d; want %d", call.statusCode, tt.status)
					}
					if call.err != tt.err {
						t.Errorf("err = %v; want %v", call.err, tt.err)
					}
				}
			}
		})
	}
}

// newTestTLSConfig creates a TLS config that accepts the test server's certificate
func newTestTLSConfig(server *httptest.Server) *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true, // Only for testing!
	}
}

// TestHTTPGetRequest tests successful GET request
func TestHTTPGetRequest(t *testing.T) {
	responseData := map[string]string{
		"id":   "123",
		"name": "Test User",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and path
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/users/123" {
			t.Errorf("expected path /v1/users/123, got %s", r.URL.Path)
		}

		// Verify headers
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("expected Accept: application/json, got %s", r.Header.Get("Accept"))
		}
		if r.Header.Get(APIKeyHeader) != "test-api-key" {
			t.Errorf("expected %s: test-api-key, got %s", APIKeyHeader, r.Header.Get(APIKeyHeader))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var response map[string]string
	err = client.get(context.Background(), "/v1/users/123", &response)
	if err != nil {
		t.Fatalf("GET request failed: %v", err)
	}

	if response["id"] != responseData["id"] {
		t.Errorf("response id = %q; want %q", response["id"], responseData["id"])
	}
	if response["name"] != responseData["name"] {
		t.Errorf("response name = %q; want %q", response["name"], responseData["name"])
	}
}

// TestHTTPPostRequest tests successful POST request
func TestHTTPPostRequest(t *testing.T) {
	requestData := map[string]string{
		"name":  "New User",
		"email": "user@example.com",
	}
	responseData := map[string]string{
		"id":   "456",
		"name": "New User",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/users" {
			t.Errorf("expected path /v1/users, got %s", r.URL.Path)
		}

		// Verify Content-Type
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Verify request body
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body["name"] != requestData["name"] {
			t.Errorf("request name = %q; want %q", body["name"], requestData["name"])
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var response map[string]string
	err = client.post(context.Background(), "/v1/users", requestData, &response)
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	if response["id"] != responseData["id"] {
		t.Errorf("response id = %q; want %q", response["id"], responseData["id"])
	}
}

// TestHTTPPutRequest tests successful PUT request
func TestHTTPPutRequest(t *testing.T) {
	requestData := map[string]string{
		"name": "Updated User",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(requestData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var response map[string]string
	err = client.put(context.Background(), "/v1/users/123", requestData, &response)
	if err != nil {
		t.Fatalf("PUT request failed: %v", err)
	}
}

// TestHTTPDeleteRequest tests successful DELETE request
func TestHTTPDeleteRequest(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/users/123" {
			t.Errorf("expected path /v1/users/123, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	err = client.delete(context.Background(), "/v1/users/123", nil)
	if err != nil {
		t.Fatalf("DELETE request failed: %v", err)
	}
}

// TestHTTPErrorResponses tests various error responses
func TestHTTPErrorResponses(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectErrType  string
		expectErrMsg   string
	}{
		{
			name:          "400 Bad Request - validation error",
			statusCode:    http.StatusBadRequest,
			responseBody:  `[{"code":"INVALID_FIELD","field":"email","message":"invalid email format"}]`,
			expectErrType: "*client.ValidationError",
			expectErrMsg:  "validation error: invalid email format (field: email)",
		},
		{
			name:          "401 Unauthorized",
			statusCode:    http.StatusUnauthorized,
			responseBody:  `{"message":"invalid credentials"}`,
			expectErrType: "*client.UnauthorizedError",
			expectErrMsg:  "unauthorized [401]: invalid credentials",
		},
		{
			name:          "404 Not Found",
			statusCode:    http.StatusNotFound,
			responseBody:  `{"message":"user not found","resource":"user"}`,
			expectErrType: "*client.NotFoundError",
			expectErrMsg:  "not found [404]: user not found (resource: user)",
		},
		{
			name:          "500 Internal Server Error",
			statusCode:    http.StatusInternalServerError,
			responseBody:  `{"message":"internal error"}`,
			expectErrType: "*client.ExceptionError",
			expectErrMsg:  "server error [500]: internal error",
		},
		{
			name:          "503 Service Unavailable",
			statusCode:    http.StatusServiceUnavailable,
			responseBody:  `{"message":"service unavailable"}`,
			expectErrType: "*client.IntegrationError",
			expectErrMsg:  "integration error [503]: service unavailable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
			if err != nil {
				t.Fatalf("failed to create client: %v", err)
			}
			defer client.Close()

			var response map[string]string
			err = client.get(context.Background(), "/test", &response)
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			if err.Error() != tt.expectErrMsg {
				t.Errorf("error message = %q; want %q", err.Error(), tt.expectErrMsg)
			}
		})
	}
}

// TestHTTPRetryLogicGET tests retry logic for GET requests
func TestHTTPRetryLogicGET(t *testing.T) {
	attemptCount := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount < 3 {
			// Return retryable error for first 2 attempts
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"message":"service temporarily unavailable"}`))
			return
		}
		// Succeed on 3rd attempt
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"success"}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var response map[string]string
	err = client.get(context.Background(), "/test", &response)
	if err != nil {
		t.Fatalf("GET request should succeed after retries: %v", err)
	}

	if attemptCount != 3 {
		t.Errorf("expected 3 attempts, got %d", attemptCount)
	}

	if response["result"] != "success" {
		t.Errorf("response result = %q; want %q", response["result"], "success")
	}
}

// TestHTTPRetryLogicPUT tests retry logic for PUT requests
func TestHTTPRetryLogicPUT(t *testing.T) {
	attemptCount := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount < 2 {
			// Return retryable error for first attempt
			w.WriteHeader(http.StatusBadGateway)
			_, _ = w.Write([]byte(`{"message":"bad gateway"}`))
			return
		}
		// Succeed on 2nd attempt
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"updated"}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	requestData := map[string]string{"field": "value"}
	var response map[string]string
	err = client.put(context.Background(), "/test", requestData, &response)
	if err != nil {
		t.Fatalf("PUT request should succeed after retry: %v", err)
	}

	if attemptCount < 2 {
		t.Errorf("expected at least 2 attempts, got %d", attemptCount)
	}
}

// TestHTTPNoRetryForPOST tests that POST requests don't retry
func TestHTTPNoRetryForPOST(t *testing.T) {
	attemptCount := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		// Always return error
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"server error"}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	requestData := map[string]string{"field": "value"}
	var response map[string]string
	err = client.post(context.Background(), "/test", requestData, &response)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if attemptCount != 1 {
		t.Errorf("expected exactly 1 attempt (no retries), got %d", attemptCount)
	}
}

// TestIdempotencyKeyHeader tests that idempotency key is sent in headers
func TestIdempotencyKeyHeader(t *testing.T) {
	testKey := "test-idempotency-key-12345"
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify idempotency key header
		receivedKey := r.Header.Get(IdempotencyKeyHeader)
		if receivedKey != testKey {
			t.Errorf("idempotency key header = %q; want %q", receivedKey, testKey)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"result":"created"}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	ctx := WithIdempotencyKey(context.Background(), testKey)
	requestData := map[string]string{"field": "value"}
	var response map[string]string
	err = client.post(ctx, "/test", requestData, &response)
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}
}

// TestJSONMarshaling tests JSON request/response marshaling
func TestJSONMarshaling(t *testing.T) {
	type TestRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	type TestResponse struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"createdAt"`
	}

	reqData := TestRequest{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	respData := TestResponse{
		ID:        "abc123",
		Name:      "John Doe",
		CreatedAt: "2024-01-01T00:00:00Z",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request is properly marshaled
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		var receivedReq TestRequest
		if err := json.Unmarshal(body, &receivedReq); err != nil {
			t.Fatalf("failed to unmarshal request: %v", err)
		}

		if receivedReq.Name != reqData.Name {
			t.Errorf("request name = %q; want %q", receivedReq.Name, reqData.Name)
		}
		if receivedReq.Email != reqData.Email {
			t.Errorf("request email = %q; want %q", receivedReq.Email, reqData.Email)
		}
		if receivedReq.Age != reqData.Age {
			t.Errorf("request age = %d; want %d", receivedReq.Age, reqData.Age)
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var response TestResponse
	err = client.post(context.Background(), "/test", reqData, &response)
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	// Verify response is properly unmarshaled
	if response.ID != respData.ID {
		t.Errorf("response ID = %q; want %q", response.ID, respData.ID)
	}
	if response.Name != respData.Name {
		t.Errorf("response Name = %q; want %q", response.Name, respData.Name)
	}
	if response.CreatedAt != respData.CreatedAt {
		t.Errorf("response CreatedAt = %q; want %q", response.CreatedAt, respData.CreatedAt)
	}
}

// TestHooksIntegration tests that hooks are called correctly during requests
func TestHooksIntegration(t *testing.T) {
	hook := &mockHook{}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"ok"}`))
	}))
	defer server.Close()

	client, err := New(
		server.URL,
		"test-api-key",
		newTestTLSConfig(server),
		WithTimeout(5*time.Second),
		WithHooks(hook),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	requestData := map[string]string{"test": "data"}
	var response map[string]string
	err = client.post(context.Background(), "/test", requestData, &response)
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	// Verify BeforeRequest was called
	if len(hook.beforeRequestCalls) != 1 {
		t.Fatalf("expected 1 BeforeRequest call, got %d", len(hook.beforeRequestCalls))
	}
	beforeCall := hook.beforeRequestCalls[0]
	if beforeCall.method != http.MethodPost {
		t.Errorf("BeforeRequest method = %q; want %q", beforeCall.method, http.MethodPost)
	}
	if beforeCall.path != "/test" {
		t.Errorf("BeforeRequest path = %q; want %q", beforeCall.path, "/test")
	}

	// Verify AfterResponse was called
	if len(hook.afterResponseCalls) != 1 {
		t.Fatalf("expected 1 AfterResponse call, got %d", len(hook.afterResponseCalls))
	}
	afterCall := hook.afterResponseCalls[0]
	if afterCall.method != http.MethodPost {
		t.Errorf("AfterResponse method = %q; want %q", afterCall.method, http.MethodPost)
	}
	if afterCall.path != "/test" {
		t.Errorf("AfterResponse path = %q; want %q", afterCall.path, "/test")
	}
	if afterCall.statusCode != http.StatusOK {
		t.Errorf("AfterResponse statusCode = %d; want %d", afterCall.statusCode, http.StatusOK)
	}
	if afterCall.err != nil {
		t.Errorf("AfterResponse err = %v; want nil", afterCall.err)
	}
}

// TestAutoIdempotency tests automatic idempotency key generation
func TestAutoIdempotency(t *testing.T) {
	var receivedKey string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedKey = r.Header.Get(IdempotencyKeyHeader)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"result":"created"}`))
	}))
	defer server.Close()

	client, err := New(
		server.URL,
		"test-api-key",
		newTestTLSConfig(server),
		WithTimeout(5*time.Second),
		WithAutoIdempotency(),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	requestData := map[string]string{"field": "value"}
	var response map[string]string
	err = client.post(context.Background(), "/test", requestData, &response)
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	// Verify an idempotency key was auto-generated and sent
	if receivedKey == "" {
		t.Error("expected auto-generated idempotency key, got empty string")
	}

	// Verify it looks like a UUID (basic check)
	if !strings.Contains(receivedKey, "-") || len(receivedKey) < 32 {
		t.Errorf("auto-generated idempotency key doesn't look like a UUID: %q", receivedKey)
	}
}

// TestAutoIdempotencyNotOverridingManual tests that manual key takes precedence
func TestAutoIdempotencyNotOverridingManual(t *testing.T) {
	manualKey := "my-custom-idempotency-key"
	var receivedKey string

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedKey = r.Header.Get(IdempotencyKeyHeader)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"result":"created"}`))
	}))
	defer server.Close()

	client, err := New(
		server.URL,
		"test-api-key",
		newTestTLSConfig(server),
		WithTimeout(5*time.Second),
		WithAutoIdempotency(),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	// Manually set idempotency key in context
	ctx := WithIdempotencyKey(context.Background(), manualKey)
	requestData := map[string]string{"field": "value"}
	var response map[string]string
	err = client.post(ctx, "/test", requestData, &response)
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	// Verify the manual key was used, not auto-generated
	if receivedKey != manualKey {
		t.Errorf("idempotency key = %q; want manual key %q", receivedKey, manualKey)
	}
}

// TestHTTPPatchRequest tests PATCH request
func TestHTTPPatchRequest(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("expected PATCH request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"patched"}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	requestData := map[string]string{"field": "value"}
	var response map[string]string
	err = client.patch(context.Background(), "/test", requestData, &response)
	if err != nil {
		t.Fatalf("PATCH request failed: %v", err)
	}

	if response["result"] != "patched" {
		t.Errorf("response result = %q; want %q", response["result"], "patched")
	}
}

// TestHTTPDeleteWithBody tests DELETE request with body
func TestHTTPDeleteWithBody(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE request, got %s", r.Method)
		}

		// Verify body is present
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}

		if len(body) == 0 {
			t.Error("expected request body, got empty")
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"deleted"}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	requestData := map[string]string{"reason": "cleanup"}
	var response map[string]string
	err = client.deleteWithBody(context.Background(), "/test", requestData, &response)
	if err != nil {
		t.Fatalf("DELETE request failed: %v", err)
	}

	if response["result"] != "deleted" {
		t.Errorf("response result = %q; want %q", response["result"], "deleted")
	}
}

// TestHTTPEmptyResponseBody tests handling of empty response body
func TestHTTPEmptyResponseBody(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		// No body
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var response map[string]string
	err = client.get(context.Background(), "/test", &response)
	if err != nil {
		t.Fatalf("GET request with empty response failed: %v", err)
	}

	// Response should be empty but not cause error
	if len(response) != 0 {
		t.Errorf("expected empty response, got %+v", response)
	}
}

// TestHTTPNilResponse tests passing nil as response parameter
func TestHTTPNilResponse(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ignored":"data"}`))
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	// Should not panic when response is nil
	err = client.get(context.Background(), "/test", nil)
	if err != nil {
		t.Fatalf("GET request with nil response failed: %v", err)
	}
}

// TestHTTPContextCancellation tests request cancellation via context
func TestHTTPContextCancellation(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate slow response
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(10*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var response map[string]string
	err = client.get(ctx, "/test", &response)
	if err == nil {
		t.Fatal("expected context cancellation error, got nil")
	}

	if !strings.Contains(err.Error(), "context") && !strings.Contains(err.Error(), "deadline") {
		t.Errorf("expected context/deadline error, got: %v", err)
	}
}

// TestPanicRecovery tests that panics in HTTP operations are recovered
func TestPanicRecovery(t *testing.T) {
	// Create a handler that will cause a panic in the client
	// We simulate this by using a hook that panics
	panicHook := &panicOnBeforeRequestHook{}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"ok"}`))
	}))
	defer server.Close()

	client, err := New(
		server.URL,
		"test-api-key",
		newTestTLSConfig(server),
		WithTimeout(5*time.Second),
		WithHooks(panicHook),
	)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	var response map[string]string
	err = client.get(context.Background(), "/test", &response)

	// Should return a PanicError, not crash
	if err == nil {
		t.Fatal("expected panic error, got nil")
	}

	// Verify it's a PanicError and can be checked with errors.Is
	if !errors.Is(err, ErrPanic) {
		t.Errorf("expected errors.Is(err, ErrPanic) to be true, got false. Error: %v", err)
	}

	// Verify it can be asserted with errors.As
	var panicErr *PanicError
	if !errors.As(err, &panicErr) {
		t.Fatalf("expected errors.As to succeed for PanicError")
	}

	// Verify the panic message is captured
	if panicErr.Message != "intentional panic for testing" {
		t.Errorf("PanicError.Message = %q; want %q", panicErr.Message, "intentional panic for testing")
	}

	// Verify stack trace is captured
	if panicErr.Stack == "" {
		t.Error("PanicError.Stack should not be empty")
	}
}

// panicOnBeforeRequestHook is a test hook that panics on BeforeRequest
type panicOnBeforeRequestHook struct{}

func (h *panicOnBeforeRequestHook) BeforeRequest(ctx context.Context, method, path string, body any) {
	panic("intentional panic for testing")
}

func (h *panicOnBeforeRequestHook) AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error) {
}
