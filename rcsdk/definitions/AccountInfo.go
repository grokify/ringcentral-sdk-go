package definitions

type AccountInfo struct {
	ServiceInfo      Account_ServiceInfo `json:"serviceInfo,omitempty"`
	Status           string              `json:"status,omitempty"`
	StatusInfo       StatusInfo          `json:"statusInfo,omitempty"`
	PartnerId        string              `json:"partnerId,omitempty"`
	Uri              string              `json:"uri,omitempty"`
	MainNumber       string              `json:"mainNumber,omitempty"`
	Operator         ExtensionInfo       `json:"operator,omitempty"`
	SetupWizardState string              `json:"setupWizardState,omitempty"`
	Id               string              `json:"id,omitempty"`
}
