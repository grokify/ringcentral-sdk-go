package definitions

type ExtensionInfo_Request_Provision struct {
	Status  string                                      `json:"status,omitempty"`
	Contact ExtensionInfo_Request_Provision_ContactInfo `json:"contact,omitempty"`
}
