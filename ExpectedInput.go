package conversationapi

type ExpectedInput struct {
	InputPrompt        *InputPrompt
	PossibleIntents    *ExpectedIntent
	SpeechBiasingHints []string
}
