package definitions

type MessageNotification struct {
	Body           MessageEvent `json:"body,omitempty"`
	Event          string       `json:"event,omitempty"`
	SubscriptionId string       `json:"subscriptionId,omitempty"`
	Timestamp      string       `json:"timestamp,omitempty"`
	Uuid           string       `json:"uuid,omitempty"`
}
