package client

import (
	"context"
	"fmt"

	"github.com/henriqueatila/evertec-golang-sdk-conta-de-pagamento/types"
)

// Webhook Operations
// Based on OpenAPI spec: EventoEmailDTO, NotificationPushRequest, EventHubRequest

// UpdateSendGridWebhook updates SendGrid webhook
// Uses EventoEmailDTO from API spec
func (c *Client) UpdateSendGridWebhook(ctx context.Context, req *types.EventoEmailDTO) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.post(ctx, "/webhook/sendgrid/update", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// NotifyArbiOperation notifies Arbi user operation
// Uses NotificationPushRequest from API spec
func (c *Client) NotifyArbiOperation(ctx context.Context, req *types.NotificationPushRequest) (*types.GenericResponse, error) {
	var response types.GenericResponse
	if err := c.post(ctx, "/webhook/arbi/notifiesUserArbiOperation", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Postpaid Notification Operations (EventHub)
// Based on OpenAPI spec: EventHubRequest

// NotifyStatementClosed notifies statement closed via EventHub
// Uses EventHubRequest from API spec
func (c *Client) NotifyStatementClosed(ctx context.Context, issuerName string, req *types.EventHubRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/postpaid/notification/eventhub/statement-closed/%s", issuerName)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// NotifyDueDate notifies due date via EventHub
// Uses EventHubRequest from API spec
func (c *Client) NotifyDueDate(ctx context.Context, issuerName string, req *types.EventHubRequest) (*types.GenericResponse, error) {
	path := fmt.Sprintf("/postpaid/notification/eventhub/due-notification/%s", issuerName)
	var response types.GenericResponse
	if err := c.post(ctx, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
