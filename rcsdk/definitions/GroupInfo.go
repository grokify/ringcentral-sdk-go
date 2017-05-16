package definitions

type GroupInfo struct {
	ContactsCount int    `json:"contactsCount,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
	Id            string `json:"id,omitempty"`
	Notes         string `json:"notes,omitempty"`
	Uri           string `json:"uri,omitempty"`
}
