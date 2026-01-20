package types

// ContaSimplificada represents simplified account information from API spec
// GET /accounts returns array of ContaSimplificada
type ContaSimplificada struct {
	AccountID         int64  `json:"accountId"`
	PersonalName      string `json:"personalName"`
	Email             string `json:"email"`
	PersonalDocument  string `json:"personalDocument"`
	PersonID          int64  `json:"personId"`
	ProductID         int32  `json:"productId"`
	AccountNumber     int64  `json:"accountNumber"`
	Status            int32  `json:"status"`
	StatusDescription string `json:"statusDescription"`
	MainAccountID     int64  `json:"mainAccountId"`
	ArrangementType   string `json:"arrangementType"`
}

// ContaSimplificadaListResponse represents paginated list of simplified accounts
// GET /accounts response
type ContaSimplificadaListResponse struct {
	Message string              `json:"message"`
	Contas  []ContaSimplificada `json:"contas"`
	DACode  int32               `json:"da_code"`
}
