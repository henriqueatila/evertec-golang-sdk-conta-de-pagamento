package observability

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	// MeterName is the name of the meter used by this SDK
	MeterName = "github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento"
)

// Metrics provides OpenTelemetry metrics for SDK operations
type Metrics struct {
	meter metric.Meter

	// Request metrics
	requestCounter  metric.Int64Counter
	requestDuration metric.Float64Histogram
	requestSize     metric.Int64Histogram
	responseSize    metric.Int64Histogram

	// Error metrics
	errorCounter metric.Int64Counter

	// Active requests gauge
	activeRequests metric.Int64UpDownCounter
}

// NewMetrics creates a new Metrics instance with default OpenTelemetry provider
func NewMetrics() (*Metrics, error) {
	return NewMetricsWithProvider(otel.GetMeterProvider())
}

// NewMetricsWithProvider creates a new Metrics instance with a custom MeterProvider
func NewMetricsWithProvider(mp metric.MeterProvider) (*Metrics, error) {
	m := &Metrics{
		meter: mp.Meter(MeterName),
	}

	var err error

	// Request counter
	m.requestCounter, err = m.meter.Int64Counter(
		"evertec.sdk.requests.total",
		metric.WithDescription("Total number of requests made by the SDK"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		return nil, err
	}

	// Request duration histogram
	m.requestDuration, err = m.meter.Float64Histogram(
		"evertec.sdk.request.duration",
		metric.WithDescription("Duration of SDK requests in milliseconds"),
		metric.WithUnit("ms"),
		metric.WithExplicitBucketBoundaries(1, 5, 10, 25, 50, 100, 250, 500, 1000, 2500, 5000, 10000),
	)
	if err != nil {
		return nil, err
	}

	// Request size histogram
	m.requestSize, err = m.meter.Int64Histogram(
		"evertec.sdk.request.size",
		metric.WithDescription("Size of SDK requests in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}

	// Response size histogram
	m.responseSize, err = m.meter.Int64Histogram(
		"evertec.sdk.response.size",
		metric.WithDescription("Size of SDK responses in bytes"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}

	// Error counter
	m.errorCounter, err = m.meter.Int64Counter(
		"evertec.sdk.errors.total",
		metric.WithDescription("Total number of errors encountered by the SDK"),
		metric.WithUnit("{error}"),
	)
	if err != nil {
		return nil, err
	}

	// Active requests gauge
	m.activeRequests, err = m.meter.Int64UpDownCounter(
		"evertec.sdk.requests.active",
		metric.WithDescription("Number of active SDK requests"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// RecordRequest records metrics for a completed request
func (m *Metrics) RecordRequest(ctx context.Context, method, endpoint string, statusCode int, duration time.Duration, reqSize, respSize int64) {
	attrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", endpoint),
		attribute.Int("http.status_code", statusCode),
	}

	m.requestCounter.Add(ctx, 1, metric.WithAttributes(attrs...))
	m.requestDuration.Record(ctx, float64(duration.Milliseconds()), metric.WithAttributes(attrs...))

	if reqSize > 0 {
		m.requestSize.Record(ctx, reqSize, metric.WithAttributes(attrs...))
	}
	if respSize > 0 {
		m.responseSize.Record(ctx, respSize, metric.WithAttributes(attrs...))
	}
}

// RecordError records an error metric
func (m *Metrics) RecordError(ctx context.Context, method, endpoint, errorType string) {
	attrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", endpoint),
		attribute.String("error.type", errorType),
	}
	m.errorCounter.Add(ctx, 1, metric.WithAttributes(attrs...))
}

// IncrementActiveRequests increments the active requests counter
func (m *Metrics) IncrementActiveRequests(ctx context.Context) {
	m.activeRequests.Add(ctx, 1)
}

// DecrementActiveRequests decrements the active requests counter
func (m *Metrics) DecrementActiveRequests(ctx context.Context) {
	m.activeRequests.Add(ctx, -1)
}

// MetricsHook implements the Hook interface for metrics collection
type MetricsHook struct {
	metrics   *Metrics
	startTime time.Time
	method    string
	endpoint  string
	reqSize   int64
}

// NewMetricsHook creates a new MetricsHook
func NewMetricsHook(metrics *Metrics) *MetricsHook {
	return &MetricsHook{
		metrics: metrics,
	}
}

// BeforeRequest records the start time and increments active requests
func (h *MetricsHook) BeforeRequest(ctx context.Context, method, endpoint string, body []byte) context.Context {
	h.startTime = time.Now()
	h.method = method
	h.endpoint = endpoint
	h.reqSize = int64(len(body))
	h.metrics.IncrementActiveRequests(ctx)
	return ctx
}

// AfterResponse records request metrics and decrements active requests
func (h *MetricsHook) AfterResponse(ctx context.Context, statusCode int, body []byte, err error) {
	duration := time.Since(h.startTime)
	respSize := int64(len(body))

	h.metrics.RecordRequest(ctx, h.method, h.endpoint, statusCode, duration, h.reqSize, respSize)

	if err != nil {
		h.metrics.RecordError(ctx, h.method, h.endpoint, errorType(err))
	}

	h.metrics.DecrementActiveRequests(ctx)
}

// errorType extracts the error type from an error
func errorType(err error) string {
	if err == nil {
		return ""
	}
	// Try to extract specific error types
	switch err.(type) {
	default:
		return "unknown"
	}
}
