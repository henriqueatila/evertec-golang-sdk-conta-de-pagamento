package types

// Institution Operations

// InstitutionRequest represents institution create/update request
type InstitutionRequest struct {
	Name        string  `json:"name"`
	Code        *string `json:"code,omitempty"`
	CNPJ        *string `json:"cnpj,omitempty"`
	Description *string `json:"description,omitempty"`
	Address     *string `json:"address,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
	Website     *string `json:"website,omitempty"`
}

// InstitutionResponse represents an institution
type InstitutionResponse struct {
	InstitutionID int64   `json:"institutionId"`
	Name          string  `json:"name"`
	Code          *string `json:"code,omitempty"`
	CNPJ          *string `json:"cnpj,omitempty"`
	Description   *string `json:"description,omitempty"`
	Address       *string `json:"address,omitempty"`
	Phone         *string `json:"phone,omitempty"`
	Email         *string `json:"email,omitempty"`
	Website       *string `json:"website,omitempty"`
	Status        InstitutionStatus  `json:"status"`
	CreatedAt     *string `json:"createdAt,omitempty"`
	UpdatedAt     *string `json:"updatedAt,omitempty"`
}
