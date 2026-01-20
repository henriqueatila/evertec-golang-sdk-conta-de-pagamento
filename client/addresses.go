package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// ListAddresses retrieves addresses for an account
func (c *Client) ListAddresses(ctx context.Context, accountID int64) ([]types.AddressResponse, error) {
	path := fmt.Sprintf("/accounts/%d/address", accountID)
	var response []types.AddressResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateAddress creates a new address for an account
func (c *Client) CreateAddress(ctx context.Context, accountID int64, req *types.AddressRequest) (*types.AddressResponse, error) {
	path := fmt.Sprintf("/accounts/%d/address", accountID)
	var response types.AddressResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAddress retrieves a specific address
func (c *Client) GetAddress(ctx context.Context, accountID, addressID int64) (*types.AddressResponse, error) {
	path := fmt.Sprintf("/accounts/%d/address/%d", accountID, addressID)
	var response types.AddressResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateAddress updates an address
func (c *Client) UpdateAddress(ctx context.Context, accountID, addressID int64, req *types.AddressRequest) (*types.AddressResponse, error) {
	path := fmt.Sprintf("/accounts/%d/address/%d", accountID, addressID)
	var response types.AddressResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteAddress deletes an address
func (c *Client) DeleteAddress(ctx context.Context, accountID, addressID int64) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/address/%d", accountID, addressID)
	var response types.GenericResponse
	if err := c.delete(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// LookupPostalCode looks up address by postal code
func (c *Client) LookupPostalCode(ctx context.Context, postalCode string) (*types.PostalCodeLookupResponse, error) {
	path := fmt.Sprintf("/address/postalcode/%s", postalCode)
	var response types.PostalCodeLookupResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
