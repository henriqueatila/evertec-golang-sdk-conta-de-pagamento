package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/observability"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// idempotencyKeyCtxKey is the context key for idempotency key
type idempotencyKeyCtxKey struct{}

// IdempotencyKeyHeader is the header name for idempotency key (required for PIX operations)
const IdempotencyKeyHeader = "idempotencyKey"

// WithIdempotencyKey adds an idempotency key to the context.
// Use this for PIX operations to prevent duplicate transactions.
func WithIdempotencyKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, idempotencyKeyCtxKey{}, key)
}

// getIdempotencyKey extracts the idempotency key from context, if present
func getIdempotencyKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(idempotencyKeyCtxKey{}).(string)
	return key, ok && key != ""
}

// do performs an HTTP request with light retries on idempotent methods (GET/PUT)
// Includes panic recovery to prevent crashes from unexpected runtime errors
func (c *Client) do(ctx context.Context, method, path string, body, response any) (err error) {
	// Panic recovery middleware
	defer func() {
		if r := recover(); r != nil {
			stack := string(debug.Stack())
			c.config.Logger.Error("panic recovered in HTTP request",
				"method", method,
				"path", path,
				"panic", r,
				"stack", stack,
			)
			err = &PanicError{
				Message: fmt.Sprintf("%v", r),
				Stack:   stack,
			}
		}
	}()

	url := joinURL(c.config.BaseURL, path)

	// Start OpenTelemetry span if tracing is enabled
	var span trace.Span
	if c.config.TracingEnabled {
		tracer := c.getTracer()
		ctx, span = tracer.Start(ctx, method+" "+path,
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(
				attribute.String("http.method", method),
				attribute.String("http.url", url),
				attribute.String("rpc.system", "http"),
				attribute.String("rpc.service", "evertec-conta-pagamento"),
			),
		)
		defer span.End()
	}

	// Auto-generate idempotency key if enabled and not already present
	if c.config.AutoIdempotency && isMutatingMethod(method) {
		if _, ok := getIdempotencyKey(ctx); !ok {
			key, err := observability.GenerateUUIDv4()
			if err == nil {
				ctx = WithIdempotencyKey(ctx, key)
				if span != nil {
					span.SetAttributes(attribute.String("http.idempotency_key", key))
				}
			}
		}
	}

	// Serialize body if provided
	var bodyBytes []byte
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			notifyHooks(ctx, c.config.Hooks, method, path, 0, 0, err)
			if span != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, "failed to marshal request body")
			}
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyBytes = b
		if span != nil {
			span.SetAttributes(attribute.Int("http.request_content_length", len(bodyBytes)))
		}
	}

	isIdempotent := method == http.MethodGet || method == http.MethodPut
	maxAttempts := 1
	if isIdempotent {
		maxAttempts = 3
	}

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		var bodyReader io.Reader
		if len(bodyBytes) > 0 {
			bodyReader = bytes.NewReader(bodyBytes)
		}

		req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
		if err != nil {
			notifyHooks(ctx, c.config.Hooks, method, path, 0, 0, err)
			if span != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, "failed to create request")
			}
			return fmt.Errorf("failed to create request: %w", err)
		}

		if bodyReader != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", c.config.UserAgent)
		req.Header.Set(APIKeyHeader, c.config.APIKey)

		// Add idempotency key header if present in context
		if idempotencyKey, ok := getIdempotencyKey(ctx); ok {
			req.Header.Set(IdempotencyKeyHeader, idempotencyKey)
		}

		// Inject trace context into headers for distributed tracing
		if c.config.TracingEnabled {
			otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
		}

		startTime := time.Now()
		for _, hook := range c.config.Hooks {
			hook.BeforeRequest(ctx, method, path, body)
		}

		resp, err := c.http.Do(req)
		duration := time.Since(startTime)

		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}
		notifyHooks(ctx, c.config.Hooks, method, path, statusCode, duration, err)

		if err != nil {
			lastErr = fmt.Errorf("request failed: %w", err)
			if span != nil {
				span.RecordError(err)
				span.SetAttributes(
					attribute.Int("http.retry_count", attempt-1),
				)
			}
			if isIdempotent && attempt < maxAttempts {
				time.Sleep(time.Duration(attempt) * 100 * time.Millisecond)
				continue
			}
			if span != nil {
				span.SetStatus(codes.Error, "request failed")
			}
			return lastErr
		}

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			if span != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, "failed to read response body")
			}
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if span != nil {
			span.SetAttributes(
				attribute.Int("http.status_code", resp.StatusCode),
				attribute.Int64("http.response_time_ms", duration.Milliseconds()),
				attribute.Int("http.response_content_length", len(respBody)),
			)
		}

		c.config.Logger.Debug("HTTP response",
			"method", method,
			"path", path,
			"status", resp.StatusCode,
			"duration", duration,
		)

		if resp.StatusCode >= 400 {
			parseErr := parseErrorResponse(resp, respBody)
			if isIdempotent && attempt < maxAttempts && isRetryableStatus(resp.StatusCode) {
				lastErr = parseErr
				time.Sleep(time.Duration(attempt) * 100 * time.Millisecond)
				continue
			}
			if span != nil {
				span.RecordError(parseErr)
				span.SetStatus(codes.Error, fmt.Sprintf("HTTP %d", resp.StatusCode))
			}
			return parseErr
		}

		if response != nil && len(respBody) > 0 {
			if err := json.Unmarshal(respBody, response); err != nil {
				if span != nil {
					span.RecordError(err)
					span.SetStatus(codes.Error, "failed to unmarshal response")
				}
				return fmt.Errorf("failed to unmarshal response: %w", err)
			}
		}

		if span != nil {
			span.SetStatus(codes.Ok, "")
		}
		return nil
	}

	if lastErr != nil {
		if span != nil {
			span.SetStatus(codes.Error, "max retries exceeded")
		}
		return lastErr
	}
	return errors.New("request failed without response")
}

// getTracer returns the appropriate tracer based on config
func (c *Client) getTracer() trace.Tracer {
	if c.config.TracerProvider != nil {
		return c.config.TracerProvider.Tracer(observability.TracerName)
	}
	return otel.Tracer(observability.TracerName)
}

func isMutatingMethod(method string) bool {
	return method == http.MethodPost || method == http.MethodPut ||
		method == http.MethodPatch || method == http.MethodDelete
}

func isRetryableStatus(status int) bool {
	switch status {
	case http.StatusTooManyRequests, http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return true
	default:
		return false
	}
}

func joinURL(base, path string) string {
	if base == "" {
		return path
	}
	return strings.TrimRight(base, "/") + "/" + strings.TrimLeft(path, "/")
}

func notifyHooks(ctx context.Context, hooks []Hook, method, path string, status int, duration time.Duration, err error) {
	for _, hook := range hooks {
		hook.AfterResponse(ctx, method, path, status, duration, err)
	}
}

// get performs a GET request
func (c *Client) get(ctx context.Context, path string, response any) error {
	return c.do(ctx, http.MethodGet, path, nil, response)
}

// post performs a POST request
func (c *Client) post(ctx context.Context, path string, body, response any) error {
	return c.do(ctx, http.MethodPost, path, body, response)
}

// put performs a PUT request
func (c *Client) put(ctx context.Context, path string, body, response any) error {
	return c.do(ctx, http.MethodPut, path, body, response)
}

// delete performs a DELETE request
func (c *Client) delete(ctx context.Context, path string, response any) error {
	return c.do(ctx, http.MethodDelete, path, nil, response)
}

// deleteWithBody performs a DELETE request with a body
func (c *Client) deleteWithBody(ctx context.Context, path string, body, response any) error {
	return c.do(ctx, http.MethodDelete, path, body, response)
}

// patch performs a PATCH request
func (c *Client) patch(ctx context.Context, path string, body, response any) error {
	return c.do(ctx, http.MethodPatch, path, body, response)
}
