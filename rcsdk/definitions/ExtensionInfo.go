package definitions

type ExtensionInfo struct {
	Contact          ContactInfo                   `json:"contact,omitempty"`
	Departments      []DepartmentInfo              `json:"departments,omitempty"`
	ExtensionNumber  string                        `json:"extensionNumber,omitempty"`
	ID               string                        `json:"id,omitempty"`
	Name             string                        `json:"name,omitempty"`
	PartnerID        string                        `json:"partnerId,omitempty"`
	Permissions      ExtensionPermissions          `json:"permissions,omitempty"`
	ProfileImage     ProfileImageInfo              `json:"profileImage,omitempty"`
	References       []ReferenceInfo               `json:"references,omitempty"`
	RegionalSettings RegionalSettings              `json:"regionalSettings,omitempty"`
	ServiceFeatures  []ExtensionServiceFeatureInfo `json:"serviceFeatures,omitempty"`
	SetupWizardState string                        `json:"setupWizardState,omitempty"`
	Status           string                        `json:"status,omitempty"`
	StatusInfo       StatusInfo                    `json:"statusInfo,omitempty"`
	Type             string                        `json:"type,omitempty"`
	URI              string                        `json:"uri,omitempty"`
}
