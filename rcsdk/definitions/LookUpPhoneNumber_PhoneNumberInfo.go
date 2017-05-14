package definitions

type LookUpPhoneNumber_PhoneNumberInfo struct {
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	FormattedNumber string `json:"formattedNumber,omitempty"`
	VanityPattern   string `json:"vanityPattern,omitempty"`
	Rank            int    `json:"rank,omitempty"`
}
