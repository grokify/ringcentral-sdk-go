package definitions

type MessageEvent struct {
	ExtensionId int             `json:"extensionId,omitempty"`
	LastUpdated string          `json:"lastUpdated,omitempty"`
	Changes     []MessageChange `json:"changes,omitempty"`
}
