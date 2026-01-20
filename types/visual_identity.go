package types

// Visual Identity Operations (Email Configuration)
// Based on openapi-backoffice-identidade-visual.json

// EmailVisualIdentityRequest represents email visual identity creation request (API: CreateEmailVisualIdentityRequest)
type EmailVisualIdentityRequest struct {
	HeaderColor string `json:"headerColor"` // Required
	FooterColor string `json:"footerColor"` // Required
	LogoBase64  string `json:"logoBase64"`  // Required (data URI format: "data:image/png;base64,...")
}

// UpdateEmailVisualIdentityRequest represents email visual identity update request (API: UpdateEmailVisualIdentityRequest)
// All fields are optional in update
type UpdateEmailVisualIdentityRequest struct {
	HeaderColor *string `json:"headerColor,omitempty"`
	FooterColor *string `json:"footerColor,omitempty"`
	LogoBase64  *string `json:"logoBase64,omitempty"`
}

// EmailVisualIdentityResponse represents email visual identity response (API: GetEmailVisualIdentityResponse)
type EmailVisualIdentityResponse struct {
	HeaderColor              string `json:"headerColor"`              // Required
	FooterColor              string `json:"footerColor"`              // Required
	LogoEmailVisualIdentity  string `json:"logoEmailVisualIdentity"`  // Required
}
