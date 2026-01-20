package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewHandler(t *testing.T) {
	handler := NewHandler()
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
	if handler.logger == nil {
		t.Error("expected non-nil logger")
	}
}

func TestHandlerOptions(t *testing.T) {
	var pixCalled bool
	var scheduledCalled bool
	var blockCalled bool
	var retainedCalled bool
	var automaticCalled bool
	var claimCalled bool

	handler := NewHandler(
		OnPixMovement(func(e *PixMovementEvent) error {
			pixCalled = true
			return nil
		}),
		OnScheduledPix(func(e *ScheduledPixEvent) error {
			scheduledCalled = true
			return nil
		}),
		OnPrecautionaryBlock(func(e *PrecautionaryBlockEvent) error {
			blockCalled = true
			return nil
		}),
		OnRetainedValue(func(e *RetainedValueEvent) error {
			retainedCalled = true
			return nil
		}),
		OnAutomaticPix(func(e *AutomaticPixEvent) error {
			automaticCalled = true
			return nil
		}),
		OnClaimNotification(func(e *ClaimNotificationEvent) error {
			claimCalled = true
			return nil
		}),
	)

	if handler.onPixMovement == nil {
		t.Error("expected onPixMovement to be set")
	}
	if handler.onScheduledPix == nil {
		t.Error("expected onScheduledPix to be set")
	}
	if handler.onPrecautionaryBlock == nil {
		t.Error("expected onPrecautionaryBlock to be set")
	}
	if handler.onRetainedValue == nil {
		t.Error("expected onRetainedValue to be set")
	}
	if handler.onAutomaticPix == nil {
		t.Error("expected onAutomaticPix to be set")
	}
	if handler.onClaimNotification == nil {
		t.Error("expected onClaimNotification to be set")
	}

	// Test that handlers are callable
	_ = handler.onPixMovement(&PixMovementEvent{})
	_ = handler.onScheduledPix(&ScheduledPixEvent{})
	_ = handler.onPrecautionaryBlock(&PrecautionaryBlockEvent{})
	_ = handler.onRetainedValue(&RetainedValueEvent{})
	_ = handler.onAutomaticPix(&AutomaticPixEvent{})
	_ = handler.onClaimNotification(&ClaimNotificationEvent{})

	if !pixCalled || !scheduledCalled || !blockCalled || !retainedCalled || !automaticCalled || !claimCalled {
		t.Error("expected all handlers to be called")
	}
}

func TestHandlePixMovement(t *testing.T) {
	var receivedEvent *PixMovementEvent

	handler := NewHandler(
		OnPixMovement(func(e *PixMovementEvent) error {
			receivedEvent = e
			return nil
		}),
	)

	event := PixMovementEvent{
		AccountID:       12345,
		AuthorizationID: 67890,
		Value:           100.50,
		MovementType:    PixMovementTypeManual,
		EndToEnd:        "E123456789",
	}

	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/pix-movement", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.HandlePixMovement(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if receivedEvent == nil {
		t.Fatal("expected event to be received")
	}
	if receivedEvent.AccountID != event.AccountID {
		t.Errorf("expected AccountID %d, got %d", event.AccountID, receivedEvent.AccountID)
	}
}

func TestHandleScheduledPix(t *testing.T) {
	var receivedEvent *ScheduledPixEvent

	handler := NewHandler(
		OnScheduledPix(func(e *ScheduledPixEvent) error {
			receivedEvent = e
			return nil
		}),
	)

	event := ScheduledPixEvent{
		AccountID: 12345,
		Success:   true,
		Value:     50.00,
		EndToEnd:  "E987654321",
	}

	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/scheduled-pix", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.HandleScheduledPix(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if receivedEvent == nil {
		t.Fatal("expected event to be received")
	}
}

func TestHandlePrecautionaryBlock(t *testing.T) {
	var receivedEvent *PrecautionaryBlockEvent

	handler := NewHandler(
		OnPrecautionaryBlock(func(e *PrecautionaryBlockEvent) error {
			receivedEvent = e
			return nil
		}),
	)

	event := PrecautionaryBlockEvent{
		AccountID:                  12345,
		Type:                       PrecautionaryBlockTypeBlock,
		Value:                      1000.00,
		PrecautionaryTransactionID: 11111,
	}

	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/precautionary-block", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.HandlePrecautionaryBlock(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if receivedEvent == nil {
		t.Fatal("expected event to be received")
	}
}

func TestHandleRetainedValue(t *testing.T) {
	var receivedEvent *RetainedValueEvent

	handler := NewHandler(
		OnRetainedValue(func(e *RetainedValueEvent) error {
			receivedEvent = e
			return nil
		}),
	)

	event := RetainedValueEvent{
		AccountID:           12345,
		Value:               500.00,
		OriginTransactionID: 99999,
	}

	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/retained-value", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.HandleRetainedValue(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if receivedEvent == nil {
		t.Fatal("expected event to be received")
	}
}

func TestHandleAutomaticPix(t *testing.T) {
	var receivedEvent *AutomaticPixEvent

	handler := NewHandler(
		OnAutomaticPix(func(e *AutomaticPixEvent) error {
			receivedEvent = e
			return nil
		}),
	)

	event := AutomaticPixEvent{
		AccountID:        12345,
		NotificationType: AutomaticPixAdesao,
		Amount:           200.00,
	}

	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/automatic-pix", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.HandleAutomaticPix(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if receivedEvent == nil {
		t.Fatal("expected event to be received")
	}
}

func TestHandleClaimNotification(t *testing.T) {
	var receivedEvent *ClaimNotificationEvent

	handler := NewHandler(
		OnClaimNotification(func(e *ClaimNotificationEvent) error {
			receivedEvent = e
			return nil
		}),
	)

	event := ClaimNotificationEvent{
		AccountID:   12345,
		ClaimID:     "claim-123",
		ClaimType:   ClaimTypePortability,
		ClaimStatus: ClaimStatusOpen,
		PixKey:      "user@email.com",
		KeyType:     "EMAIL",
	}

	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/claim-notification", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.HandleClaimNotification(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
	if receivedEvent == nil {
		t.Fatal("expected event to be received")
	}
}

func TestHandlerMethodNotAllowed(t *testing.T) {
	handler := NewHandler()

	req := httptest.NewRequest(http.MethodGet, "/webhook/pix-movement", nil)
	rec := httptest.NewRecorder()

	handler.HandlePixMovement(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", rec.Code)
	}
}

func TestHandlerNoCallback(t *testing.T) {
	handler := NewHandler() // No callbacks set

	event := PixMovementEvent{AccountID: 12345}
	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/pix-movement", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.HandlePixMovement(rec, req)

	// Should still return OK when no handler is set
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

// Idempotency Store Tests

func TestInMemoryIdempotencyStore(t *testing.T) {
	store := NewInMemoryIdempotencyStore(1 * time.Hour)

	// Test Store and Exists
	if err := store.Store("event-1"); err != nil {
		t.Errorf("unexpected error storing event: %v", err)
	}

	exists, err := store.Exists("event-1")
	if err != nil {
		t.Errorf("unexpected error checking existence: %v", err)
	}
	if !exists {
		t.Error("expected event-1 to exist")
	}

	exists, err = store.Exists("event-2")
	if err != nil {
		t.Errorf("unexpected error checking existence: %v", err)
	}
	if exists {
		t.Error("expected event-2 to not exist")
	}
}

func TestInMemoryIdempotencyStoreExpiration(t *testing.T) {
	store := NewInMemoryIdempotencyStore(1 * time.Millisecond)

	// Store an event
	if err := store.Store("event-expire"); err != nil {
		t.Errorf("unexpected error storing event: %v", err)
	}

	// Wait for expiration
	time.Sleep(5 * time.Millisecond)

	// Should be expired now
	exists, err := store.Exists("event-expire")
	if err != nil {
		t.Errorf("unexpected error checking existence: %v", err)
	}
	if exists {
		t.Error("expected event to be expired")
	}
}

func TestInMemoryIdempotencyStoreCleanup(t *testing.T) {
	store := NewInMemoryIdempotencyStore(1 * time.Millisecond)

	// Store multiple events
	_ = store.Store("event-1")
	_ = store.Store("event-2")
	_ = store.Store("event-3")

	// Wait for expiration
	time.Sleep(5 * time.Millisecond)

	// Cleanup should remove expired entries
	store.Cleanup()

	// All should be cleaned up
	if len(store.seen) != 0 {
		t.Errorf("expected 0 entries after cleanup, got %d", len(store.seen))
	}
}

func TestWithIdempotencyStore(t *testing.T) {
	store := NewInMemoryIdempotencyStore(1 * time.Hour)
	handler := NewHandler(
		WithIdempotencyStore(store),
	)

	if handler.idempotencyStore == nil {
		t.Error("expected idempotencyStore to be set")
	}
}

// Event Type Tests

func TestEventTypes(t *testing.T) {
	tests := []struct {
		eventType EventType
		expected  string
	}{
		{EventTypePixMovement, "movimento_pix"},
		{EventTypeScheduledPix, "pix_agendado_executado"},
		{EventTypePrecautionaryBlock, "notifica_bloqueio_cautelar"},
		{EventTypeRetainedValue, "valor_retido"},
		{EventTypeAutomaticPix, "notificacao_pix_automatico"},
		{EventTypeClaimNotification, "notifica_reivindicacao"},
	}

	for _, tt := range tests {
		if string(tt.eventType) != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, string(tt.eventType))
		}
	}
}

func TestPixMovementTypes(t *testing.T) {
	tests := []struct {
		movementType PixMovementType
		expected     string
	}{
		{PixMovementTypeManual, "MANUAL"},
		{PixMovementTypeQRCode, "QR_CODE"},
		{PixMovementTypeKey, "CHAVE"},
		{PixMovementTypeRefund, "DEVOLUCAO"},
		{PixMovementTypeInitPay, "INIC_PAG"},
	}

	for _, tt := range tests {
		if string(tt.movementType) != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, string(tt.movementType))
		}
	}
}

func TestClaimTypes(t *testing.T) {
	if string(ClaimTypePortability) != "PORTABILITY" {
		t.Errorf("expected PORTABILITY, got %s", string(ClaimTypePortability))
	}
	if string(ClaimTypeOwnership) != "OWNERSHIP" {
		t.Errorf("expected OWNERSHIP, got %s", string(ClaimTypeOwnership))
	}
}

func TestClaimStatuses(t *testing.T) {
	tests := []struct {
		status   ClaimStatus
		expected string
	}{
		{ClaimStatusOpen, "OPEN"},
		{ClaimStatusConfirmed, "CONFIRMED"},
		{ClaimStatusCancelled, "CANCELLED"},
		{ClaimStatusCompleted, "COMPLETED"},
	}

	for _, tt := range tests {
		if string(tt.status) != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, string(tt.status))
		}
	}
}

// Panic Recovery Tests

func TestWithPanicRecovery(t *testing.T) {
	handler := NewHandler(WithPanicRecovery())
	if !handler.panicRecovery {
		t.Error("expected panicRecovery to be true")
	}
}

func TestPanicRecoveryEnabled(t *testing.T) {
	handler := NewHandler(
		WithPanicRecovery(),
		OnPixMovement(func(e *PixMovementEvent) error {
			panic("test panic")
		}),
	)

	event := PixMovementEvent{AccountID: 12345}
	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/pix-movement", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	// Should not panic due to recovery
	handler.HandlePixMovement(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rec.Code)
	}

	// Check response body contains error
	var resp map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Errorf("failed to parse response: %v", err)
	}
	if resp["error"] != "internal server error" {
		t.Errorf("expected 'internal server error', got '%s'", resp["error"])
	}
}

func TestPanicRecoveryDisabled(t *testing.T) {
	handler := NewHandler(
		// No WithPanicRecovery() - panic recovery disabled
		OnPixMovement(func(e *PixMovementEvent) error {
			panic("test panic")
		}),
	)

	event := PixMovementEvent{AccountID: 12345}
	body, _ := json.Marshal(event)
	req := httptest.NewRequest(http.MethodPost, "/webhook/pix-movement", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	// Should panic since recovery is disabled
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when recovery is disabled")
		}
	}()

	handler.HandlePixMovement(rec, req)
}
