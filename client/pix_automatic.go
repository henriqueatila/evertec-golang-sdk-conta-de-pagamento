package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// PIX Automático (Recorrência) Operations
// Based on OpenAPI spec: openapi-pix-automatico.json

// StartAutomaticPix starts a new automatic/recurring PIX
func (c *Client) StartAutomaticPix(ctx context.Context, req *types.StartAutomaticPixRequest) (*types.StartAutomaticPixResponse, error) {
	var response types.StartAutomaticPixResponse
	if err := c.post(ctx, "/pix/automatic", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// RejectAutomaticPix rejects an automatic PIX request
func (c *Client) RejectAutomaticPix(ctx context.Context, req *types.RejectAutomaticPixRequest) (*types.RejectAutomaticPixResponse, error) {
	var response types.RejectAutomaticPixResponse
	if err := c.post(ctx, "/pix/automatic/reject", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// AcceptQRCodeJourneyThree accepts QR code in journey three flow with PIX payment
func (c *Client) AcceptQRCodeJourneyThree(ctx context.Context, req *types.QRCodeAcceptJourneyThreeRequest) (*types.QRCodeAcceptJourneyThreeResponse, error) {
	var response types.QRCodeAcceptJourneyThreeResponse
	if err := c.post(ctx, "/pix/automatic/qr-code/journey-three/accept", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// AcceptAutomaticPixQRCode accepts automatic PIX via QR code
func (c *Client) AcceptAutomaticPixQRCode(ctx context.Context, req *types.QRCodeUserAcceptRequest) (*types.QRCodeUserResponse, error) {
	var response types.QRCodeUserResponse
	if err := c.post(ctx, "/pix/automatic/qr-code/accept", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateAutomaticPixContract creates an automatic PIX contract
func (c *Client) CreateAutomaticPixContract(ctx context.Context, req *types.CreateAutomaticPixContractRequest) (*types.CreateAutomaticPixContractResponse, error) {
	var response types.CreateAutomaticPixContractResponse
	if err := c.post(ctx, "/pix/automatic/contract", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelAutomaticPixCharge cancels an automatic PIX charge
func (c *Client) CancelAutomaticPixCharge(ctx context.Context, req *types.CancelAutomaticPixChargeRequest) (*types.CancelAutomaticPixChargeResponse, error) {
	var response types.CancelAutomaticPixChargeResponse
	if err := c.post(ctx, "/pix/automatic/charge/cancel", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelAutomaticPix cancels an automatic PIX subscription
func (c *Client) CancelAutomaticPix(ctx context.Context, req *types.CancelAutomaticPixRequest) (*types.CancelAutomaticPixResponse, error) {
	var response types.CancelAutomaticPixResponse
	if err := c.post(ctx, "/pix/automatic/cancel", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// AcceptAutomaticPix accepts an automatic PIX request
func (c *Client) AcceptAutomaticPix(ctx context.Context, req *types.AcceptAutomaticPixRequest) (*types.AutomaticPixResponse, error) {
	var response types.AutomaticPixResponse
	if err := c.post(ctx, "/pix/automatic/accept", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAutomaticPixCharges lists automatic PIX charges for an account
func (c *Client) ListAutomaticPixCharges(ctx context.Context, accountID int64, params *types.ListAutomaticPixParams) (*types.AutomaticPixChargeListResponse, error) {
	path := fmt.Sprintf("/pix/automatic/charge/account/%d%s", accountID, params.QueryString())
	var response types.AutomaticPixChargeListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAutomaticPixByAccount lists automatic PIX subscriptions for an account
func (c *Client) ListAutomaticPixByAccount(ctx context.Context, accountID int64, params *types.ListAutomaticPixParams) (*types.AutomaticPixListResponse, error) {
	path := fmt.Sprintf("/pix/automatic/account/%d%s", accountID, params.QueryString())
	var response types.AutomaticPixListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAutomaticPixRecurrence retrieves a specific automatic PIX recurrence
func (c *Client) GetAutomaticPixRecurrence(ctx context.Context, accountID int64, recurrenceID string, isPayer *bool) (*types.AutomaticPixResponse, error) {
	path := fmt.Sprintf("/pix/automatic/account/%d/recurrence/%s", accountID, recurrenceID)
	if isPayer != nil {
		path += fmt.Sprintf("?isPayer=%t", *isPayer)
	}
	var response types.AutomaticPixResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
