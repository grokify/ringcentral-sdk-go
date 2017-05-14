package requests

type SubscriptionPostRequestBody struct {
	DeliveryMode Subscription_Request_DeliveryMode `json:"deliveryMode,omitempty"`
	EventFilters []string                          `json:"eventFilters,omitempty"`
}
