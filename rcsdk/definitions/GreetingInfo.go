package definitions

type GreetingInfo struct {
	Preset PresetInfo `json:"preset,omitempty"`
	Type   string     `json:"type,omitempty"`
}
