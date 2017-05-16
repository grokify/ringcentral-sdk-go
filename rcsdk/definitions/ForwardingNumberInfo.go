package definitions

type ForwardingNumberInfo struct {
	Features    string `json:"features,omitempty"`
	FlipNumber  int    `json:"flipNumber,omitempty"`
	Id          string `json:"id,omitempty"`
	Label       string `json:"label,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Uri         string `json:"uri,omitempty"`
}
