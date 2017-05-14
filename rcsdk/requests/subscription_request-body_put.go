package requests

type SubscriptionPutRequestBody struct {
	EventFilters	[]string	`json:"eventFilters,omitempty"`
}