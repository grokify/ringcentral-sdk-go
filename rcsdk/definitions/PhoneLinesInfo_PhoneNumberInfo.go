package definitions

type PhoneLinesInfo_PhoneNumberInfo struct {
	Status      string      `json:"status,omitempty"`
	Type        string      `json:"type,omitempty"`
	UsageType   string      `json:"usageType,omitempty"`
	Id          string      `json:"id,omitempty"`
	Country     CountryInfo `json:"country,omitempty"`
	Location    string      `json:"location,omitempty"`
	PaymentType string      `json:"paymentType,omitempty"`
	PhoneNumber string      `json:"phoneNumber,omitempty"`
}
