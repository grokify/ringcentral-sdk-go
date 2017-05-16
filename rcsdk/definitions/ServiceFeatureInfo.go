package definitions

type ServiceFeatureInfo struct {
	Enabled     bool   `json:"enabled,omitempty"`
	FeatureName string `json:"featureName,omitempty"`
}
