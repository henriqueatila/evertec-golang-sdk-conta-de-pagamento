package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestProposalsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListProposals", method: http.MethodGet, path: "/proposal",
			response: []types.ProposalResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListProposals(ctx, nil)
				return err
			},
		},
		{
			name: "GetProposal", method: http.MethodGet, path: "/proposal/1",
			response: &types.ProposalDetailResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProposal(ctx, 1)
				return err
			},
		},
		{
			name: "GetProposalImages", method: http.MethodGet, path: "/proposal/1/images",
			response: []types.ProposalImage{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProposalImages(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateProposal", method: http.MethodPut, path: "/proposal/1/update",
			response: &types.ProposalDetailResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateProposal(ctx, 1, &types.UpdateProposalRequest{})
				return err
			},
		},
		{
			name: "UpdateProposalImages", method: http.MethodPut, path: "/proposal/1/update/images",
			response: &types.ProposalDetailResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateProposalImages(ctx, 1, []types.UpdateProposalImageRequest{})
				return err
			},
		},
		{
			name: "ResendProposal", method: http.MethodPut, path: "/proposal/1/resend",
			response: &types.ProposalDetailResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ResendProposal(ctx, 1)
				return err
			},
		},
		{
			name: "GetProposalTypeStatus", method: http.MethodGet, path: "/proposal/type-status",
			response: []types.ProposalTypeStatus{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProposalTypeStatus(ctx)
				return err
			},
		},
		{
			name: "GetLastProposal", method: http.MethodGet, path: "/proposal/last/12345678901",
			response: &types.ProposalDetailResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetLastProposal(ctx, "12345678901")
				return err
			},
		},
		{
			name: "ListLegalEntityProposals", method: http.MethodGet, path: "/proposal/legalEntities",
			response: []types.LegalEntityProposalResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListLegalEntityProposals(ctx, nil)
				return err
			},
		},
		{
			name: "GetLegalEntityProposal", method: http.MethodGet, path: "/proposal/legalEntityProposal/1",
			response: &types.LegalEntityProposalDetailResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetLegalEntityProposal(ctx, 1)
				return err
			},
		},
	})
}
