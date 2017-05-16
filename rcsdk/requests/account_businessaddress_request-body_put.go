package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountBusinessaddressPutRequestBody struct {
	BusinessAddress definitions.BusinessAddressInfo `json:"businessAddress,omitempty"`
	Company         string                          `json:"company,omitempty"`
	Email           string                          `json:"email,omitempty"`
}
