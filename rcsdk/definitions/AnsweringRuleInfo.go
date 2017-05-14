package definitions

type AnsweringRuleInfo struct {
	Type                    string `json:"type,omitempty"`
	Name                    string `json:"name,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	CalledNumbers           `json:"calledNumbers,omitempty"`
	Callers                 `json:"callers,omitempty"`
	CallHandlingAction      string                      `json:"callHandlingAction,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
	Uri                     string                      `json:"uri,omitempty"`
	Id                      string                      `json:"id,omitempty"`
	Schedule                ScheduleInfo                `json:"schedule,omitempty"`
	Forwarding              ForwardingInfo              `json:"forwarding,omitempty"`
	Voicemail               VoicemailInfo               `json:"voicemail,omitempty"`
	Greetings               `json:"greetings,omitempty"`
}
