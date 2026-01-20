package observability

import (
	"context"
	"crypto/rand"
	"fmt"
)

// IdempotencyKeyHeader is the HTTP header used for idempotency keys
const IdempotencyKeyHeader = "X-Idempotency-Key"

// idempotencyKeyContextKey is the context key for idempotency keys
type idempotencyKeyContextKey struct{}

// GenerateUUIDv4 generates a new UUID v4 string
func GenerateUUIDv4() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %w", err)
	}

	// Set version (4) and variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant 10

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:16],
	), nil
}

// MustGenerateUUIDv4 generates a new UUID v4 string, panics on error
func MustGenerateUUIDv4() string {
	uuid, err := GenerateUUIDv4()
	if err != nil {
		panic(err)
	}
	return uuid
}

// WithIdempotencyKey adds an idempotency key to the context
func WithIdempotencyKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, idempotencyKeyContextKey{}, key)
}

// WithNewIdempotencyKey generates a new UUID v4 and adds it to the context
func WithNewIdempotencyKey(ctx context.Context) (context.Context, string, error) {
	key, err := GenerateUUIDv4()
	if err != nil {
		return ctx, "", err
	}
	return WithIdempotencyKey(ctx, key), key, nil
}

// GetIdempotencyKey retrieves the idempotency key from context
func GetIdempotencyKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(idempotencyKeyContextKey{}).(string)
	return key, ok
}

// IdempotencyProvider generates and manages idempotency keys
type IdempotencyProvider struct {
	autoGenerate bool
}

// IdempotencyOption configures an IdempotencyProvider
type IdempotencyOption func(*IdempotencyProvider)

// WithAutoGenerate enables automatic idempotency key generation
func WithAutoGenerate(enabled bool) IdempotencyOption {
	return func(p *IdempotencyProvider) {
		p.autoGenerate = enabled
	}
}

// NewIdempotencyProvider creates a new IdempotencyProvider
func NewIdempotencyProvider(opts ...IdempotencyOption) *IdempotencyProvider {
	p := &IdempotencyProvider{
		autoGenerate: true, // Auto-generate by default
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// GetOrGenerate returns existing key from context or generates a new one
func (p *IdempotencyProvider) GetOrGenerate(ctx context.Context) (context.Context, string, error) {
	// Check if key already exists in context
	if key, ok := GetIdempotencyKey(ctx); ok {
		return ctx, key, nil
	}

	// Generate new key if auto-generate is enabled
	if p.autoGenerate {
		return WithNewIdempotencyKey(ctx)
	}

	return ctx, "", nil
}

// IdempotencyHook implements the Hook interface for idempotency key injection
type IdempotencyHook struct {
	provider *IdempotencyProvider
	key      string
}

// NewIdempotencyHook creates a new IdempotencyHook
func NewIdempotencyHook(provider *IdempotencyProvider) *IdempotencyHook {
	return &IdempotencyHook{
		provider: provider,
	}
}

// BeforeRequest injects idempotency key into context
func (h *IdempotencyHook) BeforeRequest(ctx context.Context, method, endpoint string, body []byte) context.Context {
	// Only add idempotency key for mutating operations
	if method == "POST" || method == "PUT" || method == "PATCH" || method == "DELETE" {
		newCtx, key, _ := h.provider.GetOrGenerate(ctx)
		h.key = key
		return newCtx
	}
	return ctx
}

// AfterResponse does nothing for idempotency
func (h *IdempotencyHook) AfterResponse(ctx context.Context, statusCode int, body []byte, err error) {
	// No-op
}

// GetKey returns the key used in the last request
func (h *IdempotencyHook) GetKey() string {
	return h.key
}
