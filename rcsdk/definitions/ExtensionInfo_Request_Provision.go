package definitions

type ExtensionInfo_Request_Provision struct {
	Contact ExtensionInfo_Request_Provision_ContactInfo `json:"contact,omitempty"`
	Status  string                                      `json:"status,omitempty"`
}
