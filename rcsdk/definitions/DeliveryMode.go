package definitions

type DeliveryMode struct {
	EncryptionKey       string `json:"encryptionKey,omitempty"`
	TransportType       string `json:"transportType,omitempty"`
	Encryption          bool   `json:"encryption,omitempty"`
	Address             string `json:"address,omitempty"`
	SubscriberKey       string `json:"subscriberKey,omitempty"`
	SecretKey           string `json:"secretKey,omitempty"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
}
