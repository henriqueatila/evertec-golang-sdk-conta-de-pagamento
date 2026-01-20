package client

import (
	"context"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Authorizer Operations (Card Authorization)

// DoSummaryPurchase processes a summary purchase authorization
func (c *Client) DoSummaryPurchase(ctx context.Context, req *types.SummaryPurchaseRequest) (*types.AuthorizationResponse, error) {
	var response types.AuthorizationResponse
	if err := c.post(ctx, "/summary-purchases", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelSummaryPurchase cancels a summary purchase
func (c *Client) CancelSummaryPurchase(ctx context.Context, req *types.CancelPurchaseRequest) (*types.AuthorizationResponse, error) {
	var response types.AuthorizationResponse
	if err := c.post(ctx, "/summary-purchases/cancel", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DoSummaryChargeback processes a summary chargeback
func (c *Client) DoSummaryChargeback(ctx context.Context, req *types.ChargebackRequest) (*types.AuthorizationResponse, error) {
	var response types.AuthorizationResponse
	if err := c.post(ctx, "/summary-chargebacks", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelSummaryChargeback cancels a summary chargeback
func (c *Client) CancelSummaryChargeback(ctx context.Context, req *types.CancelChargebackRequest) (*types.AuthorizationResponse, error) {
	var response types.AuthorizationResponse
	if err := c.post(ctx, "/summary-chargebacks/cancel", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
