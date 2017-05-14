package definitions

type ExtensionInfo_Request_PasswordPIN struct {
	Password string `json:"password,omitempty"`
	IvrPin   string `json:"ivrPin,omitempty"`
}
