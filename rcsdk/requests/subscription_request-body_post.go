package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type SubscriptionPostRequestBody struct {
	DeliveryMode definitions.Subscription_Request_DeliveryMode `json:"deliveryMode,omitempty"`
	EventFilters []string                                      `json:"eventFilters,omitempty"`
}
