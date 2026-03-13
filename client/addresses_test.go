package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestAddressesEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListAddresses", method: http.MethodGet, path: "/accounts/1/address",
			response: []types.AddressResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAddresses(ctx, 1)
				return err
			},
		},
		{
			name: "CreateAddress", method: http.MethodPost, path: "/accounts/1/address",
			response: &types.AddressResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateAddress(ctx, 1, &types.AddressRequest{})
				return err
			},
		},
		{
			name: "GetAddress", method: http.MethodGet, path: "/accounts/1/address/100",
			response: &types.AddressResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAddress(ctx, 1, 100)
				return err
			},
		},
		{
			name: "UpdateAddress", method: http.MethodPut, path: "/accounts/1/address/100",
			response: &types.AddressResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAddress(ctx, 1, 100, &types.AddressRequest{})
				return err
			},
		},
		{
			name: "DeleteAddress", method: http.MethodDelete, path: "/accounts/1/address/100",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DeleteAddress(ctx, 1, 100)
				return err
			},
		},
		{
			name: "LookupPostalCode", method: http.MethodGet, path: "/address/postalcode/01310100",
			response: &types.PostalCodeLookupResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.LookupPostalCode(ctx, "01310100")
				return err
			},
		},
	})
}
