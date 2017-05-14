package definitions

type DialInNumbers struct {
	PhoneNumber     string                    `json:"phoneNumber,omitempty"`
	FormattedNumber string                    `json:"formattedNumber,omitempty"`
	Location        string                    `json:"location,omitempty"`
	Country         DialInNumbers_CountryInfo `json:"country,omitempty"`
}
