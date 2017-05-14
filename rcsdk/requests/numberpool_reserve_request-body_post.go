package requests

type NumberpoolReservePostRequestBody struct {
	Records []ReservePhoneNumber_Request_ReserveRecord `json:"records,omitempty"`
}
