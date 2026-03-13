package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestProductsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListProducts", method: http.MethodGet, path: "/products",
			response: &types.ProductListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListProducts(ctx)
				return err
			},
		},
		{
			name: "GetProduct", method: http.MethodGet, path: "/products/1",
			response: &types.ProductResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProduct(ctx, 1)
				return err
			},
		},
		{
			name: "CreateProduct", method: http.MethodPost, path: "/products",
			response: &types.ProductResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateProduct(ctx, &types.CreateProductRequest{})
				return err
			},
		},
		{
			name: "UpdateProduct", method: http.MethodPut, path: "/products/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateProduct(ctx, 1, &types.UpdateProductRequest{})
			},
		},
		{
			name: "GetProductLimitScheduling", method: http.MethodGet, path: "/products/1/limit-scheduling",
			response: &types.ProductLimitSchedulingResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProductLimitScheduling(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateProductLimitScheduling", method: http.MethodPut, path: "/products/1/limit-scheduling",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateProductLimitScheduling(ctx, 1, &types.ProductLimitSchedulingRequest{})
			},
		},
		{
			name: "ListPaysmartProducts", method: http.MethodGet, path: "/paysmart/products",
			response: []types.PaysmartProductResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListPaysmartProducts(ctx)
				return err
			},
		},
		{
			name: "GetPaysmartProduct", method: http.MethodGet, path: "/paysmart/products/1",
			response: &types.PaysmartProductResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPaysmartProduct(ctx, 1)
				return err
			},
		},
		{
			name: "CreatePaysmartProduct", method: http.MethodPost, path: "/paysmart/products",
			response: &types.PaysmartProductResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePaysmartProduct(ctx, &types.CreatePaysmartProductRequest{})
				return err
			},
		},
		{
			name: "UpdatePaysmartProduct", method: http.MethodPut, path: "/paysmart/products/1",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdatePaysmartProduct(ctx, 1, &types.UpdatePaysmartProductRequest{})
			},
		},
		{
			name: "SearchProductLimits", method: http.MethodPost, path: "/products/limits/search",
			response: []types.ProductLimitResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SearchProductLimits(ctx, &types.SearchProductLimitRequest{})
				return err
			},
		},
		{
			name: "UpdateProductLimit", method: http.MethodPut, path: "/products/1/limits/pix",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateProductLimit(ctx, 1, "pix", &types.ProductLimitRequest{})
			},
		},
	})
}
