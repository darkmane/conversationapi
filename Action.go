package conversationapi

type Action struct {
	Type   *ActionType `json:"type"`
	Button *Button     `json:"button"`
}
