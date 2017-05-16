package definitions

type ModelInfo struct {
	Addons []AddonInfo `json:"addons,omitempty"`
	Id     string      `json:"id,omitempty"`
	Name   string      `json:"name,omitempty"`
}
