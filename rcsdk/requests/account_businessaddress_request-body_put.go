package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountBusinessaddressPutRequestBody struct {
	Company         string                          `json:"company,omitempty"`
	Email           string                          `json:"email,omitempty"`
	BusinessAddress definitions.BusinessAddressInfo `json:"businessAddress,omitempty"`
}
