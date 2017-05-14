package definitions

type PhoneNumberInfo struct {
	Id          string      `json:"id,omitempty"`
	Country     CountryInfo `json:"country,omitempty"`
	Features    `json:"features,omitempty"`
	Location    string                        `json:"location,omitempty"`
	Extension   PhoneNumberInfo_ExtensionInfo `json:"extension,omitempty"`
	PaymentType string                        `json:"paymentType,omitempty"`
	PhoneNumber string                        `json:"phoneNumber,omitempty"`
	Status      string                        `json:"status,omitempty"`
	Type        string                        `json:"type,omitempty"`
	UsageType   string                        `json:"usageType,omitempty"`
}
