package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionConferencingPutRequestBody struct {
	AllowJoinBeforeHost bool                                           `json:"allowJoinBeforeHost,omitempty"`
	PhoneNumbers        []definitions.Conferencing_Request_PhoneNumber `json:"phoneNumbers,omitempty"`
}
