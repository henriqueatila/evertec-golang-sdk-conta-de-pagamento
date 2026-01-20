package types

import "time"

// TokenValidationRequest represents TokenValidationRequest from API spec
// POST /tokens/cellphone/validate, POST /tokens/email/validate
type TokenValidationRequest struct {
	TokenID   int64  `json:"tokenId"`
	TokenCode string `json:"tokenCode"`
}

// ImageFormat represents supported image formats for identification
type ImageFormat string

const (
	ImageFormatJPEG ImageFormat = "jpeg"
	ImageFormatPNG  ImageFormat = "png"
	ImageFormatPDF  ImageFormat = "pdf"
	ImageFormatCSV  ImageFormat = "csv"
)

// UserIdentificationImage represents user identification image (API: UserIdentificationImage)
type UserIdentificationImage struct {
	Image     string      `json:"image"`     // Required (enum: Base64)
	Format    ImageFormat `json:"format"`    // Required (pattern: jpeg|png|pdf|csv)
	FileName  string      `json:"fileName"`  // Required
	UpdatedAt *string     `json:"updatedAt,omitempty"`
}

// CreateCreditEngineInfoRequest represents credit engine info request (API: CreateCreditEngineInfoRequest)
type CreateCreditEngineInfoRequest struct {
	ResidenceDuration      int32   `json:"residenceDuration"`      // Required (min: 0)
	ResidenceType          string  `json:"residenceType"`          // Required (enum: Aluguel, Proprio, Outro)
	RentAmount             float64 `json:"rentAmount"`             // Required (min: 0)
	MaritalStatus          string  `json:"maritalStatus"`          // Required (enum: Solteiro(a), Casado(a), Viuvo(a))
	AdditionalExpenses     string  `json:"additionalExpenses"`     // Required
	ContractType           string  `json:"contractType"`           // Required (enum: CLT, PJ, Outro)
	MinimumIncome          float64 `json:"minimumIncome"`          // Required (min: 0)
	MaximumIncome          float64 `json:"maximumIncome"`          // Required (min: 0)
	ProfessionID           int32   `json:"professionId"`           // Required
	CurrentCompanyName     string  `json:"currentCompanyName"`     // Required
	CurrentCompanyDuration int32   `json:"currentCompanyDuration"` // Required (min: 0)
	JobTitle               string  `json:"jobTitle"`               // Required (length: 1-30)
}

// ProposalAccountRequest represents ProposalAccountRequest from API spec
// POST /accounts/proposal
type ProposalAccountRequest struct {
	PersonalDocument     string                  `json:"personalDocument"`       // required (CPF)
	PersonalName         string                  `json:"personalName"`           // required, pattern
	ProductID            int32                   `json:"productID"`              // required, int32
	BirthDate            string                  `json:"birthDate"`              // required
	FatherName           *string                 `json:"fatherName,omitempty"`
	SocialName           *string                 `json:"socialName,omitempty"`
	MotherName           string                  `json:"motherName"`             // required
	LocalAreaCodeCellPhone int32                 `json:"localAreaCodeCellPhone"` // required, int32
	CellPhoneNumber      int32                   `json:"cellPhoneNumber"`        // required, int32
	PersonalEmail        string                  `json:"personalEmail"`          // required
	Address              *AddressRequest         `json:"address"`        // Required (pointer allows omission in partial flows)
	CellPhoneToken       *TokenValidationRequest `json:"cellPhoneToken"` // Required (pointer allows omission in partial flows)
	EmailToken           *TokenValidationRequest `json:"emailToken"`     // Required (pointer allows omission in partial flows)
	BranchID             int32                   `json:"branchId"`               // required, int32
	UserIdentificationImages []UserIdentificationImage `json:"userIdentificationImages,omitempty"` // Fixed: was []string
	Ownership            *int32                  `json:"ownership,omitempty"`    // New field
	IsCorporateAccount   *bool                   `json:"isCorporateAccount,omitempty"` // New field (default: false)
	CreditEngineInfo     *CreateCreditEngineInfoRequest `json:"creditEngineInfo,omitempty"` // New field
	CreditLimit          *CreateAccountCreditLimitPostPaidRequest `json:"creditLimit,omitempty"` // New field
	Mobile               *bool                   `json:"mobile,omitempty"`
	Password             *string                 `json:"password,omitempty"`
	MainAccountID        *int64                  `json:"mainAccountId,omitempty"`
	Nationality          *string                 `json:"nationality,omitempty"`
	IdentityDocument     *string                 `json:"identityDocument,omitempty"` // maxLength: 20
	IssuingAuthority     *string                 `json:"issuingAuthority,omitempty"` // enum
	IdentityDocumentUF   *string                 `json:"identityDocumentUF,omitempty"` // enum
	GenderID             *int32                  `json:"genderId,omitempty"`         // int32
	ArrangementType      *string                 `json:"arrangementType,omitempty"`
	ExternalProductID    *int64                  `json:"externalProductId,omitempty"`
}

// UpdateAccountRequest represents UpdateAccountRequest from API spec
// PUT /accounts
type UpdateAccountRequest struct {
	DocID                  *string                 `json:"docId,omitempty"`
	AccountID              int64                   `json:"accountId"`
	Name                   *string                 `json:"name,omitempty"`
	TradingName            *string                 `json:"tradingName,omitempty"`
	CorporateName          *string                 `json:"corporateName,omitempty"`
	FatherName             *string                 `json:"fatherName,omitempty"`
	MotherName             *string                 `json:"motherName,omitempty"`
	SocialName             *string                 `json:"socialName,omitempty"`
	LocalAreaCodeCellPhone *int32                  `json:"localAreaCodeCellPhone,omitempty"`
	CellPhoneNumber        *int32                  `json:"cellPhoneNumber,omitempty"`
	PersonalEmail          *string                 `json:"personalEmail,omitempty"`
	ProposalID             *int64                  `json:"proposalId,omitempty"`
	CellPhoneToken         *TokenValidationRequest `json:"cellPhoneToken,omitempty"`
	EmailToken             *TokenValidationRequest `json:"emailToken,omitempty"`
	SkipAPIMobile          *bool                   `json:"skipApiMobile,omitempty"`
	ArrangementType        *string                 `json:"arrangementType,omitempty"`
	BirthDate              *string                 `json:"birthDate,omitempty"`
}

// UpdateAccountNameRequest represents the request to update account name
type UpdateAccountNameRequest struct {
	Name string `json:"name"`
}

// ChangePasswordRequest represents the request to change user password
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// LinkAccountRequest represents the request to link sub-accounts
type LinkAccountRequest struct {
	SubAccountID int64 `json:"subAccountId"`
}

// UnlinkAccountRequest represents the request to unlink accounts
type UnlinkAccountRequest struct {
	MainAccountID int64 `json:"mainAccountId"`
	SubAccountID  int64 `json:"subAccountId"`
}

// VerifyAccountExistsRequest represents the request to verify account existence
type VerifyAccountExistsRequest struct {
	Document string `json:"document"`
}

// AccountDataResponse represents the account data returned by the API
type AccountDataResponse struct {
	AccountID    int64         `json:"accountId"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Document     string        `json:"document"`
	DocumentType DocumentType  `json:"documentType"`
	Phone        string        `json:"phone"`
	BirthDate    *string       `json:"birthDate,omitempty"`
	MotherName   *string       `json:"motherName,omitempty"`
	Gender       *Gender       `json:"gender,omitempty"`
	Profession   *string       `json:"profession,omitempty"`
	Status       AccountStatus `json:"status"`
	AccountType  AccountType   `json:"accountType"`
	CreatedAt    *time.Time    `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time    `json:"updatedAt,omitempty"`

	// Company fields (when AccountType is COMPANY)
	CompanyName           *string `json:"companyName,omitempty"`
	TradeName             *string `json:"tradeName,omitempty"`
	CompanyDocument       *string `json:"companyDocument,omitempty"`
	CompanyFoundationDate *string `json:"companyFoundationDate,omitempty"`

	// Balance information
	Balance *int64 `json:"balance,omitempty"` // Amount in cents

	// Additional info
	MainAccountID *int64 `json:"mainAccountId,omitempty"`
}

// AccountListResponse represents a paginated list of accounts
// Deprecated: Use ContaSimplificadaListResponse which matches the official API spec
type AccountListResponse struct {
	Accounts []AccountDataResponse `json:"items"`
	Total    *int                  `json:"total,omitempty"`
	Page     *int                  `json:"page,omitempty"`
}

// BalanceResponse represents the account balance
// GET /accounts/{accountId}/balance response from API spec
type BalanceResponse struct {
	Message   string    `json:"message"`
	IDAccount int64     `json:"idAccount"`
	Balance   int64     `json:"balance"`   // Amount in cents
	DateTime  time.Time `json:"dateTime"`
	DACode    int32     `json:"da_code"`
}

// StatementEntry represents a single transaction in the statement (EntriesDTO from API spec)
// GET /accounts/{accountId}/statement response entries
type StatementEntry struct {
	// Required fields
	TransactionID          string    `json:"transactionId"`          // string per API spec (not int64)
	TransactionDate        time.Time `json:"transactionDate"`
	TransactionDescription string    `json:"transactionDescription"`
	Amount                 *AmountDTO `json:"amount"`
	DebitOrCredit          string    `json:"debit_or_credit"`

	// Optional fields
	CardID                   *string             `json:"cardId,omitempty"`
	InternationalTransaction *bool               `json:"internationalTransaction,omitempty"`
	AmountDollar             *AmountDTO          `json:"amountDollar,omitempty"`
	CredentialType           *string             `json:"credencialType,omitempty"` // Note: API uses "credencialType" (typo preserved)
	Fees                     []FeeDTO            `json:"fees,omitempty"`
	InstallmentNumber        *int32              `json:"installmentNumber,omitempty"`
	InstallmentDetails       *InstallmentDetailsDTO `json:"installmentDetails,omitempty"`
	LastFourDigits           *string             `json:"last_four_digits,omitempty"`
}

// StatementResponse represents the account statement
// GET /accounts/{accountId}/statement response from API spec
type StatementResponse struct {
	Message      string           `json:"message"`
	Total        int32            `json:"total"`
	AccountID    int64            `json:"accountId"`
	Entries      []StatementEntry `json:"entries"`
	Count        int32            `json:"count"`
	DACode       int32            `json:"da_code"`
	StartDate    *string          `json:"startDate,omitempty"` // Format: YYYY-MM-DD
	EndDate      *string          `json:"endDate,omitempty"`   // Format: YYYY-MM-DD
}

// AddressRequest represents AddressRequest from API spec
type AddressRequest struct {
	StreetAvenue           string  `json:"streetAvenue"`                     // required
	Number                 string  `json:"number"`                           // required, maxLength: 5
	PostalCode             string  `json:"postalCode"`                       // required
	Complement             *string `json:"complement,omitempty"`
	Neighborhood           *string `json:"neighborhood,omitempty"`
	City                   *string `json:"city,omitempty"`
	State                  *string `json:"state,omitempty"`
	SendCardAccountAddress *bool   `json:"sendCardAccountAddress,omitempty"`
	AutoComplete           *bool   `json:"autoComplete,omitempty"`
}

// AddressResponse represents an address returned by the API
type AddressResponse struct {
	AddressID    int64        `json:"addressId"`
	PostalCode   string       `json:"postalCode"`
	Street       string       `json:"street"`
	Number       string       `json:"number"`
	Complement   *string      `json:"complement,omitempty"`
	Neighborhood string       `json:"neighborhood"`
	City         string       `json:"city"`
	State        string       `json:"state"`
	Country      *string      `json:"country,omitempty"`
	AddressType  *AddressType `json:"addressType,omitempty"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
}

// PostalCodeLookupResponse represents address information from postal code lookup
type PostalCodeLookupResponse struct {
	PostalCode   string  `json:"postalCode"`
	Street       string  `json:"street"`
	Neighborhood string  `json:"neighborhood"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      *string `json:"country,omitempty"`
}

// TokenOperationRequest represents token generation/validation request
type TokenOperationRequest struct {
	Token *string `json:"token,omitempty"`
	Data  *string `json:"data,omitempty"`
}

// TokenOperationResponse represents token operation response
type TokenOperationResponse struct {
	Token     *string `json:"token,omitempty"`
	Valid     bool    `json:"valid"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// ProposalDataResponse represents account proposal data
type ProposalDataResponse struct {
	ProposalID  int64          `json:"proposalId"`
	AccountID   int64          `json:"accountId"`
	Status      ProposalStatus `json:"status"`
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	Document    string           `json:"document"`
	Phone       string           `json:"phone"`
	BirthDate   *string          `json:"birthDate,omitempty"`
	MotherName  *string          `json:"motherName,omitempty"`
	Address     *AddressResponse `json:"address,omitempty"`
	CreatedAt   *string          `json:"createdAt,omitempty"`
	ProcessedAt *string          `json:"processedAt,omitempty"`
}

// CreateCompanyAccountRequest represents company account creation request
type CreateCompanyAccountRequest struct {
	// Company information
	CompanyName           string  `json:"companyName"`
	TradeName             *string `json:"tradeName,omitempty"`
	CompanyDocument       string  `json:"companyDocument"` // CNPJ
	CompanyFoundationDate *string `json:"companyFoundationDate,omitempty"`

	// Representative information
	RepresentativeName      string  `json:"representativeName"`
	RepresentativeDocument  string  `json:"representativeDocument"` // CPF
	RepresentativeEmail     string  `json:"representativeEmail"`
	RepresentativePhone     string  `json:"representativePhone"`
	RepresentativeBirthDate *string `json:"representativeBirthDate,omitempty"`

	// Address
	Address *AddressRequest `json:"address,omitempty"`

	// Password
	Password *string `json:"password,omitempty"`
}
