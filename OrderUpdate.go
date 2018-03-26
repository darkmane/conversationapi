package conversationapi

type OrderUpdate struct {
	GoogleOrderId          string                 `json:"googleOrderId"`
	ActionOrderId          string                 `json:"actionOrderId"`
	OrderState             *OrderState            `json:"orderState,omitempty"`
	OrderManagementActions *[]Action              `json:"orderManagementActions,omitempty"`
	Receipt                *Receipt               `json:"receipt,omitempty"`
	UpdateTime             string                 `json:"updateTime"`
	TotalPrice             *Price                 `json:"totalPrice,omitempty"`
	LineItemUpdates        *[]LineItemUpdate      `json:"lineItemUpdates,omitempty"`
	UserNotification       *UserNotification      `json:"userNotification,omitempty"`
	InfoExtension          map[string]interface{} `json:"infoExtension"`
}
