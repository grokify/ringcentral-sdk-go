package definitions

type AnsweringRuleInfo struct {
	Uri                     string                         `json:"uri,omitempty"`
	Type                    string                         `json:"type,omitempty"`
	CalledNumbers           []AnsweringRuleInfo_CalleeInfo `json:"calledNumbers,omitempty"`
	Forwarding              ForwardingInfo                 `json:"forwarding,omitempty"`
	Voicemail               VoicemailInfo                  `json:"voicemail,omitempty"`
	Greetings               []GreetingInfo                 `json:"greetings,omitempty"`
	Id                      string                         `json:"id,omitempty"`
	Name                    string                         `json:"name,omitempty"`
	Enabled                 bool                           `json:"enabled,omitempty"`
	Schedule                ScheduleInfo                   `json:"schedule,omitempty"`
	Callers                 []AnsweringRuleInfo_CallerInfo `json:"callers,omitempty"`
	CallHandlingAction      string                         `json:"callHandlingAction,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo    `json:"unconditionalForwarding,omitempty"`
}
