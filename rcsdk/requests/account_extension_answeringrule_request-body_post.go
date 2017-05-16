package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionAnsweringrulePostRequestBody struct {
	CallHandlingAction      string                                  `json:"callHandlingAction,omitempty"`
	CalledNumbers           []definitions.CalledNumberInfo          `json:"calledNumbers,omitempty"`
	Callers                 []definitions.CallersInfo               `json:"callers,omitempty"`
	Enabled                 bool                                    `json:"enabled,omitempty"`
	Forwarding              definitions.ForwardingInfo              `json:"forwarding,omitempty"`
	Name                    string                                  `json:"name,omitempty"`
	Schedule                definitions.AnsweringRule_ScheduleInfo  `json:"schedule,omitempty"`
	Type                    string                                  `json:"type,omitempty"`
	UnconditionalForwarding definitions.UnconditionalForwardingInfo `json:"unconditionalForwarding,omitempty"`
	VoiceMail               definitions.VoicemailInfo               `json:"voiceMail,omitempty"`
}
