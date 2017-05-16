package definitions

type MeetingServiceInfo struct {
	SupportUri           string           `json:"supportUri,omitempty"`
	IntlDialInNumbersUri string           `json:"intlDialInNumbersUri,omitempty"`
	ExternalUserInfo     ExternalUserInfo `json:"externalUserInfo,omitempty"`
	DialInNumbers        DialInNumbers    `json:"dialInNumbers,omitempty"`
	Uri                  string           `json:"uri,omitempty"`
}
