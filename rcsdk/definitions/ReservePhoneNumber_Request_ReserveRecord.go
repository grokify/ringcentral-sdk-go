package definitions

type ReservePhoneNumber_Request_ReserveRecord struct {
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	ReservedTill string `json:"reservedTill,omitempty"`
}
