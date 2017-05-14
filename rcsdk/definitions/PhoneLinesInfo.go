package definitions

type PhoneLinesInfo struct {
	LineType  string                         `json:"lineType,omitempty"`
	PhoneInfo PhoneLinesInfo_PhoneNumberInfo `json:"phoneInfo,omitempty"`
}
