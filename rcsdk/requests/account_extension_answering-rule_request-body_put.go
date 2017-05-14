package requests

type AccountExtensionAnswering-rulePutRequestBody struct {
	Enabled	bool	`json:"enabled,omitempty"`
	Name	string	`json:"name,omitempty"`
	Forwarding	ForwardingInfo	`json:"forwarding,omitempty"`
	Greetings	[]GreetingInfo	`json:"greetings,omitempty"`
}