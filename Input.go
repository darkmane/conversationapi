package conversationapi

type Input struct {
	RawInputs *[]RawInput `json:"rawInputs,omitempty"`
	Intent    string      `json:"intent"`
	Arguments *[]Argument `json:"arguments,omitempty"`
}
