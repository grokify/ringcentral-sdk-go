package definitions

type AnsweringRuleInfo struct {
	CallHandlingAction      string                         `json:"callHandlingAction,omitempty"`
	CalledNumbers           []AnsweringRuleInfo_CalleeInfo `json:"calledNumbers,omitempty"`
	Callers                 []AnsweringRuleInfo_CallerInfo `json:"callers,omitempty"`
	Enabled                 bool                           `json:"enabled,omitempty"`
	Forwarding              ForwardingInfo                 `json:"forwarding,omitempty"`
	Greetings               []GreetingInfo                 `json:"greetings,omitempty"`
	Id                      string                         `json:"id,omitempty"`
	Name                    string                         `json:"name,omitempty"`
	Schedule                ScheduleInfo                   `json:"schedule,omitempty"`
	Type                    string                         `json:"type,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo    `json:"unconditionalForwarding,omitempty"`
	Uri                     string                         `json:"uri,omitempty"`
	Voicemail               VoicemailInfo                  `json:"voicemail,omitempty"`
}
