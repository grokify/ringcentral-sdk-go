package definitions

type DetailedPresenceWithSIPNotification struct {
	SubscriptionId string                       `json:"subscriptionId,omitempty"`
	Timestamp      string                       `json:"timestamp,omitempty"`
	Body           DetailedPresencewithSIPEvent `json:"body,omitempty"`
	Uuid           string                       `json:"uuid,omitempty"`
	Event          string                       `json:"event,omitempty"`
}
