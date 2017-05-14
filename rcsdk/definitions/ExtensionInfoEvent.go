package definitions

type ExtensionInfoEvent struct {
	ExtensionId string `json:"extensionId,omitempty"`
	EventType   string `json:"eventType,omitempty"`
}
