package definitions

type ExtensionInfo_Request_ContactInfo struct {
	Contact          ContactInfo                                        `json:"contact,omitempty"`
	Department       string                                             `json:"department,omitempty"`
	RegionalSettings ExtensionInfo_Request_ContactInfo_RegionalSettings `json:"regionalSettings,omitempty"`
	SetupWizardState string                                             `json:"setupWizardState,omitempty"`
}
