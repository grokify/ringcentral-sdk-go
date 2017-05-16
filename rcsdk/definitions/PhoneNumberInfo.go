package definitions

type PhoneNumberInfo struct {
	Country     CountryInfo                   `json:"country,omitempty"`
	Extension   PhoneNumberInfo_ExtensionInfo `json:"extension,omitempty"`
	Features    []string                      `json:"features,omitempty"`
	Id          string                        `json:"id,omitempty"`
	Location    string                        `json:"location,omitempty"`
	PaymentType string                        `json:"paymentType,omitempty"`
	PhoneNumber string                        `json:"phoneNumber,omitempty"`
	Status      string                        `json:"status,omitempty"`
	Type        string                        `json:"type,omitempty"`
	UsageType   string                        `json:"usageType,omitempty"`
}
