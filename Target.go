package conversationapi

type Target struct {
	UserId   string   `json:"userId"`
	Intent   string   `json:"intent"`
	Argument Argument `json:"argument"`
}
