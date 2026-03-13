package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestDepositsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListDepositOrders", method: http.MethodGet, path: "/accounts/1/deposits/order",
			response: &types.DepositOrdersResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListDepositOrders(ctx, 1, nil)
				return err
			},
		},
		{
			name: "CreateDepositOrder", method: http.MethodPost, path: "/accounts/1/deposits/order",
			response: &types.CreateDepositOrderResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateDepositOrder(ctx, 1, &types.CreateDepositOrderRequest{})
				return err
			},
		},
		{
			name: "ListActiveDepositOrders", method: http.MethodGet, path: "/accounts/1/deposits/order/active",
			response: &types.DepositOrdersResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListActiveDepositOrders(ctx, 1)
				return err
			},
		},
		{
			name: "CancelDepositOrder", method: http.MethodDelete, path: "/accounts/1/deposits/order/100",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelDepositOrder(ctx, 1, 100)
				return err
			},
		},
	})
}
