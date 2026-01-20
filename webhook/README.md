# Webhook Package

Webhook server for receiving PIX-focused event notifications from Evertec Conta de Pagamento API.

## Features

- **6 Event Types**: PIX movements, scheduled PIX, precautionary blocks, retained values, automatic PIX, claim notifications
- **Idempotent Processing**: Prevents duplicate event processing with pluggable storage backends
- **Thread-Safe**: Concurrent request handling with mutex-protected state
- **Graceful Shutdown**: Context-based cleanup with HTTP server shutdown
- **Hooks**: Pre/post processing for logging, metrics, custom logic
- **mTLS Support**: StartTLS for mutual TLS authentication
- **Background Cleanup**: Automatic expiration of old idempotency entries

## Quick Start

```go
package main

import (
    "context"
    "log"

    "github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/webhook"
)

// Implement custom handler
type MyHandler struct {
    webhook.BaseHandler // Embed for default no-ops
}

func (h *MyHandler) OnPIXMovement(ctx context.Context, event *webhook.PIXMovementEvent) error {
    log.Printf("PIX %s: R$%.2f", event.MovementType, float64(event.Value)/100)
    // Process event...
    return nil
}

func main() {
    handler := &MyHandler{}
    server := webhook.NewServer(handler)

    log.Fatal(server.Start(":8080"))
}
```

## Event Types

### 1. PIXMovementEvent
Triggered when PIX transaction occurs (sent or received).

**Endpoint**: `/webhook/movimento_pix`

```go
type PIXMovementEvent struct {
    AccountID       int64
    AuthorizationID string
    Payer           *Party
    Recipient       *Party
    Value           int64  // in cents
    EndToEnd        string
    MovementType    string // RECEIVED, SENT
}
```

### 2. ScheduledPIXExecutedEvent
Triggered when scheduled PIX is executed.

**Endpoint**: `/webhook/pix_agendado_executado`

```go
type ScheduledPIXExecutedEvent struct {
    AccountID     int64
    Success       bool
    Value         int64
    EndToEnd      string
    TransactionID int64
    Document      string
    ErrorMessage  string
}
```

### 3. PrecautionaryBlockEvent
Triggered when account/funds are blocked or unblocked.

**Endpoint**: `/webhook/notifica_bloqueio_cautelar`

```go
type PrecautionaryBlockEvent struct {
    Type                       string // BLOCK, UNBLOCK
    AccountID                  int64
    Value                      int64
    PrecautionaryTransactionID string
}
```

### 4. RetainedValueEvent
Triggered when funds are retained pending refund.

**Endpoint**: `/webhook/valor_retido`

```go
type RetainedValueEvent struct {
    AccountID           int64
    Value               int64
    OriginTransactionID int64
}
```

### 5. AutomaticPIXEvent
Triggered for automatic PIX subscription events.

**Endpoint**: `/webhook/notificacao_pix_automatico`

```go
type AutomaticPIXEvent struct {
    Type      string // See AutomaticPIXType enum
    AccountID int64
    PixKey    string
    Value     int64
    Document  string
}
```

**Types**:
- `ADESAO_PIX_AUTOMATICO` - Subscription
- `CANCELAMENTO_PIX_AUTOMATICO` - Cancellation
- `FLUXO_COMPLETO_PIX_AUTOMATICO` - Full flow
- `FALHA_AGENDAMENTO_PAGAMENTO_AUTOMATICO_SEM_RETENTATIVA` - Payment failure
- `CANCELAMENTO_COBRANCA_PIX_AUTOMATICO` - Billing cancellation

### 6. ClaimNotificationEvent
Triggered when PIX key claim is created or modified.

**Endpoint**: `/webhook/notifica_reivindicacao`

```go
type ClaimNotificationEvent struct {
    ClaimID     string
    ClaimType   string
    ClaimStatus string
    KeyType     string
    KeyValue    string
    AccountID   int64
}
```

## Configuration Options

### WithLogger
Custom logger for server events.

```go
server := webhook.NewServer(handler,
    webhook.WithLogger(slog.New(slog.NewJSONHandler(os.Stdout, nil))))
```

### WithIdempotencyTTL
Set time-to-live for idempotency entries (default: 24h).

```go
server := webhook.NewServer(handler,
    webhook.WithIdempotencyTTL(12 * time.Hour))
```

### WithIdempotencyStore
Inject custom storage backend (Redis, database, etc.).

```go
redisStore := NewRedisIdempotencyStore(redisClient)
server := webhook.NewServer(handler,
    webhook.WithIdempotencyStore(redisStore))
```

### WithHooks
Add pre/post processing hooks for logging, metrics.

```go
loggingHook := func(eventType string, event interface{}) error {
    log.Printf("Event: %s", eventType)
    return nil
}

server := webhook.NewServer(handler,
    webhook.WithHooks(loggingHook))
```

### WithMaxBodySize
Limit request body size (default: 1MB).

```go
server := webhook.NewServer(handler,
    webhook.WithMaxBodySize(2 * 1024 * 1024)) // 2MB
```

### EnableIdempotencyCleanup
Enable background cleanup of expired entries.

```go
server := webhook.NewServer(handler,
    webhook.EnableIdempotencyCleanup(1 * time.Hour))
```

## Custom Idempotency Store

Implement the `IdempotencyStore` interface for custom backends:

```go
type RedisIdempotencyStore struct {
    client *redis.Client
    ttl    time.Duration
}

func (s *RedisIdempotencyStore) IsProcessed(ctx context.Context, eventID string) (bool, error) {
    exists, err := s.client.Exists(ctx, eventID).Result()
    return exists > 0, err
}

func (s *RedisIdempotencyStore) MarkProcessed(ctx context.Context, eventID, eventType string) error {
    return s.client.Set(ctx, eventID, eventType, s.ttl).Err()
}

func (s *RedisIdempotencyStore) Cleanup(ctx context.Context) error {
    return nil // Redis handles expiration automatically
}
```

## Graceful Shutdown

```go
server := webhook.NewServer(handler)

// Start in goroutine
go func() {
    if err := server.Start(":8080"); err != nil {
        log.Fatal(err)
    }
}()

// Wait for signal...
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)
<-c

// Shutdown with timeout
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

if err := server.Shutdown(ctx); err != nil {
    log.Printf("Shutdown error: %v", err)
}
```

## mTLS Server

```go
server := webhook.NewServer(handler)
log.Fatal(server.StartTLS(":8443", "cert.pem", "key.pem"))
```

## Testing

Run tests:
```bash
go test ./webhook/ -v
```

With race detector:
```bash
go test ./webhook/ -race -v
```

Coverage:
```bash
go test ./webhook/ -cover
```

## Thread Safety

All operations are thread-safe:
- Idempotency store uses `sync.RWMutex`
- Concurrent request handling verified with 100-request test
- Race detector passes all tests

## Performance

- Target: < 100ms processing time per event
- Idempotency check: O(1) with in-memory store
- Background cleanup: Configurable interval, non-blocking
