package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionFaxPostRequestBody struct {
	To                []definitions.CallerInfo `json:"to,omitempty"`
	FaxResolution     string                   `json:"faxResolution,omitempty"`
	SendTime          string                   `json:"sendTime,omitempty"`
	CoverIndex        int                      `json:"coverIndex,omitempty"`
	CoverPageText     string                   `json:"coverPageText,omitempty"`
	OriginalMessageId string                   `json:"originalMessageId,omitempty"`
}
