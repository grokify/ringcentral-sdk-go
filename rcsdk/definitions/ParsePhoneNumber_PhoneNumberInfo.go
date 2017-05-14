package definitions

type ParsePhoneNumber_PhoneNumberInfo struct {
	Normalized             string `json:"normalized,omitempty"`
	Dialable               string `json:"dialable,omitempty"`
	FormattedInternational string `json:"formattedInternational,omitempty"`
	Special                bool   `json:"special,omitempty"`
	FormattedNational      string `json:"formattedNational,omitempty"`
	OriginalString         string `json:"originalString,omitempty"`
	AreaCode               string `json:"areaCode,omitempty"`
	Country                `json:"country,omitempty"`
	E164                   string `json:"e164,omitempty"`
}
