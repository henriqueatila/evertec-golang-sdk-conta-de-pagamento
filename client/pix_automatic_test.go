package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestPixAutomaticEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "StartAutomaticPix", method: http.MethodPost, path: "/pix/automatic",
			response: &types.StartAutomaticPixResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.StartAutomaticPix(ctx, &types.StartAutomaticPixRequest{})
				return err
			},
		},
		{
			name: "RejectAutomaticPix", method: http.MethodPost, path: "/pix/automatic/reject",
			response: &types.RejectAutomaticPixResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.RejectAutomaticPix(ctx, &types.RejectAutomaticPixRequest{})
				return err
			},
		},
		{
			name: "AcceptQRCodeJourneyThree", method: http.MethodPost, path: "/pix/automatic/qr-code/journey-three/accept",
			response: &types.QRCodeAcceptJourneyThreeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AcceptQRCodeJourneyThree(ctx, &types.QRCodeAcceptJourneyThreeRequest{})
				return err
			},
		},
		{
			name: "AcceptAutomaticPixQRCode", method: http.MethodPost, path: "/pix/automatic/qr-code/accept",
			response: &types.QRCodeUserResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AcceptAutomaticPixQRCode(ctx, &types.QRCodeUserAcceptRequest{})
				return err
			},
		},
		{
			name: "CreateAutomaticPixContract", method: http.MethodPost, path: "/pix/automatic/contract",
			response: &types.CreateAutomaticPixContractResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateAutomaticPixContract(ctx, &types.CreateAutomaticPixContractRequest{})
				return err
			},
		},
		{
			name: "CancelAutomaticPixCharge", method: http.MethodPost, path: "/pix/automatic/charge/cancel",
			response: &types.CancelAutomaticPixChargeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelAutomaticPixCharge(ctx, &types.CancelAutomaticPixChargeRequest{})
				return err
			},
		},
		{
			name: "CancelAutomaticPix", method: http.MethodPost, path: "/pix/automatic/cancel",
			response: &types.CancelAutomaticPixResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelAutomaticPix(ctx, &types.CancelAutomaticPixRequest{})
				return err
			},
		},
		{
			name: "AcceptAutomaticPix", method: http.MethodPost, path: "/pix/automatic/accept",
			response: &types.AutomaticPixResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AcceptAutomaticPix(ctx, &types.AcceptAutomaticPixRequest{})
				return err
			},
		},
		{
			name: "ListAutomaticPixCharges", method: http.MethodGet, path: "/pix/automatic/charge/account/1",
			response: &types.AutomaticPixChargeListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAutomaticPixCharges(ctx, 1, &types.ListAutomaticPixParams{})
				return err
			},
		},
		{
			name: "ListAutomaticPixByAccount", method: http.MethodGet, path: "/pix/automatic/account/1",
			response: &types.AutomaticPixListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAutomaticPixByAccount(ctx, 1, &types.ListAutomaticPixParams{})
				return err
			},
		},
		{
			name: "GetAutomaticPixRecurrence", method: http.MethodGet, path: "/pix/automatic/account/1/recurrence/rec-abc",
			response: &types.AutomaticPixResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAutomaticPixRecurrence(ctx, 1, "rec-abc", nil)
				return err
			},
		},
	})
}
