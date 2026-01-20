package types

import "fmt"

// ValidationError represents a single field validation error (HTTP 400)
type ValidationError struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors represents an array of validation errors (HTTP 400)
type ValidationErrors []ValidationError

// Error implements the error interface
func (v ValidationErrors) Error() string {
	if len(v) == 0 {
		return "validation error"
	}
	return fmt.Sprintf("validation failed: %s", v[0].Message)
}

// BusinessError represents a business rule violation (HTTP 409)
type BusinessError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface
func (b BusinessError) Error() string {
	return fmt.Sprintf("[%s] %s", b.Code, b.Message)
}

// APIError represents a server error (HTTP 500/503)
type APIError struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

// Error implements the error interface
func (a APIError) Error() string {
	if a.Code != "" {
		return fmt.Sprintf("[%d] [%s] %s", a.StatusCode, a.Code, a.Message)
	}
	return fmt.Sprintf("[%d] %s", a.StatusCode, a.Message)
}

// IntegrationError represents an external service error (HTTP 503)
type IntegrationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface
func (i IntegrationError) Error() string {
	return fmt.Sprintf("integration error: [%s] %s", i.Code, i.Message)
}

// Common error codes as constants
const (
	// Account errors
	ErrAccountNotFound       = "account.not.found"
	ErrAccountNotActive      = "account.not.active"
	ErrProposalAlreadyExists = "already.exists.proposal"

	// Transaction errors
	ErrInsufficientFunds       = "insufficient.funds"
	ErrTransactionCodeNotFound = "transaction.code.not.found"
	ErrSchedulingInvalidDate   = "scheduling.not.in.valid.date"

	// PIX errors
	ErrPixLimitExceed          = "limit.pix.exceed"
	ErrDuplicatedEndToEnd      = "duplicated.end.to.end"
	ErrPixKeyAlreadyRegistered = "pix.key.already.registered"
	ErrPixDeviceNotFound       = "pix.device.not.found"
	ErrFraudDetected           = "fraud.detected"

	// QR Code errors
	ErrQRCodeInvalidFormat = "qrcode.invalid.format"
	ErrQRCodeExpired       = "qr.code.expired"

	// Data validation errors
	ErrEmailInvalid     = "email.invalid"
	ErrCellphoneInvalid = "cellphone.invalid"
	ErrDocumentInvalid  = "document.invalid"
	ErrInvalidField     = "invalid.field"

	// Integration errors
	ErrIntegrationError = "integration.error"
	ErrUnknownError     = "unknow.error"
)
