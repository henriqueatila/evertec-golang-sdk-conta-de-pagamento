package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// InternalTransfer performs an internal transfer between accounts
func (c *Client) InternalTransfer(ctx context.Context, accountID int64, req *types.InternalTransferRequest) (*types.InternalTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/transfer", accountID)
	var response types.InternalTransferResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// InternalTransferArrangement performs an internal transfer via arrangement
func (c *Client) InternalTransferArrangement(ctx context.Context, accountID int64, req *types.InternalTransferRequest) (*types.InternalTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/transfer/arrangement", accountID)
	var response types.InternalTransferResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BankTransfer performs an external TED/DOC transfer
func (c *Client) BankTransfer(ctx context.Context, accountID int64, req *types.BankTransferRequest) (*types.BankTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/banktransfer", accountID)
	var response types.BankTransferResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelScheduledTransfer cancels a scheduled external transfer
func (c *Client) CancelScheduledTransfer(ctx context.Context, accountID, schedulingID int64) (*types.CancelTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/banktransfer/scheduled/cancel/%d", accountID, schedulingID)
	var response types.CancelTransferResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListScheduledTransfers retrieves scheduled external transfers
func (c *Client) ListScheduledTransfers(ctx context.Context, accountID int64) (*types.ScheduledTransfersResponse, error) {
	path := fmt.Sprintf("/accounts/%d/banktransfer/scheduled", accountID)
	var response types.ScheduledTransfersResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BatchInternalTransfer performs batch internal transfers
func (c *Client) BatchInternalTransfer(ctx context.Context, accountID int64, req *types.BatchTransferRequest) (*types.BatchTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/transfer/batch", accountID)
	var response types.BatchTransferResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBatchTransfers retrieves batch transfer status
func (c *Client) GetBatchTransfers(ctx context.Context, accountID int64) ([]types.BatchTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/transfer/batch", accountID)
	var response []types.BatchTransferResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetBatchTransferStatus retrieves specific batch transfer status
func (c *Client) GetBatchTransferStatus(ctx context.Context, accountID int64, processingCode string) (*types.BatchTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/transfer/batch/%s", accountID, processingCode)
	var response types.BatchTransferResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CheckRecipientAccount verifies a recipient account
func (c *Client) CheckRecipientAccount(ctx context.Context, accountID, recipientAccountID int64) (*types.CheckRecipientAccountResponse, error) {
	path := fmt.Sprintf("/accounts/%d/transfers/checkRecipientAccount/%d", accountID, recipientAccountID)
	var response types.CheckRecipientAccountResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelInternalTransfer cancels an internal transfer
func (c *Client) CancelInternalTransfer(ctx context.Context, accountID int64, req *types.CancelInternalTransferRequest) (*types.CancelInternalTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cancelTransfer", accountID)
	var response types.CancelInternalTransferResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// TransferByID performs a transfer using document ID
func (c *Client) TransferByID(ctx context.Context, document string, req *types.TransferByIDRequest) (*types.InternalTransferResponse, error) {
	path := fmt.Sprintf("/accounts/%s/transfer/idid", document)
	var response types.InternalTransferResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
