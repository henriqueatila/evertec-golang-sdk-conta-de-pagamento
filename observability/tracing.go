// Package observability provides OpenTelemetry tracing, metrics, and logging for the SDK.
package observability

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	// TracerName is the name of the tracer used by this SDK
	TracerName = "github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento"
	// TracerVersion is the version of the tracer
	TracerVersion = "1.0.0"
)

// Tracer provides OpenTelemetry tracing for SDK operations
type Tracer struct {
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
}

// NewTracer creates a new Tracer with default OpenTelemetry provider
func NewTracer() *Tracer {
	return &Tracer{
		tracer:     otel.Tracer(TracerName, trace.WithInstrumentationVersion(TracerVersion)),
		propagator: otel.GetTextMapPropagator(),
	}
}

// NewTracerWithProvider creates a new Tracer with a custom TracerProvider
func NewTracerWithProvider(tp trace.TracerProvider) *Tracer {
	return &Tracer{
		tracer:     tp.Tracer(TracerName, trace.WithInstrumentationVersion(TracerVersion)),
		propagator: otel.GetTextMapPropagator(),
	}
}

// StartSpan starts a new span for an SDK operation
func (t *Tracer) StartSpan(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, name, opts...)
}

// StartHTTPSpan starts a new span for an HTTP request with standard HTTP attributes
func (t *Tracer) StartHTTPSpan(ctx context.Context, method, path string) (context.Context, trace.Span) {
	ctx, span := t.tracer.Start(ctx, method+" "+path,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			attribute.String("http.method", method),
			attribute.String("http.url", path),
			attribute.String("rpc.system", "http"),
			attribute.String("rpc.service", "evertec-conta-pagamento"),
		),
	)
	return ctx, span
}

// EndSpan ends a span with optional error
func (t *Tracer) EndSpan(span trace.Span, err error) {
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "")
	}
	span.End()
}

// RecordHTTPResponse records HTTP response attributes on a span
func (t *Tracer) RecordHTTPResponse(span trace.Span, statusCode int, duration time.Duration) {
	span.SetAttributes(
		attribute.Int("http.status_code", statusCode),
		attribute.Int64("http.response_time_ms", duration.Milliseconds()),
	)
}

// InjectHTTPHeaders injects trace context into HTTP headers for propagation
func (t *Tracer) InjectHTTPHeaders(ctx context.Context, req *http.Request) {
	t.propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))
}

// ExtractHTTPHeaders extracts trace context from HTTP headers
func (t *Tracer) ExtractHTTPHeaders(ctx context.Context, req *http.Request) context.Context {
	return t.propagator.Extract(ctx, propagation.HeaderCarrier(req.Header))
}

// SpanAttributes provides common attributes for SDK spans
type SpanAttributes struct {
	Operation   string
	AccountID   *int64
	Endpoint    string
	Method      string
	RequestSize int
}

// ToOtelAttributes converts SpanAttributes to OpenTelemetry attributes
func (a SpanAttributes) ToOtelAttributes() []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		attribute.String("sdk.operation", a.Operation),
		attribute.String("http.method", a.Method),
		attribute.String("http.url", a.Endpoint),
	}
	if a.AccountID != nil {
		attrs = append(attrs, attribute.Int64("account.id", *a.AccountID))
	}
	if a.RequestSize > 0 {
		attrs = append(attrs, attribute.Int("http.request_content_length", a.RequestSize))
	}
	return attrs
}
