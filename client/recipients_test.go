package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestRecipientsEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "GetRecipients", method: http.MethodGet, path: "/accounts/1/recipients",
			response: &types.GetRecipientsResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRecipients(ctx, 1)
				return err
			},
		},
		{
			name: "GetRecipient", method: http.MethodGet, path: "/accounts/1/recipients/100",
			response: &types.GetRecipientResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetRecipient(ctx, 1, 100)
				return err
			},
		},
		{
			name: "CreateRecipient", method: http.MethodPost, path: "/accounts/1/recipients",
			response: &types.GetRecipientResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CreateRecipient(ctx, 1, &types.CreateRecipientRequest{})
				return err
			},
		},
		{
			name: "UpdateRecipient", method: http.MethodPut, path: "/accounts/1/recipients",
			response: &types.GetRecipientResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateRecipient(ctx, 1, &types.UpdateRecipientRequest{})
				return err
			},
		},
		{
			name: "DeleteRecipient", method: http.MethodDelete, path: "/accounts/1/recipients/100",
			response: nil,
			call: func(ctx context.Context, c *Client) error {
				return c.DeleteRecipient(ctx, 1, 100)
			},
		},
		{
			name: "GetLastTransactionError", method: http.MethodGet, path: "/accounts/1/feedback/lastTransactionError",
			response: &types.FeedbackResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetLastTransactionError(ctx, 1)
				return err
			},
		},
		{
			name: "SendFeedback", method: http.MethodPost, path: "/accounts/feedback/send",
			response: &types.ContaDigitalResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SendFeedback(ctx, &types.FeedbackRequest{})
				return err
			},
		},
		{
			name: "SendStatementFeedback", method: http.MethodPost, path: "/accounts/feedback/statement/send",
			response: &types.ContaDigitalResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SendStatementFeedback(ctx, &types.FeedbackStatementRequest{})
				return err
			},
		},
	})
}
