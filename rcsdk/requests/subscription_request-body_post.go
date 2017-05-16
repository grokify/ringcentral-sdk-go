package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type SubscriptionPostRequestBody struct {
	EventFilters []string                                      `json:"eventFilters,omitempty"`
	DeliveryMode definitions.Subscription_Request_DeliveryMode `json:"deliveryMode,omitempty"`
}
