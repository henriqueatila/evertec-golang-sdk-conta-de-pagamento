package types

import (
	"fmt"
	"net/url"
)

// PIX Automático (Recorrência) - Automatic/Recurring PIX
// Based on OpenAPI spec: openapi-pix-automatico.json

// FrequencyType represents recurrence frequency
type FrequencyType string

const (
	FrequencyAnnual    FrequencyType = "MIAN"
	FrequencyMonthly   FrequencyType = "MNTH"
	FrequencyQuarterly FrequencyType = "QURT"
	FrequencyWeekly    FrequencyType = "WEEK"
	FrequencyYearly    FrequencyType = "YEAR"
)

// RecurrenceStatus represents recurrence status
type RecurrenceStatus string

const (
	RecurrenceStatusPending   RecurrenceStatus = "PDNG"
	RecurrenceStatusConfirmed RecurrenceStatus = "CFDB"
	RecurrenceStatusCancelled RecurrenceStatus = "CCLD"
)

// RejectReason represents rejection reason codes
type RejectReason string

const (
	RejectReasonAP13 RejectReason = "AP13"
	RejectReasonAP14 RejectReason = "AP14"
)

// CancellationReason represents cancellation reason codes
type CancellationReason string

const (
	CancellationReasonACCL CancellationReason = "ACCL"
	CancellationReasonCPCL CancellationReason = "CPCL"
	CancellationReasonDCSD CancellationReason = "DCSD"
	CancellationReasonERSL CancellationReason = "ERSL"
	CancellationReasonFRUD CancellationReason = "FRUD"
	CancellationReasonNRES CancellationReason = "NRES"
	CancellationReasonPCFD CancellationReason = "PCFD"
	CancellationReasonSLCR CancellationReason = "SLCR"
	CancellationReasonSLDB CancellationReason = "SLDB"
)

// ChargeCancelReason represents charge cancel reason codes
type ChargeCancelReason string

const (
	ChargeCancelReasonSLBD ChargeCancelReason = "SLBD"
	ChargeCancelReasonFAIL ChargeCancelReason = "FAIL"
)

// JourneyType represents automatic PIX journey type
type JourneyType string

const (
	JourneyTypeAUT2 JourneyType = "AUT2"
	JourneyTypeAUT3 JourneyType = "AUT3"
	JourneyTypeAUT4 JourneyType = "AUT4"
)

// StartAutomaticPixRequest represents request to start automatic PIX
type StartAutomaticPixRequest struct {
	AccountID                  int64         `json:"accountId"`
	PayerBranch                *string       `json:"payerBranch,omitempty"`
	PayerAccount               string        `json:"payerAccount"`
	DebtorDocument             *string       `json:"debtorDocument,omitempty"`
	PayerDocument              *string       `json:"payerDocument,omitempty"`
	RecurrenceEndDate          *string       `json:"recurrenceEndDate,omitempty"` // Format: YYYY-MM-DD
	RecurrenceStartDate        string        `json:"recurrenceStartDate"`         // Format: YYYY-MM-DD
	RecurrenceID               string        `json:"recurrenceId"`                // Max 29 chars
	ExpirationDate             *string       `json:"expirationDate,omitempty"`    // ISO 8601 datetime
	Description                *string       `json:"description,omitempty"`
	IndicatorFloorMaximumValue *bool         `json:"indicatorFloorMaximumValue,omitempty"`
	DebtorName                 *string       `json:"debtorName,omitempty"`
	ContractNumber             string        `json:"contractNumber"`
	PayerParticipant           string        `json:"payerParticipant"`
	FloorMaximumValue          *float64      `json:"floorMaximumValue,omitempty"`
	FrequencyType              FrequencyType `json:"frequencyType"`
	Value                      *float64      `json:"value,omitempty"`
}

// StartAutomaticPixResponse represents response from starting automatic PIX
type StartAutomaticPixResponse struct {
	RecurrenceID        string `json:"recurrenceId"`
	RecurrenceRequestID string `json:"recurrenceRequestId"`
	PainID              int64  `json:"painId"`
}

// RejectAutomaticPixRequest represents request to reject automatic PIX
type RejectAutomaticPixRequest struct {
	AccountID    int64        `json:"accountId"`
	RecurrenceID string       `json:"recurrenceId"` // Max 29 chars
	RejectReason RejectReason `json:"rejectReason"` // AP13 or AP14
}

// RejectAutomaticPixResponse represents response from rejecting automatic PIX
type RejectAutomaticPixResponse struct {
	RecurrenceID        string `json:"recurrenceId"`
	StatusInformationID string `json:"statusInformationId"`
}

// PixPaymentRequest represents PIX payment request (for automatic PIX journey three)
type PixPaymentRequestAutomatic struct {
	AccountID                 int64    `json:"accountId"`
	RecipientInstitutionCode  string   `json:"recipientInstitutionCode"`
	RecipientBranchCode       string   `json:"recipientBranchCode"`
	RecipientAccountNumber    string   `json:"recipientAccountNumber"`
	RecipientAccountType      string   `json:"recipientAccountType"` // CACC, SLRY, SVGS, TRAN
	RecipientCpfCnpj          string   `json:"recipientCpfCnpj"`
	RecipientName             string   `json:"recipientName"`
	PayerName *string `json:"payerName,omitempty"`
	// InternalReference uses "internalReferente" (API typo preserved for compatibility)
	InternalReference *string `json:"internalReferente,omitempty"`
	OperationAmount           float64  `json:"operationAmount"`
	EndToEnd                  *string  `json:"endToEnd,omitempty"`
	RecipientAddressingKey    *string  `json:"recipientAddressingKey,omitempty"`
	FreeField                 *string  `json:"freeField,omitempty"`
	SchedulingDate            *string  `json:"schedulingDate,omitempty"`
	TransactionPurpose        *string  `json:"transactionPurpose,omitempty"` // TROCO, SAQUE
	WithdrawalServiceProvider *string  `json:"withdrawalServiceProvider,omitempty"`
	AgentMode                 *string  `json:"agentMode,omitempty"` // AGFSS, AGTEC, AGTOT
	CashMoney                 *float64 `json:"cashMoney,omitempty"`
	Latitude                  *string  `json:"latitude,omitempty"`
	Longitude                 *string  `json:"longitude,omitempty"`
	SaveContact               *bool    `json:"saveContact,omitempty"`
}

// QRCodeUserAcceptRequest represents QR code acceptance request
type QRCodeUserAcceptRequest struct {
	AccountID                   int64         `json:"accountId"`
	DebtorDocument              *string       `json:"debtorDocument,omitempty"`
	PayerDocument               string        `json:"payerDocument"`
	ReceiverDocument            *string       `json:"receiverDocument,omitempty"`
	RecurrenceEndDate           *string       `json:"recurrenceEndDate,omitempty"`
	RequestCreationDateTime     string        `json:"requestCreationDateTime"`     // ISO 8601
	CreationDateTimeForIssuance string        `json:"creationDateTimeForIssuance"` // ISO 8601
	RecurrenceStartDate         string        `json:"recurrenceStartDate"`         // YYYY-MM-DD
	Description                 *string       `json:"description,omitempty"`
	EndToEnd                    *string       `json:"endToEnd,omitempty"`
	RecurrenceID                string        `json:"recurrenceId"`
	DebtorName                  *string       `json:"debtorName,omitempty"`
	ReceiverName                string        `json:"receiverName"`
	ContractNumber              string        `json:"contractNumber"`
	ReceiverParticipant         string        `json:"receiverParticipant"`
	FrequencyType               FrequencyType `json:"frequencyType"`
	RecurrenceType              *string       `json:"recurrenceType,omitempty"`
	Value                       *float64      `json:"value,omitempty"`
	AuthorizedValue             *float64      `json:"authorizedValue,omitempty"`
	IbgeMunicipalCode           *string       `json:"ibgeMunicipalCode,omitempty"`
	Journey                     JourneyType   `json:"journey"` // AUT2, AUT3, AUT4
}

// QRCodeAcceptJourneyThreeRequest represents QR code journey three acceptance
type QRCodeAcceptJourneyThreeRequest struct {
	Pain012    QRCodeUserAcceptRequest    `json:"pain012"`
	PixPayment PixPaymentRequestAutomatic `json:"pixPayment"`
}

// QRCodeAcceptJourneyThreeResponse represents response from journey three acceptance
type QRCodeAcceptJourneyThreeResponse struct {
	IDTransaction       int64  `json:"idTransaction"`
	AuthenticationCode  string `json:"authenticationCode"`
	MessageID           string `json:"messageId"`
	RecurrenceID        string `json:"recurrenceId"`
	StatusInformationID string `json:"statusInformationId"`
}

// QRCodeUserResponse represents response from QR code acceptance
type QRCodeUserResponse struct {
	MessageID           string `json:"messageId"`
	RecurrenceID        string `json:"recurrenceId"`
	StatusInformationID string `json:"statusInformationId"`
}

// CreateAutomaticPixContractRequest represents automatic PIX contract creation
type CreateAutomaticPixContractRequest struct {
	AccountID      int64  `json:"accountId"`
	RecurrenceID   string `json:"recurrenceId"`
	ContractNumber string `json:"contractNumber"`
}

// CreateAutomaticPixContractResponse represents automatic PIX contract response
type CreateAutomaticPixContractResponse struct {
	RecurrenceID string `json:"recurrenceId"`
	ContractID   int64  `json:"contractId"`
}

// CancelAutomaticPixChargeRequest represents request to cancel automatic PIX charge
type CancelAutomaticPixChargeRequest struct {
	ID        int64              `json:"id"`
	AccountID int64              `json:"accountId"`
	Reason    ChargeCancelReason `json:"reason"` // SLBD or FAIL
}

// CancelAutomaticPixChargeResponse represents response from cancelling charge
type CancelAutomaticPixChargeResponse struct {
	CancellationID string `json:"cancellationId"`
}

// CancelAutomaticPixRequest represents request to cancel automatic PIX
type CancelAutomaticPixRequest struct {
	AccountID          int64              `json:"accountId"`
	RecurrenceID       string             `json:"recurrenceId"`
	CancellationReason CancellationReason `json:"cancellationReason"`
}

// CancelAutomaticPixResponse represents response from cancelling automatic PIX
type CancelAutomaticPixResponse struct {
	RecurrenceID       string `json:"recurrenceId"`
	CancellationInfoID string `json:"cancellationInfoId"`
	PainID             int64  `json:"painId"`
}

// AcceptAutomaticPixRequest represents request to accept automatic PIX
type AcceptAutomaticPixRequest struct {
	AccountID    int64    `json:"accountId"`
	RecurrenceID string   `json:"recurrenceId"`
	ZipCode      *string  `json:"zipCode,omitempty"` // 8 digits
	Value        *float64 `json:"value,omitempty"`
}

// AutomaticPixResponse represents full automatic PIX response
type AutomaticPixResponse struct {
	RecurrenceID             string           `json:"recurrenceId"`
	RecurrenceRequestID      *string          `json:"recurrenceRequestId,omitempty"`
	Status                   RecurrenceStatus `json:"status"`
	FrequencyType            FrequencyType    `json:"frequencyType,omitempty"`
	RecurrenceStartDate      *string       `json:"recurrenceStartDate,omitempty"`
	RecurrenceEndDate        *string       `json:"recurrenceEndDate,omitempty"`
	Value                    *float64      `json:"value,omitempty"`
	FloorMaximumValue        *float64      `json:"floorMaximumValue,omitempty"`
	ContractNumber           *string       `json:"contractNumber,omitempty"`
	Description              *string       `json:"description,omitempty"`
	PayerDocument            *string       `json:"payerDocument,omitempty"`
	PayerName                *string       `json:"payerName,omitempty"`
	ReceiverDocument         *string       `json:"receiverDocument,omitempty"`
	ReceiverName             *string       `json:"receiverName,omitempty"`
	ReceiverParticipant      *string       `json:"receiverParticipant,omitempty"`
	CreationDateTime         *string       `json:"creationDateTime,omitempty"`
	LastModificationDateTime *string       `json:"lastModificationDateTime,omitempty"`
}

// ListAutomaticPixParams represents query params for listing automatic PIX
type ListAutomaticPixParams struct {
	Inactive         *bool             `json:"inactive,omitempty"`
	RecurrenceID     *string           `json:"recurrenceId,omitempty"`
	Page             *int              `json:"page,omitempty"`
	Size             *int              `json:"size,omitempty"`
	RecurrenceStatus *RecurrenceStatus `json:"recurrenceStatus,omitempty"`
	IsPayer          *bool             `json:"isPayer,omitempty"`
}

// QueryString returns query string for list automatic PIX params
func (p *ListAutomaticPixParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.Inactive != nil {
		params.Set("inactive", fmt.Sprintf("%t", *p.Inactive))
	}
	if p.RecurrenceID != nil {
		params.Set("recurrenceId", *p.RecurrenceID)
	}
	if p.Page != nil {
		params.Set("page", fmt.Sprintf("%d", *p.Page))
	}
	if p.Size != nil {
		params.Set("size", fmt.Sprintf("%d", *p.Size))
	}
	if p.RecurrenceStatus != nil {
		params.Set("recurrenceStatus", string(*p.RecurrenceStatus))
	}
	if p.IsPayer != nil {
		params.Set("isPayer", fmt.Sprintf("%t", *p.IsPayer))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// AutomaticPixListResponse represents list of automatic PIX
type AutomaticPixListResponse struct {
	Recurrences   []AutomaticPixResponse `json:"recurrences"`
	Page          int                    `json:"page"`
	Size          int                    `json:"size"`
	TotalElements int64                  `json:"totalElements"`
	TotalPages    int                    `json:"totalPages"`
	HasNext       bool                   `json:"hasNext"`
}

// AutomaticPixPaymentScheduleDTO represents a payment schedule
type AutomaticPixPaymentScheduleDTO struct {
	ID                 int64            `json:"id"`
	RecurrenceID       string           `json:"recurrenceId"`
	ScheduledDate      string           `json:"scheduledDate"`
	Value              float64          `json:"value"`
	Status             RecurrenceStatus `json:"status"`
	PaymentDate        *string `json:"paymentDate,omitempty"`
	EndToEnd           *string `json:"endToEnd,omitempty"`
	AuthenticationCode *string `json:"authenticationCode,omitempty"`
}

// AutomaticPixChargeListResponse represents list of automatic PIX charges
type AutomaticPixChargeListResponse struct {
	Recurrences   []AutomaticPixPaymentScheduleDTO `json:"recurrences"`
	Page          int                              `json:"page"`
	Size          int                              `json:"size"`
	TotalElements int64                            `json:"totalElements"`
	TotalPages    int                              `json:"totalPages"`
	HasNext       bool                             `json:"hasNext"`
}

// Deprecated types for backward compatibility (aliases)
type AutomaticPixRequest = StartAutomaticPixRequest
type AutomaticPixContractRequest = CreateAutomaticPixContractRequest
type AutomaticPixContractResponse = CreateAutomaticPixContractResponse
type AcceptQRCodeJourneyThreeRequest = QRCodeAcceptJourneyThreeRequest
type AcceptQRCodeRequest = QRCodeUserAcceptRequest
type AutomaticPixChargeResponse = AutomaticPixPaymentScheduleDTO
