package definitions

type ExtensionInfo_Request_ContactInfo struct {
	SetupWizardState string                                             `json:"setupWizardState,omitempty"`
	Department       string                                             `json:"department,omitempty"`
	Contact          ContactInfo                                        `json:"contact,omitempty"`
	RegionalSettings ExtensionInfo_Request_ContactInfo_RegionalSettings `json:"regionalSettings,omitempty"`
}
