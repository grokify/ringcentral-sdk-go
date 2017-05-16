package definitions

type ReservePhoneNumber_Response_ReserveRecord struct {
	Error           string `json:"error,omitempty"`
	FormattedNumber string `json:"formattedNumber,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ReservationId   string `json:"reservationId,omitempty"`
	ReservedTill    string `json:"reservedTill,omitempty"`
	Status          string `json:"status,omitempty"`
}
