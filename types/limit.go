package types

// Limit types based on OpenAPI spec: openapi-cadastral.json

// LimitRequest represents LimitRequest from API spec
// PUT /accounts/limit/{accountId}/{limitType}
type LimitRequest struct {
	DayLimit            *float64 `json:"dayLimit,omitempty"`
	DayTransactionLimit *float64 `json:"dayTransactionLimit,omitempty"`
	NightLimit          *float64 `json:"nightLimit,omitempty"`
	NightTransactionLimit *float64 `json:"nightTransactionLimit,omitempty"`
}

// LimitNightTimeRequest represents LimitNightTimeRequest from API spec
// PUT /accounts/limit/{accountId}/{limitType}/startNightTime
type LimitNightTimeRequest struct {
	StartNightTime string `json:"startNightTime"` // Format: HH:MM
}

// UpdateNightTimeLimitRequest alias for LimitNightTimeRequest
type UpdateNightTimeLimitRequest = LimitNightTimeRequest

// UpdateLimitRequest alias for LimitRequest
type UpdateLimitRequest = LimitRequest

// LimitDto represents LimitDto from API spec
type LimitDto struct {
	StartNightTime        string  `json:"startNightTime"`
	DayLimit              float64 `json:"dayLimit"`              // number
	DayTransactionLimit   float64 `json:"dayTransactionLimit"`   // number
	NightLimit            float64 `json:"nightLimit"`            // number
	NightTransactionLimit float64 `json:"nightTransactionLimit"` // number
	Status                int32   `json:"status"`                // int32
}

// LimitResponse represents LimitResponse from API spec
// GET /accounts/limit/{accountId}/{limitType}/getLimit
type LimitResponse struct {
	Message string    `json:"message"`
	Limit   *LimitDto `json:"limit,omitempty"`
	DaCode  int32     `json:"da_code"` // int32
}

// TransactionLimit represents TransactionLimit from API spec
type TransactionLimit struct {
	TotalLimitFormatted           float64 `json:"totalLimitFormatted"`           // number
	TotalNightLimitFormatted      float64 `json:"totalNightLimitFormatted"`      // number
	TotalLimit                    int64   `json:"totalLimit"`                    // int64
	TotalNightLimit               int64   `json:"totalNightLimit"`               // int64
	LimitTransactionFormatted     float64 `json:"limitTransactionFormatted"`     // number
	NightLimitTransactionFormatted float64 `json:"nightLimitTransactionFormatted"` // number
	LimitTransaction              int64   `json:"limitTransaction"`              // int64
	NightLimitTransaction         int64   `json:"nightLimitTransaction"`         // int64
}

// GetLimitResponse represents GetLimitResponse from API spec
// GET /accounts/limit/{accountId}/{limitType}/getMaximumLimitIssuer
type GetLimitResponse struct {
	Message string            `json:"message"`
	Limit   *TransactionLimit `json:"limit,omitempty"`
	DaCode  int32             `json:"da_code"` // int32
}

// MaximumLimitResponse alias for GetLimitResponse
type MaximumLimitResponse = GetLimitResponse

// UpdateProductLimitRequest represents UpdateProductLimitRequest from API spec
// PUT /limit/{limitType}/productLimit
type UpdateProductLimitRequest struct {
	Environment            string   `json:"environment"`
	RecurrenceType         string   `json:"recurrenceType"`
	DayLimit               *float64 `json:"dayLimit,omitempty"`
	DayTransactionLimit    *float64 `json:"dayTransactionLimit,omitempty"`
	NightLimit             *float64 `json:"nightLimit,omitempty"`
	NightTransactionLimit  *float64 `json:"nightTransactionLimit,omitempty"`
}

// ProductLimitRequest alias for UpdateProductLimitRequest
type ProductLimitRequest = UpdateProductLimitRequest

// SearchProductLimitRequest represents SearchProductLimitRequest from API spec
// POST /limit/{limitType}/searchProductLimit
type SearchProductLimitRequest struct {
	Environment    *string `json:"environment,omitempty"`
	RecurrenceType *string `json:"recurrenceType,omitempty"`
}

// ProductLimitResponse - Note: API returns ContaDigitalGenericResponse
// This is kept for backwards compatibility but the actual API returns generic response
type ProductLimitResponse struct {
	Message string `json:"message,omitempty"`
	DaCode  int32  `json:"da_code,omitempty"`
}

// GetProductLimitDaySchedulingResponse represents GetProductLimitDaySchedulingResponse from API spec
// GET /product/limitScheduling/{productId}
type GetProductLimitDaySchedulingResponse struct {
	Message         string                 `json:"message"`
	LimitScheduling *LimitDaySchedulingDTO `json:"limitScheduling,omitempty"`
	DaCode          int32                  `json:"da_code"`
}

// LimitDaySchedulingDTO represents LimitDaySchedulingDTO from API spec
type LimitDaySchedulingDTO struct {
	ProductID      int64  `json:"productId"`
	Monday         bool   `json:"monday"`
	Tuesday        bool   `json:"tuesday"`
	Wednesday      bool   `json:"wednesday"`
	Thursday       bool   `json:"thursday"`
	Friday         bool   `json:"friday"`
	Saturday       bool   `json:"saturday"`
	Sunday         bool   `json:"sunday"`
	StartNightTime string `json:"startNightTime"`
}

// UpdateProductLimitDaySchedulingRequest represents UpdateProductLimitDaySchedulingRequest from API spec
// PUT /product/limitScheduling
type UpdateProductLimitDaySchedulingRequest struct {
	ProductID      int64   `json:"productId"`
	Monday         *bool   `json:"monday,omitempty"`
	Tuesday        *bool   `json:"tuesday,omitempty"`
	Wednesday      *bool   `json:"wednesday,omitempty"`
	Thursday       *bool   `json:"thursday,omitempty"`
	Friday         *bool   `json:"friday,omitempty"`
	Saturday       *bool   `json:"saturday,omitempty"`
	Sunday         *bool   `json:"sunday,omitempty"`
	StartNightTime *string `json:"startNightTime,omitempty"`
}

// UpdateProductLimitDaySchedulingResponse represents UpdateProductLimitDaySchedulingResponse from API spec
type UpdateProductLimitDaySchedulingResponse struct {
	Message         string                 `json:"message"`
	LimitScheduling *LimitDaySchedulingDTO `json:"limitScheduling,omitempty"`
	DaCode          int32                  `json:"da_code"`
}
