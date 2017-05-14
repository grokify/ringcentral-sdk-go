package definitions

type ForwardingNumberInfo struct {
	Features    string `json:"features,omitempty"`
	FlipNumber  int    `json:"flipNumber,omitempty"`
	Id          string `json:"id,omitempty"`
	Uri         string `json:"uri,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Label       string `json:"label,omitempty"`
}
