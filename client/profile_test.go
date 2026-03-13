package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestProfileEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetProfilePicture", method: http.MethodGet, path: "/accounts/1/picture",
			response: &ProfilePictureResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProfilePicture(ctx, 1)
				return err
			},
		},
		{
			name: "UploadProfilePicture", method: http.MethodPost, path: "/accounts/1/picture",
			response: &ProfilePictureResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UploadProfilePicture(ctx, 1, &UploadProfilePictureRequest{})
				return err
			},
		},
		{
			name: "DeleteProfilePicture", method: http.MethodDelete, path: "/accounts/1/picture",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.DeleteProfilePicture(ctx, 1)
			},
		},
		{
			name: "SaveDocumentImage", method: http.MethodPost, path: "/accounts/1/doc",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SaveDocumentImage(ctx, 1, &DocumentImageRequest{})
				return err
			},
		},
		{
			name: "UpdateDocumentImage", method: http.MethodPut, path: "/accounts/1/doc",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateDocumentImage(ctx, 1, &DocumentImageRequest{})
				return err
			},
		},
		{
			name: "GetDocumentImages", method: http.MethodGet, path: "/accounts/1/doc/rg/approved",
			response: []DocumentImageResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetDocumentImages(ctx, 1, types.DocType("rg"), "approved")
				return err
			},
		},
		{
			name: "GetCreditEngineInfo", method: http.MethodGet, path: "/accounts/1/creditEngineInfo",
			response: &types.CreditEngineInfoResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCreditEngineInfo(ctx, 1)
				return err
			},
		},
		{
			name: "CreateCreditEngineInfo", method: http.MethodPost, path: "/accounts/1/creditEngineInfo",
			response: &types.CreditEngineInfoResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateCreditEngineInfo(ctx, 1, &types.CreditEngineInfoRequest{})
				return err
			},
		},
	})
}
