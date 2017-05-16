package definitions

type TokenInfo struct {
	Access_token             string `json:"access_token,omitempty"`
	Endpoint_id              string `json:"endpoint_id,omitempty"`
	Expires_in               int    `json:"expires_in,omitempty"`
	Owner_id                 string `json:"owner_id,omitempty"`
	Refresh_token            string `json:"refresh_token,omitempty"`
	Refresh_token_expires_in int    `json:"refresh_token_expires_in,omitempty"`
	Scope                    string `json:"scope,omitempty"`
	Token_type               string `json:"token_type,omitempty"`
}
