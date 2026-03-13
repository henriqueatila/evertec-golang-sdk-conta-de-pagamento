package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestMedEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListInfractionReports", method: http.MethodGet, path: "/pix/infraction-reports",
			response: &types.ListInfractionReportsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInfractionReports(ctx, nil)
				return err
			},
		},
		{
			name: "CreateInfractionReport", method: http.MethodPost, path: "/pix/infraction-reports",
			response: &types.InfractionReportResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateInfractionReport(ctx, &types.InfractionReportRequest{})
				return err
			},
		},
		{
			name: "CloseInfractionReport", method: http.MethodPost, path: "/pix/infraction-reports/close",
			response: &types.InfractionReportResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CloseInfractionReport(ctx, &types.CloseInfractionReportRequest{})
				return err
			},
		},
		{
			name: "GetInfractionReport", method: http.MethodGet, path: "/pix/infraction-reports/report-abc",
			response: &types.InfractionReportResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetInfractionReport(ctx, "report-abc")
				return err
			},
		},
		{
			name: "CancelInfractionReport", method: http.MethodPut, path: "/pix/infraction-reports/cancel/report-abc",
			response: &types.InfractionReportResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelInfractionReport(ctx, "report-abc")
				return err
			},
		},
		{
			name: "CreateRefundSolicitation", method: http.MethodPost, path: "/pix/refunds/create",
			response: &types.RefundResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateRefundSolicitation(ctx, &types.RefundSolicitationRequest{})
				return err
			},
		},
		{
			name: "CloseRefundSolicitation", method: http.MethodPost, path: "/pix/refunds/close",
			response: &types.RefundResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CloseRefundSolicitation(ctx, &types.CloseRefundRequest{})
				return err
			},
		},
		{
			name: "ListRefundSolicitations", method: http.MethodGet, path: "/pix/refunds",
			response: &types.ListRefundsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListRefundSolicitations(ctx, nil)
				return err
			},
		},
		{
			name: "GetRefundSolicitation", method: http.MethodGet, path: "/pix/refunds/refund-abc",
			response: &types.RefundResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRefundSolicitation(ctx, "refund-abc")
				return err
			},
		},
		{
			name: "CancelRefundSolicitation", method: http.MethodPut, path: "/pix/refunds/refund-abc/cancel",
			response: &types.RefundResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelRefundSolicitation(ctx, "refund-abc")
				return err
			},
		},
	})
}

func TestMedDeprecatedAliases(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CloseRefund", method: http.MethodPost, path: "/pix/refunds/close",
			response: &types.RefundResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CloseRefund(ctx, &types.CloseRefundRequest{})
				return err
			},
		},
		{
			name: "GetRefund", method: http.MethodGet, path: "/pix/refunds/refund-xyz",
			response: &types.RefundResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRefund(ctx, "refund-xyz")
				return err
			},
		},
		{
			name: "CancelRefund", method: http.MethodPut, path: "/pix/refunds/refund-xyz/cancel",
			response: &types.RefundResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelRefund(ctx, "refund-xyz")
				return err
			},
		},
	})
}
