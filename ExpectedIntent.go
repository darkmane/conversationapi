package conversationapi

type ExpectedIntent struct {
	Intent         string                 `json:"intent"`
	InputValueData map[string]interface{} `json:"inputValueData"`
	ParameterName  string                 `json:"parameterName",omitempty`
}
