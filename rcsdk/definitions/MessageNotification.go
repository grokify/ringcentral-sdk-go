package definitions

type MessageNotification struct {
	Body           MessageEvent `json:"body,omitempty"`
	Uuid           string       `json:"uuid,omitempty"`
	Event          string       `json:"event,omitempty"`
	SubscriptionId string       `json:"subscriptionId,omitempty"`
	Timestamp      string       `json:"timestamp,omitempty"`
}
