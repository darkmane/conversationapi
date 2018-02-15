package conversationapi

type SimpleResponse struct {
	TextToSpeech string `json:"TextToSpeech"`
	Ssml         string `json:"ssml",omitempty`
	DisplayText  string `json:"displayText",omitempty`
}
