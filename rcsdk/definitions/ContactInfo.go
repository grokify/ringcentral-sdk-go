package definitions

type ContactInfo struct {
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	Company         string             `json:"company,omitempty"`
	Email           string             `json:"email,omitempty"`
}
