package conversationapi

type InputPrompt struct {
	RichInitialPrompt *RichResponse
	NoInputPrompts    *[]SimpleResponse
}
