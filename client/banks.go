package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// ListBanks retrieves available banks for TED
func (c *Client) ListBanks(ctx context.Context, params *types.ListBanksParams) (*types.BanksResponse, error) {
	path := "/banks"
	if params != nil {
		path += params.QueryString()
	}
	var response types.BanksResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetScheduledOperations retrieves all scheduled operations for an account
func (c *Client) GetScheduledOperations(ctx context.Context, accountID int64) (*types.ScheduledOperationsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/scheduleds", accountID)
	var response types.ScheduledOperationsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CheckAPIStatus checks API availability
func (c *Client) CheckAPIStatus(ctx context.Context) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.get(ctx, "/status", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CheckIntegrationStatus checks integration module status
func (c *Client) CheckIntegrationStatus(ctx context.Context) (*types.IntegrationStatusResponse, error) {
	var response types.IntegrationStatusResponse
	if err := c.get(ctx, "/status/integrationModules", &response); err != nil {
		return nil, err
	}
	return &response, nil
}
