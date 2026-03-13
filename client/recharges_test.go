package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestRechargesEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "DoRecharge", method: http.MethodPost, path: "/accounts/1/recharges",
			response: &types.DoRechargeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoRecharge(ctx, 1, &types.DoRechargeRequest{})
				return err
			},
		},
		{
			name: "GetRechargeValues", method: http.MethodGet, path: "/accounts/1/recharges/availableValues/11/999999999",
			response: &types.RechargeValuesResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRechargeValues(ctx, 1, "11", "999999999")
				return err
			},
		},
		{
			name: "DoVoucherRecharge", method: http.MethodPost, path: "/accounts/1/eletronicVouchers",
			response: &types.DoRechargeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoVoucherRecharge(ctx, 1, &types.DoVoucherRechargeRequest{})
				return err
			},
		},
		{
			name: "GetVoucherProviders", method: http.MethodGet, path: "/accounts/1/eletronicVouchers/providers",
			response: &types.VoucherProvidersResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetVoucherProviders(ctx, 1)
				return err
			},
		},
	})
}
