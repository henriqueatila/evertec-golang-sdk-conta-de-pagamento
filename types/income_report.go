package types

// Income Report Operations (Informe de Rendimentos)
// Based on openapi-informe-rendimentos.json

// AccountBalanceByDateResponse represents balance at a specific date (API: AccountBalanceByDateResponse)
type AccountBalanceByDateResponse struct {
	Balance *float64 `json:"balance,omitempty"`
	Date    *string  `json:"date,omitempty"` // date format
}

// AccountBalanceByYearResponse represents yearly balance for an account (API: AccountBalanceByYearResponse)
type AccountBalanceByYearResponse struct {
	AccountID        *int64                        `json:"accountId,omitempty"`
	Document         *string                       `json:"document,omitempty"`
	Name             *string                       `json:"name,omitempty"`
	BeginningOfYear  *AccountBalanceByDateResponse `json:"beginningOfTheYear,omitempty"`
	EndOfYear        *AccountBalanceByDateResponse `json:"endOfTheYear,omitempty"`
}

// ListAccountBalanceByYearResponse represents list of yearly balances with pagination (API: ListAccountBalanceByYearResponse)
type ListAccountBalanceByYearResponse struct {
	IncomeReports []AccountBalanceByYearResponse `json:"incomeReports"` // Required
	HasNextPage   bool                           `json:"hasNextPage"`   // Required
	Page          int32                          `json:"page"`          // Required
	TotalPages    int32                          `json:"totalPages"`    // Required
	TotalElements int64                          `json:"totalElements"` // Required
}

// IncomeReportResponse represents an income report for tax purposes (legacy/extended)
type IncomeReportResponse struct {
	AccountID   int64  `json:"accountId"`
	Year        int    `json:"year"`
	ReportData  string `json:"reportData"` // Base64 encoded PDF or JSON data
	GeneratedAt string `json:"generatedAt"`
}

// YearlyBalanceResponse represents yearly balance summary (SDK extended type)
type YearlyBalanceResponse struct {
	AccountID      int64 `json:"accountId"`
	Year           int   `json:"year"`
	OpeningBalance int64 `json:"openingBalance"` // Amount in cents
	ClosingBalance int64 `json:"closingBalance"` // Amount in cents
	TotalCredits   int64 `json:"totalCredits"`   // Amount in cents
	TotalDebits    int64 `json:"totalDebits"`    // Amount in cents
}

// AllAccountsYearlyBalanceResponse represents yearly balances for all accounts (SDK extended type)
type AllAccountsYearlyBalanceResponse struct {
	Year     int                     `json:"year"`
	Accounts []YearlyBalanceResponse `json:"accounts"`
	Total    *int                    `json:"total,omitempty"`
}
