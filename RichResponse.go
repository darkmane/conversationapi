package conversationapi

type RichResponse struct {
	Items             *[]Item            `json:"items,omitempty"`
	Suggestions       *[]Suggestion      `json:"suggestions,omitempty"`
	LinkOutSuggestion *LinkOutSuggestion `json:"linkOutSuggestion,omitempty"`
}
