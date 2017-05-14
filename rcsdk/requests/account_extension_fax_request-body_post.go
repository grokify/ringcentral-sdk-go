package requests

type AccountExtensionFaxPostRequestBody struct {
	FaxResolution     string       `json:"faxResolution,omitempty"`
	SendTime          string       `json:"sendTime,omitempty"`
	CoverIndex        int          `json:"coverIndex,omitempty"`
	CoverPageText     string       `json:"coverPageText,omitempty"`
	OriginalMessageId string       `json:"originalMessageId,omitempty"`
	To                []CallerInfo `json:"to,omitempty"`
}
