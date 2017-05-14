package requests

type AccountExtensionRingoutPostRequestBody struct {
	To         RingOut_Request_To          `json:"to,omitempty"`
	CallerId   RingOut_Request_To          `json:"callerId,omitempty"`
	PlayPrompt bool                        `json:"playPrompt,omitempty"`
	Country    RingOut_Request_CountryInfo `json:"country,omitempty"`
	From       RingOut_Request_From        `json:"from,omitempty"`
}
