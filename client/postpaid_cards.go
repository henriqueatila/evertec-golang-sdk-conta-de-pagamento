package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Post-paid Card Operations

// GetPostPaidVirtualCards retrieves virtual post-paid cards
func (c *Client) GetPostPaidVirtualCards(ctx context.Context, accountID int64) (*types.VirtualCardsResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/virtual", accountID)
	var response types.VirtualCardsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidPhysicalCards retrieves physical post-paid cards
func (c *Client) GetPostPaidPhysicalCards(ctx context.Context, accountID int64) (*types.AccountCardsResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/physical", accountID)
	var response types.AccountCardsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePostPaidCard creates a new post-paid card
func (c *Client) CreatePostPaidCard(ctx context.Context, req *types.CreateCardRequest) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.post(ctx, "/postpaid/cards/new", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePostPaidVirtualCard creates a virtual post-paid card for account
func (c *Client) CreatePostPaidVirtualCard(ctx context.Context, accountID int64, req *types.CreateVirtualCardRequest) (*types.VirtualCardResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/new/virtual", accountID)
	var response types.VirtualCardResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BlockPostPaidCard blocks a post-paid card
func (c *Client) BlockPostPaidCard(ctx context.Context, accountID, cardID int64, req *types.BlockCardRequest) (*types.BlockCardResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/block/%d", accountID, cardID)
	var response types.BlockCardResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UnblockPostPaidCard unblocks a post-paid card
func (c *Client) UnblockPostPaidCard(ctx context.Context, accountID, cardID int64) (*types.UnblockCardResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/unblock/%d", accountID, cardID)
	var response types.UnblockCardResponse
	if err := c.post(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ActivatePostPaidCard activates a post-paid card
func (c *Client) ActivatePostPaidCard(ctx context.Context, accountID, cardID int64, req *types.ActivateCardRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/activate/%d", accountID, cardID)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ChangePostPaidCardPin changes a post-paid card PIN
func (c *Client) ChangePostPaidCardPin(ctx context.Context, accountID, cardID int64, req *types.ChangeCardPinRequest) (*types.ChangeCardPinResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/changePin/%d", accountID, cardID)
	var response types.ChangeCardPinResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ValidatePostPaidCardPin validates and changes post-paid card PIN
func (c *Client) ValidatePostPaidCardPin(ctx context.Context, accountID, cardID int64, req *types.ChangeCardPinRequest) (*types.ChangeCardPinResponse, error) {
	path := fmt.Sprintf("/postpaid/cards/%d/changeValidatePin/%d", accountID, cardID)
	var response types.ChangeCardPinResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidCardSettings retrieves card settings
// Returns PostpaidCardSettingsResponse from types
func (c *Client) GetPostPaidCardSettings(ctx context.Context, accountID, cardID int64) (*types.PostpaidCardSettingsResponse, error) {
	path := fmt.Sprintf("/postpaid/account/%d/cards/%d/settings", accountID, cardID)
	var response types.PostpaidCardSettingsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePostPaidCardSettings updates card settings
// Uses PostpaidCardSettingsRequest from types
func (c *Client) UpdatePostPaidCardSettings(ctx context.Context, accountID, cardID int64, req *types.PostpaidCardSettingsRequest) (*types.PostpaidCardSettingsResponse, error) {
	path := fmt.Sprintf("/postpaid/account/%d/cards/%d/settings", accountID, cardID)
	var response types.PostpaidCardSettingsResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ResetPostPaidCardSettings resets card settings to defaults
func (c *Client) ResetPostPaidCardSettings(ctx context.Context, accountID, cardID int64) (*types.PostpaidCardSettingsResponse, error) {
	path := fmt.Sprintf("/postpaid/account/%d/cards/%d/settings/reset", accountID, cardID)
	var response types.PostpaidCardSettingsResponse
	if err := c.patch(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
