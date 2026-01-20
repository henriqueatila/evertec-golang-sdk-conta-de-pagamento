package client

import (
	"context"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Reference Data / Enum Operations

// GetStates retrieves all Brazilian states (UF)
func (c *Client) GetStates(ctx context.Context) ([]types.StateResponse, error) {
	var response []types.StateResponse
	if err := c.get(ctx, "/enum/uf", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetProfessions retrieves all available professions
func (c *Client) GetProfessions(ctx context.Context) ([]types.ProfessionResponse, error) {
	var response []types.ProfessionResponse
	if err := c.get(ctx, "/enum/profession", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetIssuingAuthorities retrieves all document issuing authorities
func (c *Client) GetIssuingAuthorities(ctx context.Context) ([]types.IssuingAuthorityResponse, error) {
	var response []types.IssuingAuthorityResponse
	if err := c.get(ctx, "/enum/issuingAuthority", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetGenders retrieves all available gender options
func (c *Client) GetGenders(ctx context.Context) ([]types.GenderResponse, error) {
	var response []types.GenderResponse
	if err := c.get(ctx, "/enum/gender", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCountries retrieves all countries with phone codes
// Returns PhoneCodeCountryResponse from API spec
func (c *Client) GetCountries(ctx context.Context) ([]types.PhoneCodeCountryResponse, error) {
	var response []types.PhoneCodeCountryResponse
	if err := c.get(ctx, "/country", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetAllBanks retrieves all banks (alternative to ListBanks)
func (c *Client) GetAllBanks(ctx context.Context) ([]types.BankResponse, error) {
	var response []types.BankResponse
	if err := c.get(ctx, "/banco", &response); err != nil {
		return nil, err
	}
	return response, nil
}
