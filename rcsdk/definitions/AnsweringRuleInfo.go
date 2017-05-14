package definitions

type AnsweringRuleInfo struct {
	CalledNumbers           `json:"calledNumbers,omitempty"`
	Callers                 `json:"callers,omitempty"`
	CallHandlingAction      string         `json:"callHandlingAction,omitempty"`
	Forwarding              ForwardingInfo `json:"forwarding,omitempty"`
	Greetings               `json:"greetings,omitempty"`
	Uri                     string                      `json:"uri,omitempty"`
	Type                    string                      `json:"type,omitempty"`
	Schedule                ScheduleInfo                `json:"schedule,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
	Voicemail               VoicemailInfo               `json:"voicemail,omitempty"`
	Id                      string                      `json:"id,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	Enabled                 bool                        `json:"enabled,omitempty"`
}
