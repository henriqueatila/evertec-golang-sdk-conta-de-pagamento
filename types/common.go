package types

import "time"

// GenericResponse represents a standard API success response
type GenericResponse struct {
	Message string `json:"message"`
	DACode  int32  `json:"da_code"`
}

// TransactionResponse represents a standard transaction response
type TransactionResponse struct {
	Message            string     `json:"message"`
	TransactionID      *int64     `json:"transactionId,omitempty"`
	AuthenticationCode *string    `json:"authenticationCode,omitempty"`
	FreeField          *string    `json:"freeField,omitempty"`
	DateTimeTransfer   *time.Time `json:"dateTimeTransfer,omitempty"`
	Amount             *int64     `json:"amount,omitempty"`
	DACode             int32      `json:"da_code"`
}

// ListResponse represents a paginated list response
type ListResponse struct {
	Items any `json:"items"`
	Total *int        `json:"total,omitempty"`
	Page  *int        `json:"page,omitempty"`
}

// StatusResponse represents API status response
type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
