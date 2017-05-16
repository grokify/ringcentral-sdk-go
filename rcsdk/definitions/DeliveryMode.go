package definitions

type DeliveryMode struct {
	Address             string `json:"address,omitempty"`
	Encryption          bool   `json:"encryption,omitempty"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
	EncryptionKey       string `json:"encryptionKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	SubscriberKey       string `json:"subscriberKey,omitempty"`
	TransportType       string `json:"transportType,omitempty"`
}
