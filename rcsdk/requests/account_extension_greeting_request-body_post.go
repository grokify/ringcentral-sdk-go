package requests

type AccountExtensionGreetingPostRequestBody struct {
	Type          string                               `json:"type,omitempty"`
	AnsweringRule CustomGreetingInfo_AnsweringRuleInfo `json:"answeringRule,omitempty"`
}
