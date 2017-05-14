package definitions

type PersonalContactInfo struct {
	Email2          string             `json:"email2,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	MiddleName      string             `json:"middleName,omitempty"`
	Company         string             `json:"company,omitempty"`
	CompanyPhone    string             `json:"companyPhone,omitempty"`
	CarPhone        string             `json:"carPhone,omitempty"`
	CallbackPhone   string             `json:"callbackPhone,omitempty"`
	Email3          string             `json:"email3,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	Availability    string             `json:"availability,omitempty"`
	NickName        string             `json:"nickName,omitempty"`
	HomePhone2      string             `json:"homePhone2,omitempty"`
	BusinessFax     string             `json:"businessFax,omitempty"`
	OtherPhone      string             `json:"otherPhone,omitempty"`
	OtherAddress    ContactAddressInfo `json:"otherAddress,omitempty"`
	Notes           string             `json:"notes,omitempty"`
	JobTitle        string             `json:"jobTitle,omitempty"`
	Id              string             `json:"id,omitempty"`
	Email           string             `json:"email,omitempty"`
	HomeAddress     ContactAddressInfo `json:"homeAddress,omitempty"`
	WebPage         string             `json:"webPage,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	HomePhone       string             `json:"homePhone,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	BusinessPhone2  string             `json:"businessPhone2,omitempty"`
	AssistantPhone  string             `json:"assistantPhone,omitempty"`
	OtherFax        string             `json:"otherFax,omitempty"`
	Birthday        string             `json:"birthday,omitempty"`
	Url             string             `json:"url,omitempty"`
	MobilePhone     string             `json:"mobilePhone,omitempty"`
}
