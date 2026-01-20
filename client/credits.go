package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// UpdateCreditExpiration updates credit expiration
func (c *Client) UpdateCreditExpiration(ctx context.Context, accountID int64, req *types.UpdateCreditExpirationRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/credits/expiration", accountID)
	var response types.GenericResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetUsableCredits retrieves usable credits
func (c *Client) GetUsableCredits(ctx context.Context, accountID int64, dateLimit *string) (*types.CreditsInfoResponse, error) {
	path := fmt.Sprintf("/accounts/%d/credits/to/use", accountID)
	if dateLimit != nil {
		path += "?dateLimit=" + *dateLimit
	}
	var response types.CreditsInfoResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRefundableCredits retrieves refundable credits
func (c *Client) GetRefundableCredits(ctx context.Context, accountID int64) (*types.CreditsInfoResponse, error) {
	path := fmt.Sprintf("/accounts/%d/credits/refund", accountID)
	var response types.CreditsInfoResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetExpiredCredits retrieves expired credits
func (c *Client) GetExpiredCredits(ctx context.Context, accountID int64) (*types.CreditsInfoResponse, error) {
	path := fmt.Sprintf("/accounts/%d/credits/expired", accountID)
	var response types.CreditsInfoResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
