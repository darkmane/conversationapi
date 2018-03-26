package conversationapi

type Action struct {
	Type   *ActionType `json:"type,omitempty"`
	Button *Button     `json:"button,omitempty"`
}
