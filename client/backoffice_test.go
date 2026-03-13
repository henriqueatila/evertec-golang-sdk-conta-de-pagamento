package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestBackofficeEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListAccountsBackoffice", method: http.MethodGet, path: "/backoffice/accounts",
			response: &types.AccountListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAccountsBackoffice(ctx, nil)
				return err
			},
		},
		{
			name: "ProcessProposalManually", method: http.MethodPost, path: "/backoffice/proposals/process",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ProcessProposalManually(ctx, &types.ProposalProcessingRequest{})
				return err
			},
		},
		{
			name: "CreateMobileAccount", method: http.MethodPut, path: "/backoffice/accounts/createMobileAccount",
			response: &types.CreateAccountResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateMobileAccount(ctx, &types.CreateMobileAccountRequest{})
				return err
			},
		},
		{
			name: "CreateBiroAnalysis", method: http.MethodPost, path: "/backoffice/biro/analysis",
			response: &types.BiroAnalysisResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBiroAnalysis(ctx, &types.BiroAnalysisRequest{})
				return err
			},
		},
		{
			name: "GetBiroAnalysis", method: http.MethodGet, path: "/backoffice/biro/analysis/1",
			response: &types.BiroAnalysisResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBiroAnalysis(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateBiroAnalysis", method: http.MethodPut, path: "/backoffice/biro/analysis/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateBiroAnalysis(ctx, 1, &types.UpdateBiroAnalysisRequest{})
			},
		},
		{
			name: "BindProcessorAccount", method: http.MethodPost, path: "/backoffice/processor/account/bind",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BindProcessorAccount(ctx, &types.BindProcessorAccountRequest{})
				return err
			},
		},
		{
			name: "BindProcessorCard", method: http.MethodPut, path: "/backoffice/processor/account/card",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BindProcessorCard(ctx, &types.BindProcessorCardRequest{})
				return err
			},
		},
		{
			name: "SyncProcessorAccount", method: http.MethodPut, path: "/backoffice/processor/account/1/synchronize",
			response: &types.SyncProcessorResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SyncProcessorAccount(ctx, 1)
				return err
			},
		},
		{
			name: "GetPixScanConfiguration", method: http.MethodGet, path: "/backoffice/pix/scan/configuration",
			response: &types.PixScanConfigurationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPixScanConfiguration(ctx)
				return err
			},
		},
		{
			name: "UpdatePixScanConfiguration", method: http.MethodPut, path: "/backoffice/pix/scan/configuration",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdatePixScanConfiguration(ctx, &types.UpdatePixScanConfigurationRequest{})
			},
		},
		{
			name: "ListHceDevices", method: http.MethodGet, path: "/backoffice/hce/devices",
			response: []types.HceDeviceResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListHceDevices(ctx, nil)
				return err
			},
		},
		{
			name: "GetHceDevice", method: http.MethodGet, path: "/backoffice/hce/devices/dev-123",
			response: &types.HceDeviceResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetHceDevice(ctx, "dev-123")
				return err
			},
		},
		{
			name: "BlockHceDevice", method: http.MethodPost, path: "/backoffice/hce/devices/dev-123/block",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.BlockHceDevice(ctx, "dev-123")
			},
		},
		{
			name: "UnblockHceDevice", method: http.MethodPost, path: "/backoffice/hce/devices/dev-123/unblock",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UnblockHceDevice(ctx, "dev-123")
			},
		},
		{
			name: "GetDailyStatement", method: http.MethodGet, path: "/backoffice/statements/daily/2024-01-01",
			response: &types.DailyStatementResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetDailyStatement(ctx, "2024-01-01")
				return err
			},
		},
		{
			name: "GetIssuerBalance", method: http.MethodGet, path: "/backoffice/issuer/balance",
			response: &types.IssuerBalanceResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetIssuerBalance(ctx)
				return err
			},
		},
		{
			name: "ResetAccountLoginTime", method: http.MethodPut, path: "/backoffice/accounts/resetLoginTime/1",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ResetAccountLoginTime(ctx, 1)
				return err
			},
		},
		{
			name: "SyncProcessorCard", method: http.MethodPut, path: "/backoffice/processor/card/1/synchronize",
			response: &types.SyncProcessorResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SyncProcessorCard(ctx, 1)
				return err
			},
		},
		{
			name: "ListHceOverview", method: http.MethodGet, path: "/backoffice/hce",
			response: []types.HceDeviceResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListHceOverview(ctx)
				return err
			},
		},
		{
			name: "ListDailyStatements", method: http.MethodGet, path: "/dailyStatement/list",
			response: &types.DailyStatementListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListDailyStatements(ctx)
				return err
			},
		},
		{
			name: "DeleteAllDailyStatements", method: http.MethodDelete, path: "/dailyStatement/all",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DeleteAllDailyStatements(ctx)
				return err
			},
		},
		{
			name: "ListBiroAnalyses", method: http.MethodGet, path: "/backoffice/accounts/biroAnalysis",
			response: []types.BiroAnalysisResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBiroAnalyses(ctx)
				return err
			},
		},
		{
			name: "GetBiroAnalysisByProposal", method: http.MethodGet, path: "/backoffice/accounts/biroAnalysis/proposal/1",
			response: &types.BiroAnalysisResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBiroAnalysisByProposal(ctx, 1)
				return err
			},
		},
		{
			name: "DeletePaysmartProduct", method: http.MethodDelete, path: "/backoffice/product/paysmart/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.DeletePaysmartProduct(ctx, 1)
			},
		},
	})
}
