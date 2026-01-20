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

// ===== Bill Payment Tests =====

// TestPayBill tests paying a bill
func TestPayBill(t *testing.T) {
	requestData := &types.BillPaymentRequest{
		Barcode: "34191790010104351004791020150008291070026000",
		Amount:  intPtr(5000),
	}

	responseData := &types.BillPaymentResponse{
		Message:            "Payment processed successfully",
		AuthenticationCode: strPtr("AUTH123456"),
		AmountPaid:         intPtr(5000),
		PaymentDate:        timePtr(time.Now()),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/bill/payment" {
			t.Errorf("expected path /bill/payment, got %s", r.URL.Path)
		}

		var body types.BillPaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.Barcode != requestData.Barcode {
			t.Errorf("barcode = %q; want %q", body.Barcode, requestData.Barcode)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.PayBill(context.Background(), requestData)
	if err != nil {
		t.Fatalf("PayBill failed: %v", err)
	}

	if response.Message != responseData.Message {
		t.Errorf("message = %q; want %q", response.Message, responseData.Message)
	}
}

// TestPayBillBatch tests batch bill payment
func TestPayBillBatch(t *testing.T) {
	requestData := &types.BillPaymentRequest{
		Barcode: "34191790010104351004791020150008291070026000",
		Amount:  intPtr(5000),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/bill/payment/batch" {
			t.Errorf("expected path /bill/payment/batch, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.BillPaymentResponse{Message: "Batch processed"})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.PayBillBatch(context.Background(), requestData)
	if err != nil {
		t.Fatalf("PayBillBatch failed: %v", err)
	}

	if response.Message != "Batch processed" {
		t.Errorf("message = %q; want %q", response.Message, "Batch processed")
	}
}

// TestGetBillInfo tests retrieving bill information
func TestGetBillInfo(t *testing.T) {
	requestData := &types.GetBillInfoRequest{
		Barcode: "34191790010104351004791020150008291070026000",
	}

	responseData := &types.GetBillInfoResponse{
		OriginalValue: 5000,
		DueDate:       "2024-12-31",
		Assignor:      strPtr("Test Company"),
		Payer:         strPtr("John Doe"),
		Amount:        intPtr(5000),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/bill/info" {
			t.Errorf("expected path /bill/info, got %s", r.URL.Path)
		}

		var body types.GetBillInfoRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.Barcode != requestData.Barcode {
			t.Errorf("barcode = %q; want %q", body.Barcode, requestData.Barcode)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetBillInfo(context.Background(), requestData)
	if err != nil {
		t.Fatalf("GetBillInfo failed: %v", err)
	}

	if response.OriginalValue != responseData.OriginalValue {
		t.Errorf("originalValue = %d; want %d", response.OriginalValue, responseData.OriginalValue)
	}
	if response.DueDate != responseData.DueDate {
		t.Errorf("dueDate = %q; want %q", response.DueDate, responseData.DueDate)
	}
}

// TestCancelScheduledBill tests canceling a scheduled bill
func TestCancelScheduledBill(t *testing.T) {
	accountID := int64(12345)
	schedulingID := int64(67890)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		expectedPath := "/bill/account/12345/scheduling/67890"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.CancelBillResponse{
			Message:            "Bill canceled successfully",
			AuthenticationCode: strPtr("AUTH789"),
		})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CancelScheduledBill(context.Background(), accountID, schedulingID)
	if err != nil {
		t.Fatalf("CancelScheduledBill failed: %v", err)
	}

	if response.Message != "Bill canceled successfully" {
		t.Errorf("message = %q; want %q", response.Message, "Bill canceled successfully")
	}
}

// TestListScheduledBills tests listing scheduled bills
func TestListScheduledBills(t *testing.T) {
	accountID := int64(12345)

	responseData := &types.ScheduledBillsResponse{
		Bills: []types.ScheduledBill{
			{
				SchedulingID: 1,
				Digitable:    "34191790010104351004791020150008291070026000",
				Amount:       5000,
				DueDate:      "2024-12-31",
				Status:       "PENDING",
			},
			{
				SchedulingID: 2,
				Digitable:    "34191790010104351004791020150008291070027000",
				Amount:       3000,
				DueDate:      "2025-01-15",
				Status:       "PENDING",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/bill/account/12345/schedules"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListScheduledBills(context.Background(), accountID)
	if err != nil {
		t.Fatalf("ListScheduledBills failed: %v", err)
	}

	if len(response.Bills) != 2 {
		t.Errorf("bills count = %d; want 2", len(response.Bills))
	}
	if response.Bills[0].SchedulingID != 1 {
		t.Errorf("first bill schedulingID = %d; want 1", response.Bills[0].SchedulingID)
	}
}

// TestPayBillByAccount tests paying a bill via account endpoint
func TestPayBillByAccount(t *testing.T) {
	accountID := int64(12345)
	requestData := &types.BillPaymentRequest{
		Barcode: "34191790010104351004791020150008291070026000",
		Amount:  intPtr(5000),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/billpayment"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.BillPaymentResponse{Message: "Payment successful"})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.PayBillByAccount(context.Background(), accountID, requestData)
	if err != nil {
		t.Fatalf("PayBillByAccount failed: %v", err)
	}

	if response.Message != "Payment successful" {
		t.Errorf("message = %q; want %q", response.Message, "Payment successful")
	}
}

// TestGetBillInfoByAccount tests getting bill info via account endpoint
func TestGetBillInfoByAccount(t *testing.T) {
	accountID := int64(12345)
	requestData := &types.GetBillInfoRequest{
		Barcode: "34191790010104351004791020150008291070026000",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/billpayment"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.GetBillInfoResponse{
			OriginalValue: 5000,
			DueDate:       "2024-12-31",
		})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetBillInfoByAccount(context.Background(), accountID, requestData)
	if err != nil {
		t.Fatalf("GetBillInfoByAccount failed: %v", err)
	}

	if response.OriginalValue != 5000 {
		t.Errorf("originalValue = %d; want 5000", response.OriginalValue)
	}
}

// TestCancelScheduledBillByAccount tests canceling scheduled bill via account endpoint
func TestCancelScheduledBillByAccount(t *testing.T) {
	accountID := int64(12345)
	schedulingID := int64(67890)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/billpayment/scheduled/67890"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.CancelBillResponse{Message: "Canceled"})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CancelScheduledBillByAccount(context.Background(), accountID, schedulingID)
	if err != nil {
		t.Fatalf("CancelScheduledBillByAccount failed: %v", err)
	}

	if response.Message != "Canceled" {
		t.Errorf("message = %q; want %q", response.Message, "Canceled")
	}
}

// TestPayBillBatchByAccount tests batch bill payment via account endpoint
func TestPayBillBatchByAccount(t *testing.T) {
	accountID := int64(12345)
	requestData := &types.BatchBillPaymentRequest{
		Payments: []types.BillPaymentRequest{
			{
				Barcode: "34191790010104351004791020150008291070026000",
				Amount:  intPtr(5000),
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/billpayment/batch"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		var body types.BatchBillPaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if len(body.Payments) != 1 {
			t.Errorf("payments count = %d; want 1", len(body.Payments))
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.BillPaymentResponse{Message: "Batch completed"})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.PayBillBatchByAccount(context.Background(), accountID, requestData)
	if err != nil {
		t.Fatalf("PayBillBatchByAccount failed: %v", err)
	}

	if response.Message != "Batch completed" {
		t.Errorf("message = %q; want %q", response.Message, "Batch completed")
	}
}

// TestListScheduledBillsByAccount tests listing scheduled bills via account endpoint
func TestListScheduledBillsByAccount(t *testing.T) {
	accountID := int64(12345)

	totalCount := 1
	responseData := &types.ScheduledBillPaymentListResponse{
		Payments: []types.ScheduledBillPaymentResponse{
			{
				SchedulingID:  1,
				Barcode:       "34191790010104351004791020150008291070026000",
				Amount:        5000,
				ScheduledDate: "2024-12-31",
				Status:        "PENDING",
			},
		},
		Total: &totalCount,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/billpayment/scheduled"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListScheduledBillsByAccount(context.Background(), accountID)
	if err != nil {
		t.Fatalf("ListScheduledBillsByAccount failed: %v", err)
	}

	if len(response.Payments) != 1 {
		t.Errorf("payments count = %d; want 1", len(response.Payments))
	}
	if response.Payments[0].SchedulingID != 1 {
		t.Errorf("schedulingID = %d; want 1", response.Payments[0].SchedulingID)
	}
}

// ===== Deposit Order Tests =====

// TestListDepositOrders tests listing deposit orders
func TestListDepositOrders(t *testing.T) {
	accountID := int64(12345)
	params := &types.ListDepositOrdersParams{
		StartDate: strPtr("2024-01-01"),
		EndDate:   strPtr("2024-12-31"),
	}

	responseData := &types.DepositOrdersResponse{
		Deposits: []types.DepositOrder{
			{
				ID:       1,
				Amount:   10000,
				Covenant: "TEST_COVENANT",
				Status:   "ACTIVE",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/deposits/order"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Verify query parameters
		if r.URL.Query().Get("startDate") != "2024-01-01" {
			t.Errorf("startDate = %q; want %q", r.URL.Query().Get("startDate"), "2024-01-01")
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListDepositOrders(context.Background(), accountID, params)
	if err != nil {
		t.Fatalf("ListDepositOrders failed: %v", err)
	}

	if len(response.Deposits) != 1 {
		t.Errorf("deposits count = %d; want 1", len(response.Deposits))
	}
	if response.Deposits[0].ID != 1 {
		t.Errorf("deposit ID = %d; want 1", response.Deposits[0].ID)
	}
}

// TestListDepositOrdersWithoutParams tests listing deposit orders without params
func TestListDepositOrdersWithoutParams(t *testing.T) {
	accountID := int64(12345)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/deposits/order"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		// Query string should be empty
		if r.URL.RawQuery != "" {
			t.Errorf("expected empty query string, got %q", r.URL.RawQuery)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.DepositOrdersResponse{Deposits: []types.DepositOrder{}})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	_, err = client.ListDepositOrders(context.Background(), accountID, nil)
	if err != nil {
		t.Fatalf("ListDepositOrders failed: %v", err)
	}
}

// TestCreateDepositOrder tests creating a deposit order
func TestCreateDepositOrder(t *testing.T) {
	accountID := int64(12345)
	expiryHours := 24
	requestData := &types.CreateDepositOrderRequest{
		Amount:      10000,
		Description: strPtr("Test deposit"),
		ExpiryHours: &expiryHours,
	}

	responseData := &types.CreateDepositOrderResponse{
		DepositOrderID:     67890,
		DateTimeExpiration: timePtr(time.Now().Add(24 * time.Hour)),
		Status:             "PENDING",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/deposits/order"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		var body types.CreateDepositOrderRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.Amount != requestData.Amount {
			t.Errorf("amount = %d; want %d", body.Amount, requestData.Amount)
		}
		if *body.Description != *requestData.Description {
			t.Errorf("description = %q; want %q", *body.Description, *requestData.Description)
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CreateDepositOrder(context.Background(), accountID, requestData)
	if err != nil {
		t.Fatalf("CreateDepositOrder failed: %v", err)
	}

	if response.DepositOrderID != responseData.DepositOrderID {
		t.Errorf("depositOrderID = %d; want %d", response.DepositOrderID, responseData.DepositOrderID)
	}
	if response.Status != responseData.Status {
		t.Errorf("status = %q; want %q", response.Status, responseData.Status)
	}
}

// TestListActiveDepositOrders tests listing active deposit orders
func TestListActiveDepositOrders(t *testing.T) {
	accountID := int64(12345)

	responseData := &types.DepositOrdersResponse{
		Deposits: []types.DepositOrder{
			{
				ID:       1,
				Amount:   10000,
				Covenant: "TEST",
				Status:   "ACTIVE",
			},
			{
				ID:       2,
				Amount:   20000,
				Covenant: "TEST2",
				Status:   "ACTIVE",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/deposits/order/active"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListActiveDepositOrders(context.Background(), accountID)
	if err != nil {
		t.Fatalf("ListActiveDepositOrders failed: %v", err)
	}

	if len(response.Deposits) != 2 {
		t.Errorf("deposits count = %d; want 2", len(response.Deposits))
	}
	for _, deposit := range response.Deposits {
		if deposit.Status != "ACTIVE" {
			t.Errorf("deposit status = %q; want ACTIVE", deposit.Status)
		}
	}
}

// TestCancelDepositOrder tests canceling a deposit order
func TestCancelDepositOrder(t *testing.T) {
	accountID := int64(12345)
	depositOrderID := int64(67890)

	responseData := &types.GenericResponse{
		Message: "Deposit order canceled successfully",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/deposits/order/67890"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CancelDepositOrder(context.Background(), accountID, depositOrderID)
	if err != nil {
		t.Fatalf("CancelDepositOrder failed: %v", err)
	}

	if response.Message != responseData.Message {
		t.Errorf("message = %q; want %q", response.Message, responseData.Message)
	}
}

// ===== Bankslip Tests =====

// TestListBankslipsForAccount tests listing bankslips for an account
func TestListBankslipsForAccount(t *testing.T) {
	accountID := int64(12345)

	responseData := &types.BankslipsResponse{
		Bankslips: []types.BankslipItem{
			{
				ID:            1,
				DigitableLine: "34191790010104351004791020150008291070026000",
				Barcode:       "34191.79001 01043.510047 91020.150008 2 91070026000",
				Amount:        5000,
				DueDate:       "2024-12-31",
				Status:        "PENDING",
				DownloadURL:   strPtr("https://example.com/bankslip/1.pdf"),
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/bankslip"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListBankslips(context.Background(), accountID)
	if err != nil {
		t.Fatalf("ListBankslips failed: %v", err)
	}

	if len(response.Bankslips) != 1 {
		t.Errorf("bankslips count = %d; want 1", len(response.Bankslips))
	}
	if response.Bankslips[0].ID != 1 {
		t.Errorf("bankslip ID = %d; want 1", response.Bankslips[0].ID)
	}
}

// TestCreateBankslipForAccount tests creating a bankslip for an account
func TestCreateBankslipForAccount(t *testing.T) {
	accountID := int64(12345)
	requestData := &types.CreateBankslipRequest{
		Amount:      5000,
		DueDate:     "2024-12-31",
		Description: strPtr("Payment for services"),
	}

	responseData := &types.CreateBankslipResponse{
		IDBankslip:          67890,
		DigitableLine:       "34191790010104351004791020150008291070026000",
		DueDate:             "2024-12-31",
		BankslipDownloadURL: strPtr("https://example.com/bankslip/67890.pdf"),
		Barcode:             strPtr("34191.79001 01043.510047 91020.150008 2 91070026000"),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/bankslip"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		var body types.CreateBankslipRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.Amount != requestData.Amount {
			t.Errorf("amount = %d; want %d", body.Amount, requestData.Amount)
		}
		if body.DueDate != requestData.DueDate {
			t.Errorf("dueDate = %q; want %q", body.DueDate, requestData.DueDate)
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CreateBankslip(context.Background(), accountID, requestData)
	if err != nil {
		t.Fatalf("CreateBankslip failed: %v", err)
	}

	if response.IDBankslip != responseData.IDBankslip {
		t.Errorf("idBankslip = %d; want %d", response.IDBankslip, responseData.IDBankslip)
	}
	if response.DigitableLine != responseData.DigitableLine {
		t.Errorf("digitableLine = %q; want %q", response.DigitableLine, responseData.DigitableLine)
	}
}

// TestListBankslipsByStatus tests listing bankslips by status
func TestListBankslipsByStatus(t *testing.T) {
	accountID := int64(12345)
	status := "PENDING"

	responseData := &types.BankslipsResponse{
		Bankslips: []types.BankslipItem{
			{
				ID:            1,
				DigitableLine: "34191790010104351004791020150008291070026000",
				Barcode:       "34191.79001 01043.510047 91020.150008 2 91070026000",
				Amount:        5000,
				DueDate:       "2024-12-31",
				Status:        types.BankSlipStatusPending,
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/bankslip/PENDING"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListBankslipsByStatus(context.Background(), accountID, status)
	if err != nil {
		t.Fatalf("ListBankslipsByStatus failed: %v", err)
	}

	if len(response.Bankslips) != 1 {
		t.Errorf("bankslips count = %d; want 1", len(response.Bankslips))
	}
	if response.Bankslips[0].Status != types.BankSlipStatusPending {
		t.Errorf("bankslip status = %q; want %q", response.Bankslips[0].Status, types.BankSlipStatusPending)
	}
}

// TestListBankslipsByStatusAndDate tests listing bankslips by status and date
func TestListBankslipsByStatusAndDate(t *testing.T) {
	accountID := int64(12345)
	status := "PENDING"
	createdAt := "2024-12-01"

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/bankslip/PENDING/2024-12-01"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&types.BankslipsResponse{Bankslips: []types.BankslipItem{}})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	_, err = client.ListBankslipsByStatusAndDate(context.Background(), accountID, status, createdAt)
	if err != nil {
		t.Fatalf("ListBankslipsByStatusAndDate failed: %v", err)
	}
}

// TestCreateBankslipV2 tests creating a bankslip using v2 API
func TestCreateBankslipV2(t *testing.T) {
	requestData := &types.BankslipV2Request{
		AccountID:     12345,
		Amount:        5000,
		DueDate:       "2024-12-31",
		PayerDocument: "12345678901",
		PayerName:     "John Doe",
		Description:   strPtr("Payment for services"),
		Instructions:  strPtr("Pay before due date"),
	}

	responseData := &types.BankslipV2Response{
		BankslipID:    "BS123456",
		Barcode:       "34191.79001 01043.510047 91020.150008 2 91070026000",
		DigitableLine: "34191790010104351004791020150008291070026000",
		PDFBase64:     "JVBERi0xLjQKJeLjz9MK...",
		Amount:        5000,
		DueDate:       "2024-12-31",
		Status:        "PENDING",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/bankslip/v2/generate" {
			t.Errorf("expected path /bankslip/v2/generate, got %s", r.URL.Path)
		}

		var body types.BankslipV2Request
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.AccountID != requestData.AccountID {
			t.Errorf("accountID = %d; want %d", body.AccountID, requestData.AccountID)
		}
		if body.PayerDocument != requestData.PayerDocument {
			t.Errorf("payerDocument = %q; want %q", body.PayerDocument, requestData.PayerDocument)
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CreateBankslipV2(context.Background(), requestData)
	if err != nil {
		t.Fatalf("CreateBankslipV2 failed: %v", err)
	}

	if response.BankslipID != responseData.BankslipID {
		t.Errorf("bankslipID = %q; want %q", response.BankslipID, responseData.BankslipID)
	}
	if response.Status != responseData.Status {
		t.Errorf("status = %q; want %q", response.Status, responseData.Status)
	}
	if response.Amount != responseData.Amount {
		t.Errorf("amount = %d; want %d", response.Amount, responseData.Amount)
	}
}
