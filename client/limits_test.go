package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestLimitsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetAccountLimit", method: http.MethodGet, path: "/accounts/limit/1/pix/getLimit",
			response: &types.LimitResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountLimit(ctx, 1, types.LimitType("pix"))
				return err
			},
		},
		{
			name: "UpdateAccountLimit", method: http.MethodPut, path: "/accounts/limit/1/pix",
			response: &types.LimitResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAccountLimit(ctx, 1, types.LimitType("pix"), &types.UpdateLimitRequest{})
				return err
			},
		},
		{
			name: "UpdateAccountNightTimeLimit", method: http.MethodPut, path: "/accounts/limit/1/pix/startNightTime",
			response: &types.LimitResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAccountNightTimeLimit(ctx, 1, types.LimitType("pix"), &types.UpdateNightTimeLimitRequest{})
				return err
			},
		},
		{
			name: "GetMaximumLimitIssuer", method: http.MethodGet, path: "/accounts/limit/1/pix/getMaximumLimitIssuer",
			response: &types.MaximumLimitResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetMaximumLimitIssuer(ctx, 1, types.LimitType("pix"))
				return err
			},
		},
		{
			name: "GetAccountFees", method: http.MethodGet, path: "/accounts/1/fees",
			response: &types.ContaDigitalGenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountFees(ctx, 1)
				return err
			},
		},
		{
			name: "GetCardIssuanceFee", method: http.MethodGet, path: "/accounts/1/fees/cardissuer",
			response: &types.ContaDigitalGenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardIssuanceFee(ctx, 1)
				return err
			},
		},
		{
			name: "GetCardReissueFee", method: http.MethodGet, path: "/accounts/1/fees/cardreissue",
			response: &types.ContaDigitalGenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardReissueFee(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateProductLimitByType", method: http.MethodPut, path: "/limit/pix/productLimit",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateProductLimitByType(ctx, types.LimitType("pix"), &types.ProductLimitRequest{})
			},
		},
		{
			name: "SearchProductLimitByType", method: http.MethodPost, path: "/limit/pix/searchProductLimit",
			response: []types.ProductLimitResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SearchProductLimitByType(ctx, types.LimitType("pix"), &types.SearchProductLimitRequest{})
				return err
			},
		},
	})
}
