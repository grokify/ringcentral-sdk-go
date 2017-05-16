package definitions

type ConferencingInfo_PhoneNumberInfo struct {
	Country     ConferencingInfo_PhoneNumberInfo_CountryInfo `json:"country,omitempty"`
	Default     bool                                         `json:"default,omitempty"`
	HasGreeting bool                                         `json:"hasGreeting,omitempty"`
	Location    string                                       `json:"location,omitempty"`
	PhoneNumber string                                       `json:"phoneNumber,omitempty"`
}
