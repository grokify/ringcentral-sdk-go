package definitions

type ReservePhoneNumber_Response_ReserveRecord struct {
	FormattedNumber string `json:"formattedNumber,omitempty"`
	ReservedTill    string `json:"reservedTill,omitempty"`
	ReservationId   string `json:"reservationId,omitempty"`
	Status          string `json:"status,omitempty"`
	Error           string `json:"error,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
}
