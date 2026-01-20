package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Sentinel errors for error type checking with errors.Is()
var (
	ErrAPI               = errors.New("api error")
	ErrValidation        = errors.New("validation error")
	ErrBusinessRule      = errors.New("business rule error")
	ErrException         = errors.New("server error")
	ErrIntegration       = errors.New("integration error")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrNotFound          = errors.New("not found")
	ErrUnprocessable     = errors.New("unprocessable entity")
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrMethodNotAllowed  = errors.New("method not allowed")
	ErrPreconditionFailed = errors.New("precondition failed")
	ErrThirdParty        = errors.New("third party error")
	ErrPanic             = errors.New("panic recovered")
)

// APIError represents an error response from the Evertec API
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
	Err        error  `json:"-"` // Wrapped error for error chaining
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error [%d]: %s (code: %s)", e.StatusCode, e.Message, e.Code)
}

func (e *APIError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrAPI
}

// ValidationError represents a 400 Bad Request response with field-level validation errors
type ValidationError struct {
	StatusCode int                `json:"statusCode"`
	Errors     []ValidationDetail `json:"errors"`
	Err        error              `json:"-"`
}

type ValidationDetail struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	if len(e.Errors) == 0 {
		return "validation error"
	}
	return fmt.Sprintf("validation error: %s (field: %s)", e.Errors[0].Message, e.Errors[0].Field)
}

func (e *ValidationError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrValidation
}

// BusinessRuleError represents a 409 Conflict response
type BusinessRuleError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *BusinessRuleError) Error() string {
	return fmt.Sprintf("business rule error [%d]: %s (code: %s)", e.StatusCode, e.Message, e.Code)
}

func (e *BusinessRuleError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrBusinessRule
}

// ExceptionError represents a 500 Internal Server Error response
type ExceptionError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *ExceptionError) Error() string {
	return fmt.Sprintf("server error [%d]: %s", e.StatusCode, e.Message)
}

func (e *ExceptionError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrException
}

// IntegrationError represents a 503 Service Unavailable response
type IntegrationError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *IntegrationError) Error() string {
	return fmt.Sprintf("integration error [%d]: %s", e.StatusCode, e.Message)
}

func (e *IntegrationError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrIntegration
}

// UnauthorizedError represents a 401 Unauthorized response
type UnauthorizedError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("unauthorized [%d]: %s", e.StatusCode, e.Message)
}

func (e *UnauthorizedError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrUnauthorized
}

// ForbiddenError represents a 403 Forbidden response
type ForbiddenError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *ForbiddenError) Error() string {
	return fmt.Sprintf("forbidden [%d]: %s", e.StatusCode, e.Message)
}

func (e *ForbiddenError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrForbidden
}

// NotFoundError represents a 404 Not Found response
type NotFoundError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Resource   string `json:"resource,omitempty"`
	Err        error  `json:"-"`
}

func (e *NotFoundError) Error() string {
	if e.Resource != "" {
		return fmt.Sprintf("not found [%d]: %s (resource: %s)", e.StatusCode, e.Message, e.Resource)
	}
	return fmt.Sprintf("not found [%d]: %s", e.StatusCode, e.Message)
}

func (e *NotFoundError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrNotFound
}

// UnprocessableEntityError represents a 422 Unprocessable Entity response
type UnprocessableEntityError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *UnprocessableEntityError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("unprocessable entity [%d]: %s (code: %s)", e.StatusCode, e.Message, e.Code)
	}
	return fmt.Sprintf("unprocessable entity [%d]: %s", e.StatusCode, e.Message)
}

func (e *UnprocessableEntityError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrUnprocessable
}

// InsufficientFundsError represents a 402 Payment Required response
type InsufficientFundsError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	Required   int64  `json:"required,omitempty"`
	Available  int64  `json:"available,omitempty"`
	Err        error  `json:"-"`
}

func (e *InsufficientFundsError) Error() string {
	if e.Required > 0 && e.Available >= 0 {
		return fmt.Sprintf("insufficient funds [%d]: %s (required: %d, available: %d)", e.StatusCode, e.Message, e.Required, e.Available)
	}
	return fmt.Sprintf("insufficient funds [%d]: %s", e.StatusCode, e.Message)
}

func (e *InsufficientFundsError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrInsufficientFunds
}

// MethodNotAllowedError represents a 405 Method Not Allowed response
type MethodNotAllowedError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *MethodNotAllowedError) Error() string {
	return fmt.Sprintf("method not allowed [%d]: %s", e.StatusCode, e.Message)
}

func (e *MethodNotAllowedError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrMethodNotAllowed
}

// PreconditionFailedError represents a 412 Precondition Failed response
type PreconditionFailedError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

func (e *PreconditionFailedError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("precondition failed [%d]: %s (code: %s)", e.StatusCode, e.Message, e.Code)
	}
	return fmt.Sprintf("precondition failed [%d]: %s", e.StatusCode, e.Message)
}

func (e *PreconditionFailedError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrPreconditionFailed
}

// ThirdPartyError represents a 424 Failed Dependency response (third-party service failure)
type ThirdPartyError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	Service    string `json:"service,omitempty"`
	Err        error  `json:"-"`
}

func (e *ThirdPartyError) Error() string {
	if e.Service != "" {
		return fmt.Sprintf("third party error [%d]: %s (service: %s)", e.StatusCode, e.Message, e.Service)
	}
	return fmt.Sprintf("third party error [%d]: %s", e.StatusCode, e.Message)
}

func (e *ThirdPartyError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrThirdParty
}

// PanicError represents a recovered panic during HTTP operations
type PanicError struct {
	Message string
	Stack   string
	Err     error
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic recovered: %s", e.Message)
}

func (e *PanicError) Unwrap() error {
	if e.Err != nil {
		return e.Err
	}
	return ErrPanic
}

// parseErrorResponse parses the HTTP response into an appropriate error type
func parseErrorResponse(resp *http.Response, body []byte) error {
	statusCode := resp.StatusCode

	switch statusCode {
	case http.StatusBadRequest:
		// Try parsing as validation error array
		var validationErrors []ValidationDetail
		if err := json.Unmarshal(body, &validationErrors); err == nil && len(validationErrors) > 0 {
			return &ValidationError{
				StatusCode: statusCode,
				Errors:     validationErrors,
			}
		}

		// Fallback to generic API error
		var apiErr APIError
		if err := json.Unmarshal(body, &apiErr); err == nil {
			apiErr.StatusCode = statusCode
			return &apiErr
		}

	case http.StatusPaymentRequired:
		var errResp struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			Required  int64  `json:"required"`
			Available int64  `json:"available"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil {
			return &InsufficientFundsError{
				StatusCode: statusCode,
				Code:       errResp.Code,
				Message:    errResp.Message,
				Required:   errResp.Required,
				Available:  errResp.Available,
			}
		}
		return &InsufficientFundsError{
			StatusCode: statusCode,
			Message:    "insufficient funds",
		}

	case http.StatusUnauthorized:
		var errResp struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil && errResp.Message != "" {
			return &UnauthorizedError{
				StatusCode: statusCode,
				Message:    errResp.Message,
			}
		}
		return &UnauthorizedError{
			StatusCode: statusCode,
			Message:    "authentication required",
		}

	case http.StatusForbidden:
		var errResp struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil && errResp.Message != "" {
			return &ForbiddenError{
				StatusCode: statusCode,
				Message:    errResp.Message,
			}
		}
		return &ForbiddenError{
			StatusCode: statusCode,
			Message:    "access denied",
		}

	case http.StatusNotFound:
		var errResp struct {
			Message  string `json:"message"`
			Resource string `json:"resource"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil {
			return &NotFoundError{
				StatusCode: statusCode,
				Message:    errResp.Message,
				Resource:   errResp.Resource,
			}
		}
		return &NotFoundError{
			StatusCode: statusCode,
			Message:    "resource not found",
		}

	case http.StatusMethodNotAllowed:
		var errResp struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil && errResp.Message != "" {
			return &MethodNotAllowedError{
				StatusCode: statusCode,
				Message:    errResp.Message,
			}
		}
		return &MethodNotAllowedError{
			StatusCode: statusCode,
			Message:    "method not allowed",
		}

	case http.StatusConflict:
		var businessErr struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &businessErr); err == nil {
			return &BusinessRuleError{
				StatusCode: statusCode,
				Code:       businessErr.Code,
				Message:    businessErr.Message,
			}
		}

	case http.StatusPreconditionFailed:
		var errResp struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil {
			return &PreconditionFailedError{
				StatusCode: statusCode,
				Code:       errResp.Code,
				Message:    errResp.Message,
			}
		}
		return &PreconditionFailedError{
			StatusCode: statusCode,
			Message:    "precondition failed",
		}

	case http.StatusUnprocessableEntity:
		var errResp struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil {
			return &UnprocessableEntityError{
				StatusCode: statusCode,
				Code:       errResp.Code,
				Message:    errResp.Message,
			}
		}
		return &UnprocessableEntityError{
			StatusCode: statusCode,
			Message:    string(body),
		}

	case http.StatusFailedDependency:
		var errResp struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Service string `json:"service"`
		}
		if err := json.Unmarshal(body, &errResp); err == nil {
			return &ThirdPartyError{
				StatusCode: statusCode,
				Code:       errResp.Code,
				Message:    errResp.Message,
				Service:    errResp.Service,
			}
		}
		return &ThirdPartyError{
			StatusCode: statusCode,
			Message:    "third party service failure",
		}

	case http.StatusInternalServerError:
		var exceptionErr struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &exceptionErr); err == nil {
			return &ExceptionError{
				StatusCode: statusCode,
				Message:    exceptionErr.Message,
			}
		}

	case http.StatusServiceUnavailable:
		var integrationErr struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &integrationErr); err == nil {
			return &IntegrationError{
				StatusCode: statusCode,
				Message:    integrationErr.Message,
			}
		}
	}

	// Generic error for unhandled cases
	return &APIError{
		StatusCode: statusCode,
		Message:    string(body),
	}
}
