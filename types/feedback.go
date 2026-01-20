package types

// Feedback types based on OpenAPI spec: openapi-cadastral.json

// FeedbackRequest represents FeedbackRequest from API spec
// POST /accounts/feedback/send
// All fields except images are required
type FeedbackRequest struct {
	AccountID          int64    `json:"accountId"`          // required
	Message            string   `json:"message"`            // required, minLength: 10
	AppVersion         string   `json:"appVersion"`         // required
	OperationalSystem  string   `json:"operationalSystem"`  // required
	Amount             float64  `json:"amount"`             // required (number)
	DateAndHourProblem string   `json:"dateAndHourProblem"` // required
	Images             []string `json:"images,omitempty"`
}

// FeedbackStatementRequest represents FeedbackStatementRequest from API spec
// POST /accounts/feedback/statement/send
// All fields except images are required
type FeedbackStatementRequest struct {
	OperationalSystem string   `json:"operationalSystem"` // required
	AppVersion        string   `json:"appVersion"`        // required
	TransactionID     int64    `json:"transactionId"`     // required
	Message           string   `json:"message"`           // required, minLength: 10
	Images            []string `json:"images,omitempty"`
}

// FeedbackResponse represents FeedbackResponse from API spec
// GET /accounts/{accountId}/feedback/lastTransactionError
// All fields are required
type FeedbackResponse struct {
	LastTransactionErrorAmount      float64 `json:"lastTransactionErrorAmount"`      // required (number)
	LastTransactionErrorDateAndHour string  `json:"lastTransactionErrorDateAndHour"` // required (date-time)
}

// ContaDigitalResponse represents ContaDigitalResponse from API spec
// Generic response for many operations
type ContaDigitalResponse struct {
	Result  bool   `json:"result"`
	Message any    `json:"message,omitempty"` // Can be string or object
	Code    *int   `json:"code,omitempty"`
}
