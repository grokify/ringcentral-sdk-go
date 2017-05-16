package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionFaxPostRequestBody struct {
	CoverIndex        int                      `json:"coverIndex,omitempty"`
	CoverPageText     string                   `json:"coverPageText,omitempty"`
	FaxResolution     string                   `json:"faxResolution,omitempty"`
	OriginalMessageId string                   `json:"originalMessageId,omitempty"`
	SendTime          string                   `json:"sendTime,omitempty"`
	To                []definitions.CallerInfo `json:"to,omitempty"`
}
