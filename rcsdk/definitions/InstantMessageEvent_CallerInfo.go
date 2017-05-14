package definitions

type InstantMessageEvent_CallerInfo struct {
	Name            string `json:"name,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	Location        string `json:"location,omitempty"`
}
