package types

// Product Limit Scheduling Operations

// ProductLimitSchedulingRequest represents product limit scheduling update
type ProductLimitSchedulingRequest struct {
	LimitDays int `json:"limitDays"`
}

// ProductLimitSchedulingResponse represents product limit scheduling
type ProductLimitSchedulingResponse struct {
	ProductID int64 `json:"productId"`
	LimitDays int   `json:"limitDays"`
}

// Paysmart Product Operations

// PaysmartProductResponse represents paysmart product
type PaysmartProductResponse struct {
	PaysmartProductID int64   `json:"paysmartProductId"`
	Name              string  `json:"name"`
	Description       *string `json:"description,omitempty"`
	Type              *string `json:"type,omitempty"`
	Status            ProductStatus  `json:"status"`
	ProcessorCode     *string `json:"processorCode,omitempty"`
	CreatedAt         *string `json:"createdAt,omitempty"`
	UpdatedAt         *string `json:"updatedAt,omitempty"`
}

// CreatePaysmartProductRequest represents paysmart product creation
type CreatePaysmartProductRequest struct {
	Name          string  `json:"name"`
	Description   *string `json:"description,omitempty"`
	Type          *string `json:"type,omitempty"`
	ProcessorCode *string `json:"processorCode,omitempty"`
}

// UpdatePaysmartProductRequest represents paysmart product update
type UpdatePaysmartProductRequest struct {
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	Status        *string `json:"status,omitempty"`
	ProcessorCode *string `json:"processorCode,omitempty"`
}

// Note: SearchProductLimitRequest is defined in limit.go with official API schema
