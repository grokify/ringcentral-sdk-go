package definitions

type AnsweringRuleInfo struct {
	Callers                 `json:"callers,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
	Uri                     string                      `json:"uri,omitempty"`
	Id                      string                      `json:"id,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	Schedule                ScheduleInfo                `json:"schedule,omitempty"`
	Forwarding              ForwardingInfo              `json:"forwarding,omitempty"`
	Voicemail               VoicemailInfo               `json:"voicemail,omitempty"`
	Greetings               `json:"greetings,omitempty"`
	Type                    string `json:"type,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	CalledNumbers           `json:"calledNumbers,omitempty"`
	CallHandlingAction      string `json:"callHandlingAction,omitempty"`
}
