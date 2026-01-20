package types

import "time"

// ProposalStatus represents the status of a proposal
type ProposalStatus string

const (
	ProposalStatusPending  ProposalStatus = "PENDING"
	ProposalStatusApproved ProposalStatus = "APPROVED"
	ProposalStatusRejected ProposalStatus = "REJECTED"
	ProposalStatusCanceled ProposalStatus = "CANCELED"
)

// ProposalResponse represents a proposal
type ProposalResponse struct {
	ProposalID  int64          `json:"proposalId"`
	Status      ProposalStatus `json:"status"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Document    string         `json:"document"`
	Phone       string         `json:"phone"`
	BirthDate   *string        `json:"birthDate,omitempty"`
	AccountType AccountType    `json:"accountType"`

	// Company fields
	CompanyName *string `json:"companyName,omitempty"`
	TradeName   *string `json:"tradeName,omitempty"`

	// Address
	Address *AddressResponse `json:"address,omitempty"`

	// Document images
	HasDocumentImages bool `json:"hasDocumentImages"`

	// Dates
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	ApprovedAt *time.Time `json:"approvedAt,omitempty"`
	RejectedAt *time.Time `json:"rejectedAt,omitempty"`

	// Rejection reason
	RejectionReason *string `json:"rejectionReason,omitempty"`
}

// ProposalListResponse represents a list of proposals
type ProposalListResponse struct {
	Proposals []ProposalResponse `json:"items"`
	Total     *int               `json:"total,omitempty"`
}

// UpdateProposalRequest represents a request to update a proposal (API: UpdateProposalRequest)
type UpdateProposalRequest struct {
	Name        *string         `json:"name,omitempty"`
	SocialName  *string         `json:"socialName,omitempty"`  // New field
	ProductID   *int32          `json:"productId,omitempty"`   // New field
	Ownership   *int32          `json:"ownerShip,omitempty"`   // New field (note: API uses ownerShip)
	BirthDate   *string         `json:"birthDate,omitempty"`
	FatherName  *string         `json:"fatherName,omitempty"`  // New field
	MotherName  *string         `json:"motherName,omitempty"`  // New field
	Nationality *string         `json:"nationality,omitempty"` // New field
	Valid       *bool           `json:"valid,omitempty"`       // New field
	Email       *string         `json:"email,omitempty"`
	Phone       *string         `json:"phone,omitempty"`
	Address     *AddressRequest `json:"address,omitempty"` // Should be UpdateProposalAddressRequest
}

// ProcessProposalRequest represents a request to approve or reject a proposal
type ProcessProposalRequest struct {
	ProposalID      int64   `json:"proposalId"`
	Action          string  `json:"action"` // "APPROVE" or "REJECT"
	RejectionReason *string `json:"rejectionReason,omitempty"`
}

// ResendProposalRequest represents a request to resend a proposal
type ResendProposalRequest struct {
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// ProposalImageType represents the type of proposal document image
type ProposalImageType string

const (
	ProposalImageTypeFront          ProposalImageType = "FRONT"
	ProposalImageTypeBack           ProposalImageType = "BACK"
	ProposalImageTypeSelfie         ProposalImageType = "SELFIE"
	ProposalImageTypeProofOfAddress ProposalImageType = "PROOF_OF_ADDRESS"
)

// ProposalImageResponse represents a proposal document image
type ProposalImageResponse struct {
	ImageID    int64             `json:"imageId"`
	ProposalID int64             `json:"proposalId"`
	ImageType  ProposalImageType `json:"imageType"`
	ImageURL   string            `json:"imageUrl"`
	Status     ProposalStatus    `json:"status"` // e.g., "PENDING", "APPROVED", "REJECTED"
	UploadedAt *time.Time        `json:"uploadedAt,omitempty"`
}

// ProposalImageListResponse represents a list of proposal images
type ProposalImageListResponse struct {
	Images []ProposalImageResponse `json:"items"`
}
