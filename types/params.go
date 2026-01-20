package types

import (
	"fmt"
	"net/url"
)

// ListAccountsParams represents query parameters for listing accounts
type ListAccountsParams struct {
	Status      *AccountStatus `json:"status,omitempty"`
	AccountType *AccountType   `json:"accountType,omitempty"`
	Document    *string        `json:"document,omitempty"`
	Name        *string        `json:"name,omitempty"`
	First       *int           `json:"first,omitempty"`
	Max         *int           `json:"max,omitempty"`
}

func (p *ListAccountsParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.Status != nil {
		params.Set("status", string(*p.Status))
	}
	if p.AccountType != nil {
		params.Set("accountType", string(*p.AccountType))
	}
	if p.Document != nil {
		params.Set("document", *p.Document)
	}
	if p.Name != nil {
		params.Set("name", *p.Name)
	}
	if p.First != nil {
		params.Set("first", fmt.Sprintf("%d", *p.First))
	}
	if p.Max != nil {
		params.Set("max", fmt.Sprintf("%d", *p.Max))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// StatementParams represents query parameters for account statement
type StatementParams struct {
	StartDate *string `json:"startDate,omitempty"` // Format: YYYY-MM-DD
	EndDate   *string `json:"endDate,omitempty"`   // Format: YYYY-MM-DD
	OrderType *string `json:"orderType,omitempty"` // ASC or DESC
	IsPix     *bool   `json:"isPix,omitempty"`
	Type      *string `json:"type,omitempty"` // credit or debit
	First     *int    `json:"first,omitempty"`
	Max       *int    `json:"max,omitempty"`
}

func (p *StatementParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.StartDate != nil {
		params.Set("startDate", *p.StartDate)
	}
	if p.EndDate != nil {
		params.Set("endDate", *p.EndDate)
	}
	if p.OrderType != nil {
		params.Set("orderType", *p.OrderType)
	}
	if p.IsPix != nil {
		params.Set("isPix", fmt.Sprintf("%t", *p.IsPix))
	}
	if p.Type != nil {
		params.Set("type", *p.Type)
	}
	if p.First != nil {
		params.Set("first", fmt.Sprintf("%d", *p.First))
	}
	if p.Max != nil {
		params.Set("max", fmt.Sprintf("%d", *p.Max))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// ListCardsParams represents query parameters for listing cards
type ListCardsParams struct {
	Status   *CardStatus `json:"status,omitempty"`
	CardType *CardType   `json:"cardType,omitempty"`
	First    *int        `json:"first,omitempty"`
	Max      *int        `json:"max,omitempty"`
}

func (p *ListCardsParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.Status != nil {
		params.Set("status", string(*p.Status))
	}
	if p.CardType != nil {
		params.Set("cardType", string(*p.CardType))
	}
	if p.First != nil {
		params.Set("first", fmt.Sprintf("%d", *p.First))
	}
	if p.Max != nil {
		params.Set("max", fmt.Sprintf("%d", *p.Max))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// SearchCardsParams represents query parameters for searching cards
type SearchCardsParams struct {
	Status       *CardStatus   `json:"status,omitempty"`
	CardType     *CardType     `json:"cardType,omitempty"`
	CardCategory *CardCategory `json:"cardCategory,omitempty"`
	AccountID    *int64        `json:"accountId,omitempty"`
	First        *int          `json:"first,omitempty"`
	Max          *int          `json:"max,omitempty"`
}

func (p *SearchCardsParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.Status != nil {
		params.Set("status", string(*p.Status))
	}
	if p.CardType != nil {
		params.Set("cardType", string(*p.CardType))
	}
	if p.CardCategory != nil {
		params.Set("cardCategory", string(*p.CardCategory))
	}
	if p.AccountID != nil {
		params.Set("accountId", fmt.Sprintf("%d", *p.AccountID))
	}
	if p.First != nil {
		params.Set("first", fmt.Sprintf("%d", *p.First))
	}
	if p.Max != nil {
		params.Set("max", fmt.Sprintf("%d", *p.Max))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// ListProposalsParams represents query parameters for listing proposals
type ListProposalsParams struct {
	Status   *string `json:"status,omitempty"`
	Document *string `json:"document,omitempty"`
	First    *int    `json:"first,omitempty"`
	Max      *int    `json:"max,omitempty"`
}

func (p *ListProposalsParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.Status != nil {
		params.Set("status", *p.Status)
	}
	if p.Document != nil {
		params.Set("document", *p.Document)
	}
	if p.First != nil {
		params.Set("first", fmt.Sprintf("%d", *p.First))
	}
	if p.Max != nil {
		params.Set("max", fmt.Sprintf("%d", *p.Max))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// ListDepositOrdersParams represents query parameters for listing deposit orders
type ListDepositOrdersParams struct {
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
	OrderType *string `json:"orderType,omitempty"`
	First     *int    `json:"first,omitempty"`
	Max       *int    `json:"max,omitempty"`
}

func (p *ListDepositOrdersParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.StartDate != nil {
		params.Set("startDate", *p.StartDate)
	}
	if p.EndDate != nil {
		params.Set("endDate", *p.EndDate)
	}
	if p.OrderType != nil {
		params.Set("orderType", *p.OrderType)
	}
	if p.First != nil {
		params.Set("first", fmt.Sprintf("%d", *p.First))
	}
	if p.Max != nil {
		params.Set("max", fmt.Sprintf("%d", *p.Max))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}

// ListBanksParams represents query parameters for listing banks
type ListBanksParams struct {
	OrderType *string `json:"orderType,omitempty"`
	First     *int    `json:"first,omitempty"`
	Max       *int    `json:"max,omitempty"`
}

func (p *ListBanksParams) QueryString() string {
	if p == nil {
		return ""
	}
	params := url.Values{}
	if p.OrderType != nil {
		params.Set("orderType", *p.OrderType)
	}
	if p.First != nil {
		params.Set("first", fmt.Sprintf("%d", *p.First))
	}
	if p.Max != nil {
		params.Set("max", fmt.Sprintf("%d", *p.Max))
	}
	if len(params) == 0 {
		return ""
	}
	return "?" + params.Encode()
}
