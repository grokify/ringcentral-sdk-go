package definitions

type ExtensionListEvent struct {
	EventType   string `json:"eventType,omitempty"`
	ExtensionId string `json:"extensionId,omitempty"`
}
