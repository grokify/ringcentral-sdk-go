package definitions

type ContactInfo struct {
	FirstName       string             `json:"firstName,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	Company         string             `json:"company,omitempty"`
	Email           string             `json:"email,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
}
