package requests

type OauthTokenPostRequestBody struct {
	Access_token_ttl  int    `json:"access_token_ttl,omitempty"`
	Endpoint_id       string `json:"endpoint_id,omitempty"`
	Extension         string `json:"extension,omitempty"`
	Grant_type        string `json:"grant_type,omitempty"`
	Password          string `json:"password,omitempty"`
	Refresh_token_ttl int    `json:"refresh_token_ttl,omitempty"`
	Scope             string `json:"scope,omitempty"`
	Username          string `json:"username,omitempty"`
}
