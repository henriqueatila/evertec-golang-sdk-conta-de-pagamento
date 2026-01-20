package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// GetAccount retrieves account data by ID
func (c *Client) GetAccount(ctx context.Context, accountID int64) (*types.AccountDataResponse, error) {
	path := fmt.Sprintf("/accounts/%d", accountID)
	var response types.AccountDataResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListAccounts retrieves a list of accounts with optional filters
func (c *Client) ListAccounts(ctx context.Context, params *types.ListAccountsParams) (*types.AccountListResponse, error) {
	path := "/accounts"
	if params != nil {
		path += params.QueryString()
	}
	var response types.AccountListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateAccount creates a new account (proposal)
func (c *Client) CreateAccount(ctx context.Context, req *types.ProposalAccountRequest) (*types.CreateAccountResponse, error) {
	var response types.CreateAccountResponse
	if err := c.post(ctx, "/accounts", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateAccount updates an existing account
func (c *Client) UpdateAccount(ctx context.Context, accountID int64, req *types.UpdateAccountRequest) error {
	path := fmt.Sprintf("/accounts/%d", accountID)
	return c.put(ctx, path, req, nil)
}

// LinkAccounts links a sub-account to a main account
func (c *Client) LinkAccounts(ctx context.Context, mainAccountID int64, req *types.LinkAccountRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/accounts/%d/link", mainAccountID)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UnlinkAccounts unlinks accounts
func (c *Client) UnlinkAccounts(ctx context.Context, req *types.UnlinkAccountRequest) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.post(ctx, "/accounts/unlink", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// VerifyAccountExists checks if an account exists for a document
func (c *Client) VerifyAccountExists(ctx context.Context, req *types.VerifyAccountExistsRequest) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.post(ctx, "/accounts/verify/exists", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAccountBalance retrieves account balance
func (c *Client) GetAccountBalance(ctx context.Context, accountID int64) (*types.BalanceResponse, error) {
	path := fmt.Sprintf("/accounts/%d/balance", accountID)
	var response types.BalanceResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAccountStatement retrieves account statement
func (c *Client) GetAccountStatement(ctx context.Context, accountID int64, params *types.StatementParams) (*types.StatementResponse, error) {
	path := fmt.Sprintf("/accounts/%d/statement", accountID)
	if params != nil {
		path += params.QueryString()
	}
	var response types.StatementResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTransactionDetails retrieves details of a specific transaction
func (c *Client) GetTransactionDetails(ctx context.Context, accountID, transactionID int64) (*types.StatementEntry, error) {
	path := fmt.Sprintf("/accounts/%d/statement/%d", accountID, transactionID)
	var response types.StatementEntry
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetCorporateAccounts retrieves corporate accounts by document
func (c *Client) GetCorporateAccounts(ctx context.Context, document string) (*types.CorporateAccountsResponse, error) {
	path := fmt.Sprintf("/accounts/%s/corporate", document)
	var response types.CorporateAccountsResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListBlockedAccounts retrieves list of blocked accounts
func (c *Client) ListBlockedAccounts(ctx context.Context) (*types.AccountListResponse, error) {
	var response types.AccountListResponse
	if err := c.get(ctx, "/accounts/list/block", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// TokenGenerateAndValidate generates or validates a token for an operation
func (c *Client) TokenGenerateAndValidate(ctx context.Context, operation, target string, req *types.TokenOperationRequest) (*types.TokenOperationResponse, error) {
	path := fmt.Sprintf("/accounts/tokens/%s/%s", operation, target)
	var response types.TokenOperationResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAccountProposalData retrieves account proposal data by account ID
func (c *Client) GetAccountProposalData(ctx context.Context, accountID int64) (*types.ProposalDataResponse, error) {
	path := fmt.Sprintf("/accounts/%d/proposalAccount/data", accountID)
	var response types.ProposalDataResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateCompanyAccount creates a new company account
func (c *Client) CreateCompanyAccount(ctx context.Context, req *types.CreateCompanyAccountRequest) (*types.CreateAccountResponse, error) {
	var response types.CreateAccountResponse
	if err := c.post(ctx, "/accounts/company", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetTransactionsByType retrieves transactions by transaction type
func (c *Client) GetTransactionsByType(ctx context.Context, transactionType string) (*types.TransactionListResponse, error) {
	path := fmt.Sprintf("/transactions/%s", transactionType)
	var response types.TransactionListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListBalanceLocks retrieves all balance locks
func (c *Client) ListBalanceLocks(ctx context.Context) (*types.BalanceLockListResponse, error) {
	var response types.BalanceLockListResponse
	if err := c.get(ctx, "/accounts/balanceLock/list", &response); err != nil {
		return nil, err
	}
	return &response, nil
}
