package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Profile Picture Operations

// ProfilePictureResponse represents profile picture data
type ProfilePictureResponse struct {
	AccountID   int64   `json:"accountId"`
	PictureURL  *string `json:"pictureUrl,omitempty"`
	PictureData *string `json:"pictureData,omitempty"` // Base64 encoded
	MimeType    *string `json:"mimeType,omitempty"`
}

// UploadProfilePictureRequest represents a profile picture upload request
type UploadProfilePictureRequest struct {
	PictureData string `json:"pictureData"` // Base64 encoded image
	MimeType    string `json:"mimeType"`    // e.g., "image/jpeg", "image/png"
}

// GetProfilePicture retrieves the profile picture for an account
func (c *Client) GetProfilePicture(ctx context.Context, accountID int64) (*ProfilePictureResponse, error) {
	path := fmt.Sprintf("/accounts/%d/picture", accountID)
	var response ProfilePictureResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UploadProfilePicture uploads a profile picture for an account
func (c *Client) UploadProfilePicture(ctx context.Context, accountID int64, req *UploadProfilePictureRequest) (*ProfilePictureResponse, error) {
	path := fmt.Sprintf("/accounts/%d/picture", accountID)
	var response ProfilePictureResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteProfilePicture deletes the profile picture for an account
func (c *Client) DeleteProfilePicture(ctx context.Context, accountID int64) error {
	path := fmt.Sprintf("/accounts/%d/picture", accountID)
	return c.delete(ctx, path, nil)
}

// Document Image Operations

// DocumentImageRequest represents a document image upload request
type DocumentImageRequest struct {
	ImageData string `json:"imageData"` // Base64 encoded image
	MimeType  string `json:"mimeType"`  // e.g., "image/jpeg", "image/png"
}

// DocumentImageResponse represents document image data
type DocumentImageResponse struct {
	AccountID  int64         `json:"accountId"`
	DocType    types.DocType `json:"docType"`
	Status     string        `json:"status"`
	ImageURL   *string       `json:"imageUrl,omitempty"`
	ImageData  *string       `json:"imageData,omitempty"` // Base64 encoded
	MimeType   *string       `json:"mimeType,omitempty"`
	VerifiedAt *string       `json:"verifiedAt,omitempty"`
}

// SaveDocumentImage saves a document image for an account
func (c *Client) SaveDocumentImage(ctx context.Context, accountID int64, req *DocumentImageRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/doc", accountID)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateDocumentImage updates a document image for an account
func (c *Client) UpdateDocumentImage(ctx context.Context, accountID int64, req *DocumentImageRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/doc", accountID)
	var response types.GenericResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetDocumentImages retrieves document images by type and status
func (c *Client) GetDocumentImages(ctx context.Context, accountID int64, docType types.DocType, status string) ([]DocumentImageResponse, error) {
	path := fmt.Sprintf("/accounts/%d/doc/%s/%s", accountID, docType, status)
	var response []DocumentImageResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Credit Engine Operations

// GetCreditEngineInfo retrieves credit engine information for an account
func (c *Client) GetCreditEngineInfo(ctx context.Context, accountID int64) (*types.CreditEngineInfoResponse, error) {
	path := fmt.Sprintf("/accounts/%d/creditEngineInfo", accountID)
	var response types.CreditEngineInfoResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateCreditEngineInfo creates credit engine information for an account
func (c *Client) CreateCreditEngineInfo(ctx context.Context, accountID int64, req *types.CreditEngineInfoRequest) (*types.CreditEngineInfoResponse, error) {
	path := fmt.Sprintf("/accounts/%d/creditEngineInfo", accountID)
	var response types.CreditEngineInfoResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
