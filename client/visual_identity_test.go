package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestVisualIdentityEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetEmailVisualIdentity", method: http.MethodGet, path: "/backoffice/email-config/identity-visual",
			response: &types.EmailVisualIdentityResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetEmailVisualIdentity(ctx)
				return err
			},
		},
		{
			name: "UpdateEmailVisualIdentity", method: http.MethodPut, path: "/backoffice/email-config/identity-visual",
			response: &types.EmailVisualIdentityResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateEmailVisualIdentity(ctx, &types.EmailVisualIdentityRequest{})
				return err
			},
		},
		{
			name: "CreateEmailVisualIdentity", method: http.MethodPost, path: "/backoffice/email-config/identity-visual",
			response: &types.EmailVisualIdentityResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateEmailVisualIdentity(ctx, &types.EmailVisualIdentityRequest{})
				return err
			},
		},
		{
			name: "DeleteEmailVisualIdentity", method: http.MethodDelete, path: "/backoffice/email-config/identity-visual",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.DeleteEmailVisualIdentity(ctx)
			},
		},
	})
}
