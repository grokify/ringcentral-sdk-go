package definitions

type MessageInfo_CallerInfo struct {
	Name            string `json:"name,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	Location        string `json:"location,omitempty"`
	MessageStatus   string `json:"messageStatus,omitempty"`
	FaxErrorCode    string `json:"faxErrorCode,omitempty"`
}
