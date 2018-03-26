package conversationapi

type AppResponse struct {
	ConversationToken  string             `json:"conversationToken"`
	UserStorage        string             `json:"userStorage"`
	ResetUserStorage   bool               `json:"resetUserStorage"`
	ExpectUserResponse bool               `json:"expectUserResponse"`
	ExpectedInputs     *[]ExpectedInput   `json:"expectedInputs,omitempty"`
	FinalResponse      *FinalResponse     `json:"finalResponse,omitempty"`
	CustomPushMessage  *CustomPushMessage `json:"customPushMessage,omitempty"`
	IsInSandbox        bool               `json:"isInSandbox,omitempty"`
}
