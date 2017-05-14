package requests

type OauthAuthorizePostRequestBody struct {
	Response_type string `json:"response_type,omitempty"`
	Client_id     string `json:"client_id,omitempty"`
	Redirect_uri  string `json:"redirect_uri,omitempty"`
	State         string `json:"state,omitempty"`
}
