package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Backoffice Account Operations

// ListAccountsBackoffice lists accounts with backoffice filters
func (c *Client) ListAccountsBackoffice(ctx context.Context, params *types.ListAccountsBackofficeRequest) (*types.AccountListResponse, error) {
	path := "/backoffice/accounts"
	if params != nil {
		qs := buildBackofficeAccountsQueryString(params)
		if qs != "" {
			path += "?" + qs
		}
	}
	var response types.AccountListResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ProcessProposalManually processes a proposal manually (approve/reject)
func (c *Client) ProcessProposalManually(ctx context.Context, req *types.ProposalProcessingRequest) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.post(ctx, "/backoffice/proposals/process", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateMobileAccount creates a mobile account via backoffice
func (c *Client) CreateMobileAccount(ctx context.Context, req *types.CreateMobileAccountRequest) (*types.CreateAccountResponse, error) {
	var response types.CreateAccountResponse
	if err := c.put(ctx, "/backoffice/accounts/createMobileAccount", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Biro Analysis (Credit Bureau) Operations

// CreateBiroAnalysis creates a credit bureau analysis
func (c *Client) CreateBiroAnalysis(ctx context.Context, req *types.BiroAnalysisRequest) (*types.BiroAnalysisResponse, error) {
	var response types.BiroAnalysisResponse
	if err := c.post(ctx, "/backoffice/biro/analysis", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetBiroAnalysis retrieves a biro analysis by ID
func (c *Client) GetBiroAnalysis(ctx context.Context, analysisID int64) (*types.BiroAnalysisResponse, error) {
	path := fmt.Sprintf("/backoffice/biro/analysis/%d", analysisID)
	var response types.BiroAnalysisResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateBiroAnalysis updates a biro analysis
func (c *Client) UpdateBiroAnalysis(ctx context.Context, analysisID int64, req *types.UpdateBiroAnalysisRequest) error {
	path := fmt.Sprintf("/backoffice/biro/analysis/%d", analysisID)
	return c.put(ctx, path, req, nil)
}

// Processor Operations

// BindProcessorAccount binds an account to a processor account
func (c *Client) BindProcessorAccount(ctx context.Context, req *types.BindProcessorAccountRequest) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.post(ctx, "/backoffice/processor/account/bind", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BindProcessorCard binds a card to a processor card
func (c *Client) BindProcessorCard(ctx context.Context, req *types.BindProcessorCardRequest) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.put(ctx, "/backoffice/processor/account/card", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SyncProcessorAccount synchronizes account with processor
func (c *Client) SyncProcessorAccount(ctx context.Context, accountID int64) (*types.SyncProcessorResponse, error) {
	path := fmt.Sprintf("/backoffice/processor/account/%d/synchronize", accountID)
	var response types.SyncProcessorResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PIX Scan Configuration

// GetPixScanConfiguration retrieves PIX scan configuration
func (c *Client) GetPixScanConfiguration(ctx context.Context) (*types.PixScanConfigurationResponse, error) {
	var response types.PixScanConfigurationResponse
	if err := c.get(ctx, "/backoffice/pix/scan/configuration", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdatePixScanConfiguration updates PIX scan configuration
func (c *Client) UpdatePixScanConfiguration(ctx context.Context, req *types.UpdatePixScanConfigurationRequest) error {
	return c.put(ctx, "/backoffice/pix/scan/configuration", req, nil)
}

// HCE Device Operations

// ListHceDevices lists HCE devices
func (c *Client) ListHceDevices(ctx context.Context, params *types.ListHceDevicesParams) ([]types.HceDeviceResponse, error) {
	path := "/backoffice/hce/devices"
	if params != nil {
		qs := params.QueryString()
		if qs != "" {
			path += qs
		}
	}
	var response []types.HceDeviceResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetHceDevice retrieves an HCE device by ID
func (c *Client) GetHceDevice(ctx context.Context, deviceID string) (*types.HceDeviceResponse, error) {
	path := fmt.Sprintf("/backoffice/hce/devices/%s", deviceID)
	var response types.HceDeviceResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BlockHceDevice blocks an HCE device
func (c *Client) BlockHceDevice(ctx context.Context, deviceID string) error {
	path := fmt.Sprintf("/backoffice/hce/devices/%s/block", deviceID)
	return c.post(ctx, path, nil, nil)
}

// UnblockHceDevice unblocks an HCE device
func (c *Client) UnblockHceDevice(ctx context.Context, deviceID string) error {
	path := fmt.Sprintf("/backoffice/hce/devices/%s/unblock", deviceID)
	return c.post(ctx, path, nil, nil)
}

// Daily Statement Operations

// GetDailyStatement retrieves daily statement for a date
func (c *Client) GetDailyStatement(ctx context.Context, date string) (*types.DailyStatementResponse, error) {
	path := fmt.Sprintf("/backoffice/statements/daily/%s", date)
	var response types.DailyStatementResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Issuer Balance Operations

// GetIssuerBalance retrieves total issuer balance
func (c *Client) GetIssuerBalance(ctx context.Context) (*types.IssuerBalanceResponse, error) {
	var response types.IssuerBalanceResponse
	if err := c.get(ctx, "/backoffice/issuer/balance", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ResetAccountLoginTime resets the login time for an account
func (c *Client) ResetAccountLoginTime(ctx context.Context, accountID int64) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/backoffice/accounts/resetLoginTime/%d", accountID)
	var response types.GenericResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SyncProcessorCard synchronizes card with processor
func (c *Client) SyncProcessorCard(ctx context.Context, cardID int64) (*types.SyncProcessorResponse, error) {
	path := fmt.Sprintf("/backoffice/processor/card/%d/synchronize", cardID)
	var response types.SyncProcessorResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListHceOverview retrieves HCE overview/list
func (c *Client) ListHceOverview(ctx context.Context) ([]types.HceDeviceResponse, error) {
	var response []types.HceDeviceResponse
	if err := c.get(ctx, "/backoffice/hce", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// ListDailyStatements retrieves list of daily statements
func (c *Client) ListDailyStatements(ctx context.Context) (*types.DailyStatementListResponse, error) {
	var response types.DailyStatementListResponse
	if err := c.get(ctx, "/dailyStatement/list", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteAllDailyStatements deletes all daily statements
func (c *Client) DeleteAllDailyStatements(ctx context.Context) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.delete(ctx, "/dailyStatement/all", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListBiroAnalyses retrieves all biro analyses
func (c *Client) ListBiroAnalyses(ctx context.Context) ([]types.BiroAnalysisResponse, error) {
	var response []types.BiroAnalysisResponse
	if err := c.get(ctx, "/backoffice/accounts/biroAnalysis", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetBiroAnalysisByProposal retrieves biro analysis by proposal ID
func (c *Client) GetBiroAnalysisByProposal(ctx context.Context, proposalID int64) (*types.BiroAnalysisResponse, error) {
	path := fmt.Sprintf("/backoffice/accounts/biroAnalysis/proposal/%d", proposalID)
	var response types.BiroAnalysisResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeletePaysmartProduct deletes a paysmart product
func (c *Client) DeletePaysmartProduct(ctx context.Context, productID int64) error {
	path := fmt.Sprintf("/backoffice/product/paysmart/%d", productID)
	return c.delete(ctx, path, nil)
}

// Helper function to build backoffice accounts query string
func buildBackofficeAccountsQueryString(params *types.ListAccountsBackofficeRequest) string {
	if params == nil {
		return ""
	}
	var parts []string
	if params.Document != nil {
		parts = append(parts, fmt.Sprintf("document=%s", *params.Document))
	}
	if params.Name != nil {
		parts = append(parts, fmt.Sprintf("name=%s", *params.Name))
	}
	if params.Email != nil {
		parts = append(parts, fmt.Sprintf("email=%s", *params.Email))
	}
	if params.Status != nil {
		parts = append(parts, fmt.Sprintf("status=%s", *params.Status))
	}
	if params.Page != nil {
		parts = append(parts, fmt.Sprintf("page=%d", *params.Page))
	}
	if params.PageSize != nil {
		parts = append(parts, fmt.Sprintf("pageSize=%d", *params.PageSize))
	}
	if len(parts) == 0 {
		return ""
	}
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += "&"
		}
		result += p
	}
	return result
}
