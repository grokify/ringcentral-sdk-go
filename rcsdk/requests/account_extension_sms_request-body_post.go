package requests

type AccountExtensionSmsPostRequestBody struct {
	From	CallerInfo	`json:"from,omitempty"`
	To	[]CallerInfo	`json:"to,omitempty"`
	Text	string	`json:"text,omitempty"`
}