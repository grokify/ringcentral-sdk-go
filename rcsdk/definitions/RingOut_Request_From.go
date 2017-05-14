package definitions

type RingOut_Request_From struct {
	PhoneNumber        string `json:"phoneNumber,omitempty"`
	ForwardingNumberId string `json:"forwardingNumberId,omitempty"`
}
