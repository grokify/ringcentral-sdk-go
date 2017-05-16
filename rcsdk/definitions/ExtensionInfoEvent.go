package definitions

type ExtensionInfoEvent struct {
	EventType   string `json:"eventType,omitempty"`
	ExtensionId string `json:"extensionId,omitempty"`
}
