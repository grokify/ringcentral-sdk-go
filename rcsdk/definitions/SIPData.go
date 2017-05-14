package definitions

type SIPData struct {
	ToTag     string `json:"toTag,omitempty"`
	FromTag   string `json:"fromTag,omitempty"`
	RemoteUri string `json:"remoteUri,omitempty"`
	LocalUri  string `json:"localUri,omitempty"`
}
