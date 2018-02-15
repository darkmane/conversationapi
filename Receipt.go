package conversationapi

type Receipt struct {
	ConfirmedActioOrderId string `json:"confirmedActioOrderId"`
	UserVisibileOrderId   string `json:"userVisibileOrderId"`
}
