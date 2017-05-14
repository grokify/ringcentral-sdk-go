package definitions

type MeetingServiceInfo struct {
	ExternalUserInfo     ExternalUserInfo `json:"externalUserInfo,omitempty"`
	DialInNumbers        DialInNumbers    `json:"dialInNumbers,omitempty"`
	Uri                  string           `json:"uri,omitempty"`
	SupportUri           string           `json:"supportUri,omitempty"`
	IntlDialInNumbersUri string           `json:"intlDialInNumbersUri,omitempty"`
}
