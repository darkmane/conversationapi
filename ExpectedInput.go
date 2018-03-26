package conversationapi

type ExpectedInput struct {
	InputPrompt        *InputPrompt    `json:"inputPrompt,omitempty"`
	PossibleIntents    *ExpectedIntent `json:"possibleIntents,omitempty"`
	SpeechBiasingHints []string        `json:"speechBiasingHints,omitempty"`
}
