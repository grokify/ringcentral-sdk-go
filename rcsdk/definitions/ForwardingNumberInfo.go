package definitions

type ForwardingNumberInfo struct {
	FlipNumber  int    `json:"flipNumber,omitempty"`
	Id          string `json:"id,omitempty"`
	Uri         string `json:"uri,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Label       string `json:"label,omitempty"`
	Features    string `json:"features,omitempty"`
}
