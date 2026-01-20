# Evertec Conta de Pagamento SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento.svg)](https://pkg.go.dev/github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento)
[![CI](https://github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/actions/workflows/ci.yml/badge.svg)](https://github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/henriqueatila/evertec-golang-sdk-conta-de-pagamento/branch/main/graph/badge.svg)](https://codecov.io/gh/henriqueatila/evertec-golang-sdk-conta-de-pagamento)
[![Go Report Card](https://goreportcard.com/badge/github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento)](https://goreportcard.com/report/github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go SDK for the [Evertec Conta de Pagamento](https://paysmart-api.gitlab.io/conta-de-pagamentos/) payment account API.

## Features

- **287 API Methods** — Complete coverage of 15 API domains
- **mTLS Authentication** — Secure mutual TLS as required by the API
- **Webhook Handler** — 6 event types with idempotency support
- **Observability** — slog logging, metrics hooks, OpenTelemetry tracing
- **Type-Safe** — Typed request/response structs, 31 enum types
- **Typed Errors** — 13 error types for all HTTP status codes (400-503)

## Installation

```bash
go get github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento
```

Requires Go 1.21+

## Quick Start

Create the client with mTLS certificates:

```go
package main

import (
    "context"
    "log"

    "github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/client"
)

func main() {
    c, err := client.NewWithCertFiles(
        client.HomologBaseURL,
        "your-api-key",
        "certs/client.crt",
        "certs/client.key",
        "certs/ca.crt",
    )
    if err != nil {
        log.Fatal(err)
    }
    defer c.Close()

    ctx := context.Background()
    account, err := c.GetAccount(ctx, 12345)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Account: %+v", account)
}
```

## API Reference

### Core Domains
- **Accounts** — Account creation, balance, statements, status management
- **PIX Keys** — Create, delete, portability claims
- **PIX Transactions** — Payments, refunds, QR codes, receipts
- **PIX Automatico** — Recurring payment authorizations
- **Cards** — Physical/virtual card management, PIN, activation
- **Transfers** — Internal transfers, TED, DOC, batch operations
- **Bills** — Bill payments and scheduling
- **Bankslips** — Deposit bankslip generation (v1/v2)
- **MED** — Fraud reporting and refund solicitations per BACEN
- **Backoffice** — BIRO analysis, processors, branches, institutions
- Plus 5 more domains (Limits, Recipients, Travel, Profile, Visual Identity)

## Webhook Handler

Process asynchronous notifications:

```go
import (
    "log"
    "net/http"

    "github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/webhook"
)

func main() {
    handler := webhook.NewHandler(
        webhook.OnPixMovement(func(event *webhook.PixMovementEvent) error {
            log.Printf("[WEBHOOK] PIX %s: R$%.2f", event.MovementType, event.Value)
            return nil
        }),
        webhook.OnPrecautionaryBlock(func(event *webhook.PrecautionaryBlockEvent) error {
            log.Printf("[WEBHOOK] Block: %s", event.Type)
            return nil
        }),
    )

    http.HandleFunc("/webhook/pix-movement", handler.HandlePixMovement)
    http.HandleFunc("/webhook/precautionary-block", handler.HandlePrecautionaryBlock)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Supported events: PixMovement, ScheduledPix, PrecautionaryBlock, RetainedValue, AutomaticPix, ClaimNotification

## Error Handling

Handle API errors with type checking or sentinel errors:

```go
var validErr *client.ValidationError
if errors.As(err, &validErr) {
    for _, e := range validErr.Errors {
        log.Printf("Field %s: %s", e.Field, e.Message)
    }
}

var bizErr *client.BusinessRuleError
if errors.As(err, &bizErr) {
    log.Printf("Business Error [%s]: %s", bizErr.Code, bizErr.Message)
}

// Or use sentinel errors for quick checks:
if errors.Is(err, client.ErrNotFound) {
    // handle not found
}
if errors.Is(err, client.ErrInsufficientFunds) {
    // handle insufficient funds
}
```

## Observability

### Logging

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
c, _ := client.NewWithCertFiles(baseURL, apiKey, cert, key, ca,
    client.WithLogger(logger),
)
```

### OpenTelemetry

```go
c, _ := client.NewWithCertFiles(baseURL, apiKey, cert, key, ca,
    client.WithTracing(),
    client.WithMetrics(),
)
```

## Configuration

```go
// Environment
c, _ := client.NewWithCertFiles(
    client.HomologBaseURL,  // or client.ProductionBaseURL
    apiKey, cert, key, ca,
    client.WithTimeout(60*time.Second),
    client.WithAutoIdempotency(),
)

// Idempotency (safe retries)
ctx := client.WithIdempotencyKey(context.Background(), "uuid-v4-key")
resp, err := c.DoPixPayment(ctx, req)
```

## Testing

```bash
go test ./...              # All tests
go test -race ./...        # With race detection
go test -cover ./...       # With coverage
golangci-lint run          # Linting
```

## Security

- mTLS enforced on all requests
- TLS 1.2+ minimum
- X-API-KEY header authentication
- Idempotency keys for PIX operations
- Typed errors with BACEN compliance

## Project Statistics

| Metric | Value |
|--------|-------|
| Total Code | ~21K LOC |
| API Methods | 287 |
| Enum Types | 31 |
| Error Types | 13 |
| Dependencies | 1 (google/uuid) |
| Go Version | 1.21+ |

## License

MIT License - see [LICENSE](LICENSE) for details.
