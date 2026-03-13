package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestCardsAdditionalEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateCard", method: http.MethodPost, path: "/accounts/1/cards/new",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateCard(ctx, 1, &types.CreateCardRequest{})
				return err
			},
		},
		{
			name: "CreateCardBackoffice", method: http.MethodPost, path: "/accounts/1/cards/new/backoffice",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateCardBackoffice(ctx, 1, &types.CreateCardRequest{})
				return err
			},
		},
		{
			name: "ReissueCard", method: http.MethodPut, path: "/accounts/1/cards/100/reissue",
			response: &types.BlockCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ReissueCard(ctx, 1, 100, &types.BlockCardRequest{})
				return err
			},
		},
		{
			name: "ReissueCardBackoffice", method: http.MethodPut, path: "/accounts/1/cards/100/reissue/backoffice",
			response: &types.BlockCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ReissueCardBackoffice(ctx, 1, 100, &types.BlockCardRequest{})
				return err
			},
		},
		{
			name: "ActivateCard", method: http.MethodPut, path: "/accounts/1/cards/100/activate",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ActivateCard(ctx, 1, 100, &types.ActivateCardRequest{})
				return err
			},
		},
		{
			name: "ChangeCardPin", method: http.MethodPut, path: "/accounts/1/cards/100/changePin",
			response: &types.ChangeCardPinResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangeCardPin(ctx, 1, 100, &types.ChangeCardPinRequest{})
				return err
			},
		},
		{
			name: "UpdateVirtualCardTag", method: http.MethodPut, path: "/accounts/1/cards/100/tag",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateVirtualCardTag(ctx, 1, 100, &types.UpdateVirtualCardTagRequest{})
			},
		},
		{
			name: "CreateVirtualCardFromPhysical", method: http.MethodPost, path: "/accounts/1/cards/100/virtual",
			response: &types.VirtualCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateVirtualCardFromPhysical(ctx, 1, 100)
				return err
			},
		},
		{
			name: "GetVirtualCards", method: http.MethodGet, path: "/accounts/1/cards/100/virtual",
			response: &types.VirtualCardsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetVirtualCards(ctx, 1, 100)
				return err
			},
		},
		{
			name: "ListAllVirtualCards", method: http.MethodGet, path: "/accounts/1/cards/virtual",
			response: &types.VirtualCardsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAllVirtualCards(ctx, 1)
				return err
			},
		},
		{
			name: "GetCardReplacementInfo", method: http.MethodGet, path: "/accounts/1/cards/100/replacement",
			response: &types.ReplacementCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardReplacementInfo(ctx, 1, 100)
				return err
			},
		},
		{
			name: "RequestCardReplacement", method: http.MethodPost, path: "/accounts/1/cards/100/replacement",
			response: &types.ReplacementCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.RequestCardReplacement(ctx, 1, 100)
				return err
			},
		},
		{
			name: "BindAnonymousCard", method: http.MethodPost, path: "/accounts/1/cards/bindAnonymousCard",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BindAnonymousCard(ctx, 1, &types.BindAnonymousCardRequest{})
				return err
			},
		},
		{
			name: "SearchCards", method: http.MethodGet, path: "/cards",
			response: &types.AccountCardsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SearchCards(ctx, nil)
				return err
			},
		},
		{
			name: "GetCardConfiguration", method: http.MethodGet, path: "/accounts/1/cards/100/getCardConfiguration",
			response: &types.CardConfigurationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardConfiguration(ctx, 1, 100)
				return err
			},
		},
		{
			name: "GetDefaultCardConfiguration", method: http.MethodGet, path: "/accounts/1/cards/getDefaultCardConfiguration",
			response: &types.DefaultCardConfigurationResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetDefaultCardConfiguration(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateDefaultCardConfiguration", method: http.MethodPut, path: "/accounts/1/cards/updateDefaultCardConfiguration",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateDefaultCardConfiguration(ctx, 1, &types.CardConfigurationRequest{})
			},
		},
		{
			name: "ConfigureCard", method: http.MethodPut, path: "/accounts/1/cards/configureCard",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.ConfigureCard(ctx, 1, &types.CardConfigurationRequest{})
			},
		},
		{
			name: "GetCardPaysmart", method: http.MethodGet, path: "/accounts/1/cards/paysmart/ps-abc",
			response: &types.CardPaysmartResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardPaysmart(ctx, 1, "ps-abc")
				return err
			},
		},
	})
}
