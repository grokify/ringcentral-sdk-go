package definitions

type InstantMessageNotification struct {
	Timestamp      string              `json:"timestamp,omitempty"`
	Body           InstantMessageEvent `json:"body,omitempty"`
	Uuid           string              `json:"uuid,omitempty"`
	Event          string              `json:"event,omitempty"`
	SubscriptionId string              `json:"subscriptionId,omitempty"`
}
