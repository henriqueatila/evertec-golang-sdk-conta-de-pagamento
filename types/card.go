package types

import "time"

// CreateCardRequest represents the request to create a new debit card
type CreateCardRequest struct {
	Tag *string `json:"tag,omitempty"` // Card nickname/label
}

// CreateVirtualCardRequest represents the request to create a virtual card
type CreateVirtualCardRequest struct {
	Tag            *string `json:"tag,omitempty"`
	PhysicalCardID *int64  `json:"physicalCardId,omitempty"` // If creating from physical card
}

// BlockCardRequest represents BlockCardRequest from API spec
// POST /cards/{accountId}/block/{cardId}
type BlockCardRequest struct {
	DestinationStatus int32  `json:"destinationStatus"` // int32
	BlockMotivation   string `json:"blockMotivation"`
}

// UnblockCardRequest - Note: No official schema in API, endpoint may not require body
type UnblockCardRequest struct{}

// ActivateCardRequest represents ActivateCardRequest from API spec
// POST /cards/{accountId}/activate/{cardId}
type ActivateCardRequest struct {
	Last4Digits string `json:"last4Digits"` // string
	CancelOlds  bool   `json:"cancelOlds"`  // boolean
	Password    string `json:"password"`    // string
}

// ChangeCardPinRequest represents ChangeCardPinRequest from API spec
// POST /cards/{accountId}/changePin/{cardId}
type ChangeCardPinRequest struct {
	NewPin        string `json:"newPin"`        // string
	ConfirmNewPin string `json:"confirmNewPin"` // string
}

// UpdateCardTagRequest represents the request to update card tag/nickname
type UpdateCardTagRequest struct {
	Tag string `json:"tag"`
}

// ReissueCardRequest represents the request to reissue a card
type ReissueCardRequest struct {
	Reason *string `json:"reason,omitempty"`
}

// CardConfigurationRequest represents the request to configure card settings
type CardConfigurationRequest struct {
	// Transaction type limits (amount in cents)
	ContactlessLimit      *int64 `json:"contactlessLimit,omitempty"`
	OnlinePurchaseEnabled *bool  `json:"onlinePurchaseEnabled,omitempty"`
	ContactlessEnabled    *bool  `json:"contactlessEnabled,omitempty"`
	InternationalEnabled  *bool  `json:"internationalEnabled,omitempty"`
	WithdrawalEnabled     *bool  `json:"withdrawalEnabled,omitempty"`

	// Daily limits (amount in cents)
	DailyPurchaseLimit   *int64 `json:"dailyPurchaseLimit,omitempty"`
	DailyWithdrawalLimit *int64 `json:"dailyWithdrawalLimit,omitempty"`
}

// CardConfigurationResponse represents the card configuration settings
type CardConfigurationResponse struct {
	ContactlessLimit      *int64 `json:"contactlessLimit,omitempty"`
	OnlinePurchaseEnabled *bool  `json:"onlinePurchaseEnabled,omitempty"`
	ContactlessEnabled    *bool  `json:"contactlessEnabled,omitempty"`
	InternationalEnabled  *bool  `json:"internationalEnabled,omitempty"`
	WithdrawalEnabled     *bool  `json:"withdrawalEnabled,omitempty"`
	DailyPurchaseLimit    *int64 `json:"dailyPurchaseLimit,omitempty"`
	DailyWithdrawalLimit  *int64 `json:"dailyWithdrawalLimit,omitempty"`
}

// CardResponse represents CardVo from API spec
// GET /cards/{accountId}/{cardId} response
type CardResponse struct {
	// Required fields from CardVo
	ID                   int64      `json:"id"`
	PersonID             int64      `json:"personID"`
	CardType             string     `json:"cardType"`
	FunctionType         string     `json:"functionType"`
	Modality             string     `json:"modality"`
	ExpirationDate       time.Time  `json:"expirationDate"`
	ExpirationDateString *string    `json:"expirationDateString,omitempty"`
	ExpirationDateMMAA   *string    `json:"expirationDateMMAA,omitempty"`
	Status               int32      `json:"status"`
	StatusDescription    *string    `json:"statusDescription,omitempty"`
	StatusDateTime       *time.Time `json:"statusDateTime,omitempty"`
	PrintedName          *string    `json:"printedName,omitempty"`
	BIN                  *string    `json:"bin,omitempty"`
	LastFour             *string    `json:"lastFour,omitempty"`
	Ownership            *string    `json:"ownership,omitempty"`
	ExternalCardID       *string    `json:"externalCardId,omitempty"`
	PAN                  *string    `json:"pan,omitempty"`
	CVV                  *string    `json:"cvv,omitempty"`
	ActivateDate         *time.Time `json:"activateDate,omitempty"`
	EmbossingFileName    *string    `json:"embossingFileName,omitempty"`
	Tag                  *string    `json:"tag,omitempty"`
	AccountID            int64      `json:"accountId"`
	IssuerID             *int32     `json:"issuerId,omitempty"`

	// Security controls
	SecurityControls *CardSecurityControlsVo `json:"securityControls,omitempty"`
}

// CardSecurityControlsVo represents card security controls from API spec
type CardSecurityControlsVo struct {
	ContactlessEnabled   *bool  `json:"contactlessEnabled,omitempty"`
	OnlineEnabled        *bool  `json:"onlineEnabled,omitempty"`
	InternationalEnabled *bool  `json:"internationalEnabled,omitempty"`
	WithdrawalEnabled    *bool  `json:"withdrawalEnabled,omitempty"`
	ContactlessLimit     *int64 `json:"contactlessLimit,omitempty"`
}

// CardListResponse represents a list of cards
type CardListResponse struct {
	Cards []CardResponse `json:"items"`
	Total *int           `json:"total,omitempty"`
}

// CardSearchRequest represents filters for searching cards
type CardSearchRequest struct {
	Status       *CardStatus   `json:"status,omitempty"`
	CardType     *CardType     `json:"cardType,omitempty"`
	CardCategory *CardCategory `json:"cardCategory,omitempty"`
	AccountID    *int64        `json:"accountId,omitempty"`
	MaskedNumber *string       `json:"maskedNumber,omitempty"`
}

// PostpaidCardRequest represents the request to create a post-paid card
type PostpaidCardRequest struct {
	AccountID   int64   `json:"accountId"`
	CreditLimit int64   `json:"creditLimit"` // Amount in cents
	Tag         *string `json:"tag,omitempty"`
}

// PostpaidCardSettingsRequest represents settings for post-paid card
type PostpaidCardSettingsRequest struct {
	OnlinePurchaseEnabled *bool `json:"onlinePurchaseEnabled,omitempty"`
	InternationalEnabled  *bool `json:"internationalEnabled,omitempty"`
	ContactlessEnabled    *bool `json:"contactlessEnabled,omitempty"`
	WithdrawalEnabled     *bool `json:"withdrawalEnabled,omitempty"`
}

// PostpaidCardSettingsResponse represents settings for post-paid card
type PostpaidCardSettingsResponse struct {
	CardID                int64 `json:"cardId"`
	OnlinePurchaseEnabled bool  `json:"onlinePurchaseEnabled"`
	InternationalEnabled  bool  `json:"internationalEnabled"`
	ContactlessEnabled    bool  `json:"contactlessEnabled"`
	WithdrawalEnabled     bool  `json:"withdrawalEnabled"`
}

// Note: Travel-related types (NotifyTripRequest, TravelNoticeResponse, CountryListResponse, Pais)
// are defined in types/travel.go based on the official OpenAPI spec

// CardPaysmartResponse represents a Paysmart card
type CardPaysmartResponse struct {
	CardIDPaysmart string     `json:"cardIdPaysmart"`
	CardID         int64      `json:"cardId"`
	AccountID      int64      `json:"accountId"`
	Status         CardStatus `json:"status"`
	MaskedNumber   *string    `json:"maskedNumber,omitempty"`
	ExpiryDate     *string    `json:"expiryDate,omitempty"`
	ProcessorCode  *string    `json:"processorCode,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
}

// DefaultCardConfigurationResponse represents default card configuration
type DefaultCardConfigurationResponse struct {
	ContactlessLimit      *int64 `json:"contactlessLimit,omitempty"`
	OnlinePurchaseEnabled bool   `json:"onlinePurchaseEnabled"`
	ContactlessEnabled    bool   `json:"contactlessEnabled"`
	InternationalEnabled  bool   `json:"internationalEnabled"`
	WithdrawalEnabled     bool   `json:"withdrawalEnabled"`
	DailyPurchaseLimit    *int64 `json:"dailyPurchaseLimit,omitempty"`
	DailyWithdrawalLimit  *int64 `json:"dailyWithdrawalLimit,omitempty"`
}
