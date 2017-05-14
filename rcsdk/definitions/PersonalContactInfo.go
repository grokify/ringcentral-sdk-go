package definitions

type PersonalContactInfo struct {
	JobTitle        string             `json:"jobTitle,omitempty"`
	OtherPhone      string             `json:"otherPhone,omitempty"`
	CallbackPhone   string             `json:"callbackPhone,omitempty"`
	Email3          string             `json:"email3,omitempty"`
	Notes           string             `json:"notes,omitempty"`
	Id              string             `json:"id,omitempty"`
	HomePhone       string             `json:"homePhone,omitempty"`
	HomeAddress     ContactAddressInfo `json:"homeAddress,omitempty"`
	MiddleName      string             `json:"middleName,omitempty"`
	BusinessPhone2  string             `json:"businessPhone2,omitempty"`
	MobilePhone     string             `json:"mobilePhone,omitempty"`
	BusinessAddress ContactAddressInfo `json:"businessAddress,omitempty"`
	OtherAddress    ContactAddressInfo `json:"otherAddress,omitempty"`
	WebPage         string             `json:"webPage,omitempty"`
	Company         string             `json:"company,omitempty"`
	Email2          string             `json:"email2,omitempty"`
	Birthday        string             `json:"birthday,omitempty"`
	Availability    string             `json:"availability,omitempty"`
	BusinessPhone   string             `json:"businessPhone,omitempty"`
	CompanyPhone    string             `json:"companyPhone,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	NickName        string             `json:"nickName,omitempty"`
	HomePhone2      string             `json:"homePhone2,omitempty"`
	Email           string             `json:"email,omitempty"`
	Url             string             `json:"url,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	BusinessFax     string             `json:"businessFax,omitempty"`
	AssistantPhone  string             `json:"assistantPhone,omitempty"`
	CarPhone        string             `json:"carPhone,omitempty"`
	OtherFax        string             `json:"otherFax,omitempty"`
}
