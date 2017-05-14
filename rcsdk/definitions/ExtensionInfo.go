package definitions

type ExtensionInfo struct {
	Id               string           `json:"id,omitempty"`
	RegionalSettings RegionalSettings `json:"regionalSettings,omitempty"`
	ServiceFeatures  `json:"serviceFeatures,omitempty"`
	Uri              string      `json:"uri,omitempty"`
	PartnerId        string      `json:"partnerId,omitempty"`
	Contact          ContactInfo `json:"contact,omitempty"`
	Name             string      `json:"name,omitempty"`
	StatusInfo       StatusInfo  `json:"statusInfo,omitempty"`
	Type             string      `json:"type,omitempty"`
	SetupWizardState string      `json:"setupWizardState,omitempty"`
	Status           string      `json:"status,omitempty"`
	Departments      `json:"departments,omitempty"`
	ExtensionNumber  string               `json:"extensionNumber,omitempty"`
	Permissions      ExtensionPermissions `json:"permissions,omitempty"`
	ProfileImage     ProfileImageInfo     `json:"profileImage,omitempty"`
	References       `json:"references,omitempty"`
}
