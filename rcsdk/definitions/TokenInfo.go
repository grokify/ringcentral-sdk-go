package definitions

type TokenInfo struct {
	Token_type               string `json:"token_type,omitempty"`
	Owner_id                 string `json:"owner_id,omitempty"`
	Endpoint_id              string `json:"endpoint_id,omitempty"`
	Access_token             string `json:"access_token,omitempty"`
	Expires_in               int    `json:"expires_in,omitempty"`
	Refresh_token            string `json:"refresh_token,omitempty"`
	Refresh_token_expires_in int    `json:"refresh_token_expires_in,omitempty"`
	Scope                    string `json:"scope,omitempty"`
}
