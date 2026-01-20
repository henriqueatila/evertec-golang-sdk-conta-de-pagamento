package types

// Branch (Sales Channel) Operations

// BranchRequest represents a branch create/update request (API: BranchRequest)
type BranchRequest struct {
	BranchID    int64  `json:"branchId"`    // Required (unified to int64 for consistency)
	BranchName  string `json:"branchName"`  // Required
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Active      *bool   `json:"active,omitempty"`
	Address     *string `json:"address,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
}

// BranchResponse represents a branch (API: BranchResponse)
type BranchResponse struct {
	Message     string  `json:"message"`    // API response wrapper field
	BranchID    int64   `json:"branchId"`
	BranchName  string  `json:"branchName"` // Renamed from Name to match spec
	DaCode      int32   `json:"da_code"`    // API response wrapper field
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Active      bool    `json:"active"`
	Address     *string `json:"address,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	Email       *string `json:"email,omitempty"`
	CreatedAt   *string `json:"createdAt,omitempty"`
	UpdatedAt   *string `json:"updatedAt,omitempty"`
}
