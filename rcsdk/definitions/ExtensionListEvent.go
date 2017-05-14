package definitions

type ExtensionListEvent struct {
	ExtensionId string `json:"extensionId,omitempty"`
	EventType   string `json:"eventType,omitempty"`
}
