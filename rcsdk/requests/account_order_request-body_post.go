package requests

type AccountOrderPostRequestBody struct {
	Devices	[]DeviceInfo	`json:"devices,omitempty"`
}