package conversationapi

type RawInput struct {
	InputType InputType `json:"inputType"`
	Query     string    `json:"query"`
}
