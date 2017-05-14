package definitions

type ConferencingInfo_PhoneNumberInfo struct {
	Default     bool                                         `json:"default,omitempty"`
	HasGreeting bool                                         `json:"hasGreeting,omitempty"`
	Location    string                                       `json:"location,omitempty"`
	PhoneNumber string                                       `json:"phoneNumber,omitempty"`
	Country     ConferencingInfo_PhoneNumberInfo_CountryInfo `json:"country,omitempty"`
}
