package definitions

type AccountInfo struct {
	PartnerId        string              `json:"partnerId,omitempty"`
	SetupWizardState string              `json:"setupWizardState,omitempty"`
	Id               string              `json:"id,omitempty"`
	Uri              string              `json:"uri,omitempty"`
	Operator         ExtensionInfo       `json:"operator,omitempty"`
	StatusInfo       StatusInfo          `json:"statusInfo,omitempty"`
	MainNumber       string              `json:"mainNumber,omitempty"`
	ServiceInfo      Account_ServiceInfo `json:"serviceInfo,omitempty"`
	Status           string              `json:"status,omitempty"`
}
