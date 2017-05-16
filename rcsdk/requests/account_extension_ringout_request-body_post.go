package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionRingoutPostRequestBody struct {
	CallerId   definitions.RingOut_Request_To          `json:"callerId,omitempty"`
	Country    definitions.RingOut_Request_CountryInfo `json:"country,omitempty"`
	From       definitions.RingOut_Request_From        `json:"from,omitempty"`
	PlayPrompt bool                                    `json:"playPrompt,omitempty"`
	To         definitions.RingOut_Request_To          `json:"to,omitempty"`
}
