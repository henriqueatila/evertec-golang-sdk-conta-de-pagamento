package types

import "time"

// BankSlipStatus represents the status of a bank slip
type BankSlipStatus string

const (
	BankSlipStatusPending  BankSlipStatus = "PENDING"
	BankSlipStatusPaid     BankSlipStatus = "PAID"
	BankSlipStatusCanceled BankSlipStatus = "CANCELED"
	BankSlipStatusExpired  BankSlipStatus = "EXPIRED"
)

// CreateBankSlipRequest represents a request to generate a bank slip (boleto)
type CreateBankSlipRequest struct {
	Amount        int64   `json:"amount"`  // Amount in cents
	DueDate       string  `json:"dueDate"` // Format: YYYY-MM-DD
	Description   *string `json:"description,omitempty"`
	PayerName     *string `json:"payerName,omitempty"`
	PayerDocument *string `json:"payerDocument,omitempty"`

	// Fine and interest (optional)
	FinePercentage     *float64 `json:"finePercentage,omitempty"`     // e.g., 2.0 for 2%
	InterestPercentage *float64 `json:"interestPercentage,omitempty"` // Daily interest, e.g., 0.033 for 1% per month
	DiscountAmount     *int64   `json:"discountAmount,omitempty"`     // Discount in cents
	DiscountDueDate    *string  `json:"discountDueDate,omitempty"`    // Format: YYYY-MM-DD
}

// BankSlipResponse represents a bank slip (boleto)
type BankSlipResponse struct {
	BankSlipID    int64          `json:"bankSlipId"`
	AccountID     int64          `json:"accountId"`
	Amount        int64          `json:"amount"` // Amount in cents
	Status        BankSlipStatus `json:"status"`
	Barcode       string         `json:"barcode"`
	DigitableLine string         `json:"digitableLine"`
	DueDate       string         `json:"dueDate"`
	Description   *string        `json:"description,omitempty"`
	PayerName     *string        `json:"payerName,omitempty"`
	PayerDocument *string        `json:"payerDocument,omitempty"`

	// Fine and interest
	FinePercentage     *float64 `json:"finePercentage,omitempty"`
	InterestPercentage *float64 `json:"interestPercentage,omitempty"`
	DiscountAmount     *int64   `json:"discountAmount,omitempty"`
	DiscountDueDate    *string  `json:"discountDueDate,omitempty"`

	// Payment info
	PaidAmount *int64  `json:"paidAmount,omitempty"` // Amount in cents
	PaidDate   *string `json:"paidDate,omitempty"`

	// URLs
	PDFUrl    *string `json:"pdfUrl,omitempty"`
	QRCodeUrl *string `json:"qrCodeUrl,omitempty"`

	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// BankSlipListResponse represents a list of bank slips
type BankSlipListResponse struct {
	BankSlips []BankSlipResponse `json:"items"`
	Total     *int               `json:"total,omitempty"`
}

// ===============================================
// Collections V2 API Types (openapi-cobranca.json)
// ===============================================

// FeeType represents the type of fee calculation (V2 API)
type FeeType string

const (
	FeeTypeFixed      FeeType = "FIXED"
	FeeTypePercentage FeeType = "PERCENTAGE"
)

// CustomerRequestV2 represents customer information for V2 bank slip (API: CustomerRequestV2)
type CustomerRequestV2 struct {
	Document      string  `json:"document"`      // Required - CPF or CNPJ
	Name          string  `json:"name"`          // Required
	Email         string  `json:"email"`         // Required
	Address       string  `json:"address"`       // Required - Street address
	AddressNumber string  `json:"adressNumber"`  // Required (note: API typo preserved)
	Complement    *string `json:"complement,omitempty"`
	Neighborhood  string  `json:"neighborhood"`  // Required
	ZipCode       string  `json:"zipCode"`       // Required
	City          string  `json:"city"`          // Required
	State         string  `json:"state"`         // Required (2-letter UF code)
}

// DiscountRequestV2 represents discount configuration for V2 bank slip (API: DiscountRequestV2)
type DiscountRequestV2 struct {
	Amount           float64 `json:"amount"`           // Required - Discount amount
	DueDateLimitDays int     `json:"dueDateLimitDays"` // Required - Days before due date
	Type             FeeType `json:"type"`             // FIXED or PERCENTAGE
}

// FineRequestV2 represents fine configuration for V2 bank slip (API: FineRequestV2)
type FineRequestV2 struct {
	Amount float64 `json:"amount"` // Required - Fine amount
	Type   FeeType `json:"type"`   // FIXED or PERCENTAGE
}

// InterestRequestV2 represents interest configuration for V2 bank slip (API: InterestRequestV2)
type InterestRequestV2 struct {
	Amount float64 `json:"amount"` // Required - Daily interest amount
}

// CreateBankSlipRequestV2 represents the V2 bank slip creation request (API: CreateBankSlipRequestV2)
type CreateBankSlipRequestV2 struct {
	AccountID   int64              `json:"accountId"`   // Required
	Amount      float64            `json:"amount"`      // Required
	DueDate     string             `json:"dueDate"`     // Required (date format)
	Description string             `json:"description"` // Required
	Customer    CustomerRequestV2  `json:"customer"`    // Required
	Discount    *DiscountRequestV2 `json:"discount,omitempty"`
	Fine        *FineRequestV2     `json:"fine,omitempty"`
	Interest    *InterestRequestV2 `json:"interest,omitempty"`
}

// CreateBankSlipResponseV2 represents the V2 bank slip creation response (API: CreateBankSlipResponseV2)
type CreateBankSlipResponseV2 struct {
	IDBankslip          int64   `json:"idBankslip"`
	AccountID           int64   `json:"accountId"`
	CreatedDate         string  `json:"createdDate"`
	Amount              float64 `json:"amount"`
	DueDate             string  `json:"dueDate"`
	DigitableLine       string  `json:"digitableLine"`
	Description         string  `json:"description"`
	PaymentStatus       PaymentStatus  `json:"paymentStatus"`
	BankslipDownloadURL *string `json:"bankslipDownloadUrl,omitempty"`
}
