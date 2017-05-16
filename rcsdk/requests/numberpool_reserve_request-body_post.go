package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type NumberpoolReservePostRequestBody struct {
	Records []definitions.ReservePhoneNumber_Request_ReserveRecord `json:"records,omitempty"`
}
