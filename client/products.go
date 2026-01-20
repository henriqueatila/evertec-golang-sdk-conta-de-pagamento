package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Product Operations

// ListProducts retrieves all products
func (c *Client) ListProducts(ctx context.Context) (*types.ProductListResponse, error) {
	var response types.ProductListResponse
	if err := c.get(ctx, "/products", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetProduct retrieves a product by ID
func (c *Client) GetProduct(ctx context.Context, productID int64) (*types.ProductResponse, error) {
	path := fmt.Sprintf("/products/%d", productID)
	var response types.ProductResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateProduct creates a new product
func (c *Client) CreateProduct(ctx context.Context, req *types.CreateProductRequest) (*types.ProductResponse, error) {
	var response types.ProductResponse
	if err := c.post(ctx, "/products", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateProduct updates an existing product
func (c *Client) UpdateProduct(ctx context.Context, productID int64, req *types.UpdateProductRequest) error {
	path := fmt.Sprintf("/products/%d", productID)
	return c.put(ctx, path, req, nil)
}

// Product Limit Scheduling Operations

// GetProductLimitScheduling retrieves product limit scheduling
func (c *Client) GetProductLimitScheduling(ctx context.Context, productID int64) (*types.ProductLimitSchedulingResponse, error) {
	path := fmt.Sprintf("/products/%d/limit-scheduling", productID)
	var response types.ProductLimitSchedulingResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateProductLimitScheduling updates product limit scheduling
func (c *Client) UpdateProductLimitScheduling(ctx context.Context, productID int64, req *types.ProductLimitSchedulingRequest) error {
	path := fmt.Sprintf("/products/%d/limit-scheduling", productID)
	return c.put(ctx, path, req, nil)
}

// Paysmart Product Operations

// ListPaysmartProducts retrieves all paysmart products
func (c *Client) ListPaysmartProducts(ctx context.Context) ([]types.PaysmartProductResponse, error) {
	var response []types.PaysmartProductResponse
	if err := c.get(ctx, "/paysmart/products", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPaysmartProduct retrieves a paysmart product by ID
func (c *Client) GetPaysmartProduct(ctx context.Context, productID int64) (*types.PaysmartProductResponse, error) {
	path := fmt.Sprintf("/paysmart/products/%d", productID)
	var response types.PaysmartProductResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreatePaysmartProduct creates a new paysmart product
func (c *Client) CreatePaysmartProduct(ctx context.Context, req *types.CreatePaysmartProductRequest) (*types.PaysmartProductResponse, error) {
	var response types.PaysmartProductResponse
	if err := c.post(ctx, "/paysmart/products", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePaysmartProduct updates an existing paysmart product
func (c *Client) UpdatePaysmartProduct(ctx context.Context, productID int64, req *types.UpdatePaysmartProductRequest) error {
	path := fmt.Sprintf("/paysmart/products/%d", productID)
	return c.put(ctx, path, req, nil)
}

// SearchProductLimits searches for product limits
func (c *Client) SearchProductLimits(ctx context.Context, req *types.SearchProductLimitRequest) ([]types.ProductLimitResponse, error) {
	var response []types.ProductLimitResponse
	if err := c.post(ctx, "/products/limits/search", req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateProductLimit updates a product limit
func (c *Client) UpdateProductLimit(ctx context.Context, productID int64, limitType string, req *types.ProductLimitRequest) error {
	path := fmt.Sprintf("/products/%d/limits/%s", productID, limitType)
	return c.put(ctx, path, req, nil)
}
