package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// ListProposals retrieves proposals with optional filters
func (c *Client) ListProposals(ctx context.Context, params *types.ListProposalsParams) ([]types.ProposalResponse, error) {
	path := "/proposal"
	if params != nil {
		path += params.QueryString()
	}
	var response []types.ProposalResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetProposal retrieves a specific proposal
func (c *Client) GetProposal(ctx context.Context, proposalID int64) (*types.ProposalDetailResponse, error) {
	path := fmt.Sprintf("/proposal/%d", proposalID)
	var response types.ProposalDetailResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetProposalImages retrieves proposal images
func (c *Client) GetProposalImages(ctx context.Context, proposalID int64) ([]types.ProposalImage, error) {
	path := fmt.Sprintf("/proposal/%d/images", proposalID)
	var response []types.ProposalImage
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateProposal updates a proposal
func (c *Client) UpdateProposal(ctx context.Context, proposalID int64, req *types.UpdateProposalRequest) (*types.ProposalDetailResponse, error) {
	path := fmt.Sprintf("/proposal/%d/update", proposalID)
	var response types.ProposalDetailResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateProposalImages updates proposal images
func (c *Client) UpdateProposalImages(ctx context.Context, proposalID int64, req []types.UpdateProposalImageRequest) (*types.ProposalDetailResponse, error) {
	path := fmt.Sprintf("/proposal/%d/update/images", proposalID)
	var response types.ProposalDetailResponse
	if err := c.put(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ResendProposal resends a proposal
func (c *Client) ResendProposal(ctx context.Context, proposalID int64) (*types.ProposalDetailResponse, error) {
	path := fmt.Sprintf("/proposal/%d/resend", proposalID)
	var response types.ProposalDetailResponse
	if err := c.put(ctx, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetProposalTypeStatus retrieves proposal type statuses
func (c *Client) GetProposalTypeStatus(ctx context.Context) ([]types.ProposalTypeStatus, error) {
	var response []types.ProposalTypeStatus
	if err := c.get(ctx, "/proposal/type-status", &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetLastProposal retrieves the last proposal for a document
func (c *Client) GetLastProposal(ctx context.Context, document string) (*types.ProposalDetailResponse, error) {
	path := fmt.Sprintf("/proposal/last/%s", document)
	var response types.ProposalDetailResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListLegalEntityProposals retrieves legal entity proposals
func (c *Client) ListLegalEntityProposals(ctx context.Context, params *types.ListProposalsParams) ([]types.LegalEntityProposalResponse, error) {
	path := "/proposal/legalEntities"
	if params != nil {
		path += params.QueryString()
	}
	var response []types.LegalEntityProposalResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetLegalEntityProposal retrieves a specific legal entity proposal
func (c *Client) GetLegalEntityProposal(ctx context.Context, proposalID int64) (*types.LegalEntityProposalDetailResponse, error) {
	path := fmt.Sprintf("/proposal/legalEntityProposal/%d", proposalID)
	var response types.LegalEntityProposalDetailResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
