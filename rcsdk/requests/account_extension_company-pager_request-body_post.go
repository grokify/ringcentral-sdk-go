package requests

type AccountExtensionCompany-pagerPostRequestBody struct {
	ReplyOn	int	`json:"replyOn,omitempty"`
	Text	string	`json:"text,omitempty"`
	To	[]CallerInfo	`json:"to,omitempty"`
	From	CallerInfo	`json:"from,omitempty"`
}