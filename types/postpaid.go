package types

// Postpaid types based on OpenAPI spec: openapi-cadastral.json

// FormattedAmountDTO represents FormattedAmountDTO from API spec
type FormattedAmountDTO struct {
	Amount       float64 `json:"amount"`       // number
	CurrencyCode int32   `json:"currencyCode"` // int32
}

// AmountDTO represents AmountDTO from API spec
type AmountDTO struct {
	Amount       int32 `json:"amount"`       // int32
	CurrencyCode int32 `json:"currencyCode"` // int32
}

// SourceAudit represents SourceAudit from API spec
type SourceAudit struct {
	Source    *string `json:"source,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// FeeDTO represents FeeDTO from API spec
type FeeDTO struct {
	Source                 *string    `json:"source,omitempty"`
	FeeType                *string    `json:"feeType,omitempty"`
	TransactionDescription *string    `json:"transactionDescripton,omitempty"` // Note: API has typo "transactionDescripton"
	Value                  *AmountDTO `json:"value,omitempty"`
}

// InstallmentDetailsDTO represents InstallmentDetailsDTO from API spec
type InstallmentDetailsDTO struct {
	TotalInstallments   *int32     `json:"totalInstallments,omitempty"`   // int32
	InstallmentNumber   *int32     `json:"installmentNumber,omitempty"`   // int32
	InstallmentAmount   *AmountDTO `json:"installmentAmount,omitempty"`
	RemainingAmount     *AmountDTO `json:"remainingAmount,omitempty"`
	OriginalTransaction *string    `json:"originalTransaction,omitempty"`
}

// TransactionsDTO represents TransactionsDTO from API spec
type TransactionsDTO struct {
	TransactionID            *string                `json:"transactionId,omitempty"`
	CardID                   *string                `json:"cardId,omitempty"`
	InternationalTransaction *bool                  `json:"internationalTransaction,omitempty"`
	TransactionDate          *string                `json:"transactionDate,omitempty"`
	TransactionDescription   *string                `json:"transactionDescription,omitempty"`
	Amount                   *AmountDTO             `json:"amount,omitempty"`
	AmountDollar             *AmountDTO             `json:"amountDollar,omitempty"`
	CredencialType           *string                `json:"credencialType,omitempty"`
	Fees                     []FeeDTO               `json:"fees,omitempty"`
	InstallmentNumber        *int32                 `json:"installmentNumber,omitempty"` // int32
	InstallmentDetails       *InstallmentDetailsDTO `json:"installmentDetails,omitempty"`
	LastFourDigits           *string                `json:"last_four_digits,omitempty"`
	DebitOrCredit            *string                `json:"debit_or_credit,omitempty"`
}

// AccountPostPaidResponse represents AccountPostPaidResponse from API spec
// GET /postpaid/account/{accountId}
type AccountPostPaidResponse struct {
	AccountID            int64               `json:"accountId"`            // int64
	PaysmartAccountID    *string             `json:"paysmartAccountId,omitempty"`
	MaxCreditInfo        *FormattedAmountDTO `json:"maxCreditInfo,omitempty"`
	CurrentLimit         *FormattedAmountDTO `json:"currentLimit,omitempty"`
	CreditLimit          *FormattedAmountDTO `json:"creditLimit,omitempty"`
	CurrentWithdrawLimit *FormattedAmountDTO `json:"currentWithdrawLimit,omitempty"`
	WithdrawLimit        *FormattedAmountDTO `json:"withdrawLimit,omitempty"`
	PaymentDue           *int32              `json:"paymentDue,omitempty"` // int32
}

// PaysmartOpenStatementResponse represents PaysmartOpenStatementResponse from API spec
// GET /postpaid/statements/{accountId}/open-statement
type PaysmartOpenStatementResponse struct {
	AccountID                    *string           `json:"accountId,omitempty"`
	PaymentDue                   *string           `json:"paymentDue,omitempty"`
	CloseDate                    *string           `json:"closeDate,omitempty"`
	PreviousBalance              *AmountDTO        `json:"previousBalance,omitempty"`
	Balance                      *AmountDTO        `json:"balance,omitempty"`
	CreditLimit                  *AmountDTO        `json:"creditLimit,omitempty"`
	WithdrawalCreditLimit        *AmountDTO        `json:"withdrawalCreditLimit,omitempty"`
	CurrentCreditLimit           *AmountDTO        `json:"currentCreditLimit,omitempty"`
	CurrentWithdrawalCreditLimit *AmountDTO        `json:"currentWithdrawalCreditLimit,omitempty"`
	TransactionsList             []TransactionsDTO `json:"transactionsList,omitempty"`
	FirstTransactionInserted     *string           `json:"firstTransactionInserted,omitempty"`
	OpeningDateTime              *string           `json:"openingDateTime,omitempty"`
	SourceAudit                  *SourceAudit      `json:"sourceAudit,omitempty"`
	QueryDate                    *string           `json:"query_date,omitempty"`
}

// DueDateOption represents due date option from API spec
// GET /postpaid/account/{accountId}/dueDates
type DueDateOption struct {
	Value       int    `json:"value"`
	Description string `json:"description"`
}
