package requests

type OauthAuthorizePostRequestBody struct {
	Client_id     string `json:"client_id,omitempty"`
	Redirect_uri  string `json:"redirect_uri,omitempty"`
	Response_type string `json:"response_type,omitempty"`
	State         string `json:"state,omitempty"`
}
