package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestQRCodeEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "PayQRCode", method: http.MethodPost, path: "/accounts/1/qrcode/payment",
			response: &types.QRCodePaymentResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PayQRCode(ctx, 1, &types.QRCodePaymentRequest{})
				return err
			},
		},
		{
			name: "PaySimpleQRCode", method: http.MethodPost, path: "/accounts/1/qrcode/simplePayment",
			response: &types.QRCodePaymentResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PaySimpleQRCode(ctx, 1, &types.SimpleQRCodePaymentRequest{})
				return err
			},
		},
		{
			name: "ParseQRCode", method: http.MethodPost, path: "/accounts/1/qrcode/parse",
			response: &types.ParseQRCodeResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ParseQRCode(ctx, 1, &types.ParseQRCodeRequest{})
				return err
			},
		},
		{
			name: "GetQRCodePublicKey", method: http.MethodGet, path: "/accounts/1/qrcode/publicKey",
			response: &types.QRCodePublicKeyResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetQRCodePublicKey(ctx, 1)
				return err
			},
		},
	})
}
