package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Branch (Sales Channel) Operations

// CreateBranch creates a new branch (sales channel)
func (c *Client) CreateBranch(ctx context.Context, req *types.BranchRequest) (*types.BranchResponse, error) {
	var response types.BranchResponse
	if err := c.post(ctx, "/branches", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBranch retrieves a branch by ID
func (c *Client) GetBranch(ctx context.Context, branchID int64) (*types.BranchResponse, error) {
	path := fmt.Sprintf("/branches/%d", branchID)
	var response types.BranchResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateBranch updates an existing branch
func (c *Client) UpdateBranch(ctx context.Context, branchID int64, req *types.BranchRequest) error {
	path := fmt.Sprintf("/branches/%d", branchID)
	return c.put(ctx, path, req, nil)
}

// ListBranches retrieves all branches
func (c *Client) ListBranches(ctx context.Context) ([]types.BranchResponse, error) {
	var response []types.BranchResponse
	if err := c.get(ctx, "/branches", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteBranch deletes a branch
func (c *Client) DeleteBranch(ctx context.Context, branchID int64) error {
	path := fmt.Sprintf("/branches/%d", branchID)
	return c.delete(ctx, path, nil)
}
