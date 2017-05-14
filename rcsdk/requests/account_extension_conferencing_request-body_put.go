package requests

type AccountExtensionConferencingPutRequestBody struct {
	PhoneNumbers	[]Conferencing_Request_PhoneNumber	`json:"phoneNumbers,omitempty"`
	AllowJoinBeforeHost	bool	`json:"allowJoinBeforeHost,omitempty"`
}