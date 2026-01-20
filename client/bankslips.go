package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// ListBankslips retrieves bankslips for an account
func (c *Client) ListBankslips(ctx context.Context, accountID int64) (*types.BankslipsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/bankslip", accountID)
	var response types.BankslipsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateBankslip creates a deposit bankslip
func (c *Client) CreateBankslip(ctx context.Context, accountID int64, req *types.CreateBankslipRequest) (*types.CreateBankslipResponse, error) {
	path := fmt.Sprintf("/accounts/%d/bankslip", accountID)
	var response types.CreateBankslipResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListBankslipsByStatus retrieves bankslips filtered by status
func (c *Client) ListBankslipsByStatus(ctx context.Context, accountID int64, status string) (*types.BankslipsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/bankslip/%s", accountID, status)
	var response types.BankslipsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListBankslipsByStatusAndDate retrieves bankslips filtered by status and date
func (c *Client) ListBankslipsByStatusAndDate(ctx context.Context, accountID int64, status, createdAt string) (*types.BankslipsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/bankslip/%s/%s", accountID, status, createdAt)
	var response types.BankslipsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateBankslipV2 creates a bankslip using the v2 API
func (c *Client) CreateBankslipV2(ctx context.Context, req *types.BankslipV2Request) (*types.BankslipV2Response, error) {
	var response types.BankslipV2Response
	if err := c.post(ctx, "/bankslip/v2/generate", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
