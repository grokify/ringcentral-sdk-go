package definitions

type ParsePhoneNumber_PhoneNumberInfo struct {
	E164                   string                         `json:"e164,omitempty"`
	FormattedInternational string                         `json:"formattedInternational,omitempty"`
	FormattedNational      string                         `json:"formattedNational,omitempty"`
	Special                bool                           `json:"special,omitempty"`
	Normalized             string                         `json:"normalized,omitempty"`
	AreaCode               string                         `json:"areaCode,omitempty"`
	Country                []ParsePhoneNumber_CountryInfo `json:"country,omitempty"`
	Dialable               string                         `json:"dialable,omitempty"`
	OriginalString         string                         `json:"originalString,omitempty"`
}
