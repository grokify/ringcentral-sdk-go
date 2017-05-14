package requests

type AccountExtensionCompanypagerPostRequestBody struct {
	From    CallerInfo   `json:"from,omitempty"`
	ReplyOn int          `json:"replyOn,omitempty"`
	Text    string       `json:"text,omitempty"`
	To      []CallerInfo `json:"to,omitempty"`
}
