package requests

type AccountBusiness-addressPutRequestBody struct {
	Email	string	`json:"email,omitempty"`
	BusinessAddress	BusinessAddressInfo	`json:"businessAddress,omitempty"`
	Company	string	`json:"company,omitempty"`
}