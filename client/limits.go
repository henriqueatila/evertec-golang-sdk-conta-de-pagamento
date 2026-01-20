package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Account Limit Operations

// GetAccountLimit retrieves the account limit for a specific type
func (c *Client) GetAccountLimit(ctx context.Context, accountID int64, limitType types.LimitType) (*types.LimitResponse, error) {
	path := fmt.Sprintf("/accounts/limit/%d/%s/getLimit", accountID, limitType)
	var response types.LimitResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateAccountLimit updates the account limit for a specific type
func (c *Client) UpdateAccountLimit(ctx context.Context, accountID int64, limitType types.LimitType, req *types.UpdateLimitRequest) (*types.LimitResponse, error) {
	path := fmt.Sprintf("/accounts/limit/%d/%s", accountID, limitType)
	var response types.LimitResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateAccountNightTimeLimit updates the night-time limit settings for a specific type
func (c *Client) UpdateAccountNightTimeLimit(ctx context.Context, accountID int64, limitType types.LimitType, req *types.UpdateNightTimeLimitRequest) (*types.LimitResponse, error) {
	path := fmt.Sprintf("/accounts/limit/%d/%s/startNightTime", accountID, limitType)
	var response types.LimitResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetMaximumLimitIssuer retrieves the maximum limit allowed by the issuer
func (c *Client) GetMaximumLimitIssuer(ctx context.Context, accountID int64, limitType types.LimitType) (*types.MaximumLimitResponse, error) {
	path := fmt.Sprintf("/accounts/limit/%d/%s/getMaximumLimitIssuer", accountID, limitType)
	var response types.MaximumLimitResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Account Fees Operations
// Based on OpenAPI spec: Returns ContaDigitalGenericResponse

// GetAccountFees retrieves the fees for an account
// Returns ContaDigitalGenericResponse from API spec
func (c *Client) GetAccountFees(ctx context.Context, accountID int64) (*types.ContaDigitalGenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/fees", accountID)
	var response types.ContaDigitalGenericResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCardIssuanceFee retrieves the card issuance fee
// Returns ContaDigitalGenericResponse from API spec
func (c *Client) GetCardIssuanceFee(ctx context.Context, accountID int64) (*types.ContaDigitalGenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/fees/cardissuer", accountID)
	var response types.ContaDigitalGenericResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCardReissueFee retrieves the card reissuance fee
// Returns ContaDigitalGenericResponse from API spec
func (c *Client) GetCardReissueFee(ctx context.Context, accountID int64) (*types.ContaDigitalGenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/fees/cardreissue", accountID)
	var response types.ContaDigitalGenericResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Product Limit Operations

// UpdateProductLimitByType updates the product limit for a specific limit type
func (c *Client) UpdateProductLimitByType(ctx context.Context, limitType types.LimitType, req *types.ProductLimitRequest) error {
	path := fmt.Sprintf("/limit/%s/productLimit", limitType)
	return c.put(ctx, path, req, nil)
}

// SearchProductLimitByType searches for product limits by type
func (c *Client) SearchProductLimitByType(ctx context.Context, limitType types.LimitType, req *types.SearchProductLimitRequest) ([]types.ProductLimitResponse, error) {
	path := fmt.Sprintf("/limit/%s/searchProductLimit", limitType)
	var response []types.ProductLimitResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return response, nil
}
