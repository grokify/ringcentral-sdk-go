package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionCompanypagerPostRequestBody struct {
	From    definitions.CallerInfo   `json:"from,omitempty"`
	ReplyOn int                      `json:"replyOn,omitempty"`
	Text    string                   `json:"text,omitempty"`
	To      []definitions.CallerInfo `json:"to,omitempty"`
}
