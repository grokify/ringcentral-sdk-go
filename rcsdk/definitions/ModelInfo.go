package definitions

type ModelInfo struct {
	Name   string      `json:"name,omitempty"`
	Addons []AddonInfo `json:"addons,omitempty"`
	Id     string      `json:"id,omitempty"`
}
