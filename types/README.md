# Types Package

This package contains all type definitions for the Evertec Paysmart Conta de Pagamentos SDK.

## Overview

- **Total:** 1,665 lines of code across 14 files
- **Convention:** JSON fields use `snake_case` (Autorizador API) or `camelCase` (other APIs)
- **Amounts:** All monetary values are `int64` representing cents
- **Optional Fields:** Use pointers with `omitempty` tag

## Files

| File | Purpose |
|------|---------|
| `common.go` | Generic response structures |
| `errors.go` | Error types (400/409/500/503) |
| `enums.go` | Status enums and constants |
| `account.go` | Account management types |
| `card.go` | Card operations (physical/virtual/postpaid) |
| `pix.go` | PIX operations (keys/claims/payments) |
| `transaction.go` | Transfers/payments/recipients |
| `authorization.go` | Autorizador API (purchases/chargebacks) |
| `deposit.go` | Deposit orders |
| `bankslip.go` | Bank slips (boletos) |
| `limit.go` | Limit management |
| `proposal.go` | Account proposals |
| `contact.go` | Contacts and credits |
| `enum_reference.go` | Reference data enums |

## Usage Examples

### Account Creation
```go
req := &types.ProposalAccountRequest{
    Name:      "João Silva",
    Email:     "joao@example.com",
    Document:  "12345678901",
    BirthDate: "1990-01-15",
    Phone:     "11987654321",
}
```

### PIX Payment
```go
req := &types.PixPaymentRequest{
    KeyType:  types.PixKeyTypeCPF,
    KeyValue: "12345678901",
    Amount:   10000, // R$ 100.00 in cents
}
```

### Card Activation
```go
req := &types.ActivateCardRequest{
    LastFourDigits: "1234",
    CVV:            "123",
}
```

### Internal Transfer
```go
req := &types.InternalTransferRequest{
    RecipientAccountID: 54321,
    TransferAmount:     50000, // R$ 500.00 in cents
}
```

## Error Handling

```go
// Validation errors (HTTP 400)
var validationErrs types.ValidationErrors

// Business errors (HTTP 409)
var businessErr types.BusinessError

// API errors (HTTP 500/503)
var apiErr types.APIError
```

## Field Naming

### Autorizador API (snake_case)
- `account_id`
- `authorization_id`
- `total_amount`
- `asset_type`

### Other APIs (camelCase in responses)
- `accountId`
- `transactionId`
- `authenticationCode`

## Amount Convention

All monetary values are `int64` in **cents**:
- R$ 1.00 = 100
- R$ 10.00 = 1000
- R$ 100.00 = 10000

## Date/Time Formats

- **Date:** `YYYY-MM-DD` (e.g., "2026-01-15")
- **DateTime:** `time.Time` (RFC3339)
- **Time:** `HH:MM` (e.g., "20:00")
