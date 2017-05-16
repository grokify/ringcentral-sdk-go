package definitions

type MessageEvent struct {
	Changes     []MessageChange `json:"changes,omitempty"`
	ExtensionId int             `json:"extensionId,omitempty"`
	LastUpdated string          `json:"lastUpdated,omitempty"`
}
