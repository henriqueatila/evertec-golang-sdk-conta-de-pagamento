package observability

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	"go.opentelemetry.io/otel/trace"
)

// ===========================
// Tracing Tests
// ===========================

func TestNewTracer(t *testing.T) {
	tracer := NewTracer()
	if tracer == nil {
		t.Fatal("NewTracer() returned nil")
	}
	if tracer.tracer == nil {
		t.Error("NewTracer() tracer field is nil")
	}
	if tracer.propagator == nil {
		t.Error("NewTracer() propagator field is nil")
	}
}

func TestNewTracerWithProvider(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
	)

	tracer := NewTracerWithProvider(tp)
	if tracer == nil {
		t.Fatal("NewTracerWithProvider() returned nil")
	}
	if tracer.tracer == nil {
		t.Error("NewTracerWithProvider() tracer field is nil")
	}
	if tracer.propagator == nil {
		t.Error("NewTracerWithProvider() propagator field is nil")
	}
}

func TestStartSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
	)

	tracer := NewTracerWithProvider(tp)
	ctx := context.Background()

	tests := []struct {
		name     string
		spanName string
		opts     []trace.SpanStartOption
	}{
		{
			name:     "basic span",
			spanName: "test-operation",
			opts:     nil,
		},
		{
			name:     "span with attributes",
			spanName: "test-operation-with-attrs",
			opts: []trace.SpanStartOption{
				trace.WithAttributes(attribute.String("key", "value")),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, span := tracer.StartSpan(ctx, tt.spanName, tt.opts...)
			if ctx == nil {
				t.Error("StartSpan() returned nil context")
			}
			if span == nil {
				t.Error("StartSpan() returned nil span")
			}
			span.End()
		})
	}
}

func TestStartHTTPSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
	)

	tracer := NewTracerWithProvider(tp)
	ctx := context.Background()

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{
			name:   "GET request",
			method: "GET",
			path:   "/api/v1/accounts",
		},
		{
			name:   "POST request",
			method: "POST",
			path:   "/api/v1/transactions",
		},
		{
			name:   "PUT request",
			method: "PUT",
			path:   "/api/v1/accounts/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, span := tracer.StartHTTPSpan(ctx, tt.method, tt.path)
			if ctx == nil {
				t.Error("StartHTTPSpan() returned nil context")
			}
			if span == nil {
				t.Error("StartHTTPSpan() returned nil span")
			}
			span.End()

			// Verify span attributes
			spans := exporter.GetSpans()
			if len(spans) == 0 {
				t.Fatal("No spans were exported")
			}

			lastSpan := spans[len(spans)-1]
			expectedName := tt.method + " " + tt.path
			if lastSpan.Name != expectedName {
				t.Errorf("Span name = %v, want %v", lastSpan.Name, expectedName)
			}

			// Verify attributes
			attrs := lastSpan.Attributes
			foundMethod := false
			foundURL := false
			for _, attr := range attrs {
				if attr.Key == "http.method" && attr.Value.AsString() == tt.method {
					foundMethod = true
				}
				if attr.Key == "http.url" && attr.Value.AsString() == tt.path {
					foundURL = true
				}
			}
			if !foundMethod {
				t.Error("http.method attribute not found or incorrect")
			}
			if !foundURL {
				t.Error("http.url attribute not found or incorrect")
			}
		})
	}
}

func TestEndSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
	)

	tracer := NewTracerWithProvider(tp)
	ctx := context.Background()

	tests := []struct {
		name           string
		err            error
		expectedStatus codes.Code
	}{
		{
			name:           "success - no error",
			err:            nil,
			expectedStatus: codes.Ok,
		},
		{
			name:           "error case",
			err:            errors.New("test error"),
			expectedStatus: codes.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, span := tracer.StartSpan(ctx, "test-span")
			tracer.EndSpan(span, tt.err)

			spans := exporter.GetSpans()
			if len(spans) == 0 {
				t.Fatal("No spans were exported")
			}

			lastSpan := spans[len(spans)-1]
			if lastSpan.Status.Code != tt.expectedStatus {
				t.Errorf("Span status = %v, want %v", lastSpan.Status.Code, tt.expectedStatus)
			}

			if tt.err != nil {
				if len(lastSpan.Events) == 0 {
					t.Error("Expected error event but found none")
				}
			}
		})
	}
}

func TestRecordHTTPResponse(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
	)

	tracer := NewTracerWithProvider(tp)
	ctx := context.Background()

	tests := []struct {
		name       string
		statusCode int
		duration   time.Duration
	}{
		{
			name:       "200 OK",
			statusCode: 200,
			duration:   100 * time.Millisecond,
		},
		{
			name:       "404 Not Found",
			statusCode: 404,
			duration:   50 * time.Millisecond,
		},
		{
			name:       "500 Internal Server Error",
			statusCode: 500,
			duration:   200 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, span := tracer.StartSpan(ctx, "test-span")
			tracer.RecordHTTPResponse(span, tt.statusCode, tt.duration)
			span.End()

			spans := exporter.GetSpans()
			if len(spans) == 0 {
				t.Fatal("No spans were exported")
			}

			lastSpan := spans[len(spans)-1]
			foundStatusCode := false
			foundDuration := false

			for _, attr := range lastSpan.Attributes {
				if attr.Key == "http.status_code" && int(attr.Value.AsInt64()) == tt.statusCode {
					foundStatusCode = true
				}
				if attr.Key == "http.response_time_ms" && attr.Value.AsInt64() == tt.duration.Milliseconds() {
					foundDuration = true
				}
			}

			if !foundStatusCode {
				t.Error("http.status_code attribute not found or incorrect")
			}
			if !foundDuration {
				t.Error("http.response_time_ms attribute not found or incorrect")
			}
		})
	}
}

func TestInjectHTTPHeaders(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exporter),
	)

	tracer := NewTracerWithProvider(tp)
	ctx := context.Background()

	ctx, span := tracer.StartSpan(ctx, "test-span")
	defer span.End()

	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Inject headers
	tracer.InjectHTTPHeaders(ctx, req)

	// The propagator.Inject might not add headers if the propagator is not properly configured
	// We just verify it doesn't panic and completes successfully
	// In real usage with a properly configured propagator (like W3C TraceContext), headers would be added
}

func TestExtractHTTPHeaders(t *testing.T) {
	tracer := NewTracer()
	ctx := context.Background()

	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add some trace headers
	req.Header.Set("traceparent", "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")

	extractedCtx := tracer.ExtractHTTPHeaders(ctx, req)
	if extractedCtx == nil {
		t.Error("ExtractHTTPHeaders() returned nil context")
	}
}

func TestSpanAttributesToOtelAttributes(t *testing.T) {
	accountID := int64(12345)

	tests := []struct {
		name           string
		attrs          SpanAttributes
		expectedCount  int
		checkAccountID bool
		checkReqSize   bool
	}{
		{
			name: "all fields set",
			attrs: SpanAttributes{
				Operation:   "CreateAccount",
				AccountID:   &accountID,
				Endpoint:    "/api/v1/accounts",
				Method:      "POST",
				RequestSize: 1024,
			},
			expectedCount:  5,
			checkAccountID: true,
			checkReqSize:   true,
		},
		{
			name: "no account ID",
			attrs: SpanAttributes{
				Operation:   "ListAccounts",
				AccountID:   nil,
				Endpoint:    "/api/v1/accounts",
				Method:      "GET",
				RequestSize: 0,
			},
			expectedCount:  3,
			checkAccountID: false,
			checkReqSize:   false,
		},
		{
			name: "with request size only",
			attrs: SpanAttributes{
				Operation:   "UpdateAccount",
				AccountID:   nil,
				Endpoint:    "/api/v1/accounts/123",
				Method:      "PUT",
				RequestSize: 512,
			},
			expectedCount:  4,
			checkAccountID: false,
			checkReqSize:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrs := tt.attrs.ToOtelAttributes()
			if len(attrs) != tt.expectedCount {
				t.Errorf("ToOtelAttributes() count = %v, want %v", len(attrs), tt.expectedCount)
			}

			// Verify required attributes
			foundOp := false
			foundMethod := false
			foundEndpoint := false
			foundAccountID := false
			foundReqSize := false

			for _, attr := range attrs {
				switch attr.Key {
				case "sdk.operation":
					if attr.Value.AsString() == tt.attrs.Operation {
						foundOp = true
					}
				case "http.method":
					if attr.Value.AsString() == tt.attrs.Method {
						foundMethod = true
					}
				case "http.url":
					if attr.Value.AsString() == tt.attrs.Endpoint {
						foundEndpoint = true
					}
				case "account.id":
					if tt.checkAccountID && attr.Value.AsInt64() == *tt.attrs.AccountID {
						foundAccountID = true
					}
				case "http.request_content_length":
					if tt.checkReqSize && int(attr.Value.AsInt64()) == tt.attrs.RequestSize {
						foundReqSize = true
					}
				}
			}

			if !foundOp {
				t.Error("sdk.operation attribute not found")
			}
			if !foundMethod {
				t.Error("http.method attribute not found")
			}
			if !foundEndpoint {
				t.Error("http.url attribute not found")
			}
			if tt.checkAccountID && !foundAccountID {
				t.Error("account.id attribute not found")
			}
			if tt.checkReqSize && !foundReqSize {
				t.Error("http.request_content_length attribute not found")
			}
		})
	}
}

// ===========================
// Metrics Tests
// ===========================

func TestNewMetrics(t *testing.T) {
	metrics, err := NewMetrics()
	if err != nil {
		t.Fatalf("NewMetrics() error = %v", err)
	}
	if metrics == nil {
		t.Fatal("NewMetrics() returned nil")
	}
	if metrics.meter == nil {
		t.Error("NewMetrics() meter field is nil")
	}
}

func TestNewMetricsWithProvider(t *testing.T) {
	tests := []struct {
		name    string
		mp      metric.MeterProvider
		wantErr bool
	}{
		{
			name:    "noop provider",
			mp:      noop.NewMeterProvider(),
			wantErr: false,
		},
		{
			name:    "sdk provider",
			mp:      sdkmetric.NewMeterProvider(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metrics, err := NewMetricsWithProvider(tt.mp)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMetricsWithProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && metrics == nil {
				t.Error("NewMetricsWithProvider() returned nil without error")
			}
		})
	}
}

func TestRecordRequest(t *testing.T) {
	reader := sdkmetric.NewManualReader()
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	metrics, err := NewMetricsWithProvider(mp)
	if err != nil {
		t.Fatalf("NewMetricsWithProvider() error = %v", err)
	}

	ctx := context.Background()

	tests := []struct {
		name       string
		method     string
		endpoint   string
		statusCode int
		duration   time.Duration
		reqSize    int64
		respSize   int64
	}{
		{
			name:       "GET request",
			method:     "GET",
			endpoint:   "/api/v1/accounts",
			statusCode: 200,
			duration:   100 * time.Millisecond,
			reqSize:    0,
			respSize:   1024,
		},
		{
			name:       "POST request",
			method:     "POST",
			endpoint:   "/api/v1/transactions",
			statusCode: 201,
			duration:   200 * time.Millisecond,
			reqSize:    512,
			respSize:   256,
		},
		{
			name:       "error response",
			method:     "GET",
			endpoint:   "/api/v1/accounts/999",
			statusCode: 404,
			duration:   50 * time.Millisecond,
			reqSize:    0,
			respSize:   128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metrics.RecordRequest(ctx, tt.method, tt.endpoint, tt.statusCode, tt.duration, tt.reqSize, tt.respSize)

			// Collect metrics
			var rm metricdata.ResourceMetrics
			err := reader.Collect(ctx, &rm)
			if err != nil {
				t.Fatalf("Failed to collect metrics: %v", err)
			}

			// Verify metrics were recorded
			if len(rm.ScopeMetrics) == 0 {
				t.Error("No metrics were recorded")
			}
		})
	}
}

func TestRecordError(t *testing.T) {
	reader := sdkmetric.NewManualReader()
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	metrics, err := NewMetricsWithProvider(mp)
	if err != nil {
		t.Fatalf("NewMetricsWithProvider() error = %v", err)
	}

	ctx := context.Background()

	tests := []struct {
		name      string
		method    string
		endpoint  string
		errorType string
	}{
		{
			name:      "timeout error",
			method:    "GET",
			endpoint:  "/api/v1/accounts",
			errorType: "timeout",
		},
		{
			name:      "network error",
			method:    "POST",
			endpoint:  "/api/v1/transactions",
			errorType: "network",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			metrics.RecordError(ctx, tt.method, tt.endpoint, tt.errorType)

			var rm metricdata.ResourceMetrics
			err := reader.Collect(ctx, &rm)
			if err != nil {
				t.Fatalf("Failed to collect metrics: %v", err)
			}

			if len(rm.ScopeMetrics) == 0 {
				t.Error("No error metrics were recorded")
			}
		})
	}
}

func TestIncrementDecrementActiveRequests(t *testing.T) {
	reader := sdkmetric.NewManualReader()
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	metrics, err := NewMetricsWithProvider(mp)
	if err != nil {
		t.Fatalf("NewMetricsWithProvider() error = %v", err)
	}

	ctx := context.Background()

	// Increment
	metrics.IncrementActiveRequests(ctx)
	metrics.IncrementActiveRequests(ctx)

	// Decrement
	metrics.DecrementActiveRequests(ctx)

	var rm metricdata.ResourceMetrics
	err = reader.Collect(ctx, &rm)
	if err != nil {
		t.Fatalf("Failed to collect metrics: %v", err)
	}

	if len(rm.ScopeMetrics) == 0 {
		t.Error("No active request metrics were recorded")
	}
}

func TestMetricsHook(t *testing.T) {
	reader := sdkmetric.NewManualReader()
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	metrics, err := NewMetricsWithProvider(mp)
	if err != nil {
		t.Fatalf("NewMetricsWithProvider() error = %v", err)
	}

	hook := NewMetricsHook(metrics)
	if hook == nil {
		t.Fatal("NewMetricsHook() returned nil")
	}
	if hook.metrics == nil {
		t.Error("MetricsHook.metrics is nil")
	}

	ctx := context.Background()
	method := "POST"
	endpoint := "/api/v1/transactions"
	body := []byte(`{"amount": 100}`)

	// Test BeforeRequest
	ctx = hook.BeforeRequest(ctx, method, endpoint, body)
	if ctx == nil {
		t.Error("BeforeRequest() returned nil context")
	}
	if hook.method != method {
		t.Errorf("BeforeRequest() method = %v, want %v", hook.method, method)
	}
	if hook.endpoint != endpoint {
		t.Errorf("BeforeRequest() endpoint = %v, want %v", hook.endpoint, endpoint)
	}
	if hook.reqSize != int64(len(body)) {
		t.Errorf("BeforeRequest() reqSize = %v, want %v", hook.reqSize, len(body))
	}

	// Test AfterResponse
	time.Sleep(10 * time.Millisecond) // Ensure some duration
	respBody := []byte(`{"id": 123}`)
	hook.AfterResponse(ctx, 201, respBody, nil)

	// Verify metrics were recorded
	var rm metricdata.ResourceMetrics
	err = reader.Collect(ctx, &rm)
	if err != nil {
		t.Fatalf("Failed to collect metrics: %v", err)
	}

	if len(rm.ScopeMetrics) == 0 {
		t.Error("No metrics were recorded by hook")
	}
}

func TestErrorType(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: "",
		},
		{
			name:     "generic error",
			err:      errors.New("test error"),
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := errorType(tt.err)
			if result != tt.expected {
				t.Errorf("errorType() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// ===========================
// Logging Tests
// ===========================

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
		opts []LoggerOption
	}{
		{
			name: "default logger",
			opts: nil,
		},
		{
			name: "with debug level",
			opts: []LoggerOption{WithLevel(slog.LevelDebug)},
		},
		{
			name: "with warn level",
			opts: []LoggerOption{WithLevel(slog.LevelWarn)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewLogger(tt.opts...)
			if logger == nil {
				t.Fatal("NewLogger() returned nil")
			}
			if logger.logger == nil {
				t.Error("NewLogger() logger field is nil")
			}
		})
	}
}

func TestWithHandler(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)

	logger := NewLogger(WithHandler(handler))
	if logger == nil {
		t.Fatal("NewLogger() with handler returned nil")
	}
	if logger.logger == nil {
		t.Error("Logger with custom handler has nil logger")
	}

	// Test that logs go to the buffer
	logger.Info("test message")
	if buf.Len() == 0 {
		t.Error("Expected log output but buffer is empty")
	}
}

func TestNewLoggerWithSlog(t *testing.T) {
	slogger := slog.New(slog.NewJSONHandler(io.Discard, nil))
	logger := NewLoggerWithSlog(slogger)
	if logger == nil {
		t.Fatal("NewLoggerWithSlog() returned nil")
	}
	if logger.logger != slogger {
		t.Error("NewLoggerWithSlog() did not use provided logger")
	}
}

func TestLoggerMethods(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := NewLogger(WithHandler(handler))

	tests := []struct {
		name   string
		logFn  func()
		expect string
	}{
		{
			name: "Debug",
			logFn: func() {
				logger.Debug("debug message", "key", "value")
			},
			expect: "debug message",
		},
		{
			name: "Info",
			logFn: func() {
				logger.Info("info message", "key", "value")
			},
			expect: "info message",
		},
		{
			name: "Warn",
			logFn: func() {
				logger.Warn("warn message", "key", "value")
			},
			expect: "warn message",
		},
		{
			name: "Error",
			logFn: func() {
				logger.Error("error message", "key", "value")
			},
			expect: "error message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			tt.logFn()
			output := buf.String()
			if !strings.Contains(output, tt.expect) {
				t.Errorf("Log output does not contain %q: %s", tt.expect, output)
			}
		})
	}
}

func TestLoggerWith(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := NewLogger(WithHandler(handler))

	childLogger := logger.With("request_id", "12345")
	if childLogger == nil {
		t.Fatal("With() returned nil")
	}

	childLogger.Info("test message")
	output := buf.String()
	if !strings.Contains(output, "request_id") {
		t.Error("Child logger does not include context attributes")
	}
	if !strings.Contains(output, "12345") {
		t.Error("Child logger does not include context values")
	}
}

func TestLogRequest(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := NewLogger(WithHandler(handler))

	ctx := context.Background()
	body := []byte(`{"key": "value"}`)

	logger.LogRequest(ctx, "POST", "/api/v1/transactions", body)

	output := buf.String()
	if !strings.Contains(output, "POST") {
		t.Error("Log output missing method")
	}
	if !strings.Contains(output, "/api/v1/transactions") {
		t.Error("Log output missing endpoint")
	}
}

func TestLogResponse(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		err        error
		expectErr  bool
	}{
		{
			name:       "success response",
			statusCode: 200,
			err:        nil,
			expectErr:  false,
		},
		{
			name:       "client error",
			statusCode: 404,
			err:        nil,
			expectErr:  true,
		},
		{
			name:       "server error",
			statusCode: 500,
			err:        errors.New("server error"),
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			handler := slog.NewJSONHandler(&buf, nil)
			logger := NewLogger(WithHandler(handler))

			ctx := context.Background()
			duration := 100 * time.Millisecond

			logger.LogResponse(ctx, "GET", "/api/v1/accounts", tt.statusCode, duration, tt.err)

			output := buf.String()
			if !strings.Contains(output, "GET") {
				t.Error("Log output missing method")
			}
			if tt.expectErr && !strings.Contains(output, "ERROR") && !strings.Contains(output, "error") {
				t.Error("Expected error level log")
			}
		})
	}
}

func TestLogRetry(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := NewLogger(WithHandler(handler))

	ctx := context.Background()
	err := errors.New("temporary failure")

	logger.LogRetry(ctx, "POST", "/api/v1/transactions", 2, err)

	output := buf.String()
	if !strings.Contains(output, "retry") || !strings.Contains(output, "WARN") {
		t.Error("Log output missing retry information")
	}
	if !strings.Contains(output, "2") {
		t.Error("Log output missing attempt number")
	}
}

func TestLogRateLimit(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, nil)
	logger := NewLogger(WithHandler(handler))

	ctx := context.Background()
	retryAfter := 30 * time.Second

	logger.LogRateLimit(ctx, "GET", "/api/v1/accounts", retryAfter)

	output := buf.String()
	if !strings.Contains(output, "rate") {
		t.Error("Log output missing rate limit information")
	}
}

func TestLoggerHook(t *testing.T) {
	var buf bytes.Buffer
	handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := NewLogger(WithHandler(handler))

	hook := NewLoggerHook(logger)
	if hook == nil {
		t.Fatal("NewLoggerHook() returned nil")
	}
	if hook.logger == nil {
		t.Error("LoggerHook.logger is nil")
	}

	ctx := context.Background()
	method := "POST"
	endpoint := "/api/v1/transactions"
	reqBody := []byte(`{"amount": 100}`)

	// Test BeforeRequest
	buf.Reset()
	ctx = hook.BeforeRequest(ctx, method, endpoint, reqBody)
	if ctx == nil {
		t.Error("BeforeRequest() returned nil context")
	}
	if !strings.Contains(buf.String(), "request") {
		t.Error("BeforeRequest() did not log request")
	}

	// Test AfterResponse
	buf.Reset()
	time.Sleep(10 * time.Millisecond)
	respBody := []byte(`{"id": 123}`)
	hook.AfterResponse(ctx, 201, respBody, nil)
	if !strings.Contains(buf.String(), "response") {
		t.Error("AfterResponse() did not log response")
	}
}

// ===========================
// Idempotency Tests
// ===========================

func TestGenerateUUIDv4(t *testing.T) {
	uuid, err := GenerateUUIDv4()
	if err != nil {
		t.Fatalf("GenerateUUIDv4() error = %v", err)
	}

	// Verify UUID format (8-4-4-4-12)
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		t.Errorf("UUID has wrong number of parts: %v", uuid)
	}
	if len(parts[0]) != 8 || len(parts[1]) != 4 || len(parts[2]) != 4 || len(parts[3]) != 4 || len(parts[4]) != 12 {
		t.Errorf("UUID parts have wrong lengths: %v", uuid)
	}

	// Verify version 4
	if parts[2][0] != '4' {
		t.Errorf("UUID is not version 4: %v", uuid)
	}

	// Test uniqueness
	uuid2, err := GenerateUUIDv4()
	if err != nil {
		t.Fatalf("GenerateUUIDv4() error = %v", err)
	}
	if uuid == uuid2 {
		t.Error("Generated UUIDs are not unique")
	}
}

func TestMustGenerateUUIDv4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("MustGenerateUUIDv4() panicked unexpectedly")
		}
	}()

	uuid := MustGenerateUUIDv4()
	if uuid == "" {
		t.Error("MustGenerateUUIDv4() returned empty string")
	}

	// Verify format
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		t.Errorf("UUID has wrong number of parts: %v", uuid)
	}
}

func TestWithIdempotencyKey(t *testing.T) {
	ctx := context.Background()
	key := "test-key-12345"

	newCtx := WithIdempotencyKey(ctx, key)
	if newCtx == nil {
		t.Fatal("WithIdempotencyKey() returned nil context")
	}

	retrievedKey, ok := GetIdempotencyKey(newCtx)
	if !ok {
		t.Error("GetIdempotencyKey() did not find key in context")
	}
	if retrievedKey != key {
		t.Errorf("GetIdempotencyKey() = %v, want %v", retrievedKey, key)
	}
}

func TestWithNewIdempotencyKey(t *testing.T) {
	ctx := context.Background()

	newCtx, key, err := WithNewIdempotencyKey(ctx)
	if err != nil {
		t.Fatalf("WithNewIdempotencyKey() error = %v", err)
	}
	if newCtx == nil {
		t.Fatal("WithNewIdempotencyKey() returned nil context")
	}
	if key == "" {
		t.Error("WithNewIdempotencyKey() returned empty key")
	}

	retrievedKey, ok := GetIdempotencyKey(newCtx)
	if !ok {
		t.Error("GetIdempotencyKey() did not find key in context")
	}
	if retrievedKey != key {
		t.Errorf("GetIdempotencyKey() = %v, want %v", retrievedKey, key)
	}
}

func TestGetIdempotencyKey(t *testing.T) {
	tests := []struct {
		name   string
		ctx    context.Context
		wantOk bool
	}{
		{
			name:   "empty context",
			ctx:    context.Background(),
			wantOk: false,
		},
		{
			name:   "context with key",
			ctx:    WithIdempotencyKey(context.Background(), "test-key"),
			wantOk: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := GetIdempotencyKey(tt.ctx)
			if ok != tt.wantOk {
				t.Errorf("GetIdempotencyKey() ok = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}

func TestNewIdempotencyProvider(t *testing.T) {
	tests := []struct {
		name             string
		opts             []IdempotencyOption
		expectAutoGen    bool
	}{
		{
			name:          "default provider",
			opts:          nil,
			expectAutoGen: true,
		},
		{
			name:          "auto-generate enabled",
			opts:          []IdempotencyOption{WithAutoGenerate(true)},
			expectAutoGen: true,
		},
		{
			name:          "auto-generate disabled",
			opts:          []IdempotencyOption{WithAutoGenerate(false)},
			expectAutoGen: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := NewIdempotencyProvider(tt.opts...)
			if provider == nil {
				t.Fatal("NewIdempotencyProvider() returned nil")
			}
			if provider.autoGenerate != tt.expectAutoGen {
				t.Errorf("autoGenerate = %v, want %v", provider.autoGenerate, tt.expectAutoGen)
			}
		})
	}
}

func TestIdempotencyProviderGetOrGenerate(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		autoGen     bool
		expectKey   bool
	}{
		{
			name:      "existing key in context",
			ctx:       WithIdempotencyKey(context.Background(), "existing-key"),
			autoGen:   true,
			expectKey: true,
		},
		{
			name:      "no key, auto-generate enabled",
			ctx:       context.Background(),
			autoGen:   true,
			expectKey: true,
		},
		{
			name:      "no key, auto-generate disabled",
			ctx:       context.Background(),
			autoGen:   false,
			expectKey: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider := NewIdempotencyProvider(WithAutoGenerate(tt.autoGen))
			newCtx, key, err := provider.GetOrGenerate(tt.ctx)
			if err != nil {
				t.Fatalf("GetOrGenerate() error = %v", err)
			}
			if newCtx == nil {
				t.Fatal("GetOrGenerate() returned nil context")
			}

			if tt.expectKey && key == "" {
				t.Error("Expected key but got empty string")
			}
			if !tt.expectKey && key != "" {
				t.Error("Did not expect key but got one")
			}
		})
	}
}

func TestIdempotencyHook(t *testing.T) {
	provider := NewIdempotencyProvider()
	hook := NewIdempotencyHook(provider)
	if hook == nil {
		t.Fatal("NewIdempotencyHook() returned nil")
	}
	if hook.provider == nil {
		t.Error("IdempotencyHook.provider is nil")
	}

	ctx := context.Background()

	tests := []struct {
		name      string
		method    string
		expectKey bool
	}{
		{
			name:      "POST request",
			method:    "POST",
			expectKey: true,
		},
		{
			name:      "PUT request",
			method:    "PUT",
			expectKey: true,
		},
		{
			name:      "PATCH request",
			method:    "PATCH",
			expectKey: true,
		},
		{
			name:      "DELETE request",
			method:    "DELETE",
			expectKey: true,
		},
		{
			name:      "GET request",
			method:    "GET",
			expectKey: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newCtx := hook.BeforeRequest(ctx, tt.method, "/api/v1/test", nil)
			if newCtx == nil {
				t.Fatal("BeforeRequest() returned nil context")
			}

			_, ok := GetIdempotencyKey(newCtx)
			if tt.expectKey && !ok {
				t.Error("Expected idempotency key in context but not found")
			}
			if !tt.expectKey && ok {
				t.Error("Did not expect idempotency key but found one")
			}

			if tt.expectKey && hook.GetKey() == "" {
				t.Error("Expected hook to store key but GetKey() returned empty")
			}
		})
	}
}

func TestIdempotencyHookAfterResponse(t *testing.T) {
	provider := NewIdempotencyProvider()
	hook := NewIdempotencyHook(provider)
	ctx := context.Background()

	// AfterResponse should be a no-op
	hook.AfterResponse(ctx, 200, nil, nil)
	// No assertion needed, just verify it doesn't panic
}

func TestIdempotencyKeyHeader(t *testing.T) {
	expected := "X-Idempotency-Key"
	if IdempotencyKeyHeader != expected {
		t.Errorf("IdempotencyKeyHeader = %v, want %v", IdempotencyKeyHeader, expected)
	}
}
