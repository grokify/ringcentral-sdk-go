package definitions

type MessageNotification struct {
	Uuid           string       `json:"uuid,omitempty"`
	Event          string       `json:"event,omitempty"`
	SubscriptionId string       `json:"subscriptionId,omitempty"`
	Timestamp      string       `json:"timestamp,omitempty"`
	Body           MessageEvent `json:"body,omitempty"`
}
