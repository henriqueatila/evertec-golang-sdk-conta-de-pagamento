package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestTravelEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetTravelNotices", method: http.MethodGet, path: "/travel/account/1/notify",
			response: &types.GetAccountTravelingResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetTravelNotices(ctx, 1)
				return err
			},
		},
		{
			name: "CreateTravelNotice", method: http.MethodPost, path: "/travel/account/1/notify",
			response: &types.TravelNoticeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateTravelNotice(ctx, 1, &types.NotifyTripRequest{})
				return err
			},
		},
		{
			name: "GetTravelCountries", method: http.MethodGet, path: "/travel/getCountries",
			response: &types.CountryListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetTravelCountries(ctx)
				return err
			},
		},
		{
			name: "ChangeAccountPassword", method: http.MethodPut, path: "/accounts/1/changeUserPassword",
			response: &types.ContaDigitalGenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangeAccountPassword(ctx, 1, &types.ChangeUserPasswordRequest{})
				return err
			},
		},
		{
			name: "ChangeAccountStatus", method: http.MethodPut, path: "/accounts/1/changeStatus/active",
			response: &types.ContaDigitalGenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangeAccountStatus(ctx, 1, types.AccountStatus("active"))
				return err
			},
		},
		{
			name: "UpdateAccountName", method: http.MethodPut, path: "/accounts/1/name",
			response: &types.ContaDigitalGenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAccountName(ctx, 1, &types.UpdateNameInAccountRequest{})
				return err
			},
		},
	})
}
