package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestBranchEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateBranch", method: http.MethodPost, path: "/branches",
			response: &types.BranchResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBranch(ctx, &types.BranchRequest{})
				return err
			},
		},
		{
			name: "GetBranch", method: http.MethodGet, path: "/branches/1",
			response: &types.BranchResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBranch(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateBranch", method: http.MethodPut, path: "/branches/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateBranch(ctx, 1, &types.BranchRequest{})
			},
		},
		{
			name: "ListBranches", method: http.MethodGet, path: "/branches",
			response: []types.BranchResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBranches(ctx)
				return err
			},
		},
		{
			name: "DeleteBranch", method: http.MethodDelete, path: "/branches/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.DeleteBranch(ctx, 1)
			},
		},
	})
}
