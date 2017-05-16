package definitions

type ContactInfo struct {
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	Company         string             `json:"company,omitempty"`
	Email           string             `json:"email,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
}
