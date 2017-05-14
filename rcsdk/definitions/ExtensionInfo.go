package definitions

type ExtensionInfo struct {
	ProfileImage     ProfileImageInfo `json:"profileImage,omitempty"`
	Name             string           `json:"name,omitempty"`
	PartnerId        string           `json:"partnerId,omitempty"`
	ServiceFeatures  `json:"serviceFeatures,omitempty"`
	Departments      `json:"departments,omitempty"`
	RegionalSettings RegionalSettings `json:"regionalSettings,omitempty"`
	SetupWizardState string           `json:"setupWizardState,omitempty"`
	StatusInfo       StatusInfo       `json:"statusInfo,omitempty"`
	Type             string           `json:"type,omitempty"`
	Uri              string           `json:"uri,omitempty"`
	References       `json:"references,omitempty"`
	ExtensionNumber  string               `json:"extensionNumber,omitempty"`
	Permissions      ExtensionPermissions `json:"permissions,omitempty"`
	Status           string               `json:"status,omitempty"`
	Id               string               `json:"id,omitempty"`
	Contact          ContactInfo          `json:"contact,omitempty"`
}
