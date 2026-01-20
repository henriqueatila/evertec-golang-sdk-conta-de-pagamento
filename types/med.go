package types

import (
	"fmt"
	"net/url"
)

// MED (Mecanismo Especial de Devolução) - Special Return Mechanism
// Required by Brazilian Central Bank for PIX fraud reporting
// Based on OpenAPI spec: openapi-med.json

// InfractionReportReason represents infraction report reason
type InfractionReportReason string

const (
	InfractionReasonRefundRequest   InfractionReportReason = "REFUND_REQUEST"
	InfractionReasonRefundCancelled InfractionReportReason = "REFUND_CANCELLED"
)

// InfractionSituation represents infraction situation
type InfractionSituation string

const (
	InfractionSituationScam             InfractionSituation = "SCAM"
	InfractionSituationAccountTakeover  InfractionSituation = "ACCOUNT_TAKEOVER"
	InfractionSituationCoercion         InfractionSituation = "COERCION"
	InfractionSituationFraudulentAccess InfractionSituation = "FRAUDULENT_ACCESS"
	InfractionSituationOther            InfractionSituation = "OTHER"
	InfractionSituationUnknown          InfractionSituation = "UNKNOWN"
)

// InfractionReportStatus represents infraction report status
type InfractionReportStatus string

const (
	InfractionStatusOpen         InfractionReportStatus = "OPEN"
	InfractionStatusAcknowledged InfractionReportStatus = "ACKNOWLEDGED"
	InfractionStatusClosed       InfractionReportStatus = "CLOSED"
	InfractionStatusCancelled    InfractionReportStatus = "CANCELLED"
)

// FraudType represents fraud type for closing infraction
type FraudType string

const (
	FraudTypeApplicationFraud FraudType = "APPLICATION_FRAUD"
	FraudTypeMuleAccount      FraudType = "MULE_ACCOUNT"
	FraudTypeScammerAccount   FraudType = "SCAMMER_ACCOUNT"
	FraudTypeOther            FraudType = "OTHER"
	FraudTypeUnknown          FraudType = "UNKNOWN"
)

// AnalysisResult represents analysis result
type AnalysisResult string

const (
	AnalysisResultAgreed    AnalysisResult = "AGREED"
	AnalysisResultDisagreed AnalysisResult = "DISAGREED"
)

// RefundReason represents refund reason
type RefundReason string

const (
	RefundReasonFraud           RefundReason = "FRAUD"
	RefundReasonOperationalFlaw RefundReason = "OPERATIONAL_FLAW"
	RefundReasonRefundCancelled RefundReason = "REFUND_CANCELLED"
)

// RefundResult represents refund result
type RefundResult string

const (
	RefundResultTotallyAccepted   RefundResult = "TOTALLY_ACCEPTED"
	RefundResultPartiallyAccepted RefundResult = "PARTIALLY_ACCEPTED"
	RefundResultRejected          RefundResult = "REJECTED"
)

// RefundRejectReason represents refund rejection reason
type RefundRejectReason string

const (
	RefundRejectNoBalance      RefundRejectReason = "NO_BALANCE"
	RefundRejectAccountClosure RefundRejectReason = "ACCOUNT_CLOSURE"
	RefundRejectOther          RefundRejectReason = "OTHER"
	RefundRejectInvalidRequest RefundRejectReason = "INVALID_REQUEST"
)

// ParticipantRole represents participant role in refund
type ParticipantRole string

const (
	ParticipantRoleRequesting ParticipantRole = "REQUESTING"
	ParticipantRoleContested  ParticipantRole = "CONTESTED"
)

// InfractionReportRequest represents an infraction report creation request
type InfractionReportRequest struct {
	TransactionID int64                  `json:"transactionId"`
	Reason        InfractionReportReason `json:"reason"`
	Situation     InfractionSituation    `json:"situation"`
	Details       *string                `json:"details,omitempty"` // Max 2000 chars
}

// InfractionReportResponse represents an infraction report (API: ExtendedInfractionReportResponse)
type InfractionReportResponse struct {
	InfractionReportID   string                 `json:"infractionReportId"`           // Required (API: id)
	Reason               InfractionReportReason `json:"reason"`                       // Required
	Situation            InfractionSituation    `json:"situation"`                    // Required
	Status               InfractionReportStatus `json:"status"`                       // Required
	ReporterIspb         string                 `json:"reporterIspb"`                 // Required (8-digit ISPB)
	CounterpartyIspb     string                 `json:"counterpartyIspb"`             // Required (8-digit ISPB)
	CreationDateTime     string                 `json:"creationDate"`                 // Required (renamed from creationDateTime)
	LastModifiedDateTime string                 `json:"lastUpdateDate"`               // Required (renamed)
	Details              *string                `json:"details,omitempty"`            // Optional (maxLength: 2000)
	FraudMarkerID        *string                `json:"fraudMarkerId,omitempty"`      // Optional
	AnalysisResult       *AnalysisResult        `json:"result,omitempty"`             // Optional (API: result)
	TransactionID        *int64                 `json:"transactionId,omitempty"`      // Extended field
	IsReporter           *bool                  `json:"isReporter,omitempty"`         // Extended field
	IsCounterparty       *bool                  `json:"isCounterparty,omitempty"`     // Extended field
	AnalysisDetails      *string                `json:"analysisDetails,omitempty"`    // Extended field
	FraudType            *FraudType             `json:"fraudType,omitempty"`          // Extended field
	EndToEnd             *string                `json:"endToEnd,omitempty"`           // Extended field
	ReportedParticipant  *string                `json:"reportedParticipant,omitempty"` // Extended field
	ReporterParticipant  *string                `json:"reporterParticipant,omitempty"` // Extended field
}

// ListInfractionReportsParams represents query params for listing infraction reports
type ListInfractionReportsParams struct {
	IncludeIndirectParticipants *bool                   `json:"includeIndirectParticipants,omitempty"`
	IsReporter                  *bool                   `json:"isReporter,omitempty"`
	IsCounterparty              *bool                   `json:"isCounterparty,omitempty"`
	Status                      *InfractionReportStatus `json:"status,omitempty"`
	IncludeDetails              *bool                   `json:"includeDetails,omitempty"`
	ModifiedAfter               *string                 `json:"modifiedAfter,omitempty"`  // ISO 8601
	ModifiedBefore              *string                 `json:"modifiedBefore,omitempty"` // ISO 8601
	Limit                       *int                    `json:"limit,omitempty"`          // Default 20
}

// QueryString returns query string for list infraction reports
func (p *ListInfractionReportsParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.IncludeIndirectParticipants != nil {
		params.Set("includeIndirectParticipants", fmt.Sprintf("%t", *p.IncludeIndirectParticipants))
	}
	if p.IsReporter != nil {
		params.Set("isReporter", fmt.Sprintf("%t", *p.IsReporter))
	}
	if p.IsCounterparty != nil {
		params.Set("isCounterparty", fmt.Sprintf("%t", *p.IsCounterparty))
	}
	if p.Status != nil {
		params.Set("status", string(*p.Status))
	}
	if p.IncludeDetails != nil {
		params.Set("includeDetails", fmt.Sprintf("%t", *p.IncludeDetails))
	}
	if p.ModifiedAfter != nil {
		params.Set("modifiedAfter", *p.ModifiedAfter)
	}
	if p.ModifiedBefore != nil {
		params.Set("modifiedBefore", *p.ModifiedBefore)
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *p.Limit))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// ListInfractionReportsResponse represents list of infraction reports (API: ListInfractionReportResponse)
type ListInfractionReportsResponse struct {
	ResponseTime      string                     `json:"responseTime"`      // Required (date-time)
	InfractionReports []InfractionReportResponse `json:"infractionReport"`  // Required (singular in API)
	HasMoreElements   bool                       `json:"hasMoreElements"`   // Required
	WarningMessage    *string                    `json:"warningMessage,omitempty"`
}

// CloseInfractionReportRequest represents a request to close an infraction report
type CloseInfractionReportRequest struct {
	InfractionReportID string         `json:"infractionReportId"`
	FraudType          *FraudType     `json:"fraudType,omitempty"`
	AnalysisResult     AnalysisResult `json:"analysisResult"`
	AnalysisDetails    *string        `json:"analysisDetails,omitempty"` // Max 2000 chars
}

// RefundSolicitationRequest represents a refund solicitation request
type RefundSolicitationRequest struct {
	TransactionID string       `json:"transactionId"` // 32 chars alphanumeric (E2E)
	RefundReason  RefundReason `json:"refundReason"`
	RefundAmount  float64      `json:"refundAmount"`
	RefundDetails *string      `json:"refundDetails,omitempty"` // Max 2000 chars
}

// RefundResponse represents a refund response (API: RefundResponse)
type RefundResponse struct {
	RefundID               string             `json:"id"`                    // Required (API: id)
	TransactionID          string             `json:"transactionId"`         // Required (32-char alphanumeric)
	Status                 RefundStatus       `json:"status"`                // Required (OPEN/CLOSED/CANCELLED)
	RefundReason           RefundReason       `json:"refundReason"`          // Required
	RefundAmount           float64            `json:"refundAmount"`          // Required
	RequestingParticipant  string             `json:"requestingParticipant"` // Required (8-digit ISPB)
	ContestedParticipant   string             `json:"contestedParticipant"`  // Required (8-digit ISPB)
	RefundTransactionID    string             `json:"refundTransactionId"`   // Required (32-char alphanumeric)
	CreationDateTime       string             `json:"creationTime"`          // Required
	LastModifiedDateTime   string             `json:"lastModified"`          // Required
	RefundDetails          *string            `json:"refundDetails,omitempty"`
	InfractionReportID     *string            `json:"infractionReportId,omitempty"`
	AnalysisResult         *RefundResult      `json:"analysisResult,omitempty"`
	AnalysisDetails        *string            `json:"analysisDetails,omitempty"`
	RefundRejectionReason  *RefundRejectReason `json:"refundRejectionReason,omitempty"`
}

// CloseRefundRequest represents a request to close a refund
type CloseRefundRequest struct {
	RefundID              string              `json:"refundId"`
	TransactionID         string              `json:"transactionId"` // 32 chars alphanumeric
	Result                RefundResult        `json:"result"`
	RefundRejectReason    *RefundRejectReason `json:"refundRejectReason,omitempty"`
	RefundAnalysisDetails *string             `json:"refundAnalysisDetails,omitempty"` // Max 2000 chars
}

// ListRefundsParams represents query params for listing refunds
type ListRefundsParams struct {
	IncludeIndirectParticipants *bool            `json:"includeIndirectParticipants,omitempty"`
	ParticipantRole             *ParticipantRole `json:"participantRole,omitempty"` // Required
	Status                      *string          `json:"status,omitempty"`
	IncludeDetails              *bool            `json:"includeDetails,omitempty"`
	ModifiedAfter               *string          `json:"modifiedAfter,omitempty"`
	ModifiedBefore              *string          `json:"modifiedBefore,omitempty"`
	Limit                       *int             `json:"limit,omitempty"` // Default 20
}

// QueryString returns query string for list refunds
func (p *ListRefundsParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.IncludeIndirectParticipants != nil {
		params.Set("includeIndirectParticipants", fmt.Sprintf("%t", *p.IncludeIndirectParticipants))
	}
	if p.ParticipantRole != nil {
		params.Set("participantRole", string(*p.ParticipantRole))
	}
	if p.Status != nil {
		params.Set("status", *p.Status)
	}
	if p.IncludeDetails != nil {
		params.Set("includeDetails", fmt.Sprintf("%t", *p.IncludeDetails))
	}
	if p.ModifiedAfter != nil {
		params.Set("modifiedAfter", *p.ModifiedAfter)
	}
	if p.ModifiedBefore != nil {
		params.Set("modifiedBefore", *p.ModifiedBefore)
	}
	if p.Limit != nil {
		params.Set("limit", fmt.Sprintf("%d", *p.Limit))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// ListRefundsResponse represents list of refunds (API: ListRefundSolicitationResponse)
type ListRefundsResponse struct {
	ResponseTime    string           `json:"responseTime"`    // Required (date-time)
	Refunds         []RefundResponse `json:"refund"`          // Required (singular in API)
	HasMoreElements bool             `json:"hasMoreElements"` // Required
	WarningMessage  *string          `json:"warningMessage,omitempty"`
}
