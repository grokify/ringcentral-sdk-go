package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionAnsweringrulePostRequestBody struct {
	Enabled                 bool                                    `json:"enabled,omitempty"`
	Type                    string                                  `json:"type,omitempty"`
	Name                    string                                  `json:"name,omitempty"`
	Callers                 []definitions.CallersInfo               `json:"callers,omitempty"`
	CallHandlingAction      string                                  `json:"callHandlingAction,omitempty"`
	Forwarding              definitions.ForwardingInfo              `json:"forwarding,omitempty"`
	VoiceMail               definitions.VoicemailInfo               `json:"voiceMail,omitempty"`
	CalledNumbers           []definitions.CalledNumberInfo          `json:"calledNumbers,omitempty"`
	Schedule                definitions.AnsweringRule_ScheduleInfo  `json:"schedule,omitempty"`
	UnconditionalForwarding definitions.UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
}
