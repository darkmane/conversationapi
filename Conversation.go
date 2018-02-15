package conversationapi

type Conversation struct {
	ConversationId    string           `json:"conversationId"`
	Type              ConversationType `json:"type"`
	ConversationToken string           `json:"conversationToken"`
}
