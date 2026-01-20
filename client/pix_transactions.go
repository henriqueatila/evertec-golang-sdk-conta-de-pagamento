package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// PIX Transaction Operations - Aligned with OpenAPI Specification

// DoPixPayment executes a PIX payment transaction (POST /pix/transactions/payment)
func (c *Client) DoPixPayment(ctx context.Context, req *types.PixPaymentRequest) (*types.PixPaymentResponse, error) {
	var response types.PixPaymentResponse
	if err := c.post(ctx, "/pix/transactions/payment", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DoPixChargeback executes a PIX chargeback (POST /pix/transactions/chargeback)
func (c *Client) DoPixChargeback(ctx context.Context, req *types.PixChargebackRequest) (*types.PixChargebackResponse, error) {
	var response types.PixChargebackResponse
	if err := c.post(ctx, "/pix/transactions/chargeback", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelPixSchedule cancels a scheduled PIX transaction (POST /pix/transactions/cancelSchedule)
func (c *Client) CancelPixSchedule(ctx context.Context, req *types.PixCancelScheduleRequest) (*types.PixCancelScheduleResponse, error) {
	var response types.PixCancelScheduleResponse
	if err := c.post(ctx, "/pix/transactions/cancelSchedule", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePrecautionaryBlock creates a precautionary block on a transaction (POST /pix/backoffice/precautionaryBlock)
func (c *Client) CreatePrecautionaryBlock(ctx context.Context, req *types.PixPrecautionaryBlockRequest) (*types.PixPrecautionaryBlockResponse, error) {
	var response types.PixPrecautionaryBlockResponse
	if err := c.post(ctx, "/pix/backoffice/precautionaryBlock", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePrecautionaryBlock updates a precautionary block (POST /pix/backoffice/precautionaryBlock/update)
func (c *Client) UpdatePrecautionaryBlock(ctx context.Context, req *types.PixUpdatePrecautionaryBlockRequest) (*types.PixUpdatePrecautionaryBlockResponse, error) {
	var response types.PixUpdatePrecautionaryBlockResponse
	if err := c.post(ctx, "/pix/backoffice/precautionaryBlock/update", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPixTransactionLimit retrieves PIX transaction limits for an account (GET /pix/transactions/{accountId}/limit)
func (c *Client) GetPixTransactionLimit(ctx context.Context, accountID int64) (*types.PixGetLimitResponse, error) {
	path := fmt.Sprintf("/pix/transactions/%d/limit", accountID)
	var response types.PixGetLimitResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPixPaymentByE2E retrieves a PIX payment by its E2E identifier (GET /pix/transactions/payment/{e2e})
func (c *Client) GetPixPaymentByE2E(ctx context.Context, e2e string) (*types.GetPixInfoResponse, error) {
	path := fmt.Sprintf("/pix/transactions/payment/%s", e2e)
	var response types.GetPixInfoResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPSPs retrieves the list of PIX Service Providers (GET /pix/psps)
func (c *Client) ListPSPs(ctx context.Context) (*types.PspListResponse, error) {
	var response types.PspListResponse
	if err := c.get(ctx, "/pix/psps", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPixKeyInfo retrieves information about a PIX key (GET /pix/keys/{accountId}/{key})
func (c *Client) GetPixKeyInfo(ctx context.Context, accountID int64, key string) (*types.SearchKeyResponse, error) {
	path := fmt.Sprintf("/pix/keys/%d/%s", accountID, key)
	var response types.SearchKeyResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
