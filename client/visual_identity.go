package client

import (
	"context"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Visual Identity Operations (Email Configuration)

// GetEmailVisualIdentity retrieves email visual identity configuration
func (c *Client) GetEmailVisualIdentity(ctx context.Context) (*types.EmailVisualIdentityResponse, error) {
	var response types.EmailVisualIdentityResponse
	if err := c.get(ctx, "/backoffice/email-config/identity-visual", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateEmailVisualIdentity updates email visual identity configuration
func (c *Client) UpdateEmailVisualIdentity(ctx context.Context, req *types.EmailVisualIdentityRequest) (*types.EmailVisualIdentityResponse, error) {
	var response types.EmailVisualIdentityResponse
	if err := c.put(ctx, "/backoffice/email-config/identity-visual", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateEmailVisualIdentity creates email visual identity configuration
func (c *Client) CreateEmailVisualIdentity(ctx context.Context, req *types.EmailVisualIdentityRequest) (*types.EmailVisualIdentityResponse, error) {
	var response types.EmailVisualIdentityResponse
	if err := c.post(ctx, "/backoffice/email-config/identity-visual", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteEmailVisualIdentity deletes email visual identity configuration
func (c *Client) DeleteEmailVisualIdentity(ctx context.Context) error {
	return c.delete(ctx, "/backoffice/email-config/identity-visual", nil)
}
