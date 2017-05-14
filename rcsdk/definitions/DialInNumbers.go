package definitions

type DialInNumbers struct {
	Country         DialInNumbers_CountryInfo `json:"country,omitempty"`
	PhoneNumber     string                    `json:"phoneNumber,omitempty"`
	FormattedNumber string                    `json:"formattedNumber,omitempty"`
	Location        string                    `json:"location,omitempty"`
}
