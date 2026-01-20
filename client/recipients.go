package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Recipient Operations
// Based on OpenAPI spec: CreateRecipient, UpdateRecipient, RecipientDTO, GetRecipientsResponse, GetRecipientResponse

// GetRecipients retrieves all recipients for an account
// Returns GetRecipientsResponse from API
func (c *Client) GetRecipients(ctx context.Context, accountID int64) (*types.GetRecipientsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/recipients", accountID)
	var response types.GetRecipientsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRecipient retrieves a specific recipient
// Returns GetRecipientResponse from API
func (c *Client) GetRecipient(ctx context.Context, accountID, recipientID int64) (*types.GetRecipientResponse, error) {
	path := fmt.Sprintf("/accounts/%d/recipients/%d", accountID, recipientID)
	var response types.GetRecipientResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateRecipient creates a new recipient
// Uses CreateRecipientRequest from API spec
func (c *Client) CreateRecipient(ctx context.Context, accountID int64, req *types.CreateRecipientRequest) (*types.GetRecipientResponse, error) {
	path := fmt.Sprintf("/accounts/%d/recipients", accountID)
	var response types.GetRecipientResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateRecipient updates an existing recipient
// Uses UpdateRecipientRequest from API spec
func (c *Client) UpdateRecipient(ctx context.Context, accountID int64, req *types.UpdateRecipientRequest) (*types.GetRecipientResponse, error) {
	path := fmt.Sprintf("/accounts/%d/recipients", accountID)
	var response types.GetRecipientResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteRecipient deletes a recipient
func (c *Client) DeleteRecipient(ctx context.Context, accountID, recipientID int64) error {
	path := fmt.Sprintf("/accounts/%d/recipients/%d", accountID, recipientID)
	return c.delete(ctx, path, nil)
}

// Feedback Operations
// Based on OpenAPI spec: FeedbackRequest, FeedbackStatementRequest, FeedbackResponse

// GetLastTransactionError retrieves the last transaction error for an account
// Returns FeedbackResponse from API
func (c *Client) GetLastTransactionError(ctx context.Context, accountID int64) (*types.FeedbackResponse, error) {
	path := fmt.Sprintf("/accounts/%d/feedback/lastTransactionError", accountID)
	var response types.FeedbackResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendFeedback sends feedback
// Uses FeedbackRequest from API spec
func (c *Client) SendFeedback(ctx context.Context, req *types.FeedbackRequest) (*types.ContaDigitalResponse, error) {
	var response types.ContaDigitalResponse
	if err := c.post(ctx, "/accounts/feedback/send", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SendStatementFeedback sends statement feedback
// Uses FeedbackStatementRequest from API spec
func (c *Client) SendStatementFeedback(ctx context.Context, req *types.FeedbackStatementRequest) (*types.ContaDigitalResponse, error) {
	var response types.ContaDigitalResponse
	if err := c.post(ctx, "/accounts/feedback/statement/send", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
