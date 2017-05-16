package definitions

type DialInNumbers struct {
	Location        string                    `json:"location,omitempty"`
	Country         DialInNumbers_CountryInfo `json:"country,omitempty"`
	PhoneNumber     string                    `json:"phoneNumber,omitempty"`
	FormattedNumber string                    `json:"formattedNumber,omitempty"`
}
