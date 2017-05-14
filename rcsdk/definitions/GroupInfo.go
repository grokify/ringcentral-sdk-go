package definitions

type GroupInfo struct {
	Id            string `json:"id,omitempty"`
	Uri           string `json:"uri,omitempty"`
	ContactsCount int    `json:"contactsCount,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	Notes         string `json:"notes,omitempty"`
}
