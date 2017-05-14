package definitions

type ExtensionServiceFeatureInfo struct {
	Enabled     bool   `json:"enabled,omitempty"`
	FeatureName string `json:"featureName,omitempty"`
	Reason      string `json:"reason,omitempty"`
}
