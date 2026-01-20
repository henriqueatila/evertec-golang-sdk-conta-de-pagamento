package types

// Recipient types based on OpenAPI spec: openapi-cadastral.json

// CreateRecipientRequest represents CreateRecipient from API spec
// POST /accounts/{accountId}/recipients
type CreateRecipientRequest struct {
	BankCode              int32  `json:"bankCode"`              // int32
	BankBranchCode        int32  `json:"bankBranchCode"`        // int32
	AccountBank           int64  `json:"accountBank"`           // int64
	AccountType           int32  `json:"accountType"`           // int32
	RecipientDocumentType int32  `json:"recipientDocumentType"` // int32
	RecipientDocument     string `json:"recipientDocument"`
	RecipientName         string `json:"recipientName"`
	Internal              int32  `json:"internal"`              // int32
}

// UpdateRecipientRequest represents UpdateRecipient from API spec
// PUT /accounts/{accountId}/recipients
type UpdateRecipientRequest struct {
	ID                    int64  `json:"id"`                    // int64 (required)
	BankCode              int32  `json:"bankCode"`              // int32
	BankBranchCode        int32  `json:"bankBranchCode"`        // int32
	AccountBank           int64  `json:"accountBank"`           // int64
	AccountType           int32  `json:"accountType"`           // int32
	RecipientDocumentType int32  `json:"recipientDocumentType"` // int32
	RecipientDocument     string `json:"recipientDocument"`
	RecipientName         string `json:"recipientName"`
	Internal              int32  `json:"internal"`              // int32
}

// RecipientDTO represents RecipientDTO from API spec
type RecipientDTO struct {
	BankCode              int32   `json:"bankCode"`              // int32
	BankBranchCode        int32   `json:"bankBranchCode"`        // int32
	AccountBank           int64   `json:"accountBank"`           // int64
	AccountType           int32   `json:"accountType"`           // int32
	RecipientDocumentType int32   `json:"recipientDocumentType"` // int32
	RecipientDocument     string  `json:"recipientDocument"`
	RecipientName         string  `json:"recipientName"`
	Internal              int32   `json:"internal"`              // int32
	ID                    int64   `json:"id"`                    // int64
	CreatedDate           *string `json:"createdDate,omitempty"` // date-time
	BankName              *string `json:"bankName,omitempty"`
}

// GetRecipientsResponse represents GetRecipientsResponse from API spec
// GET /accounts/{accountId}/recipients
type GetRecipientsResponse struct {
	Message    string         `json:"message"`
	Recipients []RecipientDTO `json:"recipients"`
	AccountID  int64          `json:"accountId"`
	DaCode     int32          `json:"da_code"` // int32
}

// GetRecipientResponse represents GetRecipientResponse from API spec
// GET /accounts/{accountId}/recipients/{recipientId}
type GetRecipientResponse struct {
	Message   string        `json:"message"`
	Recipient *RecipientDTO `json:"recipient,omitempty"`
	DaCode    int32         `json:"da_code"` // int32
}
