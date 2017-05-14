package definitions

type AccountInfo struct {
	Operator         ExtensionInfo       `json:"operator,omitempty"`
	PartnerId        string              `json:"partnerId,omitempty"`
	ServiceInfo      Account_ServiceInfo `json:"serviceInfo,omitempty"`
	SetupWizardState string              `json:"setupWizardState,omitempty"`
	Status           string              `json:"status,omitempty"`
	StatusInfo       StatusInfo          `json:"statusInfo,omitempty"`
	Uri              string              `json:"uri,omitempty"`
	MainNumber       string              `json:"mainNumber,omitempty"`
	Id               string              `json:"id,omitempty"`
}
