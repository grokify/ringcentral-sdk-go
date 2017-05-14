package requests

type AccountExtensionAnswering-rulePostRequestBody struct {
	CallHandlingAction	string	`json:"callHandlingAction,omitempty"`
	VoiceMail	VoicemailInfo	`json:"voiceMail,omitempty"`
	Type	string	`json:"type,omitempty"`
	Name	string	`json:"name,omitempty"`
	CalledNumbers	[]CalledNumberInfo	`json:"calledNumbers,omitempty"`
	Forwarding	ForwardingInfo	`json:"forwarding,omitempty"`
	UnconditionalForwarding	UnconditionalForwardingInfo	`json:"unconditionalForwarding,omitempty"`
	Enabled	bool	`json:"enabled,omitempty"`
	Callers	[]CallersInfo	`json:"callers,omitempty"`
	Schedule	AnsweringRule_ScheduleInfo	`json:"schedule,omitempty"`
}