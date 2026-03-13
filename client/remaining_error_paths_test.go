package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// TestAccountsRemainingErrorPaths tests error paths for remaining accounts methods
func TestAccountsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "LinkAccounts_error", method: http.MethodPost, path: "/accounts/1/link",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.LinkAccounts(ctx, 1, &types.LinkAccountRequest{})
				return err
			},
		},
		{
			name: "GetAccountBalance_error", method: http.MethodGet, path: "/accounts/1/balance",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountBalance(ctx, 1)
				return err
			},
		},
		{
			name: "GetTransactionDetails_error", method: http.MethodGet, path: "/accounts/1/statement/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetTransactionDetails(ctx, 1, 2)
				return err
			},
		},
		{
			name: "GetCorporateAccounts_error", method: http.MethodGet, path: "/accounts/doc123/corporate",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCorporateAccounts(ctx, "doc123")
				return err
			},
		},
		{
			name: "TokenGenerateAndValidate_error", method: http.MethodPost, path: "/accounts/tokens/generate/target",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.TokenGenerateAndValidate(ctx, "generate", "target", &types.TokenOperationRequest{})
				return err
			},
		},
		{
			name: "GetAccountProposalData_error", method: http.MethodGet, path: "/accounts/1/proposalAccount/data",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountProposalData(ctx, 1)
				return err
			},
		},
		{
			name: "GetTransactionsByType_error", method: http.MethodGet, path: "/transactions/payment",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetTransactionsByType(ctx, "payment")
				return err
			},
		},
	})
}

// TestAddressesErrorPaths tests error paths for addresses methods
func TestAddressesErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListAddresses_error", method: http.MethodGet, path: "/accounts/1/address",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAddresses(ctx, 1)
				return err
			},
		},
		{
			name: "CreateAddress_error", method: http.MethodPost, path: "/accounts/1/address",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateAddress(ctx, 1, &types.AddressRequest{})
				return err
			},
		},
		{
			name: "GetAddress_error", method: http.MethodGet, path: "/accounts/1/address/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAddress(ctx, 1, 2)
				return err
			},
		},
		{
			name: "UpdateAddress_error", method: http.MethodPut, path: "/accounts/1/address/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAddress(ctx, 1, 2, &types.AddressRequest{})
				return err
			},
		},
		{
			name: "DeleteAddress_error", method: http.MethodDelete, path: "/accounts/1/address/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DeleteAddress(ctx, 1, 2)
				return err
			},
		},
		{
			name: "LookupPostalCode_error", method: http.MethodGet, path: "/address/postalcode/12345",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.LookupPostalCode(ctx, "12345")
				return err
			},
		},
	})
}

// TestBankslipsRemainingErrorPaths tests error paths for bankslips methods
func TestBankslipsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListBankslips_error", method: http.MethodGet, path: "/accounts/1/bankslip",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBankslips(ctx, 1)
				return err
			},
		},
		{
			name: "CreateBankslip_error", method: http.MethodPost, path: "/accounts/1/bankslip",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBankslip(ctx, 1, &types.CreateBankslipRequest{})
				return err
			},
		},
		{
			name: "ListBankslipsByStatus_error", method: http.MethodGet, path: "/accounts/1/bankslip/active",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBankslipsByStatus(ctx, 1, "active")
				return err
			},
		},
		{
			name: "ListBankslipsByStatusAndDate_error", method: http.MethodGet, path: "/accounts/1/bankslip/active/2024-01-01",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBankslipsByStatusAndDate(ctx, 1, "active", "2024-01-01")
				return err
			},
		},
	})
}

// TestBillsRemainingErrorPaths tests error paths for bills methods
func TestBillsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CancelScheduledBill_error", method: http.MethodPut, path: "/bill/account/1/scheduling/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelScheduledBill(ctx, 1, 2)
				return err
			},
		},
		{
			name: "ListScheduledBills_error", method: http.MethodGet, path: "/bill/account/1/schedules",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListScheduledBills(ctx, 1)
				return err
			},
		},
		{
			name: "PayBillByAccount_error", method: http.MethodPut, path: "/accounts/1/billpayment",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PayBillByAccount(ctx, 1, &types.BillPaymentRequest{})
				return err
			},
		},
		{
			name: "GetBillInfoByAccount_error", method: http.MethodPost, path: "/accounts/1/billpayment",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBillInfoByAccount(ctx, 1, &types.GetBillInfoRequest{})
				return err
			},
		},
		{
			name: "CancelScheduledBillByAccount_error", method: http.MethodPut, path: "/accounts/1/billpayment/scheduled/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelScheduledBillByAccount(ctx, 1, 2)
				return err
			},
		},
		{
			name: "PayBillBatchByAccount_error", method: http.MethodPut, path: "/accounts/1/billpayment/batch",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PayBillBatchByAccount(ctx, 1, &types.BatchBillPaymentRequest{})
				return err
			},
		},
		{
			name: "ListScheduledBillsByAccount_error", method: http.MethodGet, path: "/accounts/1/billpayment/scheduled",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListScheduledBillsByAccount(ctx, 1)
				return err
			},
		},
	})
}

// TestCardsRemainingErrorPaths tests error paths for additional cards methods
func TestCardsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetCard_error", method: http.MethodGet, path: "/accounts/1/cards/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCard(ctx, 1, 2)
				return err
			},
		},
		{
			name: "CreateCard_error", method: http.MethodPost, path: "/accounts/1/cards/new",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateCard(ctx, 1, &types.CreateCardRequest{})
				return err
			},
		},
		{
			name: "CreateCardBackoffice_error", method: http.MethodPost, path: "/accounts/1/cards/new/backoffice",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateCardBackoffice(ctx, 1, &types.CreateCardRequest{})
				return err
			},
		},
		{
			name: "BlockCard_error", method: http.MethodPut, path: "/accounts/1/cards/2/block",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BlockCard(ctx, 1, 2, &types.BlockCardRequest{})
				return err
			},
		},
		{
			name: "UnblockCard_error", method: http.MethodPut, path: "/accounts/1/cards/2/unblock",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UnblockCard(ctx, 1, 2)
				return err
			},
		},
		{
			name: "ReissueCard_error", method: http.MethodPut, path: "/accounts/1/cards/2/reissue",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ReissueCard(ctx, 1, 2, &types.BlockCardRequest{})
				return err
			},
		},
		{
			name: "ReissueCardBackoffice_error", method: http.MethodPut, path: "/accounts/1/cards/2/reissue/backoffice",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ReissueCardBackoffice(ctx, 1, 2, &types.BlockCardRequest{})
				return err
			},
		},
		{
			name: "ActivateCard_error", method: http.MethodPut, path: "/accounts/1/cards/2/activate",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ActivateCard(ctx, 1, 2, &types.ActivateCardRequest{})
				return err
			},
		},
		{
			name: "ChangeCardPin_error", method: http.MethodPut, path: "/accounts/1/cards/2/changePin",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangeCardPin(ctx, 1, 2, &types.ChangeCardPinRequest{})
				return err
			},
		},
		{
			name: "CreateVirtualCardFromPhysical_error", method: http.MethodPost, path: "/accounts/1/cards/2/virtual",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateVirtualCardFromPhysical(ctx, 1, 2)
				return err
			},
		},
		{
			name: "GetVirtualCards_error", method: http.MethodGet, path: "/accounts/1/cards/2/virtual",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetVirtualCards(ctx, 1, 2)
				return err
			},
		},
		{
			name: "CreateVirtualCard_error", method: http.MethodPost, path: "/accounts/1/cards/virtual",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateVirtualCard(ctx, 1, &types.CreateVirtualCardRequest{})
				return err
			},
		},
		{
			name: "ListAllVirtualCards_error", method: http.MethodGet, path: "/accounts/1/cards/virtual",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAllVirtualCards(ctx, 1)
				return err
			},
		},
		{
			name: "GetCardReplacementInfo_error", method: http.MethodGet, path: "/accounts/1/cards/2/replacement",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardReplacementInfo(ctx, 1, 2)
				return err
			},
		},
		{
			name: "RequestCardReplacement_error", method: http.MethodPost, path: "/accounts/1/cards/2/replacement",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.RequestCardReplacement(ctx, 1, 2)
				return err
			},
		},
		{
			name: "BindAnonymousCard_error", method: http.MethodPost, path: "/accounts/1/cards/bindAnonymousCard",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BindAnonymousCard(ctx, 1, &types.BindAnonymousCardRequest{})
				return err
			},
		},
		{
			name: "GetCardConfiguration_error", method: http.MethodGet, path: "/accounts/1/cards/2/getCardConfiguration",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardConfiguration(ctx, 1, 2)
				return err
			},
		},
		{
			name: "GetDefaultCardConfiguration_error", method: http.MethodGet, path: "/accounts/1/cards/getDefaultCardConfiguration",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetDefaultCardConfiguration(ctx, 1)
				return err
			},
		},
		{
			name: "GetCardPaysmart_error", method: http.MethodGet, path: "/accounts/1/cards/paysmart/card-abc",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardPaysmart(ctx, 1, "card-abc")
				return err
			},
		},
	})
}

// TestContactsErrorPaths tests error paths for contacts methods
func TestContactsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListContacts_error", method: http.MethodGet, path: "/accounts/1/contact",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListContacts(ctx, 1)
				return err
			},
		},
		{
			name: "GetContactBankDetails_error", method: http.MethodGet, path: "/accounts/1/contact/2/bankdetails",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetContactBankDetails(ctx, 1, 2, "pix")
				return err
			},
		},
	})
}

// TestCreditsRemainingErrorPaths tests error paths for credits methods
func TestCreditsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "UpdateCreditExpiration_error", method: http.MethodPut, path: "/accounts/1/credits/expiration",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateCreditExpiration(ctx, 1, &types.UpdateCreditExpirationRequest{})
				return err
			},
		},
		{
			name: "GetRefundableCredits_error", method: http.MethodGet, path: "/accounts/1/credits/refund",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRefundableCredits(ctx, 1)
				return err
			},
		},
		{
			name: "GetExpiredCredits_error", method: http.MethodGet, path: "/accounts/1/credits/expired",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetExpiredCredits(ctx, 1)
				return err
			},
		},
	})
}

// TestDepositsErrorPaths tests error paths for deposits methods
func TestDepositsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateDepositOrder_error", method: http.MethodPost, path: "/accounts/1/deposits/order",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateDepositOrder(ctx, 1, &types.CreateDepositOrderRequest{})
				return err
			},
		},
		{
			name: "ListActiveDepositOrders_error", method: http.MethodGet, path: "/accounts/1/deposits/order/active",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListActiveDepositOrders(ctx, 1)
				return err
			},
		},
		{
			name: "CancelDepositOrder_error", method: http.MethodDelete, path: "/accounts/1/deposits/order/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelDepositOrder(ctx, 1, 2)
				return err
			},
		},
	})
}

// TestIncomeReportErrorPaths tests error paths for income_report methods
func TestIncomeReportErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GenerateIncomeReport_error", method: http.MethodGet, path: "/accounts/1/issuer-report/2024",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GenerateIncomeReport(ctx, 1, 2024)
				return err
			},
		},
		{
			name: "GetAccountBalanceByYear_error", method: http.MethodGet, path: "/accounts/1/balance/2024",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountBalanceByYear(ctx, 1, 2024)
				return err
			},
		},
		{
			name: "GetAllAccountsBalanceByYear_error", method: http.MethodGet, path: "/accounts/balance/2024",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAllAccountsBalanceByYear(ctx, 2024)
				return err
			},
		},
	})
}

// TestLimitsErrorPaths tests error paths for limits methods
func TestLimitsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetAccountLimit_error", method: http.MethodGet, path: "/accounts/limit/1/NOCTURNAL/getLimit",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountLimit(ctx, 1, "NOCTURNAL")
				return err
			},
		},
		{
			name: "UpdateAccountLimit_error", method: http.MethodPut, path: "/accounts/limit/1/NOCTURNAL",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAccountLimit(ctx, 1, "NOCTURNAL", &types.UpdateLimitRequest{})
				return err
			},
		},
		{
			name: "UpdateAccountNightTimeLimit_error", method: http.MethodPut, path: "/accounts/limit/1/NOCTURNAL/startNightTime",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAccountNightTimeLimit(ctx, 1, "NOCTURNAL", &types.UpdateNightTimeLimitRequest{})
				return err
			},
		},
		{
			name: "GetMaximumLimitIssuer_error", method: http.MethodGet, path: "/accounts/limit/1/NOCTURNAL/getMaximumLimitIssuer",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetMaximumLimitIssuer(ctx, 1, "NOCTURNAL")
				return err
			},
		},
		{
			name: "GetAccountFees_error", method: http.MethodGet, path: "/accounts/1/fees",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAccountFees(ctx, 1)
				return err
			},
		},
		{
			name: "GetCardIssuanceFee_error", method: http.MethodGet, path: "/accounts/1/fees/cardissuer",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardIssuanceFee(ctx, 1)
				return err
			},
		},
		{
			name: "GetCardReissueFee_error", method: http.MethodGet, path: "/accounts/1/fees/cardreissue",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCardReissueFee(ctx, 1)
				return err
			},
		},
		{
			name: "SearchProductLimitByType_error", method: http.MethodPost, path: "/limit/NOCTURNAL/searchProductLimit",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SearchProductLimitByType(ctx, "NOCTURNAL", &types.SearchProductLimitRequest{})
				return err
			},
		},
	})
}

// TestMedRemainingErrorPaths tests error paths for med methods
func TestMedRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetInfractionReport_error", method: http.MethodGet, path: "/pix/infraction-reports/rpt-1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetInfractionReport(ctx, "rpt-1")
				return err
			},
		},
		{
			name: "CancelInfractionReport_error", method: http.MethodPut, path: "/pix/infraction-reports/cancel/rpt-1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelInfractionReport(ctx, "rpt-1")
				return err
			},
		},
		{
			name: "GetRefundSolicitation_error", method: http.MethodGet, path: "/pix/refunds/ref-1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRefundSolicitation(ctx, "ref-1")
				return err
			},
		},
		{
			name: "CancelRefundSolicitation_error", method: http.MethodPut, path: "/pix/refunds/ref-1/cancel",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelRefundSolicitation(ctx, "ref-1")
				return err
			},
		},
	})
}

// TestPixRemainingErrorPaths tests error paths for additional pix methods
func TestPixRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreatePixKey_error", method: http.MethodPost, path: "/accounts/1/createKey",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePixKey(ctx, 1, &types.CreatePixKeyRequest{})
				return err
			},
		},
		{
			name: "GetPixKeys_error", method: http.MethodGet, path: "/accounts/1/getKeys",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPixKeys(ctx, 1)
				return err
			},
		},
		{
			name: "CreatePixClaim_error", method: http.MethodPost, path: "/accounts/1/createClaim",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePixClaim(ctx, 1, &types.CreatePixClaimRequest{})
				return err
			},
		},
		{
			name: "ConfirmPortability_error", method: http.MethodPost, path: "/accounts/1/confirmPortability",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ConfirmPortability(ctx, 1, &types.ConfirmPortabilityRequest{})
				return err
			},
		},
		{
			name: "CompletePortability_error", method: http.MethodPost, path: "/accounts/1/completePortability",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CompletePortability(ctx, 1, &types.CompletePortabilityRequest{})
				return err
			},
		},
		{
			name: "CancelPortability_error", method: http.MethodPost, path: "/accounts/1/cancelPortability",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelPortability(ctx, 1, &types.CancelPortabilityRequest{})
				return err
			},
		},
		{
			name: "GetRequestedClaims_error", method: http.MethodGet, path: "/accounts/1/getRequestedClaims",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRequestedClaims(ctx, 1)
				return err
			},
		},
		{
			name: "GetPixLimit_error", method: http.MethodGet, path: "/accounts/1/pix/getLimit",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPixLimit(ctx, 1)
				return err
			},
		},
		{
			name: "UpdatePixLimit_error", method: http.MethodPut, path: "/accounts/1/pix/limit",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdatePixLimit(ctx, 1, &types.PixLimitRequest{})
				return err
			},
		},
		{
			name: "UpdatePixNightTimeLimit_error", method: http.MethodPut, path: "/accounts/1/pix/limit/startNightTime",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdatePixNightTimeLimit(ctx, 1, &types.UpdatePixNightTimeLimitRequest{})
				return err
			},
		},
		{
			name: "ListPixDevices_error", method: http.MethodGet, path: "/pix/devices/list/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListPixDevices(ctx, 1)
				return err
			},
		},
		{
			name: "GetRaiseLimitRequestDetail_error", method: http.MethodGet, path: "/accounts/pix/limit/getDetailRaiseLimitRequest/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRaiseLimitRequestDetail(ctx, 1)
				return err
			},
		},
	})
}

// TestPixTransactionsRemainingErrorPaths tests error paths for pix_transactions remaining methods
func TestPixTransactionsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetPixTransactionLimit_error", method: http.MethodGet, path: "/pix/transactions/1/limit",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPixTransactionLimit(ctx, 1)
				return err
			},
		},
		{
			name: "GetPixPaymentByE2E_error", method: http.MethodGet, path: "/pix/transactions/payment/e2e-123",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPixPaymentByE2E(ctx, "e2e-123")
				return err
			},
		},
		{
			name: "GetPixKeyInfo_error", method: http.MethodGet, path: "/pix/keys/1/key-val",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPixKeyInfo(ctx, 1, "key-val")
				return err
			},
		},
	})
}

// TestPostPaidCardsRemainingErrorPaths tests error paths for postpaid_cards remaining methods
func TestPostPaidCardsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetPostPaidVirtualCards_error", method: http.MethodGet, path: "/postpaid/cards/1/virtual",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidVirtualCards(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidPhysicalCards_error", method: http.MethodGet, path: "/postpaid/cards/1/physical",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidPhysicalCards(ctx, 1)
				return err
			},
		},
		{
			name: "CreatePostPaidVirtualCard_error", method: http.MethodPost, path: "/postpaid/cards/1/new/virtual",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePostPaidVirtualCard(ctx, 1, &types.CreateVirtualCardRequest{})
				return err
			},
		},
		{
			name: "BlockPostPaidCard_error", method: http.MethodPost, path: "/postpaid/cards/1/block/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BlockPostPaidCard(ctx, 1, 2, &types.BlockCardRequest{})
				return err
			},
		},
		{
			name: "UnblockPostPaidCard_error", method: http.MethodPost, path: "/postpaid/cards/1/unblock/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UnblockPostPaidCard(ctx, 1, 2)
				return err
			},
		},
		{
			name: "ActivatePostPaidCard_error", method: http.MethodPost, path: "/postpaid/cards/1/activate/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ActivatePostPaidCard(ctx, 1, 2, &types.ActivateCardRequest{})
				return err
			},
		},
		{
			name: "ChangePostPaidCardPin_error", method: http.MethodPost, path: "/postpaid/cards/1/changePin/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangePostPaidCardPin(ctx, 1, 2, &types.ChangeCardPinRequest{})
				return err
			},
		},
		{
			name: "ValidatePostPaidCardPin_error", method: http.MethodPost, path: "/postpaid/cards/1/changeValidatePin/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ValidatePostPaidCardPin(ctx, 1, 2, &types.ChangeCardPinRequest{})
				return err
			},
		},
		{
			name: "GetPostPaidCardSettings_error", method: http.MethodGet, path: "/postpaid/account/1/cards/2/settings",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidCardSettings(ctx, 1, 2)
				return err
			},
		},
		{
			name: "UpdatePostPaidCardSettings_error", method: http.MethodPost, path: "/postpaid/account/1/cards/2/settings",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdatePostPaidCardSettings(ctx, 1, 2, &types.PostpaidCardSettingsRequest{})
				return err
			},
		},
		{
			name: "ResetPostPaidCardSettings_error", method: http.MethodPatch, path: "/postpaid/account/1/cards/2/settings/reset",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ResetPostPaidCardSettings(ctx, 1, 2)
				return err
			},
		},
	})
}

// TestPostPaidStatementsRemainingErrorPaths tests error paths for postpaid_statements remaining methods
func TestPostPaidStatementsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetPostPaidAccount_error", method: http.MethodGet, path: "/postpaid/account/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidAccount(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidDueDates_error", method: http.MethodGet, path: "/postpaid/account/1/dueDates",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidDueDates(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidCardDueDates_error", method: http.MethodGet, path: "/postpaid/account/1/cards/2/dueDates",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidCardDueDates(ctx, 1, 2)
				return err
			},
		},
		{
			name: "GetPostPaidOpenStatement_error", method: http.MethodGet, path: "/postpaid/statements/1/open-statement",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidOpenStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidClosedStatement_error", method: http.MethodGet, path: "/postpaid/statements/1/closed-statement",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidClosedStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidFutureStatement_error", method: http.MethodGet, path: "/postpaid/statements/1/future-statement",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidFutureStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidCombinedStatement_error", method: http.MethodGet, path: "/postpaid/statements/1/combined-statement",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidCombinedStatement(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidTransactions_error", method: http.MethodGet, path: "/postpaid/statements/1/transactions",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidTransactions(ctx, 1)
				return err
			},
		},
		{
			name: "GetPostPaidPossibleAdvances_error", method: http.MethodGet, path: "/postpaid/statements/1/possible-advance",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPostPaidPossibleAdvances(ctx, 1)
				return err
			},
		},
	})
}

// TestProfileErrorPaths tests error paths for profile methods
func TestProfileErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetProfilePicture_error", method: http.MethodGet, path: "/accounts/1/picture",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProfilePicture(ctx, 1)
				return err
			},
		},
		{
			name: "UploadProfilePicture_error", method: http.MethodPost, path: "/accounts/1/picture",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UploadProfilePicture(ctx, 1, &UploadProfilePictureRequest{})
				return err
			},
		},
		{
			name: "SaveDocumentImage_error", method: http.MethodPost, path: "/accounts/1/doc",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SaveDocumentImage(ctx, 1, &DocumentImageRequest{})
				return err
			},
		},
		{
			name: "UpdateDocumentImage_error", method: http.MethodPut, path: "/accounts/1/doc",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateDocumentImage(ctx, 1, &DocumentImageRequest{})
				return err
			},
		},
		{
			name: "GetDocumentImages_error", method: http.MethodGet, path: "/accounts/1/doc/passport/approved",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetDocumentImages(ctx, 1, "passport", "approved")
				return err
			},
		},
		{
			name: "GetCreditEngineInfo_error", method: http.MethodGet, path: "/accounts/1/creditEngineInfo",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCreditEngineInfo(ctx, 1)
				return err
			},
		},
		{
			name: "CreateCreditEngineInfo_error", method: http.MethodPost, path: "/accounts/1/creditEngineInfo",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateCreditEngineInfo(ctx, 1, &types.CreditEngineInfoRequest{})
				return err
			},
		},
	})
}

// TestQRCodeErrorPaths tests error paths for qrcode methods
func TestQRCodeErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "PayQRCode_error", method: http.MethodPost, path: "/accounts/1/qrcode/payment",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PayQRCode(ctx, 1, &types.QRCodePaymentRequest{})
				return err
			},
		},
		{
			name: "PaySimpleQRCode_error", method: http.MethodPost, path: "/accounts/1/qrcode/simplePayment",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PaySimpleQRCode(ctx, 1, &types.SimpleQRCodePaymentRequest{})
				return err
			},
		},
		{
			name: "ParseQRCode_error", method: http.MethodPost, path: "/accounts/1/qrcode/parse",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ParseQRCode(ctx, 1, &types.ParseQRCodeRequest{})
				return err
			},
		},
		{
			name: "GetQRCodePublicKey_error", method: http.MethodGet, path: "/accounts/1/qrcode/publicKey",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetQRCodePublicKey(ctx, 1)
				return err
			},
		},
	})
}

// TestRechargesErrorPaths tests error paths for recharges methods
func TestRechargesErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "DoRecharge_error", method: http.MethodPost, path: "/accounts/1/recharges",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoRecharge(ctx, 1, &types.DoRechargeRequest{})
				return err
			},
		},
		{
			name: "GetRechargeValues_error", method: http.MethodGet, path: "/accounts/1/recharges/availableValues/11/999999999",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRechargeValues(ctx, 1, "11", "999999999")
				return err
			},
		},
		{
			name: "DoVoucherRecharge_error", method: http.MethodPost, path: "/accounts/1/eletronicVouchers",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoVoucherRecharge(ctx, 1, &types.DoVoucherRechargeRequest{})
				return err
			},
		},
		{
			name: "GetVoucherProviders_error", method: http.MethodGet, path: "/accounts/1/eletronicVouchers/providers",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetVoucherProviders(ctx, 1)
				return err
			},
		},
	})
}

// TestTransfersErrorPaths tests error paths for transfers methods
func TestTransfersErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "InternalTransfer_error", method: http.MethodPost, path: "/accounts/1/transfer",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.InternalTransfer(ctx, 1, &types.InternalTransferRequest{})
				return err
			},
		},
		{
			name: "InternalTransferArrangement_error", method: http.MethodPost, path: "/accounts/1/transfer/arrangement",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.InternalTransferArrangement(ctx, 1, &types.InternalTransferRequest{})
				return err
			},
		},
		{
			name: "BankTransfer_error", method: http.MethodPost, path: "/accounts/1/banktransfer",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BankTransfer(ctx, 1, &types.BankTransferRequest{})
				return err
			},
		},
		{
			name: "CancelScheduledTransfer_error", method: http.MethodPut, path: "/accounts/1/banktransfer/scheduled/cancel/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelScheduledTransfer(ctx, 1, 2)
				return err
			},
		},
		{
			name: "ListScheduledTransfers_error", method: http.MethodGet, path: "/accounts/1/banktransfer/scheduled",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListScheduledTransfers(ctx, 1)
				return err
			},
		},
		{
			name: "GetBatchTransfers_error", method: http.MethodGet, path: "/accounts/1/transfer/batch",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBatchTransfers(ctx, 1)
				return err
			},
		},
		{
			name: "GetBatchTransferStatus_error", method: http.MethodGet, path: "/accounts/1/transfer/batch/proc-1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBatchTransferStatus(ctx, 1, "proc-1")
				return err
			},
		},
		{
			name: "CheckRecipientAccount_error", method: http.MethodGet, path: "/accounts/1/transfers/checkRecipientAccount/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CheckRecipientAccount(ctx, 1, 2)
				return err
			},
		},
		{
			name: "CancelInternalTransfer_error", method: http.MethodPost, path: "/accounts/1/cancelTransfer",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelInternalTransfer(ctx, 1, &types.CancelInternalTransferRequest{})
				return err
			},
		},
		{
			name: "TransferByID_error", method: http.MethodPost, path: "/accounts/doc123/transfer/idid",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.TransferByID(ctx, "doc123", &types.TransferByIDRequest{})
				return err
			},
		},
	})
}

// TestTravelErrorPaths tests error paths for travel methods
func TestTravelErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetTravelNotices_error", method: http.MethodGet, path: "/travel/account/1/notify",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetTravelNotices(ctx, 1)
				return err
			},
		},
		{
			name: "CreateTravelNotice_error", method: http.MethodPost, path: "/travel/account/1/notify",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateTravelNotice(ctx, 1, &types.NotifyTripRequest{})
				return err
			},
		},
		{
			name: "ChangeAccountPassword_error", method: http.MethodPut, path: "/accounts/1/changeUserPassword",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangeAccountPassword(ctx, 1, &types.ChangeUserPasswordRequest{})
				return err
			},
		},
		{
			name: "ChangeAccountStatus_error", method: http.MethodPut, path: "/accounts/1/changeStatus/ACTIVE",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ChangeAccountStatus(ctx, 1, types.AccountStatus("ACTIVE"))
				return err
			},
		},
		{
			name: "UpdateAccountName_error", method: http.MethodPut, path: "/accounts/1/name",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateAccountName(ctx, 1, &types.UpdateNameInAccountRequest{})
				return err
			},
		},
	})
}

// TestWebhooksRemainingErrorPaths tests error paths for webhooks methods
func TestWebhooksRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "NotifyStatementClosed_error", method: http.MethodPost, path: "/postpaid/notification/eventhub/statement-closed/issuer1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.NotifyStatementClosed(ctx, "issuer1", &types.EventHubRequest{})
				return err
			},
		},
		{
			name: "NotifyDueDate_error", method: http.MethodPost, path: "/postpaid/notification/eventhub/due-notification/issuer1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.NotifyDueDate(ctx, "issuer1", &types.EventHubRequest{})
				return err
			},
		},
	})
}

// TestPixAutomaticRemainingErrorPaths tests error paths for pix_automatic list methods
func TestPixAutomaticRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListAutomaticPixCharges_error", method: http.MethodGet, path: "/pix/automatic/charge/account/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAutomaticPixCharges(ctx, 1, nil)
				return err
			},
		},
		{
			name: "ListAutomaticPixByAccount_error", method: http.MethodGet, path: "/pix/automatic/account/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListAutomaticPixByAccount(ctx, 1, nil)
				return err
			},
		},
	})
}

// TestBranchRemainingErrorPaths tests error paths for branch remaining methods
func TestBranchRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetBranch_error", method: http.MethodGet, path: "/branches/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBranch(ctx, 1)
				return err
			},
		},
	})
}

// TestBanksRemainingErrorPaths tests error paths for banks remaining methods
func TestBanksRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetScheduledOperations_error", method: http.MethodGet, path: "/accounts/1/scheduleds",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetScheduledOperations(ctx, 1)
				return err
			},
		},
	})
}

// TestBackofficeRemainingErrorPaths tests error paths for backoffice remaining methods
func TestBackofficeRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetBiroAnalysis_error", method: http.MethodGet, path: "/backoffice/biro/analysis/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBiroAnalysis(ctx, 1)
				return err
			},
		},
		{
			name: "SyncProcessorAccount_error", method: http.MethodPut, path: "/backoffice/processor/account/1/synchronize",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SyncProcessorAccount(ctx, 1)
				return err
			},
		},
		{
			name: "GetHceDevice_error", method: http.MethodGet, path: "/backoffice/hce/devices/dev-1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetHceDevice(ctx, "dev-1")
				return err
			},
		},
		{
			name: "GetDailyStatement_error", method: http.MethodGet, path: "/backoffice/statements/daily/2024-01-01",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetDailyStatement(ctx, "2024-01-01")
				return err
			},
		},
		{
			name: "ResetAccountLoginTime_error", method: http.MethodPut, path: "/backoffice/accounts/resetLoginTime/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ResetAccountLoginTime(ctx, 1)
				return err
			},
		},
		{
			name: "SyncProcessorCard_error", method: http.MethodPut, path: "/backoffice/processor/card/1/synchronize",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SyncProcessorCard(ctx, 1)
				return err
			},
		},
		{
			name: "GetBiroAnalysisByProposal_error", method: http.MethodGet, path: "/backoffice/accounts/biroAnalysis/proposal/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBiroAnalysisByProposal(ctx, 1)
				return err
			},
		},
	})
}

// TestProductsRemainingErrorPaths tests error paths for products remaining methods
func TestProductsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetProduct_error", method: http.MethodGet, path: "/products/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProduct(ctx, 1)
				return err
			},
		},
		{
			name: "GetProductLimitScheduling_error", method: http.MethodGet, path: "/products/1/limit-scheduling",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProductLimitScheduling(ctx, 1)
				return err
			},
		},
		{
			name: "GetPaysmartProduct_error", method: http.MethodGet, path: "/paysmart/products/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPaysmartProduct(ctx, 1)
				return err
			},
		},
	})
}

// TestProposalsRemainingErrorPaths tests error paths for proposals remaining methods
func TestProposalsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetProposal_error", method: http.MethodGet, path: "/proposal/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProposal(ctx, 1)
				return err
			},
		},
		{
			name: "GetProposalImages_error", method: http.MethodGet, path: "/proposal/1/images",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProposalImages(ctx, 1)
				return err
			},
		},
		{
			name: "UpdateProposal_error", method: http.MethodPut, path: "/proposal/1/update",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateProposal(ctx, 1, &types.UpdateProposalRequest{})
				return err
			},
		},
		{
			name: "UpdateProposalImages_error", method: http.MethodPut, path: "/proposal/1/update/images",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateProposalImages(ctx, 1, []types.UpdateProposalImageRequest{})
				return err
			},
		},
		{
			name: "ResendProposal_error", method: http.MethodPut, path: "/proposal/1/resend",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ResendProposal(ctx, 1)
				return err
			},
		},
		{
			name: "GetLastProposal_error", method: http.MethodGet, path: "/proposal/last/doc-123",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetLastProposal(ctx, "doc-123")
				return err
			},
		},
		{
			name: "GetLegalEntityProposal_error", method: http.MethodGet, path: "/proposal/legalEntityProposal/1",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetLegalEntityProposal(ctx, 1)
				return err
			},
		},
	})
}

// TestRecipientsRemainingErrorPaths tests error paths for recipients remaining methods
func TestRecipientsRemainingErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetRecipients_error", method: http.MethodGet, path: "/accounts/1/recipients",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRecipients(ctx, 1)
				return err
			},
		},
		{
			name: "GetRecipient_error", method: http.MethodGet, path: "/accounts/1/recipients/2",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRecipient(ctx, 1, 2)
				return err
			},
		},
		{
			name: "CreateRecipient_error", method: http.MethodPost, path: "/accounts/1/recipients",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateRecipient(ctx, 1, &types.CreateRecipientRequest{})
				return err
			},
		},
		{
			name: "UpdateRecipient_error", method: http.MethodPut, path: "/accounts/1/recipients",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateRecipient(ctx, 1, &types.UpdateRecipientRequest{})
				return err
			},
		},
		{
			name: "GetLastTransactionError_error", method: http.MethodGet, path: "/accounts/1/feedback/lastTransactionError",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetLastTransactionError(ctx, 1)
				return err
			},
		},
	})
}
