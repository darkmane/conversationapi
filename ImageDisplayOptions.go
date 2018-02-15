package conversationapi

import "strings"

type ImageDisplayOptions int

const (
	DEFAULT ImageDisplayOptions = iota
	WHITE
	CROPPED
)

func (ido *ImageDisplayOptions) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch {
	case str == "DEFAULT":
		*ido = DEFAULT
	case str == "WHITE":
		*ido = WHITE
	case str == "CROPPED":
		*ido = CROPPED
	default:
		*ido = DEFAULT
	}

	return nil
}
