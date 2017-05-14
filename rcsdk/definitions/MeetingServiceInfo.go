package definitions

type MeetingServiceInfo struct {
	IntlDialInNumbersUri string           `json:"intlDialInNumbersUri,omitempty"`
	ExternalUserInfo     ExternalUserInfo `json:"externalUserInfo,omitempty"`
	DialInNumbers        DialInNumbers    `json:"dialInNumbers,omitempty"`
	Uri                  string           `json:"uri,omitempty"`
	SupportUri           string           `json:"supportUri,omitempty"`
}
