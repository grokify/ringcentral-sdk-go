package definitions

type ExtensionInfo struct {
	PartnerId        string                        `json:"partnerId,omitempty"`
	ProfileImage     ProfileImageInfo              `json:"profileImage,omitempty"`
	SetupWizardState string                        `json:"setupWizardState,omitempty"`
	Status           string                        `json:"status,omitempty"`
	StatusInfo       StatusInfo                    `json:"statusInfo,omitempty"`
	Uri              string                        `json:"uri,omitempty"`
	Departments      []DepartmentInfo              `json:"departments,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Permissions      ExtensionPermissions          `json:"permissions,omitempty"`
	References       []ReferenceInfo               `json:"references,omitempty"`
	RegionalSettings RegionalSettings              `json:"regionalSettings,omitempty"`
	Contact          ContactInfo                   `json:"contact,omitempty"`
	ExtensionNumber  string                        `json:"extensionNumber,omitempty"`
	ServiceFeatures  []ExtensionServiceFeatureInfo `json:"serviceFeatures,omitempty"`
	Id               string                        `json:"id,omitempty"`
	Type             string                        `json:"type,omitempty"`
}
