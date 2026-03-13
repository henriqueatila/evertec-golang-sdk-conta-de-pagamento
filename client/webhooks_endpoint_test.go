package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

func TestWebhooksEndpoints(t *testing.T) {
	runEndpointTests(t, []testEndpoint{
		{
			name: "UpdateSendGridWebhook", method: http.MethodPost, path: "/webhook/sendgrid/update",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateSendGridWebhook(ctx, &types.EventoEmailDTO{})
				return err
			},
		},
		{
			name: "NotifyArbiOperation", method: http.MethodPost, path: "/webhook/arbi/notifiesUserArbiOperation",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.NotifyArbiOperation(ctx, &types.NotificationPushRequest{})
				return err
			},
		},
		{
			name: "NotifyStatementClosed", method: http.MethodPost, path: "/postpaid/notification/eventhub/statement-closed/issuer1",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.NotifyStatementClosed(ctx, "issuer1", &types.EventHubRequest{})
				return err
			},
		},
		{
			name: "NotifyDueDate", method: http.MethodPost, path: "/postpaid/notification/eventhub/due-notification/issuer1",
			response: &types.GenericResponse{},
			call: func(ctx context.Context, c *Client) error {
				_, err := c.NotifyDueDate(ctx, "issuer1", &types.EventHubRequest{})
				return err
			},
		},
	})
}
