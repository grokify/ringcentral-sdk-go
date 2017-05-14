package definitions

type LookUpPhoneNumber_PhoneNumberInfo struct {
	Rank            int    `json:"rank,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	FormattedNumber string `json:"formattedNumber,omitempty"`
	VanityPattern   string `json:"vanityPattern,omitempty"`
}
