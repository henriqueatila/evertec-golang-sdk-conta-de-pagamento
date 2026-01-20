package types

// Travel types based on OpenAPI spec: openapi-cadastral.json

// NotifyTripRequest represents NotifyTripRequest from API spec
// POST /travel/account/{accountId}/notify
type NotifyTripRequest struct {
	TravelStartDate string  `json:"travelStartDate"` // Format: date
	TravelEndDate   string  `json:"travelEndDate"`   // Format: date
	Countries       []int64 `json:"countries"`       // Country IDs (int64)
	CardIdList      []int64 `json:"cardIdList"`      // Card IDs (int64)
}

// TravelNoticeResponse represents TravelNoticeResponse from API spec
type TravelNoticeResponse struct {
	Cards          []int64  `json:"cards"`
	CountryCodes   []string `json:"countryCodes"`
	TravelNoticeID string   `json:"travelNoticeId"`
	BeginDate      string   `json:"beginDate"`
	EndDate        string   `json:"endDate"`
}

// GetAccountTravelingResponse represents GetAccountTravelingResponse from API spec
// GET /travel/account/{accountId}/notify
type GetAccountTravelingResponse struct {
	TravelNotices []TravelNoticeResponse `json:"travelNotices"`
}

// Pais represents Pais from API spec (country)
type Pais struct {
	ID         int64  `json:"id"`         // int64
	NamePt     string `json:"namePt"`
	Name       string `json:"name"`
	Alpha2Code string `json:"alpha2Code"`
	Alpha3Code string `json:"alpha3Code"`
	Code       int64  `json:"code"`       // int64
	PhoneCode  string `json:"phoneCode"`
}

// CountryListResponse represents CountryListResponse from API spec
// GET /travel/getCountries
type CountryListResponse struct {
	Message   *string `json:"message,omitempty"`
	Code      *int    `json:"code,omitempty"`
	Countries []Pais  `json:"countries"`
	Total     int     `json:"total"`
}

// ChangeUserPasswordRequest represents ChangeUserPasswordRequest from API spec
// PUT /accounts/{accountId}/changeUserPassword
type ChangeUserPasswordRequest struct {
	HashPassword string `json:"hashPassword"`
}

// UpdateNameInAccountRequest represents UpdateNameInAccountRequest from API spec
// PUT /accounts/{accountId}/name
type UpdateNameInAccountRequest struct {
	Name string `json:"name"`
}

// ContaDigitalGenericResponse represents ContaDigitalGenericResponse from API spec
// Generic response for many operations
type ContaDigitalGenericResponse struct {
	Message *string `json:"message,omitempty"`
	DaCode  *int    `json:"da_code,omitempty"`
}
