package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// PayBill processes a bill payment
func (c *Client) PayBill(ctx context.Context, req *types.BillPaymentRequest) (*types.BillPaymentResponse, error) {
	var response types.BillPaymentResponse
	if err := c.put(ctx, "/bill/payment", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PayBillBatch processes batch bill payments
func (c *Client) PayBillBatch(ctx context.Context, req *types.BillPaymentRequest) (*types.BillPaymentResponse, error) {
	var response types.BillPaymentResponse
	if err := c.put(ctx, "/bill/payment/batch", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBillInfo retrieves bill information
func (c *Client) GetBillInfo(ctx context.Context, req *types.GetBillInfoRequest) (*types.GetBillInfoResponse, error) {
	var response types.GetBillInfoResponse
	if err := c.post(ctx, "/bill/info", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelScheduledBill cancels a scheduled bill payment
func (c *Client) CancelScheduledBill(ctx context.Context, accountID, schedulingID int64) (*types.CancelBillResponse, error) {
	path := fmt.Sprintf("/bill/account/%d/scheduling/%d", accountID, schedulingID)
	var response types.CancelBillResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListScheduledBills retrieves scheduled bill payments
func (c *Client) ListScheduledBills(ctx context.Context, accountID int64) (*types.ScheduledBillsResponse, error) {
	path := fmt.Sprintf("/bill/account/%d/schedules", accountID)
	var response types.ScheduledBillsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Account Bill Payment Operations (Alternative endpoints)

// PayBillByAccount processes a bill payment via account endpoint
func (c *Client) PayBillByAccount(ctx context.Context, accountID int64, req *types.BillPaymentRequest) (*types.BillPaymentResponse, error) {
	path := fmt.Sprintf("/accounts/%d/billpayment", accountID)
	var response types.BillPaymentResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBillInfoByAccount retrieves bill information via account endpoint
func (c *Client) GetBillInfoByAccount(ctx context.Context, accountID int64, req *types.GetBillInfoRequest) (*types.GetBillInfoResponse, error) {
	path := fmt.Sprintf("/accounts/%d/billpayment", accountID)
	var response types.GetBillInfoResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelScheduledBillByAccount cancels a scheduled bill payment via account endpoint
func (c *Client) CancelScheduledBillByAccount(ctx context.Context, accountID, schedulingID int64) (*types.CancelBillResponse, error) {
	path := fmt.Sprintf("/accounts/%d/billpayment/scheduled/%d", accountID, schedulingID)
	var response types.CancelBillResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PayBillBatchByAccount processes batch bill payments via account endpoint
func (c *Client) PayBillBatchByAccount(ctx context.Context, accountID int64, req *types.BatchBillPaymentRequest) (*types.BillPaymentResponse, error) {
	path := fmt.Sprintf("/accounts/%d/billpayment/batch", accountID)
	var response types.BillPaymentResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListScheduledBillsByAccount retrieves scheduled bill payments via account endpoint
func (c *Client) ListScheduledBillsByAccount(ctx context.Context, accountID int64) (*types.ScheduledBillPaymentListResponse, error) {
	path := fmt.Sprintf("/accounts/%d/billpayment/scheduled", accountID)
	var response types.ScheduledBillPaymentListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
