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

// TestCreatePixClaim tests creating a PIX claim
func TestCreatePixClaim(t *testing.T) {
	accountID := int64(12345)
	reqData := &types.CreatePixClaimRequest{
		KeyType:   types.PixKeyTypeCPF,
		KeyValue:  "12345678901",
		ClaimType: "PORTABILITY",
	}
	respData := &types.PixClaimResponse{
		ClaimID:   "claim123",
		KeyType:   types.PixKeyTypeCPF,
		KeyValue:  "12345678901",
		ClaimType: "PORTABILITY",
		Status:    "PENDING",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/createClaim"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		var body types.CreatePixClaimRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
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

	response, err := client.CreatePixClaim(context.Background(), accountID, reqData)
	if err != nil {
		t.Fatalf("CreatePixClaim failed: %v", err)
	}

	if response.ClaimID != respData.ClaimID {
		t.Errorf("response ClaimID = %q; want %q", response.ClaimID, respData.ClaimID)
	}
	if response.Status != respData.Status {
		t.Errorf("response Status = %q; want %q", response.Status, respData.Status)
	}
}

// TestConfirmPortability tests confirming PIX portability
func TestConfirmPortability(t *testing.T) {
	accountID := int64(12345)
	reqData := &types.ConfirmPortabilityRequest{
		ClaimID: "claim123",
	}
	respData := &types.PixClaimResponse{
		ClaimID: "claim123",
		Status:  "CONFIRMED",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/confirmPortability"
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

	response, err := client.ConfirmPortability(context.Background(), accountID, reqData)
	if err != nil {
		t.Fatalf("ConfirmPortability failed: %v", err)
	}

	if response.ClaimID != respData.ClaimID {
		t.Errorf("response ClaimID = %q; want %q", response.ClaimID, respData.ClaimID)
	}
	if response.Status != respData.Status {
		t.Errorf("response Status = %q; want %q", response.Status, respData.Status)
	}
}

// TestCompletePortability tests completing PIX portability
func TestCompletePortability(t *testing.T) {
	accountID := int64(12345)
	reqData := &types.CompletePortabilityRequest{
		ClaimID: "claim123",
	}
	respData := &types.PixClaimResponse{
		ClaimID: "claim123",
		Status:  "COMPLETED",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/completePortability"
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

	response, err := client.CompletePortability(context.Background(), accountID, reqData)
	if err != nil {
		t.Fatalf("CompletePortability failed: %v", err)
	}

	if response.Status != respData.Status {
		t.Errorf("response Status = %q; want %q", response.Status, respData.Status)
	}
}

// TestCancelPortability tests canceling PIX portability
func TestCancelPortability(t *testing.T) {
	accountID := int64(12345)
	reason := "Changed my mind"
	reqData := &types.CancelPortabilityRequest{
		ClaimID: "claim123",
		Reason:  &reason,
	}
	respData := &types.PixClaimResponse{
		ClaimID: "claim123",
		Status:  "CANCELED",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/cancelPortability"
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

	response, err := client.CancelPortability(context.Background(), accountID, reqData)
	if err != nil {
		t.Fatalf("CancelPortability failed: %v", err)
	}

	if response.Status != respData.Status {
		t.Errorf("response Status = %q; want %q", response.Status, respData.Status)
	}
}

// TestGetRequestedClaims tests retrieving PIX claims
func TestGetRequestedClaims(t *testing.T) {
	accountID := int64(12345)
	respData := &types.PixClaimListResponse{
		Claims: []types.PixClaimResponse{
			{
				ClaimID:   "claim123",
				KeyType:   types.PixKeyTypeCPF,
				KeyValue:  "12345678901",
				ClaimType: "PORTABILITY",
				Status:    "PENDING",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/getRequestedClaims"
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

	response, err := client.GetRequestedClaims(context.Background(), accountID)
	if err != nil {
		t.Fatalf("GetRequestedClaims failed: %v", err)
	}

	if len(response.Claims) != 1 {
		t.Errorf("expected 1 claim, got %d", len(response.Claims))
	}
	if response.Claims[0].ClaimID != "claim123" {
		t.Errorf("response Claims[0].ClaimID = %q; want %q", response.Claims[0].ClaimID, "claim123")
	}
}

// TestGetPixLimit tests retrieving PIX limits
func TestGetPixLimit(t *testing.T) {
	accountID := int64(12345)
	nightTimeStart := "20:00"
	nightTimeEnd := "06:00"
	respData := &types.PixLimitResponse{
		AccountID:        accountID,
		DailyLimit:       100000,
		NightlyLimit:     50000,
		TransactionLimit: 10000,
		DailyUsed:        10000,
		NightlyUsed:      0,
		DailyRemaining:   90000,
		NightlyRemaining: 50000,
		NightTimeStart:   &nightTimeStart,
		NightTimeEnd:     &nightTimeEnd,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/pix/getLimit"
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

	response, err := client.GetPixLimit(context.Background(), accountID)
	if err != nil {
		t.Fatalf("GetPixLimit failed: %v", err)
	}

	if response.AccountID != respData.AccountID {
		t.Errorf("response AccountID = %d; want %d", response.AccountID, respData.AccountID)
	}
	if response.DailyLimit != respData.DailyLimit {
		t.Errorf("response DailyLimit = %d; want %d", response.DailyLimit, respData.DailyLimit)
	}
	if response.DailyRemaining != respData.DailyRemaining {
		t.Errorf("response DailyRemaining = %d; want %d", response.DailyRemaining, respData.DailyRemaining)
	}
}

// TestUpdatePixLimit tests updating PIX limits
func TestUpdatePixLimit(t *testing.T) {
	accountID := int64(12345)
	dailyLimit := int64(200000)
	reqData := &types.PixLimitRequest{
		DailyLimit: &dailyLimit,
	}
	respData := &types.PixLimitResponse{
		AccountID:  accountID,
		DailyLimit: 200000,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/pix/limit"
		if r.URL.Path != expectedPath {
			t.Errorf("expected path %s, got %s", expectedPath, r.URL.Path)
		}

		var body types.PixLimitRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.DailyLimit == nil || *body.DailyLimit != dailyLimit {
			t.Errorf("request DailyLimit = %v; want %d", body.DailyLimit, dailyLimit)
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

	response, err := client.UpdatePixLimit(context.Background(), accountID, reqData)
	if err != nil {
		t.Fatalf("UpdatePixLimit failed: %v", err)
	}

	if response.DailyLimit != respData.DailyLimit {
		t.Errorf("response DailyLimit = %d; want %d", response.DailyLimit, respData.DailyLimit)
	}
}

// TestUpdatePixNightTimeLimit tests updating PIX night-time limits
func TestUpdatePixNightTimeLimit(t *testing.T) {
	accountID := int64(12345)
	reqData := &types.UpdatePixNightTimeLimitRequest{
		StartTime: "20:00",
		EndTime:   "06:00",
	}
	nightTimeStart := "20:00"
	nightTimeEnd := "06:00"
	respData := &types.PixLimitResponse{
		AccountID:      accountID,
		NightTimeStart: &nightTimeStart,
		NightTimeEnd:   &nightTimeEnd,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		expectedPath := "/accounts/12345/pix/limit/startNightTime"
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

	response, err := client.UpdatePixNightTimeLimit(context.Background(), accountID, reqData)
	if err != nil {
		t.Fatalf("UpdatePixNightTimeLimit failed: %v", err)
	}

	if *response.NightTimeStart != "20:00" {
		t.Errorf("response NightTimeStart = %q; want %q", *response.NightTimeStart, "20:00")
	}
}

// TestAddPixDevice tests adding a PIX device
func TestAddPixDevice(t *testing.T) {
	deviceType := "MOBILE"
	reqData := &types.PixDeviceRequest{
		DeviceID:   "device123",
		DeviceName: "My Phone",
		DeviceType: &deviceType,
	}
	respData := &types.PixDeviceResponse{
		DeviceID:   "device123",
		DeviceName: "My Phone",
		Status:     "ACTIVE",
		AccountID:  12345,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/devices" {
			t.Errorf("expected path /pix/devices, got %s", r.URL.Path)
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

	response, err := client.AddPixDevice(context.Background(), reqData)
	if err != nil {
		t.Fatalf("AddPixDevice failed: %v", err)
	}

	if response.DeviceID != respData.DeviceID {
		t.Errorf("response DeviceID = %q; want %q", response.DeviceID, respData.DeviceID)
	}
	if response.Status != respData.Status {
		t.Errorf("response Status = %q; want %q", response.Status, respData.Status)
	}
}

// TestDeletePixDevice tests deleting a PIX device
func TestDeletePixDevice(t *testing.T) {
	reqData := &types.DeletePixDeviceRequest{
		DeviceID: "device123",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/devices" {
			t.Errorf("expected path /pix/devices, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	err = client.DeletePixDevice(context.Background(), reqData)
	if err != nil {
		t.Fatalf("DeletePixDevice failed: %v", err)
	}
}

// TestBlockPixDevice tests blocking a PIX device
func TestBlockPixDevice(t *testing.T) {
	reason := "Security concern"
	reqData := &types.BlockPixDeviceRequest{
		DeviceID: "device123",
		Reason:   &reason,
	}
	respData := &types.PixDeviceResponse{
		DeviceID: "device123",
		Status:   "BLOCKED",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/devices/block" {
			t.Errorf("expected path /pix/devices/block, got %s", r.URL.Path)
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

	response, err := client.BlockPixDevice(context.Background(), reqData)
	if err != nil {
		t.Fatalf("BlockPixDevice failed: %v", err)
	}

	if response.Status != "BLOCKED" {
		t.Errorf("response Status = %q; want %q", response.Status, "BLOCKED")
	}
}

// TestUnblockPixDevice tests unblocking a PIX device
func TestUnblockPixDevice(t *testing.T) {
	reqData := &types.UnblockPixDeviceRequest{
		DeviceID: "device123",
	}
	respData := &types.PixDeviceResponse{
		DeviceID: "device123",
		Status:   "ACTIVE",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/devices/unblock" {
			t.Errorf("expected path /pix/devices/unblock, got %s", r.URL.Path)
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

	response, err := client.UnblockPixDevice(context.Background(), reqData)
	if err != nil {
		t.Fatalf("UnblockPixDevice failed: %v", err)
	}

	if response.Status != "ACTIVE" {
		t.Errorf("response Status = %q; want %q", response.Status, "ACTIVE")
	}
}

// TestListPixDevices tests listing PIX devices
func TestListPixDevices(t *testing.T) {
	accountID := int64(12345)
	respData := &types.PixDeviceListResponse{
		Devices: []types.PixDeviceResponse{
			{
				DeviceID:   "device123",
				DeviceName: "My Phone",
				Status:     "ACTIVE",
				AccountID:  accountID,
			},
			{
				DeviceID:   "device456",
				DeviceName: "My Tablet",
				Status:     "ACTIVE",
				AccountID:  accountID,
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/pix/devices/list/12345"
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

	response, err := client.ListPixDevices(context.Background(), accountID)
	if err != nil {
		t.Fatalf("ListPixDevices failed: %v", err)
	}

	if len(response.Devices) != 2 {
		t.Errorf("expected 2 devices, got %d", len(response.Devices))
	}
}

// TestListPixClaims tests listing PIX claims with filters
func TestListPixClaims(t *testing.T) {
	status := "PENDING"
	reqData := &types.ListPixClaimsRequest{
		Status: &status,
	}
	respData := &types.PixClaimListResponse{
		Claims: []types.PixClaimResponse{
			{
				ClaimID:   "claim123",
				KeyType:   types.PixKeyTypeCPF,
				KeyValue:  "12345678901",
				ClaimType: "PORTABILITY",
				Status:    "PENDING",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/pix/claim/list" {
			t.Errorf("expected path /accounts/pix/claim/list, got %s", r.URL.Path)
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

	response, err := client.ListPixClaims(context.Background(), reqData)
	if err != nil {
		t.Fatalf("ListPixClaims failed: %v", err)
	}

	if len(response.Claims) != 1 {
		t.Errorf("expected 1 claim, got %d", len(response.Claims))
	}
}

// TestCreateClaimFromKey tests creating a claim from existing key
func TestCreateClaimFromKey(t *testing.T) {
	reqData := &types.CreateClaimFromKeyRequest{
		KeyType:   types.PixKeyTypeCPF,
		KeyValue:  "12345678901",
		ClaimType: "OWNERSHIP",
	}
	respData := &types.PixClaimResponse{
		ClaimID:   "claim456",
		KeyType:   types.PixKeyTypeCPF,
		KeyValue:  "12345678901",
		ClaimType: "OWNERSHIP",
		Status:    "PENDING",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/pix/claim/key/createClaim" {
			t.Errorf("expected path /accounts/pix/claim/key/createClaim, got %s", r.URL.Path)
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

	response, err := client.CreateClaimFromKey(context.Background(), reqData)
	if err != nil {
		t.Fatalf("CreateClaimFromKey failed: %v", err)
	}

	if response.ClaimID != respData.ClaimID {
		t.Errorf("response ClaimID = %q; want %q", response.ClaimID, respData.ClaimID)
	}
}

// TestProcessLimitRequest tests processing a PIX limit request
func TestProcessLimitRequest(t *testing.T) {
	reqData := &types.ProcessLimitRequestData{
		RequestID: 123,
		Action:    "APPROVE",
	}
	respData := &types.GenericResponse{
		Message: "Request processed successfully",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/pix/limit/processLimitRequest" {
			t.Errorf("expected path /accounts/pix/limit/processLimitRequest, got %s", r.URL.Path)
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

	response, err := client.ProcessLimitRequest(context.Background(), reqData)
	if err != nil {
		t.Fatalf("ProcessLimitRequest failed: %v", err)
	}

	if response.Message != respData.Message {
		t.Errorf("response Message = %q; want %q", response.Message, respData.Message)
	}
}

// TestGetRaiseLimitRequests tests retrieving raise limit requests
func TestGetRaiseLimitRequests(t *testing.T) {
	respData := &types.RaiseLimitRequestListResponse{
		Requests: []types.RaiseLimitRequestResponse{
			{
				RequestID:      123,
				AccountID:      12345,
				RequestedLimit: 500000,
				CurrentLimit:   200000,
				Status:         "PENDING",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/pix/limit/getRaiseLimitRequests" {
			t.Errorf("expected path /accounts/pix/limit/getRaiseLimitRequests, got %s", r.URL.Path)
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

	response, err := client.GetRaiseLimitRequests(context.Background())
	if err != nil {
		t.Fatalf("GetRaiseLimitRequests failed: %v", err)
	}

	if len(response.Requests) != 1 {
		t.Errorf("expected 1 request, got %d", len(response.Requests))
	}
}

// TestGetMaximumPixLimitIssuer tests retrieving maximum PIX limit
func TestGetMaximumPixLimitIssuer(t *testing.T) {
	respData := &types.MaximumPixLimitIssuerResponse{
		MaxDailyLimit:       1000000,
		MaxNightlyLimit:     500000,
		MaxTransactionLimit: 50000,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/pix/limit/getMaximumLimitIssuer" {
			t.Errorf("expected path /accounts/pix/limit/getMaximumLimitIssuer, got %s", r.URL.Path)
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

	response, err := client.GetMaximumPixLimitIssuer(context.Background())
	if err != nil {
		t.Fatalf("GetMaximumPixLimitIssuer failed: %v", err)
	}

	if response.MaxDailyLimit != respData.MaxDailyLimit {
		t.Errorf("response MaxDailyLimit = %d; want %d", response.MaxDailyLimit, respData.MaxDailyLimit)
	}
}

// TestGetRaiseLimitRequestDetail tests retrieving raise limit request detail
func TestGetRaiseLimitRequestDetail(t *testing.T) {
	requestID := int64(123)
	respData := &types.RaiseLimitRequestResponse{
		RequestID:      123,
		AccountID:      12345,
		RequestedLimit: 500000,
		CurrentLimit:   200000,
		Status:         "PENDING",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		expectedPath := "/accounts/pix/limit/getDetailRaiseLimitRequest/123"
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

	response, err := client.GetRaiseLimitRequestDetail(context.Background(), requestID)
	if err != nil {
		t.Fatalf("GetRaiseLimitRequestDetail failed: %v", err)
	}

	if response.RequestID != respData.RequestID {
		t.Errorf("response RequestID = %d; want %d", response.RequestID, respData.RequestID)
	}
}

// TestReceivePixCallback tests receiving PIX callback
func TestReceivePixCallback(t *testing.T) {
	reqData := &types.PixCallbackRequest{
		EndToEndID:      "E123456789202401010000000001",
		Amount:          10000,
		TransactionDate: "2024-01-01T12:00:00Z",
	}
	message := "Callback received"
	respData := &types.PixCallbackResponse{
		Processed: true,
		Message:   &message,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/pix/callbacks/receive-transaction" {
			t.Errorf("expected path /pix/callbacks/receive-transaction, got %s", r.URL.Path)
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

	response, err := client.ReceivePixCallback(context.Background(), reqData)
	if err != nil {
		t.Fatalf("ReceivePixCallback failed: %v", err)
	}

	if !response.Processed {
		t.Errorf("response Processed = %v; want %v", response.Processed, true)
	}
}
