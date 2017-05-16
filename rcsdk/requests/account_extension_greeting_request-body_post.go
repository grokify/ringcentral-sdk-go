package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionGreetingPostRequestBody struct {
	AnsweringRule definitions.CustomGreetingInfo_AnsweringRuleInfo `json:"answeringRule,omitempty"`
	Type          string                                           `json:"type,omitempty"`
}
