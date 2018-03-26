package conversationapi

type LineItemUpdate struct {
	OrderState *OrderState            `json:"orderState,omitempty"`
	Price      *Price                 `json:"price,omitempty"`
	Reason     string                 `json:"reason"`
	Extension  map[string]interface{} `json:"extension"`
}
