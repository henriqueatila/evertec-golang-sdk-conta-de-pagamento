package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestInstitutionEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateInstitution", method: http.MethodPost, path: "/institutions",
			response: &types.InstitutionResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateInstitution(ctx, &types.InstitutionRequest{})
				return err
			},
		},
		{
			name: "GetInstitution", method: http.MethodGet, path: "/institutions/1",
			response: &types.InstitutionResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetInstitution(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateInstitution", method: http.MethodPut, path: "/institutions/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateInstitution(ctx, 1, &types.InstitutionRequest{})
			},
		},
		{
			name: "ListInstitutions", method: http.MethodGet, path: "/institutions",
			response: []types.InstitutionResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInstitutions(ctx)
				return err
			},
		},
		{
			name: "DeleteInstitution", method: http.MethodDelete, path: "/institutions/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.DeleteInstitution(ctx, 1)
			},
		},
	})
}
