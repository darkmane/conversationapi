package conversationapi

type InputPrompt struct {
	RichInitialPrompt *RichResponse     `json:"richInitialPrompt,omitempty"`
	NoInputPrompts    *[]SimpleResponse `json:"noInputPrompts,omitempty"`
}
