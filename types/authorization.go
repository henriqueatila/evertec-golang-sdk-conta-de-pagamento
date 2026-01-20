package types

import "time"

// Asset represents an asset used in a transaction
type Asset struct {
	Amount    int64     `json:"amount"` // Amount in cents
	AssetType AssetType `json:"asset_type"`
}

// SummaryPurchaseRequest represents a simplified card purchase request (Autorizador API)
type SummaryPurchaseRequest struct {
	AccountID       string            `json:"account_id"`
	AuthorizationID string            `json:"authorization_id"`
	Status          TransactionStatus `json:"status"`       // APPROVED or REJECTED
	TotalAmount     int64             `json:"total_amount"` // Amount in cents
	PaymentType     PaymentType       `json:"payment_type"`
	Assets          []Asset           `json:"assets"`

	// Optional fields
	MerchantName     *string    `json:"merchant_name,omitempty"`
	MerchantCategory *string    `json:"merchant_category,omitempty"`
	MerchantCity     *string    `json:"merchant_city,omitempty"`
	MerchantCountry  *string    `json:"merchant_country,omitempty"`
	TransactionDate  *time.Time `json:"transaction_date,omitempty"`
}

// CancelPurchaseRequest represents a request to cancel a purchase (Autorizador API)
type CancelPurchaseRequest struct {
	AccountID               string `json:"account_id"`
	OriginalAuthorizationID string `json:"original_authorization_id"`
	AuthorizationID         string `json:"authorization_id"` // New authorization ID for cancellation
}

// ChargebackRequest represents a chargeback operation (Autorizador API)
type ChargebackRequest struct {
	AccountID               string         `json:"account_id"`
	OriginalAuthorizationID string         `json:"original_authorization_id"`
	ChargebackID            string         `json:"chargeback_id"`
	ChargebackMode          ChargebackMode `json:"chargeback_mode"`             // "total" or "partial"
	ChargebackAmount        *int64         `json:"chargeback_amount,omitempty"` // Amount in cents (for partial)
	ChargebackReason        string         `json:"chargeback_reason"`
}

// CancelChargebackRequest represents a request to cancel a chargeback (Autorizador API)
type CancelChargebackRequest struct {
	AccountID            string `json:"account_id"`
	OriginalChargebackID string `json:"original_chargeback_id"`
	ChargebackID         string `json:"chargeback_id"` // New chargeback ID for cancellation
}

// AuthorizationResponse represents the response from authorization operations
type AuthorizationResponse struct {
	Message            string  `json:"message"`
	TransactionID      *int64  `json:"transactionId,omitempty"`
	AuthenticationCode *string `json:"authenticationCode,omitempty"`
}
