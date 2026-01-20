package types

import "time"

// DepositOrderStatus represents the status of a deposit order
type DepositOrderStatus string

const (
	DepositOrderStatusPending   DepositOrderStatus = "PENDING"
	DepositOrderStatusActive    DepositOrderStatus = "ACTIVE"
	DepositOrderStatusCompleted DepositOrderStatus = "COMPLETED"
	DepositOrderStatusCanceled  DepositOrderStatus = "CANCELED"
	DepositOrderStatusExpired   DepositOrderStatus = "EXPIRED"
)

// CreateDepositOrderRequest represents a request to create a deposit order
type CreateDepositOrderRequest struct {
	Amount      int64   `json:"amount"` // Amount in cents
	Description *string `json:"description,omitempty"`
	ExpiryHours *int    `json:"expiryHours,omitempty"` // Hours until expiration
}

// DepositOrderResponse represents a deposit order
type DepositOrderResponse struct {
	DepositOrderID int64              `json:"depositOrderId"`
	AccountID      int64              `json:"accountId"`
	Amount         int64              `json:"amount"` // Amount in cents
	Status         DepositOrderStatus `json:"status"`
	Description    *string            `json:"description,omitempty"`

	// Payment instructions
	BankCode     *string `json:"bankCode,omitempty"`
	Branch       *string `json:"branch,omitempty"`
	Account      *string `json:"account,omitempty"`
	AccountDigit *string `json:"accountDigit,omitempty"`
	Barcode      *string `json:"barcode,omitempty"`
	PixKey       *string `json:"pixKey,omitempty"`
	QRCodeData   *string `json:"qrCodeData,omitempty"`

	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}

// DepositOrderListResponse represents a list of deposit orders
type DepositOrderListResponse struct {
	Orders []DepositOrderResponse `json:"items"`
	Total  *int                   `json:"total,omitempty"`
}
