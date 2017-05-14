package definitions

type AccountInfo struct {
	Status           string              `json:"status,omitempty"`
	StatusInfo       StatusInfo          `json:"statusInfo,omitempty"`
	Uri              string              `json:"uri,omitempty"`
	MainNumber       string              `json:"mainNumber,omitempty"`
	Operator         ExtensionInfo       `json:"operator,omitempty"`
	SetupWizardState string              `json:"setupWizardState,omitempty"`
	Id               string              `json:"id,omitempty"`
	PartnerId        string              `json:"partnerId,omitempty"`
	ServiceInfo      Account_ServiceInfo `json:"serviceInfo,omitempty"`
}
