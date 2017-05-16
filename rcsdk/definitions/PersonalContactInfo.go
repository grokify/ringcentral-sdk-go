package definitions

type PersonalContactInfo struct {
	Availability    string             `json:"availability,omitempty"`
	HomePhone       string             `json:"homePhone,omitempty"`
	OtherPhone      string             `json:"otherPhone,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	MiddleName      string             `json:"middleName,omitempty"`
	Company         string             `json:"company,omitempty"`
	AssistantPhone  string             `json:"assistantPhone,omitempty"`
	JobTitle        string             `json:"jobTitle,omitempty"`
	BusinessPhone2  string             `json:"businessPhone2,omitempty"`
	CarPhone        string             `json:"carPhone,omitempty"`
	CallbackPhone   string             `json:"callbackPhone,omitempty"`
	Email3          string             `json:"email3,omitempty"`
	Url             string             `json:"url,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	OtherFax        string             `json:"otherFax,omitempty"`
	Birthday        string             `json:"birthday,omitempty"`
	Notes           string             `json:"notes,omitempty"`
	Id              string             `json:"id,omitempty"`
	NickName        string             `json:"nickName,omitempty"`
	HomePhone2      string             `json:"homePhone2,omitempty"`
	OtherAddress    ContactAddressInfo `json:"otherAddress,omitempty"`
	MobilePhone     string             `json:"mobilePhone,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	BusinessFax     string             `json:"businessFax,omitempty"`
	CompanyPhone    string             `json:"companyPhone,omitempty"`
	Email           string             `json:"email,omitempty"`
	Email2          string             `json:"email2,omitempty"`
	HomeAddress     ContactAddressInfo `json:"homeAddress,omitempty"`
	WebPage         string             `json:"webPage,omitempty"`
}
