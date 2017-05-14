package definitions

type ContactInfo struct {
	Email           string             `json:"email,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	Company         string             `json:"company,omitempty"`
}
