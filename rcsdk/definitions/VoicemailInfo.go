package definitions

type VoicemailInfo struct {
	Recipient RecipientInfo `json:"recipient,omitempty"`
	Enabled   bool          `json:"enabled,omitempty"`
}
