package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestIncomeReportEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GenerateIncomeReport", method: http.MethodGet, path: "/accounts/1/issuer-report/2024",
			response: &types.IncomeReportResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GenerateIncomeReport(ctx, 1, 2024)
				return err
			},
		},
		{
			name: "GetAccountBalanceByYear", method: http.MethodGet, path: "/accounts/1/balance/2024",
			response: &types.YearlyBalanceResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountBalanceByYear(ctx, 1, 2024)
				return err
			},
		},
		{
			name: "GetAllAccountsBalanceByYear", method: http.MethodGet, path: "/accounts/balance/2024",
			response: &types.AllAccountsYearlyBalanceResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAllAccountsBalanceByYear(ctx, 2024)
				return err
			},
		},
	})
}
