package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Travel Notification Operations
// Based on OpenAPI spec: NotifyTripRequest, TravelNoticeResponse, GetAccountTravelingResponse, CountryListResponse

// GetTravelNotices retrieves travel notices for an account
// Returns GetAccountTravelingResponse from API
func (c *Client) GetTravelNotices(ctx context.Context, accountID int64) (*types.GetAccountTravelingResponse, error) {
	path := fmt.Sprintf("/travel/account/%d/notify", accountID)
	var response types.GetAccountTravelingResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateTravelNotice creates a travel notice
// Uses NotifyTripRequest from API spec
func (c *Client) CreateTravelNotice(ctx context.Context, accountID int64, req *types.NotifyTripRequest) (*types.TravelNoticeResponse, error) {
	path := fmt.Sprintf("/travel/account/%d/notify", accountID)
	var response types.TravelNoticeResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTravelCountries retrieves all countries available for travel notification
// Returns CountryListResponse from API
func (c *Client) GetTravelCountries(ctx context.Context) (*types.CountryListResponse, error) {
	var response types.CountryListResponse
	if err := c.get(ctx, "/travel/getCountries", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Account Password Operations
// Based on OpenAPI spec: ChangeUserPasswordRequest

// ChangeAccountPassword changes the account password
// Uses ChangeUserPasswordRequest from API spec
func (c *Client) ChangeAccountPassword(ctx context.Context, accountID int64, req *types.ChangeUserPasswordRequest) (*types.ContaDigitalGenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/changeUserPassword", accountID)
	var response types.ContaDigitalGenericResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Account Status Operations

// ChangeAccountStatus changes the account status
func (c *Client) ChangeAccountStatus(ctx context.Context, accountID int64, targetStatus types.AccountStatus) (*types.ContaDigitalGenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/changeStatus/%s", accountID, targetStatus)
	var response types.ContaDigitalGenericResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Account Name Operations
// Based on OpenAPI spec: UpdateNameInAccountRequest

// UpdateAccountName updates the account holder name
// Uses UpdateNameInAccountRequest from API spec
func (c *Client) UpdateAccountName(ctx context.Context, accountID int64, req *types.UpdateNameInAccountRequest) (*types.ContaDigitalGenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/name", accountID)
	var response types.ContaDigitalGenericResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
