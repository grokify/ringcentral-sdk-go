package requests

type AccountBusinessaddressPutRequestBody struct {
	Company         string              `json:"company,omitempty"`
	Email           string              `json:"email,omitempty"`
	BusinessAddress BusinessAddressInfo `json:"businessAddress,omitempty"`
}
