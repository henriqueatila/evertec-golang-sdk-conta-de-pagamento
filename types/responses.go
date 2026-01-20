package types

import "time"

// CreateAccountResponse represents the response for account creation
// POST /accounts/proposal response from API spec
type CreateAccountResponse struct {
	AccountID                int64  `json:"accountId"`
	PersonID                 int64  `json:"personId"`
	PersonAddressID          int64  `json:"personAddressId"`
	CompanyID                int64  `json:"companyId"`
	CompanyAddressID         int64  `json:"companyAddressId"`
	IDProposal               int64  `json:"idProposal"`
	IDProposalAddress        int64  `json:"idProposalAddress"`
	IDCompanyProposal        int64  `json:"idCompanyProposal"`
	IDCompanyProposalAddress int64  `json:"idCompanyProposalAddress"`
	AccountNumber            int64  `json:"accountNumber"`
	Branch                   string `json:"branch"`
}

// CorporateAccountsResponse represents corporate accounts response
type CorporateAccountsResponse struct {
	Accounts []AccountDataResponse `json:"accounts"`
}

// AccountCardsResponse represents cards for an account
// GET /cards/{accountId} response from API spec
type AccountCardsResponse struct {
	Message   string         `json:"message"`
	Cards     []CardResponse `json:"cards"`
	DACode    int32          `json:"da_code"`
	AccountID int64          `json:"accountId"`
}

// BlockCardResponse represents block card operation response
// POST /cards/{accountId}/block/{cardId} response from API spec
type BlockCardResponse struct {
	Message string `json:"message"`
	DACode  int32  `json:"da_code"`
}

// UnblockCardResponse represents unblock card operation response
// POST /cards/{accountId}/unblock/{cardId} response from API spec
type UnblockCardResponse struct {
	Message string `json:"message"`
	DACode  int32  `json:"da_code"`
}

// ChangeCardPinResponse represents PIN change response
// POST /cards/{accountId}/changePin/{cardId} response from API spec
type ChangeCardPinResponse struct {
	Message string `json:"message"`
	DACode  int32  `json:"da_code"`
}

// UpdateVirtualCardTagRequest represents the request to update virtual card tag
type UpdateVirtualCardTagRequest struct {
	Tag string `json:"tag"`
}

// VirtualCardResponse represents a virtual card
type VirtualCardResponse struct {
	CardID         int64      `json:"cardId"`
	Tag            *string    `json:"tag,omitempty"`
	MaskedNumber   string     `json:"maskedNumber"`
	ExpiryDate     string     `json:"expiryDate"`
	CVV            string     `json:"cvv"`
	Status         CardStatus `json:"status"`
	PhysicalCardID *int64     `json:"physicalCardId,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
}

// VirtualCardsResponse represents list of virtual cards
// GET /cards/{accountId}/virtual response from API spec
type VirtualCardsResponse struct {
	Message string               `json:"message"`
	CardID  int64                `json:"cardId"`
	Cards   []VirtualCardResponse `json:"cards"`
	DACode  int32                `json:"da_code"`
}

// ReplacementCardResponse represents card replacement response
type ReplacementCardResponse struct {
	Message       string  `json:"message"`
	CardID        *int64  `json:"cardId,omitempty"`
	TrackingCode  *string `json:"trackingCode,omitempty"`
	EstimatedDate *string `json:"estimatedDate,omitempty"`
}

// BindAnonymousCardRequest represents request to bind anonymous card
type BindAnonymousCardRequest struct {
	CardNumber     string `json:"cardNumber"`
	LastFourDigits string `json:"lastFourDigits"`
}

// InternalTransferResponse represents internal transfer response
type InternalTransferResponse struct {
	Message            string     `json:"message"`
	TransactionID      int64      `json:"transactionId"`
	DateTimeTransfer   *time.Time `json:"dateTimeTransfer,omitempty"`
	Amount             int64      `json:"amount"`
	RecipientAccountID *int64     `json:"recipientAccountId,omitempty"`
	RecipientName      *string    `json:"recipientName,omitempty"`
	AuthenticationCode *string    `json:"authenticationCode,omitempty"`
}

// BankTransferRequest represents bank transfer request
type BankTransferRequest struct {
	Recipient         BankTransferRecipient `json:"recipient"`
	TransactionAmount int64                 `json:"transactionAmount"`
	SchedulingDate    *string               `json:"schedulingDate,omitempty"`
	Description       *string               `json:"description,omitempty"`
}

// BankTransferRecipient represents bank transfer recipient
type BankTransferRecipient struct {
	Name         string  `json:"name"`
	Document     string  `json:"document"`
	BankCode     string  `json:"bankCode"`
	Branch       string  `json:"branch"`
	Account      string  `json:"account"`
	AccountDigit *string `json:"accountDigit,omitempty"`
	AccountType  string  `json:"accountType"`
}

// BankTransferResponse represents bank transfer response
type BankTransferResponse struct {
	Message            string     `json:"message"`
	TransactionID      int64      `json:"transactionId"`
	DateTimeTransfer   *time.Time `json:"dateTimeTransfer,omitempty"`
	RecipientBankName  *string    `json:"recipientBankName,omitempty"`
	AuthenticationCode *string    `json:"authenticationCode,omitempty"`
}

// CancelTransferResponse represents cancel transfer response
type CancelTransferResponse struct {
	Message            string  `json:"message"`
	AuthenticationCode *string `json:"authenticationCode,omitempty"`
}

// ScheduledTransfersResponse represents scheduled transfers response
type ScheduledTransfersResponse struct {
	Transfers []ScheduledTransferResponse `json:"transfers"`
}

// BatchTransferResponse represents batch transfer response
type BatchTransferResponse struct {
	ProcessingCode string            `json:"processingCode"`
	Status         TransactionStatus `json:"status"`
	TransferDate   *time.Time          `json:"transferDate,omitempty"`
	Transactions   []BatchTransferItem `json:"transactions,omitempty"`
}

// BatchTransferItem represents a single item in batch transfer
type BatchTransferItem struct {
	RecipientAccountID int64             `json:"recipientAccountId"`
	Amount             int64             `json:"amount"`
	Status             TransactionStatus `json:"status"`
	Error              *string `json:"error,omitempty"`
}

// CancelInternalTransferRequest represents cancel internal transfer request
type CancelInternalTransferRequest struct {
	TransactionID int64 `json:"transactionId"`
}

// CancelInternalTransferResponse represents cancel internal transfer response
type CancelInternalTransferResponse struct {
	Message            string  `json:"message"`
	AuthenticationCode *string `json:"authenticationCode,omitempty"`
}

// TransferByIDRequest represents transfer by ID request
type TransferByIDRequest struct {
	TargetAccountID int64   `json:"targetAccountId"`
	Amount          int64   `json:"amount"`
	MerchantName    *string `json:"merchantName,omitempty"`
}

// BillPaymentResponse represents bill payment response
type BillPaymentResponse struct {
	Message            string     `json:"message"`
	AuthenticationCode *string    `json:"authenticationCode,omitempty"`
	AmountPaid         *int64     `json:"amountPaid,omitempty"`
	PaymentDate        *time.Time `json:"paymentDate,omitempty"`
}

// GetBillInfoResponse represents bill info response
type GetBillInfoResponse struct {
	OriginalValue int64   `json:"originalValue"`
	DueDate       string  `json:"dueDate"`
	Assignor      *string `json:"assignor,omitempty"`
	Payer         *string `json:"payer,omitempty"`
	Amount        *int64  `json:"amount,omitempty"`
	Fine          *int64  `json:"fine,omitempty"`
	Discount      *int64  `json:"discount,omitempty"`
}

// CancelBillResponse represents cancel bill response
type CancelBillResponse struct {
	Message            string  `json:"message"`
	AuthenticationCode *string `json:"authenticationCode,omitempty"`
}

// ScheduledBillsResponse represents scheduled bills response
type ScheduledBillsResponse struct {
	Bills []ScheduledBill `json:"bills"`
}

// ScheduledBill represents a scheduled bill
type ScheduledBill struct {
	SchedulingID int64            `json:"schedulingId"`
	Digitable    string           `json:"digitable"`
	Amount       int64            `json:"amount"`
	DueDate      string           `json:"dueDate"`
	Status       SchedulingStatus `json:"status"`
}

// BankslipsResponse represents bankslips response
type BankslipsResponse struct {
	Bankslips []BankslipItem `json:"bankslips"`
}

// BankslipItem represents a bankslip
type BankslipItem struct {
	ID            int64          `json:"id"`
	DigitableLine string         `json:"digitableLine"`
	Barcode       string         `json:"barcode"`
	Amount        int64          `json:"amount"`
	DueDate       string         `json:"dueDate"`
	Status        BankSlipStatus `json:"status"`
	DownloadURL   *string `json:"downloadUrl,omitempty"`
}

// CreateBankslipRequest represents create bankslip request
type CreateBankslipRequest struct {
	Amount      int64   `json:"amount"`
	DueDate     string  `json:"dueDate"`
	Description *string `json:"description,omitempty"`
}

// CreateBankslipResponse represents create bankslip response
type CreateBankslipResponse struct {
	IDBankslip          int64   `json:"idBankslip"`
	DigitableLine       string  `json:"digitableLine"`
	DueDate             string  `json:"dueDate"`
	BankslipDownloadURL *string `json:"bankslipDownloadUrl,omitempty"`
	Barcode             *string `json:"barcode,omitempty"`
}

// BankslipV2Request represents create bankslip v2 request
type BankslipV2Request struct {
	AccountID     int64   `json:"accountId"`
	Amount        int64   `json:"amount"`  // Amount in cents
	DueDate       string  `json:"dueDate"` // Format: YYYY-MM-DD
	PayerDocument string  `json:"payerDocument"`
	PayerName     string  `json:"payerName"`
	Description   *string `json:"description,omitempty"`
	Instructions  *string `json:"instructions,omitempty"`
}

// BankslipV2Response represents create bankslip v2 response
type BankslipV2Response struct {
	BankslipID    string         `json:"bankslipId"`
	Barcode       string         `json:"barcode"`
	DigitableLine string         `json:"digitableLine"`
	PDFBase64     string         `json:"pdfBase64"`
	Amount        int64          `json:"amount"`
	DueDate       string         `json:"dueDate"`
	Status        BankSlipStatus `json:"status"`
}

// DepositOrdersResponse represents deposit orders response
type DepositOrdersResponse struct {
	Deposits []DepositOrder `json:"deposits"`
}

// DepositOrder represents a deposit order
type DepositOrder struct {
	ID       int64            `json:"id"`
	Amount   int64            `json:"amount"`
	Covenant string           `json:"covenant"`
	Status   DepositOrderStatus `json:"status"`
}

// CreateDepositOrderResponse represents create deposit order response
type CreateDepositOrderResponse struct {
	DepositOrderID     int64              `json:"depositOrderId"`
	DateTimeExpiration *time.Time         `json:"dateTimeExpiration,omitempty"`
	Status             DepositOrderStatus `json:"status"`
}

// DoRechargeRequest represents recharge request
type DoRechargeRequest struct {
	AreaCode        string  `json:"areaCode"`
	PhoneNumber     string  `json:"phoneNumber"`
	RechargeValue   int64   `json:"rechargeValue"`
	FreeDescription *string `json:"freeDescription,omitempty"`
}

// DoRechargeResponse represents recharge response
type DoRechargeResponse struct {
	ProviderName       string  `json:"providerName"`
	RechargeValue      int64   `json:"rechargeValue"`
	AuthenticationCode *string `json:"authenticationCode,omitempty"`
}

// RechargeValuesResponse represents recharge values response
type RechargeValuesResponse struct {
	Vouchers   []RechargeVoucher `json:"vouchers"`
	ProviderID *string           `json:"providerId,omitempty"`
}

// RechargeVoucher represents a recharge voucher option
type RechargeVoucher struct {
	Value       int64   `json:"value"`
	Description *string `json:"description,omitempty"`
}

// DoVoucherRechargeRequest represents voucher recharge request
type DoVoucherRechargeRequest struct {
	ProviderID string  `json:"providerId"`
	Amount     int64   `json:"amount"`
	SignerCode *string `json:"signerCode,omitempty"`
}

// VoucherProvidersResponse represents voucher providers response
type VoucherProvidersResponse struct {
	Providers []VoucherProvider `json:"providers"`
}

// VoucherProvider represents a voucher provider
type VoucherProvider struct {
	ProviderID   string `json:"providerId"`
	ProviderName string `json:"providerName"`
	MinValue     *int64 `json:"minValue,omitempty"`
	MaxValue     *int64 `json:"maxValue,omitempty"`
}

// SimpleQRCodePaymentRequest represents simple QR code payment request
type SimpleQRCodePaymentRequest struct {
	QRCode  string  `json:"qrCode"`
	VCardID *string `json:"vCardId,omitempty"`
}

// QRCodePaymentResponse represents QR code payment response
type QRCodePaymentResponse struct {
	TransactionID   int64             `json:"transactionId"`
	Status          TransactionStatus `json:"status"`
	TransactionInfo *string `json:"transactionInfo,omitempty"`
}

// ParseQRCodeRequest represents parse QR code request
type ParseQRCodeRequest struct {
	QRCode string `json:"qrCode"`
}

// ParseQRCodeResponse represents parse QR code response
type ParseQRCodeResponse struct {
	MerchantName      *string `json:"merchantName,omitempty"`
	TransactionAmount *int64  `json:"transactionAmount,omitempty"`
	MerchantCity      *string `json:"merchantCity,omitempty"`
	PixKey            *string `json:"pixKey,omitempty"`
	TxID              *string `json:"txId,omitempty"`
}

// QRCodePublicKeyResponse represents QR code public key response
type QRCodePublicKeyResponse struct {
	KeyData       string     `json:"keyData"`
	KeyID         string     `json:"keyId"`
	KeyExpiration *time.Time `json:"keyExpiration,omitempty"`
}

// ProposalDetailResponse represents proposal details
type ProposalDetailResponse struct {
	ID               int64          `json:"id"`
	PersonalDocument string         `json:"personalDocument"`
	PersonalName     string         `json:"personalName"`
	Email            *string        `json:"email,omitempty"`
	Phone            *string        `json:"phone,omitempty"`
	BirthDate        *string        `json:"birthDate,omitempty"`
	Status           ProposalStatus `json:"status"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
}

// ProposalImage represents a proposal image
type ProposalImage struct {
	ImageType string `json:"imageType"`
	ImageURL  string `json:"imageUrl"`
}

// UpdateProposalImageRequest represents update proposal image request
type UpdateProposalImageRequest struct {
	ImageType string `json:"imageType"`
	ImageData string `json:"imageData"` // Base64 encoded
}

// ProposalTypeStatus represents proposal type status
type ProposalTypeStatus struct {
	Type   string         `json:"type"`
	Status ProposalStatus `json:"status"`
}

// LegalEntityProposalResponse represents legal entity proposal
type LegalEntityProposalResponse struct {
	ID          int64          `json:"id"`
	CompanyName string         `json:"companyName"`
	Document    string         `json:"document"`
	Status      ProposalStatus `json:"status"`
}

// LegalEntityProposalDetailResponse represents legal entity proposal details
type LegalEntityProposalDetailResponse struct {
	ID              int64          `json:"id"`
	CompanyName     string         `json:"companyName"`
	CompanyDocument string         `json:"companyDocument"`
	TradeName       *string        `json:"tradeName,omitempty"`
	FoundationDate  *string        `json:"foundationDate,omitempty"`
	Status          ProposalStatus `json:"status"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
}

// UpdateCreditExpirationRequest represents update credit expiration request
type UpdateCreditExpirationRequest struct {
	ExpirationDate string  `json:"expirationDate"`
	ID             int64   `json:"id"`
	ProcessingCode *string `json:"processingCode,omitempty"`
}

// CreditsInfoResponse represents credits info response
type CreditsInfoResponse struct {
	Amount       int64               `json:"amount"`
	Transactions []CreditTransaction `json:"transactions,omitempty"`
}

// CreditTransaction represents a credit transaction
type CreditTransaction struct {
	ID          int64      `json:"id"`
	Amount      int64      `json:"amount"`
	Date        *time.Time `json:"date,omitempty"`
	Description *string    `json:"description,omitempty"`
}

// PostPaidPaymentBalanceRequest represents post-paid payment balance request
type PostPaidPaymentBalanceRequest struct {
	AccountID       int64   `json:"accountId"`
	Amount          int64   `json:"amount"`
	ScheduleDate    *string `json:"scheduleDate,omitempty"`
	PaymentTypeEnum *string `json:"paymentTypeEnum,omitempty"`
}

// PostPaidPaymentResponse represents post-paid payment response
type PostPaidPaymentResponse struct {
	Message       string `json:"message"`
	IDTransaction *int64 `json:"idTransaction,omitempty"`
}

// PostPaidInstallmentSimulationRequest represents installment simulation request
type PostPaidInstallmentSimulationRequest struct {
	AccountID       int64  `json:"accountId"`
	Amount          int64  `json:"amount"`
	DownPayment     *int64 `json:"downPayment,omitempty"`
	NumInstallments []int  `json:"numInstallments"`
}

// PostPaidInstallmentSimulationResponse represents installment simulation response
type PostPaidInstallmentSimulationResponse struct {
	Simulations []InstallmentSimulation `json:"simulations"`
}

// InstallmentSimulation represents an installment simulation
type InstallmentSimulation struct {
	NumInstallments      int     `json:"numInstallments"`
	AmountPerInstallment int64   `json:"amountPerInstallment"`
	MonthlyInterest      float64 `json:"monthlyInterest"`
	TotalAmount          int64   `json:"totalAmount"`
}

// PostPaidInstallmentRequest represents installment request
type PostPaidInstallmentRequest struct {
	AccountID       int64   `json:"accountId"`
	Amount          int64   `json:"amount"`
	DownPayment     *int64  `json:"downPayment,omitempty"`
	NumInstallments int     `json:"numInstallments"`
	InvoiceDate     *string `json:"invoiceDate,omitempty"`
}

// PostPaidInstallmentPixResponse represents installment via PIX response
type PostPaidInstallmentPixResponse struct {
	QRCodeID    string `json:"qrCodeId"`
	QRCodeValue string `json:"qrCodeValue"`
}

// PostPaidInstallmentBalanceResponse represents installment via balance response
type PostPaidInstallmentBalanceResponse struct {
	AuthorizationID    int64   `json:"authorizationId"`
	AuthenticationCode *string `json:"authenticationCode,omitempty"`
}

// CancelInvoiceScheduleRequest represents cancel invoice schedule request
type CancelInvoiceScheduleRequest struct {
	AccountID  int64   `json:"accountId"`
	ScheduleID int64   `json:"scheduleId"`
	Reason     *string `json:"reason,omitempty"`
}

// BanksResponse represents banks list response
type BanksResponse struct {
	Banks []BankInfo `json:"banks"`
}

// BankInfo represents bank information
type BankInfo struct {
	CodeBank string `json:"codeBank"`
	BankName string `json:"bankName"`
}

// ScheduledOperationsResponse represents scheduled operations response
type ScheduledOperationsResponse struct {
	Bills     []ScheduledBill             `json:"bills,omitempty"`
	Transfers []ScheduledTransferResponse `json:"transfers,omitempty"`
	Pixs      []ScheduledPixOperation     `json:"pixs,omitempty"`
	Invoices  []ScheduledInvoice          `json:"invoices,omitempty"`
}

// ScheduledPixOperation represents a scheduled PIX operation
type ScheduledPixOperation struct {
	ID            int64            `json:"id"`
	Amount        int64            `json:"amount"`
	ScheduledDate string           `json:"scheduledDate"`
	Status        SchedulingStatus `json:"status"`
}

// ScheduledInvoice represents a scheduled invoice
type ScheduledInvoice struct {
	ID            int64            `json:"id"`
	Amount        int64            `json:"amount"`
	ScheduledDate string           `json:"scheduledDate"`
	Status        SchedulingStatus `json:"status"`
}

// IntegrationStatusResponse represents integration status response
type IntegrationStatusResponse struct {
	Authorization bool `json:"authorization"`
	Management    bool `json:"management"`
	Transactional bool `json:"transactional"`
}

// UpdatePostPaidAccountRequest represents post-paid account update request
type UpdatePostPaidAccountRequest struct {
	AccountID   int64  `json:"accountId"`
	DueDate     *int   `json:"dueDate,omitempty"`     // Day of month (1-28)
	CreditLimit *int64 `json:"creditLimit,omitempty"` // Amount in cents
}
