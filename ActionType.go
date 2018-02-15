package conversationapi

import "strings"

type ActionType int64

const (
	UNKNOWN ActionType = iota
	VIEW_DETAILS
	MODIFY
	CANCEL
	RETURN
	EXCHANGE
	EMAIL
	CALL
	REORDER
	REVIEW
	CUSTOMER_SERVICE
)

func (at *ActionType) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch {

	case str == "UNKNOWN":
		*at = UNKNOWN
	case str == "VIEW_DETAILS":
		*at = VIEW_DETAILS
	case str == "MODIFY":
		*at = MODIFY
	case str == "CANCEL":
		*at = CANCEL
	case str == "RETURN":
		*at = RETURN
	case str == "EXCHANGE":
		*at = EXCHANGE
	case str == "EMAIL":
		*at = EMAIL
	case str == "CALL":
		*at = CALL
	case str == "REORDER":
		*at = REORDER
	case str == "REVIEW":
		*at = REVIEW
	case str == "CUSTOMER_SERVICE":
		*at = CUSTOMER_SERVICE
	default:
		*at = UNKNOWN
	}

	return nil
}
