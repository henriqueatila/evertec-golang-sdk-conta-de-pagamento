package types

// AccountStatus represents the status of an account
type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "ACTIVE"
	AccountStatusInactive  AccountStatus = "INACTIVE"
	AccountStatusBlocked   AccountStatus = "BLOCKED"
	AccountStatusSuspended AccountStatus = "SUSPENDED"
	AccountStatusClosed    AccountStatus = "CLOSED"
	AccountStatusPending   AccountStatus = "PENDING"
)

// AccountType represents the type of account (general)
type AccountType string

const (
	AccountTypePersonal AccountType = "PERSONAL"
	AccountTypeCompany  AccountType = "COMPANY"
)

// PixAccountType represents the type of recipient account for PIX transactions
type PixAccountType string

const (
	PixAccountTypeCACC PixAccountType = "CACC" // Conta Corrente
	PixAccountTypeSLRY PixAccountType = "SLRY" // Conta Salário
	PixAccountTypeSVGS PixAccountType = "SVGS" // Conta Poupança
	PixAccountTypeTRAN PixAccountType = "TRAN" // Conta de Pagamento
)

// CardStatus represents the status of a card
type CardStatus string

const (
	CardStatusActive   CardStatus = "ACTIVE"
	CardStatusBlocked  CardStatus = "BLOCKED"
	CardStatusCanceled CardStatus = "CANCELED"
	CardStatusPending  CardStatus = "PENDING"
	CardStatusExpired  CardStatus = "EXPIRED"
)

// CardType represents the type of card
type CardType string

const (
	CardTypePhysical CardType = "PHYSICAL"
	CardTypeVirtual  CardType = "VIRTUAL"
)

// CardCategory represents the category of card
type CardCategory string

const (
	CardCategoryDebit    CardCategory = "DEBIT"
	CardCategoryCredit   CardCategory = "CREDIT"
	CardCategoryPostpaid CardCategory = "POSTPAID"
)

// TransactionStatus represents the status of a transaction
type TransactionStatus string

const (
	TransactionStatusApproved   TransactionStatus = "APPROVED"
	TransactionStatusRejected   TransactionStatus = "REJECTED"
	TransactionStatusPending    TransactionStatus = "PENDING"
	TransactionStatusCanceled   TransactionStatus = "CANCELED"
	TransactionStatusScheduled  TransactionStatus = "SCHEDULED"
	TransactionStatusProcessing TransactionStatus = "PROCESSING"
)

// TransactionType represents the type of transaction
type TransactionType string

const (
	TransactionTypeTransfer    TransactionType = "TRANSFER"
	TransactionTypeTED         TransactionType = "TED"
	TransactionTypeDOC         TransactionType = "DOC"
	TransactionTypePix         TransactionType = "PIX"
	TransactionTypeBillPayment TransactionType = "BILL_PAYMENT"
	TransactionTypePurchase    TransactionType = "PURCHASE"
	TransactionTypeRecharge    TransactionType = "RECHARGE"
	TransactionTypeChargeback  TransactionType = "CHARGEBACK"
)

// PixKeyType represents the type of PIX key
type PixKeyType string

const (
	PixKeyTypeCPF    PixKeyType = "CPF"
	PixKeyTypeCNPJ   PixKeyType = "CNPJ"
	PixKeyTypeEmail  PixKeyType = "EMAIL"
	PixKeyTypePhone  PixKeyType = "PHONE"
	PixKeyTypeRandom PixKeyType = "EVP"
)

// PixKeyStatus represents the status of a PIX key
type PixKeyStatus string

const (
	PixKeyStatusActive      PixKeyStatus = "ACTIVE"
	PixKeyStatusInactive    PixKeyStatus = "INACTIVE"
	PixKeyStatusPortability PixKeyStatus = "PORTABILITY"
	PixKeyStatusClaimed     PixKeyStatus = "CLAIMED"
)

// PixMovementType represents the type of PIX movement
type PixMovementType string

const (
	PixMovementTypeReceived PixMovementType = "RECEIVED"
	PixMovementTypeSent     PixMovementType = "SENT"
)

// PaymentType represents the type of payment
type PaymentType string

const (
	PaymentTypeDebitCard  PaymentType = "DEBIT_CARD"
	PaymentTypeCreditCard PaymentType = "CREDIT_CARD"
	PaymentTypePix        PaymentType = "PIX"
	PaymentTypeBankSlip   PaymentType = "BANK_SLIP"
	PaymentTypeBalance    PaymentType = "BALANCE"
)

// AssetType represents the type of asset used in a transaction
type AssetType string

const (
	AssetTypeDigitalAccount AssetType = "DIGITAL_ACCOUNT"
	AssetTypeVoucher        AssetType = "VOUCHER"
	AssetTypeUnknown        AssetType = "UNKNOWN"
)

// LimitType represents the type of limit
type LimitType string

const (
	LimitTypePix         LimitType = "PIX"
	LimitTypeTransfer    LimitType = "TRANSFER"
	LimitTypeTED         LimitType = "TED"
	LimitTypeBillPayment LimitType = "BILL_PAYMENT"
	LimitTypeWithdrawal  LimitType = "WITHDRAWAL"
	LimitTypePurchase    LimitType = "PURCHASE"
)

// DocumentType represents the type of document
type DocumentType string

const (
	DocumentTypeCPF      DocumentType = "CPF"
	DocumentTypeCNPJ     DocumentType = "CNPJ"
	DocumentTypeRG       DocumentType = "RG"
	DocumentTypePassport DocumentType = "PASSPORT"
)

// AddressType represents the type of address
type AddressType string

const (
	AddressTypeResidential AddressType = "RESIDENTIAL"
	AddressTypeCommercial  AddressType = "COMMERCIAL"
)

// Gender represents the gender of a person
type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther  Gender = "OTHER"
)

// ChargebackMode represents the chargeback mode
type ChargebackMode string

const (
	ChargebackModeTotal   ChargebackMode = "total"
	ChargebackModePartial ChargebackMode = "partial"
)

// ChargebackReason represents the reason code for PIX chargebacks
type ChargebackReason string

const (
	ChargebackReasonMD06 ChargebackReason = "MD06" // Reembolso solicitado pelo cliente final
	ChargebackReasonFR01 ChargebackReason = "FR01" // Fraude
	ChargebackReasonBE08 ChargebackReason = "BE08" // Erro operacional
	ChargebackReasonSL02 ChargebackReason = "SL02" // Crédito não processado
)

// TransactionPurpose represents the purpose of PIX transaction
type TransactionPurpose string

const (
	TransactionPurposeTROCO TransactionPurpose = "TROCO" // Troco
	TransactionPurposeSAQUE TransactionPurpose = "SAQUE" // Saque
)

// AgentMode represents the agent mode for withdrawal operations
type AgentMode string

const (
	AgentModeAGFSS AgentMode = "AGFSS" // Facilitador de serviço de saque
	AgentModeAGTEC AgentMode = "AGTEC" // Correspondente no País
	AgentModeAGTOT AgentMode = "AGTOT" // Outro tipo de agente
)

// PrecautionaryBlockStatus represents the status of a precautionary block
type PrecautionaryBlockStatus string

const (
	PrecautionaryBlockNotCompleted             PrecautionaryBlockStatus = "BLOQUEIO_NAO_CONCLUIDO"        // Bloqueio não concluído
	PrecautionaryBlockReleased                 PrecautionaryBlockStatus = "BLOQUEIO_LIBERADO"             // Bloqueio liberado
	PrecautionaryBlockConvertedToFraudAnalysis PrecautionaryBlockStatus = "BLOQUEIO_CONVERTIDO_MED"       // Bloqueio convertido em MED
	PrecautionaryBlockReleasedWithoutAnalysis  PrecautionaryBlockStatus = "BLOQUEIO_LIBERADO_SEM_ANALISE" // Bloqueio liberado sem análise
)

// WebhookEventType represents the type of webhook event
type WebhookEventType string

const (
	WebhookEventPixMovement          WebhookEventType = "movimento_pix"
	WebhookEventScheduledPixExecuted WebhookEventType = "pix_agendado_executado"
	WebhookEventPrecautionaryBlock   WebhookEventType = "notifica_bloqueio_cautelar"
	WebhookEventRetainedValue        WebhookEventType = "valor_retido"
	WebhookEventAutomaticPix         WebhookEventType = "notificacao_pix_automatico"
	WebhookEventClaimNotification    WebhookEventType = "notifica_reivindicacao"
)

// PrecautionaryBlockType represents the type of precautionary block
type PrecautionaryBlockType string

const (
	PrecautionaryBlockTypeBlock   PrecautionaryBlockType = "BLOCK"
	PrecautionaryBlockTypeUnblock PrecautionaryBlockType = "UNBLOCK"
)

// AutomaticPixType represents the type of automatic PIX notification
type AutomaticPixType string

const (
	AutomaticPixTypeAdesao        AutomaticPixType = "ADESAO_PIX_AUTOMATICO"
	AutomaticPixTypeCancelamento  AutomaticPixType = "CANCELAMENTO_PIX_AUTOMATICO"
	AutomaticPixTypeFluxoCompleto AutomaticPixType = "FLUXO_COMPLETO_PIX_AUTOMATICO"
)

// DocType represents the type of document image
type DocType string

const (
	DocTypeFront   DocType = "FRONT"
	DocTypeBack    DocType = "BACK"
	DocTypeSelfie  DocType = "SELFIE"
	DocTypeAddress DocType = "ADDRESS_PROOF"
	DocTypeIncome  DocType = "INCOME_PROOF"
)

// PaymentStatus represents the payment status of a bank slip or similar
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusPaid      PaymentStatus = "PAID"
	PaymentStatusCanceled  PaymentStatus = "CANCELED"
	PaymentStatusExpired   PaymentStatus = "EXPIRED"
	PaymentStatusScheduled PaymentStatus = "SCHEDULED"
)

// CreditStatus represents the status of a credit/voucher
type CreditStatus string

const (
	CreditStatusAvailable CreditStatus = "AVAILABLE"
	CreditStatusUsed      CreditStatus = "USED"
	CreditStatusExpired   CreditStatus = "EXPIRED"
)

// ProductStatus represents the status of a product
type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "ACTIVE"
	ProductStatusInactive ProductStatus = "INACTIVE"
)

// InstitutionStatus represents the status of an institution
type InstitutionStatus string

const (
	InstitutionStatusActive   InstitutionStatus = "ACTIVE"
	InstitutionStatusInactive InstitutionStatus = "INACTIVE"
)

// BiroAnalysisStatus represents the status of a BIRO credit analysis
type BiroAnalysisStatus string

const (
	BiroAnalysisStatusPending   BiroAnalysisStatus = "PENDING"
	BiroAnalysisStatusApproved  BiroAnalysisStatus = "APPROVED"
	BiroAnalysisStatusRejected  BiroAnalysisStatus = "REJECTED"
	BiroAnalysisStatusProcessing BiroAnalysisStatus = "PROCESSING"
)

// ProcessorStatus represents the status of processor sync
type ProcessorStatus string

const (
	ProcessorStatusSynced   ProcessorStatus = "SYNCED"
	ProcessorStatusPending  ProcessorStatus = "PENDING"
	ProcessorStatusError    ProcessorStatus = "ERROR"
)

// HceDeviceStatus represents the status of an HCE device
type HceDeviceStatus string

const (
	HceDeviceStatusActive   HceDeviceStatus = "ACTIVE"
	HceDeviceStatusInactive HceDeviceStatus = "INACTIVE"
	HceDeviceStatusBlocked  HceDeviceStatus = "BLOCKED"
)

// SchedulingStatus represents the status of a scheduled operation
type SchedulingStatus string

const (
	SchedulingStatusPending   SchedulingStatus = "PENDING"
	SchedulingStatusExecuted  SchedulingStatus = "EXECUTED"
	SchedulingStatusCanceled  SchedulingStatus = "CANCELED"
	SchedulingStatusFailed    SchedulingStatus = "FAILED"
)

// PixClaimStatus represents the status of a PIX claim
type PixClaimStatus string

const (
	PixClaimStatusPending   PixClaimStatus = "PENDING"
	PixClaimStatusCompleted PixClaimStatus = "COMPLETED"
	PixClaimStatusCanceled  PixClaimStatus = "CANCELED"
	PixClaimStatusRejected  PixClaimStatus = "REJECTED"
)

// PixDeviceStatus represents the status of a PIX device
type PixDeviceStatus string

const (
	PixDeviceStatusActive  PixDeviceStatus = "ACTIVE"
	PixDeviceStatusBlocked PixDeviceStatus = "BLOCKED"
)

// LimitRequestStatus represents the status of a limit increase request
type LimitRequestStatus string

const (
	LimitRequestStatusPending  LimitRequestStatus = "PENDING"
	LimitRequestStatusApproved LimitRequestStatus = "APPROVED"
	LimitRequestStatusRejected LimitRequestStatus = "REJECTED"
)

// BalanceLockStatus represents the status of a balance lock
type BalanceLockStatus string

const (
	BalanceLockStatusActive   BalanceLockStatus = "ACTIVE"
	BalanceLockStatusReleased BalanceLockStatus = "RELEASED"
	BalanceLockStatusExpired  BalanceLockStatus = "EXPIRED"
)

// RefundStatus represents the status of a MED refund
type RefundStatus string

const (
	RefundStatusOpen      RefundStatus = "OPEN"
	RefundStatusClosed    RefundStatus = "CLOSED"
	RefundStatusCancelled RefundStatus = "CANCELLED"
)
