package conversationapi

import "strings"

type ConversationType int

const (
	TYPE_UNSPECIFIED ConversationType = iota
	NEW
	ACTIVE
)

func (ct *ConversationType) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch {
	case str == "TYPE_UNSPECIFIED":
		*ct = TYPE_UNSPECIFIED
	case str == "NEW":
		*ct = NEW
	case str == "ACTIVE":
		*ct = ACTIVE
	}

	return nil
}
