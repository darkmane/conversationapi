package conversationapi

type CustomPushMessage struct {
	Target           *Target           `json:"target"`
	OrderUpdate      *OrderUpdate      `json:"orderUpdate"`
	UserNotification *UserNotification `json:"userNotification"`
}
