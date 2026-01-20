package types

import (
	"net/url"
	"strings"
	"testing"
)

// Helper functions for creating pointers
func strPtr(s string) *string             { return &s }
func intPtr(i int) *int                   { return &i }
func int64Ptr(i int64) *int64             { return &i }
func boolPtr(b bool) *bool                { return &b }
func accountStatusPtr(s AccountStatus) *AccountStatus { return &s }
func accountTypePtr(t AccountType) *AccountType { return &t }
func cardStatusPtr(s CardStatus) *CardStatus { return &s }
func cardTypePtr(t CardType) *CardType { return &t }
func cardCategoryPtr(c CardCategory) *CardCategory { return &c }
func infractionStatusPtr(s InfractionReportStatus) *InfractionReportStatus { return &s }
func participantRolePtr(r ParticipantRole) *ParticipantRole { return &r }
func recurrenceStatusPtr(s RecurrenceStatus) *RecurrenceStatus { return &s }

// Helper function to parse and validate query string
func parseQueryString(qs string) (url.Values, error) {
	if qs == "" {
		return url.Values{}, nil
	}
	if !strings.HasPrefix(qs, "?") {
		return nil, nil
	}
	return url.ParseQuery(qs[1:])
}

func TestListAccountsParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListAccountsParams
		want     string
		wantVals map[string]string // expected query param values
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListAccountsParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListAccountsParams{
				Status:      accountStatusPtr(AccountStatusActive),
				AccountType: accountTypePtr(AccountTypePersonal),
				Document:    strPtr("12345678901"),
				Name:        strPtr("John Doe"),
				First:       intPtr(10),
				Max:         intPtr(50),
			},
			wantVals: map[string]string{
				"status":      string(AccountStatusActive),
				"accountType": string(AccountTypePersonal),
				"document":    "12345678901",
				"name":        "John Doe",
				"first":       "10",
				"max":         "50",
			},
		},
		{
			name: "partial fields - status and document only",
			params: &ListAccountsParams{
				Status:   accountStatusPtr(AccountStatusActive),
				Document: strPtr("12345678901"),
			},
			wantVals: map[string]string{
				"status":   string(AccountStatusActive),
				"document": "12345678901",
			},
		},
		{
			name: "partial fields - pagination only",
			params: &ListAccountsParams{
				First: intPtr(0),
				Max:   intPtr(100),
			},
			wantVals: map[string]string{
				"first": "0",
				"max":   "100",
			},
		},
		{
			name: "name with spaces - URL encoding",
			params: &ListAccountsParams{
				Name: strPtr("John Doe Jr."),
			},
			wantVals: map[string]string{
				"name": "John Doe Jr.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			if !strings.HasPrefix(got, "?") {
				t.Errorf("QueryString() should start with '?', got: %v", got)
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}

			// Verify no extra params
			if len(vals) != len(tt.wantVals) {
				t.Errorf("QueryString() has %d params, want %d", len(vals), len(tt.wantVals))
			}
		})
	}
}

func TestStatementParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *StatementParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &StatementParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &StatementParams{
				StartDate: strPtr("2024-01-01"),
				EndDate:   strPtr("2024-12-31"),
				OrderType: strPtr("ASC"),
				IsPix:     boolPtr(true),
				Type:      strPtr("credit"),
				First:     intPtr(0),
				Max:       intPtr(100),
			},
			wantVals: map[string]string{
				"startDate": "2024-01-01",
				"endDate":   "2024-12-31",
				"orderType": "ASC",
				"isPix":     "true",
				"type":      "credit",
				"first":     "0",
				"max":       "100",
			},
		},
		{
			name: "partial fields - date range only",
			params: &StatementParams{
				StartDate: strPtr("2024-01-01"),
				EndDate:   strPtr("2024-01-31"),
			},
			wantVals: map[string]string{
				"startDate": "2024-01-01",
				"endDate":   "2024-01-31",
			},
		},
		{
			name: "isPix false",
			params: &StatementParams{
				IsPix: boolPtr(false),
			},
			wantVals: map[string]string{
				"isPix": "false",
			},
		},
		{
			name: "type debit",
			params: &StatementParams{
				Type:      strPtr("debit"),
				OrderType: strPtr("DESC"),
			},
			wantVals: map[string]string{
				"type":      "debit",
				"orderType": "DESC",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListCardsParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListCardsParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListCardsParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListCardsParams{
				Status:   cardStatusPtr(CardStatusActive),
				CardType: cardTypePtr(CardTypeVirtual),
				First:    intPtr(0),
				Max:      intPtr(50),
			},
			wantVals: map[string]string{
				"status":   string(CardStatusActive),
				"cardType": string(CardTypeVirtual),
				"first":    "0",
				"max":      "50",
			},
		},
		{
			name: "status only",
			params: &ListCardsParams{
				Status: cardStatusPtr(CardStatusBlocked),
			},
			wantVals: map[string]string{
				"status": string(CardStatusBlocked),
			},
		},
		{
			name: "cardType physical",
			params: &ListCardsParams{
				CardType: cardTypePtr(CardTypePhysical),
			},
			wantVals: map[string]string{
				"cardType": string(CardTypePhysical),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestSearchCardsParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *SearchCardsParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &SearchCardsParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &SearchCardsParams{
				Status:       cardStatusPtr(CardStatusActive),
				CardType:     cardTypePtr(CardTypeVirtual),
				CardCategory: cardCategoryPtr(CardCategoryCredit),
				AccountID:    int64Ptr(123456),
				First:        intPtr(0),
				Max:          intPtr(100),
			},
			wantVals: map[string]string{
				"status":       string(CardStatusActive),
				"cardType":     string(CardTypeVirtual),
				"cardCategory": string(CardCategoryCredit),
				"accountId":    "123456",
				"first":        "0",
				"max":          "100",
			},
		},
		{
			name: "partial fields - accountId only",
			params: &SearchCardsParams{
				AccountID: int64Ptr(999999),
			},
			wantVals: map[string]string{
				"accountId": "999999",
			},
		},
		{
			name: "partial fields - status and category",
			params: &SearchCardsParams{
				Status:       cardStatusPtr(CardStatusCanceled),
				CardCategory: cardCategoryPtr(CardCategoryDebit),
			},
			wantVals: map[string]string{
				"status":       string(CardStatusCanceled),
				"cardCategory": string(CardCategoryDebit),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListProposalsParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListProposalsParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListProposalsParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListProposalsParams{
				Status:   strPtr("APPROVED"),
				Document: strPtr("12345678901"),
				First:    intPtr(0),
				Max:      intPtr(50),
			},
			wantVals: map[string]string{
				"status":   "APPROVED",
				"document": "12345678901",
				"first":    "0",
				"max":      "50",
			},
		},
		{
			name: "status only",
			params: &ListProposalsParams{
				Status: strPtr("PENDING"),
			},
			wantVals: map[string]string{
				"status": "PENDING",
			},
		},
		{
			name: "document only",
			params: &ListProposalsParams{
				Document: strPtr("98765432100"),
			},
			wantVals: map[string]string{
				"document": "98765432100",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListDepositOrdersParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListDepositOrdersParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListDepositOrdersParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListDepositOrdersParams{
				StartDate: strPtr("2024-01-01"),
				EndDate:   strPtr("2024-12-31"),
				OrderType: strPtr("ASC"),
				First:     intPtr(10),
				Max:       intPtr(100),
			},
			wantVals: map[string]string{
				"startDate": "2024-01-01",
				"endDate":   "2024-12-31",
				"orderType": "ASC",
				"first":     "10",
				"max":       "100",
			},
		},
		{
			name: "date range only",
			params: &ListDepositOrdersParams{
				StartDate: strPtr("2024-06-01"),
				EndDate:   strPtr("2024-06-30"),
			},
			wantVals: map[string]string{
				"startDate": "2024-06-01",
				"endDate":   "2024-06-30",
			},
		},
		{
			name: "orderType DESC",
			params: &ListDepositOrdersParams{
				OrderType: strPtr("DESC"),
			},
			wantVals: map[string]string{
				"orderType": "DESC",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListBanksParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListBanksParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListBanksParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListBanksParams{
				OrderType: strPtr("ASC"),
				First:     intPtr(0),
				Max:       intPtr(50),
			},
			wantVals: map[string]string{
				"orderType": "ASC",
				"first":     "0",
				"max":       "50",
			},
		},
		{
			name: "orderType only",
			params: &ListBanksParams{
				OrderType: strPtr("DESC"),
			},
			wantVals: map[string]string{
				"orderType": "DESC",
			},
		},
		{
			name: "pagination only",
			params: &ListBanksParams{
				First: intPtr(20),
				Max:   intPtr(100),
			},
			wantVals: map[string]string{
				"first": "20",
				"max":   "100",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListInfractionReportsParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListInfractionReportsParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListInfractionReportsParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListInfractionReportsParams{
				IncludeIndirectParticipants: boolPtr(true),
				IsReporter:                  boolPtr(true),
				IsCounterparty:              boolPtr(false),
				Status:                      infractionStatusPtr(InfractionStatusOpen),
				IncludeDetails:              boolPtr(true),
				ModifiedAfter:               strPtr("2024-01-01T00:00:00Z"),
				ModifiedBefore:              strPtr("2024-12-31T23:59:59Z"),
				Limit:                       intPtr(50),
			},
			wantVals: map[string]string{
				"includeIndirectParticipants": "true",
				"isReporter":                  "true",
				"isCounterparty":              "false",
				"status":                      string(InfractionStatusOpen),
				"includeDetails":              "true",
				"modifiedAfter":               "2024-01-01T00:00:00Z",
				"modifiedBefore":              "2024-12-31T23:59:59Z",
				"limit":                       "50",
			},
		},
		{
			name: "partial fields - status and limit",
			params: &ListInfractionReportsParams{
				Status: infractionStatusPtr(InfractionStatusClosed),
				Limit:  intPtr(20),
			},
			wantVals: map[string]string{
				"status": string(InfractionStatusClosed),
				"limit":  "20",
			},
		},
		{
			name: "boolean false values",
			params: &ListInfractionReportsParams{
				IsReporter:     boolPtr(false),
				IsCounterparty: boolPtr(false),
				IncludeDetails: boolPtr(false),
			},
			wantVals: map[string]string{
				"isReporter":     "false",
				"isCounterparty": "false",
				"includeDetails": "false",
			},
		},
		{
			name: "modified date range",
			params: &ListInfractionReportsParams{
				ModifiedAfter:  strPtr("2024-06-01T00:00:00Z"),
				ModifiedBefore: strPtr("2024-06-30T23:59:59Z"),
			},
			wantVals: map[string]string{
				"modifiedAfter":  "2024-06-01T00:00:00Z",
				"modifiedBefore": "2024-06-30T23:59:59Z",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListRefundsParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListRefundsParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListRefundsParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListRefundsParams{
				IncludeIndirectParticipants: boolPtr(true),
				ParticipantRole:             participantRolePtr(ParticipantRoleRequesting),
				Status:                      strPtr("PENDING"),
				IncludeDetails:              boolPtr(true),
				ModifiedAfter:               strPtr("2024-01-01T00:00:00Z"),
				ModifiedBefore:              strPtr("2024-12-31T23:59:59Z"),
				Limit:                       intPtr(100),
			},
			wantVals: map[string]string{
				"includeIndirectParticipants": "true",
				"participantRole":             string(ParticipantRoleRequesting),
				"status":                      "PENDING",
				"includeDetails":              "true",
				"modifiedAfter":               "2024-01-01T00:00:00Z",
				"modifiedBefore":              "2024-12-31T23:59:59Z",
				"limit":                       "100",
			},
		},
		{
			name: "participant role contested",
			params: &ListRefundsParams{
				ParticipantRole: participantRolePtr(ParticipantRoleContested),
				Limit:           intPtr(20),
			},
			wantVals: map[string]string{
				"participantRole": string(ParticipantRoleContested),
				"limit":           "20",
			},
		},
		{
			name: "status only",
			params: &ListRefundsParams{
				Status: strPtr("COMPLETED"),
			},
			wantVals: map[string]string{
				"status": "COMPLETED",
			},
		},
		{
			name: "includeDetails false",
			params: &ListRefundsParams{
				IncludeDetails: boolPtr(false),
			},
			wantVals: map[string]string{
				"includeDetails": "false",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListAutomaticPixParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListAutomaticPixParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListAutomaticPixParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListAutomaticPixParams{
				Inactive:         boolPtr(false),
				RecurrenceID:     strPtr("REC123456789"),
				Page:             intPtr(1),
				Size:             intPtr(50),
				RecurrenceStatus: recurrenceStatusPtr(RecurrenceStatusConfirmed),
				IsPayer:          boolPtr(true),
			},
			wantVals: map[string]string{
				"inactive":         "false",
				"recurrenceId":     "REC123456789",
				"page":             "1",
				"size":             "50",
				"recurrenceStatus": string(RecurrenceStatusConfirmed),
				"isPayer":          "true",
			},
		},
		{
			name: "partial fields - recurrence status only",
			params: &ListAutomaticPixParams{
				RecurrenceStatus: recurrenceStatusPtr(RecurrenceStatusPending),
			},
			wantVals: map[string]string{
				"recurrenceStatus": string(RecurrenceStatusPending),
			},
		},
		{
			name: "partial fields - pagination",
			params: &ListAutomaticPixParams{
				Page: intPtr(0),
				Size: intPtr(100),
			},
			wantVals: map[string]string{
				"page": "0",
				"size": "100",
			},
		},
		{
			name: "inactive true and isPayer false",
			params: &ListAutomaticPixParams{
				Inactive: boolPtr(true),
				IsPayer:  boolPtr(false),
			},
			wantVals: map[string]string{
				"inactive": "true",
				"isPayer":  "false",
			},
		},
		{
			name: "recurrenceId only",
			params: &ListAutomaticPixParams{
				RecurrenceID: strPtr("ABC-DEF-GHI-123"),
			},
			wantVals: map[string]string{
				"recurrenceId": "ABC-DEF-GHI-123",
			},
		},
		{
			name: "cancelled status",
			params: &ListAutomaticPixParams{
				RecurrenceStatus: recurrenceStatusPtr(RecurrenceStatusCancelled),
			},
			wantVals: map[string]string{
				"recurrenceStatus": string(RecurrenceStatusCancelled),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

func TestListHceDevicesParams_QueryString(t *testing.T) {
	tests := []struct {
		name     string
		params   *ListHceDevicesParams
		want     string
		wantVals map[string]string
	}{
		{
			name:   "nil receiver",
			params: nil,
			want:   "",
		},
		{
			name:   "empty struct",
			params: &ListHceDevicesParams{},
			want:   "",
		},
		{
			name: "all fields set",
			params: &ListHceDevicesParams{
				AccountID: int64Ptr(123456),
				Status:    strPtr("ACTIVE"),
				Page:      intPtr(1),
				PageSize:  intPtr(50),
			},
			wantVals: map[string]string{
				"accountId": "123456",
				"status":    "ACTIVE",
				"page":      "1",
				"pageSize":  "50",
			},
		},
		{
			name: "accountId only",
			params: &ListHceDevicesParams{
				AccountID: int64Ptr(999999),
			},
			wantVals: map[string]string{
				"accountId": "999999",
			},
		},
		{
			name: "status only",
			params: &ListHceDevicesParams{
				Status: strPtr("INACTIVE"),
			},
			wantVals: map[string]string{
				"status": "INACTIVE",
			},
		},
		{
			name: "pagination only",
			params: &ListHceDevicesParams{
				Page:     intPtr(0),
				PageSize: intPtr(100),
			},
			wantVals: map[string]string{
				"page":     "0",
				"pageSize": "100",
			},
		},
		{
			name: "accountId and pagination",
			params: &ListHceDevicesParams{
				AccountID: int64Ptr(12345),
				Page:      intPtr(2),
				PageSize:  intPtr(25),
			},
			wantVals: map[string]string{
				"accountId": "12345",
				"page":      "2",
				"pageSize":  "25",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if tt.want != "" {
				if got != tt.want {
					t.Errorf("QueryString() = %v, want %v", got, tt.want)
				}
				return
			}

			if tt.wantVals == nil {
				if got != "" {
					t.Errorf("QueryString() = %v, want empty string", got)
				}
				return
			}

			vals, err := parseQueryString(got)
			if err != nil {
				t.Errorf("Invalid query string: %v", err)
				return
			}

			for key, expectedVal := range tt.wantVals {
				if gotVal := vals.Get(key); gotVal != expectedVal {
					t.Errorf("QueryString() param %s = %v, want %v", key, gotVal, expectedVal)
				}
			}
		})
	}
}

// Test URL encoding for special characters
func TestQueryString_URLEncoding(t *testing.T) {
	tests := []struct {
		name   string
		params *ListAccountsParams
	}{
		{
			name: "name with special characters",
			params: &ListAccountsParams{
				Name: strPtr("José & Maria da Silva"),
			},
		},
		{
			name: "document with special chars",
			params: &ListAccountsParams{
				Document: strPtr("123.456.789-01"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.QueryString()

			if !strings.HasPrefix(got, "?") {
				t.Errorf("QueryString() should start with '?'")
				return
			}

			// Verify it can be parsed
			_, err := url.ParseQuery(got[1:])
			if err != nil {
				t.Errorf("QueryString() produced invalid URL encoding: %v", err)
			}
		})
	}
}
