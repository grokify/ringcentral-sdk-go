package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountOrderPostRequestBody struct {
	Devices []definitions.DeviceInfo `json:"devices,omitempty"`
}
