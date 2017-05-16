package definitions

type ParsePhoneNumber_PhoneNumberInfo struct {
	AreaCode               string                         `json:"areaCode,omitempty"`
	Country                []ParsePhoneNumber_CountryInfo `json:"country,omitempty"`
	Dialable               string                         `json:"dialable,omitempty"`
	E164                   string                         `json:"e164,omitempty"`
	FormattedInternational string                         `json:"formattedInternational,omitempty"`
	FormattedNational      string                         `json:"formattedNational,omitempty"`
	Normalized             string                         `json:"normalized,omitempty"`
	OriginalString         string                         `json:"originalString,omitempty"`
	Special                bool                           `json:"special,omitempty"`
}
