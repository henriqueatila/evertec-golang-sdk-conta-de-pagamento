package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestPostpaidEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "PostPaidPaymentBalance", method: http.MethodPost, path: "/postpaid/payment/balance",
			response: &types.PostPaidPaymentResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidPaymentBalance(ctx, &types.PostPaidPaymentBalanceRequest{})
				return err
			},
		},
		{
			name: "PostPaidInstallmentSimulation", method: http.MethodPost, path: "/postpaid/payment/installment/simulation",
			response: &types.PostPaidInstallmentSimulationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidInstallmentSimulation(ctx, &types.PostPaidInstallmentSimulationRequest{})
				return err
			},
		},
		{
			name: "PostPaidInstallmentPix", method: http.MethodPost, path: "/postpaid/payment/installment/request/pix",
			response: &types.PostPaidInstallmentPixResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidInstallmentPix(ctx, &types.PostPaidInstallmentRequest{})
				return err
			},
		},
		{
			name: "PostPaidInstallmentAccountBalance", method: http.MethodPost, path: "/postpaid/payment/installment/request/account/balance",
			response: &types.PostPaidInstallmentBalanceResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidInstallmentAccountBalance(ctx, &types.PostPaidInstallmentRequest{})
				return err
			},
		},
		{
			name: "CancelPostPaidSchedule", method: http.MethodPost, path: "/postpaid/payment/schedule/cancel",
			response: &types.PostPaidPaymentResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelPostPaidSchedule(ctx, &types.CancelInvoiceScheduleRequest{})
				return err
			},
		},
	})
}
