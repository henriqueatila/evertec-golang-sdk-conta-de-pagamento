package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// MED (Mecanismo Especial de Devolução) Operations
// Required by Brazilian Central Bank for PIX fraud reporting
// Based on OpenAPI spec: openapi-med.json

// Infraction Reports

// ListInfractionReports lists all infraction reports with optional filters
func (c *Client) ListInfractionReports(ctx context.Context, params *types.ListInfractionReportsParams) (*types.ListInfractionReportsResponse, error) {
	path := "/pix/infraction-reports"
	if params != nil {
		path += params.QueryString()
	}
	var response types.ListInfractionReportsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateInfractionReport creates a new infraction report
func (c *Client) CreateInfractionReport(ctx context.Context, req *types.InfractionReportRequest) (*types.InfractionReportResponse, error) {
	var response types.InfractionReportResponse
	if err := c.post(ctx, "/pix/infraction-reports", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CloseInfractionReport closes an infraction report with analysis
func (c *Client) CloseInfractionReport(ctx context.Context, req *types.CloseInfractionReportRequest) (*types.InfractionReportResponse, error) {
	var response types.InfractionReportResponse
	if err := c.post(ctx, "/pix/infraction-reports/close", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInfractionReport retrieves an infraction report by ID
func (c *Client) GetInfractionReport(ctx context.Context, infractionReportID string) (*types.InfractionReportResponse, error) {
	path := fmt.Sprintf("/pix/infraction-reports/%s", infractionReportID)
	var response types.InfractionReportResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelInfractionReport cancels an infraction report
func (c *Client) CancelInfractionReport(ctx context.Context, infractionReportID string) (*types.InfractionReportResponse, error) {
	path := fmt.Sprintf("/pix/infraction-reports/cancel/%s", infractionReportID)
	var response types.InfractionReportResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Refund Solicitations

// CreateRefundSolicitation creates a new refund solicitation
func (c *Client) CreateRefundSolicitation(ctx context.Context, req *types.RefundSolicitationRequest) (*types.RefundResponse, error) {
	var response types.RefundResponse
	if err := c.post(ctx, "/pix/refunds/create", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CloseRefundSolicitation closes a refund with result
func (c *Client) CloseRefundSolicitation(ctx context.Context, req *types.CloseRefundRequest) (*types.RefundResponse, error) {
	var response types.RefundResponse
	if err := c.post(ctx, "/pix/refunds/close", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListRefundSolicitations lists all refund solicitations with optional filters
func (c *Client) ListRefundSolicitations(ctx context.Context, params *types.ListRefundsParams) (*types.ListRefundsResponse, error) {
	path := "/pix/refunds"
	if params != nil {
		path += params.QueryString()
	}
	var response types.ListRefundsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRefundSolicitation retrieves a refund by ID
func (c *Client) GetRefundSolicitation(ctx context.Context, refundID string) (*types.RefundResponse, error) {
	path := fmt.Sprintf("/pix/refunds/%s", refundID)
	var response types.RefundResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelRefundSolicitation cancels a refund
func (c *Client) CancelRefundSolicitation(ctx context.Context, refundID string) (*types.RefundResponse, error) {
	path := fmt.Sprintf("/pix/refunds/%s/cancel", refundID)
	var response types.RefundResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Deprecated: Use CloseRefundSolicitation instead
func (c *Client) CloseRefund(ctx context.Context, req *types.CloseRefundRequest) (*types.RefundResponse, error) {
	return c.CloseRefundSolicitation(ctx, req)
}

// Deprecated: Use GetRefundSolicitation instead
func (c *Client) GetRefund(ctx context.Context, refundID string) (*types.RefundResponse, error) {
	return c.GetRefundSolicitation(ctx, refundID)
}

// Deprecated: Use CancelRefundSolicitation instead
func (c *Client) CancelRefund(ctx context.Context, refundID string) (*types.RefundResponse, error) {
	return c.CancelRefundSolicitation(ctx, refundID)
}
