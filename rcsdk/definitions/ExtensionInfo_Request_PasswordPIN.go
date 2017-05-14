package definitions

type ExtensionInfo_Request_PasswordPIN struct {
	IvrPin   string `json:"ivrPin,omitempty"`
	Password string `json:"password,omitempty"`
}
