package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestBanksEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListBanks", method: http.MethodGet, path: "/banks",
			response: &types.BanksResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBanks(ctx, nil)
				return err
			},
		},
		{
			name: "GetScheduledOperations", method: http.MethodGet, path: "/accounts/1/scheduleds",
			response: &types.ScheduledOperationsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetScheduledOperations(ctx, 1)
				return err
			},
		},
		{
			name: "CheckAPIStatus", method: http.MethodGet, path: "/status",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CheckAPIStatus(ctx)
				return err
			},
		},
		{
			name: "CheckIntegrationStatus", method: http.MethodGet, path: "/status/integrationModules",
			response: &types.IntegrationStatusResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CheckIntegrationStatus(ctx)
				return err
			},
		},
	})
}
