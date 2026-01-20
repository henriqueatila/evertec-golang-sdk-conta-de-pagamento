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

// TestInternalTransferMethod tests the InternalTransfer method
func TestInternalTransferMethod(t *testing.T) {
	accountID := int64(12345)
	recipientID := int64(67890)
	amount := int64(50000) // R$ 500.00
	description := "Payment for services"

	expectedReq := &types.InternalTransferRequest{
		RecipientAccountID: recipientID,
		TransferAmount:     amount,
		FreeDescription:    &description,
	}

	expectedResp := types.InternalTransferResponse{
		Message:            "Transfer completed successfully",
		TransactionID:      987654,
		Amount:             amount,
		RecipientAccountID: &recipientID,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and path
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/transfer"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Verify request body
		var req types.InternalTransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if req.RecipientAccountID != expectedReq.RecipientAccountID {
			t.Errorf("recipientAccountID = %d; want %d", req.RecipientAccountID, expectedReq.RecipientAccountID)
		}
		if req.TransferAmount != expectedReq.TransferAmount {
			t.Errorf("transferAmount = %d; want %d", req.TransferAmount, expectedReq.TransferAmount)
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.InternalTransfer(context.Background(), accountID, expectedReq)
	if err != nil {
		t.Fatalf("InternalTransfer failed: %v", err)
	}

	if resp.TransactionID != expectedResp.TransactionID {
		t.Errorf("transactionID = %d; want %d", resp.TransactionID, expectedResp.TransactionID)
	}
	if resp.Message != expectedResp.Message {
		t.Errorf("message = %q; want %q", resp.Message, expectedResp.Message)
	}
}

// TestInternalTransferArrangement tests the InternalTransferArrangement method
func TestInternalTransferArrangement(t *testing.T) {
	accountID := int64(12345)
	recipientID := int64(67890)
	amount := int64(30000)

	expectedReq := &types.InternalTransferRequest{
		RecipientAccountID: recipientID,
		TransferAmount:     amount,
	}

	expectedResp := types.InternalTransferResponse{
		Message:       "Arrangement transfer completed",
		TransactionID: 111222,
		Amount:        amount,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/transfer/arrangement"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.InternalTransferArrangement(context.Background(), accountID, expectedReq)
	if err != nil {
		t.Fatalf("InternalTransferArrangement failed: %v", err)
	}

	if resp.TransactionID != expectedResp.TransactionID {
		t.Errorf("transactionID = %d; want %d", resp.TransactionID, expectedResp.TransactionID)
	}
}

// TestBankTransferMethod tests the BankTransfer method
func TestBankTransferMethod(t *testing.T) {
	accountID := int64(12345)
	bankName := "Banco Example"
	authCode := "AUTH123456"

	expectedReq := &types.BankTransferRequest{
		Recipient: types.BankTransferRecipient{
			Name:        "John Doe",
			Document:    "12345678900",
			BankCode:    "001",
			Branch:      "1234",
			Account:     "56789",
			AccountType: "CHECKING",
		},
		TransactionAmount: 100000,
	}

	expectedResp := types.BankTransferResponse{
		Message:            "Bank transfer completed",
		TransactionID:      555666,
		RecipientBankName:  &bankName,
		AuthenticationCode: &authCode,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/banktransfer"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Verify request body
		var req types.BankTransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if req.Recipient.BankCode != expectedReq.Recipient.BankCode {
			t.Errorf("bankCode = %s; want %s", req.Recipient.BankCode, expectedReq.Recipient.BankCode)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.BankTransfer(context.Background(), accountID, expectedReq)
	if err != nil {
		t.Fatalf("BankTransfer failed: %v", err)
	}

	if resp.TransactionID != expectedResp.TransactionID {
		t.Errorf("transactionID = %d; want %d", resp.TransactionID, expectedResp.TransactionID)
	}
	if resp.RecipientBankName == nil || *resp.RecipientBankName != bankName {
		t.Errorf("recipientBankName = %v; want %s", resp.RecipientBankName, bankName)
	}
}

// TestCancelScheduledTransfer tests the CancelScheduledTransfer method
func TestCancelScheduledTransfer(t *testing.T) {
	accountID := int64(12345)
	schedulingID := int64(999888)
	authCode := "CANCEL123"

	expectedResp := types.CancelTransferResponse{
		Message:            "Transfer cancelled successfully",
		AuthenticationCode: &authCode,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/banktransfer/scheduled/cancel/999888"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.CancelScheduledTransfer(context.Background(), accountID, schedulingID)
	if err != nil {
		t.Fatalf("CancelScheduledTransfer failed: %v", err)
	}

	if resp.Message != expectedResp.Message {
		t.Errorf("message = %q; want %q", resp.Message, expectedResp.Message)
	}
}

// TestListScheduledTransfers tests the ListScheduledTransfers method
func TestListScheduledTransfers(t *testing.T) {
	accountID := int64(12345)

	expectedResp := types.ScheduledTransfersResponse{
		Transfers: []types.ScheduledTransferResponse{
			{
				SchedulingID:  100,
				Amount:        50000,
				ScheduledDate: "2024-12-25",
				Status:        "PENDING",
			},
			{
				SchedulingID:  101,
				Amount:        75000,
				ScheduledDate: "2024-12-30",
				Status:        "PENDING",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/banktransfer/scheduled"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.ListScheduledTransfers(context.Background(), accountID)
	if err != nil {
		t.Fatalf("ListScheduledTransfers failed: %v", err)
	}

	if len(resp.Transfers) != len(expectedResp.Transfers) {
		t.Errorf("transfers count = %d; want %d", len(resp.Transfers), len(expectedResp.Transfers))
	}
	if resp.Transfers[0].SchedulingID != expectedResp.Transfers[0].SchedulingID {
		t.Errorf("schedulingID = %d; want %d", resp.Transfers[0].SchedulingID, expectedResp.Transfers[0].SchedulingID)
	}
}

// TestBatchInternalTransfer tests the BatchInternalTransfer method
func TestBatchInternalTransfer(t *testing.T) {
	accountID := int64(12345)
	processingCode := "BATCH123456"

	expectedReq := &types.BatchTransferRequest{
		Transfers: []types.InternalTransferRequest{
			{
				RecipientAccountID: 11111,
				TransferAmount:     10000,
			},
			{
				RecipientAccountID: 22222,
				TransferAmount:     20000,
			},
		},
	}

	expectedResp := types.BatchTransferResponse{
		ProcessingCode: processingCode,
		Status:         "PROCESSING",
		Transactions: []types.BatchTransferItem{
			{RecipientAccountID: 11111, Amount: 10000, Status: "PENDING"},
			{RecipientAccountID: 22222, Amount: 20000, Status: "PENDING"},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/transfer/batch"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Verify request body
		var req types.BatchTransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if len(req.Transfers) != len(expectedReq.Transfers) {
			t.Errorf("transfers count = %d; want %d", len(req.Transfers), len(expectedReq.Transfers))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.BatchInternalTransfer(context.Background(), accountID, expectedReq)
	if err != nil {
		t.Fatalf("BatchInternalTransfer failed: %v", err)
	}

	if resp.ProcessingCode != expectedResp.ProcessingCode {
		t.Errorf("processingCode = %q; want %q", resp.ProcessingCode, expectedResp.ProcessingCode)
	}
	if resp.Status != expectedResp.Status {
		t.Errorf("status = %q; want %q", resp.Status, expectedResp.Status)
	}
}

// TestGetBatchTransfers tests the GetBatchTransfers method
func TestGetBatchTransfers(t *testing.T) {
	accountID := int64(12345)

	expectedResp := []types.BatchTransferResponse{
		{
			ProcessingCode: "BATCH001",
			Status:         "COMPLETED",
		},
		{
			ProcessingCode: "BATCH002",
			Status:         "PROCESSING",
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/transfer/batch"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.GetBatchTransfers(context.Background(), accountID)
	if err != nil {
		t.Fatalf("GetBatchTransfers failed: %v", err)
	}

	if len(resp) != len(expectedResp) {
		t.Errorf("batch transfers count = %d; want %d", len(resp), len(expectedResp))
	}
	if resp[0].ProcessingCode != expectedResp[0].ProcessingCode {
		t.Errorf("processingCode = %q; want %q", resp[0].ProcessingCode, expectedResp[0].ProcessingCode)
	}
}

// TestGetBatchTransferStatus tests the GetBatchTransferStatus method
func TestGetBatchTransferStatus(t *testing.T) {
	accountID := int64(12345)
	processingCode := "BATCH123"

	expectedResp := types.BatchTransferResponse{
		ProcessingCode: processingCode,
		Status:         "COMPLETED",
		Transactions: []types.BatchTransferItem{
			{RecipientAccountID: 11111, Amount: 10000, Status: "COMPLETED"},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/transfer/batch/BATCH123"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.GetBatchTransferStatus(context.Background(), accountID, processingCode)
	if err != nil {
		t.Fatalf("GetBatchTransferStatus failed: %v", err)
	}

	if resp.ProcessingCode != expectedResp.ProcessingCode {
		t.Errorf("processingCode = %q; want %q", resp.ProcessingCode, expectedResp.ProcessingCode)
	}
	if resp.Status != expectedResp.Status {
		t.Errorf("status = %q; want %q", resp.Status, expectedResp.Status)
	}
}

// TestCheckRecipientAccount tests the CheckRecipientAccount method
func TestCheckRecipientAccount(t *testing.T) {
	accountID := int64(12345)
	recipientAccountID := int64(67890)
	recipientName := "Jane Doe"
	recipientDoc := "98765432100"
	accountStatus := "ACTIVE"

	expectedResp := types.CheckRecipientAccountResponse{
		Valid:             true,
		RecipientName:     &recipientName,
		RecipientDocument: &recipientDoc,
		AccountStatus:     &accountStatus,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/transfers/checkRecipientAccount/67890"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.CheckRecipientAccount(context.Background(), accountID, recipientAccountID)
	if err != nil {
		t.Fatalf("CheckRecipientAccount failed: %v", err)
	}

	if !resp.Valid {
		t.Error("expected valid recipient account")
	}
	if resp.RecipientName == nil || *resp.RecipientName != recipientName {
		t.Errorf("recipientName = %v; want %s", resp.RecipientName, recipientName)
	}
}

// TestCancelInternalTransfer tests the CancelInternalTransfer method
func TestCancelInternalTransfer(t *testing.T) {
	accountID := int64(12345)
	transactionID := int64(999888)
	authCode := "CANCELINT123"

	expectedReq := &types.CancelInternalTransferRequest{
		TransactionID: transactionID,
	}

	expectedResp := types.CancelInternalTransferResponse{
		Message:            "Internal transfer cancelled",
		AuthenticationCode: &authCode,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/cancelTransfer"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Verify request body
		var req types.CancelInternalTransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if req.TransactionID != transactionID {
			t.Errorf("transactionID = %d; want %d", req.TransactionID, transactionID)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.CancelInternalTransfer(context.Background(), accountID, expectedReq)
	if err != nil {
		t.Fatalf("CancelInternalTransfer failed: %v", err)
	}

	if resp.Message != expectedResp.Message {
		t.Errorf("message = %q; want %q", resp.Message, expectedResp.Message)
	}
	if resp.AuthenticationCode == nil || *resp.AuthenticationCode != authCode {
		t.Errorf("authenticationCode = %v; want %s", resp.AuthenticationCode, authCode)
	}
}

// TestTransferByID tests the TransferByID method
func TestTransferByID(t *testing.T) {
	document := "12345678900"
	targetAccountID := int64(67890)
	amount := int64(25000)
	merchantName := "Test Merchant"

	expectedReq := &types.TransferByIDRequest{
		TargetAccountID: targetAccountID,
		Amount:          amount,
		MerchantName:    &merchantName,
	}

	expectedResp := types.InternalTransferResponse{
		Message:            "Transfer by ID completed",
		TransactionID:      777888,
		Amount:             amount,
		RecipientAccountID: &targetAccountID,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345678900/transfer/idid"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Verify request body
		var req types.TransferByIDRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if req.TargetAccountID != targetAccountID {
			t.Errorf("targetAccountID = %d; want %d", req.TargetAccountID, targetAccountID)
		}
		if req.Amount != amount {
			t.Errorf("amount = %d; want %d", req.Amount, amount)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.TransferByID(context.Background(), document, expectedReq)
	if err != nil {
		t.Fatalf("TransferByID failed: %v", err)
	}

	if resp.TransactionID != expectedResp.TransactionID {
		t.Errorf("transactionID = %d; want %d", resp.TransactionID, expectedResp.TransactionID)
	}
	if resp.Amount != expectedResp.Amount {
		t.Errorf("amount = %d; want %d", resp.Amount, expectedResp.Amount)
	}
}

// TestInternalTransferWithLocation tests InternalTransfer with location data
func TestInternalTransferWithLocation(t *testing.T) {
	accountID := int64(12345)
	latitude := "-23.5505"
	longitude := "-46.6333"

	expectedReq := &types.InternalTransferRequest{
		RecipientAccountID: 67890,
		TransferAmount:     10000,
		Latitude:           &latitude,
		Longitude:          &longitude,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req types.InternalTransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		// Verify location data
		if req.Latitude == nil || *req.Latitude != latitude {
			t.Errorf("latitude = %v; want %s", req.Latitude, latitude)
		}
		if req.Longitude == nil || *req.Longitude != longitude {
			t.Errorf("longitude = %v; want %s", req.Longitude, longitude)
		}

		resp := types.InternalTransferResponse{
			Message:       "Transfer with location completed",
			TransactionID: 123456,
			Amount:        10000,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.InternalTransfer(context.Background(), accountID, expectedReq)
	if err != nil {
		t.Fatalf("InternalTransfer with location failed: %v", err)
	}

	if resp.TransactionID == 0 {
		t.Error("expected valid transaction ID")
	}
}

// TestBankTransferWithScheduling tests BankTransfer with scheduled date
func TestBankTransferWithScheduling(t *testing.T) {
	accountID := int64(12345)
	scheduledDate := "2024-12-25"

	expectedReq := &types.BankTransferRequest{
		Recipient: types.BankTransferRecipient{
			Name:        "Scheduled Recipient",
			Document:    "11122233344",
			BankCode:    "237",
			Branch:      "0001",
			Account:     "12345",
			AccountType: "SAVINGS",
		},
		TransactionAmount: 50000,
		SchedulingDate:    &scheduledDate,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req types.BankTransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		// Verify scheduling date
		if req.SchedulingDate == nil || *req.SchedulingDate != scheduledDate {
			t.Errorf("schedulingDate = %v; want %s", req.SchedulingDate, scheduledDate)
		}

		resp := types.BankTransferResponse{
			Message:       "Scheduled bank transfer created",
			TransactionID: 999000,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.BankTransfer(context.Background(), accountID, expectedReq)
	if err != nil {
		t.Fatalf("BankTransfer with scheduling failed: %v", err)
	}

	if resp.TransactionID == 0 {
		t.Error("expected valid transaction ID")
	}
}

// TestCheckRecipientAccountInvalid tests checking an invalid recipient account
func TestCheckRecipientAccountInvalid(t *testing.T) {
	accountID := int64(12345)
	recipientAccountID := int64(99999)

	expectedResp := types.CheckRecipientAccountResponse{
		Valid: false,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	resp, err := client.CheckRecipientAccount(context.Background(), accountID, recipientAccountID)
	if err != nil {
		t.Fatalf("CheckRecipientAccount failed: %v", err)
	}

	if resp.Valid {
		t.Error("expected invalid recipient account")
	}
}

// TestBatchInternalTransferEmpty tests batch transfer with no items
func TestBatchInternalTransferEmpty(t *testing.T) {
	accountID := int64(12345)

	expectedReq := &types.BatchTransferRequest{
		Transfers: []types.InternalTransferRequest{},
	}

	expectedResp := types.BatchTransferResponse{
		ProcessingCode: "BATCH_EMPTY",
		Status:         "INVALID",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchTransferRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		if len(req.Transfers) != 0 {
			t.Errorf("expected empty transfers, got %d items", len(req.Transfers))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(expectedResp)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	_, err = client.BatchInternalTransfer(context.Background(), accountID, expectedReq)
	if err == nil {
		t.Fatal("expected error for empty batch transfer")
	}
}
