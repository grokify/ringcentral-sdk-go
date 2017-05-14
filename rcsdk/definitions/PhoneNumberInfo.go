package definitions

type PhoneNumberInfo struct {
	Features    `json:"features,omitempty"`
	PhoneNumber string                        `json:"phoneNumber,omitempty"`
	UsageType   string                        `json:"usageType,omitempty"`
	Id          string                        `json:"id,omitempty"`
	Country     CountryInfo                   `json:"country,omitempty"`
	Extension   PhoneNumberInfo_ExtensionInfo `json:"extension,omitempty"`
	Location    string                        `json:"location,omitempty"`
	PaymentType string                        `json:"paymentType,omitempty"`
	Status      string                        `json:"status,omitempty"`
	Type        string                        `json:"type,omitempty"`
}
