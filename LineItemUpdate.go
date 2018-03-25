package conversationapi

type LineItemUpdate struct {
	OrderState *OrderState            `json:"orderState"`
	Price      *Price                 `json:"price"`
	Reason     string                 `json:"reason"`
	Extension  map[string]interface{} `json:"extension"`
}
