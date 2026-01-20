package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Income Report Operations (Informe de Rendimentos)

// GenerateIncomeReport generates an income report for a specific year
func (c *Client) GenerateIncomeReport(ctx context.Context, accountID int64, year int) (*types.IncomeReportResponse, error) {
	path := fmt.Sprintf("/accounts/%d/issuer-report/%d", accountID, year)
	var response types.IncomeReportResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAccountBalanceByYear retrieves account balance for a specific year
func (c *Client) GetAccountBalanceByYear(ctx context.Context, accountID int64, year int) (*types.YearlyBalanceResponse, error) {
	path := fmt.Sprintf("/accounts/%d/balance/%d", accountID, year)
	var response types.YearlyBalanceResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetAllAccountsBalanceByYear retrieves all accounts balance for a specific year
func (c *Client) GetAllAccountsBalanceByYear(ctx context.Context, year int) (*types.AllAccountsYearlyBalanceResponse, error) {
	path := fmt.Sprintf("/accounts/balance/%d", year)
	var response types.AllAccountsYearlyBalanceResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
