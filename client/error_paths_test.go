package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// TestAuthorizerErrorPaths tests error paths for authorizer methods
func TestAuthorizerErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "DoSummaryPurchase_error", method: http.MethodPost, path: "/summary-purchases",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoSummaryPurchase(ctx, &types.SummaryPurchaseRequest{})
				return err
			},
		},
		{
			name: "CancelSummaryPurchase_error", method: http.MethodPost, path: "/summary-purchases/cancel",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelSummaryPurchase(ctx, &types.CancelPurchaseRequest{})
				return err
			},
		},
		{
			name: "DoSummaryChargeback_error", method: http.MethodPost, path: "/summary-chargebacks",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoSummaryChargeback(ctx, &types.ChargebackRequest{})
				return err
			},
		},
		{
			name: "CancelSummaryChargeback_error", method: http.MethodPost, path: "/summary-chargebacks/cancel",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelSummaryChargeback(ctx, &types.CancelChargebackRequest{})
				return err
			},
		},
	})
}

// TestEnumsErrorPaths tests error paths for enum methods
func TestEnumsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetStates_error", method: http.MethodGet, path: "/enum/uf",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetStates(ctx)
				return err
			},
		},
		{
			name: "GetProfessions_error", method: http.MethodGet, path: "/enum/profession",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProfessions(ctx)
				return err
			},
		},
		{
			name: "GetIssuingAuthorities_error", method: http.MethodGet, path: "/enum/issuingAuthority",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetIssuingAuthorities(ctx)
				return err
			},
		},
		{
			name: "GetGenders_error", method: http.MethodGet, path: "/enum/gender",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetGenders(ctx)
				return err
			},
		},
		{
			name: "GetCountries_error", method: http.MethodGet, path: "/country",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetCountries(ctx)
				return err
			},
		},
		{
			name: "GetAllBanks_error", method: http.MethodGet, path: "/banco",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetAllBanks(ctx)
				return err
			},
		},
	})
}

// TestInstitutionErrorPaths tests error paths for institution methods
func TestInstitutionErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateInstitution_error", method: http.MethodPost, path: "/institutions",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateInstitution(ctx, &types.InstitutionRequest{})
				return err
			},
		},
		{
			name: "ListInstitutions_error", method: http.MethodGet, path: "/institutions",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInstitutions(ctx)
				return err
			},
		},
	})
}

// TestBranchErrorPaths tests error paths for branch methods
func TestBranchErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateBranch_error", method: http.MethodPost, path: "/branches",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBranch(ctx, &types.BranchRequest{})
				return err
			},
		},
		{
			name: "ListBranches_error", method: http.MethodGet, path: "/branches",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBranches(ctx)
				return err
			},
		},
	})
}


// TestBillsErrorPaths tests error paths for bills methods
func TestBillsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "PayBill_error", method: http.MethodPut, path: "/bill/payment",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PayBill(ctx, &types.BillPaymentRequest{})
				return err
			},
		},
		{
			name: "PayBillBatch_error", method: http.MethodPut, path: "/bill/payment/batch",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PayBillBatch(ctx, &types.BillPaymentRequest{})
				return err
			},
		},
		{
			name: "GetBillInfo_error", method: http.MethodPost, path: "/bill/info",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetBillInfo(ctx, &types.GetBillInfoRequest{})
				return err
			},
		},
	})
}

// TestBackofficeAdditionalErrorPaths tests more backoffice error paths
func TestBackofficeAdditionalErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetPixScanConfiguration_error", method: http.MethodGet, path: "/backoffice/pix/scan/configuration",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetPixScanConfiguration(ctx)
				return err
			},
		},
		{
			name: "GetIssuerBalance_error", method: http.MethodGet, path: "/backoffice/issuer/balance",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetIssuerBalance(ctx)
				return err
			},
		},
		{
			name: "ListHceOverview_error", method: http.MethodGet, path: "/backoffice/hce",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListHceOverview(ctx)
				return err
			},
		},
		{
			name: "ListDailyStatements_error", method: http.MethodGet, path: "/dailyStatement/list",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListDailyStatements(ctx)
				return err
			},
		},
		{
			name: "DeleteAllDailyStatements_error", method: http.MethodDelete, path: "/dailyStatement/all",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DeleteAllDailyStatements(ctx)
				return err
			},
		},
		{
			name: "ListBiroAnalyses_error", method: http.MethodGet, path: "/backoffice/accounts/biroAnalysis",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListBiroAnalyses(ctx)
				return err
			},
		},
		{
			name: "ProcessProposalManually_error", method: http.MethodPost, path: "/backoffice/proposals/process",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ProcessProposalManually(ctx, &types.ProposalProcessingRequest{})
				return err
			},
		},
		{
			name: "CreateMobileAccount_error", method: http.MethodPut, path: "/backoffice/accounts/createMobileAccount",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateMobileAccount(ctx, &types.CreateMobileAccountRequest{})
				return err
			},
		},
		{
			name: "CreateBiroAnalysis_error", method: http.MethodPost, path: "/backoffice/biro/analysis",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBiroAnalysis(ctx, &types.BiroAnalysisRequest{})
				return err
			},
		},
		{
			name: "BindProcessorAccount_error", method: http.MethodPost, path: "/backoffice/processor/account/bind",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BindProcessorAccount(ctx, &types.BindProcessorAccountRequest{})
				return err
			},
		},
		{
			name: "BindProcessorCard_error", method: http.MethodPut, path: "/backoffice/processor/account/card",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BindProcessorCard(ctx, &types.BindProcessorCardRequest{})
				return err
			},
		},
	})
}

// TestBanksAdditionalErrorPaths tests more banks error paths
func TestBanksAdditionalErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CheckAPIStatus_error", method: http.MethodGet, path: "/status",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CheckAPIStatus(ctx)
				return err
			},
		},
		{
			name: "CheckIntegrationStatus_error", method: http.MethodGet, path: "/status/integrationModules",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CheckIntegrationStatus(ctx)
				return err
			},
		},
	})
}

// TestMedAdditionalErrorPaths tests additional MED error paths
func TestMedAdditionalErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateInfractionReport_error", method: http.MethodPost, path: "/pix/infraction-reports",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateInfractionReport(ctx, &types.InfractionReportRequest{})
				return err
			},
		},
		{
			name: "CloseInfractionReport_error", method: http.MethodPost, path: "/pix/infraction-reports/close",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CloseInfractionReport(ctx, &types.CloseInfractionReportRequest{})
				return err
			},
		},
		{
			name: "CreateRefundSolicitation_error", method: http.MethodPost, path: "/pix/refunds/create",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateRefundSolicitation(ctx, &types.RefundSolicitationRequest{})
				return err
			},
		},
		{
			name: "CloseRefundSolicitation_error", method: http.MethodPost, path: "/pix/refunds/close",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CloseRefundSolicitation(ctx, &types.CloseRefundRequest{})
				return err
			},
		},
	})
}

// TestPixAdditionalErrorPaths tests additional Pix error paths
func TestPixAdditionalErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "AddPixDevice_error", method: http.MethodPost, path: "/pix/devices",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AddPixDevice(ctx, &types.PixDeviceRequest{})
				return err
			},
		},
		{
			name: "BlockPixDevice_error", method: http.MethodPut, path: "/pix/devices/block",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BlockPixDevice(ctx, &types.BlockPixDeviceRequest{})
				return err
			},
		},
		{
			name: "UnblockPixDevice_error", method: http.MethodPut, path: "/pix/devices/unblock",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UnblockPixDevice(ctx, &types.UnblockPixDeviceRequest{})
				return err
			},
		},
		{
			name: "ListPixClaims_error", method: http.MethodPost, path: "/accounts/pix/claim/list",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListPixClaims(ctx, &types.ListPixClaimsRequest{})
				return err
			},
		},
		{
			name: "GetRaiseLimitRequests_error", method: http.MethodGet, path: "/accounts/pix/limit/getRaiseLimitRequests",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRaiseLimitRequests(ctx)
				return err
			},
		},
		{
			name: "GetMaximumPixLimitIssuer_error", method: http.MethodGet, path: "/accounts/pix/limit/getMaximumLimitIssuer",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetMaximumPixLimitIssuer(ctx)
				return err
			},
		},
		{
			name: "ReceivePixCallback_error", method: http.MethodPost, path: "/pix/callbacks/receive-transaction",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ReceivePixCallback(ctx, &types.PixCallbackRequest{})
				return err
			},
		},
	})
}

// TestBankslipsAdditionalErrorPaths tests bankslip v2 path fix
func TestBankslipsV2ErrorPath(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateBankslipV2_error_fixed", method: http.MethodPost, path: "/bankslip/v2/generate",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateBankslipV2(ctx, &types.BankslipV2Request{})
				return err
			},
		},
	})
}

// TestPixAutomaticAdditionalErrorPaths tests error paths for pix_automatic methods
func TestPixAutomaticAdditionalErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "StartAutomaticPix_error", method: http.MethodPost, path: "/pix/automatic",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.StartAutomaticPix(ctx, &types.StartAutomaticPixRequest{})
				return err
			},
		},
		{
			name: "RejectAutomaticPix_error", method: http.MethodPost, path: "/pix/automatic/reject",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.RejectAutomaticPix(ctx, &types.RejectAutomaticPixRequest{})
				return err
			},
		},
		{
			name: "AcceptQRCodeJourneyThree_error", method: http.MethodPost, path: "/pix/automatic/qr-code/journey-three/accept",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AcceptQRCodeJourneyThree(ctx, &types.QRCodeAcceptJourneyThreeRequest{})
				return err
			},
		},
		{
			name: "AcceptAutomaticPixQRCode_error", method: http.MethodPost, path: "/pix/automatic/qr-code/accept",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AcceptAutomaticPixQRCode(ctx, &types.QRCodeUserAcceptRequest{})
				return err
			},
		},
		{
			name: "CreateAutomaticPixContract_error", method: http.MethodPost, path: "/pix/automatic/contract",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateAutomaticPixContract(ctx, &types.CreateAutomaticPixContractRequest{})
				return err
			},
		},
		{
			name: "CancelAutomaticPixCharge_error", method: http.MethodPost, path: "/pix/automatic/charge/cancel",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelAutomaticPixCharge(ctx, &types.CancelAutomaticPixChargeRequest{})
				return err
			},
		},
		{
			name: "CancelAutomaticPix_error", method: http.MethodPost, path: "/pix/automatic/cancel",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelAutomaticPix(ctx, &types.CancelAutomaticPixRequest{})
				return err
			},
		},
		{
			name: "AcceptAutomaticPix_error", method: http.MethodPost, path: "/pix/automatic/accept",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AcceptAutomaticPix(ctx, &types.AcceptAutomaticPixRequest{})
				return err
			},
		},
	})
}

// TestPixQRCodeErrorPaths tests error paths for pix_qrcode methods
func TestPixQRCodeErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreateStaticQRCode_error", method: http.MethodPost, path: "/pix/qrcodes/static",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateStaticQRCode(ctx, &types.StaticQRCodeRequest{})
				return err
			},
		},
		{
			name: "CreateDynamicQRCode_error", method: http.MethodPost, path: "/pix/qrcodes/dynamic",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateDynamicQRCode(ctx, &types.DynamicQRCodeRequest{})
				return err
			},
		},
		{
			name: "QueryQRCodeProcessing_error", method: http.MethodPost, path: "/pix/qrcodes/query-processing",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.QueryQRCodeProcessing(ctx, &types.QRCodeQueryRequest{})
				return err
			},
		},
		{
			name: "DecodeQRCodeV3_error", method: http.MethodPost, path: "/pix/qrcodes/v3/query-processing",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DecodeQRCodeV3(ctx, &types.DecodeQRCodeV3Request{})
				return err
			},
		},
	})
}

// TestPixTransactionsErrorPaths tests error paths for pix_transactions methods
func TestPixTransactionsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "DoPixPayment_error", method: http.MethodPost, path: "/pix/transactions/payment",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoPixPayment(ctx, &types.PixPaymentRequest{})
				return err
			},
		},
		{
			name: "DoPixChargeback_error", method: http.MethodPost, path: "/pix/transactions/chargeback",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DoPixChargeback(ctx, &types.PixChargebackRequest{})
				return err
			},
		},
		{
			name: "CancelPixSchedule_error", method: http.MethodPost, path: "/pix/transactions/cancelSchedule",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelPixSchedule(ctx, &types.PixCancelScheduleRequest{})
				return err
			},
		},
		{
			name: "CreatePrecautionaryBlock_error", method: http.MethodPost, path: "/pix/backoffice/precautionaryBlock",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePrecautionaryBlock(ctx, &types.PixPrecautionaryBlockRequest{})
				return err
			},
		},
		{
			name: "UpdatePrecautionaryBlock_error", method: http.MethodPost, path: "/pix/backoffice/precautionaryBlock/update",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdatePrecautionaryBlock(ctx, &types.PixUpdatePrecautionaryBlockRequest{})
				return err
			},
		},
		{
			name: "ListPSPs_error", method: http.MethodGet, path: "/pix/psps",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListPSPs(ctx)
				return err
			},
		},
	})
}

// TestPostPaidErrorPaths tests error paths for postpaid methods
func TestPostPaidErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "PostPaidPaymentBalance_error", method: http.MethodPost, path: "/postpaid/payment/balance",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidPaymentBalance(ctx, &types.PostPaidPaymentBalanceRequest{})
				return err
			},
		},
		{
			name: "PostPaidInstallmentSimulation_error", method: http.MethodPost, path: "/postpaid/payment/installment/simulation",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidInstallmentSimulation(ctx, &types.PostPaidInstallmentSimulationRequest{})
				return err
			},
		},
		{
			name: "PostPaidInstallmentPix_error", method: http.MethodPost, path: "/postpaid/payment/installment/request/pix",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidInstallmentPix(ctx, &types.PostPaidInstallmentRequest{})
				return err
			},
		},
		{
			name: "PostPaidInstallmentAccountBalance_error", method: http.MethodPost, path: "/postpaid/payment/installment/request/account/balance",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.PostPaidInstallmentAccountBalance(ctx, &types.PostPaidInstallmentRequest{})
				return err
			},
		},
		{
			name: "CancelPostPaidSchedule_error", method: http.MethodPost, path: "/postpaid/payment/schedule/cancel",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CancelPostPaidSchedule(ctx, &types.CancelInvoiceScheduleRequest{})
				return err
			},
		},
	})
}

// TestPostPaidCardsErrorPaths tests error paths for postpaid_cards methods
func TestPostPaidCardsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "CreatePostPaidCard_error", method: http.MethodPost, path: "/postpaid/cards/new",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePostPaidCard(ctx, &types.CreateCardRequest{})
				return err
			},
		},
	})
}

// TestPostPaidStatementsErrorPaths tests error paths for postpaid_statements methods
func TestPostPaidStatementsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "UpdatePostPaidAccountInfo_error", method: http.MethodPost, path: "/postpaid/account",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdatePostPaidAccountInfo(ctx, &types.UpdatePostPaidAccountRequest{})
				return err
			},
		},
	})
}

// TestProductsErrorPaths tests error paths for products methods
func TestProductsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListProducts_error", method: http.MethodGet, path: "/products",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListProducts(ctx)
				return err
			},
		},
		{
			name: "CreateProduct_error", method: http.MethodPost, path: "/products",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateProduct(ctx, &types.CreateProductRequest{})
				return err
			},
		},
		{
			name: "ListPaysmartProducts_error", method: http.MethodGet, path: "/paysmart/products",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListPaysmartProducts(ctx)
				return err
			},
		},
		{
			name: "CreatePaysmartProduct_error", method: http.MethodPost, path: "/paysmart/products",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreatePaysmartProduct(ctx, &types.CreatePaysmartProductRequest{})
				return err
			},
		},
		{
			name: "SearchProductLimits_error", method: http.MethodPost, path: "/products/limits/search",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SearchProductLimits(ctx, &types.SearchProductLimitRequest{})
				return err
			},
		},
	})
}

// TestProposalsAdditionalErrorPaths tests error paths for proposals methods
func TestProposalsAdditionalErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetProposalTypeStatus_error", method: http.MethodGet, path: "/proposal/type-status",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetProposalTypeStatus(ctx)
				return err
			},
		},
	})
}

// TestRecipientsErrorPaths tests error paths for recipients methods
func TestRecipientsErrorPaths(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "SendFeedback_error", method: http.MethodPost, path: "/accounts/feedback/send",
			status: 500, response: map[string]string{"message": "server error"},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SendFeedback(ctx, &types.FeedbackRequest{})
				return err
			},
		},
	})
}
