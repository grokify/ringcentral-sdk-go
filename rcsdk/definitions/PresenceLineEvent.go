package definitions

type PresenceLineEvent struct {
	Extension PresenceLineEvent_ExtensionInfo `json:"extension,omitempty"`
	Sequence  int                             `json:"sequence,omitempty"`
}
