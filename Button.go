package conversationapi

type Button struct {
	Title         string         `json:"title"`
	OpenUrlAction *OpenUrlAction `json:"openUrlAction"`
}
