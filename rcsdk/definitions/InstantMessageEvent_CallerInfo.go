package definitions

type InstantMessageEvent_CallerInfo struct {
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	Location        string `json:"location,omitempty"`
	Name            string `json:"name,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
}
