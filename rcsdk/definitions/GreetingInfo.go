package definitions

type GreetingInfo struct {
	Type   string     `json:"type,omitempty"`
	Preset PresetInfo `json:"preset,omitempty"`
}
