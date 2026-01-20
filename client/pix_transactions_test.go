package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// TestDoPixPayment tests executing a PIX payment
func TestDoPixPayment(t *testing.T) {
	reqData := &types.PixPaymentRequest{
		AccountID:                12345,
		RecipientInstitutionCode: "001",
		RecipientBranchCode:      "0001",
		RecipientAccountNumber:   "123456",
		RecipientAccountType:     types.PixAccountTypeCACC,
		RecipientCpfCnpj:         "12345678901",
		RecipientName:            "John Doe",
		OperationAmount:          100.50,
	}
	respData := &types.PixPaymentResponse{
		Message:            "Payment successful",
		IDTransaction:      9876,
		AuthenticationCode: "AUTH123",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/transactions/payment" {
			t.Errorf("expected path /pix/transactions/payment, got %s", r.URL.Path)
		}

		var body types.PixPaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.AccountID != reqData.AccountID {
			t.Errorf("request AccountID = %d; want %d", body.AccountID, reqData.AccountID)
		}
		if body.OperationAmount != reqData.OperationAmount {
			t.Errorf("request OperationAmount = %f; want %f", body.OperationAmount, reqData.OperationAmount)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.DoPixPayment(context.Background(), reqData)
	if err != nil {
		t.Fatalf("DoPixPayment failed: %v", err)
	}

	if response.Message != respData.Message {
		t.Errorf("response Message = %q; want %q", response.Message, respData.Message)
	}
	if response.IDTransaction != respData.IDTransaction {
		t.Errorf("response IDTransaction = %d; want %d", response.IDTransaction, respData.IDTransaction)
	}
	if response.AuthenticationCode != respData.AuthenticationCode {
		t.Errorf("response AuthenticationCode = %q; want %q", response.AuthenticationCode, respData.AuthenticationCode)
	}
}

// TestDoPixChargeback tests executing a PIX chargeback
func TestDoPixChargeback(t *testing.T) {
	reqData := &types.PixChargebackRequest{
		AccountID:     12345,
		IDTransaction: 9876,
		Amount:        100.50,
		ReasonCode:    "MD06",
	}
	respData := &types.PixChargebackResponse{
		Message: "Chargeback successful",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/transactions/chargeback" {
			t.Errorf("expected path /pix/transactions/chargeback, got %s", r.URL.Path)
		}

		var body types.PixChargebackRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.IDTransaction != reqData.IDTransaction {
			t.Errorf("request IDTransaction = %d; want %d", body.IDTransaction, reqData.IDTransaction)
		}
		if body.ReasonCode != reqData.ReasonCode {
			t.Errorf("request ReasonCode = %q; want %q", body.ReasonCode, reqData.ReasonCode)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.DoPixChargeback(context.Background(), reqData)
	if err != nil {
		t.Fatalf("DoPixChargeback failed: %v", err)
	}

	if response.Message != respData.Message {
		t.Errorf("response Message = %q; want %q", response.Message, respData.Message)
	}
}

// TestCancelPixSchedule tests canceling a scheduled PIX transaction
func TestCancelPixSchedule(t *testing.T) {
	reqData := &types.PixCancelScheduleRequest{
		AccountID:  12345,
		ScheduleID: 555,
	}
	respData := &types.PixCancelScheduleResponse{
		AccountID: int64(12345),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/transactions/cancelSchedule" {
			t.Errorf("expected path /pix/transactions/cancelSchedule, got %s", r.URL.Path)
		}

		var body types.PixCancelScheduleRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.ScheduleID != reqData.ScheduleID {
			t.Errorf("request ScheduleID = %d; want %d", body.ScheduleID, reqData.ScheduleID)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CancelPixSchedule(context.Background(), reqData)
	if err != nil {
		t.Fatalf("CancelPixSchedule failed: %v", err)
	}

	if response.AccountID == nil {
		t.Error("expected AccountID to be set")
	}
}

// TestCreatePrecautionaryBlock tests creating a precautionary block
func TestCreatePrecautionaryBlock(t *testing.T) {
	reqData := &types.PixPrecautionaryBlockRequest{
		IDTransaction: 9876,
	}
	respData := &types.PixPrecautionaryBlockResponse{
		IDAccount:     12345,
		IDTransaction: 9876,
		Message:       "Block created",
		Value:         100.50,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/backoffice/precautionaryBlock" {
			t.Errorf("expected path /pix/backoffice/precautionaryBlock, got %s", r.URL.Path)
		}

		var body types.PixPrecautionaryBlockRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.IDTransaction != reqData.IDTransaction {
			t.Errorf("request IDTransaction = %d; want %d", body.IDTransaction, reqData.IDTransaction)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CreatePrecautionaryBlock(context.Background(), reqData)
	if err != nil {
		t.Fatalf("CreatePrecautionaryBlock failed: %v", err)
	}

	if response.IDTransaction != respData.IDTransaction {
		t.Errorf("response IDTransaction = %d; want %d", response.IDTransaction, respData.IDTransaction)
	}
	if response.Message != respData.Message {
		t.Errorf("response Message = %q; want %q", response.Message, respData.Message)
	}
}

// TestUpdatePrecautionaryBlock tests updating a precautionary block
func TestUpdatePrecautionaryBlock(t *testing.T) {
	reqData := &types.PixUpdatePrecautionaryBlockRequest{
		IDOnlineTransactionLog: 111,
		PixPrecautionaryEnum:   "UNBLOCK",
	}
	respData := &types.PixUpdatePrecautionaryBlockResponse{
		Message: "Block updated",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/backoffice/precautionaryBlock/update" {
			t.Errorf("expected path /pix/backoffice/precautionaryBlock/update, got %s", r.URL.Path)
		}

		var body types.PixUpdatePrecautionaryBlockRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.IDOnlineTransactionLog != reqData.IDOnlineTransactionLog {
			t.Errorf("request IDOnlineTransactionLog = %d; want %d", body.IDOnlineTransactionLog, reqData.IDOnlineTransactionLog)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.UpdatePrecautionaryBlock(context.Background(), reqData)
	if err != nil {
		t.Fatalf("UpdatePrecautionaryBlock failed: %v", err)
	}

	if response.Message != respData.Message {
		t.Errorf("response Message = %q; want %q", response.Message, respData.Message)
	}
}

// TestGetPixTransactionLimit tests retrieving PIX transaction limits
func TestGetPixTransactionLimit(t *testing.T) {
	accountID := int64(12345)
	respData := &types.PixGetLimitResponse{
		Message: "Limits retrieved",
		PixLimitInternal: &types.PixTransactionLimit{
			TotalAvailableLimit:                   100000,
			TotalAvailableLimitFormatted:          1000.00,
			LimitPerAvailableTransaction:          10000,
			LimitPerAvailableTransactionFormatted: 100.00,
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/pix/transactions/12345/limit"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetPixTransactionLimit(context.Background(), accountID)
	if err != nil {
		t.Fatalf("GetPixTransactionLimit failed: %v", err)
	}

	if response.Message != respData.Message {
		t.Errorf("response Message = %q; want %q", response.Message, respData.Message)
	}
	if response.PixLimitInternal == nil {
		t.Error("expected PixLimitInternal to be set")
	}
}

// TestGetPixPaymentByE2E tests retrieving a PIX payment by E2E
func TestGetPixPaymentByE2E(t *testing.T) {
	e2e := "E123456789202401010000000001"
	respData := &types.GetPixInfoResponse{
		TransactionID:      9876,
		AuthenticationCode: "AUTH123",
		EndToEnd:           e2e,
		Status:             "APPROVED",
		Success:            true,
		ResultDescription:  "Payment successful",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/pix/transactions/payment/E123456789202401010000000001"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetPixPaymentByE2E(context.Background(), e2e)
	if err != nil {
		t.Fatalf("GetPixPaymentByE2E failed: %v", err)
	}

	if response.TransactionID != respData.TransactionID {
		t.Errorf("response TransactionID = %d; want %d", response.TransactionID, respData.TransactionID)
	}
	if response.EndToEnd != respData.EndToEnd {
		t.Errorf("response EndToEnd = %q; want %q", response.EndToEnd, respData.EndToEnd)
	}
	if !response.Success {
		t.Error("expected Success to be true")
	}
}

// TestListPSPs tests retrieving the list of PIX Service Providers
func TestListPSPs(t *testing.T) {
	respData := &types.PspListResponse{
		PSPs: []types.PspResponse{
			{
				CodIspb:                  "12345678",
				NomeParticipante:         "Bank Name",
				NomeResumido:             "BANK",
				TipoParticipante:         "DIRETO",
				SituacaoParticipante:     "ATIVO",
				DataHoraEnvioMensagem:    "2024-01-01T00:00:00Z",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/psps" {
			t.Errorf("expected path /pix/psps, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListPSPs(context.Background())
	if err != nil {
		t.Fatalf("ListPSPs failed: %v", err)
	}

	if len(response.PSPs) != 1 {
		t.Errorf("expected 1 PSP, got %d", len(response.PSPs))
	}
	if response.PSPs[0].CodIspb != "12345678" {
		t.Errorf("response PSPs[0].CodIspb = %q; want %q", response.PSPs[0].CodIspb, "12345678")
	}
}

// TestGetPixKeyInfo tests retrieving PIX key information
func TestGetPixKeyInfo(t *testing.T) {
	accountID := int64(12345)
	key := "12345678901"
	respData := &types.SearchKeyResponse{
		Success:           true,
		ResultDescription: "Key found",
		Chave:             &key,
		TipoChave:         strPtr("CPF"),
		Nome:              strPtr("John Doe"),
		CpfCnpj:           &key,
		Instituicao:       strPtr("001"),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/pix/keys/12345/12345678901"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(respData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetPixKeyInfo(context.Background(), accountID, key)
	if err != nil {
		t.Fatalf("GetPixKeyInfo failed: %v", err)
	}

	if !response.Success {
		t.Error("expected Success to be true")
	}
	if response.Chave == nil || *response.Chave != key {
		t.Errorf("response Chave = %v; want %q", response.Chave, key)
	}
}

