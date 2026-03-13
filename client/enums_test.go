package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestEnumsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetStates", method: http.MethodGet, path: "/enum/uf",
			response: []types.StateResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetStates(ctx)
				return err
			},
		},
		{
			name: "GetProfessions", method: http.MethodGet, path: "/enum/profession",
			response: []types.ProfessionResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProfessions(ctx)
				return err
			},
		},
		{
			name: "GetIssuingAuthorities", method: http.MethodGet, path: "/enum/issuingAuthority",
			response: []types.IssuingAuthorityResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetIssuingAuthorities(ctx)
				return err
			},
		},
		{
			name: "GetGenders", method: http.MethodGet, path: "/enum/gender",
			response: []types.GenderResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetGenders(ctx)
				return err
			},
		},
		{
			name: "GetCountries", method: http.MethodGet, path: "/country",
			response: []types.PhoneCodeCountryResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCountries(ctx)
				return err
			},
		},
		{
			name: "GetAllBanks", method: http.MethodGet, path: "/banco",
			response: []types.BankResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAllBanks(ctx)
				return err
			},
		},
	})
}
