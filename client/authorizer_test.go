package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestAuthorizerEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "DoSummaryPurchase", method: http.MethodPost, path: "/summary-purchases",
			response: &types.AuthorizationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoSummaryPurchase(ctx, &types.SummaryPurchaseRequest{})
				return err
			},
		},
		{
			name: "CancelSummaryPurchase", method: http.MethodPost, path: "/summary-purchases/cancel",
			response: &types.AuthorizationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelSummaryPurchase(ctx, &types.CancelPurchaseRequest{})
				return err
			},
		},
		{
			name: "DoSummaryChargeback", method: http.MethodPost, path: "/summary-chargebacks",
			response: &types.AuthorizationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoSummaryChargeback(ctx, &types.ChargebackRequest{})
				return err
			},
		},
		{
			name: "CancelSummaryChargeback", method: http.MethodPost, path: "/summary-chargebacks/cancel",
			response: &types.AuthorizationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelSummaryChargeback(ctx, &types.CancelChargebackRequest{})
				return err
			},
		},
	})
}
