package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestCreditsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "UpdateCreditExpiration", method: http.MethodPut, path: "/accounts/1/credits/expiration",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateCreditExpiration(ctx, 1, &types.UpdateCreditExpirationRequest{})
				return err
			},
		},
		{
			name: "GetUsableCredits", method: http.MethodGet, path: "/accounts/1/credits/to/use",
			response: &types.CreditsInfoResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetUsableCredits(ctx, 1, nil)
				return err
			},
		},
		{
			name: "GetRefundableCredits", method: http.MethodGet, path: "/accounts/1/credits/refund",
			response: &types.CreditsInfoResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRefundableCredits(ctx, 1)
				return err
			},
		},
		{
			name: "GetExpiredCredits", method: http.MethodGet, path: "/accounts/1/credits/expired",
			response: &types.CreditsInfoResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetExpiredCredits(ctx, 1)
				return err
			},
		},
	})
}
