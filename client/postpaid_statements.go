package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Post-paid Account Operations
// Based on OpenAPI spec: AccountPostPaidResponse, PaysmartOpenStatementResponse

// GetPostPaidAccount retrieves post-paid account information
// Returns AccountPostPaidResponse from API spec
func (c *Client) GetPostPaidAccount(ctx context.Context, accountID int64) (*types.AccountPostPaidResponse, error) {
	path := fmt.Sprintf("/postpaid/account/%d", accountID)
	var response types.AccountPostPaidResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePostPaidAccountInfo updates post-paid account information
func (c *Client) UpdatePostPaidAccountInfo(ctx context.Context, req *types.UpdatePostPaidAccountRequest) (*types.AccountPostPaidResponse, error) {
	var response types.AccountPostPaidResponse
	if err := c.post(ctx, "/postpaid/account", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidDueDates retrieves available due date options
// Returns DueDateOption from API spec
func (c *Client) GetPostPaidDueDates(ctx context.Context, accountID int64) ([]types.DueDateOption, error) {
	path := fmt.Sprintf("/postpaid/account/%d/dueDates", accountID)
	var response []types.DueDateOption
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPostPaidCardDueDates retrieves available due dates for a specific card
// Returns DueDateOption from API spec
func (c *Client) GetPostPaidCardDueDates(ctx context.Context, accountID, cardID int64) ([]types.DueDateOption, error) {
	path := fmt.Sprintf("/postpaid/account/%d/cards/%d/dueDates", accountID, cardID)
	var response []types.DueDateOption
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Post-paid Statement Operations
// Based on OpenAPI spec: PaysmartOpenStatementResponse

// GetPostPaidStatementByMonth retrieves statement for a specific month
// Returns PaysmartOpenStatementResponse from API spec
func (c *Client) GetPostPaidStatementByMonth(ctx context.Context, accountID int64, month, year int) (*types.PaysmartOpenStatementResponse, error) {
	path := fmt.Sprintf("/postpaid/statements/%d?month=%d&year=%d", accountID, month, year)
	var response types.PaysmartOpenStatementResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidOpenStatement retrieves the current open statement
// Returns PaysmartOpenStatementResponse from API spec
func (c *Client) GetPostPaidOpenStatement(ctx context.Context, accountID int64) (*types.PaysmartOpenStatementResponse, error) {
	path := fmt.Sprintf("/postpaid/statements/%d/open-statement", accountID)
	var response types.PaysmartOpenStatementResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidClosedStatement retrieves the most recent closed statement
// Returns PaysmartOpenStatementResponse from API spec
func (c *Client) GetPostPaidClosedStatement(ctx context.Context, accountID int64) (*types.PaysmartOpenStatementResponse, error) {
	path := fmt.Sprintf("/postpaid/statements/%d/closed-statement", accountID)
	var response types.PaysmartOpenStatementResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidFutureStatement retrieves future scheduled transactions
// Returns PaysmartOpenStatementResponse from API spec
func (c *Client) GetPostPaidFutureStatement(ctx context.Context, accountID int64) (*types.PaysmartOpenStatementResponse, error) {
	path := fmt.Sprintf("/postpaid/statements/%d/future-statement", accountID)
	var response types.PaysmartOpenStatementResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidCombinedStatement retrieves combined statement (open + closed + future)
// Returns PaysmartOpenStatementResponse from API spec
func (c *Client) GetPostPaidCombinedStatement(ctx context.Context, accountID int64) (*types.PaysmartOpenStatementResponse, error) {
	path := fmt.Sprintf("/postpaid/statements/%d/combined-statement", accountID)
	var response types.PaysmartOpenStatementResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetPostPaidTransactions retrieves post-paid transactions list
// Returns TransactionsDTO array from API spec
func (c *Client) GetPostPaidTransactions(ctx context.Context, accountID int64) ([]types.TransactionsDTO, error) {
	path := fmt.Sprintf("/postpaid/statements/%d/transactions", accountID)
	var response []types.TransactionsDTO
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetPostPaidPossibleAdvances retrieves transactions eligible for advance
// Returns TransactionsDTO array from API spec
func (c *Client) GetPostPaidPossibleAdvances(ctx context.Context, accountID int64) ([]types.TransactionsDTO, error) {
	path := fmt.Sprintf("/postpaid/statements/%d/possible-advance", accountID)
	var response []types.TransactionsDTO
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// SendPostPaidStatementEmail sends statement to email
func (c *Client) SendPostPaidStatementEmail(ctx context.Context, accountID int64, email string) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/postpaid/statements/%d/mail", accountID)
	req := map[string]string{"email": email}
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
