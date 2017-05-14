package definitions

type ReservePhoneNumber_Response_ReserveRecord struct {
	Status          string `json:"status,omitempty"`
	Error           string `json:"error,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	FormattedNumber string `json:"formattedNumber,omitempty"`
	ReservedTill    string `json:"reservedTill,omitempty"`
	ReservationId   string `json:"reservationId,omitempty"`
}
