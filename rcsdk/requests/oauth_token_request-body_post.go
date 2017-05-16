package requests

type OauthTokenPostRequestBody struct {
	Password          string `json:"password,omitempty"`
	Scope             string `json:"scope,omitempty"`
	Endpoint_id       string `json:"endpoint_id,omitempty"`
	Grant_type        string `json:"grant_type,omitempty"`
	Access_token_ttl  int    `json:"access_token_ttl,omitempty"`
	Refresh_token_ttl int    `json:"refresh_token_ttl,omitempty"`
	Username          string `json:"username,omitempty"`
	Extension         string `json:"extension,omitempty"`
}
