package conversationapi

type Money struct {
	CurrencyCode string `json:"currencyCode"`
	Units        int64  `json:"units"`
	Nanos        int64  `json:"nanos"`
}
