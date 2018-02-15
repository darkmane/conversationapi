package conversationapi

type OrderUpdate struct {
	GoogleOrderId          string                 `json:"googleOrderId"`
	ActionOrderId          string                 `json:"actionOrderId"`
	OrderState             OrderState             `json:"orderState"`
	OrderManagementActions []Action               `json:"orderManagementActions"`
	Receipt                Receipt                `json:"receipt"`
	UpdateTime             string                 `json:"updateTime"`
	TotalPrice             Price                  `json:"totalPrice"`
	LineItemUpdates        []LineItemUpdate       `json:"lineItemUpdates"`
	UserNotification       UserNotification       `json:"userNotification"`
	InfoExtension          map[string]interface{} `json:"infoExtension"`
}
