package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestPostpaidCardsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetPostPaidVirtualCards", method: http.MethodGet, path: "/postpaid/cards/1/virtual",
			response: &types.VirtualCardsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidVirtualCards(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidPhysicalCards", method: http.MethodGet, path: "/postpaid/cards/1/physical",
			response: &types.AccountCardsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidPhysicalCards(ctx, 1)
				return err
			},
		},
		{
			name: "CreatePostPaidCard", method: http.MethodPost, path: "/postpaid/cards/new",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePostPaidCard(ctx, &types.CreateCardRequest{})
				return err
			},
		},
		{
			name: "CreatePostPaidVirtualCard", method: http.MethodPost, path: "/postpaid/cards/1/new/virtual",
			response: &types.VirtualCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePostPaidVirtualCard(ctx, 1, &types.CreateVirtualCardRequest{})
				return err
			},
		},
		{
			name: "BlockPostPaidCard", method: http.MethodPost, path: "/postpaid/cards/1/block/100",
			response: &types.BlockCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BlockPostPaidCard(ctx, 1, 100, &types.BlockCardRequest{})
				return err
			},
		},
		{
			name: "UnblockPostPaidCard", method: http.MethodPost, path: "/postpaid/cards/1/unblock/100",
			response: &types.UnblockCardResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UnblockPostPaidCard(ctx, 1, 100)
				return err
			},
		},
		{
			name: "ActivatePostPaidCard", method: http.MethodPost, path: "/postpaid/cards/1/activate/100",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ActivatePostPaidCard(ctx, 1, 100, &types.ActivateCardRequest{})
				return err
			},
		},
		{
			name: "ChangePostPaidCardPin", method: http.MethodPost, path: "/postpaid/cards/1/changePin/100",
			response: &types.ChangeCardPinResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangePostPaidCardPin(ctx, 1, 100, &types.ChangeCardPinRequest{})
				return err
			},
		},
		{
			name: "ValidatePostPaidCardPin", method: http.MethodPost, path: "/postpaid/cards/1/changeValidatePin/100",
			response: &types.ChangeCardPinResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ValidatePostPaidCardPin(ctx, 1, 100, &types.ChangeCardPinRequest{})
				return err
			},
		},
		{
			name: "GetPostPaidCardSettings", method: http.MethodGet, path: "/postpaid/account/1/cards/100/settings",
			response: &types.PostpaidCardSettingsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidCardSettings(ctx, 1, 100)
				return err
			},
		},
		{
			name: "UpdatePostPaidCardSettings", method: http.MethodPost, path: "/postpaid/account/1/cards/100/settings",
			response: &types.PostpaidCardSettingsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdatePostPaidCardSettings(ctx, 1, 100, &types.PostpaidCardSettingsRequest{})
				return err
			},
		},
		{
			name: "ResetPostPaidCardSettings", method: http.MethodPatch, path: "/postpaid/account/1/cards/100/settings/reset",
			response: &types.PostpaidCardSettingsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ResetPostPaidCardSettings(ctx, 1, 100)
				return err
			},
		},
	})
}
