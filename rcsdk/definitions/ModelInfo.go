package definitions

type ModelInfo struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Addons `json:"addons,omitempty"`
}
