package conversationapi

import "time"

type Argument struct {
	Name          string     `json:"name"`
	RawText       string     `json:"rawText"`
	TextValue     string     `json:"textValue"`
	BoolValue     bool       `json:"boolValue,omitempty"`
	DatetimeValue *time.Time `json:"datetimeValue,omitempty"`
	Extension     *Extension `json:"extension,omitempty"`
}
