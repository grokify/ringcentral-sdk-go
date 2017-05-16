package definitions

type ReservePhoneNumber_Request_ReserveRecord struct {
	ReservedTill string `json:"reservedTill,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
}
