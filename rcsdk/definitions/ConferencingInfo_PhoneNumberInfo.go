package definitions

type ConferencingInfo_PhoneNumberInfo struct {
	HasGreeting bool                                         `json:"hasGreeting,omitempty"`
	Location    string                                       `json:"location,omitempty"`
	PhoneNumber string                                       `json:"phoneNumber,omitempty"`
	Country     ConferencingInfo_PhoneNumberInfo_CountryInfo `json:"country,omitempty"`
	Default     bool                                         `json:"default,omitempty"`
}
