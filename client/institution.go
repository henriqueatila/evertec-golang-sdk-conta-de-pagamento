package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Institution Operations

// CreateInstitution creates a new institution
func (c *Client) CreateInstitution(ctx context.Context, req *types.InstitutionRequest) (*types.InstitutionResponse, error) {
	var response types.InstitutionResponse
	if err := c.post(ctx, "/institutions", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetInstitution retrieves an institution by ID
func (c *Client) GetInstitution(ctx context.Context, institutionID int64) (*types.InstitutionResponse, error) {
	path := fmt.Sprintf("/institutions/%d", institutionID)
	var response types.InstitutionResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateInstitution updates an existing institution
func (c *Client) UpdateInstitution(ctx context.Context, institutionID int64, req *types.InstitutionRequest) error {
	path := fmt.Sprintf("/institutions/%d", institutionID)
	return c.put(ctx, path, req, nil)
}

// ListInstitutions retrieves all institutions
func (c *Client) ListInstitutions(ctx context.Context) ([]types.InstitutionResponse, error) {
	var response []types.InstitutionResponse
	if err := c.get(ctx, "/institutions", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteInstitution deletes an institution
func (c *Client) DeleteInstitution(ctx context.Context, institutionID int64) error {
	path := fmt.Sprintf("/institutions/%d", institutionID)
	return c.delete(ctx, path, nil)
}
