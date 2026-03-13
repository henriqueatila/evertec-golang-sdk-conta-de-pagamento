package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestPostpaidStatementsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetPostPaidAccount", method: http.MethodGet, path: "/postpaid/account/1",
			response: &types.AccountPostPaidResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidAccount(ctx, 1)
				return err
			},
		},
		{
			name: "UpdatePostPaidAccountInfo", method: http.MethodPost, path: "/postpaid/account",
			response: &types.AccountPostPaidResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdatePostPaidAccountInfo(ctx, &types.UpdatePostPaidAccountRequest{})
				return err
			},
		},
		{
			name: "GetPostPaidDueDates", method: http.MethodGet, path: "/postpaid/account/1/dueDates",
			response: []types.DueDateOption{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidDueDates(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidCardDueDates", method: http.MethodGet, path: "/postpaid/account/1/cards/100/dueDates",
			response: []types.DueDateOption{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidCardDueDates(ctx, 1, 100)
				return err
			},
		},
		{
			name: "GetPostPaidStatementByMonth", method: http.MethodGet, path: "/postpaid/statements/1",
			response: &types.PaysmartOpenStatementResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidStatementByMonth(ctx, 1, 1, 2024)
				return err
			},
		},
		{
			name: "GetPostPaidOpenStatement", method: http.MethodGet, path: "/postpaid/statements/1/open-statement",
			response: &types.PaysmartOpenStatementResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidOpenStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidClosedStatement", method: http.MethodGet, path: "/postpaid/statements/1/closed-statement",
			response: &types.PaysmartOpenStatementResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidClosedStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidFutureStatement", method: http.MethodGet, path: "/postpaid/statements/1/future-statement",
			response: &types.PaysmartOpenStatementResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidFutureStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidCombinedStatement", method: http.MethodGet, path: "/postpaid/statements/1/combined-statement",
			response: &types.PaysmartOpenStatementResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidCombinedStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidTransactions", method: http.MethodGet, path: "/postpaid/statements/1/transactions",
			response: []types.TransactionsDTO{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidTransactions(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidPossibleAdvances", method: http.MethodGet, path: "/postpaid/statements/1/possible-advance",
			response: []types.TransactionsDTO{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidPossibleAdvances(ctx, 1)
				return err
			},
		},
		{
			name: "SendPostPaidStatementEmail", method: http.MethodPost, path: "/postpaid/statements/1/mail",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SendPostPaidStatementEmail(ctx, 1, "test@example.com")
				return err
			},
		},
	})
}
