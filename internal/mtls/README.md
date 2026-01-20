# mTLS Package

The `mtls` package provides utilities for loading and configuring mutual TLS (mTLS) certificates for the Evertec API client.

## Overview

Mutual TLS (mTLS) is a security mechanism where both the client and server authenticate each other using certificates. The Evertec API requires mTLS for secure communication.

## Features

- Load TLS configuration from certificate files
- Load TLS configuration from certificate bytes
- Support for custom CA certificates
- TLS 1.2+ enforcement
- Thread-safe configuration

## Usage

### Load from Files

```go
import "github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/internal/mtls"

// Load TLS config from files
tlsConfig, err := mtls.LoadTLSConfig(
	"client.crt",  // Client certificate
	"client.key",  // Private key
	"ca.crt",      // Optional CA certificate (empty string if not needed)
)
if err != nil {
	log.Fatal(err)
}

// Use with client
c, err := client.New(baseURL, apiKey, tlsConfig)
```

### Load from Bytes

```go
import "github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/internal/mtls"

// Read certificates from files or environment
certPEM, _ := os.ReadFile("client.crt")
keyPEM, _ := os.ReadFile("client.key")
caPEM, _ := os.ReadFile("ca.crt")

// Load TLS config from bytes
tlsConfig, err := mtls.LoadTLSConfigFromBytes(certPEM, keyPEM, caPEM)
if err != nil {
	log.Fatal(err)
}
```

### Without CA Certificate

If you don't have a custom CA certificate, you can omit it:

```go
// From files
tlsConfig, err := mtls.LoadTLSConfig("client.crt", "client.key", "")

// From bytes
tlsConfig, err := mtls.LoadTLSConfigFromBytes(certPEM, keyPEM, nil)
```

## Certificate Requirements

### Client Certificate

The client certificate must:
- Be in PEM format
- Include the full certificate chain (if applicable)
- Match the private key
- Be valid (not expired)
- Be issued by a CA trusted by Evertec

### Private Key

The private key must:
- Be in PEM format
- Match the client certificate
- Be unencrypted (no passphrase)
- Use RSA or ECDSA algorithm

### CA Certificate (Optional)

If provided, the CA certificate:
- Must be in PEM format
- Is used to verify the server certificate
- If omitted, system CA pool is used

## TLS Configuration

The package creates a `tls.Config` with:

```go
&tls.Config{
	Certificates: []tls.Certificate{cert},
	MinVersion:   tls.VersionTLS12,
	RootCAs:      caCertPool,  // If CA provided
}
```

### TLS Version

- Minimum version: TLS 1.2
- Recommended: TLS 1.3 (auto-negotiated)
- TLS 1.0 and 1.1 are not supported

## Obtaining Certificates

Contact Evertec to obtain your client certificates:

1. **Request Certificates**: Contact Evertec support
2. **Receive Credentials**: Get `client.crt`, `client.key`, and optionally `ca.crt`
3. **Store Securely**: Keep certificates secure (use secrets management)
4. **Set Permissions**: Limit file permissions (e.g., `chmod 600 client.key`)

## Certificate Storage

### Local Files

```bash
# Recommended permissions
chmod 600 client.key
chmod 644 client.crt
chmod 644 ca.crt
```

### Environment Variables

```bash
export EVERTEC_CERT=$(cat client.crt)
export EVERTEC_KEY=$(cat client.key)
export EVERTEC_CA=$(cat ca.crt)
```

```go
certPEM := []byte(os.Getenv("EVERTEC_CERT"))
keyPEM := []byte(os.Getenv("EVERTEC_KEY"))
caPEM := []byte(os.Getenv("EVERTEC_CA"))

tlsConfig, err := mtls.LoadTLSConfigFromBytes(certPEM, keyPEM, caPEM)
```

### Secrets Management

For production, use a secrets management service:

- AWS Secrets Manager
- Google Cloud Secret Manager
- Azure Key Vault
- HashiCorp Vault

## Generating Test Certificates

For testing purposes only (not for production):

```bash
# Generate private key
openssl genrsa -out client.key 2048

# Generate certificate signing request
openssl req -new -key client.key -out client.csr

# Generate self-signed certificate (1 year validity)
openssl x509 -req -days 365 -in client.csr -signkey client.key -out client.crt

# Clean up
rm client.csr
```

**Warning**: Self-signed certificates will not work with the Evertec production API. Use only for local testing.

## Error Handling

Common errors:

### Failed to Load Certificate

```
failed to load client certificate: tls: failed to find any PEM data in certificate input
```

**Solution**: Ensure certificate file is in PEM format

### Failed to Load Key

```
failed to load client certificate: tls: failed to find any PEM data in key input
```

**Solution**: Ensure private key file is in PEM format and unencrypted

### Certificate/Key Mismatch

```
failed to load client certificate: tls: private key does not match public key
```

**Solution**: Ensure certificate and key are a matching pair

### Failed to Append CA

```
failed to append CA certificate to pool
```

**Solution**: Ensure CA certificate is in valid PEM format

## Security Considerations

1. **Private Key Protection**
   - Never commit keys to version control
   - Use file permissions (e.g., `chmod 600`)
   - Use secrets management in production

2. **Certificate Rotation**
   - Monitor certificate expiration
   - Implement automated rotation
   - Have a renewal process

3. **Secure Storage**
   - Encrypt at rest
   - Use hardware security modules (HSM) for production
   - Implement access controls

4. **TLS Version**
   - Always use TLS 1.2+
   - Enable TLS 1.3 if supported
   - Never downgrade to TLS 1.0/1.1

## Testing

The package includes tests for error handling:

```bash
# Run tests
go test ./internal/mtls/...

# With coverage
go test ./internal/mtls/... -cover
```

**Note**: Tests requiring valid certificates are skipped by default.

## API Reference

### LoadTLSConfig

```go
func LoadTLSConfig(certFile, keyFile, caFile string) (*tls.Config, error)
```

Loads TLS configuration from certificate files.

**Parameters:**
- `certFile`: Path to client certificate file
- `keyFile`: Path to private key file
- `caFile`: Path to CA certificate file (empty string to use system CA pool)

**Returns:**
- `*tls.Config`: Configured TLS configuration
- `error`: Error if loading fails

### LoadTLSConfigFromBytes

```go
func LoadTLSConfigFromBytes(certPEM, keyPEM, caPEM []byte) (*tls.Config, error)
```

Loads TLS configuration from certificate bytes.

**Parameters:**
- `certPEM`: Client certificate in PEM format
- `keyPEM`: Private key in PEM format
- `caPEM`: CA certificate in PEM format (nil to use system CA pool)

**Returns:**
- `*tls.Config`: Configured TLS configuration
- `error`: Error if loading fails

## Examples

### Example 1: Basic Usage

```go
tlsConfig, err := mtls.LoadTLSConfig("client.crt", "client.key", "")
if err != nil {
	log.Fatal(err)
}

c, err := client.New(baseURL, apiKey, tlsConfig)
```

### Example 2: With CA Certificate

```go
tlsConfig, err := mtls.LoadTLSConfig("client.crt", "client.key", "ca.crt")
if err != nil {
	log.Fatal(err)
}
```

### Example 3: From Environment Variables

```go
certPEM := []byte(os.Getenv("EVERTEC_CERT"))
keyPEM := []byte(os.Getenv("EVERTEC_KEY"))

tlsConfig, err := mtls.LoadTLSConfigFromBytes(certPEM, keyPEM, nil)
```

### Example 4: With Error Handling

```go
tlsConfig, err := mtls.LoadTLSConfig(certFile, keyFile, caFile)
if err != nil {
	if strings.Contains(err.Error(), "no such file") {
		log.Fatal("Certificate files not found")
	}
	if strings.Contains(err.Error(), "failed to find any PEM data") {
		log.Fatal("Invalid PEM format")
	}
	log.Fatal(err)
}
```

## See Also

- [client](../../client/README.md) - Main client package
- [Root README](../../README.md) - SDK overview
- [TLS Best Practices](https://www.ssllabs.com/projects/best-practices/)
