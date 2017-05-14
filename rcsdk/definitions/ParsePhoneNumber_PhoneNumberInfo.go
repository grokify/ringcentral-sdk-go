package definitions

type ParsePhoneNumber_PhoneNumberInfo struct {
	AreaCode               string `json:"areaCode,omitempty"`
	Country                `json:"country,omitempty"`
	Dialable               string `json:"dialable,omitempty"`
	E164                   string `json:"e164,omitempty"`
	FormattedNational      string `json:"formattedNational,omitempty"`
	OriginalString         string `json:"originalString,omitempty"`
	FormattedInternational string `json:"formattedInternational,omitempty"`
	Special                bool   `json:"special,omitempty"`
	Normalized             string `json:"normalized,omitempty"`
}
