package definitions

type AccountInfo struct {
	ID               string              `json:"id,omitempty"`
	MainNumber       string              `json:"mainNumber,omitempty"`
	Operator         ExtensionInfo       `json:"operator,omitempty"`
	PartnerID        string              `json:"partnerId,omitempty"`
	ServiceInfo      Account_ServiceInfo `json:"serviceInfo,omitempty"`
	SetupWizardState string              `json:"setupWizardState,omitempty"`
	Status           string              `json:"status,omitempty"`
	StatusInfo       StatusInfo          `json:"statusInfo,omitempty"`
	Uri              string              `json:"uri,omitempty"`
}
