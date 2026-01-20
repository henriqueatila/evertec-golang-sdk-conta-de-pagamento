package client

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
)

func TestParseErrorResponse(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantType   string
	}{
		{
			name:       "validation error",
			statusCode: 400,
			body:       `[{"code":"INVALID_FIELD","field":"email","message":"Invalid email format"}]`,
			wantType:   "*client.ValidationError",
		},
		{
			name:       "unauthorized error",
			statusCode: 401,
			body:       `{"message":"Invalid API key"}`,
			wantType:   "*client.UnauthorizedError",
		},
		{
			name:       "forbidden error",
			statusCode: 403,
			body:       `{"message":"Access denied to resource"}`,
			wantType:   "*client.ForbiddenError",
		},
		{
			name:       "not found error",
			statusCode: 404,
			body:       `{"message":"Account not found","resource":"account"}`,
			wantType:   "*client.NotFoundError",
		},
		{
			name:       "business rule error",
			statusCode: 409,
			body:       `{"code":"DUPLICATE_ACCOUNT","message":"Account already exists"}`,
			wantType:   "*client.BusinessRuleError",
		},
		{
			name:       "unprocessable entity error",
			statusCode: 422,
			body:       `{"code":"INVALID_STATE","message":"Cannot process in current state"}`,
			wantType:   "*client.UnprocessableEntityError",
		},
		{
			name:       "exception error",
			statusCode: 500,
			body:       `{"message":"Internal server error"}`,
			wantType:   "*client.ExceptionError",
		},
		{
			name:       "integration error",
			statusCode: 503,
			body:       `{"message":"Service temporarily unavailable"}`,
			wantType:   "*client.IntegrationError",
		},
		{
			name:       "generic API error",
			statusCode: 418,
			body:       `{"message":"I'm a teapot"}`,
			wantType:   "*client.APIError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       io.NopCloser(bytes.NewBufferString(tt.body)),
			}

			err := parseErrorResponse(resp, []byte(tt.body))
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			// Check error type by comparing error string format
			switch tt.wantType {
			case "*client.ValidationError":
				if _, ok := err.(*ValidationError); !ok {
					t.Errorf("expected ValidationError, got %T", err)
				}
			case "*client.UnauthorizedError":
				if _, ok := err.(*UnauthorizedError); !ok {
					t.Errorf("expected UnauthorizedError, got %T", err)
				}
			case "*client.ForbiddenError":
				if _, ok := err.(*ForbiddenError); !ok {
					t.Errorf("expected ForbiddenError, got %T", err)
				}
			case "*client.NotFoundError":
				if _, ok := err.(*NotFoundError); !ok {
					t.Errorf("expected NotFoundError, got %T", err)
				}
			case "*client.BusinessRuleError":
				if _, ok := err.(*BusinessRuleError); !ok {
					t.Errorf("expected BusinessRuleError, got %T", err)
				}
			case "*client.UnprocessableEntityError":
				if _, ok := err.(*UnprocessableEntityError); !ok {
					t.Errorf("expected UnprocessableEntityError, got %T", err)
				}
			case "*client.ExceptionError":
				if _, ok := err.(*ExceptionError); !ok {
					t.Errorf("expected ExceptionError, got %T", err)
				}
			case "*client.IntegrationError":
				if _, ok := err.(*IntegrationError); !ok {
					t.Errorf("expected IntegrationError, got %T", err)
				}
			case "*client.APIError":
				if _, ok := err.(*APIError); !ok {
					t.Errorf("expected APIError, got %T", err)
				}
			}
		})
	}
}

func TestValidationError(t *testing.T) {
	err := &ValidationError{
		StatusCode: 400,
		Errors: []ValidationDetail{
			{
				Code:    "INVALID_EMAIL",
				Field:   "email",
				Message: "Invalid email format",
			},
		},
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}

	// Should contain field and message
	if len(err.Errors) > 0 {
		detail := err.Errors[0]
		if detail.Field == "" || detail.Message == "" {
			t.Error("expected non-empty field and message")
		}
	}
}

func TestBusinessRuleError(t *testing.T) {
	err := &BusinessRuleError{
		StatusCode: 409,
		Code:       "DUPLICATE_ACCOUNT",
		Message:    "Account already exists",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestExceptionError(t *testing.T) {
	err := &ExceptionError{
		StatusCode: 500,
		Message:    "Internal server error",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestIntegrationError(t *testing.T) {
	err := &IntegrationError{
		StatusCode: 503,
		Message:    "Service unavailable",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestAPIError(t *testing.T) {
	err := &APIError{
		StatusCode: 404,
		Code:       "NOT_FOUND",
		Message:    "Resource not found",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestInsufficientFundsError(t *testing.T) {
	err := &InsufficientFundsError{
		StatusCode: 402,
		Code:       "INSUFFICIENT_FUNDS",
		Message:    "Insufficient funds",
		Required:   10000,
		Available:  5000,
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestMethodNotAllowedError(t *testing.T) {
	err := &MethodNotAllowedError{
		StatusCode: 405,
		Message:    "Method not allowed",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestPreconditionFailedError(t *testing.T) {
	err := &PreconditionFailedError{
		StatusCode: 412,
		Code:       "PRECONDITION_FAILED",
		Message:    "Precondition failed",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestThirdPartyError(t *testing.T) {
	err := &ThirdPartyError{
		StatusCode: 424,
		Code:       "THIRD_PARTY_ERROR",
		Message:    "Third party service failed",
		Service:    "payment-gateway",
	}

	errStr := err.Error()
	if errStr == "" {
		t.Error("expected non-empty error string")
	}
}

func TestParseErrorResponseNewTypes(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantType   string
	}{
		{
			name:       "insufficient funds error",
			statusCode: 402,
			body:       `{"code":"INSUFFICIENT_FUNDS","message":"Insufficient funds","required":10000,"available":5000}`,
			wantType:   "*client.InsufficientFundsError",
		},
		{
			name:       "method not allowed error",
			statusCode: 405,
			body:       `{"message":"Method not allowed"}`,
			wantType:   "*client.MethodNotAllowedError",
		},
		{
			name:       "precondition failed error",
			statusCode: 412,
			body:       `{"code":"PRECONDITION_FAILED","message":"Precondition failed"}`,
			wantType:   "*client.PreconditionFailedError",
		},
		{
			name:       "third party error",
			statusCode: 424,
			body:       `{"code":"THIRD_PARTY_ERROR","message":"Third party service failed","service":"payment-gateway"}`,
			wantType:   "*client.ThirdPartyError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       io.NopCloser(bytes.NewBufferString(tt.body)),
			}

			err := parseErrorResponse(resp, []byte(tt.body))
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			switch tt.wantType {
			case "*client.InsufficientFundsError":
				if _, ok := err.(*InsufficientFundsError); !ok {
					t.Errorf("expected InsufficientFundsError, got %T", err)
				}
			case "*client.MethodNotAllowedError":
				if _, ok := err.(*MethodNotAllowedError); !ok {
					t.Errorf("expected MethodNotAllowedError, got %T", err)
				}
			case "*client.PreconditionFailedError":
				if _, ok := err.(*PreconditionFailedError); !ok {
					t.Errorf("expected PreconditionFailedError, got %T", err)
				}
			case "*client.ThirdPartyError":
				if _, ok := err.(*ThirdPartyError); !ok {
					t.Errorf("expected ThirdPartyError, got %T", err)
				}
			}
		})
	}
}

// TestErrorsIs verifies that errors.Is() works correctly with sentinel errors
func TestErrorsIs(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		sentinel error
		want     bool
	}{
		{"APIError matches ErrAPI", &APIError{Message: "test"}, ErrAPI, true},
		{"ValidationError matches ErrValidation", &ValidationError{Errors: []ValidationDetail{}}, ErrValidation, true},
		{"BusinessRuleError matches ErrBusinessRule", &BusinessRuleError{Message: "test"}, ErrBusinessRule, true},
		{"ExceptionError matches ErrException", &ExceptionError{Message: "test"}, ErrException, true},
		{"IntegrationError matches ErrIntegration", &IntegrationError{Message: "test"}, ErrIntegration, true},
		{"UnauthorizedError matches ErrUnauthorized", &UnauthorizedError{Message: "test"}, ErrUnauthorized, true},
		{"ForbiddenError matches ErrForbidden", &ForbiddenError{Message: "test"}, ErrForbidden, true},
		{"NotFoundError matches ErrNotFound", &NotFoundError{Message: "test"}, ErrNotFound, true},
		{"UnprocessableEntityError matches ErrUnprocessable", &UnprocessableEntityError{Message: "test"}, ErrUnprocessable, true},
		{"InsufficientFundsError matches ErrInsufficientFunds", &InsufficientFundsError{Message: "test"}, ErrInsufficientFunds, true},
		{"MethodNotAllowedError matches ErrMethodNotAllowed", &MethodNotAllowedError{Message: "test"}, ErrMethodNotAllowed, true},
		{"PreconditionFailedError matches ErrPreconditionFailed", &PreconditionFailedError{Message: "test"}, ErrPreconditionFailed, true},
		{"ThirdPartyError matches ErrThirdParty", &ThirdPartyError{Message: "test"}, ErrThirdParty, true},
		{"PanicError matches ErrPanic", &PanicError{Message: "test"}, ErrPanic, true},
		{"APIError does not match ErrNotFound", &APIError{Message: "test"}, ErrNotFound, false},
		{"NotFoundError does not match ErrAPI", &NotFoundError{Message: "test"}, ErrAPI, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errors.Is(tt.err, tt.sentinel); got != tt.want {
				t.Errorf("errors.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestErrorsAs verifies that errors.As() works correctly for type assertion
func TestErrorsAs(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantMsg string
	}{
		{"APIError", &APIError{StatusCode: 400, Message: "bad request"}, "bad request"},
		{"NotFoundError", &NotFoundError{StatusCode: 404, Message: "not found"}, "not found"},
		{"ValidationError", &ValidationError{StatusCode: 400, Errors: []ValidationDetail{{Message: "invalid"}}}, "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "APIError":
				var target *APIError
				if !errors.As(tt.err, &target) {
					t.Error("errors.As() failed")
				}
				if target.Message != tt.wantMsg {
					t.Errorf("Message = %q, want %q", target.Message, tt.wantMsg)
				}
			case "NotFoundError":
				var target *NotFoundError
				if !errors.As(tt.err, &target) {
					t.Error("errors.As() failed")
				}
				if target.Message != tt.wantMsg {
					t.Errorf("Message = %q, want %q", target.Message, tt.wantMsg)
				}
			case "ValidationError":
				var target *ValidationError
				if !errors.As(tt.err, &target) {
					t.Error("errors.As() failed")
				}
				if len(target.Errors) > 0 && target.Errors[0].Message != tt.wantMsg {
					t.Errorf("Message = %q, want %q", target.Errors[0].Message, tt.wantMsg)
				}
			}
		})
	}
}

// TestUnwrapReturnsCorrectSentinel verifies Unwrap() returns the correct sentinel error
func TestUnwrapReturnsCorrectSentinel(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		sentinel error
	}{
		{"APIError", &APIError{Message: "test"}, ErrAPI},
		{"ValidationError", &ValidationError{}, ErrValidation},
		{"BusinessRuleError", &BusinessRuleError{}, ErrBusinessRule},
		{"ExceptionError", &ExceptionError{}, ErrException},
		{"IntegrationError", &IntegrationError{}, ErrIntegration},
		{"UnauthorizedError", &UnauthorizedError{}, ErrUnauthorized},
		{"ForbiddenError", &ForbiddenError{}, ErrForbidden},
		{"NotFoundError", &NotFoundError{}, ErrNotFound},
		{"UnprocessableEntityError", &UnprocessableEntityError{}, ErrUnprocessable},
		{"InsufficientFundsError", &InsufficientFundsError{}, ErrInsufficientFunds},
		{"MethodNotAllowedError", &MethodNotAllowedError{}, ErrMethodNotAllowed},
		{"PreconditionFailedError", &PreconditionFailedError{}, ErrPreconditionFailed},
		{"ThirdPartyError", &ThirdPartyError{}, ErrThirdParty},
		{"PanicError", &PanicError{}, ErrPanic},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unwrapper, ok := tt.err.(interface{ Unwrap() error })
			if !ok {
				t.Fatal("error does not implement Unwrap()")
			}
			if got := unwrapper.Unwrap(); got != tt.sentinel {
				t.Errorf("Unwrap() = %v, want %v", got, tt.sentinel)
			}
		})
	}
}

// TestPanicErrorFields verifies PanicError captures message and stack
func TestPanicErrorFields(t *testing.T) {
	err := &PanicError{
		Message: "nil pointer dereference",
		Stack:   "goroutine 1 [running]:\nmain.main()\n\t/app/main.go:10",
	}

	if err.Error() != "panic recovered: nil pointer dereference" {
		t.Errorf("Error() = %q, want %q", err.Error(), "panic recovered: nil pointer dereference")
	}

	if err.Stack == "" {
		t.Error("Stack should not be empty")
	}

	if !errors.Is(err, ErrPanic) {
		t.Error("PanicError should match ErrPanic with errors.Is()")
	}
}
