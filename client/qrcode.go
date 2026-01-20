package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// PayQRCode processes a QR code payment
func (c *Client) PayQRCode(ctx context.Context, accountID int64, req *types.QRCodePaymentRequest) (*types.QRCodePaymentResponse, error) {
	path := fmt.Sprintf("/accounts/%d/qrcode/payment", accountID)
	var response types.QRCodePaymentResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PaySimpleQRCode processes a simple QR code payment
func (c *Client) PaySimpleQRCode(ctx context.Context, accountID int64, req *types.SimpleQRCodePaymentRequest) (*types.QRCodePaymentResponse, error) {
	path := fmt.Sprintf("/accounts/%d/qrcode/simplePayment", accountID)
	var response types.QRCodePaymentResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ParseQRCode parses a QR code to extract information
func (c *Client) ParseQRCode(ctx context.Context, accountID int64, req *types.ParseQRCodeRequest) (*types.ParseQRCodeResponse, error) {
	path := fmt.Sprintf("/accounts/%d/qrcode/parse", accountID)
	var response types.ParseQRCodeResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetQRCodePublicKey retrieves the QR code public key
func (c *Client) GetQRCodePublicKey(ctx context.Context, accountID int64) (*types.QRCodePublicKeyResponse, error) {
	path := fmt.Sprintf("/accounts/%d/qrcode/publicKey", accountID)
	var response types.QRCodePublicKeyResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
