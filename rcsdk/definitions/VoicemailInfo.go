package definitions

type VoicemailInfo struct {
	Enabled   bool          `json:"enabled,omitempty"`
	Recipient RecipientInfo `json:"recipient,omitempty"`
}
