package conversationapi

import "strings"

type Price struct {
	Type   PriceType `json:"type"`
	Amount Money     `json:"amount"`
}

type PriceType int64

const (
	ESTIMATE PriceType = 1
	ACTUAL   PriceType = 2
)

func (pt *PriceType) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch {

	case str == "UNKNOWN":
		*pt = 0
	case str == "ESTIMATE":
		*pt = ESTIMATE
	case str == "ACTUAL":
		*pt = ACTUAL
	default:
		*pt = 0
	}

	return nil
}
