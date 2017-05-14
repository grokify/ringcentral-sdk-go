package requests

type AccountExtensionFaxPostRequestBody struct {
	CoverPageText	string	`json:"coverPageText,omitempty"`
	OriginalMessageId	string	`json:"originalMessageId,omitempty"`
	To	[]CallerInfo	`json:"to,omitempty"`
	FaxResolution	string	`json:"faxResolution,omitempty"`
	SendTime	string	`json:"sendTime,omitempty"`
	CoverIndex	int	`json:"coverIndex,omitempty"`
}