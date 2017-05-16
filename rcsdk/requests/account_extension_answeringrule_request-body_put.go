package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionAnsweringrulePutRequestBody struct {
	Enabled    bool                       `json:"enabled,omitempty"`
	Forwarding definitions.ForwardingInfo `json:"forwarding,omitempty"`
	Greetings  []definitions.GreetingInfo `json:"greetings,omitempty"`
	Name       string                     `json:"name,omitempty"`
}
