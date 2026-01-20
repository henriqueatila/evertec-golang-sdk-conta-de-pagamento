package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

// Handler handles incoming webhook events
type Handler struct {
	logger               *slog.Logger
	idempotencyStore     IdempotencyStore
	onPixMovement        func(*PixMovementEvent) error
	onScheduledPix       func(*ScheduledPixEvent) error
	onPrecautionaryBlock func(*PrecautionaryBlockEvent) error
	onRetainedValue      func(*RetainedValueEvent) error
	onAutomaticPix       func(*AutomaticPixEvent) error
	onClaimNotification  func(*ClaimNotificationEvent) error
}

// NewHandler creates a new webhook handler
func NewHandler(opts ...HandlerOption) *Handler {
	h := &Handler{
		logger: slog.Default(),
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// HandlerOption is a functional option for configuring the handler
type HandlerOption func(*Handler)

// WithLogger sets the logger for the handler
func WithLogger(logger *slog.Logger) HandlerOption {
	return func(h *Handler) {
		h.logger = logger
	}
}

// OnPixMovement sets the handler for PIX movement events
func OnPixMovement(fn func(*PixMovementEvent) error) HandlerOption {
	return func(h *Handler) {
		h.onPixMovement = fn
	}
}

// OnScheduledPix sets the handler for scheduled PIX events
func OnScheduledPix(fn func(*ScheduledPixEvent) error) HandlerOption {
	return func(h *Handler) {
		h.onScheduledPix = fn
	}
}

// OnPrecautionaryBlock sets the handler for precautionary block events
func OnPrecautionaryBlock(fn func(*PrecautionaryBlockEvent) error) HandlerOption {
	return func(h *Handler) {
		h.onPrecautionaryBlock = fn
	}
}

// OnRetainedValue sets the handler for retained value events
func OnRetainedValue(fn func(*RetainedValueEvent) error) HandlerOption {
	return func(h *Handler) {
		h.onRetainedValue = fn
	}
}

// OnAutomaticPix sets the handler for automatic PIX events
func OnAutomaticPix(fn func(*AutomaticPixEvent) error) HandlerOption {
	return func(h *Handler) {
		h.onAutomaticPix = fn
	}
}

// OnClaimNotification sets the handler for PIX key claim notification events
func OnClaimNotification(fn func(*ClaimNotificationEvent) error) HandlerOption {
	return func(h *Handler) {
		h.onClaimNotification = fn
	}
}

// WithIdempotencyStore sets the idempotency store for duplicate event detection.
// Use InMemoryIdempotencyStore for single-instance or implement IdempotencyStore
// interface for Redis/database backends in distributed deployments.
func WithIdempotencyStore(store IdempotencyStore) HandlerOption {
	return func(h *Handler) {
		h.idempotencyStore = store
	}
}

// HandlePixMovement handles PIX movement webhook requests
func (h *Handler) HandlePixMovement(w http.ResponseWriter, r *http.Request) {
	h.handleEvent(w, r, EventTypePixMovement)
}

// HandleScheduledPix handles scheduled PIX webhook requests
func (h *Handler) HandleScheduledPix(w http.ResponseWriter, r *http.Request) {
	h.handleEvent(w, r, EventTypeScheduledPix)
}

// HandlePrecautionaryBlock handles precautionary block webhook requests
func (h *Handler) HandlePrecautionaryBlock(w http.ResponseWriter, r *http.Request) {
	h.handleEvent(w, r, EventTypePrecautionaryBlock)
}

// HandleRetainedValue handles retained value webhook requests
func (h *Handler) HandleRetainedValue(w http.ResponseWriter, r *http.Request) {
	h.handleEvent(w, r, EventTypeRetainedValue)
}

// HandleAutomaticPix handles automatic PIX webhook requests
func (h *Handler) HandleAutomaticPix(w http.ResponseWriter, r *http.Request) {
	h.handleEvent(w, r, EventTypeAutomaticPix)
}

// HandleClaimNotification handles PIX key claim notification webhook requests
func (h *Handler) HandleClaimNotification(w http.ResponseWriter, r *http.Request) {
	h.handleEvent(w, r, EventTypeClaimNotification)
}

func (h *Handler) handleEvent(w http.ResponseWriter, r *http.Request, eventType EventType) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("Failed to read request body", "error", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	h.logger.Debug("Received webhook", "type", eventType, "body", string(body))

	var handlerErr error
	switch eventType {
	case EventTypePixMovement:
		handlerErr = h.processPixMovement(body)
	case EventTypeScheduledPix:
		handlerErr = h.processScheduledPix(body)
	case EventTypePrecautionaryBlock:
		handlerErr = h.processPrecautionaryBlock(body)
	case EventTypeRetainedValue:
		handlerErr = h.processRetainedValue(body)
	case EventTypeAutomaticPix:
		handlerErr = h.processAutomaticPix(body)
	case EventTypeClaimNotification:
		handlerErr = h.processClaimNotification(body)
	default:
		http.Error(w, "Unknown event type", http.StatusBadRequest)
		return
	}

	if handlerErr != nil {
		h.logger.Error("Handler error", "type", eventType, "error", handlerErr)
		http.Error(w, handlerErr.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) processPixMovement(body []byte) error {
	if h.onPixMovement == nil {
		return nil
	}
	var event PixMovementEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return fmt.Errorf("failed to parse PIX movement event: %w", err)
	}
	return h.onPixMovement(&event)
}

func (h *Handler) processScheduledPix(body []byte) error {
	if h.onScheduledPix == nil {
		return nil
	}
	var event ScheduledPixEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return fmt.Errorf("failed to parse scheduled PIX event: %w", err)
	}
	return h.onScheduledPix(&event)
}

func (h *Handler) processPrecautionaryBlock(body []byte) error {
	if h.onPrecautionaryBlock == nil {
		return nil
	}
	var event PrecautionaryBlockEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return fmt.Errorf("failed to parse precautionary block event: %w", err)
	}
	return h.onPrecautionaryBlock(&event)
}

func (h *Handler) processRetainedValue(body []byte) error {
	if h.onRetainedValue == nil {
		return nil
	}
	var event RetainedValueEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return fmt.Errorf("failed to parse retained value event: %w", err)
	}
	return h.onRetainedValue(&event)
}

func (h *Handler) processAutomaticPix(body []byte) error {
	if h.onAutomaticPix == nil {
		return nil
	}
	var event AutomaticPixEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return fmt.Errorf("failed to parse automatic PIX event: %w", err)
	}
	return h.onAutomaticPix(&event)
}

func (h *Handler) processClaimNotification(body []byte) error {
	if h.onClaimNotification == nil {
		return nil
	}
	var event ClaimNotificationEvent
	if err := json.Unmarshal(body, &event); err != nil {
		return fmt.Errorf("failed to parse claim notification event: %w", err)
	}
	return h.onClaimNotification(&event)
}
