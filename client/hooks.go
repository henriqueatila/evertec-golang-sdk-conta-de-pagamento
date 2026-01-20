package client

import (
	"context"
	"time"
)

// Hook provides observability into HTTP client operations.
// Implement this interface to add custom logging, metrics, tracing, etc.
type Hook interface {
	// BeforeRequest is called before making an HTTP request
	BeforeRequest(ctx context.Context, method, path string, body any)

	// AfterResponse is called after receiving an HTTP response
	AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error)
}

// NoOpHook is a hook that does nothing - useful for testing
type NoOpHook struct{}

func (h *NoOpHook) BeforeRequest(ctx context.Context, method, path string, body any) {}

func (h *NoOpHook) AfterResponse(ctx context.Context, method, path string, statusCode int, duration time.Duration, err error) {
}
