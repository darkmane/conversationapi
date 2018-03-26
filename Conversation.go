package conversationapi

type Conversation struct {
	ConversationId    string            `json:"conversationId"`
	Type              *ConversationType `json:"type,omitempty"`
	ConversationToken string            `json:"conversationToken"`
}
