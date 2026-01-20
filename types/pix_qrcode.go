package types

import (
	"fmt"
	"time"
)

// PIX QR Code Operations - Updated to match OpenAPI specification

// QRCodeFormat represents QR code format types
type QRCodeFormat int32

const (
	QRCodeFormatBRCode      QRCodeFormat = 1 // BR CODE
	QRCodeFormatBase64      QRCodeFormat = 2 // BASE64
	QRCodeFormatImageBase64 QRCodeFormat = 3 // IMAGE BASE64
)

// PayerType represents payer type (PF or PJ)
type PayerType int32

const (
	PayerTypePF PayerType = 1 // Pessoa Física
	PayerTypePJ PayerType = 2 // Pessoa Jurídica
)

// FineModality represents fine modality types
type FineModality string

const (
	FineModalityValorFixo  FineModality = "VALOR_FIXO"
	FineModalityPercentual FineModality = "PERCENTUAL"
)

// InterestModality represents interest modality types
type InterestModality string

const (
	InterestModalityValorDiasAtraso      InterestModality = "VALOR_DIAS_ATRASO"
	InterestModalityPercentualDiasAtraso InterestModality = "PERCENTUAL_DIAS_ATRASO"
	InterestModalityValorMensal          InterestModality = "VALOR_MENSAL"
	InterestModalityPercentualMensal     InterestModality = "PERCENTUAL_MENSAL"
	InterestModalityValorAnual           InterestModality = "VALOR_ANUAL"
	InterestModalityPercentualAnual      InterestModality = "PERCENTUAL_ANUAL"
	InterestModalityIsento               InterestModality = "ISENTO"
	InterestModalityNaoTemJuros          InterestModality = "NAO_TEM_JUROS"
)

// DiscountModality represents discount modality types
type DiscountModality string

const (
	DiscountModalityValorFixo                       DiscountModality = "VALOR_FIXO"
	DiscountModalityPercentual                      DiscountModality = "PERCENTUAL"
	DiscountModalityValorAntecipacaoDiaCorrido      DiscountModality = "VALOR_ANTECIPACAO_DIA_CORRIDO"
	DiscountModalityValorAntecipacaoDiaUtil         DiscountModality = "VALOR_ANTECIPACAO_DIA_UTIL"
	DiscountModalityPercentualAntecipacaoDiaCorrido DiscountModality = "PERCENTUAL_ANTECIPACAO_DIA_CORRIDO"
	DiscountModalityPercentualAntecipacaoDiaUtil    DiscountModality = "PERCENTUAL_ANTECIPACAO_DIA_UTIL"
)

// ReductionModality represents reduction modality types
type ReductionModality string

const (
	ReductionModalityValorFixo  ReductionModality = "VALOR_FIXO"
	ReductionModalityPercentual ReductionModality = "PERCENTUAL"
)

// ChangeModality represents change modality
type ChangeModality int32

const (
	ChangeModalityNotAllowed ChangeModality = 0
	ChangeModalityAllowed    ChangeModality = 1
)

// QRCodeType represents QR code type
type QRCodeType string

const (
	QRCodeTypeStatic  QRCodeType = "STATIC"
	QRCodeTypeDynamic QRCodeType = "DYNAMIC"
)

// PIXType represents PIX transaction type
type PIXType string

const (
	PIXTypeSaque  PIXType = "SAQUE"
	PIXTypeTroco  PIXType = "TROCO"
	PIXTypeNormal PIXType = "NORMAL"
)

// FlowType represents the flow type for decoded QR codes
type FlowType string

const (
	FlowTypeStatic            FlowType = "STATIC"
	FlowTypeImmediateCharge   FlowType = "IMMEDIATE_CHARGE"
	FlowTypeChargeWithDueDate FlowType = "CHARGE_WITH_DUE_DATE"
	FlowTypeAUT2              FlowType = "AUT2"
	FlowTypeAUT3              FlowType = "AUT3"
	FlowTypeAUT4              FlowType = "AUT4"
)

// AdditionalInfo represents additional information for QR code
type AdditionalInfo struct {
	Nome  string `json:"nome"`
	Valor string `json:"valor"`
}

// DiscountFixedDate represents discount fixed date
type DiscountFixedDate struct {
	Data             string  `json:"data"`
	ValorDescontoAbs float64 `json:"valorDescontoAbs,omitempty"`
	ValorPerc        float64 `json:"valorPerc,omitempty"`
}

// StaticQRCodeRequest represents a static PIX QR code creation request
// POST /pix/qrcodes/static
type StaticQRCodeRequest struct {
	AccountID     int64        `json:"accountId"`               // Required
	AddressKey    *string      `json:"addressKey,omitempty"`    // PIX key (optional)
	QRCodeFormat  QRCodeFormat `json:"qrCodeFormat"`            // Required: 1=BR CODE, 2=BASE64, 3=IMAGE BASE64
	PostalCode    *string      `json:"postalCode,omitempty"`    // Optional
	Identificador string       `json:"identificador"`           // Required: QR code identifier
	Descricao     *string      `json:"descricao,omitempty"`     // Optional
	Amount        *float64     `json:"amount,omitempty"`        // Optional: Value in reals
	InvoiceDate   *Date        `json:"invoiceDate,omitempty"`   // Optional
	InvoiceQRCode *bool        `json:"invoiceQrCode,omitempty"` // Optional
}

// DynamicQRCodeRequest represents a dynamic PIX QR code creation request
// POST /pix/qrcodes/dynamic
type DynamicQRCodeRequest struct {
	AccountID             int64               `json:"accountId"`                       // Required
	AddressKey            string              `json:"addressKey"`                      // Required: PIX key
	Amount                float64             `json:"amount"`                          // Required: minimum 1
	QRCodeFormat          QRCodeFormat        `json:"qrCodeFormat"`                    // Required: 1=BR CODE, 2=BASE64, 3=IMAGE BASE64
	PostalCode            *string             `json:"postalCode,omitempty"`            // Optional
	ExpirationSeconds     *int64              `json:"expirationSeconds,omitempty"`     // Optional
	ValityAfterExpiration *int32              `json:"valityAfterExpiration,omitempty"` // Optional: Days
	DueDate               *string             `json:"dueDate,omitempty"`               // Optional
	PayerCpfCnpj          *string             `json:"payerCpfCnpj,omitempty"`          // Optional
	PayerType             *PayerType          `json:"payerType,omitempty"`             // Optional: 1=PF, 2=PJ
	PayerName             *string             `json:"payerName,omitempty"`             // Optional
	InterestValue         *float64            `json:"interestValue,omitempty"`         // Optional
	FineAmount            *float64            `json:"fineAmount,omitempty"`            // Optional
	DiscountAmount        *float64            `json:"discountAmount,omitempty"`        // Optional
	ReductionAmount       *float64            `json:"reductionAmount,omitempty"`       // Optional
	PayerRequest          *string             `json:"payerRequest,omitempty"`          // Optional
	AdditionalInfo        []AdditionalInfo    `json:"additionalInfo,omitempty"`        // Optional
	Reusable              *bool               `json:"reusable,omitempty"`              // Optional
	InvoiceDate           *Date               `json:"invoiceDate,omitempty"`           // Optional
	InvoiceQRCode         *bool               `json:"invoiceQrCode,omitempty"`         // Optional
	PayerEmail            *string             `json:"payerEmail,omitempty"`            // Optional
	PayerZipCode          *string             `json:"payerZipCode,omitempty"`          // Optional
	PayerAddress          *string             `json:"payerAddress,omitempty"`          // Optional
	PayerCity             *string             `json:"payerCity,omitempty"`             // Optional
	PayerState            *string             `json:"payerState,omitempty"`            // Optional
	FineModality          *FineModality       `json:"fineModality,omitempty"`          // Optional: VALOR_FIXO, PERCENTUAL
	InterestModality      *InterestModality   `json:"interestModality,omitempty"`      // Optional: 8 options
	DiscountModality      *DiscountModality   `json:"discountModality,omitempty"`      // Optional: 6 options
	ReductionModality     *ReductionModality  `json:"reductionModality,omitempty"`     // Optional: VALOR_FIXO, PERCENTUAL
	DiscountFixedDate     []DiscountFixedDate `json:"discountFixedDate,omitempty"`     // Optional
	ChangeModality        *ChangeModality     `json:"changeModality,omitempty"`        // Optional: 0 or 1
}

// QRCodeResponse represents a PIX QR code response
// 200 response for both static and dynamic QR code creation
type QRCodeResponse struct {
	Message           string `json:"message"`
	QRCodeID          int64  `json:"qrCodeId"`
	QRCodeValue       string `json:"qrCodeValue"`
	InternalReference string `json:"internalReference"`
}

// QRCodeQueryProcessingRequest represents query processing request
// POST /pix/qrcodes/query-processing
type QRCodeQueryProcessingRequest struct {
	AccountID    int64        `json:"accountId"`          // Required
	Document     *string      `json:"document,omitempty"` // Optional
	QRCodeValue  string       `json:"qrCodeValue"`        // Required
	QRCodeFormat QRCodeFormat `json:"qrCodeFormat"`       // Required
}

// Estatisticas represents statistics in query processing response
type Estatisticas struct {
	QuantidadeRecebida *int32   `json:"quantidadeRecebida,omitempty"`
	ValorRecebido      *float64 `json:"valorRecebido,omitempty"`
}

// InfoAdicional represents additional info in query processing response
type InfoAdicional struct {
	Nome  string `json:"nome"`
	Valor string `json:"valor"`
}

// ErroInfo represents error information
type ErroInfo struct {
	Codigo    *string `json:"codigo,omitempty"`
	Descricao *string `json:"descricao,omitempty"`
}

// QRCodeQueryProcessingResponse represents extensive query processing response
type QRCodeQueryProcessingResponse struct {
	Success                       *bool           `json:"success,omitempty"`
	ResultDescription             *string         `json:"resultDescription,omitempty"`
	Nome                          *string         `json:"nome,omitempty"`
	NomeFantasia                  *string         `json:"nomeFantasia,omitempty"`
	CpfCnpj                       *string         `json:"cpfCnpj,omitempty"`
	NomePsp                       *string         `json:"nomePsp,omitempty"`
	CodInstituicao                *string         `json:"codInstituicao,omitempty"`
	CodAgencia                    *string         `json:"codAgencia,omitempty"`
	NroConta                      *string         `json:"nroConta,omitempty"`
	TipoConta                     *string         `json:"tipoConta,omitempty"`
	DataCriacao                   *string         `json:"dataCriacao,omitempty"`
	DataPosse                     *string         `json:"dataPosse,omitempty"`
	DataAbertura                  *string         `json:"dataAbertura,omitempty"`
	Referencia                    *string         `json:"referencia,omitempty"`
	Info                          *string         `json:"info,omitempty"`
	Valor                         *float64        `json:"valor,omitempty"`
	TipoQRCode                    *string         `json:"tipoQRCode,omitempty"`
	EndToEnd                      *string         `json:"endToEnd,omitempty"`
	Estatisticas                  *Estatisticas   `json:"estatisticas,omitempty"`
	InfoAdicional                 []InfoAdicional `json:"infoAdicional,omitempty"`
	DocumentoId                   *string         `json:"documentoId,omitempty"`
	DocumentoRevisao              *string         `json:"documentoRevisao,omitempty"`
	CalendarioExpiracaoSegundos   *int32          `json:"calendarioExpiracaoSegundos,omitempty"`
	CalendarioVencimento          *string         `json:"calendarioVencimento,omitempty"`
	Dpp                           *string         `json:"dpp,omitempty"`
	ValidadeAposVencimento        *int32          `json:"validadeAposVencimento,omitempty"`
	CalendarioApresentacao        *string         `json:"calendarioApresentacao,omitempty"`
	CalendarioCriacao             *string         `json:"calendarioCriacao,omitempty"`
	TipoPessoa                    *string         `json:"tipoPessoa,omitempty"`
	PagadorCpf                    *string         `json:"pagadorCpf,omitempty"`
	PagadorCnpj                   *string         `json:"pagadorCnpj,omitempty"`
	PagadorNome                   *string         `json:"pagadorNome,omitempty"`
	ValorFinal                    *float64        `json:"valorFinal,omitempty"`
	ValorJuros                    *float64        `json:"valorJuros,omitempty"`
	ValorMulta                    *float64        `json:"valorMulta,omitempty"`
	ValorDesconto                 *float64        `json:"valorDesconto,omitempty"`
	ValorAbatimento               *float64        `json:"valorAbatimento,omitempty"`
	SolicitacaoPagador            *string         `json:"solicitacaoPagador,omitempty"`
	CodMun                        *string         `json:"codMun,omitempty"`
	TipoChave                     *string         `json:"tipoChave,omitempty"`
	ChaveEnderecamento            *string         `json:"chaveEnderecamento,omitempty"`
	Erro                          *ErroInfo       `json:"erro,omitempty"`
	Identificador                 *string         `json:"identificador,omitempty"`
	Descricao                     *string         `json:"descricao,omitempty"`
	TipoPix                       *PIXType        `json:"tipoPix,omitempty"` // SAQUE, TROCO, NORMAL
	ValorTrocoSaque               *float64        `json:"valorTrocoSaque,omitempty"`
	ModalidadeAgente              *string         `json:"modalidadeAgente,omitempty"`
	PrestadorDoServicoDeSaque     *string         `json:"prestadorDoServicoDeSaque,omitempty"`
	RecebedorNomeFantasia         *string         `json:"recebedorNomeFantasia,omitempty"`
	RecebedorNome                 *string         `json:"recebedorNome,omitempty"`
	RecebedorCnpj                 *string         `json:"recebedorCnpj,omitempty"`
	RecebedorCpf                  *string         `json:"recebedorCpf,omitempty"`
	RecebedorLogradouro           *string         `json:"recebedorLogradouro,omitempty"`
	RecebedorCidade               *string         `json:"recebedorCidade,omitempty"`
	RecebedorUF                   *string         `json:"recebedorUF,omitempty"`
	RecebedorCep                  *string         `json:"recebedorCep,omitempty"`
	FraudIdentified               *bool           `json:"fraudIdentified,omitempty"`
	ModalidadeAlteracao           *int32          `json:"modalidadeAlteracao,omitempty"`
	ModalidadeAlteracaoTrocoSaque *int32          `json:"modalidadeAlteracaoTrocoSaque,omitempty"`
}

// DecodeQRCodeV3Request represents QR code decode v3 request
// POST /pix/qrcodes/v3/query-processing
type DecodeQRCodeV3Request struct {
	CityCode            string `json:"cityCode"`                      // Required: IBGE code 7 digits
	IntendedPaymentDate *Date  `json:"intendedPaymentDate,omitempty"` // Optional
	QRCodeData          string `json:"qrCodeData"`                    // Required
	AccountID           int64  `json:"accountId"`                     // Required
}

// BillingDueDate represents billing due date information
type BillingDueDate struct {
	// Add fields based on actual API response structure
	Vencimento *string  `json:"vencimento,omitempty"`
	Valor      *float64 `json:"valor,omitempty"`
	// Add more fields as needed from API spec
}

// ImmediateBilling represents immediate billing information
type ImmediateBilling struct {
	// Add fields based on actual API response structure
	Valor *float64 `json:"valor,omitempty"`
	// Add more fields as needed from API spec
}

// StaticQrCodeInfo represents static QR code information
type StaticQrCodeInfo struct {
	// Add fields based on actual API response structure
	Identificador *string `json:"identificador,omitempty"`
	// Add more fields as needed from API spec
}

// RecurrenceInfo represents recurrence information
type RecurrenceInfo struct {
	// Add fields based on actual API response structure
	Tipo *string `json:"tipo,omitempty"`
	// Add more fields as needed from API spec
}

// KeyInfo represents key information
type KeyInfo struct {
	TipoChave          *string `json:"tipoChave,omitempty"`
	ChaveEnderecamento *string `json:"chaveEnderecamento,omitempty"`
	// Add more fields as needed from API spec
}

// DecodeQRCodeV3Response represents QR code decode v3 response
type DecodeQRCodeV3Response struct {
	Flow             FlowType          `json:"flow"`                       // STATIC, IMMEDIATE_CHARGE, CHARGE_WITH_DUE_DATE, AUT2, AUT3, AUT4
	BillingDueDate   *BillingDueDate   `json:"billingDueDate,omitempty"`   // Optional
	ImmediateBilling *ImmediateBilling `json:"immediateBilling,omitempty"` // Optional
	StaticQrCode     *StaticQrCodeInfo `json:"staticQrCode,omitempty"`     // Optional
	Recurrence       *RecurrenceInfo   `json:"recurrence,omitempty"`       // Optional
	KeyInfo          *KeyInfo          `json:"keyInfo,omitempty"`          // Optional
	Code             *string           `json:"code,omitempty"`             // Response code
	Message          *string           `json:"message,omitempty"`          // Response message
}

// Date represents a date in YYYY-MM-DD format
type Date struct {
	time.Time
}

// MarshalJSON implements json.Marshaler for Date
func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format("2006-01-02"))), nil
}

// UnmarshalJSON implements json.Unmarshaler for Date
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	str := string(data)
	if len(str) > 2 {
		str = str[1 : len(str)-1]
	}
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// Backward compatibility type aliases
type (
	// PayerInfoRequest is deprecated, use DynamicQRCodeRequest fields directly
	PayerInfoRequest struct {
		Name     *string `json:"name,omitempty"`
		Document *string `json:"document,omitempty"`
		City     *string `json:"city,omitempty"`
	}

	// QRCodeQueryRequest is deprecated, use QRCodeQueryProcessingRequest
	QRCodeQueryRequest = QRCodeQueryProcessingRequest

	// QRCodeQueryResponse is deprecated, use QRCodeQueryProcessingResponse
	QRCodeQueryResponse = QRCodeQueryProcessingResponse
)
