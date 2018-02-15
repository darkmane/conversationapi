package conversationapi

import "strings"

type Permission int

const (
	UNSPECIFIED_PERMISSION Permission = iota
	NAME
	DEVICE_PRECISE_LOCATION
	DEVICE_COARSE_LOCATION
	UPDATE
)

func (p *Permission) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch {
	case str == "UNSPECIFIED_PERMISSION":
		*p = UNSPECIFIED_PERMISSION

	case str == "NAME":
		*p = NAME
	case str == "DEVICE_PRECISE_LOCATION":
		*p = DEVICE_PRECISE_LOCATION
	case str == "DEVICE_COARSE_LOCATION":
		*p = DEVICE_COARSE_LOCATION
	case str == "UPDATE":
		*p = UPDATE
	}

	return nil
}
