package conversationapi

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
	Ssml         string `json:"ssml",omitempty`
	DisplayText  string `json:"displayText",omitempty`
}
