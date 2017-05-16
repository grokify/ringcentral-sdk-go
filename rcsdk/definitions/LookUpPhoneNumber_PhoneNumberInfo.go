package definitions

type LookUpPhoneNumber_PhoneNumberInfo struct {
	FormattedNumber string `json:"formattedNumber,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	Rank            int    `json:"rank,omitempty"`
	VanityPattern   string `json:"vanityPattern,omitempty"`
}
