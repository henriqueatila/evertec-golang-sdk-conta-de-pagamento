package types

// Webhook types based on OpenAPI spec: openapi-cadastral.json

// EventoEmailDTO represents EventoEmailDTO from API spec
// POST /webhook/sendgrid/update
type EventoEmailDTO struct {
	Email       *string  `json:"email,omitempty"`
	Timestamp   *int64   `json:"timestamp,omitempty"`
	Event       *string  `json:"event,omitempty"`
	Category    []string `json:"category,omitempty"`
	Response    *string  `json:"response,omitempty"`
	Attempt     *int     `json:"attempt,omitempty"`
	UserAgent   *string  `json:"useragent,omitempty"`
	IP          *string  `json:"ip,omitempty"`
	URL         *string  `json:"url,omitempty"`
	Reason      *string  `json:"reason,omitempty"`
	Status      *string  `json:"status,omitempty"`
	SmtpID      *string  `json:"smtp-id,omitempty"`
	SgEventID   *string  `json:"sg_event_id,omitempty"`
	SgMessageID *string  `json:"sg_message_id,omitempty"`
	EmailID     *string  `json:"email_id,omitempty"`
}

// EventHubRequest represents EventHubRequest from API spec
// POST /postpaid/notification/eventhub/*
type EventHubRequest struct {
	AccountID *string `json:"accountId,omitempty"`
	Title     *string `json:"title,omitempty"`
}

// ReivindicacaoRequest represents ReivindicacaoRequest from API spec (claim object)
type ReivindicacaoRequest struct {
	ClaimID     *string `json:"claimId,omitempty"`
	ClaimStatus *string `json:"claimStatus,omitempty"`
	ClaimType   *string `json:"claimType,omitempty"`
}

// NotificationPushRequest represents NotificationPushRequest from API spec
// POST /webhook/arbi/notifiesUserArbiOperation
type NotificationPushRequest struct {
	Document         *string               `json:"document,omitempty"`
	Body             *string               `json:"body,omitempty"`
	TransactionID    *int64                `json:"transactionId,omitempty"`
	Title            *string               `json:"title,omitempty"`
	ArbiWebhookEnum  *string               `json:"arbiWebhookEnum,omitempty"` // Enum: CLAIM_WAITING_RESOLUTION
	ProcessStartTime *string               `json:"processStartTime,omitempty"`
	ClaimObject      *ReivindicacaoRequest `json:"claimObject,omitempty"`
}
