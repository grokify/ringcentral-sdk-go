package requests

type AccountExtensionAnsweringrulePostRequestBody struct {
	Name                    string                      `json:"name,omitempty"`
	CalledNumbers           []CalledNumberInfo          `json:"calledNumbers,omitempty"`
	Schedule                AnsweringRule_ScheduleInfo  `json:"schedule,omitempty"`
	CallHandlingAction      string                      `json:"callHandlingAction,omitempty"`
	Forwarding              ForwardingInfo              `json:"forwarding,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
	VoiceMail               VoicemailInfo               `json:"voiceMail,omitempty"`
	Enabled                 bool                        `json:"enabled,omitempty"`
	Type                    string                      `json:"type,omitempty"`
	Callers                 []CallersInfo               `json:"callers,omitempty"`
}
