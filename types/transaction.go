package types

import "time"

// InternalTransferRequest represents an internal transfer between accounts
type InternalTransferRequest struct {
	RecipientAccountID int64   `json:"recipientAccountId"`
	TransferAmount     int64   `json:"transferAmount"` // Amount in cents
	FreeDescription    *string `json:"freeDescription,omitempty"`
	ScheduledDate      *string `json:"scheduledDate,omitempty"` // Format: YYYY-MM-DD

	// Location (for fraud prevention)
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

// IDToIDTransferRequest represents a transfer using document ID
type IDToIDTransferRequest struct {
	RecipientDocument string  `json:"recipientDocument"` // CPF or CNPJ
	TransferAmount    int64   `json:"transferAmount"`    // Amount in cents
	FreeDescription   *string `json:"freeDescription,omitempty"`

	// Location (for fraud prevention)
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

// BatchTransferRequest represents a batch transfer operation
type BatchTransferRequest struct {
	ExpiresIn string                    `json:"expiresIn"` // Required (date format)
	Transfers []InternalTransferRequest `json:"transfers"` // Named 'transactions' in spec
}

// ArrangementTransferRequest represents an arrangement transfer
type ArrangementTransferRequest struct {
	RecipientAccountID int64   `json:"recipientAccountId"`
	TransferAmount     int64   `json:"transferAmount"` // Amount in cents
	ArrangementType    string  `json:"arrangementType"`
	FreeDescription    *string `json:"freeDescription,omitempty"`
}

// CancelTransferRequest represents the request to cancel a scheduled transfer
type CancelTransferRequest struct {
	TransactionID int64   `json:"transactionId"`
	Reason        *string `json:"reason,omitempty"`
}

// CheckRecipientAccountRequest represents the request to validate recipient account
type CheckRecipientAccountRequest struct {
	RecipientAccountID int64 `json:"recipientAccountId"`
}

// CheckRecipientAccountResponse represents recipient account validation result
type CheckRecipientAccountResponse struct {
	Valid             bool    `json:"valid"`
	RecipientName     *string `json:"recipientName,omitempty"`
	RecipientDocument *string `json:"recipientDocument,omitempty"`
	AccountStatus     *string `json:"accountStatus,omitempty"`
}

// ExternalTransferRequest represents a TED/DOC transfer to external bank
type ExternalTransferRequest struct {
	// Recipient bank details
	BankCode     string  `json:"bankCode"` // Bank code (e.g., "001" for Banco do Brasil)
	Branch       string  `json:"branch"`   // Branch number
	Account      string  `json:"account"`  // Account number
	AccountDigit *string `json:"accountDigit,omitempty"`
	AccountType  string  `json:"accountType"` // e.g., "CHECKING", "SAVINGS"

	// Recipient details
	RecipientDocument string `json:"recipientDocument"` // CPF or CNPJ
	RecipientName     string `json:"recipientName"`

	// Transfer details
	TransferType string  `json:"transferType"` // "TED" or "DOC"
	Amount       int64   `json:"amount"`       // Amount in cents
	Description  *string `json:"description,omitempty"`

	// Scheduling (optional)
	ScheduledDate *string `json:"scheduledDate,omitempty"` // Format: YYYY-MM-DD

	// Location (for fraud prevention)
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

// ScheduledTransferResponse represents a scheduled transfer
type ScheduledTransferResponse struct {
	SchedulingID      int64             `json:"schedulingId"`
	TransactionID     *int64            `json:"transactionId,omitempty"`
	Type              TransactionType   `json:"type"`
	Status            TransactionStatus `json:"status"`
	Amount            int64             `json:"amount"` // Amount in cents
	ScheduledDate     string            `json:"scheduledDate"`
	RecipientName     *string           `json:"recipientName,omitempty"`
	RecipientDocument *string           `json:"recipientDocument,omitempty"`
	Description       *string           `json:"description,omitempty"`
	CreatedAt         *time.Time        `json:"createdAt,omitempty"`
}

// ScheduledTransferListResponse represents a list of scheduled transfers
type ScheduledTransferListResponse struct {
	Transfers []ScheduledTransferResponse `json:"items"`
	Total     *int                        `json:"total,omitempty"`
}

// CancelScheduledTransferRequest represents the request to cancel a scheduled transfer
type CancelScheduledTransferRequest struct {
	SchedulingID int64   `json:"schedulingId"`
	Reason       *string `json:"reason,omitempty"`
}

// BillPaymentRequest represents a bill payment request
type BillPaymentRequest struct {
	Barcode       string  `json:"barcode"`                 // Bill barcode/digitable line
	Amount        *int64  `json:"amount,omitempty"`        // Amount in cents (if different from bill)
	ScheduledDate *string `json:"scheduledDate,omitempty"` // Format: YYYY-MM-DD
	Description   *string `json:"description,omitempty"`

	// Location (for fraud prevention)
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

// BatchBillPaymentRequest represents a batch bill payment request
type BatchBillPaymentRequest struct {
	Payments []BillPaymentRequest `json:"payments"`
}

// GetBillInfoRequest represents the request to fetch bill information
type GetBillInfoRequest struct {
	Barcode string `json:"barcode"`
}

// BillInfoResponse represents bill information
type BillInfoResponse struct {
	Barcode           string  `json:"barcode"`
	Amount            int64   `json:"amount"`  // Amount in cents
	DueDate           string  `json:"dueDate"` // Format: YYYY-MM-DD
	RecipientName     string  `json:"recipientName"`
	RecipientDocument string  `json:"recipientDocument"`
	Description       *string `json:"description,omitempty"`
	Fine              *int64  `json:"fine,omitempty"`        // Fine in cents
	Interest          *int64  `json:"interest,omitempty"`    // Interest in cents
	Discount          *int64  `json:"discount,omitempty"`    // Discount in cents
	TotalAmount       *int64  `json:"totalAmount,omitempty"` // Total with fees in cents
}

// ScheduledBillPaymentResponse represents a scheduled bill payment
type ScheduledBillPaymentResponse struct {
	SchedulingID  int64             `json:"schedulingId"`
	Barcode       string            `json:"barcode"`
	Amount        int64             `json:"amount"` // Amount in cents
	ScheduledDate string            `json:"scheduledDate"`
	Status        TransactionStatus `json:"status"`
	RecipientName *string           `json:"recipientName,omitempty"`
	Description   *string           `json:"description,omitempty"`
	CreatedAt     *time.Time        `json:"createdAt,omitempty"`
}

// ScheduledBillPaymentListResponse represents a list of scheduled bill payments
type ScheduledBillPaymentListResponse struct {
	Payments []ScheduledBillPaymentResponse `json:"items"`
	Total    *int                           `json:"total,omitempty"`
}

// MobileRechargeRequest represents a mobile phone recharge request
type MobileRechargeRequest struct {
	PhoneAreaCode string  `json:"phoneAreaCode"`      // Area code (e.g., "11")
	PhoneNumber   string  `json:"phoneNumber"`        // Phone number
	Amount        int64   `json:"amount"`             // Amount in cents
	Operator      *string `json:"operator,omitempty"` // Operator name (e.g., "Vivo", "Claro")
}

// AvailableRechargeValuesRequest represents the request to get available recharge values
type AvailableRechargeValuesRequest struct {
	PhoneAreaCode string `json:"phoneAreaCode"`
	PhoneNumber   string `json:"phoneNumber"`
}

// AvailableRechargeValuesResponse represents available recharge values
type AvailableRechargeValuesResponse struct {
	Operator string  `json:"operator"`
	Values   []int64 `json:"values"` // Available amounts in cents
}

// VoucherRechargeRequest represents a voucher/electronic voucher recharge
type VoucherRechargeRequest struct {
	ProviderID    string  `json:"providerId"`
	Amount        int64   `json:"amount"`                  // Amount in cents
	RecipientInfo *string `json:"recipientInfo,omitempty"` // Phone, email, or account
}

// VoucherProviderResponse represents a voucher provider
type VoucherProviderResponse struct {
	ProviderID   string  `json:"providerId"`
	ProviderName string  `json:"providerName"`
	Category     *string `json:"category,omitempty"`
	MinAmount    *int64  `json:"minAmount,omitempty"`   // Min amount in cents
	MaxAmount    *int64  `json:"maxAmount,omitempty"`   // Max amount in cents
	FixedValues  []int64 `json:"fixedValues,omitempty"` // Fixed amounts in cents
}

// RecipientRequest represents a recipient/beneficiary creation request
type RecipientRequest struct {
	// Personal/Company information
	Name         string       `json:"name"`
	Document     string       `json:"document"`
	DocumentType DocumentType `json:"documentType"`

	// Bank details (for TED/DOC)
	BankCode     *string `json:"bankCode,omitempty"`
	Branch       *string `json:"branch,omitempty"`
	Account      *string `json:"account,omitempty"`
	AccountDigit *string `json:"accountDigit,omitempty"`
	AccountType  *string `json:"accountType,omitempty"` // "CHECKING", "SAVINGS"

	// PIX key (alternative to bank details)
	PixKeyType  *PixKeyType `json:"pixKeyType,omitempty"`
	PixKeyValue *string     `json:"pixKeyValue,omitempty"`

	// Optional
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Nickname *string `json:"nickname,omitempty"` // Friendly name
}

// RecipientResponse represents a recipient/beneficiary
type RecipientResponse struct {
	RecipientID  int64        `json:"recipientId"`
	AccountID    int64        `json:"accountId"`
	Name         string       `json:"name"`
	Document     string       `json:"document"`
	DocumentType DocumentType `json:"documentType"`
	BankCode     *string      `json:"bankCode,omitempty"`
	Branch       *string      `json:"branch,omitempty"`
	Account      *string      `json:"account,omitempty"`
	AccountDigit *string      `json:"accountDigit,omitempty"`
	AccountType  *string      `json:"accountType,omitempty"`
	PixKeyType   *PixKeyType  `json:"pixKeyType,omitempty"`
	PixKeyValue  *string      `json:"pixKeyValue,omitempty"`
	Email        *string      `json:"email,omitempty"`
	Phone        *string      `json:"phone,omitempty"`
	Nickname     *string      `json:"nickname,omitempty"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time   `json:"updatedAt,omitempty"`
}

// RecipientListResponse represents a list of recipients
type RecipientListResponse struct {
	Recipients []RecipientResponse `json:"items"`
	Total      *int                `json:"total,omitempty"`
}

// BankResponse represents a bank
type BankResponse struct {
	BankCode string  `json:"bankCode"`
	BankName string  `json:"bankName"`
	ISPB     *string `json:"ispb,omitempty"` // Brazilian payment system identifier
}

// BankListResponse represents a list of banks
type BankListResponse struct {
	Banks []BankResponse `json:"items"`
}

// BalanceLockResponse represents a balance lock entry
type BalanceLockResponse struct {
	LockID     int64             `json:"lockId"`
	AccountID  int64             `json:"accountId"`
	Amount     int64             `json:"amount"` // Amount in cents
	Reason     *string           `json:"reason,omitempty"`
	Status     BalanceLockStatus `json:"status"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty"`
	ReleasedAt *time.Time `json:"releasedAt,omitempty"`
}

// BalanceLockListResponse represents a list of balance locks
type BalanceLockListResponse struct {
	Locks []BalanceLockResponse `json:"items"`
	Total *int                  `json:"total,omitempty"`
}

// TransactionListParams represents parameters for listing transactions
type TransactionListParams struct {
	StartDate *string `json:"startDate,omitempty"` // Format: YYYY-MM-DD
	EndDate   *string `json:"endDate,omitempty"`   // Format: YYYY-MM-DD
	Page      *int    `json:"page,omitempty"`
	PageSize  *int    `json:"pageSize,omitempty"`
}

// TransactionListResponse represents a list of transactions by type
type TransactionListResponse struct {
	Transactions []StatementEntry `json:"items"`
	Total        *int             `json:"total,omitempty"`
	Page         *int             `json:"page,omitempty"`
}

// PixCallbackRequest represents a PIX callback/notification request
type PixCallbackRequest struct {
	EndToEndID        string  `json:"endToEndId"`
	TransactionID     *int64  `json:"transactionId,omitempty"`
	Amount            int64   `json:"amount"` // Amount in cents
	PayerName         *string `json:"payerName,omitempty"`
	PayerDocument     *string `json:"payerDocument,omitempty"`
	ReceiverAccountID *int64  `json:"receiverAccountId,omitempty"`
	Description       *string `json:"description,omitempty"`
	TransactionDate   string  `json:"transactionDate"`
}

// PixCallbackResponse represents the response for PIX callback processing
type PixCallbackResponse struct {
	Processed bool    `json:"processed"`
	Message   *string `json:"message,omitempty"`
}

// AccountBillPaymentRequest represents bill payment request via account endpoint
type AccountBillPaymentRequest struct {
	AccountID     int64   `json:"accountId"`
	Barcode       string  `json:"barcode"`
	Amount        *int64  `json:"amount,omitempty"` // Amount in cents
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	Description   *string `json:"description,omitempty"`
}
