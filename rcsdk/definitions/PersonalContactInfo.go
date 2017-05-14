package definitions

type PersonalContactInfo struct {
	FirstName       string             `json:"firstName,omitempty"`
	NickName        string             `json:"nickName,omitempty"`
	Company         string             `json:"company,omitempty"`
	HomeAddress     ContactAddressInfo `json:"homeAddress,omitempty"`
	HomePhone       string             `json:"homePhone,omitempty"`
	AssistantPhone  string             `json:"assistantPhone,omitempty"`
	OtherPhone      string             `json:"otherPhone,omitempty"`
	OtherFax        string             `json:"otherFax,omitempty"`
	MiddleName      string             `json:"middleName,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	Id              string             `json:"id,omitempty"`
	BusinessPhone2  string             `json:"businessPhone2,omitempty"`
	MobilePhone     string             `json:"mobilePhone,omitempty"`
	CarPhone        string             `json:"carPhone,omitempty"`
	Email2          string             `json:"email2,omitempty"`
	Url             string             `json:"url,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	OtherAddress    ContactAddressInfo `json:"otherAddress,omitempty"`
	BusinessFax     string             `json:"businessFax,omitempty"`
	Email           string             `json:"email,omitempty"`
	Birthday        string             `json:"birthday,omitempty"`
	CompanyPhone    string             `json:"companyPhone,omitempty"`
	CallbackPhone   string             `json:"callbackPhone,omitempty"`
	WebPage         string             `json:"webPage,omitempty"`
	Notes           string             `json:"notes,omitempty"`
	Availability    string             `json:"availability,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	JobTitle        string             `json:"jobTitle,omitempty"`
	HomePhone2      string             `json:"homePhone2,omitempty"`
	Email3          string             `json:"email3,omitempty"`
}
