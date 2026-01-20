package observability

import (
	"context"
	"log/slog"
	"os"
	"time"
)

// Logger provides structured logging for SDK operations using slog
type Logger struct {
	logger *slog.Logger
	level  slog.Level
}

// LoggerOption configures a Logger
type LoggerOption func(*Logger)

// WithLevel sets the log level
func WithLevel(level slog.Level) LoggerOption {
	return func(l *Logger) {
		l.level = level
	}
}

// WithHandler sets a custom slog handler
func WithHandler(handler slog.Handler) LoggerOption {
	return func(l *Logger) {
		l.logger = slog.New(handler)
	}
}

// NewLogger creates a new Logger with default configuration
func NewLogger(opts ...LoggerOption) *Logger {
	l := &Logger{
		level: slog.LevelInfo,
	}

	for _, opt := range opts {
		opt(l)
	}

	if l.logger == nil {
		l.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: l.level,
		}))
	}

	return l
}

// NewLoggerWithSlog creates a Logger from an existing slog.Logger
func NewLoggerWithSlog(logger *slog.Logger) *Logger {
	return &Logger{
		logger: logger,
		level:  slog.LevelInfo,
	}
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs an info message
func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs an error message
func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// With returns a new Logger with additional attributes
func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		logger: l.logger.With(args...),
		level:  l.level,
	}
}

// LogRequest logs an HTTP request
func (l *Logger) LogRequest(ctx context.Context, method, endpoint string, body []byte) {
	l.logger.LogAttrs(ctx, slog.LevelDebug, "SDK request",
		slog.String("method", method),
		slog.String("endpoint", endpoint),
		slog.Int("body_size", len(body)),
	)
}

// LogResponse logs an HTTP response
func (l *Logger) LogResponse(ctx context.Context, method, endpoint string, statusCode int, duration time.Duration, err error) {
	level := slog.LevelInfo
	if err != nil || statusCode >= 400 {
		level = slog.LevelError
	}

	attrs := []slog.Attr{
		slog.String("method", method),
		slog.String("endpoint", endpoint),
		slog.Int("status_code", statusCode),
		slog.Duration("duration", duration),
	}

	if err != nil {
		attrs = append(attrs, slog.String("error", err.Error()))
	}

	l.logger.LogAttrs(ctx, level, "SDK response", attrs...)
}

// LogRetry logs a retry attempt
func (l *Logger) LogRetry(ctx context.Context, method, endpoint string, attempt int, err error) {
	l.logger.LogAttrs(ctx, slog.LevelWarn, "SDK retry",
		slog.String("method", method),
		slog.String("endpoint", endpoint),
		slog.Int("attempt", attempt),
		slog.String("error", err.Error()),
	)
}

// LogRateLimit logs a rate limit event
func (l *Logger) LogRateLimit(ctx context.Context, method, endpoint string, retryAfter time.Duration) {
	l.logger.LogAttrs(ctx, slog.LevelWarn, "SDK rate limited",
		slog.String("method", method),
		slog.String("endpoint", endpoint),
		slog.Duration("retry_after", retryAfter),
	)
}

// LoggerHook implements the Hook interface for logging
type LoggerHook struct {
	logger    *Logger
	startTime time.Time
	method    string
	endpoint  string
}

// NewLoggerHook creates a new LoggerHook
func NewLoggerHook(logger *Logger) *LoggerHook {
	return &LoggerHook{
		logger: logger,
	}
}

// BeforeRequest logs the request
func (h *LoggerHook) BeforeRequest(ctx context.Context, method, endpoint string, body []byte) context.Context {
	h.startTime = time.Now()
	h.method = method
	h.endpoint = endpoint
	h.logger.LogRequest(ctx, method, endpoint, body)
	return ctx
}

// AfterResponse logs the response
func (h *LoggerHook) AfterResponse(ctx context.Context, statusCode int, body []byte, err error) {
	duration := time.Since(h.startTime)
	h.logger.LogResponse(ctx, h.method, h.endpoint, statusCode, duration, err)
}
