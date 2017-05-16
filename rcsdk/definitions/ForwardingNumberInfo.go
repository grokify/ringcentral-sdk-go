package definitions

type ForwardingNumberInfo struct {
	Id          string `json:"id,omitempty"`
	Uri         string `json:"uri,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Label       string `json:"label,omitempty"`
	Features    string `json:"features,omitempty"`
	FlipNumber  int    `json:"flipNumber,omitempty"`
}
