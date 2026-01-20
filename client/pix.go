package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// PIX Key Management

// CreatePixKey creates a new PIX key for the account
func (c *Client) CreatePixKey(ctx context.Context, accountID int64, req *types.CreatePixKeyRequest) (*types.PixKeyResponse, error) {
	path := fmt.Sprintf("/accounts/%d/createKey", accountID)
	var response types.PixKeyResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeletePixKey deletes a PIX key from the account
func (c *Client) DeletePixKey(ctx context.Context, accountID int64, req *types.DeletePixKeyRequest) error {
	path := fmt.Sprintf("/accounts/%d/deleteKey", accountID)
	return c.post(ctx, path, req, nil)
}

// GetPixKeys retrieves all PIX keys for the account
func (c *Client) GetPixKeys(ctx context.Context, accountID int64) (*types.PixKeyListResponse, error) {
	path := fmt.Sprintf("/accounts/%d/getKeys", accountID)
	var response types.PixKeyListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PIX Claims (Portability/Ownership)

// CreatePixClaim creates a new PIX key claim (portability or ownership)
func (c *Client) CreatePixClaim(ctx context.Context, accountID int64, req *types.CreatePixClaimRequest) (*types.PixClaimResponse, error) {
	path := fmt.Sprintf("/accounts/%d/createClaim", accountID)
	var response types.PixClaimResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ConfirmPortability confirms a PIX key portability claim
func (c *Client) ConfirmPortability(ctx context.Context, accountID int64, req *types.ConfirmPortabilityRequest) (*types.PixClaimResponse, error) {
	path := fmt.Sprintf("/accounts/%d/confirmPortability", accountID)
	var response types.PixClaimResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CompletePortability completes a PIX key portability claim
func (c *Client) CompletePortability(ctx context.Context, accountID int64, req *types.CompletePortabilityRequest) (*types.PixClaimResponse, error) {
	path := fmt.Sprintf("/accounts/%d/completePortability", accountID)
	var response types.PixClaimResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CancelPortability cancels a PIX key portability claim
func (c *Client) CancelPortability(ctx context.Context, accountID int64, req *types.CancelPortabilityRequest) (*types.PixClaimResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cancelPortability", accountID)
	var response types.PixClaimResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRequestedClaims retrieves all PIX key claims for the account
func (c *Client) GetRequestedClaims(ctx context.Context, accountID int64) (*types.PixClaimListResponse, error) {
	path := fmt.Sprintf("/accounts/%d/getRequestedClaims", accountID)
	var response types.PixClaimListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PIX Limits

// GetPixLimit retrieves PIX transaction limits for the account
func (c *Client) GetPixLimit(ctx context.Context, accountID int64) (*types.PixLimitResponse, error) {
	path := fmt.Sprintf("/accounts/%d/pix/getLimit", accountID)
	var response types.PixLimitResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePixLimit updates PIX transaction limits for the account
func (c *Client) UpdatePixLimit(ctx context.Context, accountID int64, req *types.PixLimitRequest) (*types.PixLimitResponse, error) {
	path := fmt.Sprintf("/accounts/%d/pix/limit", accountID)
	var response types.PixLimitResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePixNightTimeLimit updates the night-time PIX limit start/end times
func (c *Client) UpdatePixNightTimeLimit(ctx context.Context, accountID int64, req *types.UpdatePixNightTimeLimitRequest) (*types.PixLimitResponse, error) {
	path := fmt.Sprintf("/accounts/%d/pix/limit/startNightTime", accountID)
	var response types.PixLimitResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PIX Devices

// AddPixDevice registers a trusted PIX device
func (c *Client) AddPixDevice(ctx context.Context, req *types.PixDeviceRequest) (*types.PixDeviceResponse, error) {
	var response types.PixDeviceResponse
	if err := c.post(ctx, "/pix/devices", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeletePixDevice removes a trusted PIX device
func (c *Client) DeletePixDevice(ctx context.Context, req *types.DeletePixDeviceRequest) error {
	return c.deleteWithBody(ctx, "/pix/devices", req, nil)
}

// BlockPixDevice blocks a trusted PIX device
func (c *Client) BlockPixDevice(ctx context.Context, req *types.BlockPixDeviceRequest) (*types.PixDeviceResponse, error) {
	var response types.PixDeviceResponse
	if err := c.put(ctx, "/pix/devices/block", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UnblockPixDevice unblocks a trusted PIX device
func (c *Client) UnblockPixDevice(ctx context.Context, req *types.UnblockPixDeviceRequest) (*types.PixDeviceResponse, error) {
	var response types.PixDeviceResponse
	if err := c.put(ctx, "/pix/devices/unblock", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListPixDevices retrieves all trusted PIX devices for the account
func (c *Client) ListPixDevices(ctx context.Context, accountID int64) (*types.PixDeviceListResponse, error) {
	path := fmt.Sprintf("/pix/devices/list/%d", accountID)
	var response types.PixDeviceListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PIX Global Limit Operations

// ListPixClaims lists all PIX claims with optional filters
func (c *Client) ListPixClaims(ctx context.Context, req *types.ListPixClaimsRequest) (*types.PixClaimListResponse, error) {
	var response types.PixClaimListResponse
	if err := c.post(ctx, "/accounts/pix/claim/list", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateClaimFromKey creates a PIX claim from an existing key
func (c *Client) CreateClaimFromKey(ctx context.Context, req *types.CreateClaimFromKeyRequest) (*types.PixClaimResponse, error) {
	var response types.PixClaimResponse
	if err := c.post(ctx, "/accounts/pix/claim/key/createClaim", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ProcessLimitRequest processes a PIX limit raise request (approve/reject)
func (c *Client) ProcessLimitRequest(ctx context.Context, req *types.ProcessLimitRequestData) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.put(ctx, "/accounts/pix/limit/processLimitRequest", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRaiseLimitRequests retrieves all PIX raise limit requests
func (c *Client) GetRaiseLimitRequests(ctx context.Context) (*types.RaiseLimitRequestListResponse, error) {
	var response types.RaiseLimitRequestListResponse
	if err := c.get(ctx, "/accounts/pix/limit/getRaiseLimitRequests", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetMaximumPixLimitIssuer retrieves the maximum PIX limit allowed by issuer
func (c *Client) GetMaximumPixLimitIssuer(ctx context.Context) (*types.MaximumPixLimitIssuerResponse, error) {
	var response types.MaximumPixLimitIssuerResponse
	if err := c.get(ctx, "/accounts/pix/limit/getMaximumLimitIssuer", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRaiseLimitRequestDetail retrieves details of a specific raise limit request
func (c *Client) GetRaiseLimitRequestDetail(ctx context.Context, requestID int64) (*types.RaiseLimitRequestResponse, error) {
	path := fmt.Sprintf("/accounts/pix/limit/getDetailRaiseLimitRequest/%d", requestID)
	var response types.RaiseLimitRequestResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReceivePixCallback processes a PIX callback/notification
func (c *Client) ReceivePixCallback(ctx context.Context, req *types.PixCallbackRequest) (*types.PixCallbackResponse, error) {
	var response types.PixCallbackResponse
	if err := c.post(ctx, "/pix/callbacks/receive-transaction", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
