package conversationapi

type RawInput struct {
	InputType *InputType `json:"inputType,omitempty"`
	Query     string     `json:"query"`
}
