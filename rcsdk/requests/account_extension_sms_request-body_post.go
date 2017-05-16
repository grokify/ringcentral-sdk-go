package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionSmsPostRequestBody struct {
	From definitions.CallerInfo   `json:"from,omitempty"`
	Text string                   `json:"text,omitempty"`
	To   []definitions.CallerInfo `json:"to,omitempty"`
}
