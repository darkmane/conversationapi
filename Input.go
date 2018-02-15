package conversationapi

type Input struct {
	RawInputs []RawInput `json:"rawInputs"`
	Intent    string     `json:"intent"`
	Arguments []Argument `json:"arguments"`
}
