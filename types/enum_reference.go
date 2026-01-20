package types

// EnumResponse represents a generic enum value
type EnumResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Active      *bool  `json:"active,omitempty"`
}

// EnumListResponse represents a list of enum values
type EnumListResponse struct {
	Items []EnumResponse `json:"items"`
}

// GenderEnumResponse represents GenderEnumResponse from API spec
// GET /enum/gender
type GenderResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// ProfessionEnumResponse represents ProfessionEnumResponse from API spec
// GET /enum/profession
type ProfessionResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// IssuingAuthorityResponse represents IssuingAuthorityResponse from API spec
// GET /enum/issuingAuthority
type IssuingAuthorityResponse struct {
	IssuingAuthorities []string `json:"issuingAuthorities"`
}

// PhoneCodeCountryResponse represents PhoneCodeCountryResponse from API spec
// GET /enum/uf and GET /country
type StateResponse struct {
	PhoneCode   string `json:"phoneCode"`
	CountryName string `json:"countryName"`
}

// PhoneCodeCountryResponse represents PhoneCodeCountryResponse from API spec
// GET /country
type PhoneCodeCountryResponse struct {
	PhoneCode   string `json:"phoneCode"`
	CountryName string `json:"countryName"`
}

// CreditEngineInfoRequest represents credit engine information creation/update
type CreditEngineInfoRequest struct {
	MonthlyIncome      *int64  `json:"monthlyIncome,omitempty"` // Amount in cents
	EmploymentStatus   *string `json:"employmentStatus,omitempty"`
	Employer           *string `json:"employer,omitempty"`
	Occupation         *string `json:"occupation,omitempty"`
	PoliticallyExposed *bool   `json:"politicallyExposed,omitempty"`
	NetWorth           *int64  `json:"netWorth,omitempty"` // Amount in cents
	SourceOfWealth     *string `json:"sourceOfWealth,omitempty"`
}

// CreditEngineInfoResponse represents credit engine information
type CreditEngineInfoResponse struct {
	AccountID          int64   `json:"accountId"`
	MonthlyIncome      *int64  `json:"monthlyIncome,omitempty"`
	EmploymentStatus   *string `json:"employmentStatus,omitempty"`
	Employer           *string `json:"employer,omitempty"`
	Occupation         *string `json:"occupation,omitempty"`
	PoliticallyExposed *bool   `json:"politicallyExposed,omitempty"`
	NetWorth           *int64  `json:"netWorth,omitempty"`
	SourceOfWealth     *string `json:"sourceOfWealth,omitempty"`
	CreditScore        *int    `json:"creditScore,omitempty"`
	CreditLimit        *int64  `json:"creditLimit,omitempty"` // Amount in cents
}

// ProductResponse represents a product
type ProductResponse struct {
	ProductID      int64   `json:"productId"`
	ProductName    string  `json:"productName"`
	ProductType    string  `json:"productType"`
	Description    *string `json:"description,omitempty"`
	Active         bool    `json:"active"`
	MonthlyFee     *int64  `json:"monthlyFee,omitempty"`     // Amount in cents
	TransactionFee *int64  `json:"transactionFee,omitempty"` // Amount in cents
	WithdrawalFee  *int64  `json:"withdrawalFee,omitempty"`  // Amount in cents
}

// ProductListResponse represents a list of products
type ProductListResponse struct {
	Products []ProductResponse `json:"items"`
	Total    *int              `json:"total,omitempty"`
}

// CreateProductRequest represents a request to create a product
type CreateProductRequest struct {
	ProductName    string  `json:"productName"`
	ProductType    string  `json:"productType"`
	Description    *string `json:"description,omitempty"`
	Active         bool    `json:"active"`
	MonthlyFee     *int64  `json:"monthlyFee,omitempty"`
	TransactionFee *int64  `json:"transactionFee,omitempty"`
	WithdrawalFee  *int64  `json:"withdrawalFee,omitempty"`
}

// UpdateProductRequest represents a request to update a product
type UpdateProductRequest struct {
	ProductName    *string `json:"productName,omitempty"`
	ProductType    *string `json:"productType,omitempty"`
	Description    *string `json:"description,omitempty"`
	Active         *bool   `json:"active,omitempty"`
	MonthlyFee     *int64  `json:"monthlyFee,omitempty"`
	TransactionFee *int64  `json:"transactionFee,omitempty"`
	WithdrawalFee  *int64  `json:"withdrawalFee,omitempty"`
}
