package definitions

type ExtensionInfo struct {
	Id               string           `json:"id,omitempty"`
	Uri              string           `json:"uri,omitempty"`
	Name             string           `json:"name,omitempty"`
	PartnerId        string           `json:"partnerId,omitempty"`
	RegionalSettings RegionalSettings `json:"regionalSettings,omitempty"`
	Contact          ContactInfo      `json:"contact,omitempty"`
	Status           string           `json:"status,omitempty"`
	Type             string           `json:"type,omitempty"`
	Departments      `json:"departments,omitempty"`
	Permissions      ExtensionPermissions `json:"permissions,omitempty"`
	ProfileImage     ProfileImageInfo     `json:"profileImage,omitempty"`
	References       `json:"references,omitempty"`
	ExtensionNumber  string `json:"extensionNumber,omitempty"`
	ServiceFeatures  `json:"serviceFeatures,omitempty"`
	SetupWizardState string     `json:"setupWizardState,omitempty"`
	StatusInfo       StatusInfo `json:"statusInfo,omitempty"`
}
