package client

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.opentelemetry.io/otel/metric/noop"
	nooptrace "go.opentelemetry.io/otel/trace/noop"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// TestOptionsUncovered tests the option functions that have 0% coverage
func TestOptionsUncovered(t *testing.T) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	t.Run("WithBaseURL", func(t *testing.T) {
		c, err := New("https://original.com", "key", tlsConfig, WithBaseURL("https://custom.com"))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if c.config.BaseURL != "https://custom.com" {
			t.Errorf("BaseURL = %q, want %q", c.config.BaseURL, "https://custom.com")
		}
	})

	t.Run("WithDefaultLogger", func(t *testing.T) {
		c, err := New("https://example.com", "key", tlsConfig, WithDefaultLogger())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if c.config.Logger == nil {
			t.Error("expected non-nil Logger")
		}
	})

	t.Run("WithTLSConfig", func(t *testing.T) {
		custom := &tls.Config{MinVersion: tls.VersionTLS13}
		c, err := New("https://example.com", "key", tlsConfig, WithTLSConfig(custom))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if c.config.TLSConfig.MinVersion != tls.VersionTLS13 {
			t.Errorf("TLSConfig.MinVersion = %d, want TLS13", c.config.TLSConfig.MinVersion)
		}
	})

	t.Run("WithServerName", func(t *testing.T) {
		c, err := New("https://example.com", "key", tlsConfig, WithServerName("custom-server"))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if c.config.ServerName != "custom-server" {
			t.Errorf("ServerName = %q, want %q", c.config.ServerName, "custom-server")
		}
	})

	t.Run("WithTracerProvider", func(t *testing.T) {
		tp := nooptrace.NewTracerProvider()
		c, err := New("https://example.com", "key", tlsConfig, WithTracerProvider(tp))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if c.config.TracerProvider == nil {
			t.Error("expected non-nil TracerProvider")
		}
	})

	t.Run("WithMeterProvider", func(t *testing.T) {
		mp := noop.NewMeterProvider()
		c, err := New("https://example.com", "key", tlsConfig, WithMeterProvider(mp))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if c.config.MeterProvider == nil {
			t.Error("expected non-nil MeterProvider")
		}
	})

	t.Run("WithTracing", func(t *testing.T) {
		c, err := New("https://example.com", "key", tlsConfig, WithTracing())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if !c.config.TracingEnabled {
			t.Error("expected TracingEnabled = true")
		}
	})

	t.Run("WithMetrics", func(t *testing.T) {
		c, err := New("https://example.com", "key", tlsConfig, WithMetrics())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if !c.config.MetricsEnabled {
			t.Error("expected MetricsEnabled = true")
		}
	})

	t.Run("WithLogger", func(t *testing.T) {
		logger := slog.Default()
		c, err := New("https://example.com", "key", tlsConfig, WithLogger(logger))
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}
		defer c.Close()
		if c.config.Logger != logger {
			t.Error("expected custom logger to be set")
		}
	})
}

// TestNoOpHook tests the NoOpHook implementation (covers BeforeRequest and AfterResponse)
func TestNoOpHook(t *testing.T) {
	hook := &NoOpHook{}
	ctx := context.Background()
	// Should not panic
	hook.BeforeRequest(ctx, "GET", "/test", nil)
	hook.AfterResponse(ctx, "GET", "/test", 200, 100*time.Millisecond, nil)
}

// TestGetTracer tests getTracer with custom TracerProvider
func TestGetTracer(t *testing.T) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
	tp := nooptrace.NewTracerProvider()

	c, err := New("https://example.com", "key", tlsConfig, WithTracerProvider(tp))
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer c.Close()

	tracer := c.getTracer()
	if tracer == nil {
		t.Error("expected non-nil tracer")
	}
}

// TestGetTracerWithTracing tests getTracer when TracingEnabled is true (uses otel.Tracer)
func TestGetTracerWithTracing(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"result":"ok"}`))
	}))
	defer server.Close()

	c, err := New(server.URL, "key", newTestTLSConfig(server),
		WithTracing(),
		WithTimeout(5*time.Second),
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer c.Close()

	var resp map[string]string
	if err := c.get(context.Background(), "/test", &resp); err != nil {
		t.Fatalf("get with tracing failed: %v", err)
	}
}

// TestNewWithCertFiles tests NewWithCertFiles with invalid paths
func TestNewWithCertFiles(t *testing.T) {
	t.Run("invalid cert files", func(t *testing.T) {
		_, err := NewWithCertFiles(
			"https://example.com",
			"api-key",
			"/nonexistent/cert.pem",
			"/nonexistent/key.pem",
			"",
		)
		if err == nil {
			t.Error("expected error for nonexistent cert files")
		}
	})
}

// TestBuildBackofficeAccountsQueryString tests the query string builder with params
func TestBuildBackofficeAccountsQueryString(t *testing.T) {
	doc := "12345678901"
	name := "John"
	email := "john@example.com"
	status := "ACTIVE"
	page := 1
	pageSize := 10

	t.Run("all params", func(t *testing.T) {
		params := &types.ListAccountsBackofficeRequest{
			Document: &doc,
			Name:     &name,
			Email:    &email,
			Status:   &status,
			Page:     &page,
			PageSize: &pageSize,
		}
		qs := buildBackofficeAccountsQueryString(params)
		if qs == "" {
			t.Error("expected non-empty query string")
		}
		// Should contain all params
		for _, expected := range []string{"document=", "name=", "email=", "status=", "page=", "pageSize="} {
			found := false
			for i := 0; i < len(qs)-len(expected)+1; i++ {
				if qs[i:i+len(expected)] == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("query string missing %q: %s", expected, qs)
			}
		}
	})

	t.Run("nil params", func(t *testing.T) {
		qs := buildBackofficeAccountsQueryString(nil)
		if qs != "" {
			t.Errorf("expected empty query string for nil params, got %q", qs)
		}
	})

	t.Run("partial params", func(t *testing.T) {
		params := &types.ListAccountsBackofficeRequest{
			Document: &doc,
		}
		qs := buildBackofficeAccountsQueryString(params)
		if qs == "" {
			t.Error("expected non-empty query string")
		}
	})
}

// TestListAccountsBackofficeWithParams exercises ListAccountsBackoffice with params
// to ensure buildBackofficeAccountsQueryString is called
func TestListAccountsBackofficeWithParams(t *testing.T) {
	doc := "12345678901"
	name := "Test"

	runEndpointTests(t, []testEndpoint{
		{
			name:     "ListAccountsBackofficeWithParams",
			method:   http.MethodGet,
			path:     "/backoffice/accounts",
			response: &types.AccountListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAccountsBackoffice(ctx, &types.ListAccountsBackofficeRequest{
					Document: &doc,
					Name:     &name,
				})
				return err
			},
		},
	})
}

// TestListHceDevicesWithParams exercises ListHceDevices with non-nil params
func TestListHceDevicesWithParams(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name:     "ListHceDevicesWithParams",
			method:   http.MethodGet,
			path:     "/backoffice/hce/devices",
			response: []types.HceDeviceResponse{},
			call: func(ctx context.Context, c *Client) error {
				params := &types.ListHceDevicesParams{}
				_, err := c.ListHceDevices(ctx, params)
				return err
			},
		},
	})
}

// TestForbiddenErrorWithoutResource tests ForbiddenError.Error() - covers line 155
func TestForbiddenErrorMethods(t *testing.T) {
	// Test with no message that would exercise the 0% path
	err := &ForbiddenError{
		StatusCode: 403,
		Message:    "access denied",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("expected non-empty error string")
	}

	// Test Unwrap with wrapped error
	wrappedErr := &ForbiddenError{
		StatusCode: 403,
		Message:    "forbidden",
		Err:        ErrForbidden,
	}
	if wrappedErr.Unwrap() != ErrForbidden {
		t.Error("expected Unwrap to return wrapped error")
	}
}

// TestUnprocessableEntityErrorMethods tests UnprocessableEntityError.Error() - covers line 196
func TestUnprocessableEntityErrorMethods(t *testing.T) {
	// Without code - exercises the else branch (line 200)
	err := &UnprocessableEntityError{
		StatusCode: 422,
		Message:    "cannot process",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("expected non-empty error string")
	}

	// With code
	errWithCode := &UnprocessableEntityError{
		StatusCode: 422,
		Code:       "INVALID_STATE",
		Message:    "invalid state",
	}
	if errWithCode.Error() == "" {
		t.Error("expected non-empty error string with code")
	}

	// Unwrap with wrapped error
	wrapped := &UnprocessableEntityError{
		StatusCode: 422,
		Message:    "test",
		Err:        ErrUnprocessable,
	}
	if wrapped.Unwrap() != ErrUnprocessable {
		t.Error("expected Unwrap to return wrapped error")
	}
}

// TestParseErrorResponseEdgeCases tests edge cases in parseErrorResponse
func TestParseErrorResponseEdgeCases(t *testing.T) {
	// Test 400 with non-array, non-object body (fallback to generic APIError)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`plain text error`))
	}))
	defer server.Close()

	c, err := New(server.URL, "key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer c.Close()

	var resp map[string]string
	reqErr := c.get(context.Background(), "/test", &resp)
	if reqErr == nil {
		t.Error("expected error for 400 response")
	}
}

// TestValidationErrorEmptyErrors tests ValidationError.Error() with empty errors slice
func TestValidationErrorEmptyErrors(t *testing.T) {
	err := &ValidationError{
		StatusCode: 400,
		Errors:     []ValidationDetail{},
	}
	msg := err.Error()
	if msg != "validation error" {
		t.Errorf("Error() = %q, want %q", msg, "validation error")
	}
}

// TestErrorsWithWrappedErr tests that Unwrap returns wrapped err when set
func TestErrorsWithWrappedErr(t *testing.T) {
	customErr := ErrAPI

	tests := []struct {
		name string
		err  error
		want error
	}{
		{"APIError with Err", &APIError{Err: customErr}, customErr},
		{"ValidationError with Err", &ValidationError{Err: customErr}, customErr},
		{"BusinessRuleError with Err", &BusinessRuleError{Err: customErr}, customErr},
		{"ExceptionError with Err", &ExceptionError{Err: customErr}, customErr},
		{"IntegrationError with Err", &IntegrationError{Err: customErr}, customErr},
		{"UnauthorizedError with Err", &UnauthorizedError{Err: customErr}, customErr},
		{"ForbiddenError with Err", &ForbiddenError{Err: customErr}, customErr},
		{"NotFoundError with Err", &NotFoundError{Err: customErr}, customErr},
		{"UnprocessableEntityError with Err", &UnprocessableEntityError{Err: customErr}, customErr},
		{"InsufficientFundsError with Err", &InsufficientFundsError{Err: customErr}, customErr},
		{"MethodNotAllowedError with Err", &MethodNotAllowedError{Err: customErr}, customErr},
		{"PreconditionFailedError with Err", &PreconditionFailedError{Err: customErr}, customErr},
		{"ThirdPartyError with Err", &ThirdPartyError{Err: customErr}, customErr},
		{"PanicError with Err", &PanicError{Err: customErr}, customErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type unwrapper interface{ Unwrap() error }
			u, ok := tt.err.(unwrapper)
			if !ok {
				t.Fatal("does not implement Unwrap")
			}
			if got := u.Unwrap(); got != tt.want {
				t.Errorf("Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestNotFoundErrorWithResource tests NotFoundError.Error() with Resource set
func TestNotFoundErrorWithResource(t *testing.T) {
	err := &NotFoundError{
		StatusCode: 404,
		Message:    "not found",
		Resource:   "account",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("expected non-empty error string")
	}
}

// TestInsufficientFundsErrorNoAmounts tests error string without Required/Available
func TestInsufficientFundsErrorNoAmounts(t *testing.T) {
	err := &InsufficientFundsError{
		StatusCode: 402,
		Message:    "insufficient funds",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("expected non-empty error string")
	}
}

// TestPreconditionFailedErrorWithCode tests PreconditionFailedError with code set
func TestPreconditionFailedErrorWithCode(t *testing.T) {
	err := &PreconditionFailedError{
		StatusCode: 412,
		Code:       "PRECONDITION",
		Message:    "precondition failed",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("expected non-empty error string")
	}
}

// TestThirdPartyErrorNoService tests ThirdPartyError.Error() without Service
func TestThirdPartyErrorNoService(t *testing.T) {
	err := &ThirdPartyError{
		StatusCode: 424,
		Message:    "third party error",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("expected non-empty error string")
	}
}

// TestAccountsErrorPaths tests error paths for account methods using 4xx responses
func TestAccountsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListAccounts_error", method: http.MethodGet, path: "/accounts",
			status:   404,
			response: map[string]string{"message": "not found"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAccounts(ctx, nil)
				return err
			},
		},
		{
			name: "GetAccountStatement_error", method: http.MethodGet, path: "/accounts/1/statement",
			status:   404,
			response: map[string]string{"message": "not found"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountStatement(ctx, 1, nil)
				return err
			},
		},
		{
			name: "ListBlockedAccounts_error", method: http.MethodGet, path: "/accounts/list/block",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBlockedAccounts(ctx)
				return err
			},
		},
		{
			name: "CreateAccount_error", method: http.MethodPost, path: "/accounts",
			status:   400,
			response: []map[string]string{{"code": "ERR", "field": "f", "message": "m"}},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateAccount(ctx, &types.ProposalAccountRequest{})
				return err
			},
		},
		{
			name: "UnlinkAccounts_error", method: http.MethodPost, path: "/accounts/unlink",
			status:   409,
			response: map[string]string{"code": "CONFLICT", "message": "conflict"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UnlinkAccounts(ctx, &types.UnlinkAccountRequest{})
				return err
			},
		},
		{
			name: "VerifyAccountExists_error", method: http.MethodPost, path: "/accounts/verify/exists",
			status:   422,
			response: map[string]string{"code": "INVALID", "message": "invalid"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.VerifyAccountExists(ctx, &types.VerifyAccountExistsRequest{})
				return err
			},
		},
		{
			name: "ListBalanceLocks_error", method: http.MethodGet, path: "/accounts/balanceLock/list",
			status:   503,
			response: map[string]string{"message": "unavailable"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBalanceLocks(ctx)
				return err
			},
		},
		{
			name: "CreateCompanyAccount_error", method: http.MethodPost, path: "/accounts/company",
			status:   400,
			response: []map[string]string{{"code": "ERR", "field": "f", "message": "m"}},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateCompanyAccount(ctx, &types.CreateCompanyAccountRequest{})
				return err
			},
		},
	})
}

// TestBanksErrorPaths tests error paths for banks/ListBanks
func TestBanksErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListBanks_error", method: http.MethodGet, path: "/banks",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBanks(ctx, nil)
				return err
			},
		},
	})
}

// TestCreditsErrorPaths tests error paths for credits
func TestCreditsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetUsableCredits_error", method: http.MethodGet, path: "/accounts/1/credits/to/use",
			status:   404,
			response: map[string]string{"message": "not found"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetUsableCredits(ctx, 1, nil)
				return err
			},
		},
	})
}

// TestProposalsErrorPaths tests error paths for proposals
func TestProposalsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListProposals_error", method: http.MethodGet, path: "/proposal",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListProposals(ctx, nil)
				return err
			},
		},
		{
			name: "ListLegalEntityProposals_error", method: http.MethodGet, path: "/proposal/legalEntities",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListLegalEntityProposals(ctx, nil)
				return err
			},
		},
	})
}

// TestCardsErrorPaths tests error paths for cards/SearchCards
func TestCardsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListCards_error", method: http.MethodGet, path: "/accounts/1/cards",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListCards(ctx, 1, nil)
				return err
			},
		},
		{
			name: "SearchCards_error", method: http.MethodGet, path: "/cards",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SearchCards(ctx, nil)
				return err
			},
		},
	})
}

// TestMedErrorPaths tests error paths for med methods
func TestMedErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListInfractionReports_error", method: http.MethodGet, path: "/pix/infraction-reports",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInfractionReports(ctx, nil)
				return err
			},
		},
		{
			name: "ListRefundSolicitations_error", method: http.MethodGet, path: "/pix/refunds",
			status:   500,
			response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListRefundSolicitations(ctx, nil)
				return err
			},
		},
	})
}

// TestPixAutomaticErrorPaths tests error paths for pix automatic
func TestPixAutomaticErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetAutomaticPixRecurrence_error", method: http.MethodGet, path: "/pix/automatic/account/1/recurrence/key123",
			status:   404,
			response: map[string]string{"message": "not found"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAutomaticPixRecurrence(ctx, 1, "key123", nil)
				return err
			},
		},
	})
}
