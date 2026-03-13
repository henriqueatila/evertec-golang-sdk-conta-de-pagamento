package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestPixQRCodeEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateStaticQRCode", method: http.MethodPost, path: "/pix/qrcodes/static",
			response: &types.QRCodeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateStaticQRCode(ctx, &types.StaticQRCodeRequest{})
				return err
			},
		},
		{
			name: "CreateDynamicQRCode", method: http.MethodPost, path: "/pix/qrcodes/dynamic",
			response: &types.QRCodeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateDynamicQRCode(ctx, &types.DynamicQRCodeRequest{})
				return err
			},
		},
		{
			name: "QueryQRCodeProcessing", method: http.MethodPost, path: "/pix/qrcodes/query-processing",
			response: &types.QRCodeQueryResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.QueryQRCodeProcessing(ctx, &types.QRCodeQueryRequest{})
				return err
			},
		},
		{
			name: "DecodeQRCodeV3", method: http.MethodPost, path: "/pix/qrcodes/v3/query-processing",
			response: &types.DecodeQRCodeV3Response{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DecodeQRCodeV3(ctx, &types.DecodeQRCodeV3Request{})
				return err
			},
		},
	})
}
