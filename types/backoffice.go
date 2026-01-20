package types

import (
	"fmt"
	"net/url"
)

// Backoffice Account Operations

// ListAccountsBackofficeRequest represents backoffice account list request
type ListAccountsBackofficeRequest struct {
	Document *string `json:"document,omitempty"`
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Status   *string `json:"status,omitempty"`
	Page     *int    `json:"page,omitempty"`
	PageSize *int    `json:"pageSize,omitempty"`
}

// ProposalStatus enum values for UpdateProposalProcessingRequest
type ProposalProcessingStatus string

const (
	ProposalProcessingStatusWaitingAutomatic    ProposalProcessingStatus = "WAITING_AUTOMATIC_ANALYSIS"
	ProposalProcessingStatusWaitingManual       ProposalProcessingStatus = "WAITING_MANUAL_ANALYSIS"
	ProposalProcessingStatusManualApproved      ProposalProcessingStatus = "MANUAL_APPROVED"
	ProposalProcessingStatusAutomaticApproved   ProposalProcessingStatus = "AUTOMATICALLY_APPROVED"
	ProposalProcessingStatusManualReproved      ProposalProcessingStatus = "MANUAL_REPROVED"
	ProposalProcessingStatusAutomaticReproved   ProposalProcessingStatus = "AUTOMATICALLY_REPROVED"
)

// UpdateProposalProcessingRequest represents proposal processing request (API: UpdateProposalProcessingRequest)
type UpdateProposalProcessingRequest struct {
	ProposalID        int64                                     `json:"proposalId"`        // Required
	ProposalStatus    ProposalProcessingStatus                  `json:"proposalStatus"`    // Required (enum)
	MaintenanceUserID string                                    `json:"maintenanceUserId"` // Required
	Reason            *string                                   `json:"reason,omitempty"`
	ArrangementType   *string                                   `json:"arrangementType,omitempty"`
	CreditLimit       *CreateAccountCreditLimitPostPaidRequest  `json:"creditLimit,omitempty"`
	Limits            []ProposalLimitDto                        `json:"limits,omitempty"`
	ProductID         *int32                                    `json:"productId,omitempty"`
}

// ProposalProcessingRequest is an alias for backward compatibility
// Deprecated: Use UpdateProposalProcessingRequest instead
type ProposalProcessingRequest = UpdateProposalProcessingRequest

// CreateAccountCreditLimitPostPaidRequest represents post-paid credit limit request (API: CreateAccountCreditLimitPostPaidRequest)
type CreateAccountCreditLimitPostPaidRequest struct {
	LimitAmount         float64 `json:"limitAmount"`         // Required (min: 0)
	WithdrawLimitAmount float64 `json:"withdrawLimitAmount"` // Required (min: 0)
	PaymentDue          int32   `json:"paymentDue"`          // Required (range: 1-28)
}

// ProposalLimitDto represents proposal limit configuration (API: ProposalLimitDto)
type ProposalLimitDto struct {
	EnvironmentType       string   `json:"environmentType"`                 // Required (pattern: ^[IES]$)
	TransactionCode       int32    `json:"transactionCode"`                 // Required (min: 0)
	DayLimit              *float64 `json:"dayLimit,omitempty"`              // Optional (min: 0)
	DayTransactionLimit   *float64 `json:"dayTransactionLimit,omitempty"`   // Optional (min: 0)
	NightLimit            *float64 `json:"nightLimit,omitempty"`            // Optional (min: 0)
	NightTransactionLimit *float64 `json:"nightTransactionLimit,omitempty"` // Optional (min: 0)
	TransactionAmount     *int32   `json:"transactionAmount,omitempty"`     // Optional (min: 0)
}

// CreateMobileAccountRequest represents mobile account creation request
type CreateMobileAccountRequest struct {
	Document string  `json:"document"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	DeviceID *string `json:"deviceId,omitempty"`
}

// Biro Analysis (Credit Bureau)

// BiroAnalysisRequest represents credit bureau analysis request
type BiroAnalysisRequest struct {
	ProposalID   int64                  `json:"proposalId"`
	AnalysisType string                 `json:"analysisType"`
	Parameters   map[string]any `json:"parameters,omitempty"`
}

// BiroAnalysisResponse represents credit bureau analysis response
type BiroAnalysisResponse struct {
	AnalysisID int64              `json:"analysisId"`
	ProposalID int64              `json:"proposalId"`
	Status     BiroAnalysisStatus `json:"status"`
	Score      *int    `json:"score,omitempty"`
	Result     *string `json:"result,omitempty"`
	Details    *string `json:"details,omitempty"`
	AnalyzedAt *string `json:"analyzedAt,omitempty"`
	ExpiresAt  *string `json:"expiresAt,omitempty"`
}

// UpdateBiroAnalysisRequest represents biro analysis update request
type UpdateBiroAnalysisRequest struct {
	Status  BiroAnalysisStatus  `json:"status"`
	Result  *string `json:"result,omitempty"`
	Details *string `json:"details,omitempty"`
}

// Processor Operations

// BindProcessorAccountRequest represents processor account binding request
type BindProcessorAccountRequest struct {
	AccountID          int64  `json:"accountId"`
	ProcessorAccountID string `json:"processorAccountId"`
}

// BindProcessorCardRequest represents processor card binding request
type BindProcessorCardRequest struct {
	AccountID       int64  `json:"accountId"`
	ProcessorCardID string `json:"processorCardId"`
	CardType        string `json:"cardType"`
}

// SyncProcessorResponse represents processor sync response
type SyncProcessorResponse struct {
	AccountID       int64           `json:"accountId"`
	ProcessorStatus ProcessorStatus `json:"processorStatus"`
	SyncedAt        string   `json:"syncedAt"`
	Changes         []string `json:"changes,omitempty"`
}

// PIX Scan Configuration

// PixScanConfigurationResponse represents PIX scan configuration
type PixScanConfigurationResponse struct {
	Enabled           bool    `json:"enabled"`
	ScanFrequency     *int    `json:"scanFrequency,omitempty"`  // in minutes
	AlertThreshold    *int64  `json:"alertThreshold,omitempty"` // amount in cents
	AutoBlockEnabled  *bool   `json:"autoBlockEnabled,omitempty"`
	NotificationEmail *string `json:"notificationEmail,omitempty"`
}

// UpdatePixScanConfigurationRequest represents PIX scan config update request
type UpdatePixScanConfigurationRequest struct {
	Enabled           *bool   `json:"enabled,omitempty"`
	ScanFrequency     *int    `json:"scanFrequency,omitempty"`
	AlertThreshold    *int64  `json:"alertThreshold,omitempty"`
	AutoBlockEnabled  *bool   `json:"autoBlockEnabled,omitempty"`
	NotificationEmail *string `json:"notificationEmail,omitempty"`
}

// HCE Devices

// HceDeviceResponse represents HCE device information
type HceDeviceResponse struct {
	DeviceID     string          `json:"deviceId"`
	AccountID    int64           `json:"accountId"`
	DeviceName   *string         `json:"deviceName,omitempty"`
	DeviceModel  *string         `json:"deviceModel,omitempty"`
	Status       HceDeviceStatus `json:"status"`
	RegisteredAt string  `json:"registeredAt"`
	LastUsedAt   *string `json:"lastUsedAt,omitempty"`
}

// ListHceDevicesParams represents HCE device list parameters
type ListHceDevicesParams struct {
	AccountID *int64  `json:"accountId,omitempty"`
	Status    *string `json:"status,omitempty"`
	Page      *int    `json:"page,omitempty"`
	PageSize  *int    `json:"pageSize,omitempty"`
}

// QueryString returns query string for HCE device list
func (p *ListHceDevicesParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.AccountID != nil {
		params.Set("accountId", fmt.Sprintf("%d", *p.AccountID))
	}
	if p.Status != nil {
		params.Set("status", *p.Status)
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%d", *p.Page))
	}
	if p.PageSize != nil {
		params.Set("pageSize", fmt.Sprintf("%d", *p.PageSize))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// Daily Statement

// DailyStatementResponse represents daily statement entry
type DailyStatementResponse struct {
	Date         string `json:"date"`
	AccountCount int    `json:"accountCount"`
	TotalBalance int64  `json:"totalBalance"` // in cents
	TotalCredits int64  `json:"totalCredits"`
	TotalDebits  int64  `json:"totalDebits"`
}

// DailyStatementListResponse represents list of daily statements
type DailyStatementListResponse struct {
	Statements []DailyStatementResponse `json:"statements"`
	Total      int                      `json:"total"`
}

// Issuer Balance

// IssuerBalanceResponse represents issuer total balance
type IssuerBalanceResponse struct {
	TotalBalance    int64  `json:"totalBalance"` // in cents
	AccountCount    int    `json:"accountCount"`
	ActiveAccounts  int    `json:"activeAccounts"`
	BlockedAccounts int    `json:"blockedAccounts"`
	CalculatedAt    string `json:"calculatedAt"`
}

// Generic Backoffice Response

// BackofficeGenericResponse represents generic backoffice response
type BackofficeGenericResponse struct {
	Success bool        `json:"success"`
	Message *string     `json:"message,omitempty"`
	Data    any `json:"data,omitempty"`
}
