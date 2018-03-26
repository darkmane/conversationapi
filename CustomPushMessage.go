package conversationapi

type CustomPushMessage struct {
	Target           *Target           `json:"target,omitempty"`
	OrderUpdate      *OrderUpdate      `json:"orderUpdate,omitempty"`
	UserNotification *UserNotification `json:"userNotification,omitempty"`
}
