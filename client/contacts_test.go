package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestContactsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "ListContacts", method: http.MethodGet, path: "/accounts/1/contact",
			response: &types.ContactListResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListContacts(ctx, 1)
				return err
			},
		},
		{
			name: "GetContactBankDetails", method: http.MethodGet, path: "/accounts/1/contact/100/bankdetails",
			response: &types.ContactBankDetailsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetContactBankDetails(ctx, 1, 100, "TED")
				return err
			},
		},
	})
}
