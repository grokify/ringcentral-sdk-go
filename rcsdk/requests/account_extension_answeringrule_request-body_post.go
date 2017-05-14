package requests

type AccountExtensionAnsweringrulePostRequestBody struct {
	Callers                 []CallersInfo               `json:"callers,omitempty"`
	CalledNumbers           []CalledNumberInfo          `json:"calledNumbers,omitempty"`
	CallHandlingAction      string                      `json:"callHandlingAction,omitempty"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
	VoiceMail               VoicemailInfo               `json:"voiceMail,omitempty"`
	Enabled                 bool                        `json:"enabled,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	Forwarding              ForwardingInfo              `json:"forwarding,omitempty"`
	Type                    string                      `json:"type,omitempty"`
	Schedule                AnsweringRule_ScheduleInfo  `json:"schedule,omitempty"`
}
