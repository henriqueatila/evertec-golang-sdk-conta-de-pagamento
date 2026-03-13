package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestBankslipsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListBankslips", method: http.MethodGet, path: "/accounts/1/bankslip",
			response: &types.BankslipsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBankslips(ctx, 1)
				return err
			},
		},
		{
			name: "CreateBankslip", method: http.MethodPost, path: "/accounts/1/bankslip",
			response: &types.CreateBankslipResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBankslip(ctx, 1, &types.CreateBankslipRequest{})
				return err
			},
		},
		{
			name: "ListBankslipsByStatus", method: http.MethodGet, path: "/accounts/1/bankslip/pending",
			response: &types.BankslipsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBankslipsByStatus(ctx, 1, "pending")
				return err
			},
		},
		{
			name: "ListBankslipsByStatusAndDate", method: http.MethodGet, path: "/accounts/1/bankslip/pending/2024-01-01",
			response: &types.BankslipsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBankslipsByStatusAndDate(ctx, 1, "pending", "2024-01-01")
				return err
			},
		},
		{
			name: "CreateBankslipV2", method: http.MethodPost, path: "/bankslip/v2/generate",
			response: &types.BankslipV2Response{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBankslipV2(ctx, &types.BankslipV2Request{})
				return err
			},
		},
	})
}
