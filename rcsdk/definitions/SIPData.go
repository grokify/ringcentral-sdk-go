package definitions

type SIPData struct {
	FromTag   string `json:"fromTag,omitempty"`
	LocalUri  string `json:"localUri,omitempty"`
	RemoteUri string `json:"remoteUri,omitempty"`
	ToTag     string `json:"toTag,omitempty"`
}
