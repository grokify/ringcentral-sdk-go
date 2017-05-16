package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionGreetingPostRequestBody struct {
	Type          string                                           `json:"type,omitempty"`
	AnsweringRule definitions.CustomGreetingInfo_AnsweringRuleInfo `json:"answeringRule,omitempty"`
}
