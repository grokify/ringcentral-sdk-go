package requests

type AccountExtensionForwardingnumberPostRequestBody struct {
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Label       string `json:"label,omitempty"`
}
