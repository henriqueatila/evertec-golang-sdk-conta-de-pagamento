package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// DoRecharge executes a mobile recharge
func (c *Client) DoRecharge(ctx context.Context, accountID int64, req *types.DoRechargeRequest) (*types.DoRechargeResponse, error) {
	path := fmt.Sprintf("/accounts/%d/recharges", accountID)
	var response types.DoRechargeResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRechargeValues retrieves available recharge amounts for a phone
func (c *Client) GetRechargeValues(ctx context.Context, accountID int64, areaCode, phoneNumber string) (*types.RechargeValuesResponse, error) {
	path := fmt.Sprintf("/accounts/%d/recharges/availableValues/%s/%s", accountID, areaCode, phoneNumber)
	var response types.RechargeValuesResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DoVoucherRecharge executes an electronic voucher recharge
func (c *Client) DoVoucherRecharge(ctx context.Context, accountID int64, req *types.DoVoucherRechargeRequest) (*types.DoRechargeResponse, error) {
	path := fmt.Sprintf("/accounts/%d/eletronicVouchers", accountID)
	var response types.DoRechargeResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetVoucherProviders retrieves voucher providers list
func (c *Client) GetVoucherProviders(ctx context.Context, accountID int64) (*types.VoucherProvidersResponse, error) {
	path := fmt.Sprintf("/accounts/%d/eletronicVouchers/providers", accountID)
	var response types.VoucherProvidersResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
