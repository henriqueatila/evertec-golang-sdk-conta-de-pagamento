package webhook

import "time"

// EventType represents the type of webhook event
type EventType string

const (
	EventTypePixMovement        EventType = "movimento_pix"
	EventTypeScheduledPix       EventType = "pix_agendado_executado"
	EventTypePrecautionaryBlock EventType = "notifica_bloqueio_cautelar"
	EventTypeRetainedValue      EventType = "valor_retido"
	EventTypeAutomaticPix       EventType = "notificacao_pix_automatico"
	EventTypeClaimNotification  EventType = "notifica_reivindicacao"
)

// PixMovementType represents the type of PIX movement
type PixMovementType string

const (
	PixMovementTypeManual  PixMovementType = "MANUAL"
	PixMovementTypeQRCode  PixMovementType = "QR_CODE"
	PixMovementTypeKey     PixMovementType = "CHAVE"
	PixMovementTypeRefund  PixMovementType = "DEVOLUCAO"
	PixMovementTypeInitPay PixMovementType = "INIC_PAG"
)

// QRCodeType represents the QR code type
type QRCodeType string

const (
	QRCodeTypeStatic  QRCodeType = "ESTATICO"
	QRCodeTypeDynamic QRCodeType = "DINAMICO"
)

// PrecautionaryBlockType represents the type of precautionary block
type PrecautionaryBlockType string

const (
	PrecautionaryBlockTypeBlock   PrecautionaryBlockType = "BLOCK"
	PrecautionaryBlockTypeUnblock PrecautionaryBlockType = "UNBLOCK"
)

// AutomaticPixNotificationType represents automatic PIX notification type
type AutomaticPixNotificationType string

const (
	AutomaticPixAdesao                         AutomaticPixNotificationType = "ADESAO_PIX_AUTOMATICO"
	AutomaticPixCancelamento                   AutomaticPixNotificationType = "CANCELAMENTO_PIX_AUTOMATICO"
	AutomaticPixCancelamentoCobranca           AutomaticPixNotificationType = "CANCELAMENTO_COBRANCA_PIX_AUTOMATICO"
	AutomaticPixFluxoCompleto                  AutomaticPixNotificationType = "FLUXO_COMPLETO_PIX_AUTOMATICO"
	AutomaticPixFalhaAgendamentoSemRetentativa AutomaticPixNotificationType = "FALHA_AGENDAMENTO_PAGAMENTO_AUTOMATICO_SEM_RETENTATIVA"
)

// BankAccount represents bank account information in webhook payloads
type BankAccount struct {
	Name        string `json:"name,omitempty"`
	Document    string `json:"document,omitempty"`
	Bank        string `json:"bank,omitempty"`
	Branch      string `json:"branch,omitempty"`
	Account     string `json:"account,omitempty"`
	AccountType string `json:"accountType,omitempty"`
}

// PixMovementEvent represents a PIX movement webhook event
type PixMovementEvent struct {
	AccountID        int64           `json:"accountId"`
	AuthorizationID  int64           `json:"authorizationId"`
	Value            float64         `json:"value"`
	BlockedValue     float64         `json:"blockedValue,omitempty"`
	MovementType     PixMovementType `json:"movementType"`
	EndToEnd         string          `json:"endToEnd"`
	EndToEndOriginal string          `json:"endToEndOriginal,omitempty"`
	PixKey           string          `json:"pixKey,omitempty"`
	QRCodeType       QRCodeType      `json:"qrCodeType,omitempty"`
	Payer            *BankAccount    `json:"payer,omitempty"`
	Recipient        *BankAccount    `json:"recipient,omitempty"`
}

// ScheduledPixEvent represents a scheduled PIX execution webhook event
type ScheduledPixEvent struct {
	AccountID     int64   `json:"accountId"`
	Success       bool    `json:"success"`
	Value         float64 `json:"value"`
	EndToEnd      string  `json:"endToEnd,omitempty"`
	TransactionID int64   `json:"transactionId,omitempty"`
	Branch        string  `json:"branch,omitempty"`
	Account       string  `json:"account,omitempty"`
	Document      string  `json:"document,omitempty"`
}

// PrecautionaryBlockEvent represents a precautionary block webhook event
type PrecautionaryBlockEvent struct {
	AccountID                  int64                  `json:"accountId"`
	Type                       PrecautionaryBlockType `json:"type"`
	Value                      float64                `json:"value"`
	PrecautionaryTransactionID int64                  `json:"precautionaryTransactionId"`
}

// RetainedValueEvent represents a retained value webhook event
type RetainedValueEvent struct {
	AccountID           int64   `json:"accountId"`
	Value               float64 `json:"value"`
	OriginTransactionID int64   `json:"originTransactionId"`
}

// AutomaticPixEvent represents an automatic PIX notification webhook event
type AutomaticPixEvent struct {
	AccountID        int64                        `json:"accountId"`
	NotificationType AutomaticPixNotificationType `json:"notificationType"`
	Amount           float64                      `json:"amount,omitempty"`
	ContractNumber   string                       `json:"contractNumber,omitempty"`
	ReceiverName     string                       `json:"receiverName,omitempty"`
	StartDate        *time.Time                   `json:"startDate,omitempty"`
	RecurrenceID     string                       `json:"recurrenceId,omitempty"`
	ErrorCode        string                       `json:"errorCode,omitempty"`
}

// ClaimType represents the type of PIX key claim
type ClaimType string

const (
	ClaimTypePortability ClaimType = "PORTABILITY"
	ClaimTypeOwnership   ClaimType = "OWNERSHIP"
)

// ClaimStatus represents the status of a PIX key claim
type ClaimStatus string

const (
	ClaimStatusOpen      ClaimStatus = "OPEN"
	ClaimStatusConfirmed ClaimStatus = "CONFIRMED"
	ClaimStatusCancelled ClaimStatus = "CANCELLED"
	ClaimStatusCompleted ClaimStatus = "COMPLETED"
)

// ClaimNotificationEvent represents a PIX key claim notification webhook event
type ClaimNotificationEvent struct {
	AccountID   int64       `json:"accountId"`
	ClaimID     string      `json:"claimId"`
	ClaimType   ClaimType   `json:"claimType"`
	ClaimStatus ClaimStatus `json:"claimStatus"`
	PixKey      string      `json:"pixKey"`
	KeyType     string      `json:"keyType"`
	Document    string      `json:"document,omitempty"`
}

// Event represents a generic webhook event
type Event struct {
	Type EventType `json:"type"`
	Data any       `json:"data"`
}

// IdempotencyStore is the interface for webhook idempotency storage.
// Implement this interface to use Redis, database, or other backends.
type IdempotencyStore interface {
	// Exists checks if an event ID has already been processed
	Exists(eventID string) (bool, error)
	// Store marks an event ID as processed
	Store(eventID string) error
}

// InMemoryIdempotencyStore is a simple in-memory implementation of IdempotencyStore.
// Suitable for single-instance deployments. For distributed systems, use Redis or database.
type InMemoryIdempotencyStore struct {
	seen map[string]time.Time
	ttl  time.Duration
}

// NewInMemoryIdempotencyStore creates a new in-memory idempotency store
func NewInMemoryIdempotencyStore(ttl time.Duration) *InMemoryIdempotencyStore {
	return &InMemoryIdempotencyStore{
		seen: make(map[string]time.Time),
		ttl:  ttl,
	}
}

// Exists checks if an event ID has already been processed
func (s *InMemoryIdempotencyStore) Exists(eventID string) (bool, error) {
	if ts, ok := s.seen[eventID]; ok {
		if time.Since(ts) < s.ttl {
			return true, nil
		}
		delete(s.seen, eventID)
	}
	return false, nil
}

// Store marks an event ID as processed
func (s *InMemoryIdempotencyStore) Store(eventID string) error {
	s.seen[eventID] = time.Now()
	return nil
}

// Cleanup removes expired entries from the store
func (s *InMemoryIdempotencyStore) Cleanup() {
	now := time.Now()
	for id, ts := range s.seen {
		if now.Sub(ts) >= s.ttl {
			delete(s.seen, id)
		}
	}
}
