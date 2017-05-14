package requests

type SubscriptionPostRequestBody struct {
	EventFilters	[]string	`json:"eventFilters,omitempty"`
	DeliveryMode	Subscription_Request_DeliveryMode	`json:"deliveryMode,omitempty"`
}