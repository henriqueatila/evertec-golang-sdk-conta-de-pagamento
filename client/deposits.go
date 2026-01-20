package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// ListDepositOrders retrieves deposit orders for an account
func (c *Client) ListDepositOrders(ctx context.Context, accountID int64, params *types.ListDepositOrdersParams) (*types.DepositOrdersResponse, error) {
	path := fmt.Sprintf("/accounts/%d/deposits/order", accountID)
	if params != nil {
		path += params.QueryString()
	}
	var response types.DepositOrdersResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateDepositOrder creates a deposit order
func (c *Client) CreateDepositOrder(ctx context.Context, accountID int64, req *types.CreateDepositOrderRequest) (*types.CreateDepositOrderResponse, error) {
	path := fmt.Sprintf("/accounts/%d/deposits/order", accountID)
	var response types.CreateDepositOrderResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListActiveDepositOrders retrieves active deposit orders
func (c *Client) ListActiveDepositOrders(ctx context.Context, accountID int64) (*types.DepositOrdersResponse, error) {
	path := fmt.Sprintf("/accounts/%d/deposits/order/active", accountID)
	var response types.DepositOrdersResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelDepositOrder cancels a deposit order
func (c *Client) CancelDepositOrder(ctx context.Context, accountID, depositOrderID int64) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/deposits/order/%d", accountID, depositOrderID)
	var response types.GenericResponse
	if err := c.delete(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
