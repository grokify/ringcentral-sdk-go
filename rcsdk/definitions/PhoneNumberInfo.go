package definitions

type PhoneNumberInfo struct {
	Extension   PhoneNumberInfo_ExtensionInfo `json:"extension,omitempty"`
	PhoneNumber string                        `json:"phoneNumber,omitempty"`
	Type        string                        `json:"type,omitempty"`
	UsageType   string                        `json:"usageType,omitempty"`
	Id          string                        `json:"id,omitempty"`
	Country     CountryInfo                   `json:"country,omitempty"`
	Features    `json:"features,omitempty"`
	Location    string `json:"location,omitempty"`
	PaymentType string `json:"paymentType,omitempty"`
	Status      string `json:"status,omitempty"`
}
