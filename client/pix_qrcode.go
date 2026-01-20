package client

import (
	"context"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// PIX QR Code Operations

// CreateStaticQRCode creates a static PIX QR code
func (c *Client) CreateStaticQRCode(ctx context.Context, req *types.StaticQRCodeRequest) (*types.QRCodeResponse, error) {
	var response types.QRCodeResponse
	if err := c.post(ctx, "/pix/qrcodes/static", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateDynamicQRCode creates a dynamic PIX QR code
func (c *Client) CreateDynamicQRCode(ctx context.Context, req *types.DynamicQRCodeRequest) (*types.QRCodeResponse, error) {
	var response types.QRCodeResponse
	if err := c.post(ctx, "/pix/qrcodes/dynamic", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// QueryQRCodeProcessing queries QR code processing status
func (c *Client) QueryQRCodeProcessing(ctx context.Context, req *types.QRCodeQueryRequest) (*types.QRCodeQueryResponse, error) {
	var response types.QRCodeQueryResponse
	if err := c.post(ctx, "/pix/qrcodes/query-processing", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DecodeQRCodeV3 decodes a PIX QR code using v3 API
func (c *Client) DecodeQRCodeV3(ctx context.Context, req *types.DecodeQRCodeV3Request) (*types.DecodeQRCodeV3Response, error) {
	var response types.DecodeQRCodeV3Response
	if err := c.post(ctx, "/pix/qrcodes/v3/query-processing", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
