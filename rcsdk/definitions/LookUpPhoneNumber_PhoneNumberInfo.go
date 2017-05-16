package definitions

type LookUpPhoneNumber_PhoneNumberInfo struct {
	FormattedNumber string `json:"formattedNumber,omitempty"`
	VanityPattern   string `json:"vanityPattern,omitempty"`
	Rank            int    `json:"rank,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
}
