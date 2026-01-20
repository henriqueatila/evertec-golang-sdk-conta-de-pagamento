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

// ========== ACCOUNTS API TESTS ==========

func TestGetAccount(t *testing.T) {
	responseData := types.AccountDataResponse{
		AccountID:    12345,
		Name:         "John Doe",
		Email:        "john@example.com",
		Document:     "12345678901",
		DocumentType: types.DocumentTypeCPF,
		Phone:        "+5511999999999",
		Status:       types.AccountStatusActive,
		AccountType:  types.AccountTypePersonal,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345" {
			t.Errorf("expected path /accounts/12345, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetAccount(context.Background(), 12345)
	if err != nil {
		t.Fatalf("GetAccount failed: %v", err)
	}

	if response.AccountID != responseData.AccountID {
		t.Errorf("AccountID = %d; want %d", response.AccountID, responseData.AccountID)
	}
	if response.Name != responseData.Name {
		t.Errorf("Name = %s; want %s", response.Name, responseData.Name)
	}
	if response.Email != responseData.Email {
		t.Errorf("Email = %s; want %s", response.Email, responseData.Email)
	}
}

func TestGetAccount_NotFound(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message":  "account not found",
			"resource": "account",
		})
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	_, err = client.GetAccount(context.Background(), 99999)
	if err == nil {
		t.Fatal("expected error for non-existent account, got nil")
	}
}

func TestListAccounts(t *testing.T) {
	responseData := types.AccountListResponse{
		Accounts: []types.AccountDataResponse{
			{
				AccountID: 1,
				Name:      "User 1",
				Email:     "user1@example.com",
			},
			{
				AccountID: 2,
				Name:      "User 2",
				Email:     "user2@example.com",
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts" {
			t.Errorf("expected path /accounts, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListAccounts(context.Background(), nil)
	if err != nil {
		t.Fatalf("ListAccounts failed: %v", err)
	}

	if len(response.Accounts) != 2 {
		t.Errorf("expected 2 accounts, got %d", len(response.Accounts))
	}
	if response.Accounts[0].AccountID != 1 {
		t.Errorf("first account ID = %d; want 1", response.Accounts[0].AccountID)
	}
}

func TestCreateAccount(t *testing.T) {
	requestData := types.ProposalAccountRequest{
		PersonalName:           "New User",
		PersonalEmail:          "newuser@example.com",
		PersonalDocument:       "12345678901",
		BirthDate:              "1990-01-01",
		MotherName:             "Test Mother",
		ProductID:              1,
		LocalAreaCodeCellPhone: 11,
		CellPhoneNumber:        999999999,
		BranchID:               1,
	}

	responseData := types.CreateAccountResponse{
		AccountID:     54321,
		PersonID:      11111,
		PersonAddressID: 22222,
		IDProposal:    33333,
		AccountNumber: 123456789,
		Branch:        "0001",
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts" {
			t.Errorf("expected path /accounts, got %s", r.URL.Path)
		}

		var body types.ProposalAccountRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.PersonalName != requestData.PersonalName {
			t.Errorf("request name = %s; want %s", body.PersonalName, requestData.PersonalName)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CreateAccount(context.Background(), &requestData)
	if err != nil {
		t.Fatalf("CreateAccount failed: %v", err)
	}

	if response.AccountID != responseData.AccountID {
		t.Errorf("AccountID = %d; want %d", response.AccountID, responseData.AccountID)
	}
	if response.Branch != responseData.Branch {
		t.Errorf("Branch = %s; want %s", response.Branch, responseData.Branch)
	}
}

func TestGetAccountBalance(t *testing.T) {
	now := time.Now()
	responseData := types.BalanceResponse{
		Message:   "Success",
		IDAccount: 12345,
		Balance:   120000, // $1200.00
		DateTime:  now,
		DACode:    100,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345/balance" {
			t.Errorf("expected path /accounts/12345/balance, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetAccountBalance(context.Background(), 12345)
	if err != nil {
		t.Fatalf("GetAccountBalance failed: %v", err)
	}

	if response.IDAccount != responseData.IDAccount {
		t.Errorf("IDAccount = %d; want %d", response.IDAccount, responseData.IDAccount)
	}
	if response.Balance != responseData.Balance {
		t.Errorf("Balance = %d; want %d", response.Balance, responseData.Balance)
	}
}

func TestGetAccountStatement(t *testing.T) {
	now := time.Now()
	responseData := types.StatementResponse{
		Message:   "Success",
		Total:     1,
		AccountID: 12345,
		Entries: []types.StatementEntry{
			{
				TransactionID:          "TXN-001",
				TransactionDate:        now,
				TransactionDescription: "Payment received",
				Amount:                 &types.AmountDTO{Amount: 50000, CurrencyCode: 986},
				DebitOrCredit:          "C",
			},
		},
		Count:  1,
		DACode: 100,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345/statement" {
			t.Errorf("expected path /accounts/12345/statement, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetAccountStatement(context.Background(), 12345, nil)
	if err != nil {
		t.Fatalf("GetAccountStatement failed: %v", err)
	}

	if response.AccountID != responseData.AccountID {
		t.Errorf("AccountID = %d; want %d", response.AccountID, responseData.AccountID)
	}
	if len(response.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(response.Entries))
	}
	if response.Entries[0].TransactionID != "TXN-001" {
		t.Errorf("TransactionID = %s; want TXN-001", response.Entries[0].TransactionID)
	}
}

// ========== CARDS API TESTS ==========

func TestGetCard(t *testing.T) {
	now := time.Now()
	responseData := types.CardResponse{
		ID:               67890,
		PersonID:         11111,
		AccountID:        12345,
		CardType:         "PHYSICAL",
		FunctionType:     "DEBIT",
		Modality:         "STANDARD",
		ExpirationDate:   now,
		Status:           1, // Active
		LastFour:         ptrString("5678"),
		PrintedName:      ptrString("JOHN DOE"),
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345/cards/67890" {
			t.Errorf("expected path /accounts/12345/cards/67890, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.GetCard(context.Background(), 12345, 67890)
	if err != nil {
		t.Fatalf("GetCard failed: %v", err)
	}

	if response.ID != responseData.ID {
		t.Errorf("ID = %d; want %d", response.ID, responseData.ID)
	}
	if response.AccountID != responseData.AccountID {
		t.Errorf("AccountID = %d; want %d", response.AccountID, responseData.AccountID)
	}
	if response.CardType != responseData.CardType {
		t.Errorf("CardType = %v; want %v", response.CardType, responseData.CardType)
	}
}

func TestListCards(t *testing.T) {
	now := time.Now()
	responseData := types.AccountCardsResponse{
		Message:   "Success",
		AccountID: 12345,
		DACode:    100,
		Cards: []types.CardResponse{
			{
				ID:             1,
				PersonID:       11111,
				AccountID:      12345,
				CardType:       "PHYSICAL",
				FunctionType:   "DEBIT",
				Modality:       "STANDARD",
				ExpirationDate: now,
				Status:         1, // Active
			},
			{
				ID:             2,
				PersonID:       11111,
				AccountID:      12345,
				CardType:       "PHYSICAL",
				FunctionType:   "DEBIT",
				Modality:       "STANDARD",
				ExpirationDate: now,
				Status:         2, // Blocked
			},
		},
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345/cards" {
			t.Errorf("expected path /accounts/12345/cards, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.ListCards(context.Background(), 12345, nil)
	if err != nil {
		t.Fatalf("ListCards failed: %v", err)
	}

	if len(response.Cards) != 2 {
		t.Errorf("expected 2 cards, got %d", len(response.Cards))
	}
	if response.Cards[0].ID != 1 {
		t.Errorf("first card ID = %d; want 1", response.Cards[0].ID)
	}
}

func TestBlockCard(t *testing.T) {
	requestData := types.BlockCardRequest{
		DestinationStatus: 1,
		BlockMotivation:   "Lost card",
	}

	responseData := types.BlockCardResponse{
		Message: "Card blocked successfully",
		DACode:  100,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345/cards/67890/block" {
			t.Errorf("expected path /accounts/12345/cards/67890/block, got %s", r.URL.Path)
		}

		var body types.BlockCardRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body.BlockMotivation != requestData.BlockMotivation {
			t.Errorf("request blockMotivation = %v; want %v", body.BlockMotivation, requestData.BlockMotivation)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.BlockCard(context.Background(), 12345, 67890, &requestData)
	if err != nil {
		t.Fatalf("BlockCard failed: %v", err)
	}

	if response.DACode != responseData.DACode {
		t.Errorf("DACode = %d; want %d", response.DACode, responseData.DACode)
	}
	if response.Message != responseData.Message {
		t.Errorf("Message = %s; want %s", response.Message, responseData.Message)
	}
}

func TestUnblockCard(t *testing.T) {
	responseData := types.UnblockCardResponse{
		Message: "Card unblocked successfully",
		DACode:  100,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345/cards/67890/unblock" {
			t.Errorf("expected path /accounts/12345/cards/67890/unblock, got %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.UnblockCard(context.Background(), 12345, 67890)
	if err != nil {
		t.Fatalf("UnblockCard failed: %v", err)
	}

	if response.DACode != responseData.DACode {
		t.Errorf("DACode = %d; want %d", response.DACode, responseData.DACode)
	}
	if response.Message != responseData.Message {
		t.Errorf("Message = %s; want %s", response.Message, responseData.Message)
	}
}

func TestCreateVirtualCard(t *testing.T) {
	requestData := types.CreateVirtualCardRequest{
		Tag: ptrString("Shopping card"),
	}

	responseData := types.VirtualCardResponse{
		CardID:       99999,
		MaskedNumber: "1234 **** **** 9999",
		ExpiryDate:   "12/27",
		CVV:          "123",
		Status:       types.CardStatusActive,
	}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/accounts/12345/cards/virtual" {
			t.Errorf("expected path /accounts/12345/cards/virtual, got %s", r.URL.Path)
		}

		var body types.CreateVirtualCardRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(responseData)
	}))
	defer server.Close()

	client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	defer client.Close()

	response, err := client.CreateVirtualCard(context.Background(), 12345, &requestData)
	if err != nil {
		t.Fatalf("CreateVirtualCard failed: %v", err)
	}

	if response.CardID != responseData.CardID {
		t.Errorf("CardID = %d; want %d", response.CardID, responseData.CardID)
	}
	if response.CVV != responseData.CVV {
		t.Errorf("CVV = %s; want %s", response.CVV, responseData.CVV)
	}
}

// ========== ERROR CASES TESTS ==========

func TestAPIErrorHandling(t *testing.T) {
	tests := []struct {
		name         string
		endpoint     string
		method       string
		statusCode   int
		responseBody string
		expectError  bool
	}{
		{
			name:         "404 Not Found",
			endpoint:     "/accounts/99999",
			method:       http.MethodGet,
			statusCode:   http.StatusNotFound,
			responseBody: `{"message":"account not found","resource":"account"}`,
			expectError:  true,
		},
		{
			name:         "400 Bad Request",
			endpoint:     "/accounts",
			method:       http.MethodPost,
			statusCode:   http.StatusBadRequest,
			responseBody: `[{"code":"INVALID_FIELD","field":"email","message":"invalid email"}]`,
			expectError:  true,
		},
		{
			name:         "500 Internal Server Error",
			endpoint:     "/accounts/12345",
			method:       http.MethodGet,
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"internal server error"}`,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				_, _ = w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client, err := New(server.URL, "test-api-key", newTestTLSConfig(server), WithTimeout(5*time.Second))
			if err != nil {
				t.Fatalf("failed to create client: %v", err)
			}
			defer client.Close()

			var response types.AccountDataResponse
			err = client.get(context.Background(), tt.endpoint, &response)

			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

// ========== HELPER FUNCTIONS ==========

func ptrString(s string) *string {
	return &s
}
