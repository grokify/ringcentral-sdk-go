package definitions

type MeetingServiceInfo struct {
	DialInNumbers        DialInNumbers    `json:"dialInNumbers,omitempty"`
	ExternalUserInfo     ExternalUserInfo `json:"externalUserInfo,omitempty"`
	IntlDialInNumbersUri string           `json:"intlDialInNumbersUri,omitempty"`
	SupportUri           string           `json:"supportUri,omitempty"`
	Uri                  string           `json:"uri,omitempty"`
}
