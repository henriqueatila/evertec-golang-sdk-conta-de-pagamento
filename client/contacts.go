package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// ListContacts retrieves contacts for an account
func (c *Client) ListContacts(ctx context.Context, accountID int64) (*types.ContactListResponse, error) {
	path := fmt.Sprintf("/accounts/%d/contact", accountID)
	var response types.ContactListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetContactBankDetails retrieves bank details for a contact
func (c *Client) GetContactBankDetails(ctx context.Context, accountID, contactID int64, transactionType string) (*types.ContactBankDetailsResponse, error) {
	path := fmt.Sprintf("/accounts/%d/contact/%d/bankdetails?transactionType=%s", accountID, contactID, transactionType)
	var response types.ContactBankDetailsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
