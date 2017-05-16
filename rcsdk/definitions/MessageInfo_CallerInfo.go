package definitions

type MessageInfo_CallerInfo struct {
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	FaxErrorCode    string `json:"faxErrorCode,omitempty"`
	Location        string `json:"location,omitempty"`
	MessageStatus   string `json:"messageStatus,omitempty"`
	Name            string `json:"name,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
}
