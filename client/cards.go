package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// ListCards retrieves cards for an account
func (c *Client) ListCards(ctx context.Context, accountID int64, params *types.ListCardsParams) (*types.AccountCardsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards", accountID)
	if params != nil {
		path += params.QueryString()
	}
	var response types.AccountCardsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCard retrieves a specific card
func (c *Client) GetCard(ctx context.Context, accountID, cardID int64) (*types.CardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d", accountID, cardID)
	var response types.CardResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateCard creates a new card for an account
func (c *Client) CreateCard(ctx context.Context, accountID int64, req *types.CreateCardRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/new", accountID)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateCardBackoffice creates a new card via backoffice
func (c *Client) CreateCardBackoffice(ctx context.Context, accountID int64, req *types.CreateCardRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/new/backoffice", accountID)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BlockCard blocks a card
func (c *Client) BlockCard(ctx context.Context, accountID, cardID int64, req *types.BlockCardRequest) (*types.BlockCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/block", accountID, cardID)
	var response types.BlockCardResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UnblockCard unblocks a card
func (c *Client) UnblockCard(ctx context.Context, accountID, cardID int64) (*types.UnblockCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/unblock", accountID, cardID)
	var response types.UnblockCardResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReissueCard reissues a card
func (c *Client) ReissueCard(ctx context.Context, accountID, cardID int64, req *types.BlockCardRequest) (*types.BlockCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/reissue", accountID, cardID)
	var response types.BlockCardResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ReissueCardBackoffice reissues a card via backoffice
func (c *Client) ReissueCardBackoffice(ctx context.Context, accountID, cardID int64, req *types.BlockCardRequest) (*types.BlockCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/reissue/backoffice", accountID, cardID)
	var response types.BlockCardResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ActivateCard activates a card
func (c *Client) ActivateCard(ctx context.Context, accountID, cardID int64, req *types.ActivateCardRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/activate", accountID, cardID)
	var response types.GenericResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ChangeCardPin changes the card PIN
func (c *Client) ChangeCardPin(ctx context.Context, accountID, cardID int64, req *types.ChangeCardPinRequest) (*types.ChangeCardPinResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/changePin", accountID, cardID)
	var response types.ChangeCardPinResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateVirtualCardTag updates the tag of a virtual card
func (c *Client) UpdateVirtualCardTag(ctx context.Context, accountID, cardID int64, req *types.UpdateVirtualCardTagRequest) error {
	path := fmt.Sprintf("/accounts/%d/cards/%d/tag", accountID, cardID)
	return c.put(ctx, path, req, nil)
}

// CreateVirtualCardFromPhysical creates a virtual card from a physical card
func (c *Client) CreateVirtualCardFromPhysical(ctx context.Context, accountID, cardID int64) (*types.VirtualCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/virtual", accountID, cardID)
	var response types.VirtualCardResponse
	if err := c.post(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetVirtualCards retrieves virtual cards for a physical card
func (c *Client) GetVirtualCards(ctx context.Context, accountID, cardID int64) (*types.VirtualCardsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/virtual", accountID, cardID)
	var response types.VirtualCardsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateVirtualCard creates a new virtual card
func (c *Client) CreateVirtualCard(ctx context.Context, accountID int64, req *types.CreateVirtualCardRequest) (*types.VirtualCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/virtual", accountID)
	var response types.VirtualCardResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAllVirtualCards retrieves all virtual cards for an account
func (c *Client) ListAllVirtualCards(ctx context.Context, accountID int64) (*types.VirtualCardsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/virtual", accountID)
	var response types.VirtualCardsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCardReplacementInfo retrieves card replacement information
func (c *Client) GetCardReplacementInfo(ctx context.Context, accountID, cardID int64) (*types.ReplacementCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/replacement", accountID, cardID)
	var response types.ReplacementCardResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// RequestCardReplacement requests a card replacement
func (c *Client) RequestCardReplacement(ctx context.Context, accountID, cardID int64) (*types.ReplacementCardResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/replacement", accountID, cardID)
	var response types.ReplacementCardResponse
	if err := c.post(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BindAnonymousCard binds an anonymous card to an account
func (c *Client) BindAnonymousCard(ctx context.Context, accountID int64, req *types.BindAnonymousCardRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/bindAnonymousCard", accountID)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SearchCards searches all cards with optional filters
func (c *Client) SearchCards(ctx context.Context, params *types.SearchCardsParams) (*types.AccountCardsResponse, error) {
	path := "/cards"
	if params != nil {
		path += params.QueryString()
	}
	var response types.AccountCardsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCardConfiguration retrieves configuration for a specific card
func (c *Client) GetCardConfiguration(ctx context.Context, accountID, cardID int64) (*types.CardConfigurationResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/%d/getCardConfiguration", accountID, cardID)
	var response types.CardConfigurationResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetDefaultCardConfiguration retrieves the default card configuration for an account
func (c *Client) GetDefaultCardConfiguration(ctx context.Context, accountID int64) (*types.DefaultCardConfigurationResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/getDefaultCardConfiguration", accountID)
	var response types.DefaultCardConfigurationResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateDefaultCardConfiguration updates the default card configuration for an account
func (c *Client) UpdateDefaultCardConfiguration(ctx context.Context, accountID int64, req *types.CardConfigurationRequest) error {
	path := fmt.Sprintf("/accounts/%d/cards/updateDefaultCardConfiguration", accountID)
	return c.put(ctx, path, req, nil)
}

// ConfigureCard configures a specific card
func (c *Client) ConfigureCard(ctx context.Context, accountID int64, req *types.CardConfigurationRequest) error {
	path := fmt.Sprintf("/accounts/%d/cards/configureCard", accountID)
	return c.put(ctx, path, req, nil)
}

// GetCardPaysmart retrieves a Paysmart card by its Paysmart ID
func (c *Client) GetCardPaysmart(ctx context.Context, accountID int64, cardIDPaysmart string) (*types.CardPaysmartResponse, error) {
	path := fmt.Sprintf("/accounts/%d/cards/paysmart/%s", accountID, cardIDPaysmart)
	var response types.CardPaysmartResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
