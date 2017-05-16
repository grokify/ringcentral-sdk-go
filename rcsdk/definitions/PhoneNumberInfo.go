package definitions

type PhoneNumberInfo struct {
	Country     CountryInfo                   `json:"country,omitempty"`
	Extension   PhoneNumberInfo_ExtensionInfo `json:"extension,omitempty"`
	Location    string                        `json:"location,omitempty"`
	PaymentType string                        `json:"paymentType,omitempty"`
	Status      string                        `json:"status,omitempty"`
	UsageType   string                        `json:"usageType,omitempty"`
	Id          string                        `json:"id,omitempty"`
	PhoneNumber string                        `json:"phoneNumber,omitempty"`
	Type        string                        `json:"type,omitempty"`
	Features    []string                      `json:"features,omitempty"`
}
