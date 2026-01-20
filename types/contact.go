package types

import "time"

// ContactResponse represents a contact
type ContactResponse struct {
	ContactID int64      `json:"contactId"`
	AccountID int64      `json:"accountId"`
	Name      string     `json:"name"`
	Document  *string    `json:"document,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Phone     *string    `json:"phone,omitempty"`
	Nickname  *string    `json:"nickname,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// ContactListResponse represents a list of contacts
type ContactListResponse struct {
	Contacts []ContactResponse `json:"items"`
	Total    *int              `json:"total,omitempty"`
}

// ContactBankDetailsResponse represents bank details of a contact
type ContactBankDetailsResponse struct {
	ContactID    int64        `json:"contactId"`
	BankCode     *string      `json:"bankCode,omitempty"`
	BankName     *string      `json:"bankName,omitempty"`
	Branch       *string      `json:"branch,omitempty"`
	Account      *string      `json:"account,omitempty"`
	AccountDigit *string      `json:"accountDigit,omitempty"`
	AccountType  *string      `json:"accountType,omitempty"`
	PixKeys      []PixKeyInfo `json:"pixKeys,omitempty"`
}

// PixKeyInfo represents PIX key information for a contact
type PixKeyInfo struct {
	KeyType  PixKeyType `json:"keyType"`
	KeyValue string     `json:"keyValue"`
}

// CreditResponse represents credit information
type CreditResponse struct {
	CreditID    int64      `json:"creditId"`
	AccountID   int64      `json:"accountId"`
	Amount      int64      `json:"amount"` // Amount in cents
	Description *string    `json:"description,omitempty"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UsedAt      *time.Time `json:"usedAt,omitempty"`
	Status      CreditStatus     `json:"status"` // e.g., "AVAILABLE", "USED", "EXPIRED"
}

// CreditListResponse represents a list of credits
type CreditListResponse struct {
	Credits []CreditResponse `json:"items"`
	Total   *int             `json:"total,omitempty"`
}
