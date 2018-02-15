package conversationapi

import "strings"

type InputType int

const (
	UNSPECIFIED_INPUT_TYPE InputType = iota
	TOUCH
	VOICE
	KEYBOARD
)

func (it *InputType) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch {
	case str == "UNSPECIFIED_INPUT_TYPE":
		*it = UNSPECIFIED_INPUT_TYPE
	case str == "TOUCH":
		*it = TOUCH
	case str == "VOICE":
		*it = VOICE
	case str == "KEYBOARD":
		*it = KEYBOARD
	}

	return nil
}
