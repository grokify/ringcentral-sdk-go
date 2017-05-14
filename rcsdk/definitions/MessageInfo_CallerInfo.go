package definitions

type MessageInfo_CallerInfo struct {
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	Location        string `json:"location,omitempty"`
	MessageStatus   string `json:"messageStatus,omitempty"`
	FaxErrorCode    string `json:"faxErrorCode,omitempty"`
	Name            string `json:"name,omitempty"`
}
