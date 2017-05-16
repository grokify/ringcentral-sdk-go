package definitions

type DialInNumbers struct {
	Country         DialInNumbers_CountryInfo `json:"country,omitempty"`
	FormattedNumber string                    `json:"formattedNumber,omitempty"`
	Location        string                    `json:"location,omitempty"`
	PhoneNumber     string                    `json:"phoneNumber,omitempty"`
}
