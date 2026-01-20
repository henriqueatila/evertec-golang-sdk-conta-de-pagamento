package types

import "time"

// CreatePixKeyRequest represents the request to create a PIX key (API: CreateKeyPixRequest)
type CreatePixKeyRequest struct {
	KeyType               PixKeyType `json:"keyType"`               // Required
	Key                   string     `json:"key"`                   // Required (renamed from keyValue)
	IgnoreTokenValidation bool       `json:"ignoreTokenValidation"` // Required
	Token                 *string    `json:"token,omitempty"`       // Optional
	TokenID               *int64     `json:"tokenId,omitempty"`     // Optional
	CountryCode           *string    `json:"countryCode,omitempty"` // Optional
}

// DeletePixKeyRequest represents the request to delete a PIX key
type DeletePixKeyRequest struct {
	KeyType PixKeyType `json:"keyType"`
	Key     string     `json:"key"` // Renamed from keyValue to match API spec
}

// PixKeyResponse represents a PIX key
type PixKeyResponse struct {
	KeyID     *int64       `json:"keyId,omitempty"`
	KeyType   PixKeyType   `json:"keyType"`
	KeyValue  string       `json:"keyValue"`
	Status    PixKeyStatus `json:"status"`
	AccountID int64        `json:"accountId"`
	CreatedAt *time.Time   `json:"createdAt,omitempty"`
}

// PixKeyListResponse represents a list of PIX keys
type PixKeyListResponse struct {
	Keys []PixKeyResponse `json:"items"`
}

// CreatePixClaimRequest represents the request to create a PIX claim
type CreatePixClaimRequest struct {
	KeyType   PixKeyType `json:"keyType"`
	KeyValue  string     `json:"keyValue"`
	ClaimType string     `json:"claimType"` // e.g., "PORTABILITY", "OWNERSHIP"
}

// ConfirmPortabilityRequest represents the request to confirm PIX portability
type ConfirmPortabilityRequest struct {
	ClaimID string `json:"claimId"`
}

// CompletePortabilityRequest represents the request to complete PIX portability
type CompletePortabilityRequest struct {
	ClaimID string `json:"claimId"`
}

// CancelPortabilityRequest represents the request to cancel PIX portability
type CancelPortabilityRequest struct {
	ClaimID string  `json:"claimId"`
	Reason  *string `json:"reason,omitempty"`
}

// PixClaimResponse represents a PIX claim
type PixClaimResponse struct {
	ClaimID    string         `json:"claimId"`
	KeyType    PixKeyType     `json:"keyType"`
	KeyValue   string         `json:"keyValue"`
	ClaimType  string         `json:"claimType"`
	Status     PixClaimStatus `json:"status"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	ResolvedAt *time.Time `json:"resolvedAt,omitempty"`
}

// PixClaimListResponse represents a list of PIX claims
type PixClaimListResponse struct {
	Claims []PixClaimResponse `json:"items"`
}

// PixPaymentRequest represents a PIX payment transaction request (POST /pix/transactions/payment)
type PixPaymentRequest struct {
	AccountID                 int64               `json:"accountId"`                           // Required
	RecipientInstitutionCode  string              `json:"recipientInstitutionCode"`            // Required
	RecipientBranchCode       string              `json:"recipientBranchCode"`                 // Required
	RecipientAccountNumber    string              `json:"recipientAccountNumber"`              // Required
	RecipientAccountType      PixAccountType      `json:"recipientAccountType"`                // Required (CACC|SLRY|SVGS|TRAN)
	RecipientCpfCnpj          string              `json:"recipientCpfCnpj"`                    // Required
	RecipientName             string              `json:"recipientName"`                       // Required
	OperationAmount           float64             `json:"operationAmount"`                     // Required
	PayerName                 *string             `json:"payerName,omitempty"`                 // Optional
	InternalReference         *string             `json:"internalReference,omitempty"`         // Optional
	EndToEnd                  *string             `json:"endToEnd,omitempty"`                  // Optional
	RecipientAddressingKey    *string             `json:"recipientAddressingKey,omitempty"`    // Optional
	FreeField                 *string             `json:"freeField,omitempty"`                 // Optional
	SchedulingDate            *string             `json:"schedulingDate,omitempty"`            // Optional (datetime)
	TransactionPurpose        *TransactionPurpose `json:"transactionPurpose,omitempty"`        // Optional (TROCO|SAQUE)
	WithdrawalServiceProvider *string             `json:"withdrawalServiceProvider,omitempty"` // Optional
	AgentMode                 *AgentMode          `json:"agentMode,omitempty"`                 // Optional (AGFSS|AGTEC|AGTOT)
	CashMoney                 *float64            `json:"cashMoney,omitempty"`                 // Optional
	Latitude                  *string             `json:"latitude,omitempty"`                  // Optional
	Longitude                 *string             `json:"longitude,omitempty"`                 // Optional
	SaveContact               *bool               `json:"saveContact,omitempty"`               // Optional
}

// QRCodeParseRequest represents the request to parse a PIX QR code
type QRCodeParseRequest struct {
	QRCodeData string `json:"qrCodeData"` // Base64 or raw QR code string
}

// QRCodeParseResponse represents parsed QR code information
type QRCodeParseResponse struct {
	KeyType           *PixKeyType `json:"keyType,omitempty"`
	KeyValue          *string     `json:"keyValue,omitempty"`
	Amount            *int64      `json:"amount,omitempty"` // Amount in cents
	RecipientName     *string     `json:"recipientName,omitempty"`
	RecipientDocument *string     `json:"recipientDocument,omitempty"`
	Description       *string     `json:"description,omitempty"`
	ExpirationDate    *time.Time  `json:"expirationDate,omitempty"`
	TransactionID     *string     `json:"transactionId,omitempty"`
}

// QRCodePaymentRequest represents a PIX payment via QR code
type QRCodePaymentRequest struct {
	QRCodeData  string  `json:"qrCodeData"`
	Amount      *int64  `json:"amount,omitempty"` // Amount in cents (if not in QR code)
	Description *string `json:"description,omitempty"`

	// Location (for fraud prevention)
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

// PixChargebackRequest represents a PIX chargeback request (POST /pix/transactions/chargeback)
type PixChargebackRequest struct {
	AccountID     int64            `json:"accountId"`            // Required
	IDTransaction int64            `json:"idTransaction"`        // Required
	Amount        float64          `json:"amount"`               // Required
	ReasonCode    ChargebackReason `json:"reasonCode"`           // Required (MD06|FR01|BE08|SL02)
	ReasonInfo    *string          `json:"reasonInfo,omitempty"` // Optional
}

// PixLimitRequest represents the request to update PIX limit
type PixLimitRequest struct {
	DailyLimit       *int64 `json:"dailyLimit,omitempty"`       // Amount in cents
	NightlyLimit     *int64 `json:"nightlyLimit,omitempty"`     // Amount in cents
	TransactionLimit *int64 `json:"transactionLimit,omitempty"` // Amount in cents
}

// PixLimitResponse represents PIX limit information
type PixLimitResponse struct {
	AccountID        int64      `json:"accountId"`
	DailyLimit       int64      `json:"dailyLimit"`               // Amount in cents
	NightlyLimit     int64      `json:"nightlyLimit"`             // Amount in cents
	TransactionLimit int64      `json:"transactionLimit"`         // Amount in cents
	DailyUsed        int64      `json:"dailyUsed"`                // Amount in cents
	NightlyUsed      int64      `json:"nightlyUsed"`              // Amount in cents
	DailyRemaining   int64      `json:"dailyRemaining"`           // Amount in cents
	NightlyRemaining int64      `json:"nightlyRemaining"`         // Amount in cents
	NightTimeStart   *string    `json:"nightTimeStart,omitempty"` // Format: HH:MM
	NightTimeEnd     *string    `json:"nightTimeEnd,omitempty"`   // Format: HH:MM
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
}

// UpdatePixNightTimeLimitRequest represents the request to update night-time PIX limit start time
type UpdatePixNightTimeLimitRequest struct {
	StartTime string `json:"startTime"` // Format: HH:MM (e.g., "20:00")
	EndTime   string `json:"endTime"`   // Format: HH:MM (e.g., "06:00")
}

// PixDeviceRequest represents the request to register a PIX device
type PixDeviceRequest struct {
	DeviceID    string  `json:"deviceId"`              // Unique device identifier
	DeviceName  string  `json:"deviceName"`            // User-friendly device name
	DeviceType  *string `json:"deviceType,omitempty"`  // e.g., "MOBILE", "TABLET"
	DeviceModel *string `json:"deviceModel,omitempty"` // e.g., "iPhone 12"
	OS          *string `json:"os,omitempty"`          // e.g., "iOS 15.0"
}

// PixDeviceResponse represents a registered PIX device
type PixDeviceResponse struct {
	DeviceID     string          `json:"deviceId"`
	DeviceName   string          `json:"deviceName"`
	DeviceType   *string         `json:"deviceType,omitempty"`
	DeviceModel  *string         `json:"deviceModel,omitempty"`
	OS           *string         `json:"os,omitempty"`
	Status       PixDeviceStatus `json:"status"` // e.g., "ACTIVE", "BLOCKED"
	AccountID    int64      `json:"accountId"`
	RegisteredAt *time.Time `json:"registeredAt,omitempty"`
	BlockedAt    *time.Time `json:"blockedAt,omitempty"`
}

// PixDeviceListResponse represents a list of PIX devices
type PixDeviceListResponse struct {
	Devices []PixDeviceResponse `json:"items"`
}

// BlockPixDeviceRequest represents the request to block a PIX device
type BlockPixDeviceRequest struct {
	DeviceID string  `json:"deviceId"`
	Reason   *string `json:"reason,omitempty"`
}

// UnblockPixDeviceRequest represents the request to unblock a PIX device
type UnblockPixDeviceRequest struct {
	DeviceID string  `json:"deviceId"`
	Reason   *string `json:"reason,omitempty"`
}

// DeletePixDeviceRequest represents the request to delete a PIX device
type DeletePixDeviceRequest struct {
	DeviceID string `json:"deviceId"`
}

// PixPublicKeyResponse represents the public key for PIX encryption
type PixPublicKeyResponse struct {
	PublicKey string     `json:"publicKey"`
	Algorithm *string    `json:"algorithm,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// ListPixClaimsRequest represents the request to list PIX claims
type ListPixClaimsRequest struct {
	Status    *string `json:"status,omitempty"`
	ClaimType *string `json:"claimType,omitempty"`
	Page      *int    `json:"page,omitempty"`
	PageSize  *int    `json:"pageSize,omitempty"`
}

// CreateClaimFromKeyRequest represents the request to create a claim from existing key
type CreateClaimFromKeyRequest struct {
	KeyType   PixKeyType `json:"keyType"`
	KeyValue  string     `json:"keyValue"`
	ClaimType string     `json:"claimType"` // PORTABILITY or OWNERSHIP
}

// ProcessLimitRequestData represents the request to process a PIX limit request
type ProcessLimitRequestData struct {
	RequestID int64   `json:"requestId"`
	Action    string  `json:"action"` // APPROVE or REJECT
	Reason    *string `json:"reason,omitempty"`
}

// RaiseLimitRequestResponse represents a raise limit request
type RaiseLimitRequestResponse struct {
	RequestID      int64              `json:"requestId"`
	AccountID      int64              `json:"accountId"`
	RequestedLimit int64              `json:"requestedLimit"` // Amount in cents
	CurrentLimit   int64              `json:"currentLimit"`   // Amount in cents
	Status         LimitRequestStatus `json:"status"`
	Reason         *string `json:"reason,omitempty"`
	CreatedAt      *string `json:"createdAt,omitempty"`
	ProcessedAt    *string `json:"processedAt,omitempty"`
}

// RaiseLimitRequestListResponse represents a list of raise limit requests
type RaiseLimitRequestListResponse struct {
	Requests []RaiseLimitRequestResponse `json:"items"`
	Total    *int                        `json:"total,omitempty"`
}

// MaximumPixLimitIssuerResponse represents the maximum PIX limit allowed by issuer
type MaximumPixLimitIssuerResponse struct {
	MaxDailyLimit       int64 `json:"maxDailyLimit"`       // Amount in cents
	MaxNightlyLimit     int64 `json:"maxNightlyLimit"`     // Amount in cents
	MaxTransactionLimit int64 `json:"maxTransactionLimit"` // Amount in cents
}

// PixPaymentResponse represents a PIX payment transaction response
type PixPaymentResponse struct {
	Message            string `json:"message"`
	IDTransaction      int64  `json:"idTransaction"`
	AuthenticationCode string `json:"authenticationCode"`
}

// PixChargebackResponse represents a PIX chargeback response
type PixChargebackResponse struct {
	Message string `json:"message"`
}

// PixCancelScheduleRequest represents a request to cancel a scheduled PIX (POST /pix/transactions/cancelSchedule)
type PixCancelScheduleRequest struct {
	AccountID  int64 `json:"accountId"`  // Required
	ScheduleID int64 `json:"scheduleId"` // Required
}

// PixCancelScheduleResponse represents the response for cancel schedule
type PixCancelScheduleResponse struct {
	AccountID any `json:"accountId"` // Can be string or int64 (API returns varying types)
}

// PixPrecautionaryBlockRequest represents a precautionary block request (POST /pix/backoffice/precautionaryBlock)
type PixPrecautionaryBlockRequest struct {
	IDTransaction int64 `json:"idTransaction"` // Required
}

// PixPrecautionaryBlockResponse represents a precautionary block response
type PixPrecautionaryBlockResponse struct {
	IDAccount     int64   `json:"idAccount"`
	IDTransaction int64   `json:"idTransaction"`
	Message       string  `json:"message"`
	Value         float64 `json:"value"`
}

// PixUpdatePrecautionaryBlockRequest represents an update to precautionary block (POST /pix/backoffice/precautionaryBlock/update)
type PixUpdatePrecautionaryBlockRequest struct {
	IDOnlineTransactionLog int64                    `json:"idOnlineTransactionLog"` // Required
	PixPrecautionaryEnum   PrecautionaryBlockStatus `json:"pixPrecautionaryEnum"`   // Required
}

// PixUpdatePrecautionaryBlockResponse represents the response for update precautionary block
type PixUpdatePrecautionaryBlockResponse struct {
	Message string `json:"message"`
}

// PixTransactionLimit represents PIX transaction limits
type PixTransactionLimit struct {
	TotalAvailableLimitFormatted          float64 `json:"totalAvailableLimitFormatted"`
	TotalAvailableLimit                   int64   `json:"totalAvailableLimit"`
	LimitPerAvailableTransaction          int64   `json:"limitPerAvailableTransaction"`
	LimitPerAvailableTransactionFormatted float64 `json:"limitPerAvailableTransactionFormatted"`
}

// PixGetLimitResponse represents PIX limit information (GET /pix/transactions/{accountId}/limit)
type PixGetLimitResponse struct {
	Message          string               `json:"message"`
	PixLimitInternal *PixTransactionLimit `json:"pixLimitInternal,omitempty"`
	PixLimitExternal *PixTransactionLimit `json:"pixLimitExternal,omitempty"`
	PixLimitWithdraw *PixTransactionLimit `json:"pixLimitWithdraw,omitempty"`
}

// PaymentOrderErrorVO represents an error in payment order
type PaymentOrderErrorVO struct {
	CodigoErro             string `json:"codigoErro"`
	CodigoErroComplementar string `json:"codigoErroComplementar"`
}

// GetPixInfoResponse represents PIX payment information (GET /pix/transactions/payment/{e2e})
type GetPixInfoResponse struct {
	TransactionID      int64                 `json:"transactionId"`
	AuthenticationCode string                `json:"authenticationCode"`
	OrderPaymentID     int64                 `json:"orderPaymentId"`
	EndToEnd           string                `json:"endToEnd"`
	Status             TransactionStatus     `json:"status"`
	Info               string                `json:"info"`
	Errors             []PaymentOrderErrorVO `json:"errors"`
	Success            bool                  `json:"success"`
	ResultDescription  string                `json:"resultDescription"`
}

// PspResponse represents a PIX Service Provider (GET /pix/psps)
type PspResponse struct {
	CodIspb                  string  `json:"codIspb"`
	NomeParticipante         string  `json:"nomeParticipante"`
	NomeResumido             string  `json:"nomeResumido"`
	TipoParticipante         string  `json:"tipoParticipante"`
	SituacaoParticipante     string  `json:"situacaoParticipante"`
	NomeAnteriorParticipante *string `json:"nomeAnteriorParticipante,omitempty"`
	DataHoraEnvioMensagem    string  `json:"dataHoraEnvioMensagem"`
	CodBanco                 *string `json:"codBanco,omitempty"`
	DataAtualizacao          *string `json:"dataAtualizacao,omitempty"`
	NomeUsuario              *string `json:"nomeUsuario,omitempty"`
	DataInclusao             *string `json:"dataInclusao,omitempty"`
	UsuarioInclusao          *string `json:"usuarioInclusao,omitempty"`
}

// PspListResponse represents a list of PSPs
type PspListResponse struct {
	PSPs []PspResponse `json:"items"`
}

// KeyStatisticDict represents statistics about a PIX key
type KeyStatisticDict map[string]any

// SearchKeyResponse represents PIX key search information (GET /pix/keys/{accountId}/{key})
type SearchKeyResponse struct {
	Success                  bool              `json:"success"`
	ResultDescription        string            `json:"resultDescription"`
	Agencia                  *string           `json:"agencia,omitempty"`
	Conta                    *string           `json:"conta,omitempty"`
	CpfCnpj                  *string           `json:"cpfCnpj,omitempty"`
	Instituicao              *string           `json:"instituicao,omitempty"`
	TipoConta                *string           `json:"tipoConta,omitempty"`
	Confirmado               *bool             `json:"confirmado,omitempty"`
	Cid                      *string           `json:"cid,omitempty"`
	Nome                     *string           `json:"nome,omitempty"`
	TipoPessoa               *string           `json:"tipoPessoa,omitempty"`
	Chave                    *string           `json:"chave,omitempty"`
	TipoChave                *string           `json:"tipoChave,omitempty"`
	Estatisticas             *KeyStatisticDict `json:"estatisticas,omitempty"`
	DataCriacao              *string           `json:"dataCriacao,omitempty"`
	DataPosse                *string           `json:"dataPosse,omitempty"`
	NomeFantasia             *string           `json:"nomeFantasia,omitempty"`
	ReivindicadaDoacao       *bool             `json:"reivindicadaDoacao,omitempty"`
	EndToEnd                 *string           `json:"endToEnd,omitempty"`
	NomePsp                  *string           `json:"nomePsp,omitempty"`
	DataAbertura             *string           `json:"dataAbertura,omitempty"`
	ReivindicacaoAbertaDesde *string           `json:"reivindicacaoAbertaDesde,omitempty"`
	FraudIdentified          *bool             `json:"fraudIdentified,omitempty"`
}
