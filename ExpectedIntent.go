package conversationapi

type ExpectedIntent struct {
	Intent         string                 `json:"intent,omitempty"`
	InputValueData map[string]interface{} `json:"inputValueData,omitempty"`
	ParameterName  string                 `json:"parameterName,omitempty"`
}
