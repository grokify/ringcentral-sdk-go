package definitions

type PersonalContactInfo struct {
	AssistantPhone  string             `json:"assistantPhone,omitempty"`
	Availability    string             `json:"availability,omitempty"`
	Birthday        string             `json:"birthday,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	BusinessFax     string             `json:"businessFax,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	BusinessPhone2  string             `json:"businessPhone2,omitempty"`
	CallbackPhone   string             `json:"callbackPhone,omitempty"`
	CarPhone        string             `json:"carPhone,omitempty"`
	Company         string             `json:"company,omitempty"`
	CompanyPhone    string             `json:"companyPhone,omitempty"`
	Email           string             `json:"email,omitempty"`
	Email2          string             `json:"email2,omitempty"`
	Email3          string             `json:"email3,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	HomeAddress     ContactAddressInfo `json:"homeAddress,omitempty"`
	HomePhone       string             `json:"homePhone,omitempty"`
	HomePhone2      string             `json:"homePhone2,omitempty"`
	Id              string             `json:"id,omitempty"`
	JobTitle        string             `json:"jobTitle,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	MiddleName      string             `json:"middleName,omitempty"`
	MobilePhone     string             `json:"mobilePhone,omitempty"`
	NickName        string             `json:"nickName,omitempty"`
	Notes           string             `json:"notes,omitempty"`
	OtherAddress    ContactAddressInfo `json:"otherAddress,omitempty"`
	OtherFax        string             `json:"otherFax,omitempty"`
	OtherPhone      string             `json:"otherPhone,omitempty"`
	Url             string             `json:"url,omitempty"`
	WebPage         string             `json:"webPage,omitempty"`
}
