package conversationapi

type RichResponse struct {
	Items             *[]Item
	Suggestions       *[]Suggestion
	LinkOutSuggestion *LinkOutSuggestion
}
