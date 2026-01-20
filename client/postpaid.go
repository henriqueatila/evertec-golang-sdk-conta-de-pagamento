package client

import (
	"context"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// PostPaidPaymentBalance processes payment using account balance
func (c *Client) PostPaidPaymentBalance(ctx context.Context, req *types.PostPaidPaymentBalanceRequest) (*types.PostPaidPaymentResponse, error) {
	var response types.PostPaidPaymentResponse
	if err := c.post(ctx, "/postpaid/payment/balance", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PostPaidInstallmentSimulation simulates installment payment
func (c *Client) PostPaidInstallmentSimulation(ctx context.Context, req *types.PostPaidInstallmentSimulationRequest) (*types.PostPaidInstallmentSimulationResponse, error) {
	var response types.PostPaidInstallmentSimulationResponse
	if err := c.post(ctx, "/postpaid/payment/installment/simulation", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PostPaidInstallmentPix requests installment via PIX
func (c *Client) PostPaidInstallmentPix(ctx context.Context, req *types.PostPaidInstallmentRequest) (*types.PostPaidInstallmentPixResponse, error) {
	var response types.PostPaidInstallmentPixResponse
	if err := c.post(ctx, "/postpaid/payment/installment/request/pix", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PostPaidInstallmentAccountBalance requests installment via account balance
func (c *Client) PostPaidInstallmentAccountBalance(ctx context.Context, req *types.PostPaidInstallmentRequest) (*types.PostPaidInstallmentBalanceResponse, error) {
	var response types.PostPaidInstallmentBalanceResponse
	if err := c.post(ctx, "/postpaid/payment/installment/request/account/balance", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelPostPaidSchedule cancels installment schedule
func (c *Client) CancelPostPaidSchedule(ctx context.Context, req *types.CancelInvoiceScheduleRequest) (*types.PostPaidPaymentResponse, error) {
	var response types.PostPaidPaymentResponse
	if err := c.post(ctx, "/postpaid/payment/schedule/cancel", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
