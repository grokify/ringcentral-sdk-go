package definitions

type ServiceFeatureInfo struct {
	FeatureName string `json:"featureName,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}
